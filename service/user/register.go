package user

import (
	"net/http"

	"github.com/0pen-source/Carpooling/dao"
	"github.com/0pen-source/Carpooling/models"
	"github.com/0pen-source/Carpooling/service/common"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	payload := models.UserMessage{}
	if err := c.Bind(&payload); err != nil {
		return
	}
	user := models.User{
		Phone:    payload.Phone,
		Password: payload.Password,
		Nickname: payload.Nickname,
	}
	err := dao.SaveUser(user)
	response := models.Response{}
	phonetest := models.PhoneTestResponse{}
	if err != nil {
		response.Code = http.StatusNotFound
		response.Message = "注册失败，请重试"
		phonetest.Status = false
		response.Data = phonetest
		if c.GetBool("testing") {
			c.JSON(http.StatusOK, response)
			return
		}
		c.Render(http.StatusOK, common.NewEncryptedJSONRender(response, []byte(dao.Config.Checkcode)))
		return
	}
	response.Code = http.StatusOK
	response.Message = "注册成功，请登录"
	phonetest.Status = true
	response.Data = phonetest
	if c.GetBool("testing") {
		c.JSON(http.StatusOK, response)
		return
	}
	c.Render(http.StatusOK, common.NewEncryptedJSONRender(response, []byte(dao.Config.Checkcode)))
}
