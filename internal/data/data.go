package data

import (
	"fmt"

	"github.com/buqiuwenda/gin-template/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Data 基础设施聚合（GORM 等）
type Data struct {
	db *gorm.DB
}

func NewData(cfg *config.Config) (*Data, func(), error) {
	if cfg.Data.Source == "" {
		return nil, nil, fmt.Errorf("data.source is required")
	}

	db, err := openDB(cfg)
	if err != nil {
		return nil, nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, err
	}
	if err := sqlDB.Ping(); err != nil {
		return nil, nil, fmt.Errorf("ping database: %w", err)
	}

	cleanup := func() {
		_ = sqlDB.Close()
	}
	return &Data{db: db}, cleanup, nil
}

func openDB(cfg *config.Config) (*gorm.DB, error) {
	gormCfg := &gorm.Config{}
	if cfg.Server.Mode == "debug" {
		gormCfg.Logger = logger.Default.LogMode(logger.Info)
	}

	switch cfg.Data.Driver {
	case "", "mysql":
		return gorm.Open(mysql.Open(cfg.Data.Source), gormCfg)
	default:
		return nil, fmt.Errorf("unsupported data driver: %s", cfg.Data.Driver)
	}
}

func (d *Data) DB() *gorm.DB {
	return d.db
}
