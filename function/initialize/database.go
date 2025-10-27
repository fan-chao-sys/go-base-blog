package initialize

import (
	"os"
	"path/filepath"

	"go-base-blog/function/model"
	utilLog "go-base-blog/function/utils"

	toml "github.com/pelletier/go-toml/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var config Config
var db *gorm.DB

type Config struct {
	MysqlConfig  `toml:"mysql"`
	ServerConfig `toml:"server"`
}

type ServerConfig struct {
	Port    int `toml:"port"`
	Timeout int `toml:"timeout"` // 表示秒数，使用时转换为 time.Duration(timeout) * time.Second
}

type MysqlConfig struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	Name     string `toml:"name"`
	Password string `toml:"password"`
	Charset  string `toml:"charset"`
	Database string `toml:"database"`
}

func ReadJoinMysqlAddress() string {
	// 尝试从常见位置加载配置：优先使用 project 根的 config/db.toml，兼容 function 子目录下的 config
	candidates := []string{"config/db.toml", "function/config/db.toml", "db.toml"}
	var cfgPath string
	for _, c := range candidates {
		if _, err := os.Stat(c); err == nil {
			cfgPath = c
			break
		}
	}
	if cfgPath == "" {
		utilLog.LogInfo(utilLog.FormatMessage("未找到配置文件，候选位置: %v", candidates))
		os.Exit(1)
	}

	// 读取文件内容并使用 pelletier/go-toml/v2 解析
	data, err := os.ReadFile(filepath.Clean(cfgPath))
	if err != nil {
		utilLog.LogInfo(utilLog.FormatMessage("读取配置文件失败: %v", err))
		os.Exit(1)
	}
	if err := toml.Unmarshal(data, &config); err != nil {
		utilLog.LogInfo(utilLog.FormatMessage("解析配置文件失败: %v", err))
		os.Exit(1)
	}
	mysql := &config.MysqlConfig
	server := &config.ServerConfig
	// 使用配置信息
	utilLog.LogInfo(utilLog.FormatMessage("数据库地址-ip:%s, 端口=%s\n", mysql.Host, mysql.Port))
	utilLog.LogInfo(utilLog.FormatMessage("服务器端口: %d", server.Port))

	// 	dbAddress := "root:root@tcp(127.0.0.1:3306)/go_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	// 组装 MySQL DSN: user:password@tcp(host:port)/dbname?charset=...&parseTime=True&loc=Local
	str := mysql.Name + ":" + mysql.Password + "@tcp(" + mysql.Host + ":" + mysql.Port + ")/" + mysql.Database + "?charset=" + mysql.Charset + "&parseTime=True&loc=Local"
	utilLog.LogInfo(utilLog.FormatMessage("数据库mysql连接组装地址: %s\n", str))
	return str
}

func DBInit() error {
	var err error
	dataBase, err := gorm.Open(mysql.Open(ReadJoinMysqlAddress()), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败:" + err.Error())
	}

	db = dataBase
	err = dataBase.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{}, &model.Log{})
	if err != nil {

		return err
	}
	return err
}

func GetDB() *gorm.DB {
	return db
}
