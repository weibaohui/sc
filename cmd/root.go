package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"sc/config"
	"sc/counter"
	"sc/file"
)

var ignoreHide = true
var debug = false
var path string

var rootCmd = &cobra.Command{
	Use:   "sc",
	Short: "统计源码行数",
	Long:  "按文件夹统计源码行数",
	Run: func(cmd *cobra.Command, args []string) {
		config.Instant().SetConfig(ignoreHide, debug)
		initFolder := &file.Folder{
			FullPath: path,
			Hidden:   false,
		}
		initFolder.Execute()
		fmt.Println(counter.Instant().Sum())
	},
}

// Execute 执行
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&debug, "debug", "d", false, "调试")
	rootCmd.Flags().StringVarP(&path, "path", "p", ".", "扫描路径")
}
