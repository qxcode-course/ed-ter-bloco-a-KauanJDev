package main

import "fmt"

func main() {
	var numBombs, numAttributes int
	fmt.Scan(&numBombs, &numAttributes)
	bombs := make([][]int, numBombs)
	for i := 0; i < numBombs; i++ {
		bombs[i] = make([]int, numAttributes)
		for j := 0; j < numAttributes; j++ {
			var valor int
			fmt.Scan(&valor)
			bombs[i][j] = valor
		}
	}
	fmt.Println(maxDetonation(bombs))
}

func maxDetonation(bombs [][]int) int {
	n := len(bombs)
	graph := make(map[int][]int)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j && canDetonate(bombs[i], bombs[j]) {
				graph[i] = append(graph[i], j)
			}
		}
	}

	max := 0

	for i := 0; i < n; i++ {
		visited := bfs(graph, i)
		if len(visited) > max {
			max = len(visited)
		}
	}
	return max
}

func bfs(graph map[int][]int, start int) map[int]bool {
	visited := make(map[int]bool)
	queue := []int{start}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if !visited[current] {
			visited[current] = true
			for _, neighbor := range graph[current] {
				if !visited[neighbor] {
					queue = append(queue, neighbor)
				}
			}
		}
	}
	return visited
}

func canDetonate(bomb1, bomb2 []int) bool {
	xi, yi, ri := bomb1[0], bomb1[1], bomb1[2]
	xj, yj := bomb2[0], bomb2[1]

	dx := xi - xj
	dy := yi - yj

	distanciaQuadrado := dx*dx + dy*dy
	raioQuadrado := ri * ri

	return distanciaQuadrado <= raioQuadrado
}
