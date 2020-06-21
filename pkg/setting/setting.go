package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type App struct {
	PageSize        int
	JwtSecret       string
	JwtExpireTime   time.Duration
	RuntimeRootPath string

	ImagePrefixUrl string
	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExt  []string

	ApkSavePath string
	ApkAllowExt string
	AppStoreUrl string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string

	WechatAppID  string
	WechatSecret string
	QQAppID      string
	QQAppKey     string
}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var AppSetting = &App{}

var ServerSetting = &Server{}

var DatabaseSetting = &Database{}

/**
初始化配置信息
*/
func SetUp() {
	Cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatal("获取ini配置失败")
	}

	err = Cfg.Section("app").MapTo(AppSetting)
	if err != nil {
		log.Fatalf("配置文件映射app Section失败: %v", err)
	}
	// 设置允许上传的最大图片
	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024

	err = Cfg.Section("server").MapTo(ServerSetting)
	if err != nil {
		log.Fatalf("配置文件映射server Section失败: %v", err)
	}
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second

	err = Cfg.Section("database").MapTo(DatabaseSetting)
	if err != nil {
		log.Fatalf("配置文件映射database Section失败: %v", err)
	}
}
