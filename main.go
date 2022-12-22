package main

import (
	"context"
	"go_code/gintest/bootstrap"
	"go_code/gintest/bootstrap/gdb"
	"go_code/gintest/bootstrap/glog"
	"go_code/gintest/bootstrap/gredis"
	"go_code/gintest/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	// 1. 加载配置
	// 2. 初始化日志 zap
	// 3. 初始化 MySQL sqlx
	// 4. 初始化 Redis
	bootstrap.Initialize("./config/dev_conf.yaml")
	defer glog.Sync()

	// 关闭 MySQL
	defer gdb.Close()
	// 关闭 redis
	defer gredis.Close()

	// 5. 注册路由
	r := routes.RegisterRouters()

	// 6. 启动服务，优雅关机
	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(bootstrap.Config.Port),
		Handler: r,
	}

	go func() {
		// 开启一个 goroutine 启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			glog.SL.Error("Listen failed, ", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	log.Println("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

	log.Println("Server exiting")
}
