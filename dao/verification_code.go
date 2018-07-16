package dao

import "fmt"

const (
	CODE_REDIS_KEY = "code"
)

func GetCode(phone string) (string, error) {
	return redisManager.GetKey(fmt.Sprintf("%s_%s", CODE_REDIS_KEY, phone))

}
func SaveCode(phone, code string) {
	redisManager.SetKey(fmt.Sprintf("%s_%s", CODE_REDIS_KEY, phone), code)

}
