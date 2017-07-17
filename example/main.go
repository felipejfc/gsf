package main

import (
	"github.com/felipejfc/gsf"
)

func main() {
	server := gsf.NewServer("0.0.0.0", 6868)
	err := server.Start()
	if err != nil {
		panic(err)
	}
}
