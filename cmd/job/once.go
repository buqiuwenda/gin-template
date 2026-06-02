package job

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "job",
		Short: "运行后台/一次性任务",
	}
	cmd.AddCommand(&cobra.Command{
		Use:   "once",
		Short: "执行一次性任务示例",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("job once: TODO implement")
			return nil
		},
	})
	return cmd
}
