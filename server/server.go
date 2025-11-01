/* ----- ----- ----- ----- */
// server.go
// Do not distribute or modify
// Author: DragonTaki (https://github.com/DragonTaki)
// Create Date: 2025/11/01
// Update Date: 2025/11/01
// Version: v1.0
/* ----- ----- ----- ----- */

package server

import (
	"Chinese-Chess-v3-Sever/logger"
	"fmt"
	"net"
	"sync"
	"time"
)

const timeoutLimit = time.Minute

type Server struct {
	clients map[*Client]bool
	mu      sync.Mutex
}

func NewServer() *Server {
	return &Server{
		clients: make(map[*Client]bool),
	}
}

// 處理新連線
func (s *Server) HandleNewClient(conn net.Conn) {
	client := NewClient(conn, s)
	client.LastSeen = time.Now()

	s.mu.Lock()
	s.clients[client] = true
	s.mu.Unlock()

	logger.Infof("New client connected: %s", conn.RemoteAddr().String())

	go s.monitorHeartbeat(client) // 啟動心跳檢查
	client.Listen()
}

// 廣播訊息給所有玩家
func (s *Server) Broadcast(sender *Client, msg string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for c := range s.clients {
		if c != sender {
			c.Send(fmt.Sprintf("[%s] %s", sender.Addr, msg))
		}
	}
}

// 客戶端離線時移除
func (s *Server) RemoveClient(c *Client) {
	s.mu.Lock()
	delete(s.clients, c)
	s.mu.Unlock()
	logger.Warnf("Client disconnected: %s\n", c.Addr)
}

// --- 心跳監控 ---
func (s *Server) monitorHeartbeat(c *Client) {
	ticker := time.NewTicker(10 * time.Second) // 每 10 秒檢查一次
	defer ticker.Stop()

	for range ticker.C {
		if time.Since(c.LastSeen) > timeoutLimit {
			logger.Warnf("Client timed out: %s", c.Addr)
			c.Conn.Close()    // 關閉連線
			s.RemoveClient(c) // 從伺服器移除
			break
		}
	}
}
