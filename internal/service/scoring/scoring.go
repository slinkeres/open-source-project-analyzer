package scoring

import (
	"time"

	"github.com/slinkeres/open-source-project-analyzer/internal/models"
)

type calculatedScore struct {
	ActivityScore      int
	DocumentationScore int
	CommunityScore     int
	MaintenanceScore   int
	OverallScore       int
	Reasons            []string
}

func Calculate(repository *models.StoredRepositoryResponse) *models.RepositoryScore {
	score := calculateScore(repository)
	return &models.RepositoryScore{
		ActivityScore:      score.ActivityScore,
		DocumentationScore: score.DocumentationScore,
		CommunityScore:     score.CommunityScore,
		MaintenanceScore:   score.MaintenanceScore,
		OverallScore:       score.OverallScore,
		Reasons:            score.Reasons,
	}
}

func calculateScore(repository *models.StoredRepositoryResponse) calculatedScore {
	reasons := []string{}

	activityScore := calculateActivityScore(repository.PushedAt)
	if activityScore >= 85 {
		reasons = append(reasons, "The repository has recent development activity")
	}

	documentationScore := calculateDocumentationScore(repository.Readme)
	if documentationScore >= 85 {
		reasons = append(reasons, "The repository has substantial documentation")
	}

	communityScore := calculateCommunityScore(repository.Stars)
	if communityScore >= 85 {
		reasons = append(reasons, "The repository has a large community")
	}

	maintenanceScore := calculateMaintenanceScore(
		repository.Archived,
		repository.HasIssues,
		repository.License != nil,
	)
	if maintenanceScore >= 85 {
		reasons = append(reasons, "The repository appears actively maintained")
	}

	overallScore := (activityScore + documentationScore + communityScore + maintenanceScore) / 4

	return calculatedScore{
		ActivityScore:      activityScore,
		DocumentationScore: documentationScore,
		CommunityScore:     communityScore,
		MaintenanceScore:   maintenanceScore,
		OverallScore:       overallScore,
		Reasons:            reasons,
	}
}

func calculateActivityScore(pushedAt string) int {
	parse, err := time.Parse(time.RFC3339, pushedAt)
	if err != nil {
		return 40
	}
	daysSincePush := int(time.Since(parse).Hours() / 24)
	switch {
	case daysSincePush <= 30:
		return 100
	case daysSincePush <= 90:
		return 85
	case daysSincePush <= 180:
		return 70
	case daysSincePush <= 365:
		return 50
	default:
		return 25
	}
}

func calculateDocumentationScore(readme string) int {
	length := len(readme)
	switch {
	case length >= 5000:
		return 100
	case length >= 2000:
		return 85
	case length >= 500:
		return 65
	case length >= 1:
		return 40
	default:
		return 10
	}
}

func calculateCommunityScore(stars int) int {
	switch {
	case stars >= 10000:
		return 100
	case stars >= 1000:
		return 90
	case stars >= 100:
		return 75
	case stars >= 10:
		return 60
	default:
		return 35
	}
}

func calculateMaintenanceScore(archived, hasIssues, hasLicense bool) int {
	if archived {
		return 20
	}

	score := 60
	if hasIssues {
		score += 20
	}
	if hasLicense {
		score += 20
	}
	return score
}
