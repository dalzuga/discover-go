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
	defer respvar.Body.Close()

	if err != nil {
		fmt.Println("error1:", err)
	}

	fmt.Println("status code is", respvar.Status)

	var respvarStruct JsonOmdbStruct

	if err := json.NewDecoder(respvar.Body).Decode(&respvarStruct); err != nil {
		fmt.Println("error2:", err)
	}

	// respvarMarshaled, err := json.Marshal(respvarStruct)
	//
	// if err != nil {
	// 	fmt.Println("can't marshal:", err)
	// }

	// respvarMarshaled, err := json.Marshal(&respvarStruct)
	//
	// if err != nil {
	// 	fmt.Println("error3:", err)
	// }

	respvarMarshaledIndent, err := json.MarshalIndent(&respvarStruct, "", "  ")

	if err != nil {
		fmt.Println("error3:", err)
	}

	//fmt.Printf("respvar: %v\n", respvar.Body)
	//fmt.Printf("jsonMarshaled:\n%s\n", respvarMarshaled)
	fmt.Printf("jsonMarshaledIndent:\n%s\n", respvarMarshaledIndent)
	//fmt.Printf("struct output: %+v\ntime: %v\n", respvarStruct, timenow)
	//fmt.Printf("respvarStruct:\n%#v\n", respvarStruct)
	fmt.Printf("time: %v\n", timenow)
	//fmt.Printf("json output: %s\ntime: %v\n", respvarMarshaled, timenow)
}
