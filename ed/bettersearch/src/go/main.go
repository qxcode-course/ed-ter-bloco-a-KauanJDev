package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BetterSearch(slice []int, value int) (bool, int) {
	for i, elem := range slice {
		if elem == value {
			return true, i
		}
		if elem > value {
			return false, i
		}
	}
	return false, len(slice)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	slice := []int{}
	for _, elem := range parts[1 : len(parts)-1] {
		value, _ := strconv.Atoi(elem)
		slice = append(slice, value)
	}
	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	found, result := BetterSearch(slice, value)
	if found {
		fmt.Println("V", result)
	} else {
		fmt.Println("F", result)
	}
}
