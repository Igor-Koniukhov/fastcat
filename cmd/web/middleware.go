package main

import (
	"fmt"
	"net/http"
)



func  WriteToPage(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		fmt.Println("Hit the page\n")
		next.ServeHTTP(w, r)
	})
}

func  WriteTo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		fmt.Println("Hit the page2\n")
		next.ServeHTTP(w, r)
	})
}






