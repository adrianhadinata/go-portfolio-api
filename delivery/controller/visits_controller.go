package controller

import (
	"net/http"
	"portfolio-api/entity"
	"portfolio-api/usecase"
	"portfolio-api/utils/common"

	"github.com/gin-gonic/gin"
)

type VisitsController interface {
	Insert(ctx *gin.Context)
	Route()
}

type visitsController struct {
	uc usecase.VisitsUsecase
	rg *gin.RouterGroup
}

func (c *visitsController) Insert(ctx *gin.Context) {
	var payload entity.Visitor

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	response, err := c.uc.Insert(payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendSingleResponse(ctx, http.StatusCreated, "success", response)
}

func (c *visitsController) Route() {
	
	visitors := c.rg.Group("/visitors")
	{
		visitors.POST("/", c.Insert)
	}
	
}

func NewVisitsController(uc usecase.VisitsUsecase, rg *gin.RouterGroup) VisitsController {
	return &visitsController{uc: uc, rg: rg}
}