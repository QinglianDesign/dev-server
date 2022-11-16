package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func data_collection(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
	})
}
