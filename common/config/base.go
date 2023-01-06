package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// 初始化参数
var (
	ServiceName string
	SqliteInfo  Sqlite
)

// Config 基础配置
type config struct {
	Base   Base
	Sqlite Sqlite
}

// Base 搜索引擎
type Base struct {
	ServiceName string
}

// TiDB 数据库配置
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

// configInit 获取配置文件
func configInit() {
	var conf config
	conf.getConfig("./config.yml")
	ServiceName = conf.Base.ServiceName
	SqliteInfo = conf.Sqlite
}
