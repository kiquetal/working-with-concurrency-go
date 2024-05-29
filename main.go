package main

import (
	"fmt"
	"sync"
)

var msg string

func updateMessage(s string, wg * sync.WaitGroup) {

	msg = s
	wg.Done()

}
func printMessage() {
	fmt.Println(msg)

}

var wg sync.WaitGroup
func main() {

	msg = "Hello world!"

	wg.Add(1)
	go updateMessage("Hello, universe!",&wg)
	wg.Wait()
	printMessage()
	wg.Add(1)
	go updateMessage("Hello, cosmos!",&wg)
	wg.Wait()
	printMessage()
	wg.Add(1)
	go updateMessage("Hello, world!",&wg)
	wg.Wait()
	printMessage()
	
	



}
