package main

import "fmt"

func main(){
	var num1 float32
	var num2 float32
	var symbol rune
	var res float32
	fmt.Println("Write your operation for 2 numbers (eg. 2.5 + 3.3)")
	fmt.Println("Available Operations:\n1. Addition (+)\n2. Subtraction (-)\n3. Multiplication (* OR X)\n4. Division")
	fmt.Scanf("%f %c %f", &num1, &symbol ,&num2)
	switch symbol {
	case '+':
		res = num1 + num2
	case '-':
		res = num1 - num2
	case '*':
		res = num1 * num2
	case 'X':
		res = num1 * num2
	case '/':
		res = num1 / num2
	default:
		fmt.Println("Enter valid operation")

	}
	fmt.Printf("Your answer is: %f", res)
}