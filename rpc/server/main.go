package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

// Item data being stored in database
type Item struct {
	Title string
	Body  string
}

// API for create methods
type API int

var database []Item

// GetDB get the whole database
func (a *API) GetDB(_ string, reply *[]Item) error {
	*reply = database
	log.Println("GetDB")
	return nil
}

// GetByName find  item based on title
func (a *API) GetByName(title string, reply *Item) error {
	var getItem Item
	for _, val := range database {
		if val.Title == title {
			getItem = val
		}
	}
	*reply = getItem
	if getItem == (Item{}) {
		return errors.New("No item found")
	}
	return nil
}

// CreateItem create Item
func (a *API) CreateItem(item Item, reply *Item) error {
	if item == (Item{}) {
		return errors.New("Empty Item")
	}
	database = append(database, item)
	*reply = item
	log.Println("CreateItem")
	return nil
}

// EditItem edit Item
func (a *API) EditItem(edit Item, reply *Item) error {
	var changedItem Item
	for idx, val := range database {
		if val.Title == edit.Title {
			database[idx] = edit
			changedItem = edit
		}
	}
	*reply = changedItem
	if changedItem == (Item{}) {
		return errors.New("No item found to edit")
	}
	log.Println("EditItem")
	return nil
}

// DeleteItem delete Item
func (a *API) DeleteItem(item Item, reply *Item) error {
	var delItem Item
	for idx, val := range database {
		if val.Title == item.Title && val.Body == item.Body {
			database = append(database[:idx], database[idx+1:]...)
			delItem = item
		}
	}
	*reply = delItem
	if delItem == (Item{}) {
		return errors.New("No item found to be delete")
	}
	log.Println("DeleteItem")
	return nil
}

func main() {
	api := new(API)
	err := rpc.Register(api)
	if err != nil {
		panic(err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	log.Printf("Serving rpc on port %d", 4040)

	err = http.Serve(listener, nil)
	if err != nil {
		panic(err)
	}

}
