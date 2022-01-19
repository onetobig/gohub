package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"gohub/bootstrap"
	btsConfig "gohub/config"
	"gohub/pkg/config"
)

func init() {
	// 加载 config 目录下的配置信息
	btsConfig.Initialize()
}

func main() {
	// 配置初始化
	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件，，如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()
	config.InitConfig(env)

	fmt.Println(config.Get("app.port"))

	r := gin.New()

	bootstrap.SetUpDB()
	bootstrap.SetUpRoute(r)
	err := r.Run(":" + config.Get("app.port"))
	if err != nil {
		fmt.Println(err.Error())
	}

}
