package main

import "fmt"

func main() {
	var n, soma int
	fmt.Scan(&n, &soma)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		arr[i] = num
	}
	fmt.Println(subsetSum(arr, soma))
}

func subsetSum(arr []int, soma int) bool {
	return backtrack(arr, soma, 0)
}

func backtrack(arr []int, soma int, index int) bool {
	if soma == 0 {
		return true
	}
	if index == len(arr) {
		return false
	}
	if backtrack(arr, soma-arr[index], index+1) {
		return true
	}
	return backtrack(arr, soma, index+1)
}
