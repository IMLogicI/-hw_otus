package main

import (
	"golang.org/x/example/stringutil"
	"log"
)

const hello = "Hello, OTUS!"

func main() {
	log.Println(stringutil.Reverse(hello))
}
