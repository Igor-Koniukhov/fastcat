package main

import (
	"fmt"
	"github.com/igor-koniukhov/fastcat/internal/parser"
	"time"
)

func RunUpToDateSuppliersInfo(t time.Duration) {
	rest := parser.NewRestMenuRepository(&app)
	parser.NewRestMenu(rest)
	suppliersInfo := parser.ParserRestMenu.GetListSuppliers()
	app.ChanelSupplierId = make(chan int, len(suppliersInfo.Restaurants))
	app.ChanelLockUnlock = make(chan int, 1)

	for {
		parser.ParserRestMenu.ParsedDataWriter()
		fmt.Println("Menu is up-to-date ")
		time.Sleep(time.Second * t)
	}
}
