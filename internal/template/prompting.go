package template

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/caffeine-addictt/waku/cmd/options"
	"github.com/caffeine-addictt/waku/internal/license"
	"github.com/caffeine-addictt/waku/internal/log"
	"github.com/caffeine-addictt/waku/internal/searching"
	"github.com/caffeine-addictt/waku/internal/sorting"
	"github.com/caffeine-addictt/waku/internal/types"
	"github.com/caffeine-addictt/waku/internal/utils"
	"github.com/caffeine-addictt/waku/pkg/config"
	"github.com/charmbracelet/huh"
)

// PromptForStyle prompts user to select style
func PromptForStyle(styles config.TemplateStyles, setK *types.CleanString, setV *config.TemplateStyle) *huh.Select[string] {
	opts := make([]string, 0, len(styles))

	for n := range styles {
		opts = append(opts, strings.ToLower(string(n)))
	}
	sorting.QuicksortASC(opts)

	if options.NewOpts.Style.Value() != "" {
		if err := validateStyle(&styles, opts, options.NewOpts.Style.Value(), setK, setV); err == nil {
			return nil
		}
	}

	return huh.NewSelect[string]().Title("The style to use").Options(huh.NewOptions(opts...)...).Validate(func(s string) error {
		return validateStyle(&styles, opts, s, setK, setV)
	})
}

func validateStyle(ll *config.TemplateStyles, optsL []string, val string, setK *types.CleanString, setV *config.TemplateStyle) error {
	val = strings.ToLower(val)

	i := searching.BinarySearchAuto(optsL, val)
	if i == -1 {
		return fmt.Errorf("unknown style: %s", val)
	}

	for n, v := range *ll {
		if strings.EqualFold(string(n), val) {
			*setV = v
			*setK = n
			break
		}
	}

	return nil
}

// PromptForLicense prompts user to select license
func PromptForLicense(value *license.License) (*huh.Select[string], error) {
	fetchedL, err := license.GetLicenses()
	if err != nil {
		return nil, err
	}

	licenses := make([]string, 0, len(*fetchedL))
	for _, v := range *fetchedL {
		licenses = append(licenses, strings.ToLower(v.Name), strings.ToLower(v.Spdx))
	}
	sorting.QuicksortASC(licenses)

	if options.NewOpts.License.Value() != "" {
		if err := validateLicense(fetchedL, licenses, strings.ToLower(options.NewOpts.License.Value()), value); err == nil {
			return nil, nil
		}
	}

	return huh.NewSelect[string]().Title("Your project license").Options(huh.NewOptions(licenses...)...).Validate(func(s string) error {
		return validateLicense(fetchedL, licenses, s, value)
	}), nil
}

func validateLicense(ll *[]license.License, optsL []string, val string, setV *license.License) error {
	val = strings.ToLower(val)

	i := searching.BinarySearchAuto(optsL, val)
	if i == -1 {
		return fmt.Errorf("unknown license: %s", val)
	}

	for _, v := range *ll {
		if strings.EqualFold(v.Name, val) || strings.EqualFold(v.Spdx, val) {
			*setV = v
			break
		}
	}
	return nil
}

// PromptForProjectName prompts user to enter project name
// or returns the name from options if it's provided.
func PromptForProjectName(name, projectRootDir *string) *huh.Input {
	if options.NewOpts.Name.Value() != "" {
		if err := validateProjectName(options.NewOpts.Name.Value(), name, projectRootDir); err == nil {
			log.Debugf("name prefilled and is valid: %s\n", options.NewOpts.Name.Value())
			return nil
		}
	}

	return huh.NewInput().Title("Name of your project").Validate(func(s string) error {
		return validateProjectName(s, name, projectRootDir)
	})
}

func validateProjectName(s string, name, projectRootDir *string) error {
	s = strings.TrimSpace(s)
	if s == "" {
		return fmt.Errorf("project name cannot be empty")
	}

	if !options.NewOpts.AllowSpaces {
		s = strings.ReplaceAll(s, " ", "-")
	}
	pDir, err := filepath.Rel(".", s)
	if err != nil {
		return err
	}

	if ok, err := utils.PathExists(pDir); err != nil {
		return err
	} else if ok {
		return fmt.Errorf("directory '%s' already exists", s)
	}

	*projectRootDir = pDir
	*name = s
	return nil
}
