package main

import (
	"fyne.io/fyne/v2/widget"
)

func setupEditor() *widget.Entry {
	editor := widget.NewMultiLineEntry()
	editor.SetPlaceHolder("Enter your code here...")
	return editor
}
