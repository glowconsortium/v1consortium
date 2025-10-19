package riverjobsv2

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// InMemoryWorkflowStore is a simple in-memory implementation of WorkflowStore
// This should be replaced with a database-backed implementation in production
type InMemoryWorkflowStore struct {
	workflows map[string]*WorkflowStatus // workflowID -> status
	hashIndex map[string]string          // inputHash -> workflowID for running workflows
	mutex     sync.RWMutex
}

// NewInMemoryWorkflowStore creates a new in-memory workflow store
func NewInMemoryWorkflowStore() *InMemoryWorkflowStore {
	return &InMemoryWorkflowStore{
		workflows: make(map[string]*WorkflowStatus),
		hashIndex: make(map[string]string),
	}
}

// SaveWorkflowStatus saves or updates workflow status
func (s *InMemoryWorkflowStore) SaveWorkflowStatus(ctx context.Context, status *WorkflowStatus) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Update or create workflow status
	s.workflows[status.WorkflowID] = status

	// Update hash index for running workflows
	if status.Status == "running" {
		hashKey := s.makeHashKey(status.WorkflowName, status.Input)
		if hashKey != "" {
			s.hashIndex[hashKey] = status.WorkflowID
		}
	}

	return nil
}

// GetWorkflowStatus retrieves workflow status by ID
func (s *InMemoryWorkflowStore) GetWorkflowStatus(ctx context.Context, workflowID string) (*WorkflowStatus, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	status, exists := s.workflows[workflowID]
	if !exists {
		return nil, fmt.Errorf("workflow not found: %s", workflowID)
	}

	// Return a copy to avoid concurrent modification
	statusCopy := *status
	return &statusCopy, nil
}

// GetRunningWorkflowByHash finds a running workflow with the same input hash
func (s *InMemoryWorkflowStore) GetRunningWorkflowByHash(ctx context.Context, workflowName, inputHash string) (*WorkflowStatus, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	// Check if there's a running workflow with this hash
	workflowID, exists := s.hashIndex[inputHash]
	if !exists {
		return nil, nil // No running workflow found
	}

	// Get the workflow status
	status, exists := s.workflows[workflowID]
	if !exists {
		// Clean up orphaned hash index entry
		delete(s.hashIndex, inputHash)
		return nil, nil
	}

	// Only return if it's still running
	if status.Status == "running" {
		statusCopy := *status
		return &statusCopy, nil
	}

	// Clean up hash index for non-running workflow
	delete(s.hashIndex, inputHash)
	return nil, nil
}

// MarkWorkflowCompleted marks a workflow as completed
func (s *InMemoryWorkflowStore) MarkWorkflowCompleted(ctx context.Context, workflowID string, finalContext WorkflowContext) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	status, exists := s.workflows[workflowID]
	if !exists {
		return fmt.Errorf("workflow not found: %s", workflowID)
	}

	// Update status
	now := time.Now()
	status.Status = "completed"
	status.Context = finalContext
	status.CompletedAt = &now
	status.ErrorMessage = ""

	// Remove from hash index since it's no longer running
	hashKey := s.makeHashKey(status.WorkflowName, status.Input)
	if hashKey != "" {
		delete(s.hashIndex, hashKey)
	}

	return nil
}

// MarkWorkflowFailed marks a workflow as failed
func (s *InMemoryWorkflowStore) MarkWorkflowFailed(ctx context.Context, workflowID string, errorMessage string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	status, exists := s.workflows[workflowID]
	if !exists {
		return fmt.Errorf("workflow not found: %s", workflowID)
	}

	// Update status
	now := time.Now()
	status.Status = "failed"
	status.CompletedAt = &now
	status.ErrorMessage = errorMessage

	// Remove from hash index since it's no longer running
	hashKey := s.makeHashKey(status.WorkflowName, status.Input)
	if hashKey != "" {
		delete(s.hashIndex, hashKey)
	}

	return nil
}

// makeHashKey creates a hash key for the given workflow name and input
func (s *InMemoryWorkflowStore) makeHashKey(workflowName string, input map[string]interface{}) string {
	hash, err := GenerateInputHash(workflowName, input)
	if err != nil {
		return ""
	}
	return hash
}
