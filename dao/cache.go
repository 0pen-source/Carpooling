package dao

import (
	_ "database/sql"

	"github.com/astaxie/beego/cache"
	_ "github.com/go-sql-driver/mysql"

	"github.com/0pen-source/Carpooling/models"
	"github.com/jmoiron/sqlx"
)

var (
	cacheDB                  *sqlx.DB
	mem_cache                cache.Cache
	CACHE_KEY_GEO_PREIFX     = `cache_tvgeo:`
	CACHE_KEY_GRID_PREIFX    = `cache_tvgrid:`
	CACHE_KEY_GEOHASH_PREIFX = `cache_tvgeohash:`
)

func init() {
	mem_cache, _ = cache.NewCache("memory", `{"interval":5}`)
}
func InitializeCache(db models.DB) {
	cacheDB = sqlx.MustConnect("mysql", db.DataSourceName).Unsafe()
	cacheDB.SetMaxOpenConns(db.MaxIdleConns)
	cacheDB.SetMaxIdleConns(db.MaxIdleConns)
}
