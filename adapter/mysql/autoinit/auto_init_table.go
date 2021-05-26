package autoinit

import (
	"context"
	"myGo/adapter/log"

	"github.com/jinzhu/gorm"
)

var injectors []func(db *gorm.DB)

// RegisterInjector 注册回调
func RegisterInjector(f func(*gorm.DB)) {
	injectors = append(injectors, f)
}

// 执行回调
func callInjector(db *gorm.DB) {
	for _, v := range injectors {
		v(db)
	}
}

// SetupTableModel 自动初始化表结构
func SetupTableModel(ctx context.Context, db *gorm.DB, model interface{}) {
	err := db.AutoMigrate(model)
	if err != nil {
		log.Fatal(ctx, err)
	}
}
