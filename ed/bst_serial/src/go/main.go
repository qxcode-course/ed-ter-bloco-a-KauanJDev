package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (node *Node) insert(value int) {
	if value == node.Value {
		return
	}
	if value < node.Value {
		if node.Left == nil {
			node.Left = &Node{Value: value}
		} else {
			node.Left.insert(value)
		}
	} else {
		if node.Right == nil {
			node.Right = &Node{Value: value}
		} else {
			node.Right.insert(value)
		}
	}
}

func BstInsert(values []int) *Node {
	if len(values) == 0 {
		return nil
	}
	node := &Node{Value: values[0]}
	for _, value := range values[1:] {
		node.insert(value)
	}
	return node
}

// Dica: crie um vetor compartilhado e vá preenchendo conforme anda na recursão
// Depois use o strings.Join para gerar o serial
func Serialize(root *Node) string {
	if root == nil {
		return ""
	}

	if root.Left != nil {
		Serialize(root.Left)
	}
	if root.Right != nil {
		Serialize(root.Right)
	}

	return strings.Join(vector, " ")
}

// -----------------------------------------------------------------------------------
func BShow(node *Node, history string) {
	if node != nil && (node.Left != nil || node.Right != nil) {
		BShow(node.Left, history+"l")
	}
	for i := 0; i < len(history)-1; i++ {
		if history[i] != history[i+1] {
			fmt.Print("│   ")
		} else {
			fmt.Print("    ")
		}
	}
	if history != "" {
		if history[len(history)-1] == 'l' {
			fmt.Print("╭───")
		} else {
			fmt.Print("╰───")
		}
	}
	if node == nil {
		fmt.Println("#")
		return
	}
	fmt.Println(node.Value)
	if node.Left != nil || node.Right != nil {
		BShow(node.Right, history+"r")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	values := make([]int, 0, len(parts))
	for _, elem := range parts {
		v, err := strconv.Atoi(elem)
		if err == nil {
			values = append(values, v)
		}
	}
	root := BstInsert(values)
	BShow(root, "") // Chama a função de impressão formatada
	fmt.Println(Serialize((root)))
}
