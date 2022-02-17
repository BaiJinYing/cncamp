package main

import "fmt"

func main() {
	var strArray = []string{"I", "am", "stupid", "and", "weak"}
	fmt.Println(strArray)
	for i, str := range strArray {
		if str == "stupid" {
			strArray[i] = "smart"
		}
		if str == "weak" {
			strArray[i] = "strong"
		}
	}
	fmt.Println(strArray)
}
