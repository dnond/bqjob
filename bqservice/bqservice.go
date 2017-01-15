package bqservice

import (
	"context"
	"fmt"

	bq "google.golang.org/api/bigquery/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/transport"
)

type BqService struct {
	*bq.Service
	Context context.Context
}

const prodAddr = "https://www.googleapis.com/bigquery/v2/"
const Scope = "https://www.googleapis.com/auth/bigquery"
const userAgent = "gcloud-golang-bigquery/20160429"

func NewBqService(serviceAccountFile string) (*BqService, error) {
	ctx := context.Background()
	o := []option.ClientOption{
		option.WithEndpoint(prodAddr),
		option.WithScopes(Scope),
		option.WithUserAgent(userAgent),
		option.WithServiceAccountFile(serviceAccountFile),
	}
	httpClient, endpoint, _ := transport.NewHTTPClient(ctx, o...)

	service, err := bq.New(httpClient)
	if err != nil {
		return nil, fmt.Errorf("constructing bigquery client: %v", err)
	}
	service.BasePath = endpoint

	_bqservice := &BqService{
		Service: service,
		Context: ctx,
	}
	return _bqservice, nil
}
