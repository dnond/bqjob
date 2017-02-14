package jobinfo

func (jobinfo *JobInfo) Stop(jobId string) (string, error) {
	var err error

	jobCancelCall := jobinfo.Cancel(jobinfo.projectId, jobId)
	_, err = jobCancelCall.Do()
	if err != nil {
		return "", err
	}

	return "accepted", nil
}
