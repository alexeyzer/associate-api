package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	SessionKey = "Authorization"
)

type config struct {
	GRPC struct {
		ProductAPI string `yaml:"product_api"`
	}
	Auth struct {
		HeaderSessionKey string `yaml:"header_session_key"`
		SessionKey       string `yaml:"session_key"`
		LogoutKey        string `yaml:"logout_key"`
	}
	Database struct {
		Dsn      string `yaml:"dsn"`
		Dbname   string `yaml:"dbname"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Ssl      string `yaml:"ssl"`
	}
	Redis struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Password string `yaml:"password"`
	}
	App struct {
		HttpPort string `yaml:"http_port"`
		GrpcPort string `yaml:"grpc_port"`
		Env      string `yaml:"env"`
	}
}

var Config = config{}

func ReadConf(filename string) error {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	c := &Config
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return fmt.Errorf("in file %q: %v", filename, err)
	}
	c.App.Env = os.Getenv("ENV")
	if c.App.Env == "PROD" {
		c.Database.Host = "database"
		c.Redis.Host = "redis"
	}

	return nil
}
