package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Conf 全局变量，用来保存程序的所有配置信息
var Conf = new(AppConfig)

type AppConfig struct {
	Name    string `mapstructure:"name"`
	Mode    string `mapstructure:"mode"`
	Version string `mapstructure:"version"`
	Port    int    `mapstructure:"port"`

	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DbName       string `mapstructure:"db_name"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `matstructure:"port"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

func Init(filePath string) (err error) {
	// 方式1：直接指定配置文件路径（相对路径或者绝对路径）
	// 相对路径：相对执行的可执行文件的相对路径
	// viper.SetConfigFile("./conf/config.yaml")
	// 绝对路径：系统中实际的文件路径
	// viper.SetConfigFile("/Users/qiaopengjun/Desktop/web_app2 /conf/config.yaml")

	// 方式2：指定配置文件名和配置文件的位置，viper 自行查找可用的配置文件
	// 配置文件名不需要带后缀
	// 配置文件位置可配置多个
	// 注意：viper 是根据文件名查找，配置目录里不要有同名的配置文件。
	// 例如：在配置目录 ./conf 中不要同时存在 config.yaml、config.json

	// 读取配置文件
	viper.SetConfigFile(filePath) // 指定配置文件路径
	//viper.SetConfigName("config")        // 配置文件名称(无扩展名)
	//viper.AddConfigPath(".")             // 指定查找配置文件的路径（这里使用相对路径）可以配置多个
	//viper.AddConfigPath("./conf")        // 指定查找配置文件的路径（这里使用相对路径）可以配置多个
	// SetConfigType设置远端源返回的配置类型，例如:“json”。
	// 基本上是配合远程配置中心使用的，告诉viper 当前的数据使用什么格式去解析
	//viper.SetConfigType("yaml")

	err = viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {            // 处理读取配置文件的错误
		fmt.Printf("viper.ReadInConfig failed, error: %v\n", err)
		return
	}

	// 把读取到的配置信息反序列化到 Conf 变量中
	if err = viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper unmarshal failed, error: %v\n", err)
		return
	}

	// 实时监控配置文件的变化 WatchConfig 开始监视配置文件的更改。
	viper.WatchConfig()
	// OnConfigChange设置配置文件更改时调用的事件处理程序。
	// 当配置文件变化之后调用的一个回调函数
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		if err = viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper unmarshal OnConfigChange failed, error: %v\n", err)
		}
	})

	return
}
