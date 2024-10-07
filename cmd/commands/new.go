package commands

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/caffeine-addictt/waku/cmd/cleanup"
	"github.com/caffeine-addictt/waku/cmd/options"
	"github.com/caffeine-addictt/waku/internal/errors"
	"github.com/caffeine-addictt/waku/internal/git"
	"github.com/caffeine-addictt/waku/internal/license"
	"github.com/caffeine-addictt/waku/internal/log"
	"github.com/caffeine-addictt/waku/internal/template"
	"github.com/caffeine-addictt/waku/internal/types"
	"github.com/caffeine-addictt/waku/internal/utils"
	"github.com/caffeine-addictt/waku/pkg/config"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var NewCmd = &cobra.Command{
	Use:           "new",
	Aliases:       []string{"init"},
	Short:         "create a new project",
	Long:          "Create a new project from a template",
	SilenceErrors: true,
	SilenceUsage:  true,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return options.NewOpts.Validate()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		var name string
		var projectRootDir string
		var license license.License

		log.Debugln("Creating name and license prompts...")
		licenseSelect, err := template.PromptForLicense(&license)
		if err != nil {
			return errors.ToWakuError(err)
		}

		if err := huh.NewForm(
			huh.NewGroup(template.PromptForProjectName(&name, &projectRootDir)),
			huh.NewGroup(licenseSelect),
		).WithAccessible(options.GlobalOpts.Accessible).Run(); err != nil {
			return errors.ToWakuError(err)
		}

		log.Infof("creating project in '%s'...\n", projectRootDir)
		if err := os.Mkdir(projectRootDir, utils.DirPerms); err != nil {
			return errors.ToWakuError(err)
		}
		cleanup.ScheduleError(func() error {
			log.Debugf("removing project dir: %s\n", projectRootDir)
			if err := os.RemoveAll(projectRootDir); err != nil {
				return errors.NewWakuErrorf("failed to cleanup project dir: %v", err)
			}
			return nil
		})

		// Clone repo
		tmpDir, err := options.NewOpts.CloneRepo()
		if err != nil {
			return errors.NewWakuErrorf("could not clone git repo: %s", err)
		}
		cleanup.Schedule(func() error {
			log.Debugf("removing tmp dir: %s\n", tmpDir)
			if err := os.RemoveAll(tmpDir); err != nil {
				return errors.NewWakuErrorf("failed to cleanup tmp dir: %v", err)
			}
			return nil
		})

		// Resolve dir
		rootDir := tmpDir
		if options.NewOpts.Directory.Value() != "" {
			rootDir = filepath.Join(tmpDir, options.NewOpts.Directory.Value())
			log.Debugf("resolved directory to: %s\n", rootDir)

			ok, err := utils.IsDir(rootDir)
			if err != nil {
				return errors.ToWakuError(err)
			}
			if !ok {
				return errors.NewWakuErrorf("directory '%s' does not exist", options.NewOpts.Directory.Value())
			}
		}

		// Parse template.json
		log.Infoln("Parsing config...")
		configFilePath, tmpl, err := template.ParseConfig(rootDir)
		if err != nil {
			return errors.ToWakuError(err)
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
				return errors.ToWakuError(err)
			}

			rootDir = filepath.Join(rootDir, styleInfo.Source.String())
		}
		log.Debugf("resolved style to: %s\n", rootDir)

		// Handle license stuff
		licenseText, err := license.GetLicenseText()
		if err != nil {
			return errors.NewWakuErrorf("failed to get license text: %v\n", err)
		}

		// Handle prompts
		log.Debugln("resolving prompts...")
		extraPrompts := map[string]config.TemplatePrompt{}
		if tmpl.Prompts != nil {
			for _, ask := range *tmpl.Prompts {
				extraPrompts[string(ask.Key)] = ask
			}
		}
		if tmpl.Styles != nil && styleInfo.Prompts != nil {
			for _, ask := range *styleInfo.Prompts {
				extraPrompts[string(ask.Key)] = ask
			}
		}

		licenseTmpl := make(map[string]string, len(license.Wants))
		for _, v := range license.Wants {
			licenseTmpl[v] = fmt.Sprintf("Value for license %s?", v)
			delete(extraPrompts, v)
		}
		log.Debugf("resolved prompts to: %v\n", extraPrompts)

		prompts := make([]*huh.Group, 0, len(extraPrompts))
		finalTmpl := make(map[string]any, len(extraPrompts)+len(licenseTmpl))

		for _, v := range extraPrompts {
			prompts = append(prompts, huh.NewGroup(v.GetPrompt(finalTmpl)))
		}
		for n, v := range licenseTmpl {
			prompts = append(prompts, huh.NewGroup(huh.NewText().Title(v).Validate(func(s string) error {
				s = strings.TrimSpace(s)
				if s == "" {
					return fmt.Errorf("cannot be empty")
				}

				licenseTmpl[n] = s
				finalTmpl[n] = s
				return nil
			})))
		}

		log.Debugf("resolved prompt groups to: %v\n", prompts)
		if err := huh.NewForm(prompts...).WithAccessible(options.GlobalOpts.Accessible).Run(); err != nil {
			return errors.ToWakuError(err)
		}

		// Get file paths
		log.Infoln("Getting file paths...")
		paths, err := utils.WalkDirRecursive(rootDir)
		if err != nil {
			return errors.ToWakuError(err)
		}

		// Handle ignores
		log.Infoln("Applying ignores...")
		ignoreRules := types.NewSet(
			".git/",
			"LICENSE*",
			configFilePath,
		)
		if tmpl.Ignore != nil {
			ignoreRules.Union(types.Set[string](*tmpl.Ignore))
		}
		if tmpl.Setup != nil {
			ignoreRules.Add(tmpl.Setup.Any)
			ignoreRules.Add(tmpl.Setup.Linux)
			ignoreRules.Add(tmpl.Setup.Darwin)
			ignoreRules.Add(tmpl.Setup.Windows)
		}
		if tmpl.Styles != nil && styleInfo.Ignore != nil {
			ignoreRules.Union(types.Set[string](*styleInfo.Ignore))

			if styleInfo.Setup != nil {
				ignoreRules.Add(styleInfo.Setup.Any)
				ignoreRules.Add(styleInfo.Setup.Linux)
				ignoreRules.Add(styleInfo.Setup.Darwin)
				ignoreRules.Add(styleInfo.Setup.Windows)
			}
		}

		// account for template.json having a '!.git/'
		ignoreRules = template.ResolveIncludes(ignoreRules, types.NewSet(".git/", "LICENSE"))
		ignoredPaths := template.ResolveIncludes(types.NewSet(paths...), ignoreRules)

		log.Debugf("resolved files to write: %v", ignoredPaths)

		// Handle writing files
		cmd.Println("writing files...")
		finalTmpl["Name"] = name
		finalTmpl["License"] = license.Name
		finalTmpl["Spdx"] = license.Spdx
		log.Debugf("final template: %v\n", finalTmpl)

		if err := WriteFiles(rootDir, projectRootDir, ignoredPaths.ToSlice(), licenseText, finalTmpl, licenseTmpl); err != nil {
			return errors.NewWakuErrorf("failed to write files: %s\n", err)
		}

		if options.NewOpts.NoGit {
			log.Infoln("skipping git initialization")
		} else {
			if err := git.Init(projectRootDir); err != nil {
				fmt.Printf("failed to initialize git: %s\n", err)
				return errors.NewWakuErrorf("failed to initialize git: %s\n", err)
			}
		}

		return nil
	},
}

func init() {
	AddNewCmdFlags(NewCmd)
}

func AddNewCmdFlags(cmd *cobra.Command) {
	cmd.Flags().VarP(&options.NewOpts.Repo, "repo", "r", "source repository to template from")
	cmd.Flags().VarP(&options.NewOpts.Branch, "branch", "b", "branch to clone from")
	cmd.Flags().VarP(&options.NewOpts.Directory, "directory", "D", "directory where config is located")
	cmd.Flags().VarP(&options.NewOpts.Name, "name", "n", "name of the project")
	cmd.Flags().VarP(&options.NewOpts.License, "license", "l", "license to use for the project")
	cmd.Flags().VarP(&options.NewOpts.Style, "style", "S", "which style to use")
	cmd.Flags().BoolVarP(&options.NewOpts.NoGit, "no-git", "G", options.NewOpts.NoGit, "whether to not initialize git")
}

func WriteFiles(tmpRoot, projectRoot string, paths []string, licenseText string, tmpl map[string]any, licenseTmpl map[string]string) error {
	var wg sync.WaitGroup
	wg.Add(len(paths) + 1)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	errChan := make(chan error, 1)
	log.Infof("waiting for %d files to write\n", len(paths))

	for _, path := range paths {
		tmpPath := filepath.Join(tmpRoot, path)
		newPath := filepath.Join(projectRoot, path)
		log.Debugf("resolved %s -> %s\n", tmpPath, newPath)

		// write dirs
		dir := filepath.Dir(newPath)
		if dir != "." {
			if err := os.MkdirAll(dir, utils.DirPerms); err != nil {
				return errors.NewWakuErrorf("failed to create directory at %s: %s", dir, err)
			}
		}

		// write files
		go func() {
			defer wg.Done()

			tmpFile, err := os.Open(filepath.Clean(tmpPath))
			if err != nil {
				errChan <- errors.NewWakuErrorf("failed to open file for reading at %s: %v", tmpPath, err)
				return
			}
			defer tmpFile.Close()
			log.Debugf("opened file for reading: %s\n", tmpPath)

			newFile, err := os.OpenFile(filepath.Clean(newPath), os.O_TRUNC|os.O_CREATE|os.O_WRONLY, utils.FilePerms)
			if err != nil {
				errChan <- errors.NewWakuErrorf("failed to open file for writing at %s: %v", newPath, err)
				return
			}
			defer newFile.Close()
			log.Debugf("opened file for writing: %s\n", newPath)

			reader := bufio.NewScanner(tmpFile)
			writer := bufio.NewWriter(newFile)
			if err := utils.ParseTemplateFile(ctx, tmpl, reader, writer); err != nil {
				errChan <- errors.NewWakuErrorf("failed to parse template from %s to %s: %v", tmpPath, newPath, err)
				return
			}

			log.Debugf("wrote file: %s\n", newPath)
		}()
	}

	go func() {
		defer wg.Done()

		newLicenseText := utils.ParseLicenseText(licenseTmpl, licenseText)

		newPath := filepath.Join(projectRoot, "LICENSE")
		log.Infof("writing to %s\n", newPath)

		newFile, err := os.OpenFile(filepath.Clean(newPath), os.O_TRUNC|os.O_CREATE|os.O_WRONLY, utils.FilePerms)
		if err != nil {
			errChan <- errors.ToWakuError(err)
			return
		}
		defer newFile.Close()
		log.Debugf("opened file for writing: %s\n", newPath)

		if _, err := newFile.WriteString(newLicenseText); err != nil {
			errChan <- errors.NewWakuErrorf("failed to write license text at %s", newPath)
			return
		}

		if err := newFile.Sync(); err != nil {
			errChan <- errors.NewWakuErrorf("failed to flush buffer for %s", newPath)
			return
		}

		log.Debugf("wrote file: %s\n", newPath)
	}()

	// handle canceling if anything goes wrong
	var exitErr error
	go func() {
		log.Infoln("watching for errors")
		if err := <-errChan; err != nil {
			cancel()
			exitErr = err
		}
	}()

	wg.Wait()
	close(errChan)

	log.Infoln("all files written")
	return exitErr
}
