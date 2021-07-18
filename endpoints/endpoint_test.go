package endpoint

import (
	"errors"
	"net/http"
	"testing"
)

type MockedClient struct {
}

func (client *MockedClient) HandleFunc(add string, h HandlerFunc) {
}

func (client MockedClient) ListenAndServe(add string, h http.Handler) error {
	return errors.New("error raised")
}

func TestHandlerHome(t *testing.T) {
	mockedClient := MockedClient{}
	err := Register(&mockedClient)
	if err != nil {
		t.Errorf("got error expected nil")
	}

}
