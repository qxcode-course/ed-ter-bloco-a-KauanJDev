package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MultiSet struct {
	data     []int
	size     int
	capacity int
}

func NewMultiSet(capacity int) *MultiSet {
	return &MultiSet{
		data:     make([]int, 0, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (ms *MultiSet) Insert(value int) {
	posicao := ms.size
	for i := 0; i < posicao; i++ {
		if ms.data[i] > value {
			posicao = i
			break
		}
	}

	ms.data = append(ms.data, 0)
	for i := ms.size; i > posicao; i-- {
		ms.data[i] = ms.data[i-1]
	}
	ms.data[posicao] = value
	ms.size++
}

func (ms *MultiSet) String() string {
	string := ""
	string += fmt.Sprintf("[")
	for i := 0; i < ms.size; i++ {
		if i < ms.size-1 {
			string += fmt.Sprintf("%v, ", ms.data[i])
		} else {
			string += fmt.Sprintf("%v", ms.data[i])
		}
	}
	string += fmt.Sprintf("]")
	return string
}

func (ms *MultiSet) Contains(value int) bool {
	for i := 0; i < ms.size; i++ {
		if ms.data[i] == value {
			return true
		}
	}
	return false
}

func (ms *MultiSet) Erase(value int) error {
	if !ms.Contains(value) {
		return errors.New("value not found")
	}
	for i := 0; i < ms.size; i++ {
		if ms.data[i] == value {
			for j := i; j < ms.size-1; j++ {
				ms.data[j] = ms.data[j+1]
			}
			ms.size--
			return nil
		}
	}
	return nil
}

func (ms *MultiSet) Count(value int) int {
	count := 0
	for i := 0; i < ms.size; i++ {
		if ms.data[i] == value {
			count++
		}
	}
	return count
}

func (ms *MultiSet) Unique() int {
	unique := make([]int, 0, ms.size)
	for i := 0; i < ms.size; i++ {
		if !contains(unique, ms.data[i]) {
			unique = append(unique, ms.data[i])
		}
	}
	return len(unique)
}

func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func (ms *MultiSet) Clear() {
	ms.data = make([]int, 0, ms.capacity)
	ms.size = 0
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	ms := NewMultiSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(value)
		case "insert":
			for _, part := range args[1:] {
				value, _ := strconv.Atoi(part)
				ms.Insert(value)
			}
		case "show":
			fmt.Println(ms.String())
		case "erase":
			value, _ := strconv.Atoi(args[1])
			err := ms.Erase(value)
			if err != nil {
				fmt.Println(err)
			}

		case "contains":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Contains(value))
		case "count":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Count(value))
		case "unique":
			fmt.Printf("%d\n", ms.Unique())
		case "clear":
			ms.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
