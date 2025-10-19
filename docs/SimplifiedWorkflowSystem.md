# Simplified Workflow System - V1 Consortium

## Overview

The workflow system has been redesigned to eliminate complexity and reduce boilerplate code. The new system uses a **choreography** approach where each step is self-contained and knows how to enqueue the next step.

## Key Components

### 1. SimpleWorkflowManager (`workflow_manager_simple.go`)
- **Purpose**: Manages workflow lifecycle and state
- **Responsibilities**:
  - Register workflow definitions
  - Start workflows by enqueueing the first step
  - Track workflow state in database
  - Handle workflow completion/failure

### 2. StepArgs (`types.go`)
- **Purpose**: Standardized job arguments for all workflow steps
- **Benefits**: 
  - Single args type instead of multiple step-specific types
  - Contains both workflow-level and step-specific data
  - Simplifies worker implementation

```go
type StepArgs struct {
    WorkflowID    string                 `json:"workflow_id"`
    WorkflowType  string                 `json:"workflow_type"`
    StepName      string                 `json:"step_name"`
    OrgID         string                 `json:"org_id"`
    UserID        string                 `json:"user_id,omitempty"`
    WorkflowInput map[string]interface{} `json:"workflow_input"` // All workflow data
    StepInput     map[string]interface{} `json:"step_input"`     // Step-specific data
}
```

### 3. BaseStepWorker (`step_worker.go`)
- **Purpose**: Common functionality for all step workers
- **Features**:
  - Implements River worker interface
  - Handles step execution orchestration
  - Updates workflow state
  - Enqueues next steps automatically

### 4. Step Workers (`workflow/signup/steps/`)
- **Purpose**: Individual step implementations
- **Pattern**: Each worker implements `StepWorker` interface with `Execute()` method
- **Self-contained**: Each step handles its own business logic and returns output for next steps

### 5. Workflow Definitions (`workflow/signup/workflow.go`)
- **Purpose**: Simple data structures defining workflow steps and configuration
- **No Logic**: Pure configuration, no execution logic

## How It Works

### 1. Workflow Start
```go
// Start workflow
workflowID, err := workflowManager.StartWorkflow(ctx, "user_signup", input, orgID, userID)
```

### 2. Step Execution Flow
```
1. SimpleWorkflowManager.StartWorkflow()
   ↓
2. Enqueue first step job with StepArgs
   ↓
3. BaseStepWorker.Work() 
   ↓
4. Concrete step's Execute() method
   ↓
5. BaseStepWorker updates workflow state
   ↓
6. BaseStepWorker enqueues next steps
   ↓
7. Process repeats until no more steps
```

### 3. Self-Contained Steps
Each step worker:
- Receives `StepArgs` with all workflow data
- Executes its business logic
- Returns output data for next steps
- Next steps are automatically enqueued by `BaseStepWorker`

## Benefits of New Approach

### ✅ Eliminated Complexity
- **Removed**: Complex orchestrator with generic step execution
- **Removed**: Multiple job args types (ValidateStepArgs, CreateUserStepArgs, etc.)
- **Removed**: Legacy worker wrapper layer (usersignup_new.go)

### ✅ Reduced Boilerplate
- Single `StepArgs` type for all steps
- Common worker base handles River interface requirements
- Simple workflow definitions without execution logic

### ✅ Better Separation of Concerns
- **SimpleWorkflowManager**: Workflow lifecycle
- **Step Workers**: Business logic only
- **Workflow Definitions**: Configuration only

### ✅ Easier to Understand
- Clear data flow from step to step
- No complex orchestration logic
- Each component has single responsibility

## Example Usage

### Starting a Workflow
```go
func StartUserSignup(ctx context.Context, email, password string) (string, error) {
    input := map[string]interface{}{
        "input": map[string]interface{}{
            "email":      email,
            "password":   password,
            "first_name": "John",
            "last_name":  "Doe",
            "role":       "employee",
        },
    }
    
    return workflowManager.StartWorkflow(ctx, "user_signup", input, orgID, userID)
}
```

### Implementing a Step Worker
```go
func (w *ValidateStepWorker) Execute(ctx context.Context, args riverjobs.StepArgs) (map[string]interface{}, error) {
    // Get input data
    inputData := args.WorkflowInput["input"].(map[string]interface{})
    
    // Validate business logic
    if err := validateEmail(inputData["email"].(string)); err != nil {
        return nil, err
    }
    
    // Return output for next steps
    return map[string]interface{}{
        "validated_input": inputData,
    }, nil
}
```

## Migration Notes

### Removed Files
- `internal/logic/riverjobUsersignup/usersignup_new.go` - Legacy wrapper worker
- `internal/pkg/riverjobs/orchestrator.go` - Complex orchestrator

### Updated Files
- `cmd_river_worker.go` - Simplified worker registration
- All step workers now use `StepArgs` instead of specific types

### New Architecture
The system now follows a **choreography pattern** where:
- Each step knows what comes next (defined in workflow definition)
- Steps are autonomous and self-contained
- No central orchestrator managing step execution
- Workflow state is tracked in database

This approach is simpler, more maintainable, and easier to extend with new workflows and steps.