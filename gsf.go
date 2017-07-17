package gsf

import (
	"fmt"

	"github.com/felipejfc/gsf/connector"
	"github.com/felipejfc/gsf/deserializer"
)

// Server struct
type Server struct {
	connector    connector.Connector
	deserializer deserializer.Deserializer
	host         string
	port         int
}

// NewServer ctor
func NewServer(host string, port int) *Server {
	s := &Server{
		connector:    connector.NewTCPConnector(),
		deserializer: deserializer.NewJSONDeserializer(),
		host:         host,
		port:         port,
	}
	s.initialize()
	return s
}

func (s *Server) initialize() {
	clientHandler := NewClientHandler(s)
	s.connector.SetOnNewClientConnectionCB(clientHandler.onNewClientConnection)
	s.connector.SetOnClientDisconnectCB(clientHandler.onClientDisconnect)
}

// SetDeserializer setter
func (s *Server) SetDeserializer(d deserializer.Deserializer) {
	s.deserializer = d
}

// SetConnector setter
func (s *Server) SetConnector(c connector.Connector) {
	s.connector = c
}

// Start starts the server
func (s *Server) Start() error {
	fmt.Printf("listening at %s:%d\n", s.host, s.port)
	err := s.connector.Listen(s.host, s.port)
	if err != nil {
		return err
	}
	return nil
}
