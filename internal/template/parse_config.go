package template

import (
	"encoding/json"
	"os"
	"io"
	"path/filepath"
	"strings"

	"github.com/caffeine-addictt/waku/internal/log"
	"github.com/caffeine-addictt/waku/pkg/config"
	"gopkg.in/yaml.v3"
)

func ParseConfig(filePath string) (*config.TemplateJson, error) {
	file, err := os.Open(filepath.Clean(filePath))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	log.Debugf("reading config file at %v\n", filePath)
	data, err := io.ReadAll(file)
	if err != nil {
		return "", nil, err
	}

	// Unmarshal JSON data
	var template config.TemplateJson
	if strings.HasSuffix(path, "json") {
		log.Debugln("unmarshalling JSON data from " + path)
		if err := json.Unmarshal(data, &template); err != nil {
			return "", nil, err
		}
	} else {
		log.Debugln("unmarshalling YAML data from " + path)
		if err := yaml.Unmarshal(data, &template); err != nil {
			return "", nil, err
		}
	}

	log.Debugf("Unmarshalled JSON data: %+v\n", template)
	log.Infoln("Validating JSON data from " + filePath)
	if err := template.Validate(filepath.Dir(filePath)); err != nil {
		return nil, err
	}

	return &template, nil
}
