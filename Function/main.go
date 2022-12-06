package main

import "fmt"

// Simple function

func helloWorld() {
	fmt.Println("Hello World")
}

// Parameters

func sayHello(user string) string {
	var text = "Hello " + user
	return text
}

// used return to stoped proses

func getThree() {
	for i := 0; i < 7; i++ {
		if i == 3 {
			fmt.Println("Hello ini tiga\n")
			return
		}
	}
	fmt.Println("Tidak akan di eksekusi")
}

// Multiple return

var waifus = map[int]string{
	1: "Rem",
	2: "chitoge",
	3: "Onodera",
}

func getWaifu(id int) (int, string) {
	var validWaifu string
	for idWaifu, waifu := range waifus {
		if id == idWaifu {
			validWaifu = waifu
		}
	}
	return id, validWaifu
}

// Variadic function

func calculate(numbers ...int) int {
	var total int
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	helloWorld()
	fmt.Println(sayHello("Re"))

	getThree()

	var id, waifu = getWaifu(1)
	fmt.Printf("%v-->%v\n", id, waifu)

	fmt.Printf("\n%v\n", calculate(1, 2, 3, 4, 5, 6, 7))

	// Closure

	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}
	var numbers = []int{2, 3, 4, 3, 4, 2, 3}
	var min, max = getMinMax(numbers)
	fmt.Printf("data : %v\nmin : %v\nmax : %v\n", numbers, min, max)
}
