package main

import "fmt"

func printNumeros(numeros []int, N int, E int) {
	fmt.Print("[ ")
	for i := 1; i <= N; i++ {
		if numeros[i] == 0 {
			continue
		}

		if numeros[i] == E {
			fmt.Printf("%d> ", numeros[i])
		} else {
			fmt.Printf("%d ", numeros[i])
		}
	}
	fmt.Print("]\n")
}

func proximo(numeros []int, N int, atual int) int {
	for i := atual + 1; i <= N; i++ {
		if numeros[i] != 0 {
			return i
		}
	}
	for i := 1; i <= N; i++ {
		if numeros[i] != 0 {
			return i
		}
	}
	return -1
}

func validos(numeros []int, N int) int {
	validos := 0
	for i := 1; i <= N; i++ {
		if numeros[i] != 0 {
			validos++
		}
	}
	return validos
}

func main() {
	var N, E int
	fmt.Scan(&N, &E)

	numeros := make([]int, N+1)
	var espada int

	for i := 1; i <= N; i++ {
		numeros[i] = i
		if i == E {
			espada = i
		}
	}
	for validos(numeros, N) > 0 {
		printNumeros(numeros, N, espada)
		if validos(numeros, N) == 1 {
			break
		}
		proximoIdx := proximo(numeros, N, espada)
		numeros[proximoIdx] = 0

		espada = proximo(numeros, N, proximoIdx)
	}
}
