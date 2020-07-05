package middleware

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"reflect"
	"strings"

	c "github.com/damilarelana/goCYOA/core"
)

// Handler struct for the story rendering via CustomMux()
type handler struct {
	s *c.Story
	t *template.Template
}

// HandlerOptions to implement functional option changes
type HandlerOptions func(h *handler) {

	func 
	t          *template.Template
	pathParser func(r *http.Request) string // handles the parsing of which URL path to return
}

pathParser (r *http.Request) string {
	// dynamically change chapters via url path
	path := strings.TrimSpace(r.URL.Path) // extract url path
	if path == "" || path == "/" {        // ensures that root path always starts at the first chapter
		path = "/intro"
	}
	path = path[1:]
	return path
}
// CustomHandler ...
func CustomHandler(s *c.Story, templateAsString string, opts ...HandlerOptions) http.Handler { // returns an interface
	return handler{
		s,
		InitTemplateForWeb(templateAsString),
	}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) { // ensure that it implements that http.Handler interface
	path := pathParser(r) // extract the path
	chapter, ok := (*h.s)[path]
	if ok {
		err := h.t.Execute(w, chapter)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Handler method, failed to render 'Chapter'\n", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "'Chapter' data not found", http.StatusNotFound)
}

// urlShortenerHomepage handler
func storyHomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		custom404PageHandler(w, r, http.StatusNotFound)
		return
	}
	dataHomePage := "CYOA homepage"
	io.WriteString(w, dataHomePage)
}

// custom404PageHandler defines custom 404 page
func custom404PageHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.Header().Set("Content-Type", "text/html") // set the content header type
	w.WriteHeader(status)                       // this automatically generates a 404 status code
	if reflect.DeepEqual(status, http.StatusNotFound) {
		data404Page := "This page does not exist ... 404!" // custom error message content
		io.WriteString(w, data404Page)
	}
}
