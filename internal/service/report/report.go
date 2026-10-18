package report

import (
	"context"
	"math"

	"github.com/slinkeres/open-source-project-analyzer/internal/models"
	"github.com/slinkeres/open-source-project-analyzer/internal/repository/postgres"
)

type Service struct {
	projectRepo *postgres.ProjectRepository
	scoreRepo   *postgres.ScoreRepository
	aiRepo      *postgres.AiRepository
}

func NewService(
	projectRepo *postgres.ProjectRepository,
	scoreRepo *postgres.ScoreRepository,
	aiRepo *postgres.AiRepository,
) *Service {
	return &Service{
		projectRepo: projectRepo,
		scoreRepo:   scoreRepo,
		aiRepo:      aiRepo,
	}
}

func (s *Service) GetReport(ctx context.Context, repositoryId int) (*models.RepositoryReportResponse, error) {
	repo, err := s.projectRepo.GetByID(ctx, repositoryId)
	if err != nil {
		return nil, err
	}

	score, err := s.scoreRepo.GetByID(ctx, repositoryId)
	if err != nil {
		return nil, err
	}

	ai, err := s.aiRepo.GetByID(ctx, repositoryId)
	if err != nil {
		return nil, err
	}

	aiAverage := float64(ai.DocumentationQualityScore+ai.OnboardingScore+ai.BeginnerFriendlinessScore) / 3.0
	finalScore := int(math.Round((float64(score.OverallScore) + aiAverage) / 2.0))

	return &models.RepositoryReportResponse{
		Repository:         *repo,
		DeterministicScore: *score,
		AiAnalysis:         *ai,
		FinalScore:         finalScore,
	}, nil
}
