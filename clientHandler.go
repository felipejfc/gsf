package gsf

import (
	"fmt"

	"github.com/felipejfc/gsf/connector"
	"github.com/felipejfc/gsf/session"
)

// ClientHandler struct
type ClientHandler struct {
	// TODO use goroutine safe map
	sessionMap map[string]session.Session
	server     *Server
}

// NewClientHandler ctor
func NewClientHandler(s *Server) *ClientHandler {
	return &ClientHandler{
		sessionMap: map[string]session.Session{},
		server:     s,
	}
}

func (h *ClientHandler) onNewClientConnection(c connector.Client) {
	s := session.NewSession(c)
	fmt.Printf("new session: %s\n", s.ID)
	go h.handleClientMessages(s)
}

func (h *ClientHandler) onClientDisconnect(c connector.Client, e error) {
	fmt.Printf("session closed: %s\n", e.Error())
}

func (h *ClientHandler) handleClientMessages(s *session.Session) {
	// see type if req increment nbr last req
	// deserialize
	// route
	defer s.ClientConn.Close()
	for {
		data, err := s.ClientConn.Read()
		if err != nil {
			break
		}
		fmt.Printf("data rcv sid(%s): %s", s.ID, string(data))
		msg, err := h.server.deserializer.Deserialize(data)
		if err != nil {
			fmt.Printf("error deserializing message: %d\n", err.Error())
		} else {
			fmt.Printf("message (route: %s, data: %s)\n", msg.Route, msg.Data)
		}
		// TODO route messages
	}
}
