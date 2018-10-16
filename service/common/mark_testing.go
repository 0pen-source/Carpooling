package common

import (
	"fmt"
	"net/http"

	"github.com/0pen-source/Carpooling/dao"
	"github.com/0pen-source/Carpooling/models"
	"github.com/gin-gonic/gin"
)

// MarkTesting returns a middleware that marks the request as testing
func MarkTesting() gin.HandlerFunc {
	return func(c *gin.Context) {
		testing := c.Request.Header.Get("Origin") == "test"

		c.Set("testing", testing)

		c.Next()
	}
}

// MarkTesting returns a middleware that marks the request as testing
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		payload := models.Phonetest{}
		requestToken := c.Request.Header.Get("Token")
		if err := c.Bind(&payload); err != nil {
			return
		}
		token, errs := dao.GetToken(payload.Phone)

		response := models.Response{}
		if errs != nil {
			response.Message = fmt.Sprintf("%s_%s", "用户token失效", errs.Error())
			response.Code = 400
			c.Render(http.StatusBadRequest, NewEncryptedJSONRender(response, []byte(dao.Config.Checkcode)))
			return
		}
		if token != requestToken {
			response.Message = "用户token失效"
			response.Code = 400
			c.Render(http.StatusBadRequest, NewEncryptedJSONRender(response, []byte(dao.Config.Checkcode)))

			return

		}
	}
}
