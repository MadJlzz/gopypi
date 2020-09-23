package template

import (
	"html/template"
	"io"
)

var t = template.New("index")

// Generate writes the output of a template into a io.Writer.
//
// If an error occurs during template parsing or execution, it will return an error.
// Otherwise, this function will return nil.
func Generate(w io.Writer, filepath string) error {
	tmpl, err := t.ParseFiles(filepath)
	if err != nil {
		return err
	}
	if err = tmpl.ExecuteTemplate(w, "index", nil); err != nil {
		return err
	}
	return nil
}
