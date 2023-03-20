package servehtml

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHTML(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", gin.H{})
}
