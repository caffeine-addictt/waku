package commands

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/caffeine-addictt/template/cmd/options"
	"github.com/caffeine-addictt/template/cmd/template"
	"github.com/caffeine-addictt/template/cmd/utils"
	"github.com/caffeine-addictt/template/cmd/utils/types"
	"github.com/charmbracelet/huh/spinner"
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
		tmpDir, err := cloneGitRepo()
		if err != nil {
			cmd.PrintErrf("Could not clone git repo: %s", err)
			os.Exit(1)
		}
		gracefullyCleanupDir(tmpDir)
		defer cleanupDir(tmpDir)

		// Resolve dir
		rootDir := tmpDir
		if options.NewOpts.Directory.Value() != "" {
			rootDir = filepath.Join(tmpDir, options.NewOpts.Directory.Value())

			ok, err := utils.IsDir(rootDir)
			if err != nil {
				cmd.PrintErrln(err)
				os.Exit(1)
			}
			if !ok {
				cmd.PrintErrf("directory '%s' does not exist\n", options.NewOpts.Directory.Value())
				os.Exit(1)
			}
		}

		// Parse template.json
		tmpl, err := template.ParseConfig(rootDir)
		if err != nil {
			cmd.PrintErrln(err)
			os.Exit(1)
		}

		// TODO: handle Prompts
		// TODO: handle writing files in async

		// Get file paths
		paths, err := utils.WalkDirRecursive(rootDir)
		if err != nil {
			cmd.PrintErrln(err)
			os.Exit(1)
		}

		// Handle ignores
		if tmpl.Ignore != nil {
			pathsSet := template.ResolveIncludes(types.NewSet(paths...), types.Set[string](*tmpl.Ignore))
			paths = pathsSet.ToSlice()
		}
	},
}

func init() {
	AddNewCmdFlags(NewCmd)
}

func AddNewCmdFlags(cmd *cobra.Command) {
	cmd.Flags().VarP(&options.NewOpts.Repo, "repo", "r", "community source repository for templates")
	cmd.Flags().VarP(&options.NewOpts.Branch, "branch", "b", "branch to clone from [default: main/master]")
	cmd.Flags().VarP(&options.NewOpts.Directory, "directory", "D", "which directory of the template to use [default: /]")
}

// For cloning git repo with spinner
func cloneGitRepo() (string, error) {
	outCh := make(chan string, 1)
	errCh := make(chan error, 1)

	err := spinner.New().Action(func() { options.NewOpts.CloneRepo(outCh, errCh) }).Run()
	if err != nil {
		return "", err
	}

	return <-outCh, <-errCh
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

	os.Exit(0)
}
