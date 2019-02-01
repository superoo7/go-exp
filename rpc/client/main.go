package main

import (
	"log"
	"net/rpc"
)

// Item data being stored in database
type Item struct {
	Title string
	Body  string
}

func main() {
	var reply Item
	var replyDB []Item

	client, err := rpc.DialHTTP("tcp", "localhost:4040")
	if err != nil {
		panic(err)
	}

	a := Item{"First", "First thing"}

	client.Call("API.CreateItem", a, &reply)

	log.Printf("%s", reply)

	client.Call("API.GetDB", "", &replyDB)

	log.Printf("%s", replyDB)

}
