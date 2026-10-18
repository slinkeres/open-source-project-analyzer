package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	custom_errors "github.com/slinkeres/open-source-project-analyzer/internal/errors"
	"github.com/slinkeres/open-source-project-analyzer/internal/service/aiservice"
)

type AiHandler struct {
	service *aiservice.Service
}

func NewAiHandler(service *aiservice.Service) *AiHandler{
	return &AiHandler{
		service: service,
	}
}

func (h *AiHandler) Analyze(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		custom_errors.Respond(c, fmt.Errorf("%w: invalid repository id", custom_errors.ErrBadRequest))
		return
	}
	result, err := h.service.Analyze(c.Request.Context(), id)
	if err != nil{
		custom_errors.Respond(c, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *AiHandler) GetByID(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		custom_errors.Respond(c, fmt.Errorf("%w: invalid repository id", custom_errors.ErrBadRequest))
		return
	}
	result, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil{
		custom_errors.Respond(c, err)
		return
	}
	c.JSON(http.StatusOK, result)
}