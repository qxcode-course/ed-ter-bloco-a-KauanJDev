package main

import "fmt"

func main() {
	var gomos, posicaoX, posicaoY int
	var direcao string
	var antX, antY, tempX, tempY int
	fmt.Scan(&gomos, &direcao)
	for i := 0; i < gomos; i++ {
		fmt.Scan(&posicaoX, &posicaoY)
		if i == 0 {
			antX = posicaoX
			antY = posicaoY

			switch direcao {
			case "D":
				posicaoY++
			case "U":
				posicaoY--
			case "R":
				posicaoX++
			case "L":
				posicaoX--
			}
		} else {
			tempX = posicaoX
			tempY = posicaoY
			posicaoX = antX
			posicaoY = antY
			antX = tempX
			antY = tempY
		}

		fmt.Printf("%d %d\n", posicaoX, posicaoY)
	}
}
