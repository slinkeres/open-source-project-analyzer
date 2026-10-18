package models

import "time"

type AiRepositoryAnalysis struct {
	Summary                   string   `json:"summary"`
	DocumentationQualityScore int      `json:"documentation_quality_score"`
	OnboardingScore           int      `json:"onboarding_score"`
	BeginnerFriendlinessScore int      `json:"beginner_friendliness_score"`
	Strengths                 []string `json:"strengths"`
	Weaknesses                []string `json:"weaknesses"`
	SuggestedImprovements     []string `json:"suggested_improvements"`
}

type RepositoryAiAnalysisResponse struct {
	RepositoryID              int       `json:"repository_id"`
	Summary                   string    `json:"summary"`
	DocumentationQualityScore int       `json:"documentation_quality_score"`
	OnboardingScore           int       `json:"onboarding_score"`
	BeginnerFriendlinessScore int       `json:"beginner_friendliness_score"`
	Strengths                 []string  `json:"strengths"`
	Weaknesses                []string  `json:"weaknesses"`
	SuggestedImprovements     []string  `json:"suggested_improvements"`
	Model                     string    `json:"model"`
	AnalyzedAt                time.Time `json:"analyzed_at"`
}
