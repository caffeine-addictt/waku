package template

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/caffeine-addictt/waku/internal/types"
	"github.com/caffeine-addictt/waku/pkg/config"
	"github.com/caffeine-addictt/waku/pkg/log"
)

// GetStyleResources returns the list of resources that should be copied
// over to the new templated project.
//
// This also accounts for extended files and ignore rules.
//
// `configParent` should be the parent directory of the config file
func GetStyleResources(c *config.TemplateJson, s *config.TemplateStyle, configParentDir string) ([]types.StyleResource, error) {
	ignoreRules := types.NewSet(".git/", "LICENSE*")
	if c.Ignore != nil {
		ignoreRules.Union(types.Set[string](*c.Ignore))
	}
	if s.Ignore != nil {
		si := make(types.Set[string], len(*s.Ignore))
		for path := range *s.Ignore {
			si.Add(filepath.Join(s.Source.String(), path))
		}

		ignoreRules.Union(si)
	}

	// account for !.git/ in ignore rules
	ignoreRules = ResolveGlobs(ignoreRules, types.NewSet(".git/"))
	log.Debugf("ignore rules: %v\n", ignoreRules)

	stylePaths, err := getResourcePaths(filepath.Join(configParentDir, s.Source.String()))
	if err != nil {
		return nil, err
	}
	paths := make(types.Set[string], len(stylePaths))
	for p := range stylePaths {
		paths.Add(filepath.Join(s.Source.String(), p))
	}
	log.Debugf("unfiltered paths: %v\n", paths)

	filteredPaths := ResolveGlobs(paths, ignoreRules)
	log.Debugf("filtered paths: %v\n", filteredPaths)

	resources := make([]types.StyleResource, 0, len(filteredPaths))
	for v := range filteredPaths {
		parts := strings.Split(v, "/")
		resources = append(resources, types.StyleResource{
			TemplateResourceRelPath: v,
			TemplatedProjectRelPath: strings.Join(parts[min(len(parts), 1):], "/"),
		})
	}

	return resources, err
}

func getResourcePaths(root string) (types.Set[string], error) {
	paths := types.NewSet[string]()
	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		relPath = filepath.ToSlash(relPath)

		// skip root
		if relPath == "." {
			return nil
		}

		if !d.IsDir() {
			paths.Add(relPath)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return paths, nil
}
