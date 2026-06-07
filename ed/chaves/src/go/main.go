package main

import "fmt"

func main() {
	var mm, nn int
	resultados := [15][2]int{}

	for i := range resultados {
		fmt.Scan(&mm, &nn)
		resultados[i][0] = mm
		resultados[i][1] = nn
	}
	campeao := resultadoJogos(resultados)
	fmt.Println(campeao)
}

func resultadoJogos(resultados [15][2]int) string {
	fila := NewQueue[string]()
	fila.Enqueue("A")
	fila.Enqueue("B")
	fila.Enqueue("C")
	fila.Enqueue("D")
	fila.Enqueue("E")
	fila.Enqueue("F")
	fila.Enqueue("G")
	fila.Enqueue("H")
	fila.Enqueue("I")
	fila.Enqueue("J")
	fila.Enqueue("K")
	fila.Enqueue("L")
	fila.Enqueue("M")
	fila.Enqueue("N")
	fila.Enqueue("O")
	fila.Enqueue("P")
	indice := 0
	for fila.Len() > 1 {
		timeEsquerda := fila.Dequeue()
		timeDireita := fila.Dequeue()

		mm, nn := resultados[indice][0], resultados[indice][1]
		indice++
		if mm > nn {
			fila.Enqueue(timeEsquerda)
		} else {
			fila.Enqueue(timeDireita)
		}
	}
	return fila.Dequeue()
}
