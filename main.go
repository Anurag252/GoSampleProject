package main

import (
	"fmt"
	"net/http"

	endpoint "ReferenceProject.com/endpoints"
)

//Client struct encapsulated httpclient
type client struct {
}

func (client *client) HandleFunc(add string, h endpoint.HandlerFunc) {
	fmt.Println("in func")
	http.HandleFunc(add, h)
}

func (client *client) ListenAndServe(add string, h http.Handler) error {
	fmt.Println("in func")
	err := http.ListenAndServe(add, h)
	return err
}

func main() {

	clientTest := &client{}
	if &clientTest == nil {
		fmt.Print("error occured")
	}
	err := endpoint.Register(clientTest)
	fmt.Println("Anuragsdf")
	if err != nil {
		fmt.Print("error occured")
	}

}
