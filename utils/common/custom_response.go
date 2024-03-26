package common

import (
	model "portfolio-api/utils/model"

	"github.com/gin-gonic/gin"
)

func SendSingleResponse(ctx *gin.Context, code int, description string, data any) {

	ctx.JSON(code, model.SingleResponse{
		Status: model.Status{
			Code:        code,
			Description: description,
		},
		Data: data,
	})
}

func SendPagedResponse(ctx *gin.Context, code int, description string, data any, paging any) {
	ctx.JSON(code, model.PagedResponse{
		Status: model.Status{
			Code:        code,
			Description: description,
		},
		Data:   data,
		Paging: paging,
	})
 }