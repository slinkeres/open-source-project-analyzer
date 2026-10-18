package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBURL           string
    Port            string
    GroqAPIKey      string
    GroqModel       string
    GroqBaseURL     string
    GitHubBaseURL   string
    GitHubUserAgent string

}

func Load() (*Config, error) {
	_ = godotenv.Load() 
	env := os.Getenv("ENV")
	var dbURL string
	switch env {
	case "localhost":
		dbURL = os.Getenv("POSTGRES_STRING_LOCALHOST")
	case "docker":
		dbURL = os.Getenv("POSTGRES_STRING_DOCKER")
	default:
		dbURL = os.Getenv("POSTGRES_STRING_DOCKER")
	}
	if dbURL == "" {
		return nil, fmt.Errorf("database URL is empty (check ENV and POSTGRES_STRING_* variables)") 
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	groqKey := os.Getenv("GROQ_API_KEY")
	if groqKey == ""{
		return nil, fmt.Errorf("groq_api_key is empty")
	}
	return &Config{
        DBURL:           dbURL,
        Port:            port,
        GroqAPIKey:      groqKey,
        GroqModel:       os.Getenv("GROQ_MODEL"),
        GroqBaseURL:     "https://api.groq.com/openai/v1",
        GitHubBaseURL:   os.Getenv("GITHUB_BASE_URL"),
        GitHubUserAgent: os.Getenv("GITHUB_USER_AGENT"),
    }, nil

}
