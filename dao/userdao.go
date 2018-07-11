package dao

import (
	"encoding/json"
	"time"

	"github.com/0pen-source/Carpooling/models"
)

// GetUser _
func GetUser(phone string) (user models.User, err error) {
	var userStr string
	query := "SELECT * FROM `user` where phone = ?"

	userStr, err = redisManager.GetKey(phone)
	if err != nil {
		err = cacheDB.Get(&user, query, phone)
		if err == nil {
			mem_cache.Put(phone, user, time.Second*5)
			userbyte, _ := json.Marshal(user)
			redisManager.SetKey(phone, string(userbyte))
		}
		return user, err

	}

	json.Unmarshal([]byte(userStr), &user)
	return user, nil

}
