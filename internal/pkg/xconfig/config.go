package xconfig

import (
	"github.com/buqiuwenda/gin-template/internal/conf"
	"github.com/buqiuwenda/seal/config"
	"github.com/buqiuwenda/seal/config/file"
	"github.com/buqiuwenda/seal/log"
	"github.com/buqiuwenda/seal/pkg/env"
)

type BaseSource struct {
	FileName string // 本地文件名(主要用于本地加载)
	DataId   string // nacos dataId(用于远程加载)
}

type ConfigSource struct {
	ServerName string
	BaseSource []BaseSource
}

func getConfigFile() []ConfigSource {
	configMap := []ConfigSource{
		{
			ServerName: "gin-template",
			BaseSource: []BaseSource{
				{
					FileName: "config.yaml",
					DataId:   "config.yaml",
				},
			},
		},
	}
	return configMap
}

func getClientSources(configSource []ConfigSource) []config.Source {
	filePrefix := "configs/" + env.Mode()

	var sources []config.Source

	for _, s := range configSource {
		for _, b := range s.BaseSource {
			log.Infof("读取文件：%s", filePrefix+"/"+b.FileName)
			sources = append(sources, file.NewSource(filePrefix+"/"+b.FileName))
		}
	}
	return sources
}

func getConfigSource() config.Config {
	sources := getClientSources(getConfigFile())
	return config.New(
		config.WithSource(
			sources...,
		),
	)
}

func InitConfig() (bc conf.Bootstrap, err error) {
	c := getConfigSource()

	defer c.Close()

	if err = c.Load(); err != nil {
		log.Fatal("加载配置失败", err)
	}

	if err = c.Scan(&bc); err != nil {
		log.Fatal("获取配置失败", err)
	}

	return
}
