package websockify

import (
	"net"
	"sync"

	logger "github.com/isayme/go-logger"
	"golang.org/x/net/websocket"
)

func Proxy(client *websocket.Conn, remote *net.TCPConn) {
	// see https://stackoverflow.com/a/75418345/1918831
	wg := sync.WaitGroup{}
	wg.Add(2)

	clientAddr := client.Request().RemoteAddr

	go func() {
		defer wg.Done()

		var err error
		var n int64
		n, err = Copy(remote, client)
		logger.Debugw("copy from client end", "n", n, "err", err, "client", clientAddr)
		remote.CloseWrite()
	}()

	go func() {
		defer wg.Done()

		var err error
		var n int64
		n, err = Copy(client, remote)
		logger.Debugw("copy from remote end", "client", clientAddr, "n", n, "err", err)
		client.WriteClose(0)
	}()

	wg.Wait()
}
