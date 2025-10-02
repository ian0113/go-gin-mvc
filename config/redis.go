package config

type RedisConfig struct {
	HostName string
	HostPort uint16
	Password string
	DB       int
}

var (
	globalDefaultRedisConfig = RedisConfig{
		HostName: "redis",
		HostPort: 6379,
		Password: "mypassword",
		DB:       0,
	}
	_ SubConfigInf = &RedisConfig{}
)

func (x *RedisConfig) Default() {
	*x = globalDefaultRedisConfig
}

func (x *RedisConfig) Restore() {
	if x.HostPort == 0 {
		x.HostPort = globalDefaultRedisConfig.HostPort
	}
}
