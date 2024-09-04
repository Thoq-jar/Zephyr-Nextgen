package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func banner() {
	banner := `
+----------------------------------+
+ ▗▄▄▄▄▖▗▄▄▄▖▗▄▄▖ ▗▖ ▗▖▗▖  ▗▖▗▄▄▖  +
+    ▗▞▘▐▌   ▐▌ ▐▌▐▌ ▐▌ ▝▚▞▘ ▐▌ ▐▌ +
+  ▗▞▘  ▐▛▀▀▘▐▛▀▘ ▐▛▀▜▌  ▐▌  ▐▛▀▚▖ +
+ ▐▙▄▄▄▖▐▙▄▄▖▐▌   ▐▌ ▐▌  ▐▌  ▐▌ ▐▌ +
+----------------------------------+

You are using: Zephyr Editor.

`
	fmt.Print(banner)
}

func main() {
	banner()
	setupAndRun()
}

func setupAndRun() {
	App := app.New()
	Zephyr := App.NewWindow("Zephyr")

	editor := setupEditor()
	menu := setupMenu(Zephyr, editor)

	Zephyr.SetMainMenu(menu)
	content := container.NewBorder(nil, nil, nil, nil, editor)
	Zephyr.SetContent(content)
	Zephyr.Resize(fyne.NewSize(1600, 900))
	Zephyr.ShowAndRun()
}
