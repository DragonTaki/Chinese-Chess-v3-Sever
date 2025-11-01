/* ----- ----- ----- ----- */
// auth.go
// Do not distribute or modify
// Author: DragonTaki (https://github.com/DragonTaki)
// Create Date: 2025/11/01
// Update Date: 2025/11/01
// Version: v1.0
/* ----- ----- ----- ----- */

package server

import (
	"bufio"
	"time"

	"Chinese-Chess-v3-Sever/logger"
)

type AuthMessage struct {
	Type     string `json:"type"`
	SenderId string `json:"id"`
	Version  string `json:"version"`
}

// Authenticate 監控 client 在 timeout 內是否傳送驗證訊息 (HELLO)
// 成功驗證回傳 true，失敗或超時回傳 false
func Authenticate(c *Client, timeout time.Duration) bool {
	authCh := make(chan bool, 1)

	go func() {
		scanner := bufio.NewScanner(c.Connection)
		for scanner.Scan() {
			line := scanner.Text()

			// Deserialize JSON packet
			pkt, err := DeserializePacket(line)
			if err != nil {
				logger.Warnf("Invalid packet from %s: %v", c.RemoteAddr, err)
				continue
			}

			// Check if Auth packet
			if pkt.Type != PacketTypeAuthRequest {
				logger.Warnf("Unexpected packet type from %s: %s", c.RemoteAddr, pkt.Type)
				continue
			}

			// SenderId cannot be empty
			if pkt.SenderId == "" {
				logger.Warnf("Missing senderId from %s", c.RemoteAddr)
				authCh <- false
				return
			}

			// Check if version match
			if pkt.Data != ServerVersion {
				logger.Warnf("Version mismatch from %s: %s != %s",
					c.RemoteAddr, pkt.Data, ServerVersion)
				authCh <- false
				return
			}

			// Auth success
			c.SenderId = pkt.SenderId
			c.IsAuthenticated = true
			c.LastSeenAt = time.Now()

			respPkt := CreatePacket(
				PacketTypeAuthResponse,
				"Server",
				"",
				AuthSuccessString,
				"",
			)
			c.SendPacket(respPkt)

			authCh <- true
			return
		}
	}()

	select {
	case ok := <-authCh:
		if ok {
			logger.Infof("Client %s authenticated successfully", c.RemoteAddr)
		}
		return ok

	case <-time.After(timeout):
		logger.Warnf("Client %s failed to authenticate in time", c.RemoteAddr)
		return false
	}
}
