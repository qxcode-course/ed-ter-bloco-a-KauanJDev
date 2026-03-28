package main

import "fmt"

func devePermancer(fila []int, saiu []int) []int {
	for i := 0; i < len(fila); i++ {
		for j := 0; j < len(saiu); j++ {
			if fila[i] == saiu[j] {
				fila[i] = 0
				break
			}
		}
	}
	return fila
}

func main() {
	var n, m int
	fmt.Scan(&n)
	fila := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&fila[i])
	}

	fmt.Scan(&m)
	sairam := make([]int, m)

	for i := 0; i < m; i++ {
		fmt.Scan(&sairam[i])
	}

	novaFila := make([]int, n)

	novaFila = devePermancer(fila, sairam)

	for i := 0; i < n; i++ {
		if novaFila[i] != 0 {
			fmt.Print(novaFila[i], " ")
		}
	}
	fmt.Print("\n")
}
