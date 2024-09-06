package commands

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/caffeine-addictt/template/cmd/config"
	"github.com/caffeine-addictt/template/cmd/license"
	"github.com/caffeine-addictt/template/cmd/options"
	"github.com/caffeine-addictt/template/cmd/template"
	"github.com/caffeine-addictt/template/cmd/utils"
	"github.com/caffeine-addictt/template/cmd/utils/types"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var NewCmd = &cobra.Command{
	Use:     "new",
	Aliases: []string{"init"},
	Short:   "create a new project",
	Long:    "Create a new project from a template",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return options.NewOpts.Validate()
	},
	Run: func(cmd *cobra.Command, args []string) {
		exitCode := 0

		var name string
		var projectRootDir string
		var license license.License

		licenseSelect, err := template.PromptForLicense(&license)
		if err != nil {
			cmd.PrintErrln(err)
			exitCode = 1
			return
		}

		if err := huh.NewForm(
			huh.NewGroup(template.PromptForProjectName(&name, &projectRootDir)),
			huh.NewGroup(licenseSelect),
		).WithAccessible(options.GlobalOpts.Accessible).Run(); err != nil {
			cmd.PrintErrln(err)
			exitCode = 1
			return
		}

		options.Infof("creating project in '%s'...\n", projectRootDir)
		if err := os.Mkdir(projectRootDir, os.ModePerm); err != nil {
			cmd.PrintErrln(err)
			exitCode = 1
			return
		}

		// Clone repo
		tmpDir, err := options.NewOpts.CloneRepo()
		if err != nil {
			cmd.PrintErrf("could not clone git repo: %s", err)
			exitCode = 1
			return
		}
		gracefullyCleanupDir(tmpDir)
		defer func() {
			cleanupDir(tmpDir)
			os.Exit(exitCode)
		}()

		// Resolve dir
		rootDir := tmpDir
		if options.NewOpts.Directory.Value() != "" {
			rootDir = filepath.Join(tmpDir, options.NewOpts.Directory.Value())
			options.Debugf("resolved directory to: %s\n", rootDir)

			ok, err := utils.IsDir(rootDir)
			if err != nil {
				cmd.PrintErrln(err)
				exitCode = 1
				return
			}
			if !ok {
				cmd.PrintErrf("directory '%s' does not exist\n", options.NewOpts.Directory.Value())
				exitCode = 1
				return
			}
		}

		// Parse template.json
		options.Infoln("Parsing template.json...")
		tmpl, err := template.ParseConfig(filepath.Join(rootDir, "template.json"))
		if err != nil {
			cmd.PrintErrln(err)
			exitCode = 1
			return
		}

		// TODO: handle Prompts
		// TODO: handle writing files in async
		// Resolve style to use
		var style types.CleanString
		var styleInfo config.TemplateStyle

		if tmpl.Styles != nil && len(*tmpl.Styles) == 1 {
			for s, v := range *tmpl.Styles {
				style = s
				styleInfo = v
				rootDir = filepath.Join(rootDir, v.Source.String())
				break
			}
		} else if tmpl.Styles != nil {
			if err := huh.NewForm(huh.NewGroup(
				template.PromptForStyle(*tmpl.Styles, &style, &styleInfo),
			)).WithAccessible(options.GlobalOpts.Accessible).Run(); err != nil {
				cmd.PrintErrln(err)
				exitCode = 1
				return
			}

			rootDir = filepath.Join(rootDir, styleInfo.Source.String())
		}
		options.Debugf("resolved style to: %s\n", rootDir)

		// Handle prompts
		options.Debugln("resolving prompts...")
		extraPrompts := map[string]string{}
		if tmpl.Prompts != nil {
			for val, ask := range *tmpl.Prompts {
				extraPrompts[string(val)] = string(ask)
			}
		}
		if tmpl.Styles != nil && styleInfo.Prompts != nil {
			for val, ask := range *styleInfo.Prompts {
				extraPrompts[string(val)] = string(ask)
			}
		}
		options.Debugf("resolved prompts to: %v\n", extraPrompts)

		prompts := make([]*huh.Group, len(extraPrompts))
		for n, v := range extraPrompts {
			prompts = append(prompts, huh.NewGroup(huh.NewText().Title(v).Validate(func(s string) error {
				s = strings.TrimSpace(s)
				if s == "" {
					return fmt.Errorf("cannot be empty")
				}

				extraPrompts[n] = s
				return nil
			})))
		}
		if err := huh.NewForm(prompts...).WithAccessible(options.GlobalOpts.Accessible).Run(); err != nil {
			cmd.PrintErrln(err)
			exitCode = 1
			return
		}

		// Get file paths
		options.Infoln("Getting file paths...")
		paths, err := utils.WalkDirRecursive(projectRootDir)
		if err != nil {
			cmd.PrintErrln(err)
			exitCode = 1
			return
		}

		// Handle ignores
		if tmpl.Ignore != nil {
			options.Infoln("Applying ignores...")
			pathsSet := template.ResolveIncludes(types.NewSet(paths...), types.Set[string](*tmpl.Ignore))
			paths = pathsSet.ToSlice()
		}
		options.Debugf("Resolved files to write: %v", paths)

		// cmd.Printf("%v", paths)
	},
}

func init() {
	AddNewCmdFlags(NewCmd)
}

func AddNewCmdFlags(cmd *cobra.Command) {
	cmd.Flags().VarP(&options.NewOpts.Repo, "repo", "r", "source repository to template from")
	cmd.Flags().VarP(&options.NewOpts.Branch, "branch", "b", "branch to clone from")
	cmd.Flags().VarP(&options.NewOpts.Directory, "directory", "D", "directory where 'template.json' is located")
	cmd.Flags().VarP(&options.NewOpts.Name, "name", "n", "name of the project")
	cmd.Flags().VarP(&options.NewOpts.License, "license", "l", "license to use for the project")
	cmd.Flags().VarP(&options.NewOpts.Style, "style", "S", "which style to use")
}

// To catch interrupts and gracefully cleanup
func gracefullyCleanupDir(dir string) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Printf("%v received, cleaning up...\n", sig)
		cleanupDir(dir)
	}()
}

func cleanupDir(dir string) {
	if err := os.RemoveAll(dir); err != nil {
		fmt.Printf("Failed to clean up %s: %s\n", dir, err)
		os.Exit(1)
		return
	}
}
