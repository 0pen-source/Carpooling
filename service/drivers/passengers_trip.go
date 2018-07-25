package drivers

import (
	"net/http"
	"time"

	"github.com/0pen-source/Carpooling/dao"
	"github.com/0pen-source/Carpooling/models"
	"github.com/0pen-source/Carpooling/service/common"
	"github.com/gin-gonic/gin"
)

func CreatTrip(c *gin.Context) {
	payload := models.TripMessage{}
	if err := c.Bind(&payload); err != nil {
		return
	}
	trip := models.DriverTrip{
		UserName:        payload.Username,
		NickName:        payload.Nickname,
		Phone:           payload.Phone,
		CreateTime:      time.Now().Unix(),
		TravelTime:      payload.TravelTime,
		TravelTimeTitle: payload.TravelTimeTitle,
		From:            payload.From,
		FromLat:         payload.FromLat,
		FromLon:         payload.FromLon,
		DestinationLat:  payload.DestinationLat,
		DestinationLon:  payload.DestinationLon,
		Destination:     payload.Destination,
		PayPrice:        payload.PayPrice,
		Surplus:         payload.Surplus,
	}
	trip, err := dao.SaveDriverTrip(trip)
	response := models.Response{}
	phonetest := models.PhoneTestResponse{}
	if err != nil {
		response.Code = http.StatusNotFound
		response.Message = "创建行程失败，请重试"
		phonetest.Status = false
		response.Data = phonetest
		c.JSON(http.StatusOK, response)
		return
	}
	response.Code = http.StatusOK
	response.Message = "创建行程成功"
	phonetest.Status = true
	phonetest.ID = trip.Guid
	response.Data = phonetest

	if c.GetBool("testing") {
		c.JSON(http.StatusOK, response)
		return
	}
	c.Render(http.StatusOK, common.NewEncryptedJSONRender(response, []byte(dao.Config.Checkcode)))
}
