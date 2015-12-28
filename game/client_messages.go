package main

import (
	// "encoding/json"
	"github.com/etinlb/falcon/core_lib"
	// "github.com/etinlb/falcon/logger"
	// "github.com/etinlb/falcon/network"
	// "github.com/gorilla/websocket"
)

type UpdateMessage struct {
	core_lib.BaseRectData
	Id string
}
