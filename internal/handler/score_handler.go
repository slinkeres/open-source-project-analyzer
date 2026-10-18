package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	custom_errors "github.com/slinkeres/open-source-project-analyzer/internal/errors"
	"github.com/slinkeres/open-source-project-analyzer/internal/service/score"
)



type ScoreHandler struct {
	service *score.Service
}

func NewScoreHandler(service *score.Service) *ScoreHandler{
	return &ScoreHandler{service: service}
}

func (h *ScoreHandler) Calculate(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		custom_errors.Respond(c, fmt.Errorf("%w: invalid repository id", custom_errors.ErrBadRequest))
		return
	}
	score, err := h.service.Calculate(c.Request.Context(), id)
	if err != nil{
		custom_errors.Respond(c, err)
		return
	}
	c.JSON(http.StatusOK, score)
}

func (h *ScoreHandler) GetByID(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		custom_errors.Respond(c, fmt.Errorf("%w: invalid repository id", custom_errors.ErrBadRequest))
		return
	}
	score, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		custom_errors.Respond(c, err)
		return
	}
	c.JSON(http.StatusOK, score)
}