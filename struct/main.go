package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	var magi person
	magi.firstName="Magizhalan"
	magi.lastName="Vijay"
	fmt.Println(magi)
}