package passengers

import (
	"fmt"
	"net/http"

	"github.com/0pen-source/Carpooling/dao"
	"github.com/0pen-source/Carpooling/models"
	"github.com/0pen-source/Carpooling/service/common"
	"github.com/gin-gonic/gin"
)

func Connected(c *gin.Context) {
	payload := models.Connected{}
	if err := c.Bind(&payload); err != nil {
		return
	}
	trip := models.AlreadyConnDriver{
		Phone: payload.Phone,
		Guid:  payload.Guid,
	}
	dao.SaveConnPassengers(trip)

	response := models.Response{}
	index := models.IndexResponse{}
	response.Code = http.StatusOK
	response.Message = "Connected"
	response.Data = index

	if c.GetBool("testing") {
		c.JSON(http.StatusOK, response)
		return
	}
	c.Render(http.StatusOK, common.NewEncryptedJSONRender(response, []byte(dao.Config.Checkcode)))
}

func GetConnecteds(c *gin.Context) {
	payload := models.Connected{}
	if err := c.Bind(&payload); err != nil {
		return
	}
	trip := models.User{
		Phone: payload.Phone,
	}
	trips, _ := dao.GetConnPassengers(trip)

	response := models.Response{}
	var index []models.ResponseTrip
	if trips != nil {
		fmt.Println("不为null")
		index = trips
	}
	response.Code = http.StatusOK
	response.Message = "GetConnecteds"
	response.Data = index

	if c.GetBool("testing") {
		c.JSON(http.StatusOK, response)
		return
	}
	c.Render(http.StatusOK, common.NewEncryptedJSONRender(response, []byte(dao.Config.Checkcode)))
}
