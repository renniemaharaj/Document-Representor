package main

import (
	Representor "dynaweb/representor"
)

func main() {
	// println("Beginning Entry")
	document := Representor.DefaultDocument()
	document.Head.Title = "Document Representor!"
	document.Head.Description = "AI Generated html file! let;s gooo"
	document.Head.Author = "WEBAPPS"
	document.Head.Keywords = "ai,generated,webapps"
	// jsonString, err := json.Marshal(document)
	// if err == nil {
	// 	println(string(jsonString))
	// }
	println(document.AIGenerateMarkup())
}
