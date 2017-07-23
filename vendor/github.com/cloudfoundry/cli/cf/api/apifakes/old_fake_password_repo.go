package apifakes

import "github.com/cloudfoundry/cli/cf/errors"

type OldFakePasswordRepo struct {
	Score          string
	ScoredPassword string

	UpdateUnauthorized bool
	UpdateNewPassword  string
	UpdateOldPassword  string
}

func (repo *OldFakePasswordRepo) UpdatePassword(old string, new string) (apiErr error) {
	repo.UpdateOldPassword = old
	repo.UpdateNewPassword = new

	if repo.UpdateUnauthorized {
		apiErr = errors.NewHttpError(401, "unauthorized", "Authorization Failed")
	}

	return
}
