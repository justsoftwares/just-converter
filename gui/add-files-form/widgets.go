package add_files_form

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
	"github.com/justsoftwares/justui"
	"github.com/justsoftwares/justui/widgets/form"
	"image"
	"just-converter/storage"
)

func Widget(gtx layout.Context) layout.Dimensions {
	storage.GUI.AddFrameEventHandlers(justui.EventHandler{
		Event:   data.SubmitClickable.Clicked,
		Handler: submitClickableClicked,
	})
	f := form.NewForm(
		storage.GUI.Theme,
		"Add files",
		nil,
		material.Button(storage.GUI.Theme, data.SubmitClickable, "Add list").Layout,
	)
	f.AddField("Files editor", filesEditorWidget)
	f.AddField("Files list", listWidget)
	f.AddField("Stream specifiers", material.Editor(storage.GUI.Theme, data.StreamSpecifiersEditor, "0:a:0").Layout)
	f.AddField("Decoding parameters", material.Editor(storage.GUI.Theme, data.DecodingEditor, "-ss 10").Layout)
	f.AddField("Encoding parameters", material.Editor(storage.GUI.Theme, data.EncodingEditor, "-vf scale=500:-1").Layout)
	return f.Layout(gtx)
}

func listWidget(gtx layout.Context) layout.Dimensions {
	storage.GUI.AddFrameEventHandlers(justui.EventHandler{
		Event:   data.EditorToFilesClickable.Clicked,
		Handler: editorToFilesClickableClicked,
	}, justui.EventHandler{
		Event:   data.SelectFilesClickable.Clicked,
		Handler: selectFilesClickableClicked,
	}, justui.EventHandler{
		Event:   data.ClearClickable.Clicked,
		Handler: ClearClickableClicked,
	})
	gtx.Constraints.Min = image.Pt(400, 200)
	gtx.Constraints.Max = gtx.Constraints.Min
	rootLayout := &layout.Flex{
		Axis: layout.Vertical,
	}
	buttons := func(gtx layout.Context) layout.Dimensions {
		rootLayout := &layout.Flex{
			Axis: layout.Horizontal,
		}
		return rootLayout.Layout(gtx,
			layout.Rigid(material.Button(
				storage.GUI.Theme,
				data.SelectFilesClickable,
				"Select files",
			).Layout),
			layout.Rigid(material.Button(
				storage.GUI.Theme,
				data.EditorToFilesClickable,
				"Add files from editor",
			).Layout),
			layout.Rigid(material.Button(
				storage.GUI.Theme,
				data.ClearClickable,
				"Clear",
			).Layout),
		)
	}
	list := func(gtx layout.Context) layout.Dimensions {
		return material.List(storage.GUI.Theme, data.FilesList).Layout(gtx,
			len(data.FilesListData),
			func(gtx layout.Context, index int) layout.Dimensions {
				return material.Body1(storage.GUI.Theme, data.FilesListData[index]).Layout(gtx)
			})
	}
	return rootLayout.Layout(gtx,
		layout.Rigid(buttons),
		layout.Rigid(list))
}

func filesEditorWidget(gtx layout.Context) layout.Dimensions {
	gtx.Constraints.Min = image.Pt(400, 200)
	gtx.Constraints.Max = gtx.Constraints.Min
	return material.Editor(storage.GUI.Theme, data.FilesListEditor, "C:\\folder\\file.mov").Layout(gtx)
}
