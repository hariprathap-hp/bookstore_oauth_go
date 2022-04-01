package http_handler

import (
	"fmt"
	"net/http"
	"strings"
	"test3/hariprathap-hp/bookstore_oauth_go/src/domain/access_token"
	"test3/hariprathap-hp/bookstore_oauth_go/src/token_service"

	"github.com/gin-gonic/gin"
)

type HTTPHandler interface {
	GetbyID(c *gin.Context)
	Create(c *gin.Context)
}

type httpHandler struct {
	service token_service.Service
}

func NewHTTPHandler(service token_service.Service) HTTPHandler {
	return &httpHandler{
		service: service,
	}
}

func (h *httpHandler) GetbyID(c *gin.Context) {
	accessTokenId := strings.TrimSpace(c.Param("access_token_id"))
	access_token, err := h.service.GetbyID(accessTokenId)
	if err != nil {
		c.JSON(err.Status, err)
	}
	c.JSON(http.StatusOK, access_token)
}

func (h *httpHandler) Create(c *gin.Context) {
	var at access_token.AccessToken
	bindErr := c.ShouldBindJSON(&at)
	if bindErr != nil {
		c.JSON(http.StatusNotImplemented, bindErr)
	}
	err := h.service.Create(at)
	if err != nil {
		c.JSON(http.StatusNotImplemented, err)
	}
	fmt.Println(at)
}
