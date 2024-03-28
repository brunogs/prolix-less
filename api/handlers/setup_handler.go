package handlers

import "github.com/gin-gonic/gin"

type GinHandler struct {
	//TODO
}

func (h GinHandler) SetupEndpoints(r *gin.Engine) {
	// endpoints
	r.GET("/", h.Index)
	r.POST("/video", h.Video)

	// static content
	r.Static("/css", "web/static")
	r.LoadHTMLGlob("web/templates/*")
}
