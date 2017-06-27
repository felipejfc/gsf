package main

import (
	"github.com/felipejfc/gsf"
	"github.com/felipejfc/gsf/connector"
)

func main() {
	connector := connector.NewTCPConnector()
	server := gsf.NewServer(connector, "0.0.0.0", 6868)
	err := server.Start()
	if err != nil {
		panic(err)
	}
}
