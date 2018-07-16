package models

// Configuration model
type Configuration struct {
	Address             string `mapstructure:"address" validate:"required"`
	CarpoolingDatabases DB     `mapstructure:"carpooling_databases" validate:"required"`
	MODE                string `mapstructure:"mode" validate:"required"`
	Checkcode           string `mapstructure:"checkcode" validate:"required"`
	AccessKey           string `mapstructure:"accessKey" validate:"required"`
	Secret              string `mapstructure:"secret" validate:"required"`
	Sign                string `mapstructure:"sign" validate:"required"`
	TemplateID          string `mapstructure:"templateId" validate:"required"`
	VerificationCodeURL string `mapstructure:"verificationcodeurl" validate:"required"`
	Redis struct {
		URL      string `mapstructure:"url" validate:"required,redis_url"`
		PoolSize int    `mapstructure:"pool_size" validate:"required,min=1"`
	} `mapstructure:"redis"`
	OSSConfig struct {
		AccessKeyId     string `mapstructure:"AccessKeyId" validate:"required"`
		AccessKeySecret string `mapstructure:"AccessKeySecret" validate:"required"`
		EndPoint        string `mapstructure:"EndPoint" validate:"required"`
		BucketIDCards   string `mapstructure:"BucketIDCards" validate:"required"`
	} `mapstructure:"oss_config"`
}

// DB model
type DB struct {
	DataSourceName string `mapstructure:"dsn" validate:"required,dsn"`
	MaxOpenConns   int    `mapstructure:"max_open_conns" validate:"required,min=1"`
	MaxIdleConns   int    `mapstructure:"max_idle_conns" validate:"required,min=1,ltefield=MaxOpenConns"`
}
