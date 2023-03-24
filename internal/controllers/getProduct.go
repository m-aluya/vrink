package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProduct(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Get products")

}
