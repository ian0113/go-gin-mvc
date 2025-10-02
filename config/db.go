package config

type DBConfig struct {
	HostName string
	HostPort uint16
	Username string
	Password string
	Database string
}

var (
	globalDefaultDBConfig = DBConfig{
		HostName: "mysql",
		HostPort: 3306,
		Username: "admin",
		Password: "admin",
		Database: "defaultdb",
	}
	_ SubConfigInf = &DBConfig{}
)

func (x *DBConfig) Default() {
	*x = globalDefaultDBConfig
}

func (x *DBConfig) Restore() {
	// TODO: fix config
}
