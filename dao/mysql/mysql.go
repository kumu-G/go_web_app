package mysql

import (
	"fmt"
	"web_app2/settings"

	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql" // 匿名导入 自动执行 init()
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func Init(cfg *settings.MySQLConfig) (err error) {
	//DSN (Data Source Name) Sprintf根据格式说明符进行格式化，并返回结果字符串。
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
	)
	// 连接到数据库并使用ping进行验证。
	// 也可以使用 MustConnect MustConnect连接到数据库，并在出现错误时恐慌 panic。
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect DB failed", zap.Error(err))
		return
	}
	db.SetMaxOpenConns(cfg.MaxOpenConns) // 设置数据库的最大打开连接数。
	db.SetMaxIdleConns(cfg.MaxIdleConns) // 设置空闲连接池中的最大连接数。
	return
}

func Close() {
	_ = db.Close()
}
