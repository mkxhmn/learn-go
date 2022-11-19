package main

import (
	"fmt"
	"net/http"
	"time"
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

	// remember, fmt.Println(<-c) will block the main thread until it receives a value from the channel
	// what happen if we print the second value from channel?
	// surprisingly, it will print the second value from channel
	// fmt.Println(<-c)

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	for l := range c {
		time.Sleep(5 * time.Second)
		go sequentialCheckLink(l, c)
	}
}

// sequential approach
func sequentialCheckLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- "yup it is up"
}
