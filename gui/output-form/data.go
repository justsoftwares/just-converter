package output_form

import "gioui.org/widget"

var data = struct {
	SubmitClickable              *widget.Clickable
	FolderEditor, EncodingEditor *widget.Editor
}{
	SubmitClickable: &widget.Clickable{},
	FolderEditor:    &widget.Editor{},
	EncodingEditor:  &widget.Editor{},
}
