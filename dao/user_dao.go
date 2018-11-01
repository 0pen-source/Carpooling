package dao

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/0pen-source/Carpooling/models"
	"github.com/rs/xid"
)

const (
	Token_REDIS_KEY = "token"
)

var USERS_PORTRAIT = []string{
	"https://users-portrait.oss-cn-beijing.aliyuncs.com/25542158565e1df6f1f4dfe196f68263.jpg",
	"https://users-portrait.oss-cn-beijing.aliyuncs.com/46e6a933d8149f7c03e4f900431a619c.jpeg",
	"https://users-portrait.oss-cn-beijing.aliyuncs.com/8e92f8f6edd94a6c5b26db8ce8f89731.jpeg",
	"https://users-portrait.oss-cn-beijing.aliyuncs.com/head1.jpg",
	"https://users-portrait.oss-cn-beijing.aliyuncs.com/head2.jpg",
	"https://users-portrait.oss-cn-beijing.aliyuncs.com/head3.jpg",
	"https://users-portrait.oss-cn-beijing.aliyuncs.com/head4.jpg",
	"https://users-portrait.oss-cn-beijing.aliyuncs.com/head5.jpg",
	"https://users-portrait.oss-cn-beijing.aliyuncs.com/head6.jpg",
	"https://users-portrait.oss-cn-beijing.aliyuncs.com/head7.jpg",
}

// GetUser _
func GetUser(phone string) (user models.User, err error) {
	var userStr string
	query := "SELECT * FROM `user` where phone = ?"

	userStr, err = redisManager.GetKey(phone)
	if userStr != "" {
		json.Unmarshal([]byte(userStr), &user)
	} else {
		err = cacheDB.Get(&user, query, phone)
		if err == nil {
			userbyte, _ := json.Marshal(user)
			redisManager.SetKey(phone, string(userbyte))
		}
	}
	fmt.Println(user)
	return user, err

}

// SaveUser _
func SaveUser(user models.User) (error, models.User) {
	user.Guid = xid.New().String()
	if user.Nickname == user.Phone || user.Nickname == "" {
		user.Nickname = strings.Join([]string{user.Phone[:4], "***", user.Phone[7:]}, "")
	}
	if user.PortraitURL == "" {
		n := rand.Intn(len(USERS_PORTRAIT))
		user.PortraitURL = USERS_PORTRAIT[n]
	}
	_, err := cacheDB.NamedExec("INSERT INTO user (phone, password, nickname, guid,portrait_url) VALUES (:phone, :password, :nickname, :guid,:portrait_url)", user)
	if err == nil {
		UpdateUserRedis(user.Phone)
	}
	return err, user

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
	if user.LastLon > 0 {
		core_strings = append(core_strings, "last_lon = :last_lon")
	}
	if user.LastLat > 0 {
		core_strings = append(core_strings, "last_lat = :last_lat")
	}
	sql += strings.Join(core_strings, ",")
	sql += " where phone = :phone"
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
