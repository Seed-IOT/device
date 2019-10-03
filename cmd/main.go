package main

import (
	"device/config"
	"device/log"
	"device/model"
	"device/web"

	"os"
	"os/signal"

	_ "device/docs"
)

func main() {

	log.Init()

	cfg, err := config.New()
	if err != nil {
		log.GlobalLog.Fatalf("Failed to reading config file, %s\n", err)
	}

	service, err := model.New(cfg.Database)
	if err != nil {
		log.GlobalLog.Fatalf("Failed to initialize model for operating all service, %s\n", err)
	}

	server := web.NewServer(cfg, service)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	// 日志中间件
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.GlobalLog.Fatalf("Failed to listen for http server, %s\n", err)
		}
	}()

	defer service.DB.Close()

	log.GlobalLog.Println("device is running")
	<-quit
	log.GlobalLog.Println("device is stopped")
}
