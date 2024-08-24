package main

import (
	Representor "dynaweb/Representor"
	"os"
)

func main() {
	document := Representor.DefaultDocument()
	document.Head.Title = "Document Representor!"
	document.Head.Description = "AI Generated html file! let's gooo"
	document.Head.Author = "WEBAPPS"
	document.Head.Keywords = "ai,generated,webapps"
	os.WriteFile("index.html", []byte(document.AIGenerateMarkup()), 0775)
	// document.ExportMarkup("index.html")
}
