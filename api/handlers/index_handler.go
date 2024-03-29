package handlers

import (
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func (h GinHandler) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func (h GinHandler) Video(c *gin.Context) {
	parameters := make(map[string]string)
	if err := c.Bind(parameters); err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
	}
	if videoUrl, ok := parameters["video-url"]; ok {
		urlParsed, err := url.Parse(videoUrl)
		if err != nil {
			log.Println(err)
			c.Status(http.StatusInternalServerError)
		}
		videoID := urlParsed.Query().Get("v")
		parameters["videoID"] = videoID
	}
	c.HTML(http.StatusOK, "video.html", parameters)
}
