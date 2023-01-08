package main

import (
	"stash-scrapers/app"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// p := &app.Performers{
	// 	ID:   15,
	// 	Name: "加護範子",
	// }
	app.MinnanoRun()
}
