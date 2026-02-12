package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// complex64 complex128

	// Using var
	var name string = "abby"
	var age int = 37
	const isCool = true

	name2 := "lawless" //Can't do this type of assignment outside a fnctn (no need to declare data type)
	num, mail := 17, "@gmail.com"

	fmt.Println(name, age, isCool, name2, num, mail)
	fmt.Printf("%T\n", name2) //prints data type
}