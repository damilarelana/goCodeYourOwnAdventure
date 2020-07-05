package middleware

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"reflect"

	c "github.com/damilarelana/goCYOA/core"
)

// Handler struct for the story rendering via CustomMux()
type handler struct {
	s *c.Story
	t *template.Template
}

// CustomHandler ...
func CustomHandler(s *c.Story, templateAsString string) http.Handler { // returns an interface
	return handler{
		s,
		InitTemplateForWeb(templateAsString),
	}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) { // ensure that it implements that http.Handler interface
	err := h.t.Execute(w, (*h.s)["intro"])
	if err != nil {
		errMsgHandler(fmt.Sprintf("CustomerHandler's method `handler` failed to render to webserver %s\n", err.Error()))
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
