package main
import "fmt"

func main() {
    var name string
    var age int
    fmt.Print("Enter your name and age: ")
    fmt.Scanf("%s %d", &name, &age) // You will have to input in the same line
    fmt.Printf("%s is %d years old.\n", name, age)
}   