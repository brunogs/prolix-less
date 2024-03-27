package handlers

import "github.com/gin-gonic/gin"

type GinHandler struct {
	//TODO
}

func (h GinHandler) SetupEndpoints(r *gin.Engine) {
	r.GET("/", h.Index)

	r.Static("/css", "web/static")
	r.LoadHTMLGlob("web/templates/*")
}
