package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(100)

	if randomNumber > 50 {
		fmt.Println("my random number is", randomNumber, "and is greater than 50")
	} else {
		fmt.Printf("my random number is %d and is less than 50\n", randomNumber)
	}

	school := "Holberton School"

	if school == "Holberton School" {
		fmt.Println("I am a student of the", school)
	} else {
		fmt.Println("I am NOT a student of the", school)
	}

	var beautifulWeather bool // how do i set var type explicitly?

	beautifulWeather = true

	if beautifulWeather {
		fmt.Println("yay weather!")
	} else {
		fmt.Printf("nay weather :(\n")
	}

	holbertonFounders := [3]string{}

	holbertonFounders[0] = "Rudy Rigot"
	holbertonFounders[1] = "Sylvain Kalache"
	holbertonFounders[2] = "Julien Barbier"

	//	fmt.Printf("%v is a founder\n", holbertonFounders)

	for i := 0; i < 3; i++ {
		fmt.Printf("%s is a founder\n", holbertonFounders[i])
	}
}
