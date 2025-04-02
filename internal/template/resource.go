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

	// TODO: get style to prepend source dir
	if s.Ignore != nil {
		ignoreRules.Union(types.Set[string](*s.Ignore))
	}

	// account for !.git/ in ignore rules
	ignoreRules = ResolveGlobs(ignoreRules, types.NewSet(".git/"))
	log.Debugf("ignore rules: %v\n", ignoreRules)

	paths := types.NewSet[string]()
	err := filepath.WalkDir(configParentDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(configParentDir, path)
		if err != nil {
			return err
		}
		relPath = filepath.ToSlash(relPath)

		// skip root
		if relPath == "." {
			return nil
		}

		if d.IsDir() {
			paths.Add(relPath + "/")
		} else {
			paths.Add(relPath)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	log.Debugf("unfiltered paths: %v\n", paths)

	filteredPaths := ResolveGlobs(paths, ignoreRules)
	log.Debugf("filtered paths: %v\n", filteredPaths)

	resources := make([]types.StyleResource, 0, len(filteredPaths))
	for v := range filteredPaths {
		sr := types.StyleResource{
			StyleRelPath:            v,
			TemplatedProjectRelPath: strings.Join(strings.Split(v, "/")[1:], "/"), // remove first dir level
		}

		if strings.HasSuffix(v, "/") {
			sr.Kind = types.DirStyleResourceKind
		} else {
			sr.Kind = types.FileStyleResourceKind
		}

		resources = append(resources, sr)
	}

	return resources, err
}
