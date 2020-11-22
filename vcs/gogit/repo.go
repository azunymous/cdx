// Package gogit interacts with git and git tags via the gogit git implementation
package gogit

import (
	"github.com/azunymous/cdx/parse"
	"github.com/blang/semver/v4"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/sirupsen/logrus"
	"regexp"
	"sort"
)

// Repo is a VCS repository that can be manipulated
type Repo struct {
	gitRepo *git.Repository
	log     logrus.StdLogger
}

// NewRepo returns a new repository, search recursively upwards for a git repository
func NewRepo() (*Repo, error) {
	gr, err := git.PlainOpenWithOptions(".", &git.PlainOpenOptions{DetectDotGit: true})

	if err != nil {
		return nil, err
	}
	return &Repo{gitRepo: gr, log: logrus.New()}, nil
}

// TagsForHead returns sorted version tags at HEAD
// If no promotion stage is provided, only unpromoted tags are returned.
// Only the first provided promotion stage is used for filtering.
func (r *Repo) TagsForHead(module string, stage ...string) ([]string, error) {
	suffix := ""
	if len(stage) > 0 && stage[0] != "" {
		suffix = "\\+" + stage[0]
	}

	current, err := r.gitRepo.ResolveRevision("HEAD")
	if err != nil {
		return nil, err
	}
	tags, err := r.gitRepo.Tags()
	if err != nil {
		return nil, err
	}

	regex, err := regexp.Compile("^" + module + "-[0-9]+\\.[0-9]+\\.[0-9]+" + suffix + "$")
	if err != nil {
		return nil, err
	}

	var t []string
	_ = tags.ForEach(func(reference *plumbing.Reference) error {
		if (reference.Hash() == *current && regex.MatchString(reference.Name().Short())) || r.annotatedMatch(reference, current, regex) {
			t = append(t, reference.Name().Short())
		}
		return nil
	})
	sort.Strings(t)
	return t, nil
}

// TagsForModule returns sorted all semantic version tags for a module and a promotion stage.
// If no promotion stage is provided, only unpromoted tags are returned.
// Only the first provided promotion stage is used for filtering.
func (r *Repo) TagsForModule(module string, stage ...string) ([]string, error) {
	suffix := ""
	if len(stage) > 0 && stage[0] != "" {
		suffix = "\\+" + stage[0]
	}

	tags, err := r.gitRepo.Tags()
	if err != nil {
		return nil, err
	}

	var t []string
	regex, err := regexp.Compile("^" + module + "-[0-9]+\\.[0-9]+\\.[0-9]+" + suffix + "$")
	if err != nil {
		return nil, err
	}
	_ = tags.ForEach(func(reference *plumbing.Reference) error {
		if regex.MatchString(reference.Name().Short()) {
			t = append(t, reference.Name().Short())
		}
		return nil
	})
	sortSemanticVers(t)
	return t, nil
}

func sortSemanticVers(t []string) {
	sort.Slice(t, func(i, j int) bool {
		iVer := semver.MustParse(parse.Version(t[i]))
		jVer := semver.MustParse(parse.Version(t[j]))
		return iVer.LE(jVer)
	})
}

func (r *Repo) annotatedMatch(reference *plumbing.Reference, hash *plumbing.Hash, regex *regexp.Regexp) bool {
	obj, err := r.gitRepo.TagObject(reference.Hash())
	if err != nil {
		return false
	}
	commit, err := obj.Commit()
	if err != nil {
		return false
	}
	if commit.Hash == *hash && regex.MatchString(obj.Name) {
		return true
	}
	return false
}
