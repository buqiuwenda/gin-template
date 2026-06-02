package data

import (
	"database/sql"

	"github.com/buqiuwenda/gin-template/internal/config"
)

// Data 基础设施聚合（DB、Cache 等）
type Data struct {
	db *sql.DB
}

func NewData(cfg *config.Config) (*Data, func(), error) {
	// TODO: 按 cfg.Data.Driver 初始化真实连接；模板阶段允许无 DB 启动
	_ = cfg
	return &Data{}, func() {}, nil
}

func (d *Data) DB() *sql.DB {
	return d.db
}
