package jobinfo

import (
	"bqjob/utils"
	"strconv"

	bq "google.golang.org/api/bigquery/v2"
)

func (jobinfo *JobInfo) detail(job *bq.Job, outputs map[string]string) error {
	if job.Configuration != nil {
		appendConfiguration(job, outputs)
	}

	if job.Statistics != nil {
		appendStatistics(job, outputs)
	}

	if job.Status != nil {
		appendStatus(job, outputs)
	}

	return nil
}

//////////////

func appendConfiguration(job *bq.Job, outputs map[string]string) {
	if job.Configuration.Load != nil {
		outputs["DatasetId"] = job.Configuration.Load.DestinationTable.DatasetId
		outputs["TableId"] = job.Configuration.Load.DestinationTable.TableId
		outputs["MaxBadRecords"] = utils.Int64ToString(job.Configuration.Load.MaxBadRecords)

		if job.Configuration.Load.Schema != nil {
			schema, _ := job.Configuration.Load.Schema.MarshalJSON()
			outputs["Schema"] = string(schema)
		}
	}
}

func appendStatistics(job *bq.Job, outputs map[string]string) {
	creationTime, _ := utils.MsToTime(job.Statistics.CreationTime)
	outputs["CreationTime"] = creationTime.Format("2006-01-02 15:04:05")

	startTime, _ := utils.MsToTime(job.Statistics.StartTime)
	outputs["StartTime"] = startTime.Format("2006-01-02 15:04:05")

	endTime, _ := utils.MsToTime(job.Statistics.EndTime)
	outputs["EndTime"] = endTime.Format("2006-01-02 15:04:05")

	if job.Statistics.Load != nil {
		appendStatisticsLoad(job, outputs)
	}
}

func appendStatisticsLoad(job *bq.Job, outputs map[string]string) {
	outputs["InputFileBytes"] = utils.Int64ToString(job.Statistics.Load.InputFileBytes)
	outputs["InputFiles"] = utils.Int64ToString(job.Statistics.Load.InputFiles)
	outputs["OutputBytes"] = utils.Int64ToString(job.Statistics.Load.OutputBytes)
	outputs["OutputRows"] = utils.Int64ToString(job.Statistics.Load.OutputRows)
}

func appendStatus(job *bq.Job, outputs map[string]string) {
	outputs["State"] = job.Status.State
	prefix := ""

	if job.Status.ErrorResult != nil {
		prefix = "ErrorResult"
		errorMap := errorProtoToMap(job.Status.ErrorResult, prefix)

		utils.MergeMap(outputs, errorMap)
	}

	if job.Status.Errors != nil {
		appendStatusError(job, outputs)
	}
}

func appendStatusError(job *bq.Job, outputs map[string]string) {
	prefix := "ErrorResult"
	index := 0
	for _, errInStatus := range job.Status.Errors {
		index++
		prefix = "errInStatus" + strconv.Itoa(index)
		errorMap := errorProtoToMap(errInStatus, prefix)
		utils.MergeMap(outputs, errorMap)
	}
}

func errorProtoToMap(errProto *bq.ErrorProto, prefix string) map[string]string {
	results := make(map[string]string)

	key := func(elemName string) string {
		return prefix + ":" + elemName
	}

	results[key("DebugInfo")] = errProto.DebugInfo
	results[key("Location")] = errProto.Location
	results[key("Message")] = errProto.Message
	results[key("Reason")] = errProto.Reason

	return results
}
