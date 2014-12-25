package main

import (
	"fmt"

	"github.com/bndr/gopencils"
)

type Repo struct {
	*Project
	Name     string
	Resource *gopencils.Resource
}

func (repo *Repo) GetPullRequest(id int64, commit string) PullRequest {
	var res *gopencils.Resource

	if commit == "" {
		res = repo.Resource.Res("pull-requests").Id(fmt.Sprint(id))
	} else {
		res = repo.Resource.Res("pull-requests").Id(fmt.Sprint(id)).Res("commits").Id(commit)
	}

	return PullRequest{
		Repo:     repo,
		Id:       id,
		Resource: res,
	}
}

func (repo *Repo) ListPullRequest(state string) ([]PullRequest, error) {
	reply := struct {
		Size       int
		Limit      int
		IsLastPage bool
		Values     []PullRequest
	}{}

	query := map[string]string{
		"state": state,
	}

	err := repo.DoGet(repo.Resource.Res("pull-requests", &reply), query)
	if err != nil {
		return nil, err
	}

	return reply.Values, nil
}
