package main

import (
	"errors"
	"fmt"
)

func greeting(name string) string {
	return "hiiiiii " + name
}

func main(){
	var numerator = 11
	var denominator = 0
	fmt.Println(greeting(":3"))
	var result, remainder, err = intDivision(numerator, denominator)

	if(err!=nil){
		fmt.Println(err.Error())
	}else{
		fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
	}
}

func intDivision(numerator int, denominator int) (int, int, error){
	var err error
	if denominator==0{
		err = errors.New("Cannot divide by 0")
		return 0, 0, err
	}

	var result int = numerator/denominator
	var remainder int = numerator%denominator
	return result, remainder, err
}