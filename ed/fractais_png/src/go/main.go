package main

import (
	"fmt"
	"math/rand"
)

func ri(inf, sup int) float64 {
	return float64(rand.Intn(sup-inf+1) + inf)
}

func arvore(pen *Pen, distancia float64) {
	if distancia < 10 {
		return
	}

	ang_dir := ri(10, 40)
	ang_esq := ri(10, 40)

	pen.SetLineWidth(distancia / 5)
	pen.SetRGB(0, 0, 0)
	pen.Walk(distancia)
	pen.Right(ang_dir)
	arvore(pen, distancia*(ri(80, 85)/100))
	pen.Left(ang_dir + ang_esq)
	arvore(pen, distancia*(ri(80, 85)/100))
	pen.Right(ang_esq)
	pen.Walk(-distancia)

}
func circulo(pen *Pen, raio float64) {
	if raio < 2 {
		return
	}
	pen.SetRGB(0, 0, 0)
	pen.DrawCircle(raio)

	raioFilho := raio / 3.0

	for i := 0; i < 6; i++ {
		pen.Up()
		pen.Walk(raio)
		pen.Down()

		circulo(pen, raioFilho)

		pen.Up()
		pen.Walk(-raio)
		pen.Down()

		pen.Right(60)
	}
}
func gelo(pen *Pen, linha float64) {
	if linha < 2 {
		return
	}
	pen.SetRGB(0, 0, 0)
	for i := 0; i < 5; i++ {
		pen.Right(72)
		pen.Walk(linha)
		gelo(pen, linha/3)
		pen.Walk(-linha)
	}

}
func quadrado(pen *Pen, lado float64) {
	if lado < 4 {
		return
	}

	pen.SetRGB(0, 0, 0)
	pen.DrawRect(lado, lado)

	sub := lado / 2
	sub2 := sub / 2

	pen.Up()
	pen.SetHeading(180)
	pen.Walk(sub2)
	pen.SetHeading(90)
	pen.Walk(sub2)
	pen.Down()
	quadrado(pen, sub)

	pen.Up()
	pen.SetHeading(270)
	pen.Walk(sub2)
	pen.SetHeading(0)
	pen.Walk(sub2)

	pen.SetHeading(0)
	pen.Walk(lado - sub2)
	pen.SetHeading(90)
	pen.Walk(sub2)
	pen.Down()
	quadrado(pen, sub)

	pen.Up()
	pen.SetHeading(270)
	pen.Walk(sub2)
	pen.SetHeading(180)
	pen.Walk(lado - sub2)

	pen.SetHeading(180)
	pen.Walk(sub2)
	pen.SetHeading(270)
	pen.Walk(lado - sub2)
	pen.Down()
	quadrado(pen, sub)

	pen.Up()
	pen.SetHeading(90)
	pen.Walk(lado - sub2)
	pen.SetHeading(0)
	pen.Walk(sub2)

	pen.SetHeading(0)
	pen.Walk(lado - sub2)
	pen.SetHeading(270)
	pen.Walk(lado - sub2)
	pen.Down()
	quadrado(pen, sub)

	pen.Up()
	pen.SetHeading(90)
	pen.Walk(lado - sub2)
	pen.SetHeading(180)
	pen.Walk(lado - sub2)
	pen.Down()

}
func espiral(pen *Pen, linha float64) {
	if linha < 2 {
		return
	}
	pen.SetRGB(ri(0, 255), ri(0, 255), ri(0, 255))
	pen.Walk(linha)
	pen.Right(90)
	espiral(pen, linha-10)
}
func carpete(pen *Pen, linha float64) {
	if linha < 2 {
		return
	}
	sub := linha / 3

	pen.SetRGB(255, 255, 255)
	pen.FillSquare(linha, linha)
	pen.SetRGB(0, 0, 0)
	pen.Up()
	pen.SetHeading(0)
	pen.Walk(sub)
	pen.SetHeading(270)
	pen.Walk(sub)
	pen.FillSquare(sub, sub)
	pen.SetHeading(90)
	pen.Walk(sub)
	pen.SetHeading(180)
	pen.Walk(sub)

	carpete(pen, sub)
	pen.Up()
	pen.SetHeading(0)
	pen.Walk(sub)
	pen.Down()
	carpete(pen, sub)
	pen.Up()
	pen.SetHeading(0)
	pen.Walk(sub)
	pen.Down()
	carpete(pen, sub)

	pen.Up()
	pen.SetHeading(180)
	pen.Walk(sub * 2)
	pen.SetHeading(270)
	pen.Walk(sub)
	pen.Down()
	carpete(pen, sub)

	pen.Up()
	pen.SetHeading(0)
	pen.Walk(sub * 2)
	pen.Down()
	carpete(pen, sub)

	pen.Up()
	pen.SetHeading(180)
	pen.Walk(sub * 2)
	pen.SetHeading(270)
	pen.Walk(sub)
	pen.Down()
	carpete(pen, sub)

	pen.Up()
	pen.SetHeading(0)
	pen.Walk(sub)
	pen.Down()
	carpete(pen, sub)

	pen.Up()
	pen.SetHeading(0)
	pen.Walk(sub)
	pen.Down()
	carpete(pen, sub)

	pen.Up()
	pen.SetHeading(180)
	pen.Walk(sub * 2)
	pen.SetHeading(90)
	pen.Walk(sub * 2)
	pen.Down()
}
func triangulo(pen *Pen, lado float64) {
	if lado < 2 {
		return
	}

	metade := lado / 2

	pen.SetRGB(0, 0, 0)
	for range 3 {
		pen.Walk(lado)
		pen.Right(120)
	}

	triangulo(pen, metade)

	pen.Up()
	pen.Walk(metade)
	pen.Down()
	triangulo(pen, metade)

	pen.Up()
	pen.Right(180)
	pen.Walk(metade)
	pen.Left(180)
	pen.Down()

	pen.Up()
	pen.Right(60)
	pen.Walk(metade)
	pen.Left(60)
	pen.Down()
	triangulo(pen, metade)

	pen.Up()
	pen.Left(120)
	pen.Walk(metade)
	pen.Right(120)
	pen.Down()

}
func ramo(pen *Pen, linha float64) {
	if linha < 2 {
		return
	}

	quarto := linha / 4

	pen.SetRGB(255, 0, 0)
	pen.Walk(linha)
	pen.Up()
	pen.Walk(-linha)
	pen.Walk(quarto)
	pen.Right(60)
	pen.Down()
	ramo(pen, quarto)
	pen.Left(120)
	ramo(pen, quarto)

	pen.Right(60)
	pen.Walk(-quarto)
	pen.Up()
	pen.Walk(quarto * 2)
	pen.Right(60)
	pen.Down()
	ramo(pen, quarto)
	pen.Left(120)
	ramo(pen, quarto)

	pen.Right(60)
	pen.Walk(-quarto * 2)
	pen.Up()
	pen.Walk(quarto * 3)
	pen.Right(60)
	pen.Down()
	ramo(pen, quarto)
	pen.Left(120)
	ramo(pen, quarto)

	pen.Right(60)
	pen.Walk(-quarto * 3)
	pen.Up()

	pen.Walk(quarto * 4)
	pen.Right(60)
	pen.Down()
	ramo(pen, quarto)
	pen.Left(120)
	ramo(pen, quarto)

	pen.Right(60)
	pen.Walk(-quarto * 4)

}

func main() {
	pen := NewPen(500, 500)
	//gelo(pen, 200)
	//circulo(pen, 200)
	//arvore(pen, 100)
	//quadrado(pen, 50)
	//espiral(pen, 300)
	//carpete(pen, 100)
	//triangulo(pen, 300)
	pen.Up()
	pen.SetHeading(90)
	pen.SetPosition(250, 500)
	pen.Down()
	ramo(pen, 300)
	pen.SavePNG("tree.png")
	fmt.Printf("salvou")
}
