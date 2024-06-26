package handlers

import "github.com/gin-gonic/gin"

type GinHandler struct {
	//TODO
}

func (h GinHandler) SetupEndpoints(r *gin.Engine) {
	// endpoints
	r.GET("/", h.Index)
	r.POST("/video", h.Video)
	r.GET("/video_transcript/:id", h.VideoTranscript)

	// static content
	r.Static("/css", "web/static")
	r.Static("/img", "web/static")

	r.LoadHTMLGlob("web/templates/*")
}
