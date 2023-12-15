package output_form

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
	"github.com/justsoftwares/justui"
	"github.com/justsoftwares/justui/widgets/form"
	"just-converter/storage"
)

func Widget(gtx layout.Context) layout.Dimensions {
	storage.GUI.AddFrameEventHandlers(justui.EventHandler{
		Event:   data.SubmitClickable.Clicked,
		Handler: SubmitClickableClicked,
	})
	f := form.NewForm(
		storage.GUI.Theme,
		"Output options",
		nil,
		material.Button(storage.GUI.Theme, data.SubmitClickable, "Convert").Layout,
	)
	f.AddField("Folder", material.Editor(storage.GUI.Theme, data.FolderEditor, "C:\\output-folder\\").Layout)
	f.AddField("Encoding parameters", material.Editor(storage.GUI.Theme, data.EncodingEditor, "-c copy").Layout)
	return f.Layout(gtx)
}
