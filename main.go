package main

import (
    "embed"
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
    "text/template"
)

//go:embed templates/*
var templateFS embed.FS

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: create-go-app <project-name>")
        return
    }
    appName := os.Args[1]
    projectPath := filepath.Join(".", appName)

    dirs := []string{
        "src", "utils", "logger",
    }

    for _, dir := range dirs {
        os.MkdirAll(filepath.Join(projectPath, dir), 0755)
    }

    files := map[string]string{
        "templates/main.go.tmpl":       "src/main.go",
        "templates/helper.go.tmpl":     "utils/helper.go",
        "templates/logger.go.tmpl":     "logger/logger.go",
        "templates/.gitignore.tmpl":    ".gitignore",
        "templates/.env.tmpl":          ".env",
        "templates/.env.example.tmpl":  ".env.example",
        "templates/README.md.tmpl":     "README.md",
    }

    for tmplFile, outFile := range files {
        tmplBytes, _ := templateFS.ReadFile(tmplFile)
        t := template.Must(template.New("").Parse(string(tmplBytes)))
        outPath := filepath.Join(projectPath, outFile)
        outF, _ := os.Create(outPath)
        t.Execute(outF, map[string]string{"AppName": appName})
        outF.Close()
    }

    // Initialize go.mod
    cmd := exec.Command("go", "mod", "init", appName)
    cmd.Dir = projectPath
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Run()

    // Initialize git
    gitInit := exec.Command("git", "init")
    gitInit.Dir = projectPath
    gitInit.Stdout = os.Stdout
    gitInit.Stderr = os.Stderr
    gitInit.Run()

    fmt.Println("Project", appName, "created and initialized with git successfully!")
}
