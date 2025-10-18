package steps

import (
	"context"
	"fmt"
	"time"
	"v1consortium/internal/pkg/riverjobs"
)

// BaseStep provides common functionality for all signup steps
type BaseStep struct {
	name        string
	queue       string
	timeout     time.Duration
	retryPolicy riverjobs.RetryPolicy
}

// NewBaseStep creates a new base step
func NewBaseStep(name, queue string, timeout time.Duration, retryPolicy riverjobs.RetryPolicy) *BaseStep {
	return &BaseStep{
		name:        name,
		queue:       queue,
		timeout:     timeout,
		retryPolicy: retryPolicy,
	}
}

// GetName returns the step name
func (s *BaseStep) GetName() string {
	return s.name
}

// GetQueue returns the queue this step should run in
func (s *BaseStep) GetQueue() string {
	return s.queue
}

// GetTimeout returns the maximum execution time
func (s *BaseStep) GetTimeout() time.Duration {
	return s.timeout
}

// GetRetryPolicy returns the retry configuration
func (s *BaseStep) GetRetryPolicy() riverjobs.RetryPolicy {
	return s.retryPolicy
}

// GetPreconditions returns conditions that must be met
func (s *BaseStep) GetPreconditions() []riverjobs.Precondition {
	return []riverjobs.Precondition{}
}

// IsRetryable determines if an error is worth retrying
func (s *BaseStep) IsRetryable(err error) bool {
	if stepError, ok := err.(*riverjobs.StepError); ok {
		return stepError.Retryable
	}
	// Default: network and database errors are retryable
	return true
}

// Compensate provides default compensation (no-op)
func (s *BaseStep) Compensate(ctx context.Context, input riverjobs.StepInput) error {
	// Default implementation does nothing
	return nil
}

// Execute must be implemented by concrete steps
func (s *BaseStep) Execute(ctx context.Context, input riverjobs.StepInput) (riverjobs.StepResult, error) {
	return riverjobs.StepResult{}, fmt.Errorf("execute method must be implemented by concrete step")
}
