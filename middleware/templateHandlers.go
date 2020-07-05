package middleware

import (
	"bytes"
	"fmt"
	"html/template"
	"os"

	c "github.com/damilarelana/goCYOA/core"
)

// ConvertTemplateToString ...
func ConvertTemplateToString(t *template.Template) string {
	var temp bytes.Buffer
	var emptyChapterStruct c.Chapter
	err := t.Execute(&temp, emptyChapterStruct)
	if err != nil {
		errMsgHandler(fmt.Sprintf("Failed to render goHTML template to a string %s\n", err.Error()))
	}
	return temp.String()
}

// ParseTemplate ...
func ParseTemplate(templateFilename *string) *template.Template {
	t, err := template.ParseFiles(*templateFilename)
	if err != nil {
		errMsgHandler(fmt.Sprintf("Failed to parse goHTML file %s\n", err.Error()))
	}
	return t
}

// RenderToStdout ...
func RenderToStdout(t *template.Template, story c.Story) {
	for _, s := range story {
		err := t.Execute(os.Stdout, s)
		if err != nil {
			errMsgHandler(fmt.Sprintf("Failed to render goHTML file %s\n", err.Error()))
		}
	}
}

// InitTemplateForWeb ...
func InitTemplateForWeb(templateAsString string) *template.Template {
	t := template.Must(template.New("").Parse(templateAsString))
	return t
}
