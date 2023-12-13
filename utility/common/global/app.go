package global

import (
	"sync"
	"time"
)

type app struct {
	Name    string
	Build   string
	Version string
	Date    time.Time

	Copyright string

	// 启动时间
	LaunchTime time.Time
	Uptime     time.Duration

	Env string

	Host string
	Port string

	BaseURL string

	// CDN 资源域名
	CDNHttp  string
	CDNHttps string

	Domain string

	locker sync.Mutex
}

var App = &app{}
