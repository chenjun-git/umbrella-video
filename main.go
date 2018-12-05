package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/chenjun-git/umbrella-common/monitor"
	commonUtils "github.com/chenjun-git/umbrella-common/utils"

	"github.com/chenjun-git/umbrella-video/common"
	"github.com/chenjun-git/umbrella-video/db"
	"github.com/chenjun-git/umbrella-video/handler"
)

var (
	BuildTime    = "No Build Time"
	BuildGitHash = "No Build Git Hash"
	BuildGitTag  = "No Build Git Tag"
)

func main() {
	initConfig()
	initMonitor()

	startHTTPServer()
}

func initConfig() {
	configPath := flag.String("config", "", "config file's path")
	flag.Parse()

	common.InitConfig(*configPath)

	// init mysql and redis
	db.InitMySQL(common.Config.Mysql)
}

func initMonitor() {
	monitor.Init(common.Config.Monitor.Namespace, common.Config.Monitor.Subsystem)
	monitor.Monitor.SetVersion(monitor.Version{GitHash: BuildGitHash, GitTag: BuildGitTag, BuildTime: BuildTime})
}

func startHTTPServer() {
	router := handler.RegisterVideoRouter()
	commonUtils.RegisterPProf(router)
	monitor.RegisterHandlers(router)

	httpServer := &http.Server{
		Addr:    common.Config.HTTP.Listen,
		Handler: router,
	}
	fmt.Printf("start http server, listen: %s!\n", common.Config.HTTP.Listen)
	if err := httpServer.ListenAndServe(); err != nil {
		fmt.Printf("start http server failed, err : %+v\n", err)
		panic("start http server failed")
	}
}
