package processor

import "fmt"

func ShowPriceAvg(priceChannel <-chan float64, done chan<- bool) {
	var totalPrice float64
	countPrices := 0.0
	for price := range priceChannel {
		totalPrice += price
		countPrices++
		avgPrices := totalPrice / countPrices
		fmt.Printf("Preço recebido: R$ %.2f | Preço médio até agora: R$ %.2f\n", price, avgPrices)
	}

	done <- true
}
