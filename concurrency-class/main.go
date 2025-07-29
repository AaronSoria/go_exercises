package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	go greet("hola", done)
	go greet("tarolas", done)
	go slowGreet("estoy pensando", done)
	go greet("tengo hambre", done)

	for range done {
	}
}

func greet(phrase string, doneChannel chan bool) {
	fmt.Println("Hello you said: ", phrase)
	doneChannel <- true
}

func slowGreet(phrase string, doneChannel chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello you said: ", phrase)
	doneChannel <- true
	close(doneChannel)
}
