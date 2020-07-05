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

// HandlerOptions is a struct of type `http.Handler` that is `optionally` used to extend (i.e. decorate) and existing function (which returns `http.Handler`
type HandlerOptions func(h *handler)

// CustomHandler ...
func CustomHandler(s *c.Story, templateAsString string, opts ...HandlerOptions) http.Handler { // returns an interface
	h := handler{ // set the default handler [prior to being changed by functional options]
		s,
		InitTemplateForWeb(templateAsString),
	}

	/* apply the functional options by
	 	- iterating over the `Opts`
		- then passing in the `original handler` for it to be `decorated` i.e. functionally extended by `Opts`
		- then initiate this by calling `x.WithTemplate()` somewhere else in the code
	*/
	for _, opt := range opts {
		opt(&h)
	}
	return h
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) { // ensure that it implements that http.Handler interface
	path := h.pathParser(r) // extract the path
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

// pathParse() ...
// method attached to the h handler
func (h handler) pathParser(r *http.Request) (path string) {
	// dynamically change chapters via url path
	path = strings.TrimSpace(r.URL.Path) // extract url path
	if path == "" || path == "/" {       // ensures that root path always starts at the first chapter
		path = "/intro"
	}
	return path[1:]
}

// WithTemplate defines a functional option behaviour when user provides a template
func WithTemplate(t *template.Template) HandlerOptions {
	return func(h *handler) { // use the user defined `t` to adjust `h.t` i.e. the `InitTemplateForWeb`
		h.t = t
	}
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
