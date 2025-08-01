//go:build !solution

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func checkErr(err error) {
	if err != nil {
		msg := "fetch:"
		log.SetPrefix("")
		log.SetFlags(0)
		log.Fatalf("%s %v", msg, err)
	}
}

func checkUrl(url string) {
	response, err := http.Get(url)
	checkErr(err)
	body, err := io.ReadAll(response.Body)
	response.Body.Close()
	checkErr(err)
	fmt.Printf("%s", body)
}

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}
	for _, url := range os.Args[1:] {
		checkUrl(url)
	}
}
