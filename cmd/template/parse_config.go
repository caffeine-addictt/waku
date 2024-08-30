package template

import (
	"bufio"
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/caffeine-addictt/template/cmd/config"
)

func ParseConfig(filePath string) (*config.TemplateJson, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	var template config.TemplateJson
	var jsonData string

	// Read the entire file content
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		jsonData += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// Unmarshal JSON data
	if err := json.Unmarshal([]byte(jsonData), &template); err != nil {
		return nil, err
	}

	if err := template.Validate(filepath.Dir(filePath)); err != nil {
		return nil, err
	}

	return &template, nil
}
