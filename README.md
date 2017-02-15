# bqjob

## list BigQuery jobs(error only)

```
bqjob ls --service_account_credential_file=/hogehoge/service_account_credential.json  --project_id=your-project-id
```

### wip!

added target flag

```
bqjob ls --service_account_credential_file=/hogehoge/service_account_credential.json  --project_id=your-project-id --target=ALL
```

targetにerror以外を設定すると、すべてのjobをリストします


## show JobDetail

```
bqjob show --service_account_credential_file=/hogehoge/service_account_credential.json  --project_id=your-project-id --job_id=job_abcdefdf1234
```

## cancel job

```
bqjob cancel --service_account_credential_file=/hogehoge/service_account_credential.json  --project_id=your-project-id --job_id=job_abcdefdf1234
```
