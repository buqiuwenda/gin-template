package web

import (
	"github.com/buqiuwenda/seal/log"

	"github.com/buqiuwenda/gin-template/internal/conf"
	"github.com/buqiuwenda/gin-template/internal/pkg/xconfig"
	"github.com/spf13/cobra"
)

var bc *conf.Bootstrap

func init() {
	cfg, err := xconfig.InitConfig()
	if err != nil {
		log.Fatalf("Failed to init config: %v", err)
	}
	bc = &cfg
}

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "http",
		Short: "启动 HTTP API 服务（Gin）",
		RunE: func(cmd *cobra.Command, args []string) error {
			httpSrv, cleanup, err := InitializeHTTPServer(bc)
			if err != nil {
				return err
			}
			defer cleanup()
			return httpSrv.Run()
		},
	}
	return cmd
}
