package run

import (
	"fmt"
	"os"
	"sort"
	"text/template"

	kcl "kcl-lang.io/kcl-go"
)

func Run(manifestPath, brewfilePath, checkerPath string) {
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

	sort.Slice(packages, func(i, j int) bool {
		pi := packages[i].(map[string]interface{})["check_installed"].(string)
		pj := packages[j].(map[string]interface{})["check_installed"].(string)
		return pi < pj
	})

	brewTemplate := `
{{- range $pkg := . }}brew "{{ $pkg.name }}"
{{ end }}`

	tmpl, err := template.New("brew").Parse(brewTemplate)
	if err != nil {
		fmt.Println(err)
		return
	}

	brewfile, err := os.Create(brewfilePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer brewfile.Close()

	err = tmpl.Execute(brewfile, packages)
	if err != nil {
		fmt.Println(err)
		return
	}

	checkerTemplate := `#!/usr/bin/env bash
set -x
{{ range $pkg := . }}{{ $pkg.check_installed }}
{{ end }}`

	tmpl, err = template.New("checker").Parse(checkerTemplate)
	if err != nil {
		fmt.Println(err)
		return
	}

	checker, err := os.Create(checkerPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer checker.Close()

	var filteredPackages []interface{}
	for _, pkg := range packages {
		if p, ok := pkg.(map[string]interface{}); ok && p["check_installed"] != "" {
			filteredPackages = append(filteredPackages, pkg)
		}
	}
	err = tmpl.Execute(checker, filteredPackages)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = os.Chmod(checkerPath, 0o755)
	if err != nil {
		fmt.Println(err)
		return
	}
}
