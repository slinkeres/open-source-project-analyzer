package ai

import (
	"fmt"

	"github.com/slinkeres/open-source-project-analyzer/internal/models"
)

func buildPrompt(repository *models.StoredRepositoryResponse, readme string) string {
	description := "No description"
	if repository.Description != nil {
		description = *repository.Description
	}

	language := "Unknown"
	if repository.Language != nil {
		language = *repository.Language
	}

	return fmt.Sprintf(`Analyze this open-source GitHub repository.

Respond ONLY with a JSON object using EXACTLY these keys:
{
  "summary": "string, concise technical summary",
  "documentation_quality_score": number 0-100,
  "onboarding_score": number 0-100,
  "beginner_friendliness_score": number 0-100,
  "strengths": ["string", ...],
  "weaknesses": ["string", ...],
  "suggested_improvements": ["string", ...]
}

Do not invent features that are not present.

Repository: %s
Description: %s
Language: %s
Stars: %d
Forks: %d
Open issues: %d

README:

%s`, repository.FullName, description, language, repository.Stars, repository.Forks, repository.OpenIssues, readme)
}
