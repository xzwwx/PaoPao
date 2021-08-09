package main

import "time"

const (
	RoomWidth  = 200
	RoomHeight = 200
	CellWidth  = 5
	CellHeight = 5

	RoomSize = 200
)

type Scene struct {
	players  map[uint64]*ScenePlayer
	room     *Room
	Obstacle *map[uint32]*Obstacle

	gameMap Map
	// Scene info
	sceneWidth  float64
	sceneHeight float64
	now         time.Time
	startTime   time.Time
	frame       uint32
	ballID      uint32
	pool        *BallPool // Player Pool

	msgBytes []byte
	// temple things
	rangeBalls     []*Bomb //map
	rangePlayers   []*ScenePlayer
	rangeObstacles []*Obstacle
}

func NewScene(room *Room) *Scene {
	scene := &Scene{
		room:    room,
		players: make(map[uint64]*ScenePlayer),
	}
	scene.Init()
	return scene
}

func (this *Scene) Init() {
	this.Obstacle = GenerateRandMap()
	// Init map
	this.gameMap = Map{}
}

func (this *Scene) AddPlayer(p *PlayerTask) {
	this.players[p.id] = NewScenePlayer(p, this)
}

func (this *Scene) SendRoomMsg() {
	for _, p := range this.players {
		p.sendSceneMsg(this)
	}
}

func (this *Scene) UpdatePlayers(per float64) {

	// Depose player logic
	//if this.room.roomType
	//for i := 0; i < len(this.players); i++ {
	//	player, _ := this.players[]
	//}
	for _, player := range this.players {
		player.Update(per, this)
	}

}

// Check mapcell   0: null  1:wall  2:Obstacle
func (this *Scene) GetCellState(x, y float64) int32 {

	return this.gameMap.gamemap[x][y]
}

// 保证位置不超出地图范围
func (this *Scene) BorderCheck(pos *Vector2) {
	if pos.x < 0 {
		pos.x = 0
	} else if pos.x >= this.sceneWidth {
		pos.x = this.sceneWidth - 0.01
	}
	if pos.y < 0 {
		pos.y = 0
	} else if pos.y >= this.sceneHeight {
		pos.y = this.sceneHeight - 0.01
	}
}
