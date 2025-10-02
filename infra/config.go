package infra

import (
	"fmt"
	"os"

	"github.com/ian0113/go-gin-mvc/config"

	"github.com/goccy/go-yaml"
)

var (
	globalConfig *config.Config
)

func NewConfig(path, cfgName string) *config.Config {
	cfg := &config.Config{}
	filePath := fmt.Sprintf("%s/%s.yaml", path, cfgName)
	data, err := os.ReadFile(filePath)
	if err != nil {
		cfg.Default()
		err = writeConfig(path, cfgName, cfg)
		if err != nil {
			panic(err)
		}
		return cfg
	}
	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		cfg.Default()
		err = writeConfig(path, cfgName, cfg)
		if err != nil {
			panic(err)
		}
		return cfg
	}
	cfg.Restore()
	err = writeConfig(path, cfgName, cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}

func InitConfig(path, cfgName string) {
	globalConfig = NewConfig(path, cfgName)
}

func GetConfig() *config.Config {
	return globalConfig
}

func writeConfig(path, cfgName string, cfg *config.Config) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		if mkErr := os.MkdirAll(path, 0755); mkErr != nil {
			return mkErr
		} else {
			return err
		}
	}
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	filePath := fmt.Sprintf("%s/%s.yaml", path, cfgName)
	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
