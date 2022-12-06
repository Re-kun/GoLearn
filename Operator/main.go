package main

import "fmt"

func main() {

	var a int = 7
	var b int = 2

	// Arimatika

	fmt.Printf("\n%v + %v : %v\n", a, b, a+b)
	fmt.Printf("%v - %v : %v\n", a, b, a-b)
	fmt.Printf("%v * %v : %v\n", a, b, a*b)
	fmt.Printf("%v / %v : %v\n", a, b, a/b)
	fmt.Printf("%v %% %v : %v\n", a, b, a%b)

	// Comparation

	fmt.Printf("\n%v == %v : %v\n", a, a, a == a)
	fmt.Printf("%v != %v : %v\n", a, b, a != b)
	fmt.Printf("%v < %v : %v\n", a, b, a < b)
	fmt.Printf("%v <= %v : %v\n", a, b, a <= b)
	fmt.Printf("%v > %v : %v\n", a, b, a > b)
	fmt.Printf("%v >= %v : %v\n", a, b, a >= b)

	// Logic Operator
	var isTrue bool = true
	var isFalse bool = false

	// And
	fmt.Println("\n----[AND]------")
	fmt.Printf("\n%v && %v: %v", isTrue, isFalse, isTrue && isFalse)
	fmt.Printf("\n%v && %v: %v", isFalse, isFalse, isFalse && isFalse)
	fmt.Printf("\n%v && %v: %v\n", isFalse, isTrue, isFalse && isTrue)

	// Or
	fmt.Println("\n----[OR]------")
	fmt.Printf("\n%v || %v: %v", isTrue, isFalse, isTrue || isFalse)
	fmt.Printf("\n%v || %v: %v", isFalse, isFalse, isFalse || isFalse)
	fmt.Printf("\n%v || %v: %v", isFalse, isTrue, isFalse || isTrue)
	fmt.Printf("\n%v || %v: %v\n", isTrue, isTrue, isTrue || isTrue)

	// Reverse
	fmt.Println("\n----[REVERSE]------")
	var reverseStatus = !isFalse
	fmt.Printf("\n!isFalse : %v", reverseStatus)
}
