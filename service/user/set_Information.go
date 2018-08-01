package user

import (
	"net/http"

	"github.com/0pen-source/Carpooling/dao"
	"github.com/0pen-source/Carpooling/models"
	"github.com/0pen-source/Carpooling/service/common"
	"github.com/gin-gonic/gin"
)

func SetInformation(c *gin.Context) {
	payload := models.UserMessage{}
	if err := c.Bind(&payload); err != nil {
		return
	}

	user := models.User{
		Phone:    payload.Phone,
		Password: payload.Password,
		Nickname: payload.Nickname,
		Username: payload.Username,
		Sex:      payload.Sex,
		LastLat:  payload.LastLat,
		LastLon:  payload.LastLon,
	}
	err := dao.UpdateUser(user)
	response := models.Response{}
	phonetest := models.PhoneTestResponse{}
	if err != nil {
		response.Code = http.StatusNotFound
		response.Message = "更新失败，请重试"
		phonetest.Status = false
		response.Data = phonetest
		c.Render(http.StatusOK, common.NewEncryptedJSONRender(response, []byte(dao.Config.Checkcode)))
		return
	}
	response.Code = http.StatusOK
	response.Message = "更新成功"
	phonetest.Status = true
	response.Data = phonetest
	if c.GetBool("testing") {
		c.JSON(http.StatusOK, response)
		return
	}
	c.Render(http.StatusOK, common.NewEncryptedJSONRender(response, []byte(dao.Config.Checkcode)))
}
