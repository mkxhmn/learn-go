package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := logWriter{}

	if _, err := io.Copy(lw, resp.Body); err != nil {
		log.Fatal(err)
	}
}

func (logWriter) Write(bs []byte) (int, error) {
	n := len(bs)

	fmt.Println("just wrote this many bytes: ", n)
	return n, nil
}
