package dto

import (
	"github.com/gin-gonic/gin"

	"github.com/phn00dev/go-URL-Shortener/internal/model"
)

// UrlResponse struct'yny saklamak
type UrlResponse struct {
	ID          int    `json:"id"`
	OriginalUrl string `json:"original_url"`
	ShortUrl    string `json:"short_url"`
	UserID      int    `json:"user_id"`
	ClickCount  int    `json:"click_count"`
	CreatedAt   string `json:"created_at"`
}

func GetAllUserUrlResponse(c *gin.Context, urls []model.Url) []UrlResponse {
	var urlResponses []UrlResponse
	domain := c.Request.Host

	for _, url := range urls {
		urlResponse := UrlResponse{
			ID:          url.ID,
			OriginalUrl: url.OriginalUrl,
			ShortUrl:    "http://" + domain + "/" + url.ShortUrl,
			UserID:      url.UserID,
			ClickCount:  url.ClickCount,
			CreatedAt:   url.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		urlResponses = append(urlResponses, urlResponse)
	}

	return urlResponses
}

func GetOneUserUrlResponse(c *gin.Context, url *model.Url) *UrlResponse {
	domain := c.Request.Host
	return &UrlResponse{
		ID:          url.ID,
		OriginalUrl: url.OriginalUrl,
		ShortUrl:    "http://" + domain + "/" + url.ShortUrl,
		UserID:      url.UserID,
		ClickCount:  url.ClickCount,
		CreatedAt:   url.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
