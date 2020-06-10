package main

import (
	"api/models"
	"api/pkg/logging"
	"api/pkg/setting"
	"api/routers"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func init() {
	setting.SetUp() //初始化配置文件
	logging.SetUp()
	models.SetUp()
}

func main() {
	log.Println("Hello, api 正在启动中...")
	// log.Println("服务监听端口为:", setting.ServerSetting.HttpPort) //测试能否打印出ini配置文件设置的信息
	router := routers.InitRouter() //初始化路由

	// router.GET("/test", func(context *gin.Context) {
	// 	context.JSON(e.SUCCESS, gin.H{
	// 		"Code": e.SUCCESS,
	// 		"Msg":  e.GetMsg(e.SUCCESS),
	// 		"Data": "返回数据成功",
	// 	})
	// })

	s := &http.Server{
		Addr:              fmt.Sprintf(":%d", setting.ServerSetting.HttpPort), //设置服务端口号
		Handler:           router,
		ReadHeaderTimeout: setting.ServerSetting.ReadTimeout,
		WriteTimeout:      setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes:    1 << 20,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("程序服务关闭退出")
}
