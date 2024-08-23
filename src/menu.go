package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func createMenu(Zephyr fyne.Window) *fyne.Menu {
	return fyne.NewMenu("File",
		fyne.NewMenuItem("Toggle Light/Dark Mode", func() {
			toggleTheme(Zephyr)
		}),
	)
}

func setupMenu(Zephyr fyne.Window, editor *widget.Entry) *fyne.MainMenu {
	return fyne.NewMainMenu(
		createMenu(Zephyr),
		fyne.NewMenu("Edit",
			fyne.NewMenuItem("Open", openFile(Zephyr, editor)),
			fyne.NewMenuItem("Save", saveFile(Zephyr, editor)),
		),
	)
}
