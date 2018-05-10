package charts

import (
	"log"
	"github.com/CryptoTradingBot/exchanges/models"
	"github.com/CryptoTradingBot/exchanges/bitmex"
)

// получение графика валют, @symbol {string} валютная пара
func GetCharts(currencyPair string, timeframe string, count int, bitmex bitmex.Bitmex) ([]models.Candle) {
	response, err := bitmex.GetCandles(currencyPair, timeframe, count)

	if err != nil {
		log.Println("error: ", err)
	}
	log.Printf("tradeBin: %+v\n", response)
	return response
}
