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
	"github.com/caffeine-addictt/waku/cmd/ui"
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

		initialPrompts := make([]*huh.Group, 0, 2)
		err := ui.RunWithSpinner("setting things up...", func() error {
			log.Debugln("creating name and license prompts...")
			namePrompt := template.PromptForProjectName(&name, &projectRootDir)
			if namePrompt != nil {
				initialPrompts = append(initialPrompts, huh.NewGroup(namePrompt))
			}

			if options.NewOpts.NoLicense {
				log.Debugln("no-license is set, skipping license prompt...")
				return nil
			}

			licenseSelect, err := template.PromptForLicense(&license)
			if err != nil {
				return errors.ToWakuError(err)
			}
			if licenseSelect != nil {
				initialPrompts = append(initialPrompts, huh.NewGroup(licenseSelect))
			}

			return nil
		})
		if err != nil {
			return errors.ToWakuError(err)
		}

		log.Debugln("running prompts...")
		if err := huh.NewForm(initialPrompts...).WithAccessible(options.GlobalOpts.Accessible).Run(); err != nil {
			return errors.ToWakuError(err)
		}

		err = ui.RunWithSpinner(fmt.Sprintf("creating project at '%s'...", projectRootDir), func() error {
			if err := os.Mkdir(projectRootDir, utils.DirPerms); err != nil {
				return err
			}

			cleanup.ScheduleError(func() error {
				log.Debugf("removing project dir: %s\n", projectRootDir)
				if err := os.RemoveAll(projectRootDir); err != nil {
					return errors.NewWakuErrorf("failed to cleanup project dir: %v", err)
				}
				return nil
			})

			return nil
		})
		if err != nil {
			return errors.ToWakuError(err)
		}

		// Clone repo
		var rootDir string
		var tmpDir string
		err = ui.RunWithSpinner("retrieving template...", func() error {
			_tmpDir, err := options.NewOpts.GetSource()
			if err != nil {
				return nil
			}
			tmpDir = _tmpDir

			// Resolve dir
			rootDir = tmpDir
			if options.NewOpts.Directory.Value() != "" {
				rootDir = filepath.Join(tmpDir, options.NewOpts.Directory.Value())
				log.Debugf("resolved directory to: %s\n", rootDir)

				ok, err := utils.IsDir(rootDir)
				if err != nil {
					return err
				}
				if !ok {
					return errors.NewWakuErrorf("directory '%s' does not exist", options.NewOpts.Directory.Value())
				}
			}

			return nil
		})
		if err != nil {
			return errors.ToWakuError(err)
		}

		// Parse template.json
		var configFilePath string
		var wakuTemplate *config.TemplateJson

		err = ui.RunWithSpinner("parsing waku config...", func() error {
			configFilePath, wakuTemplate, err = template.ParseConfig(rootDir)
			if err != nil {
				return err
			}

			return err
		})
		if err != nil {
			return errors.ToWakuError(err)
		}

		// Resolve prompts
		styleDir, styleInfo, prompts, err := resolveTemplateStylePrompts(wakuTemplate, rootDir)
		if err != nil {
			return errors.ToWakuError(err)
		}

		var licenseText string
		licenseTmpl := make(map[string]string, len(license.Wants))

		if !options.NewOpts.NoLicense {
			err = ui.RunWithSpinner("resolving license...", func() error {
				licenseText, err = license.GetLicenseText()
				if err != nil {
					return errors.NewWakuErrorf("failed to get license text: %v\n", err)
				}

				for _, v := range license.Wants {
					licenseTmpl[v] = fmt.Sprintf("Value for license %s?", v)
					delete(prompts, v)
				}

				return nil
			})
			if err != nil {
				return errors.ToWakuError(err)
			}
		}
		log.Debugf("resolved prompts to: %v\n", prompts)

		stylePromptGroups := make([]*huh.Group, 0, len(prompts))
		finalTemplateData := make(map[string]any, len(prompts)+len(licenseTmpl))

		err = ui.RunWithSpinner("collecting prompts...", func() error {
			for _, v := range prompts {
				stylePromptGroups = append(stylePromptGroups, huh.NewGroup(v.GetPrompt(finalTemplateData)))
			}
			for n, v := range licenseTmpl {
				stylePromptGroups = append(stylePromptGroups, huh.NewGroup(huh.NewText().Title(v).Validate(func(s string) error {
					s = strings.TrimSpace(s)
					if s == "" {
						return fmt.Errorf("cannot be empty")
					}

					licenseTmpl[n] = s
					finalTemplateData[n] = s
					return nil
				})))
			}
			return nil
		})
		if err != nil {
			return errors.ToWakuError(err)
		}

		log.Debugf("resolved prompt groups to: %v\n", stylePromptGroups)
		if err := huh.NewForm(stylePromptGroups...).WithAccessible(options.GlobalOpts.Accessible).Run(); err != nil {
			return errors.ToWakuError(err)
		}

		// Get file paths
		var styleFilePaths []string
		var configRelPath string
		err = ui.RunWithSpinner("collecting files...", func() error {
			styleFilePaths, err = utils.WalkDirRecursive(styleDir)
			if err != nil {
				return err
			}
			log.Debugf("resolved file paths in style: %v\n", styleFilePaths)

			configRelPath, err = filepath.Rel(tmpDir, configFilePath)
			if err != nil {
				return err
			}
			log.Debugf("resolved config rel path to: %s\n", configRelPath)

			return err
		})
		if err != nil {
			return errors.ToWakuError(err)
		}

		// Handle ignores
		var filePathsToWrite types.Set[string]
		err = ui.RunWithSpinner("filtering files...", func() error {
			log.Infoln("applying ignore rules...")
			ignoreRules := types.NewSet(
				".git/",
				"LICENSE*",
				configRelPath,
			)
			if wakuTemplate.Ignore != nil {
				ignoreRules.Union(types.Set[string](*wakuTemplate.Ignore))
			}
			if wakuTemplate.Setup != nil {
				ignoreRules.Add(wakuTemplate.Setup.Any)
				ignoreRules.Add(wakuTemplate.Setup.Linux)
				ignoreRules.Add(wakuTemplate.Setup.Darwin)
				ignoreRules.Add(wakuTemplate.Setup.Windows)
			}
			if wakuTemplate.Styles != nil && styleInfo.Ignore != nil {
				ignoreRules.Union(types.Set[string](*styleInfo.Ignore))

				if styleInfo.Setup != nil {
					ignoreRules.Add(styleInfo.Setup.Any)
					ignoreRules.Add(styleInfo.Setup.Linux)
					ignoreRules.Add(styleInfo.Setup.Darwin)
					ignoreRules.Add(styleInfo.Setup.Windows)
				}
			}

			// account for template.json having a '!.git/'
			ignoreRules = template.ResolveGlobs(ignoreRules, types.NewSet(".git/", "LICENSE"))
			log.Debugf("ignore rules applied: %v\n", ignoreRules)

			filePathsToWrite = template.ResolveGlobs(types.NewSet(styleFilePaths...), ignoreRules)
			log.Debugf("resolved files to write: %v\n", filePathsToWrite)

			return nil
		})
		if err != nil {
			return errors.ToWakuError(err)
		}

		// Handle writing files
		err = ui.RunWithSpinner("writing files...", func() error {
			finalTemplateData["Name"] = name
			if !options.NewOpts.NoLicense {
				finalTemplateData["License"] = license.Name
				finalTemplateData["Spdx"] = license.Spdx
			}
			log.Debugf("final template data: %v\n", finalTemplateData)

			return WriteFiles(styleDir, projectRootDir, filePathsToWrite.ToSlice(), licenseText, finalTemplateData, licenseTmpl)
		})
		if err != nil {
			return errors.NewWakuErrorf("failed to write files: %s\n", err)
		}

		if options.NewOpts.NoGit {
			log.Infoln("skipping git initialization")
		} else {
			err = ui.RunWithSpinner("initializing git...", func() error {
				return git.Init(projectRootDir)
			})
			if err != nil {
				fmt.Printf("failed to initialize git: %s\n", err)
				return errors.NewWakuErrorf("failed to initialize git: %s\n", err)
			}
		}

		dirPath, err := filepath.Abs(projectRootDir)
		if err != nil {
			return errors.ToWakuError(err)
		}

		log.Printf("Project created at: %s\n", dirPath)
		return nil
	},
}

func init() {
	AddNewCmdFlags(NewCmd)
}

func AddNewCmdFlags(cmd *cobra.Command) {
	cmd.Flags().VarP(&options.NewOpts.Repo, "repo", "r", "source repository to template from")
	cmd.Flags().VarP(&options.NewOpts.Source, "source", "s", "source repository to template from")
	cmd.Flags().VarP(&options.NewOpts.Branch, "branch", "b", "branch to clone from")
	cmd.Flags().VarP(&options.NewOpts.Directory, "directory", "D", "directory where config is located")
	cmd.Flags().VarP(&options.NewOpts.Name, "name", "n", "name of the project")
	cmd.Flags().VarP(&options.NewOpts.License, "license", "l", "license to use for the project")
	cmd.Flags().VarP(&options.NewOpts.Style, "style", "S", "which style to use")
	cmd.Flags().BoolVarP(&options.NewOpts.NoGit, "no-git", "G", options.NewOpts.NoGit, "whether to not initialize git")
	cmd.Flags().BoolVarP(&options.NewOpts.NoLicense, "no-license", "L", options.NewOpts.NoLicense, "whether to not include a license")

	if err := cmd.Flags().MarkDeprecated("repo", "Please use --source instead."); err != nil {
		panic(err)
	}
	cmd.MarkFlagsMutuallyExclusive("source", "repo")
	cmd.MarkFlagsMutuallyExclusive("license", "no-license")
}

func resolveTemplateStylePrompts(wakuTemplate *config.TemplateJson, rootDir string) (styleRoot string, style *config.TemplateStyle, prompts map[string]config.TemplatePrompt, err error) {
	if len(wakuTemplate.Styles) == 0 {
		return "", nil, nil, errors.NewWakuErrorf("no styles found in")
	}

	var styleName types.CleanString
	var styleInfo config.TemplateStyle
	if len(wakuTemplate.Styles) == 1 {
		for s, v := range wakuTemplate.Styles {
			styleName = s
			styleInfo = v
			styleRoot = filepath.Join(rootDir, v.Source.String())
			break
		}
	} else {
		if err := huh.NewForm(huh.NewGroup(
			template.PromptForStyle(wakuTemplate.Styles, &styleName, &styleInfo),
		)).WithAccessible(options.GlobalOpts.Accessible).Run(); err != nil {
			return "", nil, nil, err
		}

		styleRoot = filepath.Join(rootDir, styleInfo.Source.String())
	}
	log.Debugf("resolved style to: %s\n", rootDir)

	// Handle prompts
	prompts = make(map[string]config.TemplatePrompt, len(styleInfo.Prompts))
	log.Debugln("resolving prompts...")
	if wakuTemplate.Prompts != nil {
		for _, ask := range wakuTemplate.Prompts {
			prompts[string(ask.Key)] = ask
		}
	}
	if wakuTemplate.Styles != nil && styleInfo.Prompts != nil {
		for _, ask := range styleInfo.Prompts {
			prompts[string(ask.Key)] = ask
		}
	}
	log.Debugf("resolved style prompts to: %v\n", prompts)

	style = &styleInfo
	return
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

	if options.NewOpts.NoLicense {
		wg.Done()
	} else {
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
	}

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
