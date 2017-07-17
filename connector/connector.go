package connector

// Client interface
type Client interface {
	Send(data []byte) error
	Read() ([]byte, error)
	Close()
}

// Connector Iface
type Connector interface {
	Listen(host string, port int) error
	SetOnNewClientConnectionCB(func(c Client))
	SetOnClientDisconnectCB(func(c Client, err error))
}
