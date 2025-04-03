package ui

import (
	"fmt"

	"github.com/caffeine-addictt/waku/cmd/options"
	"github.com/caffeine-addictt/waku/pkg/log"
	"github.com/charmbracelet/huh/spinner"
	"github.com/charmbracelet/lipgloss"
)

func RunWithSpinner(title string, fn func() error) error {
	if log.GetLevel() <= log.INFO {
		log.Infoln(title)
		return fn()
	}

	var retErr error
	spin := buildSpinner(title).Action(func() {
		retErr = fn()
	})

	if err := spin.Run(); err != nil {
		return err
	}

	return retErr
}

func buildSpinner(title string) *spinner.Spinner {
	return spinner.New().Type(spinner.MiniDot).Style(
		lipgloss.NewStyle().Foreground(lipgloss.Color("200")),
	).Accessible(options.GlobalOpts.Accessible).Title(fmt.Sprintf(" %s", title))
}
