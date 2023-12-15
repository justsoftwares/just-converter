package output_form

import (
	"gioui.org/io/event"
	"gioui.org/layout"
	"just-converter/storage"
	"just-converter/transcoder"
)

func SubmitClickableClicked(_ layout.Context, _ event.Event) {
	transcoder.Transcode(storage.Data)
}
