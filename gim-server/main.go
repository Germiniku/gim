/*
 * @Descripttion:
 * @Author: Sun
 * @Date: 2020-10-14 14:17:55
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2020-10-15 11:07:35
 */
package main

import (
	"flag"
	"fmt"
	"gim-server/cache"
	"gim-server/conf"
	"gim-server/model"
	"gim-server/router"
	"os"
)

func main() {
	var (
		path string
		err  error
	)
	flag.StringVar(&path, "f", "./conf.toml", "配置文件")
	// 初始化配置文件
	fmt.Println(path)
	if err = conf.Init(path); err != nil {
		os.Exit(1)
	}
	// 初始化服务
	if err = cache.Init(); err != nil {
		os.Exit(1)
	}
	if err = model.Init(); err != nil {
		os.Exit(1)
	}
	server := router.Init()

	server.Run()
}
