package repository

import (
	"context"

	"github.com/slinkeres/open-source-project-analyzer/internal/models"
	"github.com/slinkeres/open-source-project-analyzer/internal/repository/postgres"
	"github.com/slinkeres/open-source-project-analyzer/internal/service/github"
)

type Service struct {
	githubClient *github.Client
	projectRepo  *postgres.ProjectRepository
}

func NewService(githubClient *github.Client, projectRepo *postgres.ProjectRepository) *Service {
	return &Service{
		githubClient: githubClient,
		projectRepo:  projectRepo,
	}
}

func (s *Service) Preview(rawURL string) (*models.RepositoryPreviewResponse, error){
	return s.githubClient.ParseURL(rawURL)
}

func (s *Service) Fetch(ctx context.Context, rawURL string) (*models.StoredRepositoryResponse, error){
	preview, err := s.githubClient.ParseURL(rawURL)
	if err != nil{
		return nil, err
	}
	githubRepo, err := s.githubClient.FetchRepository(ctx, preview.Owner, preview.Name)
	if err != nil{
		return nil,err
	}
	readme, err := s.githubClient.FetchReadme(ctx, preview.Owner, preview.Name)
	if err != nil{
		return nil, err
	}
	var license *string
	if githubRepo.License != nil{
		license = &githubRepo.License.SpdxID
	}
	repository := &models.StoredRepositoryResponse{
		GithubID:        githubRepo.ID,
        Owner:           githubRepo.Owner.Login,
        Name:            githubRepo.Name,
        FullName:        githubRepo.FullName,
        Description:     githubRepo.Description,
        Stars:           githubRepo.StargazersCount,
        Forks:           githubRepo.ForksCount,
        OpenIssues:      githubRepo.OpenIssuesCount,
        Subscribers:     githubRepo.SubscribersCount,
        Language:        githubRepo.Language,
        License:         license,
        Archived:        githubRepo.Archived,
        HasIssues:       githubRepo.HasIssues,
        URL:             githubRepo.HTMLURL,
        DefaultBranch:   githubRepo.DefaultBranch,
        Topics:          githubRepo.Topics,
        GithubCreatedAt: githubRepo.CreatedAt,
        GithubUpdatedAt: githubRepo.UpdatedAt,
        PushedAt:        githubRepo.PushedAt,
        Readme:          readme,
	}
	id, err := s.projectRepo.Upsert(ctx, repository)
	if err != nil{
		return nil, err
	}

	repository.DatabaseID = id
	return repository, nil
}

func (s *Service) List(ctx context.Context) ([]models.RepositoryListItemResponse, error){
	return s.projectRepo.GetAll(ctx)
}

func (s *Service) GetByID(ctx context.Context, id int) (*models.StoredRepositoryResponse, error){
	return s.projectRepo.GetByID(ctx, id)
}