package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T)  {

	syncCh := make(chan struct{})
	quoteCh := make(chan []byte)

	StartTimer(syncCh, quoteCh)

	serv := httptest.NewServer(getHandler(syncCh, quoteCh))
	defer serv.Close()
	resp, err := http.Get(serv.URL)

	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Error("Status not OK")
	}
	defer resp.Body.Close()

}
