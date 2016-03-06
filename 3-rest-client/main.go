package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {

	//timenow := time.Now()

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

	u, err := url.Parse("http://www.omdbapi.com/?i=tt0372784&plot=short&r=json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The movie is '%s' and the actors are %s.\n", respvarStruct.Title, respvarStruct.Actors)

	u.Scheme = "http"
	u.Host = "omdbapi.com"
	q := u.Query()
	q.Set("i", "tt3682448")
	u.RawQuery = q.Encode()

	address := u.String()

	fmt.Printf("%s\n", address)

	respvar, err = http.Get(address)
	defer respvar.Body.Close()

	if err != nil {
		fmt.Println("error1:", err)
	}

	fmt.Println("status code is", respvar.Status)

	if err = json.NewDecoder(respvar.Body).Decode(&respvarStruct); err != nil {
		fmt.Println("error2:", err)
	}

	fmt.Println("Parsed URL", u)
	fmt.Printf("The movie is '%s' and the actors are %s.\n", respvarStruct.Title, respvarStruct.Actors)
}
