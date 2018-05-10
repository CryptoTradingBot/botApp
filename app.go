package main

import (
	"github.com/CryptoTradingBot/exchanges/bitmex"
	"github.com/CryptoTradingBot/exchanges/config"
	"github.com/CryptoTradingBot/botApp/charts"
)

const (
	// fields for test verification bitmex methods
	Key       = "yHCVhvImDYXCyDVJCiJPjUDT"
	KeySecret = "b6LvJGLI3-cZJ5geVhRvccdE_w9y94opWVzaX5mGPReRJHMK"
	SYMBOL    = "XBTUSD"
)

type Bot struct {
	exchange   bitmex.Bitmex
}

var bot Bot

var exchangeConfig config.ExchangeConfig

func main() {
	// set minimal settings
	exchangeConfig := config.ExchangeConfig{
		Enabled:true,
		Verbose:true,
		AuthenticatedAPISupport:true,
		APIKey: Key,
		APISecret: KeySecret,
		UseSandbox: true,
	}

	bot.exchange.SetDefaults()
	// setup
	bot.exchange.Setup(exchangeConfig)


	var b = bot.exchange

	charts.GetCharts(SYMBOL, "5m", 100, b)





}
