package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type TranscriptProcess struct {
	Message string `json:"message"`
	VideoID string `json:"video_id"`
}

func Transcript(videoUrl string) (string, error) {
	response, err := http.Post("http://prolix_transcript:5000/process_video?video_url="+videoUrl, "application/json", nil)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return "", err
	}
	fmt.Println(string(body))
	var resp TranscriptProcess
	if err = json.Unmarshal(body, &resp); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return "", err
	}
	fmt.Println(resp)
	return resp.VideoID, nil
}
