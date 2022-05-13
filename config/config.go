package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

var (
	conf *Config
)

// Config ...
type Config struct {
	Port              int           `envconfig:"PORT" default:"8080"`
	DBName            string        `envconfig:"PRODUCT_DB" default:"local_sample"`
	DBHost            string        `envconfig:"PRODUCT_DB_HOST" default:"localhost"`
	DBPort            int           `envconfig:"PRODUCT_DB_HOST" default:"5432"`
	DBUsername        string        `envconfig:"PRODUCT_DB_USERNAME" default:"postgres"`
	DBPassword        string        `envconfig:"PRODUCT_DB_PASS" default:"password"`
	DBConnMaxLifetime time.Duration `envconfig:"DB_CONN_MAX_LIFETIME" default:"1m"`
	DBConnMaxIdleTime time.Duration `envconfig:"DB_CONN_MAX_IDLE_TIME" default:"1m"`
	DBConnMaxIdle     int           `envconfig:"DB_CONN_MAX_IDLE" default:"10"`
	DBConnMaxOpen     int           `envconfig:"DB_CONN_MAX_OPEN" default:"50"`

	JWTSecret          string        `envconfig:"JWT_SECRET" default:"simplerestapi"`
	JWTExpiredDuration time.Duration `envconfig:"JWT_EXPIRED_DURATION" default:"168h"`

	ContentHost string `envconfig:"CONTENT_HOST" default:"http://10.64.21.25:30001"`

	LogLevel string `envconfig:"LOG_LEVEL" default:"debug"`
	LogDSN   string `envconfig:"LOG_DSN" default:""`
}

// Init ...
func Init() {
	conf = new(Config)
	envconfig.MustProcess("", conf)
}

// Get ...
func Get() *Config {
	return conf
}
