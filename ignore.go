package the_platinum_searcher

import (
	"path/filepath"

	"github.com/monochromegane/go-gitignore"
	"github.com/monochromegane/go-home"
)

type ignoreMatchers []gitignore.IgnoreMatcher

func (im ignoreMatchers) Match(path string, isDir bool) bool {
	for _, ig := range im {
		if ig == nil {
			return false
		}
		if ig.Match(path, isDir) {
			return true
		}
	}
	return false
}

func newIgnoreMatchers(path string, ignores []string) ignoreMatchers {
	var matchers ignoreMatchers
	for _, i := range ignores {
		if matcher, err := gitignore.NewGitIgnore(filepath.Join(path, i)); err == nil {
			matchers = append(matchers, matcher)
		}
	}
	return matchers
}

func homePtIgnore(base string) gitignore.IgnoreMatcher {
	if homeDir := home.Dir(); homeDir != "" {
		if matcher, err := gitignore.NewGitIgnore(filepath.Join(homeDir, ".ptignore"), base); err == nil {
			return matcher
		}
	}
	return nil
}
