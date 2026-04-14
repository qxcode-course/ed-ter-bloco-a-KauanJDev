package main

import "fmt"

func main() {
	var gomos int
	var direcao string
	fmt.Scan(&gomos, &direcao)

	x := make([]int, gomos)
	y := make([]int, gomos)

	for i := 0; i < gomos; i++ {
		fmt.Scan(&x[i], &y[i])
	}

	for i := gomos - 1; i > 0; i-- {
		x[i] = x[i-1]
		y[i] = y[i-1]
	}

	switch direcao {
	case "D":
		y[0]++
	case "U":
		y[0]--
	case "R":
		x[0]++
	case "L":
		x[0]--
	}

	for i := 0; i < gomos; i++ {
		fmt.Printf("%d %d\n", x[i], y[i])
	}
}
