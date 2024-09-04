package main

import (
	"io"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

func openFile(Zephyr fyne.Window, editor *widget.Entry) func() {
	fileExtensions := []string{
		".go",        // Go
		".py",        // Python
		".java",      // Java
		".js",        // JavaScript
		".ts",        // TypeScript
		".c",         // C
		".cpp",       // C++
		".cs",        // C#
		".rb",        // Ruby
		".php",       // PHP
		".html",      // HTML
		".css",       // CSS
		".xml",       // XML
		".json",      // JSON
		".yaml",      // YAML
		".md",        // Markdown
		".sql",       // SQL
		".swift",     // Swift
		".kt",        // Kotlin
		".r",         // R
		".pl",        // Perl
		".sh",        // Shell script
		".bash",      // Bash
		".lua",       // Lua
		".dart",      // Dart
		".scala",     // Scala
		".groovy",    // Groovy
		".clj",       // Clojure
		".erl",       // Erlang
		".ex",        // Elixir
		".html.erb",  // HTML with embedded Ruby
		".jsp",       // JavaServer Pages
		".aspx",      // ASP.NET
		".vb",        // Visual Basic
		".f",         // Fortran
		".pas",       // Pascal
		".d",         // D
		".nim",       // Nim
		".h",         // C header
		".hpp",       // C++ header
		".m",         // Objective-C
		".mm",        // Objective-C++
		".coffee",    // CoffeeScript
		".scss",      // SASS
		".less",      // LESS
		".tsv",       // Tab-separated values
		".csv",       // Comma-separated values
		".txt",       // Text file
		".log",       // Log file
		".bat",       // Batch file
		".cmd",       // Command file
		".ps1",       // PowerShell script
		".json5",     // JSON5
		".graphql",   // GraphQL
		".proto",     // Protocol Buffers
		".hbs",       // Handlebars
		".twig",      // Twig
		".pug",       // Pug
		".vue",       // Vue.js
		".svelte",    // Svelte
		".rs",        // Rust
		".asm",       // Assembly
		".v",         // V
		".zsh",       // Z shell script
		".fish",      // Fish shell script
		".raku",      // Raku (formerly Perl 6)
		".d.ts",      // TypeScript declaration file
		".mjs",       // JavaScript module
		".cjs",       // CommonJS module
		".hxml",      // Haxe
		".hx",        // Haxe source file
		".erl",       // Erlang source file
		".exs",       // Elixir script
		".nimble",    // Nimble (Nim package)
		".f90",       // Fortran 90
		".f95",       // Fortran 95
		".vhd",       // VHDL
		".vhdl",      // VHDL
		".sv",        // SystemVerilog
		".tcl",       // Tcl
		".pl6",       // Perl 6
		".rhtml",     // Ruby HTML
		".rmd",       // R Markdown
		".ipynb",     // Jupyter Notebook
		".sml",       // Standard ML
		".ml",        // OCaml
		".dylan",     // Dylan
		".ino",       // Arduino
		".cl",        // Common Lisp
		".lisp",      // Lisp
		".racket",    // Racket
		".pike",      // Pike
		".rexx",      // REXX
		".xsl",       // XSLT
		".xslt",      // XSLT
		".wxs",       // WiX source
		".wxl",       // WiX localization
		".wxi",       // WiX include
		".wixproj",   // WiX project
		".sln",       // Visual Studio solution
		".csproj",    // C# project
		".vbproj",    // Visual Basic project
		".fsproj",    // F# project
		".xcodeproj", // Xcode project
		".gradle",    // Gradle build file
		".gyp",       // GYP build file
		".cmake",     // CMake file
		".makefile",  // Makefile
		".mk",        // Makefile
		".ninja",     // Ninja build file
		".build",     // Build file
	}
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
		openDialog.SetFilter(storage.NewExtensionFileFilter(fileExtensions))
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
