package main

import (
	"net/rpc"
	"time"

	"PaoPao/server-base/src/base/env"
	"PaoPao/server-base/src/base/gonet"
	"flag"
	"fmt"
	"strconv"

	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"

	"PaoPao/server/src/usercmd"
	"net"
)

type RCenterServer struct {
	gonet.Service
	rpcser  *gonet.TcpServer
	sockser *gonet.TcpServer
}

var serverm *RCenterServer

const RpcServiceName = "RetInfoRoom"

func RCenterServer_GetMe() *RCenterServer {
	if serverm == nil {
		serverm = &RCenterServer{
			rpcser:  &gonet.TcpServer{},
			sockser: &gonet.TcpServer{},
		}
		serverm.Derived = serverm
	}
	return serverm
}

///////rpc
type RetIntoRoom struct {
}

func (q *RetIntoRoom) RetRoom(request *usercmd.ReqIntoRoom, reply *usercmd.RetIntoFRoom) error {

	fmt.Println("Into RPC...")

	uid := request.GetUId()
	username := request.UserName
	fmt.Println(strconv.FormatInt(int64(uid), 10))
	fmt.Println(*username, "66666666")

	reply.Err = proto.Uint32(uint32(0))
	reply.RoomId = proto.Uint32(uint32(107))
	reply.Addr = proto.String("127.0.0.1:9494")

	reply.Key = proto.String(strconv.FormatInt(int64(uid), 10) + *username)

	fmt.Println("Out RPC")

	return nil
}

func Acc(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			glog.Errorln("[RCenterServer] accept error:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}

func (this *RCenterServer) Init() bool {
	rpc.RegisterName(RpcServiceName, new(RetIntoRoom))
	listener, err := net.Listen("tcp", env.Get("rcenter", "server"))
	if err != nil {
		glog.Errorln("[RCenterServer] listen error:", err)
		return false
	}
	go Acc(listener)
	return true
}

//

func (this *RCenterServer) Final() bool {
	return true
}

func (this *RCenterServer) Reload() {

}

func (this *RCenterServer) MainLoop() {
	time.Sleep(time.Second)
}

var config = flag.String("config", "", "config path")

func main() {
	flag.Parse()
	env.Load(*config)
	defer glog.Flush()
	RCenterServer_GetMe().Main()

	glog.Info("[Close] RCenterServer closed.")
}
