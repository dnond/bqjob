package jobinfo

import (
	"bqjob/bqservice"
	"context"

	bq "google.golang.org/api/bigquery/v2"
)

type JobInfo struct {
	*bq.JobsService
	Context   context.Context
	projectId string
}

func NewJobInfo(_bqservice *bqservice.BqService, _projectId string) *JobInfo {
	jobinfo := &JobInfo{
		JobsService: bq.NewJobsService(_bqservice.Service),
		Context:     _bqservice.Context,
		projectId:   _projectId,
	}
	return jobinfo
}
