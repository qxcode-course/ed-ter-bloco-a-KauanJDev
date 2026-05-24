package main

import (
	"fmt"
)

func main() {
	var sequenceString string
	var l int

	fmt.Scan(&sequenceString, &l)
	sequencia := entrada(sequenceString)
	if backtrack(0, sequencia, l) {
		for _, num := range sequencia {
			fmt.Print(num, "")
		}
		fmt.Println()
	} else {
		fmt.Println("deu ruim")
	}
}

func entrada(entrada string) []int {
	sequenciaString := make([]int, len(entrada))
	for i := 0; i < len(entrada); i++ {
		if entrada[i] == '.' {
			sequenciaString[i] = -1
		} else {
			sequenciaString[i] = int(entrada[i] - '0')
		}
	}
	return sequenciaString
}

func posicaoValida(sequencia []int, index, num, distanciaMinima int) bool {
	for distancia := 1; distancia <= distanciaMinima; distancia++ {
		esquerda := index - distancia
		direita := index + distancia

		if esquerda >= 0 && sequencia[esquerda] == num {
			return false
		}
		if direita < len(sequencia) && sequencia[direita] == num {
			return false
		}
	}
	return true
}

func backtrack(pos int, sequencia []int, distanciaMinima int) bool {
	for pos < len(sequencia) && sequencia[pos] != -1 {
		pos++
	}
	if pos == len(sequencia) {
		return true
	}

	for i := 0; i <= distanciaMinima; i++ {
		if posicaoValida(sequencia, pos, i, distanciaMinima) {
			sequencia[pos] = i
			if backtrack(pos+1, sequencia, distanciaMinima) {
				return true
			}
			sequencia[pos] = -1
		}
	}
	return false
}
