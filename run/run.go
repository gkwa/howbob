package run

import (
	"fmt"
	"os"
	"text/template"

	kcl "kcl-lang.io/kcl-go"
)

func Run(manifestPath, outputPath string) {
	result, err := kcl.Run(manifestPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	r, err := result.First().ToMap()
	if err != nil {
		panic(err)
	}

	packages := r["packages"].([]interface{})
	brewTemplate := `
{{- range $pkg := . }}brew "{{ $pkg.name }}"
{{ end }}`
	tmpl, err := template.New("brew").Parse(brewTemplate)
	if err != nil {
		fmt.Println(err)
		return
	}

	file, err := os.Create(outputPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	err = tmpl.Execute(file, packages)
	if err != nil {
		fmt.Println(err)
		return
	}
}
