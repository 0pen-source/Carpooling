package dao

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/0pen-source/Carpooling/models"
	"github.com/rs/xid"
)

const (
	Token_REDIS_KEY = "token"
)

// GetUser _
func GetUser(phone string) (user models.User, err error) {
	var userStr string
	query := "SELECT * FROM `user` where phone = ?"

	userStr, err = redisManager.GetKey(phone)
	fmt.Println(err)
	fmt.Println(userStr)
	if err != nil {
	}
	if userStr != "" {
		json.Unmarshal([]byte(userStr), &user)
	} else {
		err = cacheDB.Get(&user, query, phone)
		if err == nil {
			userbyte, _ := json.Marshal(user)
			redisManager.SetKey(phone, string(userbyte))
		}
	}

	return user, nil

}

// SaveUser _
func SaveUser(user models.User) (err error) {
	user.Guid = xid.New().String()
	_, err = cacheDB.NamedExec("INSERT INTO user (phone, password, nickname, guid) VALUES (:phone, :password, :nickname, :guid)", user)
	if err == nil {
		UpdateUserRedis(user.Phone)
	}
	return err

}

func SaveImage(path, filename, imageType, phone string) {

	PutObject(bucketIDCards, filename, fmt.Sprintf("%s/%s", path, filename))

	os.Remove(fmt.Sprintf("%s/%s", path, filename))
	user := models.User{
		Phone: phone,
	}
	if imageType == "idcard" {
		user.IDCardsURL = fmt.Sprintf("%s/%s", path, filename)
	} else if imageType == "driving" {
		user.DriverURL = fmt.Sprintf("%s/%s", path, filename)
	}
	SaveUser(user)

}

// UpdateUser _
func UpdateUser(user models.User) (err error) {
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

// UpdateUserRedis _
func UpdateUserRedis(phone string) {
	query := "SELECT * FROM `user` where phone = ?"
	var user models.User
	err := cacheDB.Get(&user, query, phone)
	if err == nil {
		userbyte, _ := json.Marshal(user)
		redisManager.SetKey(phone, string(userbyte))
	}

}

// SaveToken _
func SaveToken(phone, token string) {
	redisManager.SetKey(fmt.Sprintf("%s_%s", Token_REDIS_KEY, phone), token)

}

// GetToken _
func GetToken(phone string) (string, error) {
	return redisManager.GetKey(fmt.Sprintf("%s_%s", Token_REDIS_KEY, phone))

}
