package mysql

import (
	"flag"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"myGo/config"
	"myGo/models"
	"time"
	// 引入数据库驱动注册及初始化
	_ "github.com/go-sql-driver/mysql"
)

var (
	ormLog = flag.Bool("ormlog", false, "--ormlog")
	master *gorm.DB
)

func InitEntityDao(d *gorm.DB) {
	models.InitUserDao(d)
}

func makeDsn(user, password, host, db string, port int) string {
	// root:@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", user, password, host, port, db)
}

func InitializeMainDb(o config.ConnectionConfig) error {
	masterDB, err := gorm.Open("mysql", makeDsn(o.User, o.Password, o.Host, o.Db, o.Port))
	if err != nil {
		errStr := fmt.Sprintf("failed to open MySQL master db, error=%v", err)
		return errors.New(errStr)
	}
	master = masterDB
	if *ormLog {
		master = masterDB.Debug()
	}
	master.DB().SetMaxIdleConns(o.MaxIdle)
	master.DB().SetMaxOpenConns(o.MaxOpen)
	master.SetNowFuncOverride(func() time.Time {
		return time.Now().UTC()
	})
	master.SingularTable(true)
	// master.SetLogger(gorm.Logger{})
	return nil
}

func GetClient() *gorm.DB {
	if master == nil {
		// return InitializeDao() get global config
		return nil
	}
	return master
}
