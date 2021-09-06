package main

import (
	"fmt"
	"github.com/igor-koniukhov/fastcat/internal/parser"
	"time"
)

func RunUpToDateSuppliersInfo(t time.Duration) {
	rest := parser.NewRestMenuParser(&app)
	parser.NewRestMenu(rest)
	suppliersInfo := parser.ParseRestMenu.GetListSuppliers()
	app.ChanIdSupplier = make(chan int, len(suppliersInfo.Restaurants))
	defer close(app.ChanIdSupplier)
	for {
		parser.ParseRestMenu.ParsedDataWriter()
		fmt.Println("Menu is up-to-date ")
		time.Sleep(time.Second * t)
	}
}
