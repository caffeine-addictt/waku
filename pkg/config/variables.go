package config

import (
	"bufio"
	"bytes"
	"context"
	"strings"
	"time"

	"github.com/caffeine-addictt/waku/cmd/cleanup"
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
		Type      TemplateVarType        `json:"type" yaml:"type"`
		Format    types.PermissiveString `json:"format" yaml:"format"`
	}
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
		panic("unexpected prompt type while setting value")
	}

	return nil
}

func (t *TemplateVariable) Validate() error {
	if strings.TrimSpace(t.Format.String()) == "" {
		log.Warnf("%s\n", errors.
			NewWakuErrorf("format value is empty").
			WithMeta("variable", string(t.Key)))
	}

	if t.Separator == nil {
		d := string(DEFAULT_SEPARATOR_CHAR)
		t.Separator = &d
	}

	if err := t.Type.Validate(); err != nil {
		return err
	}

	return nil
}

func (ts *TemplateVariables) Validate() error {
	for _, s := range *ts {
		if err := s.Validate(); err != nil {
			return err
		}
	}
	return nil
}
