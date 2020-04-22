package gogit

import (
	"log"
)

func (r *Repo) OnMaster() bool {
	revHash, err := r.gitRepo.ResolveRevision("origin/master")
	CheckIfError(err)
	revCommit, err := r.gitRepo.CommitObject(*revHash)
	CheckIfError(err)

	headRef, err := r.gitRepo.Head()
	CheckIfError(err)
	headCommit, err := r.gitRepo.CommitObject(headRef.Hash())
	CheckIfError(err)
	isAncestor, err := headCommit.IsAncestor(revCommit)

	CheckIfError(err)
	return isAncestor

}

func CheckIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
