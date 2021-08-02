package main

import (
	"base/gonet"
	"context"
	"crypto/md5"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-redis/redis"
	//"base/env"
	"github.com/golang/protobuf/proto"
	"glog"
	"io"
	"log"
	"net/rpc"
	"strconv"

	//"math/rand"
	"net/http"
	//"strconv"
	"strings"
	"time"
	"usercmd"
)

type LoginServer struct {
	gonet.Service


	ServerId	uint32
	//rpcserver 	*rpc.
	//rpcserver 	*grpc.RpcServer
}

var serverm *LoginServer

func LoginServer_GetMe() *LoginServer {
	if serverm == nil {
		serverm = &LoginServer{
			//rpcserver: &grpc.RpcServer{},
		}
		serverm.Derived = serverm
	}
	return serverm
}

///////
//
type RpcReqData struct {
	uid	        uint64
	username 	string
	//...
}
func (this *RpcReqData) VerifyHandle()bool {

	return false
}

func (this *RpcReqData) OnClose(){

}

func (handler *RpcReqData) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	client, err := rpc.Dial("tcp", "127.0.0.1:9099")

	if err != nil {
		log.Fatal("dialing:", err)
	}
	fmt.Println("------Dial ok")
	defer func() {
		if err := client.Close();err == nil {
			fmt.Println(" Closed 333333333333333333333")
		}
	}()

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read request.Body failed, err:%v\n", err)
		return
	}
	fmt.Println("------Get data ok")

	//fmt.Println(body)
	//fmt.Println(string(body))

	data:= make(map[string]string)
	err = json.Unmarshal(body, &data)
	if err !=nil {
		fmt.Println("body To Json failed, err:", err)
		panic("Error")
	}
	fmt.Println("++++++++++++++",data)

	//w.Write([]byte(data["username"]))


	uid, _ := strconv.Atoi(data["uid"])
	var reply = &usercmd.RetIntoFRoom{}
	var param = &usercmd.ReqIntoRoom{
		UId: proto.Uint64(uint64(uid)),
		UserName: proto.String(data["username"]),
	}
	if reply.GetKey() == ""{
		fmt.Println("------------",uid, "--" )
	}


	err = client.Call("RetIntoRoom.RetRoom", &param, &reply)
	fmt.Println("-----Call end-------",uid, "--" )

	if err != nil {
		log.Fatal(err)
	}
	if reply.GetKey() == ""{
		fmt.Println("----////--",uid, "--" )
	}

	w.Write([]byte(reply.GetKey()))

}

////////////
func (this *LoginServer) Init() bool {


	// Start RPC
	HandleLogic := new(RpcReqData)
	//inlisten := env.Get("login", "inlisten")
	//err := this.rpcserver.BindAccept

	http.Handle("/", HandleLogic)
	fmt.Println("------Handle ok")
	go http.ListenAndServe("127.0.0.1:8090", nil)
	fmt.Println("------Listen ok")



	//if err != nil {
	//	glog.Error("[Start] Listen Error. ", inlisten, ", ", err)
	//	return false
	//}

	//Start Http server
	if !StartHttpServer() {
		glog.Error("[Start] Start Http error.")
		return false
	}
	glog.Info("[Start] Initialization successful.")
	return true

}

func (this *LoginServer) MainLoop(){
	time.Sleep(time.Second)
}

func (this *LoginServer) Final() bool {



	return true
}

func (this *LoginServer) Reload(){

}


var (
	logfile = flag.String("logfile", "","Log file name")
	config = flag.String("config", "config.json","config path")
)

func main() {
	flag.Parse()
	//
	//if !env.Load(*config){
	//	return
	//}
	//
	//loglevel := env.Get("global", "loglevel")
	//if loglevel != "" {
	//	flag.Lookup("stderrthreshold").Value.Set(loglevel)
	//}
	//
	//logtostderr := env.Get("global", "logtostderr")
	//if loglevel != "" {
	//	flag.Lookup("logtostderr").Value.Set(logtostderr)
	//}
	//
	//rand.Seed(time.Now().Unix())
	//
	//if *logfile != ""{
	//	glog.SetLogFile(*logfile)
	//}else{
	//	glog.SetLogFile(env.Get("room","log"))
	//}
	//
	//defer glog.Flush()

	LoginServer_GetMe().Main()

	glog.Info("[Close] Login Server closed.")
}





//////////////////////////////////////////////////
type User struct{
	userName string
	userEmail string
	userPasswd string
	userTel string
	userKey string
	isNew bool
}

func password_md5(password, salt string)string{
	h := md5.New()
	h.Write([]byte(salt + password))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func createClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "123456",
		DB:       0,
		PoolSize: 100,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	// 通过 cient.Ping() 来检查是否成功连接到了 redis 服务器
	_, err := client.Ping(ctx).Result()
	if err != nil{
		panic(err)
	}

	return client
}

func userLogin(client *redis.Client, userName , email, password, salt string)(isLogin, getName string){
	ctx := context.Background()
	if userName != ""{

		uidKey, err := client.Get(ctx, "userName:"+strings.ToLower(strings.TrimSpace(userName))).Result()
		if err == redis.Nil{
			return "userName does not exist", "00000000"
		}

		getpwd, err := client.HGet(ctx, uidKey, "password").Result()
		if err != nil{
			panic(err)
		}

		if getpwd == password_md5(password, salt){
			return "login sucessfully", userName
		}

	}else if email != ""{
		uidKey, err := client.Get(ctx, "email:"+email).Result()
		if err == redis.Nil{
			return "email does not exist", "00000000"
		}

		getpwd, err := client.HGet(ctx, uidKey, "password").Result()
		if err != nil{
			panic(err)
		}

		if getpwd == password_md5(password, salt){
			userName, err = client.HGet(ctx, uidKey, "userName").Result()
			if err != nil{
				panic(err)
			}
			return "login sucessfully", userName
		}

	}
	return "fail", "00000000"
}

// login POST
func login(w http.ResponseWriter, r *http.Request){
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read request.Body failed, err:%v\n", err)
		return
	}
	data:= make(map[string]string)
	err = json.Unmarshal([]byte(body), &data)
	if err !=nil {
		fmt.Println("body To Json failed, err:", err)
		panic("Error")
	}
	fmt.Println(data)

	//login
	client := createClient()
	fmt.Println(client)

	var newResister User
	newResister.userName = data["userName"]
	newResister.userEmail = data["userEmail"]
	newResister.userPasswd = data["userPasswd"]
	newResister.userTel = data["userTel"]
	newResister.userKey = data["userKey"]

	status, userName := userLogin(client, newResister.userName, "", newResister.userPasswd, newResister.userName)
	if status == "login sucessfully" {
		fmt.Println(userName)
	}
	answer := `{"Login": "ok"}`
	w.Write([]byte(answer))

}
//
//func main2(){
//	//client := createClient()
//	//fmt.Println(client)
//	//
//	//register(client)
//	//status, userName := userLogin(client,"oceanstar", "", "123456", "yan")
//	//if status == "login sucessfully" {
//	//	fmt.Println(userName)
//	//}
//
//	http.HandleFunc("/rigester", rigester)
//	http.HandleFunc("/login", login)
//	err := http.ListenAndServe(":9090", nil)
//	log.Fatal(err)
//
//}
