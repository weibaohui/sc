package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/weibaohui/sc/config"
	"github.com/weibaohui/sc/counter"
	"github.com/weibaohui/sc/file"
	"github.com/weibaohui/sc/git"
	"github.com/weibaohui/sc/utils"
)

var (
	ignoreHide = true
	debug      = false
	force      = false
	path       string
	skipSuffix string
	exclude    string
)

var rootCmd = &cobra.Command{
	Use:   "sc",
	Short: "统计源码行数",
	Long:  "按文件夹统计源码行数",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.GetInstance()
		cfg.IgnoreHide = ignoreHide
		cfg.Debug = debug
		cfg.InitPath = path
		cfg.SetSkipSuffix(skipSuffix)
		cfg.SetExclude(exclude)

		if cfg.Debug {
			if bytes, err := json.Marshal(cfg); err == nil {
				fmt.Println(string(bytes))
			}

		}

		// 检查git 是否已经安装
		if _, err := git.BinVersion(); err == nil {
			result["git"] = git.GetInstance().GoExecute().Result()
		} else {
			if cfg.Debug {
				fmt.Println("当前系统未安装git，暂不统计git信息")
			}
		}

		initFolder := &file.Folder{
			FullPath: cfg.InitPath,
			Hidden:   false,
		}
		initFolder.Execute()
		result["source"] = counter.GetInstance().Sum()

		// 输出json
		bytes, err := json.Marshal(result)
		utils.CheckIfError(err)
		fmt.Println(string(bytes))

	},
}
var result = map[string]interface{}{}

// Execute 执行
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&debug, "debug", "d", false, "调试")
	rootCmd.Flags().BoolVarP(&force, "force", "", false, "使用自定义配置覆盖默认初始配置，否则合并")
	rootCmd.Flags().StringVarP(&path, "path", "p", ".", "扫描路径")
	rootCmd.Flags().StringVarP(&skipSuffix, "skipSuffix", "", "", "跳过文件后缀列表,使用逗号分割")
	rootCmd.Flags().StringVarP(&exclude, "exclude", "", "", "跳过文件夹列表,使用逗号分割")
}
