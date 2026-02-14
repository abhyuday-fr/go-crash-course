// At high lvl think of channels as a way yo enable go routines to pass around information
/*
Main Features of a channel:
1. Hold data
2. Thread Safe
3. Listen for data
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5

func main(){
	var chickenChannel = make(chan string)
	var websites = []string{"walmart.com", "costco.com", "wholefoods.com"}
	for i:= range websites{
		go checkChickenPrices(websites[i], chickenChannel)
	}
	sendMessage(chickenChannel)
}

func checkChickenPrices(website string, chickenChannel chan string){
	for{
		time.Sleep(time.Second*1)
		var chickenPrice = rand.Float32()*20
		if chickenPrice <= MAX_CHICKEN_PRICE{
			chickenChannel <- website
			break
		}
	}
}

func sendMessage(chickenChannel chan string){
	fmt.Printf("\nFound a deal on chicken at %s", <-chickenChannel)
}

/*
func main(){
	var c = make(chan int)  // way to make a channel that holds an integer value, think of it like c : []

	/* 
	NOT the way channels are meant to be used but you can still read this for understanding why

	c <- 1 // adding value '1' to the channel (c : [1])
	var i = <- c // the value in c gets popped out and now the variable i holds that value (c : [], i = 1)
	fmt.Println(i) // This will give a deadlock error because when we write to an unbuffered channel the code will block until smthg else reads from it
	*/
	/*

	//var c = make(chan int, 5) // Buffered channel:- now it can hold multiple values, here 5.

	go process(c)
	for i:= range c{
		fmt.Println(i)
	}

}

func process(c chan int){
	defer close(c) // defer means:- do this stuff right before the fnctn ends
	for i:=0; i<5; i++{
		c <- i
	}
}
*/