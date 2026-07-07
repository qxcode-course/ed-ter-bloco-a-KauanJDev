package main

import (
	"fmt"
	"strconv"
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
	}
}

func entrada(s string) []int {
	sequencia := make([]int, len(s))
	for i, c := range s {
		if c == '.' {
			sequencia[i] = -1
		} else {
			sequencia[i], _ = strconv.Atoi(string(c))
		}
	}
	return sequencia
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
