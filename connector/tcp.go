package connector

import (
	"bufio"
	"fmt"
	"net"
)

// TCPConnector struct
type TCPConnector struct {
}

// TCPClient interface
type TCPClient struct {
	conn     net.Conn
	server   *TCPConnector
	socketID string
}

func (c *TCPClient) handleMessages() {
	reader := bufio.NewReader(c.conn)
	for {
		message, err := reader.ReadBytes('\n')
		if err != nil {
			c.conn.Close()
			c.server.onClientDisconnect(c)
			return
		}
		c.server.onMessage(c, message)
	}
}

// NewTCPConnector ctor
func NewTCPConnector() *TCPConnector {
	return &TCPConnector{}
}

func (t *TCPConnector) onClientDisconnect(c *TCPClient) {
	fmt.Printf("connection closed: %s\n", c.socketID)
}

func (t *TCPConnector) onMessage(c *TCPClient, m []byte) {
	fmt.Printf("new message from client %s: %s\n", c.socketID, string(m))
}

func (t *TCPConnector) onNewClient(c *TCPClient) {
	fmt.Printf("new client connection: %s\n", c.socketID)
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
		client := &TCPClient{
			conn:     conn,
			socketID: uuid.NewV4(),
			server:   t,
		}
		go client.handleMessages()
		t.onNewClient(client)
	}
}
