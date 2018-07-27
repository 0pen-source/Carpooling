package dao

import (
	"time"

	"github.com/0pen-source/Carpooling/models"
)

// GetRealTimeTrip _
func GetBanner() (banners []models.ResponseBanner, err error) {
	query := "SELECT * from banner ORDER BY  create_time desc limit 10;"

	banners, ok := memCache.Get(query).([]models.ResponseBanner)
	if !ok {
		err = cacheDB.Select(&banners, query)
	}
	memCache.Put(query, banners, time.Hour*1)

	return banners, nil

}
