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

func (r *Repo) HeadHash() (string, error) {
	head, err := r.gitRepo.Head()
	if err != nil {
		return "", err
	}
	return head.Hash().String(), nil
}

func CheckIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
