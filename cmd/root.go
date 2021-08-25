// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"hosts-setter/service"
	"log"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hosts-setter",
	Short: "定时更新本地的 hosts 文件，实现自定义域名的访问",
	Long: `定时更新本地的 hosts 文件，实现自定义域名的访问

用法：
注册到 windows 服务: sc.exe create HostsSetter binPath={path-to-this-project}/hosts-setter.exe -a={http://curl-router-host:port} -n={hostname}

-a 指定curl-router-host的地址
-n 指定要更新的域名/主机
-i 指定更新间隔，单位：秒
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		interval, _ := cmd.Flags().GetInt("interval")
		address, _ := cmd.Flags().GetString("address")
		if address == "" {
			log.Println("缺少 -a 参数，使用 -h 查看使用说明")
			return
		}
		hostname, _ := cmd.Flags().GetString("hostname")
		if hostname == "" {
			log.Println("缺少 -n 参数，使用 -h 查看使用说明")
			return
		}

		service.WindowsRun(address, hostname, interval)
	},
}

func init() {
	rootCmd.Flags().StringP("address", "a", "", "curl-router 的地址")
	rootCmd.Flags().StringP("hostname", "n", "", "hostname")
	rootCmd.Flags().IntP("interval", "i", 60*60, "更新间隔，单位：秒")

	// rootCmd.AddCommand(listCmd)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
