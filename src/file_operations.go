package main

import (
	"io"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

func openFile(Zephyr fyne.Window, editor *widget.Entry) func() {
	return func() {
		openDialog := dialog.NewFileOpen(
			func(reader fyne.URIReadCloser, err error) {
				if err != nil {
					dialog.ShowError(err, Zephyr)
					return
				}
				if reader != nil {
					data, err := io.ReadAll(reader)
					if err != nil {
						dialog.ShowError(err, Zephyr)
						return
					}
					editor.SetText(string(data))
				}
			}, Zephyr)
		openDialog.SetFilter(storage.NewExtensionFileFilter([]string{".go"}))
		openDialog.Show()
	}
}

func saveFile(Zephyr fyne.Window, editor *widget.Entry) func() {
	return func() {
		saveDialog := dialog.NewFileSave(
			func(writer fyne.URIWriteCloser, err error) {
				if err != nil {
					dialog.ShowError(err, Zephyr)
					return
				}
				if writer != nil {
					_, err := writer.Write([]byte(editor.Text))
					if err != nil {
						dialog.ShowError(err, Zephyr)
						return
					}
					err = writer.Close()
					if err != nil {
						dialog.ShowError(err, Zephyr)
					}
				}
			}, Zephyr)
		saveDialog.SetFileName("foo.go")
		saveDialog.Show()
	}
}
