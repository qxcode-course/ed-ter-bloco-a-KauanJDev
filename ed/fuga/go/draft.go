package main

import "fmt"

const maxPosition = 15

func main() {
	var helicoptero, policial, fugitivo, direcao int
	fmt.Scan(&helicoptero, &policial, &fugitivo, &direcao)
	if direcao == 1 {
		for fugitivo != helicoptero && fugitivo != policial {
			fugitivo++
			if fugitivo > maxPosition {
				fugitivo = 0
			}
		}
	}
	if direcao == -1 {
		for fugitivo != helicoptero && fugitivo != policial {
			fugitivo--
			if fugitivo < 0 {
				fugitivo = maxPosition
			}
		}
	}
	if fugitivo == helicoptero {
		fmt.Println("S")
	} else {
		fmt.Println("N")
	}
}
