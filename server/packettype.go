/* ----- ----- ----- ----- */
// packettype.go
// Do not distribute or modify
// Author: DragonTaki (https://github.com/DragonTaki)
// Create Date: 2025/11/01
// Update Date: 2025/11/01
// Version: v1.0
/* ----- ----- ----- ----- */

package server

type PacketType string

const (
	PacketTypeNotDefined PacketType = "NotDefined"

	// Auth
	PacketTypeAuthRequest  PacketType = "AuthRequest"
	PacketTypeAuthResponse PacketType = "AuthResponse"

	// Room
	PacketTypeJoinRoom  PacketType = "JoinRoom"
	PacketTypeLeaveRoom PacketType = "LeaveRoom"

	// Chess game
	PacketTypeStartGame  PacketType = "StartGame"
	PacketTypeEndGame    PacketType = "EndGame"
	PacketTypeGameAction PacketType = "GameAction"
	PacketTypeTimerSync  PacketType = "TimerSync"

	// Chat
	PacketTypeChat PacketType = "Chat"

	// Other
	PacketTypeServer    PacketType = "Server"
	PacketTypeHeartbeat PacketType = "Heartbeat"
	PacketTypeError     PacketType = "Error"
)
