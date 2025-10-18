// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"time"
	"v1consortium/internal/pkg/riverjobs"

	"github.com/riverqueue/river"
	"github.com/riverqueue/river/rivertype"
)

type (
	IUserSignupWorker interface {
		Middleware(job *rivertype.JobRow) []rivertype.WorkerMiddleware
		NextRetry(job *river.Job[riverjobs.UserSignupArgs]) time.Time
		Timeout(job *river.Job[riverjobs.UserSignupArgs]) time.Duration
		Work(ctx context.Context, job *river.Job[riverjobs.UserSignupArgs]) error
		FailWorkflow(ctx context.Context, workflowID string, stepName string, err error)
		StartNewFlow(ctx context.Context, args riverjobs.UserSignupArgs, opts *river.InsertOpts) (string, error)
		GetSignupStatus(ctx context.Context, workflowID string) (*riverjobs.WorkflowExecution, error)
	}
)

var (
	localUserSignupWorker IUserSignupWorker
)

func UserSignupWorker() IUserSignupWorker {
	if localUserSignupWorker == nil {
		panic("implement not found for interface IUserSignupWorker, forgot register?")
	}
	return localUserSignupWorker
}

func RegisterUserSignupWorker(i IUserSignupWorker) {
	localUserSignupWorker = i
}
