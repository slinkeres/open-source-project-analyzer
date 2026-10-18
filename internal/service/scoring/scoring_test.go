package scoring

import (
	"testing"
	"time"
)

func TestCalculateDocumentationScore(t *testing.T){
	tests := []struct{
		name string
		readme string
		expected int
	}{
		{"empty", "", 10},
		{"tiny", "fdjsk", 40},
		{"small", string(make([]byte, 600)),65 },
		{"medium", string(make([]byte, 2500)), 85},
		{"large", string(make([]byte, 6000)), 100},
	}
	for _, v := range tests{
		t.Run(v.name, func(t *testing.T){
			got := calculateDocumentationScore(v.readme)
			if got != v.expected{
				t.Errorf("got %v, want %v", got, v.expected)
			}
		})
	}
}

func TestCalculateCommunityScore(t *testing.T){
	tests := []struct{
		name string
		stars int
		expected int
	}{
		{"zero", 0, 35},
		{"small", 50, 60},
		{"medium", 500, 75},
		{"big", 5000, 90},
		{"huge", 50000, 100},
	}

	for _, v := range tests{
		t.Run(v.name, func(t *testing.T){
			got := calculateCommunityScore(v.stars)
			if got != v.expected{
				t.Errorf("got %v , expected %v", got, v.expected)
			}
		})
	}
}

func TestCalculateMaintenanceScore(t *testing.T){
	tests := []struct{
		name string
		archived bool
		hasIssues bool
		hasLicense bool
		expected int
	}{
		{"acrhived", true, true, true, 20},
		{"minimal", false, false, false, 60},
		{"no license", false, true, false, 80},
		{"full", false, true, true, 100},
	}

	for _, v := range tests{
		t.Run(v.name, func(t *testing.T){
			got := calculateMaintenanceScore(v.archived, v.hasIssues, v.hasLicense)
			if got != v.expected{
				t.Errorf("got %v, expected %v", got, v.expected)
			}
		})
	}
}

func TestCalculateActivityScore(t *testing.T){
	tests := []struct{
		name string
		pushedAt string
		expected int
	}{
		 {"invalid", "not-a-date", 40},
        {"very old", "2000-01-01T00:00:00Z", 25},
        {"recent", time.Now().AddDate(0, 0, -5).Format(time.RFC3339), 100},
        {"month ago", time.Now().AddDate(0, 0, -60).Format(time.RFC3339), 85},
	}
	for _, v := range tests{
		t.Run(v.name, func(t* testing.T){
			got := calculateActivityScore(v.pushedAt)
			if got != v.expected{
				t.Errorf("got %v, expected %v", got, v.expected)
			}
		})
	}
}