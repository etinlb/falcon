package main

import (
	"encoding/json"
	"github.com/etinlb/falcon/core_lib"
	// "github.com/etinlb/falcon/logger"
	"github.com/etinlb/falcon/network"
	// "time"
)

type GameObject interface {
	Update()
	BuildUpdateMessage() network.Message
}

// ==============TODO: Evaluate where player should be. ==============
type Player struct {
	core_lib.BaseGameObjData
	PhysicsComp *core_lib.PhysicsComponent
}

func (p *Player) Move(xAxis, yAxis float64) {
	p.PhysicsComp.Move(xAxis, yAxis)
}

func (p *Player) Update() {
	p.PhysicsComp.Move(0, 0)
}

func (p *Player) BuildUpdateMessage() network.Message {
	location := core_lib.NewVector(p.PhysicsComp.Location.X, p.PhysicsComp.Location.Y)
	veloctiy := core_lib.NewVector(p.PhysicsComp.Velocity.X, p.PhysicsComp.Velocity.Y)
	rectData := core_lib.BaseRectData{Velocity: veloctiy, Location: location}
	update := UpdateMessage{BaseRectData: rectData, Id: p.Id}

	// TODO: Check error
	rawData, _ := json.Marshal(update)
	rawJsonData := json.RawMessage(rawData)

	return network.Message{Event: "update", Data: &rawJsonData}
}

func NewPlayer(x, y float64) Player {
	physicsComponenet := core_lib.NewPhysicsComponent(x, y)
	physicsComponenet.Dimensions = core_lib.NewVector(20, 20)

	// TODO: Make a source id?
	baseData := core_lib.NewBaseGameObjData()
	playerObject := Player{
		PhysicsComp:     &physicsComponenet,
		BaseGameObjData: baseData,
	}

	return playerObject
}
