package passengers

import (
	"fmt"
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
		fmt.Println(err)
		return
	}
	trip := models.PassengersTrip{
		UserName:                   payload.Username,
		NickName:                   payload.Nickname,
		Phone:                      payload.Phone,
		CreateTime:                 time.Now().Unix(),
		TravelTime:                 payload.TravelTime,
		TravelTimeTitle:            payload.TravelTimeTitle,
		From:                       payload.From,
		FromLat:                    payload.FromLat,
		FromLon:                    payload.FromLon,
		DestinationLat:             payload.DestinationLat,
		DestinationLon:             payload.DestinationLon,
		Destination:                payload.Destination,
		PayPrice:                   payload.PayPrice,
		Surplus:                    payload.Surplus,
		FromRegion:                 payload.FromRegion,
		FromCity:                   payload.FromCity,
		FromAccurateAddress:        payload.FromAccurateAddress,
		FromVagueAddress:           payload.FromVagueAddress,
		DestinationRegion:          payload.DestinationRegion,
		DestinationCity:            payload.DestinationCity,
		DestinationAccurateAddress: payload.DestinationAccurateAddress,
		DestinationVagueAddress:    payload.DestinationVagueAddress,
		Source:                     payload.Source,
		Mileage:                    payload.Mileage,
		SeatNum:                    payload.SeatNum,
		Complete:                   payload.Complete,
		Msg:                        payload.Msg,
	}
	trip, err := dao.SavePassengersTrip(trip)
	fmt.Println(err)
	response := models.Response{}
	phonetest := models.PhoneTestResponse{}
	if err != nil {
		response.Code = http.StatusNotFound
		response.Message = "创建行程失败，请重试"
		phonetest.Status = false
		response.Data = phonetest
		c.Render(http.StatusOK, common.NewEncryptedJSONRender(response, []byte(dao.Config.Checkcode)))
		return
	}
	response.Code = http.StatusOK
	response.Message = "创建行程成功，请耐心等待车主接单"
	phonetest.Status = true
	phonetest.ID = trip.Guid
	response.Data = phonetest

	if c.GetBool("testing") {
		c.JSON(http.StatusOK, response)
		return
	}
	c.Render(http.StatusOK, common.NewEncryptedJSONRender(response, []byte(dao.Config.Checkcode)))
}
