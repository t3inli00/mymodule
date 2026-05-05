package mymodule

import "fmt"

func SayHello() {
	fmt.Println("Hello my Module")
}

func SayBye() {
	fmt.Println("Bye my Module")
}

func main() {
	SayHello()
	SayBye()
}
