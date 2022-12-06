package main

import (
	"fmt"
	"reflect"
)

func main() {

	// Array

	var fruits = [3]string{"apel", "watermelon", "grape"}
	fmt.Println("Data :", fruits)
	fmt.Println("Type :", reflect.TypeOf(fruits), "--> This is array \n")

	// Multi dimension array
	var numbers = [3][2]int{[2]int{1, 2}, [2]int{3, 4}, [2]int{5, 6}}
	fmt.Println("Data :", numbers)
	fmt.Println("Type :", reflect.TypeOf(numbers), "--> This is array \n")

	// slice
	var waifus = []string{"Rem", "Chisato", "Onodera", "Kurumi"}
	fmt.Println("Data :", waifus)
	fmt.Println("Type :", reflect.TypeOf(waifus), "--> This is slice\n")

	var newWaifus = waifus[0:3]
	fmt.Println("waifus[0:3] :", newWaifus, "\n")

	// Slice Method

	// count slice length [ len() ]
	fmt.Printf("waifus length is : %v\n", len(waifus))

	// count length or maxsimum value of slice [ cap() ]
	fmt.Println("\nnewWaifus length is :", len(newWaifus))
	fmt.Println("newWaifus cap is :", cap(newWaifus))

	// insert item to slice
	waifus = append(waifus, "Elaina")
	fmt.Println("\n", waifus, "\n")

	// Map
	var students = map[string]string{
		"Name":  "Re",
		"waifu": "Rem",
	}

	fmt.Printf("data : %v\n", students)
	fmt.Printf("tipe : %v --> This is map\n", reflect.TypeOf(students))

	// looping map
	var months = map[string]string{
		"1": "January",
		"2": "February",
		"3": "Maret",
		"4": "April",
	}

	for number, month := range months {
		fmt.Printf("%v : %v", number, month)
	}

	// delete map item

}
