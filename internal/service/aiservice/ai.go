package aiservice

import (
	"context"

	"github.com/slinkeres/open-source-project-analyzer/internal/models"
	"github.com/slinkeres/open-source-project-analyzer/internal/repository/postgres"
	"github.com/slinkeres/open-source-project-analyzer/internal/service/ai"
)
 

type Service struct {
	projectRepo *postgres.ProjectRepository
	aiRepo *postgres.AiRepository
	aiClient *ai.Client
	model string 
}

func NewService(
	projectRepo *postgres.ProjectRepository,
	aiRepo *postgres.AiRepository,
	aiClient *ai.Client,
	model string,
) *Service{
	return &Service{
		projectRepo: projectRepo,
		aiRepo: aiRepo,
		aiClient: aiClient,
		model: model,
	}
}

func (s *Service) Analyze(ctx context.Context, repositoryID int) (*models.RepositoryAiAnalysisResponse, error){
	repo, err := s.projectRepo.GetByID(ctx, repositoryID)
	if err != nil {
		return nil, err
	}

	analysis, err := s.aiClient.AnalyzeRepository(ctx, repo)
	if err != nil {
		return nil, err
	}
	result := &models.RepositoryAiAnalysisResponse{
        RepositoryID:              repositoryID,
        Summary:                   analysis.Summary,
        DocumentationQualityScore: analysis.DocumentationQualityScore,
        OnboardingScore:           analysis.OnboardingScore,
        BeginnerFriendlinessScore: analysis.BeginnerFriendlinessScore,
        Strengths:                 analysis.Strengths,
        Weaknesses:                analysis.Weaknesses,
        SuggestedImprovements:     analysis.SuggestedImprovements,
        Model:                     s.model,
    }
    return s.aiRepo.Upsert(ctx, result)
}

func (s *Service) GetByID(ctx context.Context, repositoryID int) (*models.RepositoryAiAnalysisResponse, error){
	return s.aiRepo.GetByID(ctx, repositoryID)
}