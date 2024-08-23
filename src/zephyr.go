package main

import (
	"fyne.io/fyne/v2/widget"
)

func setupEditor() *widget.Entry {
	editor := widget.NewMultiLineEntry()
	editor.SetPlaceHolder("Begin typing...")

	return editor
}
