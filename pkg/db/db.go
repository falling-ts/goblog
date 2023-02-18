package db

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"goblog/pkg/config"
	"goblog/pkg/logger"
	gormMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"time"
)

var (
	gormDB *gorm.DB
	db     *sql.DB
)

func InitDB() *sql.DB {

	var err error

	// 设置数据库连接信息
	config := mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Addr:                 "127.0.0.1:3306",
		Net:                  "tcp",
		DBName:               "goblog",
		AllowNativePasswords: true,
	}

	// 准备数据库连接池
	db, err = sql.Open("mysql", config.FormatDSN())
	logger.LogError(err)

	// 设置最大连接数
	db.SetMaxOpenConns(100)
	// 设置最大空闲连接数
	db.SetMaxIdleConns(25)
	// 设置每个链接的过期时间
	db.SetConnMaxLifetime(5 * time.Minute)

	// 尝试连接，失败会报错
	err = db.Ping()
	logger.LogError(err)

	return db
}

// InitGormDB 初始化模型
func InitGormDB() *gorm.DB {

	var err error

	// 初始化 MySQL 连接信息
	gormConfig := gormMysql.New(gormMysql.Config{
		DSN: fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&loc=Local",
			config.GetString("database.mysql.username"),
			config.GetString("database.mysql.password"),
			config.GetString("database.mysql.host"),
			config.GetString("database.mysql.port"),
			config.GetString("database.mysql.database"),
			config.GetString("database.mysql.charset")),
	})

	var level gormLogger.LogLevel
	if config.GetBool("app.debug") {
		// 读取不到数据也会显示
		level = gormLogger.Warn
	} else {
		// 只有错误才会显示
		level = gormLogger.Error
	}

	// 准备数据库连接池
	gormDB, err = gorm.Open(gormConfig, &gorm.Config{
		Logger: gormLogger.Default.LogMode(level),
	})

	logger.LogError(err)

	db, _ = gormDB.DB()

	// 设置最大连接数
	db.SetMaxOpenConns(config.GetInt("database.mysql.max_open_connections"))
	// 设置最大空闲连接数
	db.SetMaxIdleConns(config.GetInt("database.mysql.max_idle_connections"))
	// 设置每个链接的过期时间
	db.SetConnMaxLifetime(time.Duration(config.GetInt("database.mysql.max_life_seconds")) * time.Second)

	return gormDB
}
