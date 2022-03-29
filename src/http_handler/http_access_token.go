package http_handler

import (
	"net/http"
	"strings"
	"test3/hariprathap-hp/bookstore_oauth_go/src/token_service"

	"github.com/gin-gonic/gin"
)

type HTTPHandler interface {
	GetbyID(c *gin.Context)
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
