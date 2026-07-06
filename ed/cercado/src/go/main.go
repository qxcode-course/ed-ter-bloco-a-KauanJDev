package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

func (p Point) NewPoint() Point {
	return Point{x: p.x, y: p.y}
}

// NÃO ALTERE A ASSINATURA DA FUNÇÃO solve
func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}
	visited := make(map[Point]bool)
	rows := len(board)
	cols := len(board[0])
	for i := 0; i < len(board); i++ {
		for j := 0; j < cols; j++ {
			if (i == 0 || i == rows-1 || j == 0 || j == cols-1) && board[i][j] == 'O' {
				p := Point{x: i, y: j}
				p.dfs(board, i, j, visited, cols)
			}
		}
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			p := Point{x: i, y: j}
			if !visited[p] && board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func (p Point) dfs(board [][]byte, i, j int, visited map[Point]bool, cols int) {
	pos := Point{x: i, y: j}
	if i < 0 || i >= len(board) || j < 0 || j >= cols || board[i][j] != 'O' || visited[pos] {
		return
	}

	visited[pos] = true

	pos.dfs(board, i+1, j, visited, cols)
	pos.dfs(board, i-1, j, visited, cols)
	pos.dfs(board, i, j+1, visited, cols)
	pos.dfs(board, i, j-1, visited, cols)
}

// NÃO ALTERE A MAIN
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var nrows, ncols int
	fmt.Sscanf(scanner.Text(), "%d %d", &nrows, &ncols)
	board := make([][]byte, nrows)
	for i := 0; i < nrows; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}
	solve(board)
	for _, row := range board {
		fmt.Println(string(row))
	}
}
