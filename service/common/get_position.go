package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/0pen-source/Carpooling/models"
)

func GetPosition(lat, lon float64) (postion models.Position) {
	resp, err := http.Get(fmt.Sprintf("http://api.map.baidu.com/geocoder/v2/?&location=%v,%v&output=json&pois=1&ak=VjLlIwbOdpWeQlNyIudaYvVZ4vwGef1q", lat, lon))
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	fmt.Println(string(body))
	json.Unmarshal(body, &postion)
	return
}
