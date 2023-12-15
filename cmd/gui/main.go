package main

import (
	"just-converter/gui"
	"just-converter/storage"
)

func main() {
	appData := storage.NewAppData()
	gui.GetUI(appData).Run(true)
}
