package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"sync/atomic"

	api "giftForum/api"
	"giftForum/config"
	"giftForum/db"
	"giftForum/models"

	"github.com/gin-gonic/gin"
)

const confPath = "./config.json"

var rowConfig []byte

type Config struct {
	Port    string
	GinMode string
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

	//Todo: use subSystem
	g.InitializeCredentials()
	config.Initialize(buf)
	models.Initialize(buf)

	g.dbManager = &d
	return nil

}

func (g *Server) InitializeCredentials() error {
	var cred config.Credentials
	file, err := ioutil.ReadFile("./creds.json")
	if err != nil {
		fmt.Printf("InitializeCredentials File error: %v\n", err)
		return err
	}
	err = json.Unmarshal(file, &cred)
	if err != nil {
		fmt.Printf("InitializeCredentials Unmarshal error: %v\n", err)
		return  err
	}
	config.SetCredentials(cred)
	config.SetPort(g.Config.Port)
	return nil
}

//Uninitialize is a
func (g *Server) Uninitialize() error {

	models.Uninitialize()
	config.Uninitialize()
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
	if g.Config.GinMode != "" {
		gin.SetMode(g.Config.GinMode)
	}

	router.Run(addr)
	//router.Run("localhost:8081")
	
	//fmt.Println(router.ListenAndServe("addr", nil))
}

