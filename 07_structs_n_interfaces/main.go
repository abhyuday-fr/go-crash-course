package main

import "fmt"

type gasEngine struct{
	mpg uint8
	gallons uint8
	ownerinfo owner
}

type owner struct{
	name string
}

func (e gasEngine) milesLeft() uint8{ // works like function in class
	return e.gallons * e.mpg
}

type engine interface{ // more general way to define a contract of methods that a type must implement
	milesLeft() uint8
}

func canMakeit(e engine, miles uint8){
	if miles<=e.milesLeft(){
		fmt.Println("You can make it")
	}else{
		fmt.Println("Need to fuel up first")
	}
}

func main(){
	var myEngine gasEngine = gasEngine{25, 15, owner{"abby"}}
	myEngine.mpg = 20
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.ownerinfo.name)

	fmt.Printf("Total miles left in the Tank: %v", myEngine.milesLeft())

	canMakeit(myEngine, 50)
}