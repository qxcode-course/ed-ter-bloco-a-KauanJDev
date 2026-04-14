package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		fmt.Scan(&b[i])
	}

	vencedor := 0
	diferenca := -1

	for i := 0; i < n; i++ {
		if a[i] < 10 || b[i] < 10 {
			continue
		}
		pontuacao := a[i] - b[i]
		if pontuacao < 0 {
			pontuacao = -pontuacao
		}
		if diferenca < 0 || pontuacao < diferenca {
			diferenca = pontuacao
			vencedor = i
		}
	}
	if vencedor == 0 {
		fmt.Println("sem ganhador")
	} else {
		fmt.Println(vencedor)
	}
}
