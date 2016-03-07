package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {

	u, err := url.Parse("http://www.omdbapi.com/?i=tt0372784&plot=short&r=json")
	if err != nil {
		log.Fatal(err)
	}

	u.Scheme = "http"
	u.Host = "omdbapi.com"
	q := u.Query()
	q.Set("i", "tt3682448")
	q.Set("tomatoes", "true")
	u.RawQuery = q.Encode()

	address := u.String()

	fmt.Printf("%s\n", address)

	var respvarStruct JsonOmdbStruct

	respvar, err := http.Get(address)
	defer respvar.Body.Close()

	if err != nil {
		fmt.Println("error1:", err)
	}

	if err := json.NewDecoder(respvar.Body).Decode(&respvarStruct); err != nil {
		fmt.Println("error2:", err)
	}

	fmt.Println("status code is", respvar.Status)

	fmt.Println("The movie ID is", respvarStruct.ImdbID)
	fmt.Println("The parsed URL is", u)
	fmt.Printf("The movie name is '%s' and the actors are %s.\n", respvarStruct.Title, respvarStruct.Actors)
	fmt.Printf("The IMBD rating is %s with %s votes.\n", respvarStruct.ImdbRating, respvarStruct.ImdbVotes)
	fmt.Printf("The tomato rating is %s\n", respvarStruct.TomatoRating)
}
