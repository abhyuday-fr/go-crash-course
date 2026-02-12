package main

import (
	"fmt"
	"time"
)

func main() {
	var Arr [3]int32
	fmt.Println(Arr[0])
	fmt.Println(Arr[1:3])

	fmt.Println(&Arr[0])
	fmt.Println(&Arr[1])
	fmt.Println(&Arr[2])

	arr2 := [...]int32{1,2,3,4}
	fmt.Println(arr2)

	var Slice1[]int32 = []int32{8, 9}
	Slice1 = append(Slice1, 7)
	fmt.Println(len(Slice1), cap(Slice1))

	var Slice2 []int32 = make([]int32, 3, 10)
	fmt.Println(Slice2)

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"adam":23, "sarah":45}

	fmt.Println(myMap2["adam"])

	var age, ok = myMap2["jason"] // value of 'ok' is true if the key exists, false otherwise
	if ok{
		fmt.Printf("age is %v", age)
	}else{
		fmt.Println("Invalid name")
	}

	delete(myMap2, "adam")

	// iterating through map
	for name, age := range myMap2{
		fmt.Printf("Name: %v, Age: %v \n", name, age)
	}

	//iterating through array
	for i, v := range arr2{
		fmt.Printf("Index: %v, Value: %v \n", i, v)
	}


	//-------Benefit of Preallocation--------------
	var n int = 1000000
	var testSlice = []int{}
	var testSlice2 = make([]int, 0, n)
	fmt.Printf("Total time without preallocation: %v\n", timeLoop(testSlice, n))
	fmt.Printf("Total time with preallocation: %v", timeLoop(testSlice2, n))
}

func timeLoop(slice[]int, n int) time.Duration{
	var t0 = time.Now()
	for len(slice)<n{
		slice = append(slice, 1)
	}
	return time.Since(t0)
}