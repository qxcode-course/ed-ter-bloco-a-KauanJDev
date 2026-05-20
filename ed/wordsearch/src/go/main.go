package main

import (
	"bufio"
	"fmt"
	"os"
)

// Não mude a assinatura desta função, ela é a função chamada pelo LeetCode
func exist(grid [][]byte, word string) bool {
	palavraEncontrada := false
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == word[0] {
				if backtrack(grid, word, i, j, 0) {
					palavraEncontrada = true
					break
				}
			}
		}
		if palavraEncontrada {
			break
		}
	}
	return palavraEncontrada
}

func backtrack(grid [][]byte, word string, row, col int, index int) bool {
	if index == len(word) {
		return true
	}
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) || grid[row][col] != word[index] {
		return false
	}
	temp := grid[row][col]
	grid[row][col] = '#'
	found := backtrack(grid, word, row+1, col, index+1) ||
		backtrack(grid, word, row-1, col, index+1) ||
		backtrack(grid, word, row, col+1, index+1) ||
		backtrack(grid, word, row, col-1, index+1)
	grid[row][col] = temp
	return found
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)
	grid := make([][]byte, 0)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
