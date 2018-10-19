package dao

import (
	"fmt"
	"time"

	"github.com/0pen-source/Carpooling/models"
	"github.com/rs/xid"
)

// GetRealTimeTrip _
func GetRealTimePassengersTrip() (trips []models.ResponseTrip, err error) {
	query := "SELECT * FROM `passengers_trip` group by phone order by create_time desc limit 20"

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
	//query := "SELECT *," +
	//	"ROUND(6378.138 * 2 * ASIN(SQRT(POW(SIN((" +
	//	":last_lat * PI() / 180 - from_lat * PI() / 180) / 2),2" +
	//	") + COS(:last_lat * PI() / 180) * COS(from_lat * PI() / 180) * POW(SIN((:last_lon * PI() / 180 " +
	//	"- from_lon * PI() / 180) / 2), 2))) * 1000) AS juli FROM passengers_trip where travel_time>=unix_timestamp(date_add(now(), interval -1 day) )  ORDER BY juli ASC limit 20"
	query := "SELECT * FROM `passengers_trip` group by phone order by create_time desc limit 20"
	trips, ok := memCache.Get(fmt.Sprintf("%s-%s-%s", user.LastLat, user.LastLon, "driver")).([]models.ResponseTrip)
	if !ok {
		err = cacheDB.Select(&trips, query, user)
	}
	memCache.Put(fmt.Sprintf("%s-%s-%s", user.LastLat, user.LastLon, "driver"), trips, time.Minute*10)

	return trips, nil

}

func GetSearchPassengersTrips(trip models.PassengersTrip) (trips []models.ResponseTrip, err error) {
	query := "SELECT *,((ROUND(6378.138 * 2 * ASIN(SQRT(POW(SIN((? * PI() / 180 - from_lat * PI() / 180) / 2),2) + COS(? * PI() / 180) * COS(from_lat * PI() / 180) * POW(SIN((? * PI() / 180 - from_lon * PI() / 180) / 2), 2))) * 1000) ) + (ROUND(6378.138 * 2 * ASIN(SQRT(POW(SIN((? * PI() / 180 - from_lat * PI() / 180) / 2),2) + COS(? * PI() / 180) * COS(from_lat * PI() / 180) * POW(SIN((? * PI() / 180 - from_lon * PI() / 180) / 2), 2))) * 1000) )) AS distance  FROM passengers_trip WHERE travel_time>= ? group by phone   ORDER BY distance ASC limit 20"

	err = cacheDB.Select(&trips, query, trip.FromLat, trip.FromLat, trip.FromLon, trip.DestinationLat, trip.DestinationLat, trip.DestinationLon, trip.TravelTime)

	return trips, nil

}

// SavePassengersTrip _
//
func SavePassengersTrip(trip models.PassengersTrip) (models.PassengersTrip, error) {
	trip.Guid = xid.New().String()
	_, err = cacheDB.NamedExec("INSERT INTO passengers_trip "+
		"(`guid`,`username`,`nickname`,`phone`,`create_time`,`travel_time`,`travel_time_title`,`from_lon`,`from_lat`,`pay_price`,`surplus`,`destination_lon`,`destination_lat`"+
		",`from_region`,`from_city`,`from_accurate_address`,`from_vague_address`,`destination_region`,`destination_city`,`destination_accurate_address`,`destination_vague_address`,`source`,`mileage`,`seat_num`,`complete`,`msg`) VALUES "+
		"(:guid,:username,:nickname,:phone,:create_time, :travel_time,  :travel_time_title, :from_lon, :from_lat, :pay_price, :surplus,  :destination_lon, :destination_lat"+
		",:from_region,:from_city,:from_accurate_address,:from_vague_address,:destination_region,:destination_city,:destination_accurate_address,:destination_vague_address,:source,:mileage,:seat_num,:complete,:msg)", &trip)
	if err == nil {
		redisManager.UpdateObject(trip.Guid, trip)
	}
	return trip, err

}
func GetMyTrip(user models.User) (trips []models.ResponseTrip, err error) {
	query := "SELECT * FROM passengers_trip where phone=?  ORDER BY create_time desc limit 20"

	trips, ok := memCache.Get(fmt.Sprintf("%s-%s", user.Phone, "my_trip_passenger")).([]models.ResponseTrip)
	if !ok {
		err = cacheDB.Select(&trips, query, user.Phone)
	}
	fmt.Println(err)
	memCache.Put(fmt.Sprintf("%s-%s", user.Phone, "my_trip_passenger"), trips, time.Minute*10)

	return trips, nil

}
func GetPhoneBYGUID(message models.Connected) (trips []models.ResponseConnected, err error) {
	query := "SELECT * FROM passengers_trip where guid=? "
	fmt.Println("GetPhoneBYGUID---", err)
	err = cacheDB.Select(&trips, query, message.Guid)
	fmt.Println("GetPhoneBYGUID", err)
	return trips, nil

}
