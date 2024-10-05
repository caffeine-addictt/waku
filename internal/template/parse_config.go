package template

import (
	"bufio"
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/caffeine-addictt/waku/internal/log"
	"github.com/caffeine-addictt/waku/pkg/config"
)

func ParseConfig(filePath string) (*config.TemplateJson, error) {
	file, err := os.Open(filepath.Clean(filePath))
	if err != nil {
		return nil, err
	}
	defer file.Close()

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
	log.Debugln("Unmarshalling JSON data from " + filePath)
	if err := json.Unmarshal([]byte(jsonData), &template); err != nil {
		return nil, err
	}

	log.Debugf("Unmarshalled JSON data: %+v\n", template)
	log.Infoln("Validating JSON data from " + filePath)
	if err := template.Validate(filepath.Dir(filePath)); err != nil {
		return nil, err
	}

	return &template, nil
}
