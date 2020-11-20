package main

const (
	JsonPath = "port.json"
	QuotesPath = "quotes.txt"
)

func main() {
	syncCh := make(chan struct{})
	quoteCh := make(chan []byte)

	StartTimer(syncCh, quoteCh)
	StartServer(syncCh, quoteCh)
}