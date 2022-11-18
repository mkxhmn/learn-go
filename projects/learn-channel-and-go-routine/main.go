package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://google.com",
		"https://stackoverflow.com",
		"https://golang.com",
	}

	for _, link := range links {
		// remember we only use 'go' keyword in front of function call.
		// this is to invoke concurrency
		go sequentialCheckLink(link)
	}
}

// sequential approach
func sequentialCheckLink(link string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}

	fmt.Println(link, "is up!")
}
