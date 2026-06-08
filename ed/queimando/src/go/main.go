package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

func getNeighbors(l, c int) [][]int {
	return [][]int{
		{l - 1, c},
		{l + 1, c},
		{l, c - 1},
		{l, c + 1},
	}
}

func burnTrees(grid [][]rune, l, c int) {
	stack := NewStack[Pos]()
	stack.Push(Pos{l, c})

	for !stack.IsEmpty() {
		pos := stack.Pop()

		if pos.l < 0 || pos.l >= len(grid) || pos.c < 0 || pos.c >= len(grid[0]) {
			continue
		}
		if grid[pos.l][pos.c] != '#' {
			continue
		}
		grid[pos.l][pos.c] = 'o'

		for _, neighbor := range getNeighbors(pos.l, pos.c) {
			stack.Push(Pos{neighbor[0], neighbor[1]})
		}
	}
	// Essa função deve usar uma list como pilha
	// e marcar as árvores na matriz como queimados
	// Uma sugestão de como fazer isso é:
	// - adicionar a primeira posição na pilha
	// - enquanto a pilha não estiver vazia:
	//   - retirar o elemento do topo
	//   - se puder ser queimado, queime e adicione seus vizinhos à pilha

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc, lfire, cfire int
	fmt.Sscanf(line, "%d %d %d %d", &nl, &nc, &lfire, &cfire)

	grid := make([][]rune, 0, nl)
	for range nl {
		scanner.Scan()
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}
	burnTrees(grid, lfire, cfire)
	showGrid(grid)
}

func showGrid(mat [][]rune) {
	for _, linha := range mat {
		fmt.Println(string(linha))
	}
}
