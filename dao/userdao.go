package dao

import (
	"encoding/json"
	"fmt"

	"github.com/0pen-source/Carpooling/models"
	"github.com/rs/xid"
)

// GetUser _
func GetUser(phone string) (user models.User, err error) {
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

// SaveUser _
func SaveUser(user models.User) (err error) {
	user.Guid = xid.New().String()
	_, err = cacheDB.NamedExec("INSERT INTO user (phone, password, nickname,guid) VALUES (:phone, :password, :nickname, :guid)", user)
	fmt.Println(err)
	if err == nil {
		userbyte, _ := json.Marshal(user)
		redisManager.SetKey(user.Phone, string(userbyte))
	}
	return err

}
