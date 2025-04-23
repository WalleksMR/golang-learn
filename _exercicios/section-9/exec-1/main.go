package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"

	"github.com/walleksmr/learn/fluxo-controle/section-9/exec-1/models"
)

func main() {
	itens := make([]models.Mercadoria, 0)

	for i := range 10 {
		name := fmt.Sprintf("produt %v", i)
		price := float64(i + 1*rand.Intn(10))
		mercado := "MW UP"
		itens = append(itens, *models.Create(name, price, mercado))
	}

	jsonData, err := json.MarshalIndent(itens, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonData))
}
