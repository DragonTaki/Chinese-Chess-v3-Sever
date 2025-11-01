/* ----- ----- ----- ----- */
// constants.go
// Do not distribute or modify
// Author: DragonTaki (https://github.com/DragonTaki)
// Create Date: 2025/11/01
// Update Date: 2025/11/01
// Version: v1.0
/* ----- ----- ----- ----- */

package server

import "time"

const ServerVersion = "v1.0.0"
const AuthSuccessString = "Taki"

const (
	AuthTimeoutLimit             = 10 * time.Second
	ClientTimeoutLimit           = 1 * time.Minute
	ClientHeartbeatCheckInterval = 10 * time.Second
	ServerHeartbeatSendInterval  = 3 * time.Second
)
