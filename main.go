// Example taken from
//  https://tutorialedge.net/post/golang/creating-simple-web-server-with-golang/
package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
)

func main() {
	fmt.Println("Starting Webserver on Port 8081")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		doc := NewHTMLDocument()
		doc.AddHead("<title>Response Test</title>")
		doc.AddBody("<h1>Response Test</h1>")
		name, _ := os.Hostname()
		doc.AddBody(
			fmt.Sprintf(
				"<p>Running on host <code>%s</code> and routed in as <code>%s</code>. "+
					"Access method is <code>%s</code> from client <code>%s</code>.</p>",
				name,
				html.EscapeString(r.URL.Path),
				r.Method,
				r.RemoteAddr,
			),
		)

		r.ParseForm()

		if len(r.Header) > 0 {

			doc.AddBody("<h2>Headers</h2>")
			strs := []string{}
			for k, v := range r.Header {
				strs = append(strs, fmt.Sprintf("%s=%s", k, v))
			}
			sort.Strings(strs)
			doc.AddBody(fmt.Sprintf("<pre>\n%s\n</pre>", strings.Join(strs, "\n")))
		}

		if len(r.URL.Query()) > 0 {

			doc.AddBody("<h2>Request Parameter</h2>")
			strs := []string{}
			for k, v := range r.URL.Query() {
				if len(v) == 1 {
					strs = append(strs, fmt.Sprintf("%s=%s", k, v[0]))
				} else {
					strs = append(strs, fmt.Sprintf("%s=%s", k, strings.Join(v, ", ")))
				}
			}
			sort.Strings(strs)
			doc.AddBody(fmt.Sprintf("<pre>\n%s\n</pre>", strings.Join(strs, "\n")))
		}

		if len(r.Form) > 0 {
			doc.AddBody("<h2>Post Form</h2>")
			strs := []string{}
			for k, v := range r.Form {
				if len(v) == 1 {
					strs = append(strs, fmt.Sprintf("%s=%s", k, v[0]))
				} else {
					strs = append(strs, fmt.Sprintf("%s=%s", k, strings.Join(v, ", ")))
				}
			}
			sort.Strings(strs)
			doc.AddBody(fmt.Sprintf("<pre>\n%s\n</pre>", strings.Join(strs, "\n")))
		}

		fmt.Fprintf(w, doc.String())
	})

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
