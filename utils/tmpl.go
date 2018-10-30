package utils

import (
	"bytes"
	"text/template"
)

func ParseTmpl(tmpl string, data map[string]interface{}) ([]byte, error) {
	parsedTmpl, err := template.New("tmpl").Parse(tmpl)
	if err != nil {
		return nil, err
	}

	var result bytes.Buffer
	if err := parsedTmpl.Execute(&result, data); err != nil {
		return nil, err
	}

	return result.Bytes(), nil
}
