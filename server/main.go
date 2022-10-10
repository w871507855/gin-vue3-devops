package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"server/common"
	"server/helper"
	"server/logger"
	"server/router"
	"syscall"
	"time"
)

func main() {
	r := router.Router()
	common.InitConfig("conf/config.yaml")
	common.InitDB()
	// init logger
	if err := logger.InitLogger(common.Conf.Log); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	// 初始化gin框架内的校验器使用的翻译器
	helper.InitTrans("zh")
	if err := helper.InitTrans("zh"); err != nil {
		fmt.Printf("init validator trans failed, err:%v\n", err)
		return
	}
	// 初始化雪花算法
	if err := helper.InitSnowFlake(common.Conf.SnowFlake.Startime, common.Conf.SnowFlake.MachineID); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}

	// 监听并在0.0.0.0:8088上启动服务
	// r.Run("0.0.0.0:8088")
	srv := &http.Server{
		Addr:    ":8088",
		Handler: r,
	}
	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	// 创建一个接受信号的通道
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务，超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

	log.Println("Server exiting")
}
