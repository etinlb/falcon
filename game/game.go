package main

import (
	// "bytes"
	// "encoding/json"
	"flag"
	"fmt"
	// "io/ioutil"
	"net/http"
	"os"
	// "strconv"

	"github.com/etinlb/falcon/logger"
	"github.com/etinlb/falcon/network"
	"github.com/gorilla/websocket"
)

// various object maps to keep track of different types of objects
var gameObjects map[string]network.NetworkedGameObjects

// var playerObjects map[string]*Player
// var physicsComponents map[string]*PhysicsComponent

// Communication coordinator
// var channelCoordinator ComunicationChannels

// Connection structures
var connections map[*websocket.Conn]int // Maps the connection object to the client id
var clientIdMap map[int]*network.ClientData

// map that keeps track of what data came from what client

var clientBackend network.NetworkController

func cleanUpSocket(conn *websocket.Conn) *network.Message {
	logger.Info.Printf("Cleaning up connection from %s\n", conn.RemoteAddr())
	return nil
	// logger.Info.Println(len(gameObjects))
	// clientId := connections[conn]
	// clientData := clientIdMap[clientId]
	// for id, _ := range clientData.GameObjects {
	// 	Trace.Println("deleting from gameObjects map, id: %s", id)
	// 	// TODO: Need to delete all references to this object...Idk how to best
	// 	// do that.
	// 	delete(gameObjects, id)
	// 	delete(playerObjects, id)
	// 	delete(physicsComponents, id)
	// }

	// delete(clientIdMap, clientData.ClientId)
	// delete(connections, conn)
	// logger.Info.Println(len(gameObjects))

	// conn.Close()
	// printGameObjectMap()
}

func initializeGameData() {
	logger.Trace.Println("Initalizing all game data")
	// keyed by id
	gameObjects = make(map[string]network.NetworkedGameObjects)
	// playerObjects = make(map[string]*Player)
	// physicsComponents = make(map[string]*PhysicsComponent)
}

func initializeLogger() {
	// TODO: Read a config file
	logger.InitLogger(os.Stdout, os.Stdout, os.Stdout, os.Stderr)
	// InitLogger(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
}

// TODO: SHould this be in server vars?
func initializeConnectionData() {
	logger.Trace.Println("Initailize connection varibles")
	// TODO: Access if we need the clients variable
	// clients = make(map[*websocket.Conn]*ClientData)
	connections = make(map[*websocket.Conn]int)
	clientIdMap = make(map[int]*network.ClientData)
}

func main() {
	initializeLogger()

	port := flag.Int("port", 8080, "port to serve on")
	// TODO: have this address passed from the other server
	dir := flag.String("directory", "../web/", "directory of web files")
	interactive := flag.Bool("i", false, "Run with an interactive shell")
	flag.Parse()

	// =========Game Initializations============================================
	initializeGameData()

	// =========Connection Initializations======================================
	initializeConnectionData()

	clientBackend = network.NewNetworkController(HandleClientEvent,
		cleanUpSocket,
		initializeClientData)

	// handle all requests by serving a file of the same name
	fs := http.Dir(*dir)
	fileHandler := http.FileServer(fs)
	http.Handle("/", fileHandler)
	http.HandleFunc("/ws", clientBackend.WsHandler)
	// the socket to read incoming connections from the master server

	logger.Info.Printf("Running on port %d\n", *port)

	addr := fmt.Sprintf("0.0.0.0:%d", *port)
	StartGameLoop(clientBackend)

	// Add channels to the channel coordinator
	// channelCoordinator = ComunicationChannels{
	// 	moveChannel:         moveChannel,
	// 	addChannel:          addChannel,
	// 	broadcastAddChannel: broadcastAddChannel}

	if *interactive {
		go runHttpServer(addr)
		// runInteractiveMode(channelCoordinator)
	} else {
		// function runs until an exit is called
		runHttpServer(addr)
	}
}

func runHttpServer(addr string) {
	// this call blocks -- the progam runs here forever
	err := http.ListenAndServe(addr, nil)
	logger.Warning.Println(err.Error())
}

// func printGameObjectMap(gameObjects []core_lib.GameObject) {
// 	for _, obj := range gameObjects {
// 		logger.Info.Println(obj)
// 	}
// }
