package setting

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

var (
	RunMode      string
	config       Config
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Type         string
	User         string
	Password     string
	Host         string
	Name         string
)

type Config struct {
	RunMode string `yaml:"runMode"`

	App struct {
		PageSize  int    `yaml:"pageSize"`
		JwtSecret string `yaml:"jwtSecret"`
	}

	Server struct {
		Port         int           `yaml:"port"`
		ReadTimeout  time.Duration `yaml:"readTimeout"`
		WriteTimeout time.Duration `yaml:"writeTimeout"`
	}

	Database struct {
		Type     string `yaml:"type"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Name     string `yaml:"name"`
	}
}

// 初始加载配置函数
func init() {

	yamlConfig, _ := ioutil.ReadFile("conf.yml")

	_ = yaml.Unmarshal(yamlConfig, &config)

	Port = config.Server.Port
	ReadTimeout = config.Server.ReadTimeout * time.Second
	WriteTimeout = config.Server.WriteTimeout * time.Second
	dbConfig := config.Database
	Type = dbConfig.Type
	User = dbConfig.User
	Password = dbConfig.Password
	Host = dbConfig.Host
	Name = dbConfig.Name

	RunMode = config.RunMode

}
