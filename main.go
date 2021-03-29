package main

import (
	"errors"
	"fmt"
	"net/http"
)

type requetResult struct {
	url    string
	status string
}

var errRequestedFailed = errors.New("Request Failed")

func main() {

	results := make(map[string]string)

	c := make(chan requetResult)

	urls := []string{
		"https://www.google.com/",
		"https://www.airbnb.co.kr/",
	}

	for _, url := range urls {
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitURL(url string, c chan<- requetResult) {
	resp, err := http.Get(url)

	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}

	c <- requetResult{url: url, status: status}
}
