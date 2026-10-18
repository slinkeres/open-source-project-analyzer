package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	custom_errors "github.com/slinkeres/open-source-project-analyzer/internal/errors"
	"github.com/slinkeres/open-source-project-analyzer/internal/models"
)

type AiRepository struct {
	db *pgxpool.Pool
}

func NewAiRepository(db *pgxpool.Pool) *AiRepository {
	return &AiRepository{
		db: db,
	}
}

func (r *AiRepository) Upsert(ctx context.Context, ai *models.RepositoryAiAnalysisResponse) (*models.RepositoryAiAnalysisResponse, error) {
	query := `
		INSERT INTO repository_ai_analysis (
			repository_id, summary, documentation_quality_score,
			onboarding_score, beginner_friendliness_score, strengths,
			weaknesses, suggested_improvements, model
		)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)
		ON CONFLICT (repository_id)
		DO UPDATE SET
			summary = EXCLUDED.summary,
			documentation_quality_score = EXCLUDED.documentation_quality_score,
			onboarding_score = EXCLUDED.onboarding_score,
			beginner_friendliness_score = EXCLUDED.beginner_friendliness_score,
			strengths = EXCLUDED.strengths,
			weaknesses = EXCLUDED.weaknesses,
			suggested_improvements = EXCLUDED.suggested_improvements,
			model = EXCLUDED.model
		RETURNING repository_id, summary, documentation_quality_score,
			onboarding_score, beginner_friendliness_score, strengths,
			weaknesses, suggested_improvements, model, analyzed_at
	`
	var result models.RepositoryAiAnalysisResponse
	err := r.db.QueryRow(ctx, query,
		ai.RepositoryID, ai.Summary, ai.DocumentationQualityScore,
		ai.OnboardingScore, ai.BeginnerFriendlinessScore,
		ai.Strengths, ai.Weaknesses, ai.SuggestedImprovements, ai.Model,
	).Scan(
		&result.RepositoryID, &result.Summary, &result.DocumentationQualityScore,
		&result.OnboardingScore, &result.BeginnerFriendlinessScore, &result.Strengths,
		&result.Weaknesses, &result.SuggestedImprovements, &result.Model, &result.AnalyzedAt,
	)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *AiRepository) GetByID(ctx context.Context, id int) (*models.RepositoryAiAnalysisResponse, error) {
	query := `
		SELECT repository_id, summary, documentation_quality_score,
			onboarding_score, beginner_friendliness_score, strengths,
			weaknesses, suggested_improvements, model, analyzed_at
		FROM repository_ai_analysis
		WHERE repository_id = $1
	`
	var result models.RepositoryAiAnalysisResponse
	err := r.db.QueryRow(ctx, query, id).Scan(
		&result.RepositoryID, &result.Summary, &result.DocumentationQualityScore,
		&result.OnboardingScore, &result.BeginnerFriendlinessScore, &result.Strengths,
		&result.Weaknesses, &result.SuggestedImprovements, &result.Model, &result.AnalyzedAt,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, custom_errors.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &result, nil
}
