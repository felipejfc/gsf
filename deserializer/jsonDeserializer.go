package deserializer

import "encoding/json"

// JSONDeserializer struct
type JSONDeserializer struct{}

// NewJSONDeserializer ctor
func NewJSONDeserializer() *JSONDeserializer {
	return &JSONDeserializer{}
}

// Deserialize method
func (j *JSONDeserializer) Deserialize(messageBytes []byte) (*Message, error) {
	message := &Message{}
	err := json.Unmarshal(messageBytes, message)
	if err != nil {
		return nil, err
	}
	return message, nil
}
