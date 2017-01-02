package main

import (
	"bqjob/bqservice"
	"bqjob/jobinfo"
	"fmt"

	"github.com/spf13/cobra"
)

func createJobInfo(projectId string, serviceAccountCredentialFile string) (*jobinfo.JobInfo, error) {
	b, err := bqservice.NewBqService(serviceAccountCredentialFile)
	if err != nil {
		return nil, err
	}
	jobinfo := jobinfo.NewJobInfo(b, projectId)
	return jobinfo, nil
}

func main() {
	var projectId, serviceAccountCredentialFile, targetJobId string
	var rootCmd = &cobra.Command{Use: "bqjob"}
	rootCmd.PersistentFlags().StringVar(&projectId, "project_id", "", "project_id")
	rootCmd.PersistentFlags().StringVar(&serviceAccountCredentialFile, "service_account_credential_file", "", "service_account_credential_file")

	var cmdJobList = &cobra.Command{
		Use:   "ls",
		Short: "show bigquery job error list",
		Long:  "show bigquery job error list",
		Run: func(cmd *cobra.Command, args []string) {
			jobinfo, err := createJobInfo(projectId, serviceAccountCredentialFile)
			if err != nil {
				fmt.Println(err)
				return
			}

			var errorJobs []string
			errorJobs, err = jobinfo.ListErrors("", errorJobs)
			if err != nil {
				fmt.Println(err)
				return
			}

			for _, errorJob := range errorJobs {
				fmt.Println(errorJob)
			}
		},
	}
	rootCmd.AddCommand(cmdJobList)

	var cmdJobShow = &cobra.Command{
		Use:   "show",
		Short: "show bigquery job deital",
		Long:  "show bigquery job detail",
		Run: func(cmd *cobra.Command, args []string) {
			jobinfo, err := createJobInfo(projectId, serviceAccountCredentialFile)
			if err != nil {
				fmt.Println(err)
				return
			}

			jobDetail, err := jobinfo.Show(targetJobId)
			if err != nil {
				fmt.Println(err)
				return
			}
			for k, v := range jobDetail {
				fmt.Println(k + ": \x1b[31m" + v + "\x1b[0m")
			}
		},
	}
	cmdJobShow.Flags().StringVar(&targetJobId, "job_id", "", "job_id")
	rootCmd.AddCommand(cmdJobShow)

	rootCmd.Execute()
}
