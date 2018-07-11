package service

import (
	"net/http"

	"github.com/0pen-source/Carpooling/dao"
	"github.com/0pen-source/Carpooling/models"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

func Login(c *gin.Context) {
	payload := models.Login{}
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
		c.JSON(http.StatusOK, response)
		return

	}
	if user.Password != payload.Password {
		phonetest := models.PhoneTestResponse{}
		response.Code = http.StatusNotFound
		response.Message = "手机号和密码不匹配"
		phonetest.Exit = false
		response.Data = phonetest
		c.JSON(http.StatusOK, response)
		return

	}
	login := models.LoginResponse{
		Token:         xid.New().String(),
		UserName:      user.Username,
		Sex:           user.Sex,
		Uid:           user.Guid,
		Balance:       user.Balance,
		NickName:      user.Nickname,
		Last_location: user.LastLocation,
	}
	response.Code = http.StatusOK
	response.Data = login
	if c.GetBool("testing") {
		c.JSON(http.StatusOK, response)
		return
	}
	c.Render(http.StatusOK, NewEncryptedJSONRender(response, []byte(Checkcode)))
}
