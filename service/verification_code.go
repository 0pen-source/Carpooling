package service

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/0pen-source/Carpooling/dao"
	"github.com/0pen-source/Carpooling/models"
	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
)

func GetVerificationCode(c *gin.Context) {
	payload := models.Phonetest{}
	if err := c.Bind(&payload); err != nil {
		return
	}

	var code models.Code
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	_, _, errs := gorequest.New().Get(dao.Config.VerificationCodeURL).
		Param("accesskey", dao.Config.AccessKey).
		Param("secret", dao.Config.Secret).
		Param("sign", dao.Config.Sign).
		Param("templateId", dao.Config.TemplateID).
		Param("mobile", payload.Phone).
		Param("content", vcode).
		Retry(3, 2*time.Second, http.StatusInternalServerError).
		EndStruct(&code)
	response := models.Response{}
	phonetest := models.PhoneTestResponse{}
	if len(errs) != 0 || code.Msg != "SUCCESS" {
		fmt.Println(errs)
		response.Code = http.StatusNotFound
		response.Message = "验证码发送失败"
		phonetest.Status = false
		response.Data = phonetest

	} else {
		dao.SaveCode(payload.Phone, vcode)
		response.Code = http.StatusOK
		response.Message = "验证码发送成功"
		phonetest.Status = true
		response.Data = phonetest
	}

	if c.GetBool("testing") {
		c.JSON(http.StatusOK, response)
		return
	}
	c.Render(http.StatusOK, NewEncryptedJSONRender(response, []byte(dao.Config.Checkcode)))
}
func CheckCode(c *gin.Context) {
	payload := models.Phonetest{}
	if err := c.Bind(&payload); err != nil {
		return
	}

	code, _ := dao.GetCode(payload.Phone)
	response := models.Response{}
	phonetest := models.PhoneTestResponse{}
	if code != payload.VerificationCode {
		response.Code = http.StatusNotFound
		response.Message = "验证码输入错误"
		phonetest.Status = false
		response.Data = phonetest

	} else {
		response.Code = http.StatusOK
		response.Message = "验证码输入正确"
		phonetest.Status = true
		response.Data = phonetest

	}

	if c.GetBool("testing") {
		c.JSON(http.StatusOK, response)
		return
	}
	c.Render(http.StatusOK, NewEncryptedJSONRender(response, []byte(dao.Config.Checkcode)))
}
