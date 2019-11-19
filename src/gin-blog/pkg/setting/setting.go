package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"

)

type App struct {
	JwtSecret string
	PageSize int
	RuntimeRootPath string

	ImagePrefixUrl string
	ImageSavePath string
	ImageMaxSize int
	ImageAllowExts []string

	LogSavePath string
	LogSaveName string
	LogFileExt string
	TimeFormat string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}
var ServerSetting = &Server{}

type Database struct {
	Type string
	User string
	Password string
	Host string
	Name string
	TablePrefix string
}

var DatabaseSetting = &Database{}

type Redisbase struct{
	MaxIdle  int
	MaxActive int
	IdleTimeout time.Duration
	Host string
	Password string

}
var RdisSetting = &Redisbase{}

func Setup() {
	Cfg, err := ini.Load("gin-blog/conf/conf.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'gin-blog/conf/conf.ini': %v", err)
	}

	err = Cfg.Section("app").MapTo(AppSetting) //映射
	if err != nil {
		log.Fatalf("Cfg.MapTo AppSetting err: %v", err)
	}

	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024

	err = Cfg.Section("server").MapTo(ServerSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo ServerSetting err: %v", err)
	}

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.ReadTimeout * time.Second

	err = Cfg.Section("database").MapTo(DatabaseSetting) //使用 MapTo 将配置项映射到结构体上
	if err != nil {
		log.Fatalf("Cfg.MapTo DatabaseSetting err: %v", err)
	}
	err = Cfg.Section("redis").MapTo(RdisSetting)
	if err !=nil{
		log.Fatalf("cfg.mapto redissetting err:%v",err)
	}

}