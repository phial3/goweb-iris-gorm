package conf

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type sysConfig struct {
	AppName    string `json:"AppName"`
	Port       string `json:"Port"`
	DBUserName string `json:"DBUserName"`
	DBPassword string `json:"DBPassword"`
	DBIp       string `json:"DBIp"`
	DBPort     string `json:"DBPort"`
	DBName     string `json:"DBName"`
}

var SysConfig = &sysConfig{}

var cmd = &cobra.Command{}

// 从JSON文件读取配置
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

func Init(appName string) error {
	c := sysConfig{
		AppName: appName,
	}

	if err := c.initConfig(); err != nil {
		return err
	}

	initLog()
	return nil

}

func (c *sysConfig) initConfig() error {
	if c.AppName != "" {
		viper.SetConfigFile(c.AppName)
	} else {
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}

	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("irisgo")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func initLog() {
	//WithMaxAge := viper.GetDuration("log.with_max_age")
	//WithRotationTime := viper.GetDuration("log.with_rotation_time")
	//WithRotationCount := viper.GetInt("log.with_rotation_count")
	LogName := viper.GetString("log.logger_file")

	writer1 := os.Stdout

	// 设置日志输入在文件
	writer2, err := os.OpenFile(LogName, os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_SYNC, 0755)
	if err != nil {
		log.Fatalf("create file log.txt failed:%v", err)
	}

	//设置在输出日志中添加文件名和方法信息
	//logrus.SetReportCaller(true)

	//设置输出写入
	logrus.SetOutput(io.MultiWriter(writer1, writer2))
	logrus.SetLevel(logrus.DebugLevel)

	//可设置json，txt，nested（需要引入github.com/antonfisher/nested-logrus-formatter）等格式
	//logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.SetFormatter(&nested.Formatter{
		FieldsOrder:           []string{"component", "category"},
		HideKeys:              false, //是否隐藏键值
		NoColors:              false, //是否显示颜色
		NoFieldsColors:        false,
		ShowFullLevel:         false,
		TrimMessages:          false,
		CallerFirst:           false,
		CustomCallerFormatter: nil,
		TimestampFormat:       time.RFC3339, //格式化时间
	})
}
