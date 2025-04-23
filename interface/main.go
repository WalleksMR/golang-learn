package main

import (
	"fmt"
	"math"
)

type geometria interface {
	area() float64
}

func ExibirGeometria(g geometria) {
	fmt.Printf("Tipo: %T | Área: %.2f\n", g, g.area())
}

type circulo struct {
	radius float64
}

func (c *circulo) area() float64 {
	return math.Pi * c.radius * c.radius
}

type retangulo struct {
	largura, altura float64
}

func (r *retangulo) area() float64 {
	return r.altura * r.largura
}

func main() {
	circulo := &circulo{radius: 3}
	retangulo := &retangulo{largura: 1, altura: 2}

	ExibirGeometria(circulo)
	ExibirGeometria(retangulo)
}
