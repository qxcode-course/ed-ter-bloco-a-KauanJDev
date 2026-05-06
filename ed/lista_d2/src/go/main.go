ackage main

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

func NewLList() List {
	var root Node

	root = Node{
		next: &root,
		prev: &root,
	}

	return List{
		root: &root,
	}
}

func (list *List) Size() int {
	size := 0
	for i := list.root.next; i != list.root; i = i.next {
		size++
	}
	return size
}

func remove(node *Node) {
	a := node.prev
	c := node.next
	a.next = c
	c.prev = a
	node.next = nil
	node.prev = nil
}

func (list *List) Remove(node *Node) {
	remove(node)
}

func (list *List) PopFront() {
	if list.root.next != list.root {
		remove(list.root.next)
	}
}

func (list *List) PopBack() {
	if list.root.prev != list.root {
		remove(list.root.prev)
	}
}

func (list *List) Clear() {
	for i := 0; list.Size() > 0; i++ {
		list.PopFront()
	}
}

func (list *List) Front() *Node {
	if list.root.next == list.root {
		return nil
	}
	return list.root.next
}

func (list *List) End() *Node {
	if list.root.prev == list.root {
		return nil
	}
	return list.root.prev
}

func (node *Node) Next(list *List) *Node {
	if node.next == list.root {
		return nil
	}
	return node.next
}

func (node *Node) Prev(list *List) *Node {
	if node.prev == list.root {
		return nil
	}
	return node.prev
}

func (list *List) Back() *Node {
	if list.root.prev == list.root {
		return nil
	}
	return list.root.prev
}

func (list *List) Search(value int) *Node {
	for node := list.root.next; node != list.root; node = node.next {
		if node.Value == value {
			return node
		}
	}
	return nil
}

func (list *List) Insert(node *Node, value int) {
	n := node.prev
	newNode := &Node{
		Value: value,
		next:  node,
		prev:  n,
	}
	n.next = newNode
	node.prev = newNode

}

func (list *List) PushBack(value int) {
	n := &Node{Value: value}
	n.prev = list.root.prev
	n.next = list.root
	list.root.prev.next = n
	list.root.prev = n
}

func (list *List) PushFront(value int) {
	n := &Node{Value: value}
	n.next = list.root.next
	n.prev = list.root
	list.root.next.prev = n
	list.root.next = n
}

func (list *List) String() string {
	string := "["
	for i := list.root.next; i != list.root; i = i.next {
		string += fmt.Sprintf("%v", i.Value)
		if i != list.root.prev {
			string += ", "
		}
	}
	string += "]"
	return string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewLList()

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
		case "walk":
			fmt.Print("[ ")
			for node := ll.Front(); node != nil; node = node.Next(&ll) {
				fmt.Printf("%v ", node.Value)
			}
			fmt.Print("]\n[ ")
			for node := ll.Back(); node != nil; node = node.Prev(&ll) {
				fmt.Printf("%v ", node.Value)
			}
			fmt.Println("]")
		case "replace":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				node.Value = newvalue
			} else {
				fmt.Println("fail: not found")
			}
		case "insert":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Insert(node, newvalue)
			} else {
				fmt.Println("fail: not found")
			}
		case "remove":
			oldvalue, _ := strconv.Atoi(args[1])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Remove(node)
			} else {
				fmt.Println("fail: not found")
			}
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
