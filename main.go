package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/filedrive-team/filplus-info/common"
	"github.com/filedrive-team/filplus-info/jobs"
	"github.com/filedrive-team/filplus-info/log"
	"github.com/filedrive-team/filplus-info/models"
	"github.com/filedrive-team/filplus-info/routers"
	"github.com/filedrive-team/filplus-info/settings"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

var (
	showV      bool
	configFile string
	loglevel   string
	initTable  bool
)

func printVersion() {
	fmt.Printf("version: v%s.%s.%s-%s\ngithub.com/gin-gonic/gin version: %s\n",
		Major, Minjor,
		Patch, BuildVersion, gin.Version)
}

func Init(configFile string) {
	settings.Setup(configFile)
	conf := settings.AppConfig

	models.Setup(conf)
	// Initialize the global cache
	common.InitGlobalCache()
}

func main() {
	flag.BoolVar(&showV, "version", false, "print version")
	flag.StringVar(&configFile, "config", "conf/app.toml", "set config file path")
	flag.StringVar(&loglevel, "loglevel", "info", "set log level")
	flag.BoolVar(&initTable, "init", false, "init table data")
	flag.Parse()

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(fmt.Errorf("fetch work dir failed %+v", err))
	}

	if showV {
		printVersion()
		os.Exit(0)
	}

	err = log.InitLogger(dir, loglevel)
	if err != nil {
		panic(fmt.Errorf("init logger failed %+v", err))
	}

	Init(configFile)
	defer models.CloseDB()

	// Database table initialization operation
	if initTable {
		models.TruncateNotary()
		models.TruncateClient()
		models.InitNotaryList()
	}

	router := routers.InitRouter()
	conf := settings.AppConfig
	addr := fmt.Sprintf(":%d", conf.Server.HttpPort)
	maxHeaderBytes := 1 << 20

	srv := &http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    time.Duration(conf.Server.ReadTimeout * 1e9),
		WriteTimeout:   time.Duration(conf.Server.WriteTimeout * 1e9),
		MaxHeaderBytes: maxHeaderBytes,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("listen: %s\n", err)
		}
	}()

	mainCtx, mainCancel := context.WithCancel(context.Background())
	syncer := jobs.NewSyncer()
	go syncer.Run(mainCtx)

	crawler := jobs.NewCrawler()
	go crawler.Run(mainCtx)

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutdown Server ...")

	mainCancel()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server Shutdown: ", err)
	}

	logger.Info("Server exiting")
}
