package gui

import (
	"gioui.org/app"
	"gioui.org/widget/material"
	"github.com/justsoftwares/justui"
	"just-converter/storage"
)

func GetUI(appData *storage.AppData) *justui.UI {
	if storage.GUI == nil {
		w := app.NewWindow(
			app.Title("Just Converter"),
			app.MinSize(800, 600),
			app.MaxSize(800, 600),
			app.Size(1, 1),
		)
		t := material.NewTheme()
		storage.GUI = justui.NewUI(w, t, draw)
	}
	storage.Data = appData
	return storage.GUI
}
