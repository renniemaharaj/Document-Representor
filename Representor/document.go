package Representor

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

// HtmlDocument struct represents an entire HTML document.
type HtmlDocument struct {
	Language string
	Head     HtmlHead
	Body     Body
}

// Body represents the body of a document.
type Body struct {
	Elements []Element
}

// This function appends an elemement to a document body.
func (body *Body) AppendChild(element Element) {
	body.Elements = append(body.Elements, element)
}

// This function creates and returns a blank document skeleton.
func BlankDocument() HtmlDocument {
	return HtmlDocument{
		Head: HtmlHead{
			Title:   "",
			Metas:   Metas{},
			Links:   Links{},
			Scripts: Scripts{},
		},
		Body: Body{},
	}
}

// This function generates markup for an HtmlDocument and exports it to the file specified. Export as .html.
func (doc HtmlDocument) ExportMarkup(filename string) {
	html := doc.GenerateMarkup()
	fileName := filename
	err := os.WriteFile(fileName, []byte(html), 0644)
	if err != nil {
		log.Fatalf("Failed to write file: %v", err)
	}
}

// This function will run an API call on the HtmlDocument and return its markup representation
func (doc HtmlDocument) AIGenerateMarkup() string {
	// Path to the Python script
	const mScript = "m-response.py"

	// Convert struct to JSON
	jsonData, err := json.Marshal(doc)
	if err != nil {
		log.Fatalf("Failed to marshal struct: %s", err)
	}

	// Convert JSON to string
	jsonString := string(jsonData)

	if !FileExists(mScript) {
		log.Fatalf("Error: API script missing....")
	}

	// Command to run the Python script
	cmd := exec.Command("py", mScript, "\""+jsonString+"\"")

	// Get the output from the command
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Failed to execute command: %s", err)
	}

	// Return the output
	return string(output)
}

// This function generates markup for an HtmlDocument and returns it as a string.
func (doc HtmlDocument) GenerateMarkup() string {
	var sb strings.Builder

	// Add comments at the top of the document.
	sb.WriteString("<!--")
	sb.WriteString("Credit newrennie/webapps")
	sb.WriteString("-->\n")

	//Begin html markup generation.
	sb.WriteString("<!DOCTYPE html>\n")

	//Insert specified language else insert en.
	if doc.Language != "" {
		sb.WriteString(fmt.Sprintf("<html lang=\"%v\">\n", doc.Language))
	} else {
		sb.WriteString("<html lang=\"en\">\n")
	}

	//Let the document head generate itself.
	sb.WriteString(doc.Head.GenerateMarkup())

	//Begin document body
	sb.WriteString("<body>\n")

	//Let each element return their own markup generation
	for _, elem := range doc.Body.Elements {
		sb.WriteString(elem.GenerateMarkup() + "\n")
	}

	//End document body with closing tag
	sb.WriteString("</body>\n")

	//End html markup generation with closing tag
	sb.WriteString("</html>")

	return sb.String()
}
