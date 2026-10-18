package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	custom_errors "github.com/slinkeres/open-source-project-analyzer/internal/errors"
	"github.com/slinkeres/open-source-project-analyzer/internal/models"
)

type ScoreRepository struct {
	db *pgxpool.Pool
}

func NewScoreRepository(db *pgxpool.Pool) *ScoreRepository {
	return &ScoreRepository{
		db: db,
	}
}

func (r *ScoreRepository) Upsert(ctx context.Context, score *models.RepositoryScore) (*models.RepositoryScore, error) {
	query := `
        INSERT INTO repository_scores (
            repository_id, activity_score, documentation_score,
            community_score, maintenance_score, overall_score, reasons
        )
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        ON CONFLICT (repository_id)
        DO UPDATE SET
            activity_score = EXCLUDED.activity_score,
            documentation_score = EXCLUDED.documentation_score,
            community_score = EXCLUDED.community_score,
            maintenance_score = EXCLUDED.maintenance_score,
            overall_score = EXCLUDED.overall_score,
            reasons = EXCLUDED.reasons,
            calculated_at = NOW()
        RETURNING repository_id, activity_score, documentation_score,
            community_score, maintenance_score, overall_score,
            reasons, calculated_at
    `
	var result models.RepositoryScore
	err := r.db.QueryRow(ctx, query,
		score.RepositoryID, score.ActivityScore, score.DocumentationScore,
		score.CommunityScore, score.MaintenanceScore, score.OverallScore,
		score.Reasons,
	).Scan(
		&result.RepositoryID, &result.ActivityScore, &result.DocumentationScore,
		&result.CommunityScore, &result.MaintenanceScore, &result.OverallScore,
		&result.Reasons, &result.CalculatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *ScoreRepository) GetByID(ctx context.Context, id int) (*models.RepositoryScore, error) {
	query := `
		SELECT repository_id,
        activity_score,
        documentation_score,
        community_score,
        maintenance_score,
        overall_score,
        reasons,
        calculated_at
		FROM repository_scores
		WHERE repository_id = $1
	`
	var result models.RepositoryScore
	err := r.db.QueryRow(ctx, query, id).Scan(
		&result.RepositoryID, &result.ActivityScore, &result.DocumentationScore,
		&result.CommunityScore, &result.MaintenanceScore, &result.OverallScore,
		&result.Reasons, &result.CalculatedAt,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, custom_errors.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &result, nil
}
