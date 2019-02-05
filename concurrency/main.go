package main

import (
	"fmt"
	"os"

	"github.com/superoo7/go-exp/concurrency/channel"
	"github.com/superoo7/go-exp/concurrency/mutex"
)

func main() {
	if len(os.Args) != 2 {
		panic("please state type of val to execute. (Format: go run concurrency/main.go mutex)")
	}
	val := os.Args[1]
	switch val {
	case "mutex":
		mutex.Mutex()
	case "channel":
		channel.Channel()
	default:
		fmt.Print("not chosen")
	}
}
