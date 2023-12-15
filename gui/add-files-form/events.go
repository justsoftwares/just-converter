package add_files_form

import (
	"gioui.org/io/event"
	"gioui.org/layout"
	"github.com/justsoftwares/justui/widgets/explorer"
	"just-converter/storage"
	"log"
	"strings"
)

func submitClickableClicked(_ layout.Context, _ event.Event) {
	storage.Data.FilesLists = append(storage.Data.FilesLists, storage.FilesList{
		FilesList:        data.FilesListData,
		StreamSpecifiers: strings.Split(data.StreamSpecifiersEditor.Text(), "\n"),
		Decoding:         data.DecodingEditor.Text(),
		Encoding:         data.EncodingEditor.Text(),
	})
}

func editorToFilesClickableClicked(_ layout.Context, _ event.Event) {
	if data.FilesListEditor.Text() != "" {
		data.FilesListData = append(data.FilesListData, strings.Split(data.FilesListEditor.Text(), "\n")...)
		data.FilesListEditor.SetText("")
	}
}

func selectFilesClickableClicked(_ layout.Context, _ event.Event) {
	var exp *explorer.Explorer
	exp = explorer.NewExplorer(storage.GUI.Theme, func(_ layout.Context, _ event.Event) {
		log.Println(exp.Files)
		if len(exp.Files) > 0 {
			data.FilesListData = append(data.FilesListData, exp.GetSelectedFiles()...)
		}
	})
	exp.Run()
}

func ClearClickableClicked(_ layout.Context, _ event.Event) {
	data.FilesListData = []string{}
}
