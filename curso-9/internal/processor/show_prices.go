package processor

import (
	"buscador/internal/models"
	"fmt"
)

func ShowPriceAvg(priceChannel <-chan models.PriceDetail, done chan<- bool) {
	var totalPrice float64
	countPrices := 0.0
	for price := range priceChannel {
		totalPrice += price.Value
		countPrices++
		avgPrices := totalPrice / countPrices
		fmt.Printf("[%s] Nome: %s | Preço recebido: R$ %.2f | Preço médio até agora: R$ %.2f\n", price.Timestamp.Format("02/01/2006 15:04:05"), price.StoreName, price.Value, avgPrices)
	}

	done <- true
}
