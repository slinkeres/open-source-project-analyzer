package models

type RepositoryReportResponse struct {
	Repository         StoredRepositoryResponse     `json:"repository"`
	DeterministicScore RepositoryScore              `json:"deterministic_score"`
	AiAnalysis         RepositoryAiAnalysisResponse `json:"ai_analysis"`
	FinalScore         int                          `json:"final_score"`
}
