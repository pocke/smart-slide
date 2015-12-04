package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"

	"golang.org/x/net/websocket"

	"github.com/pocke/hlog"
	"github.com/rs/cors"
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

	h := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	}).Handler(hlog.Wrap(websocket.Handler(s.handler).ServeHTTP))

	go http.Serve(l, h)
	return s, nil
}

func (s *WSServer) handler(ws *websocket.Conn) {
	close := func() <-chan error {
		ch := make(chan error)
		go func() {
			_, err := ioutil.ReadAll(ws)
			ch <- err
		}()
		return ch
	}()

	for {
		select {
		case msg := <-s.ch:
			websocket.JSON.Send(ws, msg)
		case <-close:
			return
		}
	}
}

func (s *WSServer) Script() ([]byte, error) {
	b, err := Asset("assets/ws.js")
	if err != nil {
		return nil, err
	}
	f := `(function(){`
	e := `})()`
	code := fmt.Sprintf("%svar port = %d;%s%s", f, s.port, string(b), e)
	return []byte(code), nil
}
