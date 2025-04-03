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
	if s.Include != nil && len(s.Include) != 0 {
		si := make(types.Set[string], len(s.Include))

		for _, includePath := range s.Include {
			if includePath.Ignore == nil {
				continue
			}

			for path := range *includePath.Ignore {
				si.Add(filepath.Join(includePath.Source.String(), path))
			}
		}

		ignoreRules.Union(si)
	}

	// account for !.git/ in ignore rules
	ignoreRules = ResolveGlobs(ignoreRules, types.NewSet(".git/"))
	log.Debugf("ignore rules: %v\n", ignoreRules)

	includePaths := make(map[string]string, len(s.Include)) // includePath: dir
	if s.Include != nil {
		for _, includePath := range s.Include {
			log.Infof("include path: %s\n", includePath.Source.String())
			inPths, err := getResourcePaths(filepath.Join(configParentDir, includePath.Source.String()))
			if err != nil {
				return nil, err
			}

			log.Debugf("resolved include paths: %v\n", inPths)
			for p := range inPths {
				if includePath.Directory == nil {
					includePaths[filepath.Join(includePath.Source.String(), p)] = ""
				} else {
					includePaths[filepath.Join(includePath.Source.String(), p)] = includePath.Directory.String()
				}
			}
		}
	}
	stylePaths, err := getResourcePaths(filepath.Join(configParentDir, s.Source.String()))
	if err != nil {
		return nil, err
	}

	paths := make(types.Set[string], len(stylePaths)+len(includePaths))
	for p := range stylePaths {
		paths.Add(filepath.Join(s.Source.String(), p))
	}
	for p := range includePaths {
		paths.Add(p)
	}
	log.Debugf("unfiltered paths: %v\n", paths)

	filteredPaths := ResolveGlobs(paths, ignoreRules)
	log.Debugf("filtered paths: %v\n", filteredPaths)

	resources := make([]types.StyleResource, 0, len(filteredPaths))
	for v := range filteredPaths {
		parts := strings.Split(v, "/")
		parts = parts[min(len(parts), 1):]

		if dirPrepend, ok := includePaths[v]; ok {
			parts = append([]string{dirPrepend}, parts...)
		}

		resources = append(resources, types.StyleResource{
			TemplateResourceRelPath: v,
			TemplatedProjectRelPath: strings.Join(parts, "/"),
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
