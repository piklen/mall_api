package main

import (
	"log"
	"mall_api/grpcclient"
	"mall_api/routes"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	r := routes.NewRouter()

	// 启动HTTP服务器
	go func() {
		if err := r.Run(":7999"); err != nil {
			grpcclient.CloseGrpcConn() // 关闭gRPC连接
			log.Fatalf("Failed to run server: %v", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，并关闭gRPC连接
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	grpcclient.CloseGrpcConn() // 关闭gRPC连接
}
