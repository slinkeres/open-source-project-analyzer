package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	custom_errors "github.com/slinkeres/open-source-project-analyzer/internal/errors"
	"github.com/slinkeres/open-source-project-analyzer/internal/models"
)

type ProjectRepository struct {
	db *pgxpool.Pool
}

func NewProjectRepository(db *pgxpool.Pool) *ProjectRepository {
	return &ProjectRepository{
		db: db,
	}
}

func (r *ProjectRepository) Upsert(ctx context.Context, repository *models.StoredRepositoryResponse) (int, error) {
	query := `INSERT INTO repositories(  
		github_id,
        owner,
        name,
        full_name,
        description,
        stars,
        forks,
        open_issues,
        subscribers,
        language,
        license,
        archived,
        has_issues,
        url,
        default_branch,
        topics,
        github_created_at,
        github_updated_at,
        pushed_at,
        readme
    )
    VALUES (
        $1, $2, $3, $4, $5,
        $6, $7, $8, $9, $10,
        $11, $12, $13, $14, $15,
        $16, $17, $18, $19, $20
    )
    ON CONFLICT (github_id)
    DO UPDATE SET
        owner = EXCLUDED.owner,
        name = EXCLUDED.name,
        full_name = EXCLUDED.full_name,
        description = EXCLUDED.description,
        stars = EXCLUDED.stars,
        forks = EXCLUDED.forks,
        open_issues = EXCLUDED.open_issues,
        subscribers = EXCLUDED.subscribers,
        language = EXCLUDED.language,
        license = EXCLUDED.license,
        archived = EXCLUDED.archived,
        has_issues = EXCLUDED.has_issues,
        url = EXCLUDED.url,
        default_branch = EXCLUDED.default_branch,
        topics = EXCLUDED.topics,
        github_created_at = EXCLUDED.github_created_at,
        github_updated_at = EXCLUDED.github_updated_at,
        pushed_at = EXCLUDED.pushed_at,
        readme = EXCLUDED.readme,
        updated_at = NOW()
    RETURNING id`

	var id int
	err := r.db.QueryRow(ctx, query,
		repository.GithubID,
		repository.Owner,
		repository.Name,
		repository.FullName,
		repository.Description,
		repository.Stars,
		repository.Forks,
		repository.OpenIssues,
		repository.Subscribers,
		repository.Language,
		repository.License,
		repository.Archived,
		repository.HasIssues,
		repository.URL,
		repository.DefaultBranch,
		repository.Topics,
		repository.GithubCreatedAt,
		repository.GithubUpdatedAt,
		repository.PushedAt,
		repository.Readme,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *ProjectRepository) GetByID(ctx context.Context, id int) (*models.StoredRepositoryResponse, error) {
	query := `
        SELECT
            id, github_id, owner, name, full_name, description,
            stars, forks, open_issues, subscribers,
            language, license, archived, has_issues,
            url, default_branch, topics,
            github_created_at, github_updated_at, pushed_at, readme
        FROM repositories
        WHERE id = $1
    `
	var result models.StoredRepositoryResponse
	err := r.db.QueryRow(ctx, query, id).Scan(
		&result.DatabaseID, &result.GithubID, &result.Owner, &result.Name,
		&result.FullName, &result.Description, &result.Stars, &result.Forks,
		&result.OpenIssues, &result.Subscribers, &result.Language, &result.License,
		&result.Archived, &result.HasIssues, &result.URL, &result.DefaultBranch,
		&result.Topics, &result.GithubCreatedAt, &result.GithubUpdatedAt,
		&result.PushedAt, &result.Readme,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, custom_errors.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *ProjectRepository) GetAll(ctx context.Context) ([]models.RepositoryListItemResponse, error) {
	query := `
        SELECT
            id, full_name, description, stars, forks,
            open_issues, language, license, archived, default_branch
        FROM repositories
        ORDER BY stars DESC
    `
	var result []models.RepositoryListItemResponse
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var item models.RepositoryListItemResponse
		err = rows.Scan(
			&item.DatabaseID, &item.FullName, &item.Description,
			&item.Stars, &item.Forks, &item.OpenIssues,
			&item.Language, &item.License, &item.Archived, &item.DefaultBranch,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, item)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return result, nil
}
