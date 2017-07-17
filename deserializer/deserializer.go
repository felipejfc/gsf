package deserializer

// Message struct
type Message struct {
	Route string      `json:"route"`
	Data  interface{} `json:"data"`
}

// Deserializer interface
type Deserializer interface {
	Deserialize(message []byte) (*Message, error)
}
