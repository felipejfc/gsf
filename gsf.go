package gsf

import (
	"fmt"

	"github.com/felipejfc/gsf/connector"
)

// Server struct
type Server struct {
	connector connector.Connector
	host      string
	port      int
}

// NewServer ctor
func NewServer(c connector.Connector, host string, port int) *Server {
	s := &Server{
		connector: c,
		host:      host,
		port:      port,
	}
	return s
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
