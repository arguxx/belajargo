package main	
import "fmt"

func sayHello(name string) {
	fmt.Println("Hello, World!")
	fmt.Println("Hello, ", name)
}

func main() {
	sayHello("argu")
}