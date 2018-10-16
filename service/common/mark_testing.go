package common

import (
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
		response.Message = "用户token失效"
		response.Code = 400
		if errs != nil {
			if c.GetBool("testing") {
				c.JSON(http.StatusBadRequest, NewEncryptedJSONRender(response, []byte(dao.Config.Checkcode)))
				//c.AbortWithStatusJSON(400, gin.H{
				//	"error": fmt.Sprintf("%s_%s","用户token失效",errs.Error()),
				//})
				return
			}
			c.JSON(http.StatusBadRequest, NewEncryptedJSONRender(response, []byte(dao.Config.Checkcode)))
			//c.AbortWithStatusJSON(400, gin.H{
			//	"error": fmt.Sprintf("%s_%s","用户token失效",errs.Error()),
			//})
			return
		}
		if token != requestToken {
			c.JSON(http.StatusBadRequest, NewEncryptedJSONRender(response, []byte(dao.Config.Checkcode)))
			return

		}
		c.Next()
	}
}
