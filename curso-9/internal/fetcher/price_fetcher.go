package fetcher

import (
	"buscador/internal/models"
	"math/rand/v2"
	"sync"
	"time"
)

func FetchPrices(priceChannel chan<- models.PriceDetail) {
	var wg sync.WaitGroup

	wg.Add(4)

	go func() {
		defer wg.Done()
		priceChannel <- FetchPriceFromSite1()
	}()

	go func() {
		defer wg.Done()
		priceChannel <- FetchPriceFromSite2()
	}()

	go func() {
		defer wg.Done()
		priceChannel <- FetchPriceFromSite3()
	}()

	go func() {
		defer wg.Done()
		FetchAndSendMultiplePrices(priceChannel)
	}()

	wg.Wait()
	close(priceChannel)
}

func FetchPriceFromSite1() models.PriceDetail {
	time.Sleep(1 * time.Second)

	return getRandomPriceDetail()
}

func FetchPriceFromSite2() models.PriceDetail {
	time.Sleep(3 * time.Second)

	return getRandomPriceDetail()
}

func FetchPriceFromSite3() models.PriceDetail {
	time.Sleep(5 * time.Second)

	return getRandomPriceDetail()
}

func FetchAndSendMultiplePrices(priceChannel chan<- models.PriceDetail) {
	time.Sleep(6 * time.Second)
	prices := []models.PriceDetail{
		getRandomPriceDetail(),
		getRandomPriceDetail(),
		getRandomPriceDetail(),
	}

	for _, price := range prices {
		priceChannel <- price
	}
}

func getRandomPriceDetail() models.PriceDetail {
	return models.PriceDetail{
		StoreName: string(byte(rand.IntN(25) + 65)),
		Value:     rand.Float64() * 100,
		Timestamp: time.Now(),
	}
}
