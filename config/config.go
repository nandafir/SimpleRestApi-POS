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
	DBName            string        `envconfig:"DB" default:"simple_pos"`
	DBHost            string        `envconfig:"DB_HOST" default:"localhost"`
	DBPort            int           `envconfig:"DB_HOST" default:"3306"`
	DBUsername        string        `envconfig:"DB_USERNAME" default:"debian-sys-maint"`
	DBPassword        string        `envconfig:"DB_PASS" default:"uCSlljDvEyR7YJ0H"`
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
