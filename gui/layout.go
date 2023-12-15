package gui

import (
	"gioui.org/layout"
	"image"
	"just-converter/gui/add-files-form"
	"just-converter/gui/files-lists"
	"just-converter/gui/output-form"
)

func draw(gtx layout.Context) {
	rootLayer := layout.Flex{
		Axis: layout.Vertical,
	}
	rootLayer.Layout(gtx, layout.Rigid(topWidget), layout.Rigid(bottomWidget))
}

func topWidget(gtx layout.Context) layout.Dimensions {
	rootLayer := layout.Flex{
		Axis: layout.Horizontal,
	}
	return rootLayer.Layout(gtx, layout.Rigid(leftWidget), layout.Rigid(rightWidget))
}

func bottomWidget(gtx layout.Context) layout.Dimensions {
	rootLayer := layout.Flex{
		Axis: layout.Horizontal,
	}
	return rootLayer.Layout(gtx, layout.Rigid(output_form.Widget))
}

func leftWidget(gtx layout.Context) layout.Dimensions {
	gtx.Constraints.Min = image.Pt(300, 400)
	gtx.Constraints.Max = gtx.Constraints.Min
	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(gtx, layout.Rigid(files_lists.Widget))
}

func rightWidget(gtx layout.Context) layout.Dimensions {
	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(gtx, layout.Rigid(add_files_form.Widget))
}
