package main

import (
	"net"
	"net/http"

	"golang.org/x/net/websocket"

	"github.com/pocke/hlog"
)

type WSServer struct {
	port int
	ch   chan string
}

func NewWSServer() (*WSServer, error) {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		return nil, err
	}

	s := &WSServer{
		port: l.Addr().(*net.TCPAddr).Port,
		ch:   make(chan string, 0),
	}
	go http.Serve(l, hlog.Wrap(websocket.Handler(s.handler).ServeHTTP))
	return s, nil
}

func (s *WSServer) handler(ws *websocket.Conn) {
	for msg := range s.ch {
		websocket.JSON.Send(ws, msg)
	}
}
