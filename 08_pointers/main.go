package main

import "fmt"

func main(){
	var p *int32 = new(int32) // make sure u either do this or assign your pointer to something, it must not NIL
	var i int32 = 4

	fmt.Printf("The value p points to is: %v", *p) // dereferencing the pointer
	fmt.Printf("\nThe value if i is: %v", i)

	p = &i
	*p = 1 // The value of i also changes

	fmt.Printf("\nThe value p points to is: %v", *p)
	fmt.Printf("\nThe value if i is: %v", i)

	var k int32 = 2
	i = k

	fmt.Printf("\nThe value p points to is: %v", *p)
	fmt.Printf("\nThe value if i is: %v\n", i)

	//We can pass arrays with pointers in a function so a copy isn't made
	var thing1 = [5]float64{1, 2, 3, 4, 5}
	var result [5]float64 = square(&thing1)
	fmt.Println(result)
}

func square(thing2 *[5]float64) [5]float64{
	for i := range thing2{
		thing2[i] = thing2[i]*thing2[i]
	}
	return *thing2
}