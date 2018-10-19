package drivers

import (
	"net/http"
	"strings"
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
	if trip.NickName== trip.Phone || trip.NickName == ""{
		trip.NickName=strings.Join([]string{trip.Phone[:4],"***",trip.Phone[7:]},"")
	}
	trip, err := dao.SaveDriverTrip(trip)
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
func MyTrip(c *gin.Context) {
	payload := models.UserMessage{}
	if err := c.Bind(&payload); err != nil {
		return
	}
	user := models.User{
		Phone: payload.Phone,
	}
	trips, _ := dao.GetMyDriverTrip(user)
	response := models.Response{}
	var index []models.ResponseTrip
	if trips != nil {
		index = trips
		response.Data = index
	} else {
		response.Data = struct {
		}{}
	}

	if c.GetBool("testing") {
		c.JSON(http.StatusOK, response)
		return
	}
	c.Render(http.StatusOK, common.NewEncryptedJSONRender(response, []byte(dao.Config.Checkcode)))
}
func GetPhone(c *gin.Context) {
	payload := models.Connected{}
	if err := c.Bind(&payload); err != nil {
		return
	}

	trips, _ := dao.GetPhoneBYDriveerTripGUID(payload)
	response := models.Response{}
	response.Data = trips

	if c.GetBool("testing") {
		c.JSON(http.StatusOK, response)
		return
	}
	c.Render(http.StatusOK, common.NewEncryptedJSONRender(response, []byte(dao.Config.Checkcode)))
}
