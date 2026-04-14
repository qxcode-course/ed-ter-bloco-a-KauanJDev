package main

import "fmt"

func main() {
	var pessoasNaFila, deixaram_fila int

	fmt.Scan(&pessoasNaFila)
	id := make([]int, pessoasNaFila)
	for i := range pessoasNaFila {
		fmt.Scan(&id[i])
	}

	fmt.Scan(&deixaram_fila)
	id_desertores := make([]int, deixaram_fila)
	for i := range deixaram_fila {
		fmt.Scan(&id_desertores[i])
	}

	for i := 0; i < pessoasNaFila; i++ {
		for j := 0; j < deixaram_fila; j++ {
			if id[i] == id_desertores[j] {
				id[i] = 0
			}
		}
	}
	for i := range len(id) {
		if id[i] != 0 {
			fmt.Printf("%v ", id[i])
		}
	}
	fmt.Println()
}
