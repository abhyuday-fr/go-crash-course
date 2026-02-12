package main

import (
	"fmt"
	"strings"
)

func main(){
	var myString = "this is a string"
	var indexed = myString[0]
	fmt.Printf("%v, %T\n", indexed, indexed)
	for i,v := range myString{
		fmt.Println(i, v)
	}
	fmt.Println(len(myString))

	//Concatenate
	myString += ", indeed"
	fmt.Println(myString)

	// Strings are immutable

	//Faster way of concatenating: import "strings"
	var strSlice = []string{"s", "u", "b", "s", "c", "r", "i", "b", "e"}
	var strBuilder strings.Builder
	for i := range strSlice{
		strBuilder.WriteString(strSlice[i])
	}
	fmt.Printf("\n%v", myString)
}