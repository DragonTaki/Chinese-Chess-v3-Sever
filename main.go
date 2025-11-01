/* ----- ----- ----- ----- */
// main.go
// Do not distribute or modify
// Author: DragonTaki (https://github.com/DragonTaki)
// Create Date: 2025/11/01
// Update Date: 2025/11/01
// Version: v1.0
/* ----- ----- ----- ----- */

package main

import (
	"Chinese-Chess-v3-Sever/logger"
	"Chinese-Chess-v3-Sever/server"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		logger.Errorf("Failed to start server: %v", err)
	}
	defer listener.Close()

	logger.Infof("Chess server started at 127.0.0.1:8080")

	srv := server.NewServer()
	srv.StartHeartbeatSystem()

	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Errorf("Connection error:", err)
			continue
		}
		go srv.HandleNewClient(conn)
	}
}
