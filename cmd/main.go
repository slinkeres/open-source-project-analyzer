package main

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/slinkeres/open-source-project-analyzer/config"
	"github.com/slinkeres/open-source-project-analyzer/internal/handler"
	"github.com/slinkeres/open-source-project-analyzer/internal/repository/postgres"
	"github.com/slinkeres/open-source-project-analyzer/internal/service/ai"
	"github.com/slinkeres/open-source-project-analyzer/internal/service/aiservice"
	"github.com/slinkeres/open-source-project-analyzer/internal/service/github"
	"github.com/slinkeres/open-source-project-analyzer/internal/service/report"
	"github.com/slinkeres/open-source-project-analyzer/internal/service/repository"
	"github.com/slinkeres/open-source-project-analyzer/internal/service/score"
)


func main() {
	ctx := context.Background()
	cfg, err := config.Load()
	if err != nil {
		slog.Error("failed load env", "error", err)
		return
	}
	db, err := postgres.CreateDb(ctx, cfg.DBURL)
	if err != nil {
		slog.Error("failed to create db", "error", err)
		return
	}
	defer db.Close()
	slog.Info("db created")

	httpClient := &http.Client{
    	Timeout: 30 * time.Second,
	}

	githubClient := github.NewClient(httpClient, cfg.GitHubBaseURL, cfg.GitHubUserAgent)
	groqClient := ai.NewClient(httpClient, cfg.GroqAPIKey, cfg.GroqBaseURL, cfg.GroqModel)

	projectRepo := postgres.NewProjectRepository(db)
	scoreRepo := postgres.NewScoreRepository(db)
	aiRepo := postgres.NewAiRepository(db)

	repositoryService := repository.NewService(githubClient, projectRepo)
	scoreService := score.NewService(projectRepo, scoreRepo)
	aiService := aiservice.NewService(projectRepo, aiRepo, groqClient, cfg.GroqModel)
	reportService := report.NewService(projectRepo, scoreRepo, aiRepo)

	healthHandler := handler.NewHealthHandler()
	repositoryHandler := handler.NewRepositoryHandler(repositoryService)
	scoreHandler := handler.NewScoreHandler(scoreService)
	aiHandler := handler.NewAiHandler(aiService)
	reportHandler := handler.NewReportHandler(reportService)


	r := gin.Default()

	corsConfig := cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
		AllowCredentials: false,
	}
	r.Use(cors.New(corsConfig))

	r.GET("/health", healthHandler.Health)
	api := r.Group("/api/repositories")
	{
		api.POST("/preview", repositoryHandler.Preview)
		api.POST("/fetch", repositoryHandler.Fetch)
		api.GET("", repositoryHandler.List)
		api.GET("/:id", repositoryHandler.GetByID)
		api.POST("/:id/score", scoreHandler.Calculate)
		api.GET("/:id/score", scoreHandler.GetByID)
		api.POST("/:id/ai-analysis", aiHandler.Analyze)
		api.GET("/:id/ai-analysis", aiHandler.GetByID)
		api.GET("/:id/report", reportHandler.Get)
	}


	slog.Info("server starting", "port", cfg.Port)
	if err := r.Run(":"+cfg.Port); err != nil{
		slog.Error("failed to start server", "error", err)
	}
}