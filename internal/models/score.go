package models

import "time"

type RepositoryScore struct {
	RepositoryID       int       `json:"repository_id"`
	ActivityScore      int       `json:"activity_score"`
	DocumentationScore int       `json:"documentation_score"`
	CommunityScore     int       `json:"community_score"`
	MaintenanceScore   int       `json:"maintenance_score"`
	OverallScore       int       `json:"overall_score"`
	Reasons            []string  `json:"reasons"`
	CalculatedAt       time.Time `json:"calculated_at"`
}
