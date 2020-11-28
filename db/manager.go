package db

import (
	"encoding/json"

	"github.com/davecgh/go-spew/spew"
)

type DBManager struct {
}

type DBConfig struct {
	Server          string            `json:"server" env:"server"`
	Port            json.Number       `json:"port" env:"port"`
	Database        string            `json:"db" env:"db"`
	Username        string            `json:"username" env:"username"`
	Password        string            `json:"password" env:"password"`
	Parameter       map[string]string `json:"parameter" env:"parameter"`
	MaxIdleConns    int               `json:"maxidleconns" env:"maxidleconns"`
	MaxOpenConns    int               `json:"maxopenconns" env:"maxopenconns"`
	ConnMaxLifetime int               `json:"connmaxlifetime" env:"connmaxlifetime"`
}

type MSConfig struct {
	Master      DBConfig   `json:"master" env:"master"`
	Replication []DBConfig `json:"replication" env:"replication"`
}

type Config struct {
	MSConfig MSConfig `json:"db" env:"db"`
}

func (r *DBManager) Initialize(buf []byte) error {

	c := Config{}
	err := json.Unmarshal(buf, &c)
	if err != nil {
		return err
	}

	spew.Dump(c)

	return nil

}

//Uninitialize is a
func (r *DBManager) Uninitialize() error {
	return nil
}
