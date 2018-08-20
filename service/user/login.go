package user

import (
	"net/http"

	"github.com/0pen-source/Carpooling/dao"
	"github.com/0pen-source/Carpooling/models"
	"github.com/0pen-source/Carpooling/service/common"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

func Login(c *gin.Context) {
	payload := models.UserMessage{}
	if err := c.Bind(&payload); err != nil {
		return
	}
	user, err := dao.GetUser(payload.Phone)
	response := models.Response{}
	if err != nil {

		phonetest := models.PhoneTestResponse{}
		response.Code = http.StatusNotFound
		response.Message = "该电话号码未注册"
		phonetest.Exit = false
		response.Data = phonetest
		if c.GetBool("testing") {
			c.JSON(http.StatusOK, response)
			return
		}
		c.Render(http.StatusOK, common.NewEncryptedJSONRender(response, []byte(dao.Config.Checkcode)))
		return

	}
	if user.Password != payload.Password {

		phonetest := models.PhoneTestResponse{}
		response.Code = http.StatusNotFound
		response.Message = "手机号和密码不匹配"
		phonetest.Exit = false
		response.Data = phonetest
		if c.GetBool("testing") {
			c.JSON(http.StatusOK, response)
			return
		}
		c.Render(http.StatusOK, common.NewEncryptedJSONRender(response, []byte(dao.Config.Checkcode)))
		return

	}
	token := xid.New().String()
	login := models.LoginResponse{
		UserName: user.Username,
		Sex:      user.Sex,
		Uid:      user.Guid,
		Balance:  user.Balance,
		NickName: user.Nickname,
		LastLon:  user.LastLon,
		LastLat:  user.LastLat,
		Token:    token,
	}
	response.Code = http.StatusOK
	response.Data = login

	dao.SaveToken(user.Phone, token)
	c.Header("token", token)

	if c.GetBool("testing") {
		c.JSON(http.StatusOK, response)
		return
	}
	c.Render(http.StatusOK, common.NewEncryptedJSONRender(response, []byte(dao.Config.Checkcode)))
}
