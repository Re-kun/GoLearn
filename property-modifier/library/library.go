package library

import "fmt"

// Like constructor function

func init() {
	fmt.Println("HI THIS IS FORM INIT FUNCTION")
}

func SayHello() {
	fmt.Println("hello")
}
func Introduce(name string) {
	fmt.Println("My name is", name)
}
