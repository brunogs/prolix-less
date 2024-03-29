package handlers

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"prolixless/api/client"
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
		videoTranscriptID, err := client.Transcript(videoUrl)
		if err != nil {
			log.Println("failed to transcript video", err)
		}
		log.Println("processing video id=", videoTranscriptID)
		urlParsed, err := url.Parse(videoUrl)
		if err != nil {
			log.Println(err)
			c.Status(http.StatusInternalServerError)
		}
		videoID := urlParsed.Query().Get("v")
		parameters["videoID"] = videoID
		parameters["videoTranscriptID"] = videoTranscriptID
	}
	c.HTML(http.StatusOK, "video.html", parameters)
}

func (h GinHandler) VideoTranscript(c *gin.Context) {
	idStr := c.Param("id")
	file, err := os.ReadFile("/resources/" + idStr + ".txt")
	if err != nil {
		log.Println("file not found", idStr, err)
		c.Status(http.StatusNoContent)
	} else {
		fmt.Println(string(file))
		c.String(http.StatusOK, "<p>%s</p>", string(file))
	}
}
