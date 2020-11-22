package gogit

import (
	"errors"
	"github.com/sirupsen/logrus"
)

func (r *Repo) Promote(module, stage string) error {
	tagsForModule, err := r.TagsForHead(module, stage)
	if err != nil {
		return err
	}
	if len(tagsForModule) > 0 {
		r.log.Printf("Already tagged with %s, continuing", tagsForModule[0])
		return nil
	}

	tagsForHead, err := r.TagsForHead(module)
	if err != nil {
		return err
	}

	if len(tagsForHead) == 0 {
		return errors.New("un-versioned commit")
	}

	if len(tagsForHead) > 1 {
		logrus.Warnln("Multiple version tags on same commit for " + module)
	}

	current := tagsForHead[len(tagsForHead)-1]
	promoted := current + "+" + stage
	r.log.Printf("Promoting version %s -> %s", current, promoted)
	//
	revision, err := r.gitRepo.ResolveRevision("HEAD")
	if err != nil {
		return err
	}

	_, err = r.gitRepo.CreateTag(promoted, *revision, nil)
	return err
}
