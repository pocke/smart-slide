package main

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"github.com/naoina/denco"
	"github.com/pocke/hlog"
)

type Controller struct {
	ch       chan<- string
	listener net.Listener
	port     int
}

func (c *Controller) KeyHandler(w http.ResponseWriter, r *http.Request, _ denco.Params) {
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	c.ch <- string(b)
}

func (c *Controller) IndexHandler(w http.ResponseWriter, r *http.Request, _ denco.Params) {
	//TODO
}

func (c *Controller) Serve() error {
	mux := denco.NewMux()
	h, err := mux.Build([]denco.Handler{
		mux.GET("/", c.IndexHandler),
		mux.POST("/key", c.KeyHandler),
	})
	if err != nil {
		return err
	}

	return http.Serve(c.listener, hlog.Wrap(h.ServeHTTP))
}

func NewController(ch chan<- string) (*Controller, error) {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		return nil, err
	}
	c := &Controller{
		ch:       ch,
		listener: l,
	}

	go func() {
		if err := c.Serve(); err != nil {
			panic(err)
		}
	}()
	return c, nil
}
