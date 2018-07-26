package passengers

import (
	"net/http"

	"github.com/0pen-source/Carpooling/dao"
	"github.com/0pen-source/Carpooling/models"
	"github.com/0pen-source/Carpooling/service/common"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	payload := models.Index{}
	if err := c.Bind(&payload); err != nil {
		return
	}
	user := models.User{
		LastLat: payload.Lan,
		LastLon: payload.Lon,
		Phone:   payload.Phone,
	}

	realTrip, _ := dao.GetRealTimePassengersTrip()
	recommendTrip, _ := dao.GetRecommendDriverTrips(user)
	response := models.Response{}
	index := models.IndexResponse{}
	index.RealtimeOrder = realTrip
	index.RecommendOrder = recommendTrip
	response.Code = http.StatusOK
	response.Message = "Index"
	response.Data = index

	if c.GetBool("testing") {
		c.JSON(http.StatusOK, response)
		return
	}
	c.Render(http.StatusOK, common.NewEncryptedJSONRender(response, []byte(dao.Config.Checkcode)))
}
