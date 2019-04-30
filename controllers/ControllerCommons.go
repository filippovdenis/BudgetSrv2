package controllers
import "BudgetSrv2/repoInterfaces"

type RepoHolder struct {
	cardRepo repoInterfaces.CardRepo
}

var repoHolder *RepoHolder

func setRepoHolder(holder *RepoHolder) {
	repoHolder = holder
}

func getRepoHolder() *RepoHolder {
	return repoHolder
}
