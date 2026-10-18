package ai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	custom_errors "github.com/slinkeres/open-source-project-analyzer/internal/errors"
	"github.com/slinkeres/open-source-project-analyzer/internal/models"
)

type groqMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type groqResponseFormat struct {
	Type string `json:"type"`
}

type groqRequest struct {
	Model          string             `json:"model"`
	Messages       []groqMessage      `json:"messages"`
	ResponseFormat groqResponseFormat `json:"response_format"`
	Temperature    float64            `json:"temperature"`
}

type groqResponse struct {
	Choices []groqChoice `json:"choices"`
}

type groqChoice struct {
	Message groqMessage `json:"message"`
}

type Client struct {
	httpClient *http.Client
	apiKey     string
	baseURL    string
	model      string
}

func NewClient(httpClient *http.Client, apiKey, baseURL, model string) *Client {
	return &Client{
		httpClient: httpClient,
		apiKey:     apiKey,
		baseURL:    baseURL,
		model:      model,
	}
}

func (c *Client) AnalyzeRepository(
	ctx context.Context,
	repository *models.StoredRepositoryResponse,
) (*models.AiRepositoryAnalysis, error) {
	readme := repository.Readme
	runes := []rune(readme)
	if len(runes) > 18000 {
		readme = string(runes[:18000])
	}

	prompt := buildPrompt(repository, readme)

	reqBody := groqRequest{
		Model: c.model,
		Messages: []groqMessage{
			{Role: "user", Content: prompt},
		},
		ResponseFormat: groqResponseFormat{Type: "json_object"},
		Temperature:    0.2,
	}

	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("%w: failed to marshal groq request: %w", custom_errors.ErrInternal, err)
	}

	endpoint := fmt.Sprintf("%s/chat/completions", c.baseURL)
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("%w: failed to build groq request: %w", custom_errors.ErrInternal, err)
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+c.apiKey)

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("%w: failed to call groq API: %w", custom_errors.ErrExternal, err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(response.Body)
		return nil, fmt.Errorf("%w: groq returned status %d: %s", custom_errors.ErrExternal,
			response.StatusCode, string(bodyBytes))
	}

	var groqResp groqResponse
	if err := json.NewDecoder(response.Body).Decode(&groqResp); err != nil {
		return nil, fmt.Errorf("%w: failed to decode groq response: %w", custom_errors.ErrExternal, err)
	}

	if len(groqResp.Choices) == 0 {
		return nil, fmt.Errorf("%w: groq returned empty choices", custom_errors.ErrExternal)
	}

	content := groqResp.Choices[0].Message.Content
	var analysis models.AiRepositoryAnalysis
	if err := json.Unmarshal([]byte(content), &analysis); err != nil {
		return nil, fmt.Errorf("%w: failed to parse model output as JSON: %w", custom_errors.ErrExternal, err)
	}

	return &analysis, nil

}
