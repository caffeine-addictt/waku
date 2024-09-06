package commands

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
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
		if err := os.Mkdir(projectRootDir, utils.PermOwnerReadWrite); err != nil {
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

		// Handle license stuff
		licenseText, err := license.GetLicenseText()
		if err != nil {
			cmd.PrintErrf("failed to get license text: %s\n", err)
			exitCode = 1
			return
		}

		licenseWants := licenseText.GetWants()
		licenseTmpl := make(map[string]string, len(licenseWants))
		for _, v := range licenseWants {
			licenseTmpl[v] = fmt.Sprintf("Value for license %s?", v)
		}

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
		for n, v := range licenseTmpl {
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
		paths, err := utils.WalkDirRecursive(rootDir)
		if err != nil {
			cmd.PrintErrln(err)
			exitCode = 1
			return
		}

		// Handle ignores
		options.Infoln("Applying ignores...")
		ignoreRules := types.NewSet(
			".git/",
			"LICENSE*",
			"template.json",
		)
		if tmpl.Ignore != nil {
			ignoreRules.Union(types.Set[string](*tmpl.Ignore))
		}
		if tmpl.Styles != nil && styleInfo.Ignore != nil {
			ignoreRules.Union(types.Set[string](*styleInfo.Ignore))
		}

		// account for template.json having a '!.git/'
		ignoreRules = template.ResolveIncludes(ignoreRules, types.NewSet(".git/", "LICENSE"))
		ignoredPaths := template.ResolveIncludes(types.NewSet(paths...), ignoreRules)

		options.Debugf("resolved files to write: %v", ignoredPaths)

		// Handle writing files
		cmd.Println("writing files...")
		finalTmpl := extraPrompts
		finalTmpl["NAME"] = name
		finalTmpl["LICENSE"] = license.Spdx

		if err := WriteFiles(rootDir, projectRootDir, ignoredPaths.ToSlice(), licenseText.Body, finalTmpl, licenseTmpl); err != nil {
			fmt.Printf("failed to write files: %s\n", err)
			exitCode = 1
			return
		}
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

func WriteFiles(tmpRoot, projectRoot string, paths []string, licenseText string, tmpl, licenseTmpl map[string]string) error {
	var wg sync.WaitGroup
	wg.Add(len(paths) + 1)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	errChan := make(chan error, 1)

	for _, path := range paths {
		tmpPath := filepath.Join(tmpRoot, path)
		newPath := filepath.Join(projectRoot, path)
		options.Infof("resolved %s -> %s\n", tmpPath, newPath)

		// write dirs
		dir := filepath.Dir(newPath)
		if dir != "." {
			if err := os.MkdirAll(dir, utils.PermOwnerReadWrite); err != nil {
				return errors.Join(fmt.Errorf("failed to create directory at %s", dir), err)
			}
		}

		// write files
		go func() {
			defer wg.Done()

			tmpFile, err := os.Open(filepath.Clean(tmpPath))
			if err != nil {
				errChan <- err
				return
			}
			defer tmpFile.Close()
			options.Debugf("opened file for reading: %s\n", tmpPath)

			newFile, err := os.OpenFile(filepath.Clean(newPath), os.O_TRUNC|os.O_CREATE|os.O_WRONLY, utils.PermOwnerReadWrite)
			if err != nil {
				errChan <- err
				return
			}
			defer newFile.Close()
			options.Debugf("opened file for writing: %s", newPath)

			reader := bufio.NewScanner(tmpFile)
			writer := bufio.NewWriter(newFile)
			if err := utils.ParseTemplateFile(ctx, tmpl, reader, writer); err != nil {
				errChan <- err
				return
			}

			options.Debugf("flushing buffer for %s", newPath)
			if err := writer.Flush(); err != nil {
				errChan <- err
				return
			}

			options.Debugf("wrote file: %s\n", newPath)
		}()
	}

	go func() {
		defer wg.Done()

		newLicenseText := utils.ParseLicenseText(licenseTmpl, licenseText)

		newPath := filepath.Join(projectRoot, "LICENSE")
		options.Infof("writing to %s\n", newPath)

		newFile, err := os.OpenFile(filepath.Clean(newPath), os.O_TRUNC|os.O_CREATE|os.O_WRONLY, utils.PermOwnerReadWrite)
		if err != nil {
			errChan <- err
			return
		}
		defer newFile.Close()
		options.Debugf("opened file for writing: %s\n", newPath)

		if _, err := newFile.WriteString(newLicenseText); err != nil {
			errChan <- err
			return
		}

		options.Debugf("flushing buffer for %s", newPath)
		if err := newFile.Sync(); err != nil {
			errChan <- err
			return
		}

		options.Debugf("wrote file: %s\n", newPath)
	}()

	// handle canceling if anything goes wrong
	var exitErr error
	go func() {
		options.Infoln("watching for errors")
		if err := <-errChan; err != nil {
			cancel()
			exitErr = err
		}
	}()

	fmt.Printf("waiting for %d files to write\n", len(paths))
	wg.Wait()
	close(errChan)

	fmt.Println("all files written")
	return exitErr
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
