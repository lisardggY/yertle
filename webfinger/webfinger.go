package webfinger

import (
	"github.com/gin-gonic/gin"
)

type request struct {
	account string `form:"resource" binding:"required,startswith=acct:"`
}

func fromQuery(ctx *gin.Context) (request, error) {
	req := request{}
	if err := ctx.BindQuery(req); err != nil {
		return req, err
	}

	return req, nil
}

func Handler(ctx *gin.Context) {
	request, err := fromQuery(ctx)
	if err != nil {
		ctx.AbortWithError(500, err)
	}

	ctx.AsciiJSON(200, request.account)
}
