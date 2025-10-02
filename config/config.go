package config

import (
	"log"
	"reflect"
)

type SubConfigInf interface {
	Default()
	Restore()
}

type Config struct {
	App         AppConfig
	Database    DBConfig
	Redis       RedisConfig
	AuthService AuthServiceConfig
}

func (x *Config) Restore() {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if !field.CanAddr() {
			log.Printf("%s has been skipped\n", field.Type().Name())
			continue
		}
		if restorer, ok := field.Addr().Interface().(SubConfigInf); ok {
			restorer.Restore()
		} else {
			log.Printf("type %s does not implement interface X (missing method Restore)\n", field.Type().Name())
		}
	}
}

func (x *Config) Default() {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if !field.CanAddr() {
			log.Printf("%s has been skipped\n", field.Type().Name())
			continue
		}
		if fieldCfg, ok := field.Addr().Interface().(SubConfigInf); ok {
			fieldCfg.Default()
		} else {
			log.Printf("type %s does not implement interface X (missing method Default)\n", field.Type().Name())
		}
	}
}
