package main

import "fmt"

type student struct {
	name  string
	grade int
}

func (s student) getName() string {
	return s.name
}

func (s student) calculateGrade(point int) int {
	var grade int
	grade = s.grade + point

	return grade
}

// Before used method pointer (name value not change)

func (s student) changeName1(name string) {
	fmt.Println("---> on changeName1, name changed to", name)
	s.name = name
}

// After used method pointer (name value change)

func (s *student) changeName2(name string) {
	fmt.Println("---> on changeName2, name changed to", name)
	s.name = name
}

func main() {
	var re = student{"Re", 7}
	fmt.Println(re.getName())
	fmt.Println(re.calculateGrade(7), "\n")

	var s1 = student{"john wick", 21}

	fmt.Println("s1 before", s1.name)
	// john wick
	s1.changeName1("jason bourne")
	fmt.Println("s1 after changeName1", s1.name)
	// john wick
	s1.changeName2("ethan hunt")
	fmt.Println("s1 after changeName2", s1.name)
	// ethan hunt
}
