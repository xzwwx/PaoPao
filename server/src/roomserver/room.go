package main

type Room struct {
	//mutex    sync.RWMutex
	id          uint32                 //房间id
	roomType    uint32                 //房间类型
	players     map[uint32]*PlayerTask //房间内的玩家
	curNum      uint32                 //当前房间内玩家数
	bombCount 	uint32
	isStart     bool
	timeLoop    uint64
	stopCh      chan bool
	isStop      bool
	iscustom	bool

	max_num		uint32			//max player number. default :8
	totalTime 	uint64 //in second
	endchan     chan bool
}

//type PlayerOpt struct {
//	pTask 		*PlayerTask
//	pPlayer 	*ScenePlayer
//	Opts		*UserOpt
//}

func NewRoom(rtype, rid uint32, player *PlayerTask) *Room{
	room := &Room{

	}
	return room
}

func (this *Room) Start()bool{




	return true
}

func (this *Room) Stop() bool {



	return true
}

func (this *Room) IsClosed() bool {



	return true
}

func (this *Room) Loop() {


}