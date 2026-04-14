package main

import (
	"fmt"
	"sort"
)

func main() {
	var totalFigurinhas, figurinhasPossuidas int
	fmt.Scan(&totalFigurinhas)
	fmt.Scan(&figurinhasPossuidas)

	repetidas := make([]int, 0)
	faltantes := make([]int, 0)
	figurinhas := make([]int, figurinhasPossuidas)
	for i := 0; i < figurinhasPossuidas; i++ {
		fmt.Scan(&figurinhas[i])
	}
	sort.Ints(figurinhas)
	for i := 1; i < figurinhasPossuidas; i++ {
		if i >= 1 && figurinhas[i] == figurinhas[i-1] {
			repetidas = append(repetidas, figurinhas[i])
		}
	}
	for i := 1; i <= totalFigurinhas; i++ {
		encontrada := false
		for j := 0; j < len(figurinhas); j++ {
			if figurinhas[j] == i {
				encontrada = true
				break
			}
		}
		if !encontrada {
			faltantes = append(faltantes, i)
		}
	}
	if len(repetidas) == 0 {
		fmt.Printf("N")
	} else {
		for i := 0; i < len(repetidas); i++ {
			if i == len(repetidas)-1 {
				fmt.Printf("%d", repetidas[i])
				break
			}
			fmt.Printf("%d ", repetidas[i])
		}
	}
	fmt.Printf("\n")
	if len(faltantes) == 0 {
		fmt.Printf("N\n")
	} else {
		for i := 0; i < len(faltantes); i++ {
			if i == len(faltantes)-1 {
				fmt.Printf("%d\n", faltantes[i])
				break
			}
			fmt.Printf("%d ", faltantes[i])
		}
	}
}
