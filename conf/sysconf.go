package conf

import (
	"github.com/json-iterator/go"
	"io/ioutil"
)

var SysConfig = &sysConfig{}

func init() {
	//指定对应的json配置文件
	b, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic("Sys config read err")
	}
	err = jsoniter.Unmarshal(b, SysConfig)
	if err != nil {
		panic(err)
	}

}

type sysConfig struct {
	Port       string `json:"Port"`
	DBUserName string `json:"DBUserName"`
	DBPassword string `json:"DBPassword"`
	DBIp       string `json:"DBIp"`
	DBPort     string `json:"DBPort"`
	DBName     string `json:"DBName"`
}
