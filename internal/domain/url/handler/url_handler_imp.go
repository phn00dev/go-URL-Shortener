package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/phn00dev/go-URL-Shortener/internal/domain/url/service"
)

type urlHandlerImp struct {
	urlService service.UrlService
}

func NewUrlHandler(service service.UrlService) UrlHandler {
	return urlHandlerImp{
		urlService: service,
	}
}

func (urlHandler urlHandlerImp) GetAll(c *gin.Context) {

}

func (urlHandler urlHandlerImp) GetOne(c *gin.Context) {
	panic("url handler impliment")
}

func (urlHandler urlHandlerImp) Create(c *gin.Context) {
	panic("url handler impliment")
}

func (urlHandler urlHandlerImp) Delete(c *gin.Context) {
	panic("url handler impliment")
}
