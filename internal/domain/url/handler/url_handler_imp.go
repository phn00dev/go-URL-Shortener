package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/phn00dev/go-URL-Shortener/internal/domain/url/dto"
	"github.com/phn00dev/go-URL-Shortener/internal/domain/url/service"
	"github.com/phn00dev/go-URL-Shortener/internal/utils/response"
	"github.com/phn00dev/go-URL-Shortener/internal/utils/validate"
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

	urls, err := urlHandler.urlService.FindAll()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "urls error", err.Error())
		return
	}
	response.Success(c, http.StatusOK, "all urls", urls)
}

func (urlHandler urlHandlerImp) GetOne(c *gin.Context) {
	urlIdStr := c.Param("urlId")
	urlId, _ := strconv.Atoi(urlIdStr)
	// get url
	url, err := urlHandler.urlService.FindOne(urlId)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "url not found", err.Error())
		return
	}
	response.Success(c, http.StatusOK, "user url", url)
}

func (urlHandler urlHandlerImp) Create(c *gin.Context) {
	var createRequest dto.CreateUrlRequest
	userIdStr, exists := c.Get("id")
	userId := userIdStr.(int)
	if !exists {
		response.Error(c, http.StatusUnauthorized, "error auth", "User not authorized")
		return
	}

	if userId == 0 {
		response.Error(c, http.StatusUnauthorized, "error auth", "User not authorized")
		return
	}

	if err := c.ShouldBindBodyWithJSON(&createRequest); err != nil {
		response.Error(c, http.StatusBadRequest, "body parser error", err.Error())
		return
	}

	// validate error
	if err := validate.ValidateStruct(createRequest); err != nil {
		response.Error(c, http.StatusBadRequest, "validate error", err.Error())
		return
	}

	// url create
	if err := urlHandler.urlService.Create(userId, createRequest); err != nil {
		response.Error(c, http.StatusInternalServerError, "short url created error", err.Error())
		return
	}
	response.Success(c, http.StatusCreated, "short url created successfully", nil)
}

// user urls handler

func (urlHandler urlHandlerImp) GetAllUserUrls(c *gin.Context) {
	authUserId, exists := c.Get("id")
	if !exists {
		response.Error(c, http.StatusUnauthorized, "error auth", "User not authorized")
		return
	}

	authUserIdInt, ok := authUserId.(int)
	if !ok || authUserIdInt == 0 {
		response.Error(c, http.StatusUnauthorized, "error auth", "User not authorized")
		return
	}

	userUrls, err := urlHandler.urlService.FindAllUserUrls(authUserIdInt)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "something wrong", err.Error())
		return
	}
	response.Success(c, http.StatusOK, "user urls", userUrls)
}

func (urlHandler urlHandlerImp) GetOneUserUrl(c *gin.Context) {
	authUserId, exists := c.Get("id")
	if !exists {
		response.Error(c, http.StatusUnauthorized, "error auth", "User not authorized")
		return
	}

	authUserIdInt, ok := authUserId.(int)
	if !ok || authUserIdInt == 0 {
		response.Error(c, http.StatusUnauthorized, "error auth", "User not authorized")
		return
	}

	urlIdStr := c.Param("urlId")
	urlId, _ := strconv.Atoi(urlIdStr)

	url, err := urlHandler.urlService.FindOneUserUrl(authUserIdInt, urlId)
	if err != nil {
		response.Error(c, http.StatusNotFound, "url not found", err.Error())
		return
	}
	response.Success(c, http.StatusOK, "user url", url)
}

func (urlHandler urlHandlerImp) Delete(c *gin.Context) {
	// UserID almak
	userIdStr, exists := c.Get("id")
	if !exists || userIdStr == nil {
		response.Error(c, http.StatusUnauthorized, "error auth", "User not authorized")
		return
	}
	userId := userIdStr.(int)
	if userId == 0 {
		response.Error(c, http.StatusUnauthorized, "error auth", "User not authorized")
		return
	}

	// URL ID almak
	urlIdStr := c.Param("urlId")
	urlId, err := strconv.Atoi(urlIdStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid URL ID", "URL ID is not valid")
		return
	}

	// URL-i pozmak
	if err := urlHandler.urlService.Delete(userId, urlId); err != nil {
		response.Error(c, http.StatusBadRequest, "deleted error", err.Error())
		return
	}
	response.Success(c, http.StatusOK, "URL deleted successfully", nil)
}
