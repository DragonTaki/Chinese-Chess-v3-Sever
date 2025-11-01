/* ----- ----- ----- ----- */
// client.go
// Do not distribute or modify
// Author: DragonTaki (https://github.com/DragonTaki)
// Create Date: 2025/11/01
// Update Date: 2025/11/01
// Version: v1.0
/* ----- ----- ----- ----- */

package server

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

type Client struct {
	Conn     net.Conn
	Addr     string
	Srv      *Server
	LastSeen time.Time
}

func NewClient(conn net.Conn, srv *Server) *Client {
	return &Client{
		Conn: conn,
		Addr: conn.RemoteAddr().String(),
		Srv:  srv,
	}
}

func (c *Client) Listen() {
	defer func() {
		c.Conn.Close()
		c.Srv.RemoveClient(c)
	}()

	scanner := bufio.NewScanner(c.Conn)
	c.Send("Welcome to Go-Chess-Server! Type message to chat.\n")

	for scanner.Scan() {
		msg := scanner.Text()
        c.LastSeen = time.Now()
		if msg == "/quit" {
			break
		}
		c.Srv.Broadcast(c, msg)
	}
}

func (c *Client) Send(msg string) {
	fmt.Fprintln(c.Conn, msg)
}
