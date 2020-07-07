package main

import "fmt"

type HTMLDocument struct {
	head []string
	body []string
}

func NewHTMLDocument() *HTMLDocument {
	html := HTMLDocument{
		head: []string{"<link rel=\"stylesheet\" href=\"https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css\" integrity=\"sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T\" crossorigin=\"anonymous\">"},
		body: []string{},
	}
	return &html
}

func (html *HTMLDocument) AddHead(entity string) {
	html.head = append(html.head, entity)
}

func (html *HTMLDocument) AddBody(entity string) {
	html.body = append(html.body, entity)
}

func (html *HTMLDocument) String() string {
	str := "<!doctype html>\n"
	str += "<html lang=\"en\">\n"

	str += " <head>\n"
	str += "  <meta charset=\"utf-8\">\n"
	str += "  <meta name=\"viewport\" content=\"width=device-width, initial-scale=1, shrink-to-fit=no\">\n"
	for _, h := range html.head {
		str += fmt.Sprintf("  %s\n", h)
	}
	str += " </head>\n"

	str += " <body class=\"bg-light\">\n"
	str += "  <div class=\"container-fluid\">\n"
	for _, h := range html.body {
		str += fmt.Sprintf("   %s\n", h)
	}
	str += "  </div>\n"
	str += " </body>\n"

	str += "</html>\n"

	return str
}
