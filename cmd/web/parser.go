package main

import (
	"fmt"
	"github.com/igor-koniukhov/fastcat/internal/parser"
	"time"
)

func RunUpToDateSuppliersInfo(t time.Duration) {
	rest := parser.NewRestMenuRepository(&app)
	parser.NewRestMenu(rest)
	suppliersInfo := parser.ParseRestMenu.GetListSuppliers()
	app.ChanIdSupplier = make(chan int, len(suppliersInfo.Restaurants))
	defer close(app.ChanIdSupplier)
	app.ChanMutex = make(chan int, 1)
	defer close(app.ChanMutex)

	for {
		parser.ParseRestMenu.ParsedDataWriter()
		fmt.Println("Menu is up-to-date ")
		time.Sleep(time.Second * t)
	}
}
