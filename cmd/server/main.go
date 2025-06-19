package server

import (
	"fmt"
	"net"
	"net/http"

	logger "github.com/isayme/go-logger"
	"github.com/isayme/websockify-go/websockify"
	"golang.org/x/net/websocket"
)

type ServerOptions struct {
	Listen   *string
	Vnc      *string
	Web      *string
	cert     *string
	certKey  *string
	sslOnly  *bool
	fileOnly *bool
	webAuth  *string
}

func Run(options ServerOptions) {
	http.Handle("/websockify", websocket.Server{
		Handshake: handshakeWebsocket,
		Handler:   handleWebsocket(*options.Vnc),
	})

	fs := http.FileServer(http.Dir(*options.Web))
	http.Handle("/", http.StripPrefix("/", fs))

	logger.Debugw("start listen", "address", options.Listen)
	if err := http.ListenAndServe(*options.Listen, nil); err != nil {
		logger.Panicw("ListenAndServe fail", "err", err)
	}
}

func handshakeWebsocket(config *websocket.Config, req *http.Request) error {
	var err error
	config.Origin, err = websocket.Origin(config, req)
	if err == nil && config.Origin == nil {
		return fmt.Errorf("null origin")
	}
	return err
}

func handleWebsocket(target string) func(*websocket.Conn) {
	return func(ws *websocket.Conn) {
		ws.PayloadType = websocket.BinaryFrame

		client := ws.Request().RemoteAddr
		logger.Infow("new connection", "target", target, "client", client)

		conn, err := net.Dial("tcp", target)
		if err != nil {
			logger.Warnw("dial service fail", "remote", target, "err", err)
			ws.Close()
			return
		}
		defer conn.Close()

		tcpConn, _ := conn.(*net.TCPConn)
		websockify.Proxy(ws, tcpConn)

		logger.Infow("connection close", "client", client, "target", target)
	}
}
