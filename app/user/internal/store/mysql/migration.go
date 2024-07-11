package mysql

import (
	"go-cloud-storage/app/user/internal/model"
	logger2 "go-cloud-storage/pkg/logger"
	"os"
)

func migration() {
	// 自动迁移
	err := _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&model.User{})
	if err != nil {
		logger2.Errorw("mysql migration error", "err", err)
		os.Exit(0)
	}
}
