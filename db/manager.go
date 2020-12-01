package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"time"

	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var masterDB *sql.DB
var slaveDB []*sql.DB

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

	db, err := r.SqlOpen(c.MSConfig.Master)
	if err != nil {
		return err
	}

	for i, v := range c.MSConfig.Replication {

		d, err := r.SqlOpen(v)
		if err != nil {
			fmt.Println("SqlOpen when slave ", i, " :", err)
			continue
		}

		slaveDB = append(slaveDB, d)
	}

	if len(slaveDB) == 0 {
		return fmt.Errorf("No valid slave db")
	}
	masterDB = db

	return nil

}

//Uninitialize is a
func (r *DBManager) Uninitialize() error {

	if masterDB != nil {
		masterDB.Close()
		masterDB = nil
	}

	for i := range slaveDB {
		if slaveDB[i] != nil {
			slaveDB[i].Close()
		}
	}

	slaveDB = nil

	return nil
}

func MapToQueryString(m map[string]string) string {

	params := url.Values{}
	for k, v := range m {
		params.Add(k, v)
	}

	return params.Encode()
}

func (r *DBManager) SqlOpen(c DBConfig) (*sql.DB, error) {

	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", 
	c.Server, c.Port, c.Username, c.Password, c.Database)

	d, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	err = d.Ping()
	if err != nil {
		return nil, err
	}

	d.SetConnMaxLifetime(time.Second * time.Duration(c.ConnMaxLifetime))
	d.SetMaxOpenConns(c.MaxOpenConns)
	d.SetMaxIdleConns(c.MaxIdleConns)

	return d, nil
}


func (r *DBManager) MySqlOpen(c DBConfig) (*sql.DB, error) {

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		c.Username, c.Password,
		c.Server, c.Port, c.Database)

	baseURL, err := url.Parse(dataSourceName)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	baseURL.RawQuery = MapToQueryString(c.Parameter)
	dataSource := fmt.Sprint(baseURL)

	d, err := sql.Open("mysql", dataSource)
	if err != nil {
		return nil, err
	}

	err = d.Ping()
	if err != nil {
		return nil, err
	}

	d.SetConnMaxLifetime(time.Second * time.Duration(c.ConnMaxLifetime))
	d.SetMaxOpenConns(c.MaxOpenConns)
	d.SetMaxIdleConns(c.MaxIdleConns)

	return d, nil
}
