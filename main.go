package main

import "github.com/go-sample-app/app/presentation"

import "fmt"

func main() {
	e := presentation.Router()
	fmt.Println(e)
}
