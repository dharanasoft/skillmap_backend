package main

import (
	"github.com/pocketbase/pocketbase"
)

func main() {
	app := pocketbase.New()
	if err := app.Start(); err != nil {
		panic(err)
	}
}
