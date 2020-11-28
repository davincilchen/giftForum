package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"sync/atomic"

	api "giftForum/api"
	"giftForum/db"
)

const confPath = "./config.json"

var rowConfig []byte

type Config struct {
	Port string
}

// Server is a
type Server struct {
	Config Config `json:"WebServer"`

	dbManager *db.DBManager
	isClosed  int32 //non-zero: we're in Close
}

// NewServer create a new server from config.
func NewServer() *Server {

	return &Server{
		Config: Config{
			Port: ":8080",
		},
	}
}

func (g *Server) Initialize() error {

	buf, err := g.loadConfig()
	if err != nil {
		return err
	}

	d := db.DBManager{}
	err = d.Initialize(buf)
	if err != nil {
		return err
	}

	g.dbManager = &d
	return nil

}

//Uninitialize is a
func (g *Server) Uninitialize() error {

	return g.dbManager.Uninitialize()

}

//Close is a
func (g *Server) Close() {
	atomic.StoreInt32(&g.isClosed, 1)
}

//IsClose is a
func (g *Server) IsClose() bool {
	return atomic.LoadInt32(&g.isClosed) != 0
}

func (g *Server) readFile(fileName string) ([]byte, error) {
	return ioutil.ReadFile(fileName)
}

func (g *Server) loadConfig() ([]byte, error) {
	buf, err := ioutil.ReadFile(confPath)
	if err != nil {
		return nil, err
	}

	rowConfig = buf
	err = json.Unmarshal(buf, &g)
	if err != nil {
		return nil, err
	}

	if g.Config.Port == "" {
		g.Config.Port = ":8080"
	}
	if strings.Contains(g.Config.Port, ":") == false {
		g.Config.Port = ":" + g.Config.Port
	}

	return buf, nil
}

// Serve starts listen http requests
func (g *Server) Serve() {

	addr := g.Config.Port
	fmt.Printf("======= Server start to listen (%s) and serve =======\n", addr)

	router := api.Router()
	router.Run(addr)

}
