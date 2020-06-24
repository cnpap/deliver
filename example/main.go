package main

import (
	"os"

	"github.com/justlikesuolong/deliver"
)

func main() {
	d := &deliver.Deliver{
		Addr: os.Args[1],
	}
	err := d.Listen()
	if err != nil {
		panic(err)
	}
}