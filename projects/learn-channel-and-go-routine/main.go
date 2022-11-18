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

	c := make(chan string)

	for _, link := range links {
		// remember we only use 'go' keyword in front of function call.
		// this is to invoke concurrency
		go sequentialCheckLink(link, c)
	}

	fmt.Println(<-c)
}

// sequential approach
func sequentialCheckLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "might be down I think"
		return
	}

	fmt.Println(link, "is up!")
	c <- "yup it is up"
}
