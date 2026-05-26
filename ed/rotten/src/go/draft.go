package main

import "fmt"

type Oranges struct {
    linha, coluna, minuto int
}

func Front(fila []Oranges) Oranges {
    return fila[0]
}
func Pop(fila []Oranges) []Oranges {
    return fila[1:]
}
func Push(fila []Oranges, o Oranges) []Oranges {
    return append(fila, o)
}
func Vizinhos(linha, coluna int) [][]int {
    return [][]int{{linha - 1, coluna}, {linha + 1, coluna}, {linha, coluna - 1}, {linha, coluna + 1}}
}

func main() {
    var linhas, colunas,minutos int
    fmt.Scan(&linhas, &colunas)
    matriz := make([][]int, linhas)

    for i := 0; i < linhas; i++ {
        matriz[i] = make([]int, colunas)
        for j := 0; j < colunas; j++ {
            fmt.Scan(&matriz[i][j])
        }
    }

    fila := make([]Oranges, 0)
    frescas := 0
    
    for i := 0; i < linhas; i++ {
        for j := 0; j < colunas; j++ {
            if matriz[i][j] == 2 {
                fila = Push(fila, Oranges{i, j, 0})
            } else if matriz[i][j] == 1 {
                frescas++
            }
        }
    }

    minutos = bfs(matriz, fila, frescas)
    fmt.Println(minutos)
}

func bfs(matriz [][]int, fila []Oranges, frescas int) int {
    minutos := 0
    for len(fila) > 0 {
        l := Front(fila)
        fila = Pop(fila)

        for _, vizinho := range Vizinhos(l.linha, l.coluna) {
            if vizinho[0] >= 0 && vizinho[0] < len(matriz) && vizinho[1] >= 0 && vizinho[1] < len(matriz[0]) {
                if matriz[vizinho[0]][vizinho[1]] == 1 {
                    matriz[vizinho[0]][vizinho[1]] = 2
                    frescas--
                    fila = Push(fila, Oranges{vizinho[0], vizinho[1], l.minuto + 1})
                    minutos = l.minuto + 1
                }
            }
        }
    }

    if frescas > 0 {
        return -1
    }

    return minutos
}


