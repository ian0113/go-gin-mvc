package config

import "github.com/ian0113/go-gin-mvc/utils"

type AppMode uint

const (
	AppModeProduction AppMode = iota
	AppModeDevelopment
	_AppModeEnd
)

type AppConfig struct {
	Name     string
	Mode     AppMode
	HostName string
	HostPort uint16
}

var (
	globalDefaultAppConfig = AppConfig{
		Name:     "TEST_PROJECT",
		Mode:     AppModeProduction,
		HostName: "0.0.0.0",
		HostPort: 8080,
	}
	_ SubConfigInf = &AppConfig{}
)

func (x *AppConfig) Default() {
	*x = globalDefaultAppConfig
}

func (x *AppConfig) Restore() {
	if x.Name == "" {
		x.Name = globalDefaultAppConfig.Name
	}
	if x.Mode >= _AppModeEnd {
		x.Mode = globalDefaultAppConfig.Mode
	}
	if x.HostName != "" && !(utils.IsIP(x.HostName) || utils.IsHostname(x.HostName)) {
		x.HostName = globalDefaultAppConfig.HostName
	}
	if x.HostPort <= 1000 {
		x.HostPort = globalDefaultAppConfig.HostPort
	}
}
