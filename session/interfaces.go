package session

// NetworkConn interface
type NetworkConn interface {
	ID() int64
	Send([]byte) error
	Push(session *Session, route string, v interface{}) error
	Response(session *Session, v interface{}) error
	Close()
}
