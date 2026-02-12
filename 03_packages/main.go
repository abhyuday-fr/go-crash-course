package main

import (
	"fmt"
	"math"
	"github.com/abhyuday-fr/go_crash_course/go_crash_course/03_packages/strutil"
)

func main(){
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(2.7))
	fmt.Println(strutil.Reverse("olleh"))
}