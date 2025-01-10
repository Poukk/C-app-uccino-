package template

import (
	"embed"
	"os"
	"path/filepath"
	"text/template"
)

//go:embed templates/*
var templateFiles embed.FS

type Data struct {
	ProjectName string
}

func GenerateFile(templateName, destPath string, data Data) error {
	tmplContent, err := templateFiles.ReadFile(filepath.Join("templates", templateName))
	if err != nil {
		return err
	}

	tmpl, err := template.New(templateName).Parse(string(tmplContent))
	if err != nil {
		return err
	}

	f, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer f.Close()

	return tmpl.Execute(f, data)
}
