package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	custom_errors "github.com/slinkeres/open-source-project-analyzer/internal/errors"
	"github.com/slinkeres/open-source-project-analyzer/internal/models"
	"github.com/slinkeres/open-source-project-analyzer/internal/service/repository"
)

type RepositoryHandler struct {
	service *repository.Service
}

func NewRepositoryHandler(service *repository.Service) *RepositoryHandler{
	return &RepositoryHandler{
		service: service,
	}
}

func (h *RepositoryHandler) Preview (c *gin.Context){
	var req models.RepositoryUrlRequest
	if err := c.ShouldBindJSON(&req); err != nil{
		custom_errors.Respond(c, fmt.Errorf("%w: invalid request body", custom_errors.ErrBadRequest))
		return
	}
	preview, err := h.service.Preview(req.URL)
	if err != nil {
		custom_errors.Respond(c, err)
		return
	}
	c.JSON(http.StatusOK, preview)
}

func (h *RepositoryHandler) Fetch(c *gin.Context){
	var req models.RepositoryUrlRequest
	if err := c.ShouldBindJSON(&req); err != nil{
		custom_errors.Respond(c, fmt.Errorf("%w: invalid request body", custom_errors.ErrBadRequest))
		return
	}
	repository, err := h.service.Fetch(c.Request.Context(), req.URL)
	if err != nil{
		custom_errors.Respond(c, err)
		return
	}
	c.JSON(http.StatusOK, repository)
}

func (h *RepositoryHandler) List(c *gin.Context){
	items, err := h.service.List(c.Request.Context())
	if err != nil{
		custom_errors.Respond(c, err)
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *RepositoryHandler) GetByID(c *gin.Context){
	id , err := strconv.Atoi(c.Param("id"))
	if err != nil{
		custom_errors.Respond(c, fmt.Errorf("%w: invalid repository id", custom_errors.ErrBadRequest))
		return
	}
	repo, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil{
		custom_errors.Respond(c, err)
		return
	}
	c.JSON(http.StatusOK, repo)
}