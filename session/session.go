package session

import (
	"time"

	"github.com/felipejfc/gsf/connector"
	uuid "github.com/satori/go.uuid"
)

// Session struct
type Session struct {
	ID            string
	UID           string
	ClientConn    connector.Client
	ReqID         int64
	data          map[string]interface{}
	lastHeartbeat int64
}

// NewSession ctor
func NewSession(clientConn connector.Client) *Session {
	return &Session{
		ID:            uuid.NewV4().String(),
		ClientConn:    clientConn,
		data:          make(map[string]interface{}),
		lastHeartbeat: time.Now().Unix(),
		ReqID:         0,
	}
}
