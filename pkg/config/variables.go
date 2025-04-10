package config

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/caffeine-addictt/waku/cmd/cleanup"
	"github.com/caffeine-addictt/waku/internal/config"
	"github.com/caffeine-addictt/waku/internal/errors"
	"github.com/caffeine-addictt/waku/internal/types"
	"github.com/caffeine-addictt/waku/internal/utils"
	"github.com/caffeine-addictt/waku/pkg/log"
)

type (
	TemplateVariables []TemplateVariable

	TemplateVariable struct {
		value     any
		Separator *string                `json:"sep,omitempty" yaml:"sep,omitempty"`
		Key       types.CleanString      `json:"key" yaml:"key"`
		Format    types.PermissiveString `json:"fmt" yaml:"fmt"`
		Type      TemplateVarType        `json:"type" yaml:"type"`
	}
	mockTemplateVariable TemplateVariable
)

func (t *TemplateVariable) Value() any {
	return t.value
}

func (t *TemplateVariable) Set(d map[string]any) error {
	src := bufio.NewScanner(strings.NewReader(t.Format.String()))
	var output bytes.Buffer
	writer := bufio.NewWriter(&output)

	errChan := make(chan error, 1)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	cleanup.Schedule(func() error {
		cancel()
		return nil
	})

	go func() {
		if err := utils.ParseTemplateFile(ctx, d, src, writer); err != nil {
			errChan <- err
			return
		}
		errChan <- writer.Flush()
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errChan:
		if err != nil {
			return err
		}
	}

	// handle val
	switch t.Type {
	case TemplateVarTypeString:
		t.value = output.String()
	case TemplateVarTypeArray:
		t.value = strings.Split(output.String(), *t.Separator)
	default:
		panic(fmt.Sprintf("unexpected variable type while setting value: %s", t.Type))
	}

	return nil
}

func (t *TemplateVariable) unmarshal(cfg config.ConfigType, data []byte) error {
	var tv mockTemplateVariable
	if err := cfg.Unmarshal(data, &tv); err != nil {
		return err
	}

	if strings.TrimSpace(tv.Format.String()) == "" {
		log.Warnf("%s\n", errors.
			NewWakuErrorf("format value is empty").
			WithMeta("variable", string(tv.Key)))
	}

	if tv.Separator == nil {
		d := string(DEFAULT_SEPARATOR_CHAR)
		tv.Separator = &d
	}

	switch tv.Type {
	case TemplateVarTypeString, TemplateVarTypeArray:
	case "":
		tv.Type = TemplateVarTypeString
	default:
		return fmt.Errorf("%s is not a valid prompt type", tv.Type)
	}

	*t = TemplateVariable(tv)
	return nil
}

func (t *TemplateVariable) UnmarshalJSON(data []byte) error {
	return t.unmarshal(config.JsonConfig{}, data)
}

func (t *TemplateVariable) UnmarshalYAML(data []byte) error {
	return t.unmarshal(config.YamlConfig{}, data)
}
