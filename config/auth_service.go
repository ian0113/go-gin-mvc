package config

import "time"

type AuthServiceConfig struct {
	RefreshTokenExpiration time.Duration
	AccessTokenExpiration  time.Duration
}

var (
	globalDefaultAuthServiceConfig = AuthServiceConfig{
		RefreshTokenExpiration: 7 * 24 * time.Hour,
		AccessTokenExpiration:  15 * time.Minute,
	}
	_ SubConfigInf = &AuthServiceConfig{}
)

func (x *AuthServiceConfig) Default() {
	*x = globalDefaultAuthServiceConfig
}

func (x *AuthServiceConfig) Restore() {
	if x.RefreshTokenExpiration == 0 {
		x.RefreshTokenExpiration = globalDefaultAuthServiceConfig.RefreshTokenExpiration
	}
	if x.AccessTokenExpiration == 0 {
		x.AccessTokenExpiration = globalDefaultAuthServiceConfig.AccessTokenExpiration
	}
}
