package dao

import (
	_ "database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/0pen-source/Carpooling/models"
	"github.com/jmoiron/sqlx"
)

var cache *sqlx.DB

func InitializeCache(db models.DB) {
	cache = sqlx.MustConnect("mysql", db.DataSourceName).Unsafe()
	cache.SetMaxOpenConns(db.MaxIdleConns)
	cache.SetMaxIdleConns(db.MaxIdleConns)
}
