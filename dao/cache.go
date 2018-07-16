package dao

import (
	_ "database/sql"

	"github.com/astaxie/beego/cache"
	_ "github.com/go-sql-driver/mysql"

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
func InitializeCache() {
	cacheDB = sqlx.MustConnect("mysql", config.CarpoolingDatabases.DataSourceName).Unsafe()
	cacheDB.SetMaxOpenConns(config.CarpoolingDatabases.MaxIdleConns)
	cacheDB.SetMaxIdleConns(config.CarpoolingDatabases.MaxIdleConns)
}
