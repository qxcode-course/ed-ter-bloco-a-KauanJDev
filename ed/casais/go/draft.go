package main

import "fmt"

func main() {
	var animaisTotal int
	var casaisTotal int
	fmt.Scan(&animaisTotal)
	animais := make([]int, animaisTotal)
	for i := 0; i < animaisTotal; i++ {
		fmt.Scan(&animais[i])
	}
	for i := 0; i < animaisTotal; i++ {
		for j := i + 1; j < animaisTotal; j++ {
			if animais[i]+animais[j] == 0 {
				casaisTotal++
				animais[i] = 999
				animais[j] = 999
				break
			}
		}

	}
	fmt.Println(casaisTotal)
}
