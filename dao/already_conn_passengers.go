package dao

import (
	"fmt"
	"time"

	"github.com/0pen-source/Carpooling/models"
)

func GetConnPassengers(user models.User) (trips []models.ResponseTrip, err error) {
	query := "SELECT * from already_conn_driver INNER join driver_trip on already_conn_driver.guid=driver_trip.guid  where already_conn_driver.phone = ?"
	trips, ok := memCache.Get(fmt.Sprintf("%s-%s", user.Phone, "already_conn_driver")).([]models.ResponseTrip)
	if !ok {
		err = cacheDB.Select(&trips, query, user.Phone)
	}

	memCache.Put(fmt.Sprintf("%s-%s", user.Phone, "already_conn_driver"), trips, time.Minute*10)

	return trips, nil

}

// SaveConnDriver _
func SaveConnPassengers(connDriver models.AlreadyConnDriver) (error) {

	_, err = cacheDB.NamedExec("INSERT INTO already_conn_driver (phone,guid,update_time) VALUES (:phone,:guid,:update_time) ON DUPLICATE KEY UPDATE update_time=:update_time", &connDriver)
	return err

}
