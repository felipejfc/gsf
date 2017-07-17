package connector

import (
	"bufio"
	"errors"
	"fmt"
	"net"
)

// TCPConnector struct
type TCPConnector struct {
	onNewClientConnectionCB func(c Client)
	onClientDisconnectCB    func(c Client, err error)
}

// TCPClient struct
type TCPClient struct {
	conn     net.Conn
	server   *TCPConnector
	reader   *bufio.Reader
	socketID string
}

// NewTCPClient ctor
func NewTCPClient(conn net.Conn, server *TCPConnector) *TCPClient {
	return &TCPClient{
		conn:   conn,
		server: server,
		reader: bufio.NewReader(conn),
	}
}

// Close closes the tcpclient
func (c *TCPClient) Close() {
	c.conn.Close()
	c.server.onClientDisconnectCB(c, errors.New("connection closed"))
}

// Send sends a message to the socket
func (c *TCPClient) Read() ([]byte, error) {
	reader := c.reader
	message, err := reader.ReadBytes('\n')
	return message, err
}

// Send sends a message to the socket
func (c *TCPClient) Send(data []byte) error {
	_, err := c.conn.Write(data)
	return err
}

// NewTCPConnector ctor
func NewTCPConnector() *TCPConnector {
	return &TCPConnector{}
}

// Listen listens
func (t *TCPConnector) Listen(host string, port int) error {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return err
	}
	defer listener.Close()
	for {
		conn, _ := listener.Accept()
		client := NewTCPClient(conn, t)
		t.onNewClientConnectionCB(client)
	}
}

// SetOnNewClientConnectionCB sets the callback for new client connections
func (t *TCPConnector) SetOnNewClientConnectionCB(f func(c Client)) {
	t.onNewClientConnectionCB = f
}

// SetOnClientDisconnectCB sets the callback for client disconnections
func (t *TCPConnector) SetOnClientDisconnectCB(f func(c Client, err error)) {
	t.onClientDisconnectCB = f
}
