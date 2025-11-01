/* ----- ----- ----- ----- */
// packet.go
// Do not distribute or modify
// Author: DragonTaki (https://github.com/DragonTaki)
// Create Date: 2025/11/01
// Update Date: 2025/11/01
// Version: v1.0
/* ----- ----- ----- ----- */

package server

import "encoding/json"

// Packet format
type Packet struct {
	Type     PacketType `json:"type"`
	SenderId string     `json:"senderId"`
	RoomId   string     `json:"roomId"`
	Data     string     `json:"data"`
	Token    string     `json:"token"`
}

func CreatePacket(pktType PacketType, senderId, roomId, data, token string) *Packet {
	return &Packet{
		Type:     pktType,
		SenderId: senderId,
		RoomId:   roomId,
		Data:     data,
		Token:    token,
	}
}

// Serialize 將 Packet 轉成 JSON 字串
func (p *Packet) SerializePacket() string {
	b, _ := json.Marshal(p)
	return string(b)
}

// Deserialize 將 JSON 字串轉回 Packet
func DeserializePacket(jsonStr string) (*Packet, error) {
	var p Packet
	err := json.Unmarshal([]byte(jsonStr), &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}
