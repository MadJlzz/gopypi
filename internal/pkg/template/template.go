package template

import (
	"github.com/MadJlzz/gopypi/internal/pkg/utils"
	"html/template"
	"io"
	"path/filepath"
)

var indexPath = filepath.Join(utils.BasePath(), "web", "index.gohtml")

type SimpleRepositoryTemplate struct {
	tmpl *template.Template
}

func New() *SimpleRepositoryTemplate {
	funcMap := template.FuncMap{
		"safe": func(s string) template.HTMLAttr {
			return template.HTMLAttr(s)
		},
	}
	tmpl := template.Must(template.New("root").Funcs(funcMap).ParseGlob(utils.BasePath() + "/*/*.gohtml"))
	//tmpl, err := template.New("index").ParseFiles(indexPath)
	//if err != nil {
	//	return nil, err
	//}
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
