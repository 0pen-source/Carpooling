package models

// Configuration model
type Configuration struct {
	Address             string `mapstructure:"address" validate:"required"`
	CarpoolingDatabases DB     `mapstructure:"carpooling_databases" validate:"required"`
	MODE                string `mapstructure:"mode" validate:"required"`
	Checkcode           string `mapstructure:"checkcode" validate:"required"`
}

// DB model
type DB struct {
	DataSourceName string `mapstructure:"dsn" validate:"required,dsn"`
	MaxOpenConns   int    `mapstructure:"max_open_conns" validate:"required,min=1"`
	MaxIdleConns   int    `mapstructure:"max_idle_conns" validate:"required,min=1,ltefield=MaxOpenConns"`
}
