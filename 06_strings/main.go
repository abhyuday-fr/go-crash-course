package main

import "fmt"

func main(){
	var myString = "this is a string"
	var indexed = myString[0]
	fmt.Printf("%v, %T\n", indexed, indexed)
	for i,v := range myString{
		fmt.Println(i, v)
	}
}