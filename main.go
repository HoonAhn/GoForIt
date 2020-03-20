package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request Failed!")

func main() {
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
		resultStatus := "OK"
		err := hitURL(url)
		if err != nil {
			resultStatus = "FAILED"
		}
		results[url] = resultStatus
	}
	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string) error {
	fmt.Println("Checking: ", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err, resp.StatusCode)
		return errRequestFailed
	}
	fmt.Println(err)
	return nil
}
