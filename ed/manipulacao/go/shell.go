package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getMen(vet []int) []int {
	var men []int
	for i := 0; i < len(vet); i++ {
		if vet[i] >= 0 {
			men = append(men, vet[i])
		}
	}
	return men
}

func getCalmWomen(vet []int) []int {
	var calmWomen []int
	for i := 0; i < len(vet); i++ {
		if vet[i] < 0 && vet[i] > -10 {
			calmWomen = append(calmWomen, vet[i])
		}
	}
	return calmWomen
}

func sortVet(vet []int) []int {
	for i := 0; i < len(vet)-1; i++ {
		for j := 0; j < len(vet)-i-1; j++ {
			if vet[j] > vet[j+1] {
				vet[j], vet[j+1] = vet[j+1], vet[j]
			}
		}
	}
	return vet
}

func negative(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func sortStress(vet []int) []int {
	for i := 0; i < len(vet)-1; i++ {
		for j := 0; j < len(vet)-i-1; j++ {
			if negative(vet[j]) > negative(vet[j+1]) {
				vet[j], vet[j+1] = vet[j+1], vet[j]
			}
		}
	}
	return vet
}

func reverse(vet []int) []int {
	reversed := make([]int, len(vet))
	for i := 0; i < len(vet); i++ {
		for j := len(reversed) - 1; j >= 0; j-- {
			reversed[j] = vet[i]
			i++
		}
	}
	return reversed
}

func unique(vet []int) []int {
	var result []int
	for i := 0; i < len(vet); i++ {
		isUnique := true
		for j := 0; j < i; j++ {
			if vet[i] == vet[j] && i != j {
				isUnique = false
				break
			}
		}
		if isUnique {
			result = append(result, vet[i])
		}
	}
	return result
}

func repeated(vet []int) []int {
	var result []int
	for i := 0; i < len(vet); i++ {
		isRepeated := false
		for j := 0; j < i; j++ {
			if vet[i] == vet[j] && i != j {
				isRepeated = true
				break
			}
		}
		if isRepeated {
			result = append(result, vet[i])
		}
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}
