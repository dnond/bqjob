package jobinfo

import (
	"bqjob/utils"

	bq "google.golang.org/api/bigquery/v2"
)

const maxResultPerRequest = 1000 //list取得時の１回の取得数
const maxResult = 400            //結果取得数。これを超えると次の取得は行わない

func (jobinfo *JobInfo) ListJobs(pageToken string, target string, listedJobs []string) ([]string, error) {
	var err error

	jobs, nextPageToken, err := jobinfo.getList(pageToken)
	if err != nil {
		return nil, err
	}

	for _, job := range jobs {
		if target == "error" || target == "" {
			if job.Status.ErrorResult != nil {
				ctime := ""
				if job.Statistics != nil {
					creationTime, _ := utils.MsToTime(job.Statistics.CreationTime)
					ctime = creationTime.Format("2006-01-02 15:04:05")
				}
				listedJobs = append(listedJobs, "\x1b[31m"+ctime+"\x1b[0m "+job.JobReference.JobId+" "+job.Status.ErrorResult.Reason)
			}
		} else {
			ctime := ""
			if job.Statistics != nil {
				creationTime, _ := utils.MsToTime(job.Statistics.CreationTime)
				ctime = creationTime.Format("2006-01-02 15:04:05")
			}
			listedJobs = append(listedJobs, "\x1b[31m"+ctime+"\x1b[0m "+job.JobReference.JobId+" "+job.Status.State)
		}
	}

	if nextPageToken != "" && len(listedJobs) <= maxResult {
		listedJobs, err = jobinfo.ListJobs(nextPageToken, target, listedJobs)
		if err != nil {
			return nil, err
		}
	}
	return listedJobs, nil
}

func (jobinfo *JobInfo) getList(pageToken string) ([]*bq.JobListJobs, string, error) {
	req := jobinfo.List(jobinfo.projectId).
		AllUsers(true).
		MaxResults(int64(maxResultPerRequest)).
		Projection("full").
		PageToken(pageToken).
		Context(jobinfo.Context)

	res, err := req.Do()
	if err != nil {
		return nil, "", err
	}

	return res.Jobs, res.NextPageToken, nil
}
