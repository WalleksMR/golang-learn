package models

import (
	"fmt"
	"math/rand"
	"time"
)

type Mercadoria struct {
	Id        string
	Name      string
	Price     float64
	CreatedAt time.Time
	Mercado   Mercado
}

func Create(name string, price float64, mercadoName string) *Mercadoria {
	item := &Mercadoria{Id: fmt.Sprint(rand.Int63()), Name: name, Price: price, CreatedAt: time.Now(), Mercado: Mercado{
		Id: "5517686653725550694", Name: mercadoName,
	}}
	return item
}
