package dao

import (
	"fmt"
	"strconv"
	"time"

	"github.com/0pen-source/Carpooling/models"
	"github.com/rs/xid"
)

// GetRealTimeDriverTrip _
func GetRealTimeDriverTrip() (trips []models.ResponseTrip, err error) {
	query := "SELECT * from driver_trip WHERE travel_time>= ? group by phone order by create_time desc  limit 20;"

	trips, ok := memCache.Get(query).([]models.ResponseTrip)
	if !ok {
		err = cacheDB.Select(&trips, query,strconv.FormatInt(time.Now().UnixNano()/1e6, 10))
	}
	memCache.Put(query, trips, time.Hour*1)

	return trips, nil

}
func GetRecommendDriverTrips(user models.User) (trips []models.ResponseTrip, err error) {
	//query := "SELECT *,ROUND(6378.138 * 2 * ASIN(SQRT(POW(SIN((? * PI() / 180 - from_lat * PI() / 180) / 2),2) + COS(? * PI() / 180) * COS(from_lat * PI() / 180) * POW(SIN((? * PI() / 180 - from_lon * PI() / 180) / 2), 2))) * 1000) AS distance FROM driver_trip where travel_time>=unix_timestamp(date_add(now(), interval -1 day) ) ORDER BY distance ASC limit 20"
	query := "SELECT * from driver_trip WHERE travel_time>= ? group by phone order by create_time desc  limit 20;"
	//trips, ok := memCache.Get(fmt.Sprintf("%s-%s-%s", user.LastLat, user.LastLon, "driver")).([]models.ResponseTrip)
	//if !ok {
	err = cacheDB.Select(&trips, query,strconv.FormatInt(time.Now().UnixNano()/1e6, 10))
	//}
	//memCache.Put(fmt.Sprintf("%s-%s-%s", user.LastLat, user.LastLon, "driver"), trips, time.Minute*10)
	fmt.Println(err)
	return trips, nil

}

func GetSearchDriverTrips(trip models.PassengersTrip) (trips []models.ResponseTrip, err error) {
	query := "SELECT *,((ROUND(6378.138 * 2 * ASIN(SQRT(POW(SIN((? * PI() / 180 - from_lat * PI() / 180) / 2),2) + COS(? * PI() / 180) * COS(from_lat * PI() / 180) * POW(SIN((? * PI() / 180 - from_lon * PI() / 180) / 2), 2))) * 1000) ) + (ROUND(6378.138 * 2 * ASIN(SQRT(POW(SIN((? * PI() / 180 - destination_lat * PI() / 180) / 2),2) + COS(? * PI() / 180) * COS(destination_lat * PI() / 180) * POW(SIN((? * PI() / 180 - destination_lon * PI() / 180) / 2), 2))) * 1000) )) AS distance  FROM driver_trip WHERE travel_time>= ? and surplus>0 and distance <=50000 group by phone  ORDER BY distance ASC limit 20"

	err = cacheDB.Select(&trips, query, trip.FromLat, trip.FromLat, trip.FromLon, trip.DestinationLat, trip.DestinationLat, trip.DestinationLon, trip.TravelTime)
	fmt.Println(err)
	return trips, nil

}

// SaveDriverTrip _
func SaveDriverTrip(trip models.DriverTrip) (models.DriverTrip, error) {
	trip.Guid = xid.New().String()

	_, err = cacheDB.NamedExec("INSERT INTO driver_trip "+
		"(`guid`,`username`,`nickname`,`phone`,`create_time`,`travel_time`,`travel_time_title`,`from_lon`,`from_lat`,`pay_price`,`surplus`,`destination_lon`,`destination_lat`"+
		",`from_region`,`from_city`,`from_accurate_address`,`from_vague_address`,`destination_region`,`destination_city`,`destination_accurate_address`,`destination_vague_address`,`source`,`mileage`,`seat_num`,`complete`,`msg`,`from_district`,`destination_district`,`portrait_url`) VALUES "+
		"(:guid,:username,:nickname,:phone,:create_time, :travel_time,  :travel_time_title, :from_lon, :from_lat, :pay_price, :surplus,  :destination_lon, :destination_lat"+
		",:from_region,:from_city,:from_accurate_address,:from_vague_address,:destination_region,:destination_city,:destination_accurate_address,:destination_vague_address,:source,:mileage,:seat_num,:complete,:msg,:from_district,:destination_district,:portrait_url)", &trip)
	if err == nil {
		redisManager.UpdateObject(trip.Guid, trip)
	}
	fmt.Println(err)
	return trip, err

}
func GetMyDriverTrip(user models.User) (trips []models.ResponseTrip, err error) {
	query := "SELECT * FROM passengers_trip where phone=?  ORDER BY create_time desc limit 20"

	trips, ok := memCache.Get(fmt.Sprintf("%s-%s", user.Phone, "my_trip_driver")).([]models.ResponseTrip)
	if !ok {
		err = cacheDB.Select(&trips, query, user.Phone)
	}
	fmt.Println(err)
	memCache.Put(fmt.Sprintf("%s-%s", user.Phone, "my_trip_driver"), trips, time.Minute*10)

	return trips, nil

}
func GetPhoneBYDriveerTripGUID(message models.Connected) (trips []models.ResponseConnected, err error) {
	query := "SELECT * FROM driver_trip where guid=?  "
	err = cacheDB.Select(&trips, query, message.Guid)
	fmt.Println(err)
	return trips, nil

}
