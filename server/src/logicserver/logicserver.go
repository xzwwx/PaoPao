package main

import (
	"PaoPao/server-base/src/base/env"
	"PaoPao/server-base/src/base/gonet"
	"flag"
	"time"

	"github.com/golang/glog"
)

type LogicServer struct {
	gonet.Service
}

var logicServer *LogicServer

func LogicServer_GetMe() *LogicServer {
	if logicServer == nil {
		logicServer = &LogicServer{}
	}
	logicServer.Derived = logicServer
	return logicServer
}

func (this *LogicServer) Init() bool {
	if !redisMgr.NewRedisManager() || !InitHttpServer() {
		glog.Errorln("[LogicServer Init] Init error")
		return false
	}
	glog.Errorln("[LogicServer Init] Init success")
	return true
}

func (this *LogicServer) Reload() {

}

func (this *LogicServer) MainLoop() {
	time.Sleep(time.Second)
}

func (this *LogicServer) Final() bool {
	redisMgr.pool.Close()
	return true
}

var (
	port   = flag.String("port", "8000", "logicserver listen port")
	config = flag.String("config", "", "config json file path")
)

func main() {
	flag.Parse()
	env.Load(*config)
	defer glog.Flush()
	// 从命令行参数获取指定端口号

	LogicServer_GetMe().Main()
}
