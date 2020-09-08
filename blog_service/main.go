/**
 */
package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/noChaos1012/tour/blog_service/global"
	"github.com/noChaos1012/tour/blog_service/internal/model"
	"github.com/noChaos1012/tour/blog_service/internal/routers"
	"github.com/noChaos1012/tour/blog_service/pkg/logger"
	st "github.com/noChaos1012/tour/blog_service/pkg/setting"
	"github.com/noChaos1012/tour/blog_service/pkg/tracer"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	isVersion    bool
	buildTime    string
	buildVersion string
	gitCommitID  string
)

func init() {
	setupFlag()
	//初始化配置信息
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err:%v", err)
	}
	//初始化日志
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err:%v", err)
	}
	//初始化数据引擎
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("initsetupDBEngine err:%v", err)
	}
	//初始化链路跟踪
	err = setupTracer()
	if err != nil {
		log.Fatalf("initsetupTracer err:%v", err)
	}
}

func setupFlag() {
	flag.BoolVar(&isVersion, "version", false, "展示编译信息") //获取版本构建信息
	flag.Parse()
}

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer("blog-service", "127.0.0.1:6831")
	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}

//获取配置内容->全局变量初始化
func setupSetting() error {
	setting, err := st.NewSetting()
	if err != nil {
		return err
	}

	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	global.AppSetting.RequestTimeOut *= time.Second

	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	global.JWTSetting.Expire *= time.Second

	global.ServerSetting.ReadTimeOut *= time.Second
	global.ServerSetting.WriteTimeOut *= time.Second

	return nil
}

//配置全局Gorm引擎
func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}

//@title 博客系统
//@version 1.0
//@description GO编程之旅
//@termsOfService https://github.com/noChaos1012/tour
func main() {

	if isVersion {
		//如果接收到构建信息指令则打印
		fmt.Printf("build_time:%s\n", buildTime)
		fmt.Printf("build_version:%s\n", buildVersion)
		fmt.Printf("git_commit_id:%s\n", gitCommitID)
		return
	}

	global.Logger.Infof(nil, "【服务启动】%s:tour/%s", "noChaos", "blog-service")

	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeOut,
		WriteTimeout:   global.ServerSetting.WriteTimeOut,
		MaxHeaderBytes: 1 << 20,
	}

	//安全退出，在接收到退出信号时保持五秒，处理已有请求
	go func() {
		err := s.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("s.ListenAndServe err:%v", err)
		}
	}()

	quit := make(chan os.Signal)                         //用以接收信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) //接收syscall.SIGINT,syscall.SIGTERM信号
	<-quit
	log.Println("Shutting down server ....")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) //保持五秒控制时间，处理原有请求
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown ", err)
	}
	log.Println("Server exiting")

}
