package signupv2

import (
	"context"
	"fmt"
	"v1consortium/internal/pkg/riverjobsv2"
)

// Example usage of the simplified workflow pattern

// StartSignupWorkflow demonstrates how to start a signup workflow with deduplication
func StartSignupWorkflow(ctx context.Context, riverClient interface{}, store riverjobsv2.WorkflowStore, signupData SignupInput) (*riverjobsv2.StartWorkflowResult, error) {
	// Convert signup data to map for workflow input
	input := map[string]interface{}{
		"email":             signupData.Email,
		"password":          signupData.Password,
		"first_name":        signupData.FirstName,
		"last_name":         signupData.LastName,
		"role":              signupData.Role,
		"organization_data": signupData.OrganizationData,
		"metadata":          signupData.Metadata,
	}

	// Start the workflow with automatic UUID generation and deduplication
	result, err := riverjobsv2.StartWorkflow(ctx, riverClient, store, "signup", input)
	if err != nil {
		return nil, fmt.Errorf("failed to start signup workflow: %w", err)
	}

	return result, nil
}

// SetupSignupWorkflowExecutor demonstrates how to set up the workflow executor
func SetupSignupWorkflowExecutor(store riverjobsv2.WorkflowStore) *riverjobsv2.WorkflowExecutor {
	// Create the executor with store
	executor := riverjobsv2.NewWorkflowExecutor(store)

	// Register the signup workflow
	signupWorkflow := NewSignupWorkflow()
	executor.RegisterWorkflow(signupWorkflow)

	return executor
}

// Example of how this would be integrated into cmd_river_worker.go
func ExampleIntegration() {
	// This shows how simple the registration becomes:

	// 1. Create workflow store (in real usage, pass actual store)
	store := riverjobsv2.NewInMemoryWorkflowStore() // or NewDatabaseWorkflowStore(dbPool)

	// 2. Create workflow executor
	_ = SetupSignupWorkflowExecutor(store)

	// 3. Register with River (pseudo-code)
	// river.AddWorker(workers, workflowExecutor)

	// That's it! No more complex step workers, just one executor per workflow type

	fmt.Println("Workflow executor setup complete")
} // Benefits of this approach:
// 1. Everything for a workflow is in one file
// 2. Steps are defined with their execution functions right together
// 3. Flow control is explicit and easy to understand
// 4. Retry logic is per-step and configurable
// 5. Optional steps are clearly marked
// 6. Context flows naturally between steps
// 7. Easy to test individual steps
// 8. Easy to add new workflows following the same pattern
