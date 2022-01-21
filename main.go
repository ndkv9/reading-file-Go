package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	file, err := ioutil.ReadFile("./test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(file))
}
