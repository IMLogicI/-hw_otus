package main

import (
	"log"

	"golang.org/x/example/stringutil"
)

const hello = "Hello, OTUS!"

func main() {
	log.Println(stringutil.Reverse(hello))
}
