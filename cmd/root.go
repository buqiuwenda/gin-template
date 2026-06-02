package main

import (
	"os"

	"github.com/buqiuwenda/gin-template/cmd/job"
	"github.com/buqiuwenda/gin-template/cmd/web"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gin-template",
	Short: "Gin Web API 服务脚手架",
}

func init() {
	rootCmd.AddCommand(web.NewCommand())
	rootCmd.AddCommand(job.NewCommand())
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
