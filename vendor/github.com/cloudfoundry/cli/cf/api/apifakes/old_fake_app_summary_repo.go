package apifakes

import (
	"github.com/cloudfoundry/cli/cf/errors"
	"github.com/cloudfoundry/cli/cf/models"
)

type OldFakeAppSummaryRepo struct {
	GetSummariesInCurrentSpaceApps []models.Application

	GetSummaryErrorCode string
	GetSummaryAppGuid   string
	GetSummarySummary   models.Application
}

func (repo *OldFakeAppSummaryRepo) GetSummariesInCurrentSpace() (apps []models.Application, apiErr error) {
	apps = repo.GetSummariesInCurrentSpaceApps
	return
}

func (repo *OldFakeAppSummaryRepo) GetSummary(appGuid string) (summary models.Application, apiErr error) {
	repo.GetSummaryAppGuid = appGuid
	summary = repo.GetSummarySummary

	if repo.GetSummaryErrorCode != "" {
		apiErr = errors.NewHttpError(400, repo.GetSummaryErrorCode, "Error")
	}

	return
}
