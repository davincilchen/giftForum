package config

import (
	"encoding/json"
)

type GConfig struct {
	Domain string
}
type Config struct {
	GConfig GConfig `json:"general" env:"general"`
}

var gConfig GConfig

func Initialize(buf []byte) error {

	c := Config{}
	err := json.Unmarshal(buf, &c)
	if err != nil {
		return err
	}

	gConfig = c.GConfig

	return nil
}

func Uninitialize() error {

	return nil

}

func GetDomain() string {
	return gConfig.Domain
}
