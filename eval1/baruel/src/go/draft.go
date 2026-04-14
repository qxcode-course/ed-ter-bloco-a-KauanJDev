package main

import "fmt"

func main() {
	var fig_album, possui int
	fmt.Scan(&fig_album, &possui)
	figurinhas := make([]int, possui)
	for i := range possui {
		fmt.Scan(&figurinhas[i])
	}
	encontradas := make(map[int]bool)
	repetidas := make([]int, 0, possui)

	for _, fig := range figurinhas {
		if encontradas[fig] {
			repetidas = append(repetidas, fig)
		}
		encontradas[fig] = true
	}

	if len(repetidas) == 0 {
		fmt.Println("N")
	} else {
		output := fmt.Sprintf("%v", repetidas)
		fmt.Println(output[1 : len(output)-1])
	}

	unicas := make([]int, 0, fig_album)

	for i := 1; i <= fig_album; i++ {
		if !encontradas[i] {
			unicas = append(unicas, i)
		}
	}

	if len(unicas) == 0 {
		fmt.Println("N")
	} else {
		output := fmt.Sprintf("%v", unicas)
		fmt.Println(output[1 : len(output)-1])
	}

}
