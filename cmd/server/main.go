package main

import (
	"github.com/multibear95/stanovi_monitor/internal"
)

func main() {
	println("Starting listening on port 8080")
	srv := internal.NewHTTPServer(":8080")
	err := srv.ListenAndServe()
	if err != nil {
		return
	}
}
