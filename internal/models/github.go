package models

type GitHubOwner struct {
	Login string `json:"login"`
}

type GitHubLicense struct {
	SpdxID string `json:"spdx_id"`
}

type GitHubRepositoryResponse struct {
	ID               int            `json:"id"`
	Name             string         `json:"name"`
	FullName         string         `json:"full_name"`
	Owner            GitHubOwner    `json:"owner"`
	Description      *string        `json:"description"`
	StargazersCount  int            `json:"stargazers_count"`
	ForksCount       int            `json:"forks_count"`
	OpenIssuesCount  int            `json:"open_issues_count"`
	SubscribersCount int            `json:"subscribers_count"`
	Language         *string        `json:"language"`
	License          *GitHubLicense `json:"license"`
	Archived         bool           `json:"archived"`
	HasIssues        bool           `json:"has_issues"`
	HTMLURL          string         `json:"html_url"`
	DefaultBranch    string         `json:"default_branch"`
	Topics           []string       `json:"topics"`
	CreatedAt        string         `json:"created_at"`
	UpdatedAt        string         `json:"updated_at"`
	PushedAt         string         `json:"pushed_at"`
}

type GitHubReadmeResponse struct {
	DownloadURL string `json:"download_url"`
}
