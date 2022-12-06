package main

import "fmt"

func main() {
	for i := 0; i < 7; i++ {
		fmt.Println("ini index ke->", i)
	}

	fmt.Println("")

	var waifu = []string{"Rem", "Chitoge", "Onodera"}
	for i, w := range waifu {
		fmt.Printf("%s adalah waifu ke->%v\n", w, i)
	}
}
