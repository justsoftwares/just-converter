package add_files_form

import (
	"gioui.org/layout"
	"gioui.org/widget"
)

var data = struct {
	SubmitClickable, SelectFilesClickable, EditorToFilesClickable, ClearClickable *widget.Clickable
	FilesListData                                                                 []string
	FilesList                                                                     *widget.List
	FilesListEditor, StreamSpecifiersEditor, DecodingEditor, EncodingEditor       *widget.Editor
}{
	SubmitClickable:        &widget.Clickable{},
	SelectFilesClickable:   &widget.Clickable{},
	EditorToFilesClickable: &widget.Clickable{},
	ClearClickable:         &widget.Clickable{},
	FilesList: &widget.List{List: layout.List{
		Axis: layout.Vertical,
	}},
	FilesListEditor:        &widget.Editor{},
	StreamSpecifiersEditor: &widget.Editor{},
	DecodingEditor:         &widget.Editor{},
	EncodingEditor:         &widget.Editor{},
}
