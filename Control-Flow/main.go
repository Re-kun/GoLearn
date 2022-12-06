package main

import "fmt"

func main() {
	var data string
	var number int = 7
	fmt.Print("Masukan Data : ")
	fmt.Scanln(&data)

	if data == "makan" {
		fmt.Println("Mau makan apa bang ?")
	} else if data == "minum" {
		fmt.Println("Mau minum apa bang ?")
	} else {
		fmt.Println("Terimakasih sudah memesan")
	}

	// Variabel temporary

	if result := number + 2; result > 0 {
		fmt.Printf("%v < 0 : %v\n", result, result)
	}

	// Switch

	var point = 6
	switch point {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	// If u want to use many condition
	case 6, 5, 4:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	/*
		or u can do like this :
		default:
			{
				fmt.Println("not bad")
				fmt.Println("you need to learn more")
			}

		use 'fallthrough' to end statement like break
	*/
}
