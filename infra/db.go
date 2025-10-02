package infra

import (
	"fmt"

	"github.com/ian0113/go-gin-mvc/config"
	"github.com/ian0113/go-gin-mvc/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	globalDB *gorm.DB
)

func NewDB(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.HostName,
		cfg.Database.HostPort,
		cfg.Database.Database,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	models := []interface{}{
		&models.User{},
		&models.Order{},
	}
	for _, model := range models {
		err := db.AutoMigrate(model)
		if err != nil {
			panic(err)
		}
	}
	return db
}

func InitDB(cfg *config.Config) {
	globalDB = NewDB(cfg)
}

func GetDB() *gorm.DB {
	return globalDB
}
