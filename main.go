package main

import (
	"stash-scrapers/app"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// p := &app.Performers{
	// 	ID:   9,
	// 	Name: "古館びわ",
	// }
	// app.SingleTest(p)
	// app.FixAvatar(9, "H:\\download-01\\390420.jpg")
	app.MinnanoRun()
	// app.MinnanoRunAvatar()
}
