package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var (
	ServiceName string
	SqliteInfo  Sqlite
	ProxyInfo   Proxy
)

type config struct {
	Base   Base
	Sqlite Sqlite
}

type Base struct {
	ServiceName string
	Proxy       Proxy
}

type Proxy struct {
	HTTP, Socks, Type string
}

type Sqlite struct {
	Path, User, Pass, DBName string
}

func init() {
	configInit()
}

func (conf *config) getConfig(fileName string) *config {
	yamlFile, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(fileName, "yamlFile.Get err:", err)
		return nil
	}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		log.Fatal(fileName, "yamlFile.Get err:", err)
		return nil
	}
	return conf
}

func configInit() {
	var conf config
	conf.getConfig("./config.yml")
	ServiceName = conf.Base.ServiceName
	SqliteInfo = conf.Sqlite
	ProxyInfo = conf.Base.Proxy
}
