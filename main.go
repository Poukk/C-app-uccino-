package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"c-appuccino/internal/template"
	"github.com/charmbracelet/huh"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var name string
	var options []string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Project name:").
				Value(&name),

			huh.NewMultiSelect[string]().
				Options(
					huh.NewOption("Create \"src\" dir", "src").Selected(true),
					huh.NewOption("Create \"include\" dir", "include").Selected(true),
					huh.NewOption("Create Makefile", "makefile"),
					huh.NewOption("Init git repository", "git"),
				).
				Title("Options").
				Value(&options),
		),
	)

	err := form.Run()
	check(err)

	// Create project directory
	projectPath := filepath.Join(".", name)
	err = os.MkdirAll(projectPath, 0755)
	check(err)

	templateData := template.Data{
		ProjectName: name,
	}

	// Process selected options
	for _, opt := range options {
		switch opt {
		case "src":
			srcPath := filepath.Join(projectPath, "src")
			err = os.MkdirAll(srcPath, 0755)
			check(err)
			err = template.GenerateFile("main.c.tmpl", filepath.Join(srcPath, "main.c"), templateData)
			check(err)

		case "include":
			includePath := filepath.Join(projectPath, "include")
			err = os.MkdirAll(includePath, 0755)
			check(err)
			err = template.GenerateFile("example.h.tmpl", filepath.Join(includePath, "example.h"), templateData)
			check(err)

		case "makefile":
			err = template.GenerateFile("Makefile.tmpl", filepath.Join(projectPath, "Makefile"), templateData)
			check(err)

		case "git":
			cmd := exec.Command("git", "init")
			cmd.Dir = projectPath
			err = cmd.Run()
			check(err)

			err = template.GenerateFile("gitignore.tmpl", filepath.Join(projectPath, ".gitignore"), templateData)
			check(err)
		}
	}

	fmt.Printf("\nProject '%s' created successfully!\n", name)
}
