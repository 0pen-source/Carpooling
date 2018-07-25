package dao

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/0pen-source/Carpooling/models"
	"github.com/rs/xid"
)

// GetRealTimeTrip _
func GetRealTimeDriverTrip() (user models.User, err error) {
	query := "SELECT * FROM `passengers_trip` order by create_time desc limit 20"

	value, ok := memCache.Get(query).(string)
	if !ok {
		err = cacheDB.Get(&user, query)
		if err == nil {
			userbyte, _ := json.Marshal(user)
			redisManager.SetKey(query, string(userbyte))
		}
		return user, err
	}

	json.Unmarshal([]byte(value), &user)
	return user, nil

}
func GetRecommendDriverTrips(user models.User) (trips []models.ResponseTrip, err error) {
	query := "SELECT *," +
		"ROUND(6378.138 * 2 * ASIN(SQRT(POW(SIN((" +
		":last_lat * PI() / 180 - from_lat * PI() / 180) / 2),2" +
		") + COS(:last_lat * PI() / 180) * COS(from_lat * PI() / 180) * POW(SIN((:last_lon * PI() / 180 " +
		"- from_lon * PI() / 180) / 2), 2))) * 1000) AS juli FROM driver_trip ORDER BY juli ASC limit 20"
	fmt.Println(fmt.Sprintf("%s-%s-%s", user.LastLat, user.LastLon, "driver"))
	trips, ok := memCache.Get(fmt.Sprintf("%s-%s-%s", user.LastLat, user.LastLon, "driver")).([]models.ResponseTrip)
	if !ok {
		err = cacheDB.Select(&trips, query, user)
	}
	memCache.Put(fmt.Sprintf("%s-%s-%s", user.LastLat, user.LastLon, "driver"), trips, time.Minute*10)

	return trips, nil

}

// SaveDriverTrip _
func SaveDriverTrip(trip models.DriverTrip) (models.DriverTrip, error) {
	trip.Guid = xid.New().String()
	_, err = cacheDB.NamedExec("INSERT INTO driver_trip "+
		"(`guid`,`username`,`nickname`,`phone`,`create_time`,`travel_time`,`travel_time_title`,`from`,`from_lon`,`from_lat`,`destination`,`pay_price`,`surplus`,`destination_lon`,`destination_lat`) VALUES "+
		"(:guid,:username,:nickname,   :phone, :create_time, :travel_time,  :travel_time_title,:From, :from_lon, :from_lat,:destination, :pay_price, :surplus,  :destination_lon, :destination_lat)", &trip)
	if err == nil {
		redisManager.UpdateObject(trip.Guid, trip)
	}
	return trip, err

}
