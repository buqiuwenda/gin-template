package web

import (
	"log"

	"github.com/spf13/cobra"
)

var configPath string

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "web",
		Short: "启动 HTTP API 服务（Gin）",
		RunE: func(cmd *cobra.Command, args []string) error {
			application, cleanup, err := InitializeApp(configPath)
			if err != nil {
				return err
			}
			defer cleanup()
			return application.Run()
		},
	}
	cmd.Flags().StringVarP(&configPath, "config", "c", "configs/config.example.yaml", "配置文件路径")
	return cmd
}

func init() {
	cobra.OnInitialize(func() {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	})
}
