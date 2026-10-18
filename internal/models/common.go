package models

type RepositoryUrlRequest struct {
	URL string `json:"url"`
}

type RepositoryPreviewResponse struct {
	Owner         string `json:"owner"`
	Name          string `json:"name"`
	NormalizedURL string `json:"normalized_url"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
