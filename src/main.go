package main

import (
	"os"
	menu "src/Controller"
	view "src/ViewWeb"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "web" {
		view.Viewpage()
	} else {
		menu.DisplayMenu()
	}
}
