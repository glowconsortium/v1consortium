package service

import (
	"context"

	"github.com/riverqueue/river"
	"github.com/riverqueue/river/rivertype"
)

type (
	IRiverClient interface {
		Insert(ctx context.Context, args river.JobArgs, opts *river.InsertOpts) (*rivertype.JobInsertResult, error)
	}
)

var (
	localRiverClient IRiverClient
)

func RiverClient() IRiverClient {
	return localRiverClient
}

func RegisterRiverClient(client IRiverClient) {
	localRiverClient = client
}
