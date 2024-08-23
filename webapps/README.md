# Document-Representor
Represent, interact with, and serve AI-generated web documents for reports, UI or otherwise in Golang (Or in any other language because the style and approach is general)

Document Representor is a way to represent web documents as a structured format, where the main document has different parts like styles, scripts, and body content. These parts are stored in smaller structures that together make up the whole document.

The main document structure can be interacted with and modified, with functions to change styles, scripts, and body content. This structure is then converted into a JSON string or marshalled in Go and sent to an AI model to generate a file that matches the document structure.

WebSockets are used to allow real-time interactions with the document, especially for handling scripts. This approach can be used for various applications, like creating a report to display data or using a web interface for user input or information display in a GUI-deprived program.
