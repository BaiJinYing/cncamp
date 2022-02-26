package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type HttpHandler struct{}

func (h HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for k, vals := range r.Header {
		for _, v := range vals {
			w.Header().Add(k, v)
		}
	}
	version := os.Getenv("VERSION")
	w.Header().Set("version", version)

	code := 200
	message := ""

	switch {
	case r.URL.Path == "/":
		message = "Hello World"
	case r.URL.Path == "/healthz":
	default:
		code = 404
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err == nil {
		fmt.Println("请求路径:", r.URL.Path, "客户端 IP:", ip, "HTTP 返回码:", code)
	}

	w.WriteHeader(code)
	fmt.Fprintf(w, message)
}

func main() {
	handler := HttpHandler{}
	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}
	go server.ListenAndServe()

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println(<-ch)

	server.Shutdown(context.Background())
}
