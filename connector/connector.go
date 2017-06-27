package connector

// Connector Iface
type Connector interface {
	Listen(host string, port int) error
}
