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
		fmt.Println("error1:", err)
	}

	fmt.Println("status code is", err, http.StatusOK)

	var respvarStruct JsonOmdbStruct

	if err := json.NewDecoder(respvar.Body).Decode(&respvarStruct); err != nil {
		fmt.Println("error2:", err)
	}

	// respvarMarshaled, err := json.Marshal(respvarStruct)
	//
	// if err != nil {
	// 	fmt.Println("can't marshal:", err)
	// }

	fmt.Printf("struct output: %+v\ntime: %v\n", respvarStruct, timenow)
	// fmt.Printf("json output: %+v\ntime: %v\n", respvarMarshaled, timenow)
}
