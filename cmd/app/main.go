package main

import (
	"learn-fiber/config"
	"learn-fiber/internal/app"
	"log"

	"github.com/shopspring/decimal"
)

func main() {
	err := config.Init()
	if err != nil {
		log.Fatalf("config error: %s", err)
	}
	decimal.MarshalJSONWithoutQuotes = true
	app.Run()
}
