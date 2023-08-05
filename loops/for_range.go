package main

import "fmt"

func main() {
	var datas []string
	datas = []string{"nadia", "levi", "kyati"}
	for key, data := range datas {
		fmt.Println("Ini data", key, "dengan nama", data)
	}
}
