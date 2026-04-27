package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Set struct {
	data     []int
	size     int
	capacity int
}

func NewSet(capacity int) *Set {
	return &Set{
		data:     make([]int, 0, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (s *Set) Insert(value int) error {
	for i := 0; i < s.size; i++ {
		if s.data[i] == value {
			return nil
		}
	}

	posicao := s.size
	for i := 0; i < s.size; i++ {
		if s.data[i] > value {
			posicao = i
			break
		}
	}

	s.data = append(s.data, 0)
	for i := s.size; i > posicao; i-- {
		s.data[i] = s.data[i-1]
	}

	s.data[posicao] = value
	s.size++
	return nil
}

func (s *Set) Contains(value int) bool {
	for i := 0; i < s.size; i++ {
		if s.data[i] == value {
			return true
		}
	}
	return false
}

func (s *Set) Remove(value int) error {
	for i := 0; i < s.size; i++ {
		if s.data[i] == value {
			for j := i; j < s.size-1; j++ {
				s.data[j] = s.data[j+1]
			}
			s.size--
			return nil
		}
	}
	return errors.New("value not found")
}

func (s *Set) String() string {
	string := ""
	string += fmt.Sprintf("[")
	for i := 0; i < s.size; i++ {
		if i < s.size-1 {
			string += fmt.Sprintf("%v, ", s.data[i])
		} else {
			string += fmt.Sprintf("%v", s.data[i])
		}
	}
	string += fmt.Sprintf("]\n")
	return string
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewSet(0)
	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			v = NewSet(value)
		case "insert":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.Insert(value)
			}
		case "show":
			fmt.Print(v.String())
		case "erase":
			value, _ := strconv.Atoi(parts[1])
			err := v.Remove(value)
			if err != nil {
				fmt.Println(err)
			}
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			fmt.Println(v.Contains(value))
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
