package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

//app配置
type App struct {
	JwtSecret       string        //jwt密钥
	JwtExpireTime   time.Duration //jwt失效时间，单位是小时
	PageSize        int           //分页返回的数据个数
	RuntimeRootPath string        //保存文件的跟路径

	ImagePrefixUrl string   //图片url
	ImageSavePath  string   //图片要保存的路径
	ImageMaxSize   int      //图片最大尺寸
	ImageAllowExts []string //图片允许的格式   jpg jpeg png

	ApkSavePath string //apk文件路径
	ApkAllowExt string //apk文件格式
	AppStoreUrl string //iOS应用在App Store中的地址，用于版本更新

	LogSavePath string //日志文件保存的路径
	LogSaveName string //日志文件名称
	LogFileExt  string //日志文件后缀
	TimeFormat  string //文件的日期名称

	WechatAppID  string //微信的appID
	WechatSecret string //微信的secret
	QQAppID      string //QQ的appID
	QQAppKey     string //QQ的appkey
}

var AppSetting = &App{}

//服务配置
type Server struct {
	RunMode      string        //运行模式
	HttpPort     int           //端口号
	ReadTimeout  time.Duration //读取超时时间
	WriteTimeout time.Duration //写入超时时间
}

var ServerSetting = &Server{}

//数据库配置
type Database struct {
	Type        string //数据库类型
	User        string //数据库用户
	Password    string //数据库密码
	Host        string //数据库地址+端口号
	Name        string //数据库名称
	TablePrefix string //数据库数据表前缀
}

var DatabaseSetting = &Database{}

//将配置选项映射到结构体上
func SetUp() {
	Cfg, err := ini.Load("conf/app.ini") //加载配置文件ini
	if err != nil {
		log.Fatal("获取.ini配置失败")
	}

	//映射配置
	err = Cfg.Section("app").MapTo(AppSetting)
	if err != nil {
		log.Fatalf("Cfg配置文件映射 AppSetting 错误: %v", err)
	}

	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024 //设置允许上传图片的最大尺寸

	err = Cfg.Section("server").MapTo(ServerSetting)
	if err != nil {
		log.Fatalf("Cfg配置文件映射 ServerSetting 错误: %v", err)
	}

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.ReadTimeout * time.Second

	err = Cfg.Section("database").MapTo(DatabaseSetting)
	if err != nil {
		log.Fatalf("Cfg配置文件映射 DatabaseSetting 错误: %v", err)
	}
}
