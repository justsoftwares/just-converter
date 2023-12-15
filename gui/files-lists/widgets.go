package files_lists

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/widget/material"
	"just-converter/storage"
)

func Widget(gtx layout.Context) layout.Dimensions {
	return material.List(storage.GUI.Theme, data).Layout(gtx, len(storage.Data.FilesLists), func(gtx layout.Context, index int) layout.Dimensions {
		elLayout := layout.Flex{
			Axis: layout.Horizontal,
		}

		return elLayout.Layout(
			gtx,
			layout.Rigid(material.Body1(
				storage.GUI.Theme,
				fmt.Sprintf("Files count: %d", len(storage.Data.FilesLists[index].FilesList))).Layout,
			),
			layout.Rigid(material.Body1(
				storage.GUI.Theme,
				storage.Data.FilesLists[index].Encoding).Layout,
			),
		)
	})
}
