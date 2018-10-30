package passengers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/0pen-source/Carpooling/dao"
	"github.com/0pen-source/Carpooling/models"
	"github.com/0pen-source/Carpooling/service/common"
	"github.com/gin-gonic/gin"
)

func SearchTrip(c *gin.Context) {
	payload := models.TripMessage{}
	fmt.Println("SearchTrip")
	if err := c.Bind(&payload); err != nil {
		return
	}
	trip := models.PassengersTrip{
		From:                payload.From,
		FromLon:             payload.FromLon,
		FromLat:             payload.FromLat,
		FromRegion:          payload.FromRegion,
		FromCity:            payload.FromCity,
		FromDistrict:        payload.FromDistrict,
		Destination:         payload.Destination,
		DestinationLon:      payload.DestinationLon,
		DestinationLat:      payload.DestinationLat,
		DestinationRegion:   payload.DestinationRegion,
		DestinationCity:     payload.DestinationCity,
		DestinationDistrict: payload.DestinationDistrict,
		TravelTime:          payload.TravelTime,
		Surplus:             payload.Surplus,
	}
	if trip.TravelTime == "" {
		trip.TravelTime = strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	}

	recommendTrip, _ := dao.GetSearchDriverTrips(trip)

	response := models.Response{}
	index := models.IndexResponse{}
	index.RecommendOrder = recommendTrip
	response.Code = http.StatusOK
	response.Message = "searchTrip"
	response.Data = index

	if c.GetBool("testing") {
		c.JSON(http.StatusOK, response)
		return
	}
	c.Render(http.StatusOK, common.NewEncryptedJSONRender(response, []byte(dao.Config.Checkcode)))
}
