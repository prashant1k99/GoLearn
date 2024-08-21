package main

// `//go:embed` is a compiler directoive taht allows programs to include arbitrary files and folders in the Go binary at build time.

import (
	"embed"
)

// embed directives accept paths relative to the directory containing the Go source file. This directive embeds the cotents of the file into the string variable immediately following it.
//
//go:embed README.md
var fileString string

// or embed the contents of the file into a []byte.
//
//go:embed README.md
var fileByte []byte

// We can also embed multiple files or even folders with wildcards. THis uses a variable of the embed FS type, which implements a simple virtual file syste.
//
//go:embed folder/Learning.md
//go:embed folder/*.md
var folder embed.FS

func main() {
	// Print out the contents of README.md
	print(fileString)
	print(string(fileByte))
	// ### Learning Go by Practice

	// This is the practice ground for Prashant to test his GoLang code.

	// Retrieve some files from the embedded folder.
	content1, _ := folder.ReadFile("folder/Learning.md")
	print(string(content1))

	content2, _ := folder.ReadFile("folder/Learning2.md")
	print(string(content2))
}
