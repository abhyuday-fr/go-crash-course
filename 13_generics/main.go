package main

import "fmt"

type gasEngine struct{
	gallons float32
	mpg float32
}

type electricEngine struct{
	kwh float32
	mpkwh float32
}

type car [T gasEngine | electricEngine]struct{
	carMake string
	carModel string
	engine T
}

func main(){
	var gasCar = car[gasEngine]{
		carMake: "Honda",
		carModel: "Civic",
		engine: gasEngine{
			gallons: 12.4,
			mpg: 40,
		},
	}
	var electricCar = car[electricEngine]{
		carMake: "Tesla",
		carModel: "Mpdel 3",
		engine: electricEngine{
			kwh: 12.4,
			mpkwh: 40,
		},
	}

	fmt.Println(gasCar)
	fmt.Println(electricCar)
}