package run

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"text/template"

	kcl "kcl-lang.io/kcl-go"
)

func Brewfile(manifestPath, brewfilePath, checkerPath, tapsPath string) {
	result, err := kcl.Run(manifestPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Errorf("error running KCL: %w", err))
		return
	}

	r, err := result.First().ToMap()
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Errorf("error converting KCL result to map: %w", err))
		return
	}

	packages := r["packages"].([]interface{})

	sort.Slice(packages, func(i, j int) bool {
		pi := packages[i].(map[string]interface{})["check_installed"].(string)
		pj := packages[j].(map[string]interface{})["check_installed"].(string)
		return pi < pj
	})

	taps := make(map[string]bool)
	for _, pkg := range packages {
		pkgMap := pkg.(map[string]interface{})
		name := pkgMap["name"].(string)
		if strings.Contains(name, "/") {
			parts := strings.Split(name, "/")
			if len(parts) >= 2 {
				tap := parts[0] + "/" + parts[1]
				taps[tap] = true
			}
		}
	}

	brewTemplate := `
{{- range $pkg := . }}
{{- if $pkg.disabled }}# {{ end -}}
brew "{{ $pkg.name }}"
{{ end }}`

	tmpl, err := template.New("brew").Parse(brewTemplate)
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Errorf("error parsing brew template: %w", err))
		return
	}

	brewfile, err := os.Create(brewfilePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Errorf("error creating Brewfile: %w", err))
		return
	}
	defer brewfile.Close()

	err = tmpl.Execute(brewfile, packages)
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Errorf("error executing brew template: %w", err))
		return
	}

	tapsFile, err := os.Create(tapsPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Errorf("error creating taps file: %w", err))
		return
	}
	defer tapsFile.Close()

	for tap := range taps {
		_, err = fmt.Fprintf(tapsFile, "brew tap %s\n", tap)
		if err != nil {
			fmt.Fprintln(os.Stderr, fmt.Errorf("error writing to taps file: %w", err))
			return
		}
	}

	checkerTemplate := `#!/usr/bin/env bash
set -x

{{ range $pkg := . }}
{{- if $pkg.disabled }}# {{ end -}}
{{ $pkg.check_installed }}
{{ end }}`

	tmpl, err = template.New("checker").Parse(checkerTemplate)
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Errorf("error parsing checker template: %w", err))
		return
	}

	checker, err := os.Create(checkerPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Errorf("error creating version_checker.sh: %w", err))
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
		fmt.Fprintln(os.Stderr, fmt.Errorf("error executing checker template: %w", err))
		return
	}

	err = os.Chmod(checkerPath, 0o755)
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Errorf("error setting executable permissions on version_checker.sh: %w", err))
		return
	}
}
