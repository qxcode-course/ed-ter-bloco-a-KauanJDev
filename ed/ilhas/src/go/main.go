package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

func (p Pos) NewPos() Pos {
	return Pos{l: p.l, c: p.c}
}

// Não modifique a assinatura da função numIslands
// Ela é a função que será chamada no LeetCode para resolver o problema
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	visited := make(map[Pos]bool)
	cols := len(grid[0])
	count := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' && !visited[Pos{i, j}] {
				count++
				p := Pos{l: i, c: j}
				p.dfs(grid, i, j, visited, cols)
			}
		}
	}
	return count
}

func (p Pos) dfs(grid [][]byte, i int, j int, visited map[Pos]bool, cols int) {
	pos := Pos{l: i, c: j}
	if i < 0 || i >= len(grid) || j < 0 || j >= cols {
		return
	}

	if visited[pos] || grid[i][j] == '0' {
		return
	}

	visited[pos] = true

	p.dfs(grid, i, j+1, visited, cols)
	p.dfs(grid, i, j-1, visited, cols)
	p.dfs(grid, i+1, j, visited, cols)
	p.dfs(grid, i-1, j, visited, cols)
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	grid := make([][]byte, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		grid[i] = []byte(scanner.Text())
	}
	result := numIslands(grid)
	fmt.Println(result)
}
