package main

import "fmt"

func main() {
	score := 81

	// switch {
	// case score == 100:
	// 	fmt.Println("Perfect")
	// case (score < 80) && (score > 60):
	// 	fmt.Println("Not bad")
	// default:
	// 	fmt.Println("Please study harder")
	// }

	if score > 80 {
		switch {
		case score == 100:
			fmt.Println("Perfect")
		default:
			fmt.Println("Nice score")
		}
	} else {
		if score == 50 {
			fmt.Println(("not bad"))
		} else {
			fmt.Println("keep trying")
		}
	}

}
