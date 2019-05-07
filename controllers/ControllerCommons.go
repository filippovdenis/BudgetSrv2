package controllers
import "BudgetSrv2/repoInterfaces"

type RepoHolder struct {
	CardRepo repoInterfaces.CardRepo
}

var repoHolder *RepoHolder

func SetRepoHolder(holder *RepoHolder) {
	repoHolder = holder
}

func GetRepoHolder() *RepoHolder {
	return repoHolder
}
