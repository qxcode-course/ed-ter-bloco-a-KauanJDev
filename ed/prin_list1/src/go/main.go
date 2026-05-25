package main

import (
	"fmt"
)

// mostra a lista com o elemento sword destacado
func ToStr(l *DList[int], sword *DNode[int]) string {
	string := "["
	for it := l.root.next; it != l.root; it = it.next {
		if it == sword {
			string += fmt.Sprintf(" %d>", it.Value)
		} else {
			string += fmt.Sprintf(" %d", it.Value)
		}
	}
	string += " ]"
	return fmt.Sprint(string)
}

// move para frente na lista circular
func Next(l *DList[int], it *DNode[int]) *DNode[int] {
	next := it.next
	if next == it.root {
		next = next.next
	}
	return next
}

func main() {
	var qtd, chosen int
	fmt.Scan(&qtd, &chosen)
	l := NewDList[int]()
	for i := 1; i <= qtd; i++ {
		l.PushBack(i)
	}
	sword := l.Front()
	for range chosen - 1 {
		sword = Next(l, sword)
	}
	for range qtd - 1 {
		fmt.Println(ToStr(l, sword))
		l.Erase(Next(l, sword))
		sword = Next(l, sword)
	}
	fmt.Println(ToStr(l, sword))
}
