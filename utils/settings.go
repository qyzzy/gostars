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
	DbPassword string
	DbName     string

	Cache     string
	CacheHost string
	CachePort string

	FtpAddress    string
	FtpUser       string
	FtpPassword   string
	HeartbeatTime int

	AccessKey  string
	SecretKey  string
	Bucket     string
	QiniuSever string

	RabbitMQUser     string
	RabbitMQPassword string
	RabbitMQAddress  string
	RabbitMQPort     string

	ModelPath  string
	PolicyPath string
)

const (
	IsLike     = 0
	UnLike     = 1
	LikeAction = 1

	// maximum number of database operations
	Attempts = 3
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
	LoadFtp(file)
	LoadQiNiu(file)
	LoadRabbitMQ(file)
	LoadCasbin(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("")
	JwtExpireTime = file.Section("server").Key("JwtExpireTime").MustInt64(10000)
}

func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("debug")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("")
	DbPassword = file.Section("database").Key("DbPassword").MustString("")
	DbName = file.Section("database").Key("DbName").MustString("gostars")
}

func LoadCache(file *ini.File) {
	Cache = file.Section("cache").Key("Cache").MustString("debug")
	CacheHost = file.Section("cache").Key("CacheHost").MustString("localhost")
	CachePort = file.Section("cache").Key("CachePort").MustString(":6379")
}

func LoadFtp(file *ini.File) {
	FtpAddress = file.Section("ftp").Key("FtpAddress").MustString("")
	FtpUser = file.Section("ftp").Key("FtpUser").MustString("")
	FtpPassword = file.Section("ftp").Key("FtpPassword").MustString("")
	HeartbeatTime = file.Section("ftp").Key("HeartbeatTime").MustInt(120)
}

func LoadQiNiu(file *ini.File) {
	AccessKey = file.Section("qiniu").Key("AccessKey").MustString("")
	SecretKey = file.Section("qiniu").Key("SecretKey").MustString("")
	Bucket = file.Section("qiniu").Key("Bucket").MustString("")
	QiniuSever = file.Section("qiniu").Key("QiniuSever").MustString("")
}

func LoadRabbitMQ(file *ini.File) {
	RabbitMQUser = file.Section("rabbitmq").Key("RabbitMQUser").MustString("")
	RabbitMQPassword = file.Section("rabbitmq").Key("RabbitMQPassword").MustString("")
	RabbitMQAddress = file.Section("rabbitmq").Key("RabbitMQAddress").MustString("")
	RabbitMQPort = file.Section("rabbitmq").Key("RabbitMQPort").MustString("")
}

func LoadCasbin(file *ini.File) {
	ModelPath = file.Section("casbin").Key("ModelPath").MustString("")
	PolicyPath = file.Section("casbin").Key("PolicyPath").MustString("")
}
