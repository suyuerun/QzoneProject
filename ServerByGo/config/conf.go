package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	ServerHost  string
	ServerPort  int
	MysqlHost   string
	MysqlPort   int
	MysqlUser   string
	MysqlPwd    string
	MysqlDBName string
}

var GlobalConfig *Config

func init() {
	Config := new(Config)
	data, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		fmt.Printf("%v", err)
	}
	err = json.Unmarshal(data, Config)
	if err != nil {
		fmt.Printf("%v \t", err)
	} else {
		fmt.Printf("Config: %v \n", Config)
		GlobalConfig = Config
	}
}
