package dao

import (
	"fmt"
	"time"

	"github.com/0pen-source/Carpooling/models"
	"github.com/rs/xid"
)

// GetRealTimeTrip _
func GetRealTimePassengersTrip() (trips []models.ResponseTrip, err error) {
	query := "SELECT * FROM `passengers_trip` order by create_time desc limit 20"

	trips, ok := memCache.Get(query).([]models.ResponseTrip)
	if !ok {
		err = cacheDB.Select(&trips, query)
	}
	memCache.Put(query, trips, time.Hour*1)

	return trips, nil

}

//select * from users_location where latitude > '.$lat.'-1 and latitude < '.$lat.'+1 and longitude > '.$lon.'-1 and longitude < '.$lon.'+1 order by ACOS(SIN(('.$lat.' * 3.1415) / 180 ) *SIN((latitude * 3.1415) / 180 ) +COS(('.$lat.' * 3.1415) / 180 ) * COS((latitude * 3.1415) / 180 ) *COS(('.$lon.'* 3.1415) / 180 - (longitude * 3.1415) / 180 ) ) * 6380 asc limit 10
// GetPassengersTrip _
func GetRecommendPassengersTrips(user models.User) (trips []models.ResponseTrip, err error) {
	query := "SELECT *," +
		"ROUND(6378.138 * 2 * ASIN(SQRT(POW(SIN((" +
		":last_lat * PI() / 180 - from_lat * PI() / 180) / 2),2" +
		") + COS(:last_lat * PI() / 180) * COS(from_lat * PI() / 180) * POW(SIN((:last_lon * PI() / 180 " +
		"- from_lon * PI() / 180) / 2), 2))) * 1000) AS juli FROM passengers_trip ORDER BY juli ASC limit 20"

	trips, ok := memCache.Get(fmt.Sprintf("%s-%s-%s", user.LastLat, user.LastLon, "driver")).([]models.ResponseTrip)
	if !ok {
		err = cacheDB.Select(&trips, query, user)
	}
	memCache.Put(fmt.Sprintf("%s-%s-%s", user.LastLat, user.LastLon, "driver"), trips, time.Minute*10)

	return trips, nil

}

// SavePassengersTrip _ , ,`from`
//  :from , ,
func SavePassengersTrip(trip models.PassengersTrip) (models.PassengersTrip, error) {
	trip.Guid = xid.New().String()
	_, err = cacheDB.NamedExec("INSERT INTO passengers_trip "+
		"(`guid`,`username`,`nickname`,`phone`,`create_time`,`travel_time`,`travel_time_title`,`from`,`from_lon`,`from_lat`,`destination`,`pay_price`,`surplus`,`destination_lon`,`destination_lat`) VALUES "+
		"(:guid,:username,:nickname,   :phone, :create_time, :travel_time,  :travel_time_title,:From, :from_lon, :from_lat,:destination, :pay_price, :surplus,  :destination_lon, :destination_lat)", &trip)
	if err == nil {
		redisManager.UpdateObject(trip.Guid, trip)
	}
	return trip, err

}
