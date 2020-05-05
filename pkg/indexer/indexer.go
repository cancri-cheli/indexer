package indexer

import (
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"strings"
)

// Index creates index html files of liturgical books and services
func Index(site string) error {
	return nil
}

/**
FileMatcher returns all files recursively in the dir that have the specified file extension
without a dot, and match any one of the supplied expressions.
The expressions are regular expressions that will match the name of a file, but
without the extension.  The extension will be automatically added to each expression.
The expressions may be nil, in which case all file patterns will match.
*/
func FileMatcher(dir, extension string, expressions []string) ([]string, error) {
	var result []string
	extensionPattern := "\\." + extension
	// precompile the expressions
	if expressions == nil {
		expressions = []string{".*"}
	}
	patterns := make([]*regexp.Regexp, len(expressions))
	for i, e := range expressions {
		p, err := regexp.Compile(e + extensionPattern)
		if err != nil {
			return result, err
		}
		patterns[i] = p
	}
	// now walk the files and apply the regular expressions
	err := filepath.Walk(ResolvePath(dir),
		func(path string, info os.FileInfo, err error) error {
			for _, p := range patterns {
				if p.MatchString(info.Name()) {
					result = append(result, path)
				}
			}
			return nil
		})
	return result, err
}
func ResolvePath(path string) string {
	usr, _ := user.Current()
	dir := usr.HomeDir
	var newPath = ""
	if path == "~" {
		newPath = dir
	} else if strings.HasPrefix(path, "~/") {
		// Use strings.HasPrefix so we don't match paths like
		// "/something/~/something/"
		newPath = filepath.Join(dir, path[2:])
	} else {
		newPath = path
	}
	return newPath
}
