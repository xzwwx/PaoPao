package main

type UserOpt struct {
	Id         uint8
	hasMove    bool
	hasBreak   bool
	hasCombine bool
	Power      uint8
	Speed      uint16
	PosX       uint16
	PosY       uint16
	OptSeq     uint8
}

// Obstacle
type Obstacle struct {
	Id    uint32
	pos   Pos
	OType uint32 //0:wall   1: ke zha
}

type Pos struct {
	X float64
	Y float64
}

type Vector2 struct {
	x float64
	y float64
}
