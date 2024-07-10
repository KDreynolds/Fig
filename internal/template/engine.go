package template

import (
	"bytes"
	"fmt"
	"text/template"
)

// Engine represents our templating engine
type Engine struct {
	// We could add more fields here if needed in the future
}

// New creates a new templating engine
func New() *Engine {
	return &Engine{}
}

// Render processes a template string with given variables
func (e *Engine) Render(templateStr string, vars map[string]string) (string, error) {
	tmpl, err := template.New("command").Parse(templateStr)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %v", err)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, vars)
	if err != nil {
		return "", fmt.Errorf("failed to execute template: %v", err)
	}

	return buf.String(), nil
}

// RenderTask processes a task's command with its variables
func (e *Engine) RenderTask(command string, taskVars, globalVars map[string]string) (string, error) {
	// Merge global and task-specific variables, with task variables taking precedence
	vars := make(map[string]string)
	for k, v := range globalVars {
		vars[k] = v
	}
	for k, v := range taskVars {
		vars[k] = v
	}

	return e.Render(command, vars)
}
