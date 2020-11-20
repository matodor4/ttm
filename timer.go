package main

import (
	"log"
	"time"
)

var quotes []string

func StartTimer(syncCh chan struct{}, quoteCh chan []byte) {
	quoteFile, err := GetQuotes(QuotesPath, quotes)
	if err != nil {
		log.Fatal(err)
	}
	go Timer(quoteFile, syncCh, quoteCh)
}

func Timer(quotes []string, syncCh chan struct{}, quoteCh chan []byte) {
	maxCount := len(quotes) - 1
	count := 0

	//tick channel
	tick := time.Tick(time.Second)

	for  {
		select {
		case <- tick:
			count++
			if count > maxCount {
				count = 0
			}
		case <- syncCh:
			actualQuote := quotes[count]
			quoteCh <- []byte(actualQuote)
		}
	}
}
