package main

import (
	"fmt"
	"time"
)

type user struct {
	Name string `json:"name"`
	DOB  string `json:"date_of_birth"`
	City string `json:"city"`
}

func (d *user) printhello() {
	//	d.Name = "dsfs"
	fmt.Printf("Hello %s\n", d.Name)
}

func (d *user) printitall() {
	//	d.Name = "dsfs"

	fmt.Printf("%s who was born in %s would be %s\n", d.Name, d.City, d.DOB)
}

func main() {

	// var u user
	// u.Name = "Betty Holberton"
	// u.DOB = "March 7, 1917"
	// u.City = "Philadelphia"

	u := user{Name: "Betty Holberton", DOB: "March 7, 1917", City: "Philadelphia"}

	value := u.DOB
	layout := "01 07, 06"

	DOB, _ := time.Parse(layout, value)
	fmt.Printf("%s\n", u.DOB)
	fmt.Printf("%s\n", DOB)

	u.printhello()
	//	fmt.Println(u.Name)
	u.printitall()

	// t := "Thu, 05/19/2011, 10:47PM"
	//
	// layout := "Mon, 01/02/06, 03:04PM"
	// var t time.Time
	// timenow.Year() =
}
