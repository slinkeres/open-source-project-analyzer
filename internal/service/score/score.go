package score

import (
	"context"

	"github.com/slinkeres/open-source-project-analyzer/internal/models"
	"github.com/slinkeres/open-source-project-analyzer/internal/repository/postgres"
	"github.com/slinkeres/open-source-project-analyzer/internal/service/scoring"
)

type Service struct {
	projectRepo *postgres.ProjectRepository
	scoreRepo *postgres.ScoreRepository
}

func NewService(projectRepo *postgres.ProjectRepository, scoreRepo *postgres.ScoreRepository) *Service{
	return &Service{
		projectRepo: projectRepo,
		scoreRepo: scoreRepo,
	}
}

func (s *Service) Calculate(ctx context.Context, repositoryID int) (*models.RepositoryScore, error){
	repo, err := s.projectRepo.GetByID(ctx, repositoryID)
	if err != nil{
		return nil, err
	}
	score := scoring.Calculate(repo)
	score.RepositoryID = repositoryID
	return s.scoreRepo.Upsert(ctx, score)
}

func (s *Service) GetByID(ctx context.Context, repositoryID int) (*models.RepositoryScore, error){
	return s.scoreRepo.GetByID(ctx, repositoryID)
}