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
	next  *Node
	prev  *Node
}

type List struct {
	root *Node
}

func NewList() List {
	var root Node
	root = Node{
		next: &root,
		prev: &root,
	}

	return List{
		root: &root,
	}
}

func (list *List) PushFront(value int) {
	n := &Node{Value: value}
	n.next = list.root.next
	n.prev = list.root
	list.root.next.prev = n
	list.root.next = n
}

func (list *List) PushBack(value int) {
	n := &Node{Value: value}
	n.prev = list.root.prev
	n.next = list.root
	list.root.prev.next = n
	list.root.prev = n
}

func Remove(node *Node) {
	a := node.prev
	c := node.next
	a.next = c
	c.prev = a
	node.next = nil
	node.prev = nil
}

func (list *List) Size() int {
	size := 0
	for i := list.root.next; i != list.root; i = i.next {
		size++
	}
	return size
}

func (list *List) Clear() {
	for i := 0; list.Size() > 0; i++ {
		list.PopFront()
	}
}

func (list *List) PopFront() {
	if list.root.next != list.root {
		Remove(list.root.next)
	}
}
func (list *List) PopBack() {
	if list.root.prev != list.root {
		Remove(list.root.prev)
	}
}

func (list *List) String() string {
	s := "["
	for i := list.root.next; i != list.root; i = i.next {
		s += fmt.Sprint(i.Value)
		if i != list.root.prev {
			s += ", "
		}
	}
	s += "]"
	return s
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewList()

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			fmt.Println(ll.String())
		case "size":
			fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushFront(num)
			}
		case "pop_back":
			ll.PopBack()
		case "pop_front":
			ll.PopFront()
		case "clear":
			ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
