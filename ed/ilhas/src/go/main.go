package main

import (
	"bufio"
	"fmt"
	"os"
)

// Não modifique a assinatura da função numIslands
// Ela é a função que será chamada no LeetCode para resolver o problema
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	visited := make(map[int]bool)
	cols := len(grid[0])
	count := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' && !visited[i*cols+j] {
				count++
				dfs(grid, i, j, visited, cols)
			}
		}
	}
	return count
}

func dfs(grid [][]byte, i int, j int, visited map[int]bool, cols int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == '0' {
		return
	}

	key := i*cols + j
	if visited[key] {
		return
	}
	visited[key] = true

	dfs(grid, i, j+1, visited, cols)
	dfs(grid, i, j-1, visited, cols)
	dfs(grid, i+1, j, visited, cols)
	dfs(grid, i-1, j, visited, cols)
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
