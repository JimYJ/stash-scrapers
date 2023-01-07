package main

import (
	"stash-scrapers/app"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	p := &app.Performers{
		ID:   9,
		Name: "古館びわ",
	}
	app.MinnanoRun(p)
}
