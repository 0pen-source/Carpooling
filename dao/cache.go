package dao

import (
	_ "database/sql"

	"github.com/astaxie/beego/cache"
	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
)

var (
	memCache                 cache.Cache
	cacheDB                  *sqlx.DB
	CACHE_KEY_GEO_PREIFX     = `cache_tvgeo:`
	CACHE_KEY_GRID_PREIFX    = `cache_tvgrid:`
	CACHE_KEY_GEOHASH_PREIFX = `cache_tvgeohash:`
)

func init() {
	memCache, _ = cache.NewCache("memory", `{"interval":5}`)
}
func InitializeCache() {
	cacheDB = sqlx.MustConnect("mysql", Config.CarpoolingDatabases.DataSourceName).Unsafe()
	cacheDB.SetMaxOpenConns(Config.CarpoolingDatabases.MaxIdleConns)
	cacheDB.SetMaxIdleConns(Config.CarpoolingDatabases.MaxIdleConns)
}
