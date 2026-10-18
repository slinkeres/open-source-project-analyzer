package models

type RepositoryFetchResponse struct {
	DatabaseID      int      `json:"database_id"`
	GithubID        int      `json:"github_id"`
	Owner           string   `json:"owner"`
	Name            string   `json:"name"`
	FullName        string   `json:"full_name"`
	Description     *string  `json:"description"`
	Stars           int      `json:"stars"`
	Forks           int      `json:"forks"`
	OpenIssues      int      `json:"open_issues"`
	Subscribers     int      `json:"subscribers"`
	Language        *string  `json:"language"`
	License         *string  `json:"license"`
	Archived        bool     `json:"archived"`
	HasIssues       bool     `json:"has_issues"`
	URL             string   `json:"url"`
	DefaultBranch   string   `json:"default_branch"`
	Topics          []string `json:"topics"`
	GithubCreatedAt string   `json:"github_created_at"`
	GithubUpdatedAt string   `json:"github_updated_at"`
	PushedAt        string   `json:"pushed_at"`
	Readme          string   `json:"readme"`
}

type RepositoryListItemResponse struct {
	DatabaseID    int     `json:"database_id"`
	FullName      string  `json:"full_name"`
	Description   *string `json:"description"`
	Stars         int     `json:"stars"`
	Forks         int     `json:"forks"`
	OpenIssues    int     `json:"open_issues"`
	Language      *string `json:"language"`
	License       *string `json:"license"`
	Archived      bool    `json:"archived"`
	DefaultBranch string  `json:"default_branch"`
}

type StoredRepositoryResponse struct {
	DatabaseID      int      `json:"database_id"`
	GithubID        int      `json:"github_id"`
	Owner           string   `json:"owner"`
	Name            string   `json:"name"`
	FullName        string   `json:"full_name"`
	Description     *string  `json:"description"`
	Stars           int      `json:"stars"`
	Forks           int      `json:"forks"`
	OpenIssues      int      `json:"open_issues"`
	Subscribers     int      `json:"subscribers"`
	Language        *string  `json:"language"`
	License         *string  `json:"license"`
	Archived        bool     `json:"archived"`
	HasIssues       bool     `json:"has_issues"`
	URL             string   `json:"url"`
	DefaultBranch   string   `json:"default_branch"`
	Topics          []string `json:"topics"`
	GithubCreatedAt string   `json:"github_created_at"`
	GithubUpdatedAt string   `json:"github_updated_at"`
	PushedAt        string   `json:"pushed_at"`
	Readme          string   `json:"readme"`
}
