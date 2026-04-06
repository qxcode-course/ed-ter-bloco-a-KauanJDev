package main

import (
	"fmt"
)

func mdc(a, b int) int {
	resto := a % b
	if resto == 0 {
		return b
	}
	return mdc(b, resto)
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(mdc(a, b))
}
