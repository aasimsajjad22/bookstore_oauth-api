package http

import (
	"fmt"
	atDomain "github.com/aasimsajjad22/bookstore_oauth-api/src/domain/access_token"
	"github.com/aasimsajjad22/bookstore_oauth-api/src/services/access_token"
	"github.com/aasimsajjad22/bookstore_oauth-api/src/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(*gin.Context)
}
type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) Create(c *gin.Context) {
	var request atDomain.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	accessToken, err := handler.service.Create(request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, accessToken)
}
func (handler *accessTokenHandler) GetById(c *gin.Context) {
	fmt.Println(c.Param("access_token"))
	accessToken, err := handler.service.GetById(c.Param("access_token"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}
