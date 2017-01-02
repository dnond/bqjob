package main

import (
	"context"
	"log"

	"google.golang.org/api/iterator"
	"google.golang.org/api/option"

	"cloud.google.com/go/bigquery"
	"github.com/k0kubun/pp"
)

func main() {
	ctx := context.Background()

	projectID := "escape-idol-01"
	serviceAccountFile := "/Users/workman/.google/bigquery/escape_Idol-8d1c9fbb52c7.json"

	ops := option.WithServiceAccountFile(serviceAccountFile)

	client, err := bigquery.NewClient(ctx, projectID, ops)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	datasetName := "my_test"

	dataset := client.Dataset(datasetName)
	it := dataset.Tables(ctx)
	for {
		t, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		pp.Print(t)
	}

	pp.Print("end")
}
