package mysql

import (
	"fmt"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"myGo/config"
	"myGo/models"
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	// 引入数据库驱动注册及初始化
	"gorm.io/driver/mysql"
)

func InitEntityDao(d *gorm.DB) {
	models.InitUserDao(d)
	models.InitSchoolDao(d)
}

func makeDsn(user, password, host, db string, port int) string {
	// root:@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", user, password, host, port, db)
}

func InitializeMainDb(o config.ConnectionConfig) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(makeDsn(o.User, o.Password, o.Host, o.Db, o.Port)), &gorm.Config{
		NamingStrategy: &schema.NamingStrategy{
			SingularTable: true,
		},
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
		Logger: logger.Default.LogMode(logMode(o.Debug)),
	})
	if err != nil {
		errStr := fmt.Sprintf("failed to open MySQL master db, error=%v", err)
		return nil, errors.New(errStr)
	}
	mysqlDb, err := db.DB()
	if err != nil {
		return nil, err
	}
	mysqlDb.SetMaxIdleConns(o.MaxIdle)
	mysqlDb.SetMaxOpenConns(o.MaxOpen)
	return db, nil
}

func logMode(debug bool) logger.LogLevel {
	if debug {
		return logger.Info
	}
	return logger.Warn
}
