package main

import "fmt"

func calcularDivisoes(n int) []int {
	if n == 0 {
		return []int{}
	}
	divisoes := calcularDivisoes(n / 2)
	return append(divisoes, n/2)
}

func calcularResto(n int) []int {
	if n == 0 {
		return []int{}
	}
	divisao := n / 2
	restos := calcularResto(divisao)
	return append(restos, n%2)
}

func main() {
	var n int
	fmt.Scan(&n)
	divisoes := calcularDivisoes(n)
	restos := calcularResto(n)
	for _, d := range divisoes {
		fmt.Printf("%d %d\n", d, restos[0])
		restos = restos[1:]
	}

}
