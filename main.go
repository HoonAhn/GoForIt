package main

import (
	"errors"
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

var errRequestFailed = errors.New("Request Failed!")

func main() {
	ch := make(chan requestResult)
	results := map[string]string{}
	urls := []string{
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.airbnb.com/",
		"https://www.facebook.com/",
		"https://www.reddit.com/",
		"https://www.instagram.com/",
		"https://soundcloud.com/",
	}
	for _, url := range urls {
		go hitURL(url, ch)
	}
	for i := 0; i < len(urls); i++ {
		result := <-ch
		results[result.url] = result.status
	}
	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitURL(url string, ch chan<- requestResult) {
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FALILED"
	}
	ch <- requestResult{url: url, status: status}
}
