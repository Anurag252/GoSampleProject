package endpoint

import (
	"fmt"
	"net/http"
)

//HandlerFunc type
type HandlerFunc func(w http.ResponseWriter, r *http.Request)

//IClient interface
type IClient interface {
	HandleFunc(string, HandlerFunc)
	ListenAndServe(string, http.Handler) error
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func compileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

//Register all http endpoint
func Register(client IClient) error {
	fmt.Println("Anurag")
	client.HandleFunc("/", testHandler)
	client.HandleFunc("/Test", testHandler)
	client.HandleFunc("/Home", homeHandler)
	client.HandleFunc("/Index", indexHandler)
	client.HandleFunc("/Compile", compileHandler)
	err := client.ListenAndServe(":8080", nil)
	if err != nil {
		return err
	}
	return nil
}
