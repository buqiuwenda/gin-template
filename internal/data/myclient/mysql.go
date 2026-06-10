package myclient

import (
	"errors"

	"github.com/buqiuwenda/gin-template/internal/conf"
	"github.com/buqiuwenda/gin-template/internal/meta"
	"github.com/buqiuwenda/seal/pkg/client/mysql"
	"xorm.io/xorm"
)

func newMysqlClient(cliConf *conf.MysqlDBConfig) (*xorm.EngineGroup, func(), error) {
	cfgSlave := make([]struct {
		Dsn string `json:"dsn" mapstructure:"dsn"`
	}, 0)
	if len(cliConf.GetSlaves()) > 0 {
		for _, v := range cliConf.Slaves {
			cfgSlave = append(cfgSlave, struct {
				Dsn string `json:"dsn" mapstructure:"dsn"`
			}{Dsn: v.Dsn})
		}
	}

	cfg := mysql.GroupConfig{
		MaxIdle:     int(cliConf.MaxIdle),
		MaxOpen:     int(cliConf.MaxOpen),
		MaxLifetime: int(cliConf.MaxLifetime),
		Master: struct {
			Dsn string `json:"dsn" mapstructure:"dsn"`
		}{Dsn: cliConf.Master.Dsn},
		Slaves:  cfgSlave,
		IsDebug: cliConf.IsDebug,
		SSL: mysql.SSLConfig{
			Enable:  cliConf.SslEnable,
			CAFile:  cliConf.CaFile,
			TlsName: cliConf.TlsName,
		},
	}

	engineGroup, err := mysql.NewGroupClient(cfg)
	if err != nil {
		return nil, nil, err
	}

	return engineGroup, func() {
		_ = engineGroup.Close()
	}, nil
}

func NewAigcMysqlClient(c *conf.Data) (meta.MysqlWendaAigcClient, func(), error) {

	cfg := c.Mysql["wenda_aigc"]
	if cfg == nil {
		return nil, nil, errors.New("mysql config not found")
	}

	//return meta.MysqlWendaAigcClient, func(), nil
	return newMysqlClient(cfg)
}
