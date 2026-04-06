package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

func search(grid [][]rune, startPos, endPos Pos) bool {
	if startPos.l < 0 || startPos.l >= len(grid) ||
		startPos.c < 0 || startPos.c >= len(grid[startPos.l]) {
		return false
	}

	if startPos == endPos {
		grid[startPos.l][startPos.c] = '.'
		return true
	}

	if grid[startPos.l][startPos.c] != ' ' {
		return false
	}

	grid[startPos.l][startPos.c] = '.'

	proximo := vizinhos(startPos.l, startPos.c)
	if search(grid, proximo[0], endPos) ||
		search(grid, proximo[1], endPos) ||
		search(grid, proximo[2], endPos) ||
		search(grid, proximo[3], endPos) {
		return true
	}
	grid[startPos.l][startPos.c] = ' '
	return false
}

func vizinhos(l, c int) []Pos {
	return []Pos{
		{l - 1, c},
		{l + 1, c},
		{l, c - 1},
		{l, c + 1},
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nl_nc := scanner.Text()
	var nl, nc int
	fmt.Sscanf(nl_nc, "%d %d", &nl, &nc)
	grid := make([][]rune, nl)

	// Lê a gridriz
	for i := range nl {
		scanner.Scan()
		grid[i] = []rune(scanner.Text())
	}

	// Procura posições de início e endPos e conserta para _
	var startPos, endPos Pos
	for l := range nl {
		for c := range nc {
			if grid[l][c] == 'I' {
				grid[l][c] = ' '
				startPos = Pos{l, c}
			}
			if grid[l][c] == 'F' {
				grid[l][c] = ' '
				endPos = Pos{l, c}
			}
		}
	}

	search(grid, startPos, endPos)

	// Imprime o labirinto final
	for _, line := range grid {
		fmt.Println(string(line))
	}
}
