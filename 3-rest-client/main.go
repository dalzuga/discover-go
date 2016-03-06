package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {

	timenow := time.Now()

	respvar, err := http.Get("http://www.omdbapi.com/?i=tt0372784&plot=short&r=json")

	if err != nil {
		fmt.Println("error:", err)
	}

	respvarjson, err := json.Marshal(respvar)

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Printf("json output: %v\ntime: %v\n", respvarjson, timenow)
}
