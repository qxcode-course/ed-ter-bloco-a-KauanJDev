package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y, z float64

	fmt.Scan(&x, &y, &z)
	var semiPerimetro float64 = (x + y + z) / 2
	var area float64 = math.Sqrt(float64(semiPerimetro * (semiPerimetro - x) * (semiPerimetro - y) * (semiPerimetro - z)))
	fmt.Printf("%.2f\n", area)
}
