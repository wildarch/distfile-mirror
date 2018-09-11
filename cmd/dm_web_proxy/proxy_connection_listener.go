package main

import (
	"net"
	"net/http"
)

// ProxyConnectionListener implements a net.Listener that yields
// connections generated by a ProxyConnectionHijacker.
type ProxyConnectionListener struct {
	connections chan net.Conn
}

var _ ProxyConnectionHandler = &ProxyConnectionListener{}
var _ net.Listener = &ProxyConnectionListener{}

func NewProxyConnectionListener() *ProxyConnectionListener {
	return &ProxyConnectionListener{
		connections: make(chan net.Conn),
	}
}

func (l *ProxyConnectionListener) Handle(c net.Conn, r *http.Request) {
	l.connections <- c
}

func (l *ProxyConnectionListener) Accept() (net.Conn, error) {
	return <-l.connections, nil
}

func (l *ProxyConnectionListener) Close() error {
	// TODO(edsch): What should we return here?
	return nil
}

func (l *ProxyConnectionListener) Addr() net.Addr {
	// TODO(edsch): What should we return here?
	return nil
}