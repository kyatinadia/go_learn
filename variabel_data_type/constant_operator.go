package main

import (
	"fmt"
)

const (
	statusNew = iota
	status0ld
)

func main() {
	// fmt.Println(statusNew)

	// stringNumber := "1234"

	// num, _ := strconv.Atoi(stringNumber)
	// fmt.Println(num)
	// fmt.Printf("Tipe data num %T", num)

	//arimetic
	result := (2 + 2) * 3
	fmt.Println(result)

	//relational
	var firstCondition bool = 2 < 3
	var secondCondition bool = "women" == "Women"
	fmt.Println("First condition", firstCondition)
	fmt.Println("Second condition", secondCondition)

	//logical
	right := true
	wrong := false

	wrongAndRight := right && wrong
	fmt.Printf("wrong && right \t(%t) \n", wrongAndRight)

	wrongOrRight := right || wrong
	fmt.Printf("wrong || right \t(%t) \n", wrongOrRight)

}
