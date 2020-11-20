package main

import (
	"bufio"
	"log"
	"os"
)

func GetQuotes(filePath string, quotes []string) ([]string, error) {
	quotesFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer quotesFile.Close()

	scanner := bufio.NewScanner(quotesFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		quotes = append(quotes, scanner.Text())
	}
	return quotes, nil
}
