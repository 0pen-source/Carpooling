package service

import (
	"net/http"

	"github.com/0pen-source/Carpooling/dao"
	"github.com/0pen-source/Carpooling/models"
	"github.com/gin-gonic/gin"
)

func Phonetest(c *gin.Context) {
	payload := models.Phonetest{}
	if err := c.Bind(&payload); err != nil {
		return
	}
	_, err := dao.GetUser(payload.Phone)
	response := models.Response{}
	phonetest := models.PhoneTestResponse{}
	if err != nil {
		response.Code = http.StatusNotFound
		response.Message = "该电话号码未注册"
		phonetest.Exit = false
	} else {
		response.Code = http.StatusOK
		response.Message = "该电话号码已注册"
		phonetest.Exit = true
	}
	response.Data = phonetest
	if c.GetBool("testing") {
		c.JSON(http.StatusOK, response)
		return
	}
	c.Render(http.StatusOK, NewEncryptedJSONRender(response, []byte(dao.Config.Checkcode)))
}
