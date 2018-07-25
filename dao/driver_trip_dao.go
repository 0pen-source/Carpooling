package dao

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/0pen-source/Carpooling/models"
	"github.com/rs/xid"
)

// GetDriverTrip _
func GetDriverTrip(phone string) (user models.User, err error) {
	var userStr string
	query := "SELECT * FROM `user` where phone = ?"

	userStr, err = redisManager.GetKey(phone)
	if err != nil {
		err = cacheDB.Get(&user, query, phone)
		if err == nil {
			userbyte, _ := json.Marshal(user)
			redisManager.SetKey(phone, string(userbyte))
		}
		return user, err

	}

	json.Unmarshal([]byte(userStr), &user)
	return user, nil

}

// SaveDriverTrip _ , ,`from`
//  :from , ,
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

// UpdateDriverTrip _
func UpdateDriverTrip(user models.User) (err error) {
	sql := "UPDATE `user` set "
	core_strings := []string{}

	if user.Username != "" {
		core_strings = append(core_strings, "username = :username")
	}
	if user.Password != "" {
		core_strings = append(core_strings, "password = :password")
	}
	if user.Nickname != "" {
		core_strings = append(core_strings, "nickname = :nickname")
	}
	if user.Sex != 0 {
		core_strings = append(core_strings, "sex = :sex")
	}
	if user.LastLocation != "" {
		core_strings = append(core_strings, "last_location = :last_location")
	}
	sql += strings.Join(core_strings, ",")
	sql += " where phone = :phone"
	fmt.Println(sql)
	_, err = cacheDB.NamedExec(sql, user)
	UpdateUserRedis(user.Phone)
	fmt.Println(err)
	return err

}

// UpdateDriverTripRedis _
func UpdateDriverTripRedis(phone string) {
	query := "SELECT * FROM `user` where phone = ?"
	var user models.User
	err := cacheDB.Get(&user, query, phone)
	if err == nil {
		userbyte, _ := json.Marshal(user)
		redisManager.SetKey(phone, string(userbyte))
	}

}
