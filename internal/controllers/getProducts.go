package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProducts(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Get products")

}
