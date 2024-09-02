package template

import (
	"regexp"
	"strings"

	"github.com/caffeine-addictt/template/cmd/utils/types"
)

// Resolve paths to include
//
// Negation always takes priority. i.e. Set["path/to/file", "!path/**"] = Set[]
//
// Syntax:
//
//	"*" is glob for everything
//	"path/to/file" is a single file
//	"path/to/f*" is a glob for a file
//	"path/to/dir/*" is a single dir level glob
//	"path/to/dir/**" == path/to/dir/ is a recursive dir level glob
//	"!path/to/file" is a negated ignore
func ResolveIncludes(paths, ignores types.Set[string]) types.Set[string] {
	negation := types.NewSet[string]()

	// handle explicit includes
	for ignore := range ignores {
		if strings.HasPrefix(ignore, "!") {
			newIgnore := strings.TrimPrefix(ignore, "!")
			for _, p := range handleMatching(&paths, newIgnore) {
				negation.Add(p)
			}
			continue
		}
	}

	// handle "*"
	if ignores.Contains("*") {
		return negation
	}

	result := paths.Copy()
	for ignore := range ignores {
		// handle as removing
		for _, p := range handleMatching(&result, ignore) {
			result.Remove(p)
		}
	}

	return result.Union(negation)
}

func handleMatching(paths *types.Set[string], pattern string) []string {
	matching := make([]string, 0, paths.Len()/2)
	patternParts := strings.Split(pattern, "/")

	// convert pattern parts to regex-able
	nonRecursePartsCount := 0
	isRecusive := false
	newPattern := "^"

a:
	for nonRecursePartsCount < len(patternParts) {
		switch patternParts[nonRecursePartsCount] {
		case "**", "":
			isRecusive = true
			break a
		default:
			newPattern += strings.ReplaceAll(patternParts[nonRecursePartsCount], "*", ".*") + `/`
		}

		nonRecursePartsCount++
	}

	newPattern = strings.TrimSuffix(newPattern, "/") + "$"
	re := regexp.MustCompile(newPattern)

	for p := range *paths {
		pParts := strings.Split(p, "/")

		if len(pParts) < nonRecursePartsCount {
			continue
		}

		common := strings.Join(pParts[:nonRecursePartsCount], "/")
		if !re.MatchString(common) {
			continue
		}

		if !isRecusive || (isRecusive && len(pParts) > nonRecursePartsCount) {
			matching = append(matching, p)
		}
	}

	return matching
}
