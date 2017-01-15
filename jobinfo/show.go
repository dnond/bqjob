package jobinfo

func (jobinfo *JobInfo) Show(jobId string) (map[string]string, error) {
	var err error

	req := jobinfo.Get(jobinfo.projectId, jobId)
	job, err := req.Do()
	if err != nil {
		return nil, err
	}

	outputs := make(map[string]string)
	outputs["UserEmail"] = job.UserEmail

	err = jobinfo.detail(job, outputs)
	if err != nil {
		return nil, err
	}

	return outputs, nil
}
