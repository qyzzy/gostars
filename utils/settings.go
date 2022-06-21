package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
	"sync"
)

var (
	AppMode       string
	HttpPort      string
	JwtKey        string
	JwtExpireTime int64

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	Cache     string
	CacheHost string
	CachePort string
)

var once sync.Once

func init() {
	once.Do(Load)
}

func Load() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("Load failed")
		panic(err)
	}
	LoadServer(file)
	LoadData(file)
	LoadCache(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("89js82js72")
	JwtExpireTime = file.Section("server").Key("JwtExpireTime").MustInt64(10000)
}

func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("debug")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("gostars")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("admin123")
	DbName = file.Section("database").Key("DbName").MustString("gostars")
}

func LoadCache(file *ini.File) {
	Cache = file.Section("cache").Key("Cache").MustString("debug")
	CacheHost = file.Section("cache").Key("CacheHost").MustString("localhost")
	CachePort = file.Section("cache").Key("CachePort").MustString(":6379")
}
