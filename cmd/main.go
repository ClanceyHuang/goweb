package main

import (
	"fmt"
	"goweb/internal/config"
	"goweb/internal/controller"
	"goweb/internal/log"
	"goweb/internal/middleware"
	"goweb/internal/templates"
	"goweb/internal/database/mysql"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strconv"
	"github.com/joho/godotenv"
)

func main() {

	// 加载 json 配置
	config.LoadConfig("../configs/config.json")
	// 加载日志模块
	log.LoadLog()

	// 加载项目的环境配置，环境变量，数据库配置，密钥等等。
	err := godotenv.Load("../.env")
	if err != nil {
		log.Error.Println("Error loading .env file")
	}

	cur_env := os.Getenv("CUR_ENV")
	db_database := os.Getenv("DB_DATABASE")
	db_username := os.Getenv("DB_USERNAME")
	// 打印调试显示
	log.Info.Println("CUR_ENV: ", cur_env)
	log.Info.Println("DATABASE: ", db_database)
	log.Info.Println("USERNAME: ", db_username)


	// 载入模板
	templates.LoadTemplate(config.Config.Template)

	// 设置服务
	server := http.Server{
		Addr: config.Config.Address + ":" + strconv.Itoa(config.Config.Port),
		Handler: &middleware.TimeoutMiddleWare{
			Next: &middleware.LogMiddleWare{
				Next: &middleware.CrossMiddleWare{},
			},
		},
	}

	// 注册路由
	controller.RegisterRoutes()

	// 显示启动信息
	fmt.Printf("\x1b[1;32m* The Web Server was running on:\x1b[0m\t"+
		"http://%s:%d\n", "localhost", config.Config.Port)
	fmt.Printf("\x1b[1;32m* The profile situation can be view on:\x1b[0m\t"+
		"http://localhost:%d/debug/pprof\n", config.Config.PprofPort)

	// 启动性能监控服务与 web 服务
	go http.ListenAndServe(":"+strconv.Itoa(config.Config.PprofPort), nil)
	server.ListenAndServe()
}
