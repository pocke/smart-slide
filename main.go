package main

import (
	"fmt"
	"net"
	"net/http"

	"github.com/pocke/hlog"
	"github.com/skratchdot/open-golang/open"
)

func main() {
	if err := Main(); err != nil {
		panic(err)
	}
}

func Main() error {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		return err
	}
	url := fmt.Sprintf("http://127.0.0.1:%d", l.Addr().(*net.TCPAddr).Port)
	open.Run(url)
	fmt.Println(url)

	ws, err := NewWSServer()
	if err != nil {
		return err
	}
	c, err := NewController(ws.ch)
	url = fmt.Sprintf("http://127.0.0.1:%d\n", c.listener.Addr().(*net.TCPAddr).Port)
	open.Run(url)

	script, err := ws.Script()
	if err != nil {
		return err
	}

	return http.Serve(l, hlog.Wrap(func(w http.ResponseWriter, r *http.Request) {
		wj := NewWriteJacker()
		http.ServeFile(wj, r, "."+r.URL.Path)
		wj.InjectScript(w, script)
	}))
}
