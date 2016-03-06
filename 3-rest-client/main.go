package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	timenow := time.Now()

	respvar, err := http.Get("http://www.omdbapi.com/?i=tt0372784&plot=short&r=json")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("output: %v\ntime: %v\n", respvar, timenow)
}
