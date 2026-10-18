package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	custom_errors "github.com/slinkeres/open-source-project-analyzer/internal/errors"
	"github.com/slinkeres/open-source-project-analyzer/internal/service/report"
)

type ReportHandler struct {
	service *report.Service
}

func NewReportHandler(service *report.Service) *ReportHandler{
	return &ReportHandler{
		service: service,
	}
}

func (h *ReportHandler) Get(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		custom_errors.Respond(c, fmt.Errorf("%w: invalid repository id", custom_errors.ErrBadRequest))
		return
	}
	result, err := h.service.GetReport(c.Request.Context(), id)
	if err != nil{
		custom_errors.Respond(c, err)
		return
	}
	c.JSON(http.StatusOK, result)
}