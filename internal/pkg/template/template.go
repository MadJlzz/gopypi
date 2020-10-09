package template

import (
	"github.com/MadJlzz/gopypi/internal/pkg/utils"
	"html/template"
	"io"
)

type SimpleRepositoryTemplate struct {
	tmpl *template.Template
}

func New() *SimpleRepositoryTemplate {
	tmpl := template.Must(template.ParseGlob(utils.BasePath() + "/*/*.gohtml"))
	return &SimpleRepositoryTemplate{tmpl: tmpl}
}

// Execute writes the output of a template into a io.Writer.
//
// If an error occurs during execution, it will return an error.
// Otherwise, this function will return nil.
func (srt *SimpleRepositoryTemplate) Execute(w io.Writer, templateName string, data interface{}) error {
	if err := srt.tmpl.ExecuteTemplate(w, templateName, data); err != nil {
		return err
	}
	return nil
}