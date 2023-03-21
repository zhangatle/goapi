package app

import (
	"goapi/pkg/config"
	"time"
)

func IsLocal() bool {
	return config.Get("app.env") == "local"
}

func IsProduction() bool {
	return config.Get("app.env") == "production"
}

func IsTesting() bool {
	return config.Get("app.env") == "testing"
}

// TimeNowInTimezone 获取当前时间，支持时区
func TimeNowInTimezone() time.Time {
	chinaTimezone, _ := time.LoadLocation(config.Get("app.timezone"))
	return time.Now().In(chinaTimezone)
}

func URL(path string) string {
	return config.Get("app.url") + path
}

func V1URL(path string) string {
	return URL("/v1/" + path)
}
