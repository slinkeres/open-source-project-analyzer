package github

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	custom_errors "github.com/slinkeres/open-source-project-analyzer/internal/errors"
	"github.com/slinkeres/open-source-project-analyzer/internal/models"
)

type Client struct {
	httpClient *http.Client
	baseURL    string
	userAgent  string
}

func NewClient(httpClient *http.Client, baseUrl, userAgent string) *Client {
	return &Client{
		httpClient: httpClient,
		baseURL:    baseUrl,
		userAgent:  userAgent,
	}
}

func (c *Client) FetchRepository(ctx context.Context, owner, name string) (*models.GitHubRepositoryResponse, error) {
	endpoint := fmt.Sprintf("%s/repos/%s/%s", c.baseURL, owner, name)
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("User-Agent", c.userAgent)
	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("%w: failed to call GitHub API: %w", custom_errors.ErrExternal, err)
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusNotFound {
		return nil, custom_errors.ErrNotFound
	} else if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w:github returned status %d", custom_errors.ErrExternal, response.StatusCode)
	}
	var repo models.GitHubRepositoryResponse
	if err := json.NewDecoder(response.Body).Decode(&repo); err != nil {
		return nil, fmt.Errorf("%w: failed to decode GitHub response: %w", custom_errors.ErrExternal, err)
	}
	return &repo, nil
}

func (c *Client) FetchReadme(ctx context.Context, owner, name string) (string, error) {
	endpoint := fmt.Sprintf("%s/repos/%s/%s/readme", c.baseURL, owner, name)
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return "", err
	}
	request.Header.Set("User-Agent", c.userAgent)
	response, err := c.httpClient.Do(request)
	if err != nil {
		return "", fmt.Errorf("%w: failed to call GitHub API: %w", custom_errors.ErrExternal, err)
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusNotFound {
		return "", nil
	}
	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("%w:github returned status %d", custom_errors.ErrExternal, response.StatusCode)
	}
	var readmeMeta models.GitHubReadmeResponse
	if err := json.NewDecoder(response.Body).Decode(&readmeMeta); err != nil {
		return "", fmt.Errorf("%w: failed to decode readme metadata", custom_errors.ErrExternal)
	}
	if readmeMeta.DownloadURL == "" {
		return "", nil
	}
	readmeReq, err := http.NewRequestWithContext(ctx, http.MethodGet, readmeMeta.DownloadURL, nil)
	if err != nil {
		return "", err
	}
	readmeReq.Header.Set("User-Agent", c.userAgent)
	readmeResp, err := c.httpClient.Do(readmeReq)
	if err != nil {
		return "", fmt.Errorf("%w:failed to download readme, %w", custom_errors.ErrExternal, err)
	}
	defer readmeResp.Body.Close()
	if readmeResp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("%w:failed to download readme, status %d", custom_errors.ErrExternal, readmeResp.StatusCode)
	}

	body, err := io.ReadAll(readmeResp.Body)
	if err != nil {
		return "", fmt.Errorf("%w:failed to read readme body, %w", custom_errors.ErrExternal, err)
	}
	return string(body), nil
}

func (c *Client) ParseURL(rawURL string) (*models.RepositoryPreviewResponse, error) {
	rawURL = strings.TrimSpace(rawURL)
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return nil, fmt.Errorf("%w: invalid URL", custom_errors.ErrBadRequest)
	}
	if parsed.Hostname() != "github.com" {
		return nil, fmt.Errorf("%w: only GitHub URLs are supported", custom_errors.ErrBadRequest)
	}
	segments := []string{}
	for _, seg := range strings.Split(parsed.Path, "/") {
		if seg != "" {
			segments = append(segments, seg)
		}
	}
	if len(segments) < 2 {
		return nil, fmt.Errorf("%w: URL must contain owner and repository name", custom_errors.ErrBadRequest)
	}
	owner := segments[0]
	name := strings.TrimSuffix(segments[1], ".git")
	normalizedURL := fmt.Sprintf("https://github.com/%s/%s", owner, name)

	return &models.RepositoryPreviewResponse{
		Owner:         owner,
		Name:          name,
		NormalizedURL: normalizedURL,
	}, nil
}
