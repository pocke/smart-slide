package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type WriteJacker struct {
	buf    *bytes.Buffer
	header http.Header
	status int
}

func NewWriteJacker() *WriteJacker {
	return &WriteJacker{
		buf:    bytes.NewBuffer([]byte{}),
		header: make(map[string][]string),
		status: 200,
	}
}

func (wj *WriteJacker) Header() http.Header {
	return wj.header
}

func (wj *WriteJacker) WriteHeader(i int) {
	wj.status = i
}

func (wj *WriteJacker) Write(b []byte) (int, error) {
	return wj.buf.Write(b)
}

func (wj *WriteJacker) InjectScript(w http.ResponseWriter, script []byte) {
	var body []byte
	if !strings.Contains(wj.Header().Get("Content-Type"), "text/html") {
		body = wj.buf.Bytes()
	} else {
		s := strings.Replace(wj.buf.String(), "</body>", fmt.Sprintf("<script>%s</script></body>", script), 1)
		body = []byte(s)
		wj.header.Set("Content-Length", strconv.Itoa(len(body)))
	}

	for key, vals := range wj.header {
		for _, v := range vals {
			w.Header().Add(key, v)
		}
	}
	w.Write(body)
	w.WriteHeader(wj.status)
}

var _ http.ResponseWriter = &WriteJacker{}
