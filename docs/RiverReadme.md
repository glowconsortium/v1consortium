# River Jobs Implementation - V1 Consortium

## Overview

V1 Consortium uses [River](https://github.com/riverqueue/river) as its job queue system for reliable, scalable background job processing. River provides robust job scheduling, retries, and monitoring capabilities with PostgreSQL as the backend.

## Architecture

### Core Components

1. **River Manager** (`riverjobs/client.go`) - Manages River client instances and configuration
2. **Workflow Manager** (`riverjobs/workflow.go`) - Handles workflow state tracking and persistence
3. **Workflow Orchestrator** (`riverjobs/orchestrator.go`) - High-level workflow coordination
4. **Worker Base** (`riverjobs/worker.go`) - Common functionality for all workers
5. **Example Workers** (`riverjobs/examples.go`) - V1 Consortium-specific workflow implementations

### Queue Structure

The system uses multiple queues for job prioritization and organization:

- **`default`** - Standard workflow jobs (100 workers)
- **`critical`** - High-priority jobs requiring immediate processing (50 workers)
- **`scheduled`** - Time-delayed or scheduled jobs (25 workers)
- **`notification`** - Email, SMS, and notification jobs (25 workers)
- **`external`** - Third-party API integrations (10 workers)

## Workflow System

### Workflow State Management

Each workflow execution is tracked with:
- **Workflow ID** - Unique identifier for the workflow instance
- **Workflow Type** - The type of workflow (e.g., "user_signup", "drug_test_order")
- **Status** - Current state (pending, running, completed, failed)
- **Context** - JSON data passed between workflow steps
- **Steps** - Individual job executions within the workflow

### Workflow Lifecycle

1. **Start Workflow** - Create workflow execution record
2. **Execute Steps** - Run individual jobs with state tracking
3. **Update Context** - Pass data between steps
4. **Handle Errors** - Retry failed steps with exponential backoff
5. **Complete Workflow** - Mark as completed or failed

## Example Workflows

### User Signup Workflow

A multi-step process for user registration and organization setup:

```
1. UserSignupWorker
   ├── Validate signup data
   ├── Create user account (Supabase)
   └── Enqueue CreateOrganizationWorker

2. CreateOrganizationWorker
   ├── Create organization record
   ├── Set up organization defaults
   └── Enqueue ProcessSubscriptionWorker

3. ProcessSubscriptionWorker
   ├── Process payment with Stripe
   ├── Create subscription record
   └── Complete workflow
```

### Drug Testing Workflow

Handles drug test ordering and notifications:

```
1. DrugTestOrderWorker
   ├── Validate test request
   ├── Select optimal facility
   ├── Create external order (Quest Diagnostics)
   └── Enqueue SendTestNotificationWorker

2. SendTestNotificationWorker
   ├── Send notifications to employee
   ├── Send notifications to supervisor
   └── Complete workflow
```

## Worker Implementation

### Creating a New Worker

1. **Define Job Arguments** - Implement `JobArgs` interface:
```go
type MyJobArgs struct {
    BaseJobArgs
    CustomField string `json:"custom_field"`
}

func (args MyJobArgs) Kind() string { return "my_job" }
```

2. **Implement Worker** - Create worker struct with required methods:
```go
type MyWorker struct {
    WorkerBase
}

func (w *MyWorker) Work(ctx context.Context, job *river.Job[MyJobArgs]) error {
    return w.ExecuteWithWorkflowTracking(ctx, job.Args.WorkflowID, job.Args.StepName, func(ctx context.Context) (map[string]interface{}, error) {
        // Job implementation here
        return map[string]interface{}{"result": "success"}, nil
    })
}

func (w *MyWorker) Middleware(job *rivertype.JobRow) []rivertype.WorkerMiddleware {
    return []rivertype.WorkerMiddleware{}
}

func (w *MyWorker) NextRetry(job *river.Job[MyJobArgs]) time.Time {
    return time.Now().Add(time.Duration(job.Attempt) * time.Minute)
}

func (w *MyWorker) Timeout(job *river.Job[MyJobArgs]) time.Duration {
    return 5 * time.Minute
}
```

3. **Register Worker** - Add to `registerRiverWorkers` function:
```go
myWorker := &riverjobs.MyWorker{WorkerBase: workerBase}
river.AddWorker[riverjobs.MyJobArgs](riverManager.Workers, myWorker)
```

## Running the Worker

### Database Setup

Before running the River worker, ensure the database is properly set up:

1. **Install River CLI**:
```bash
go install github.com/riverqueue/river/cmd/river@latest
```

2. **Run Supabase migrations** (includes workflow tables):
```bash
supabase db push
```

3. **Run River migrations** (for job queue tables):
```bash
river migrate-up --database-url "$DATABASE_URL"
```

### Start the River Worker Command

```bash
# Development
go run main.go river_worker

# Production (with compiled binary)
./v1consortium river_worker
```

### Environment Variables

Set the following environment variables:

```bash
export DATABASE_URL="postgres://user:password@localhost:5432/v1consortium"
# or
export DB_HOST="localhost"
export DB_PORT="5432"
export DB_USER="username"
export DB_PASS="password"
export DB_NAME="v1consortium"
```

## Monitoring and Maintenance

### Background Jobs

The worker automatically runs background monitoring tasks:

- **Stuck Workflow Monitor** (every 5 minutes) - Identifies workflows running longer than 1 hour
- **Workflow Cleanup** (daily at 3 AM) - Removes completed workflows older than 30 days
- **Metrics Collection** (every 10 minutes) - Gathers workflow statistics
- **Health Checks** (every minute) - Basic system health monitoring

### Logging

All job executions are logged with:
- Job start/completion times
- Workflow and step context
- Error details for failed jobs
- Performance metrics

### Database Schema

The system requires these database tables:

```sql
-- Workflow executions tracking
CREATE TABLE workflow_executions (
    workflow_id VARCHAR PRIMARY KEY,
    workflow_type VARCHAR NOT NULL,
    status VARCHAR NOT NULL,
    context JSONB,
    started_at TIMESTAMP,
    completed_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Individual workflow steps
CREATE TABLE workflow_steps (
    step_id VARCHAR PRIMARY KEY,
    workflow_id VARCHAR REFERENCES workflow_executions(workflow_id),
    step_name VARCHAR NOT NULL,
    status VARCHAR NOT NULL,
    input_data JSONB,
    output_data JSONB,
    error_message TEXT,
    started_at TIMESTAMP,
    completed_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW()
);

-- River job tables (created automatically by River)
-- river_job, river_leader, river_migration
```

## Configuration

### Queue Configuration

Modify queue settings in the River worker command:

```go
config := &riverjobs.Config{
    DatabaseURL: dbURL,
    Queues: map[string]river.QueueConfig{
        river.QueueDefault:          {MaxWorkers: 100},
        riverjobs.QueueCritical:     {MaxWorkers: 50},
        riverjobs.QueueScheduled:    {MaxWorkers: 25},
        riverjobs.QueueNotification: {MaxWorkers: 25},
        riverjobs.QueueExternal:     {MaxWorkers: 10},
    },
    PollInterval: 5 * time.Second,
}
```

### Worker Timeouts and Retries

Configure per-worker settings:

```go
// Custom timeout
func (w *MyWorker) Timeout(job *river.Job[MyJobArgs]) time.Duration {
    return 10 * time.Minute
}

// Custom retry logic
func (w *MyWorker) NextRetry(job *river.Job[MyJobArgs]) time.Time {
    // Exponential backoff: 1min, 2min, 4min, 8min, etc.
    backoff := time.Duration(1<<job.Attempt) * time.Minute
    return time.Now().Add(backoff)
}
```

## API Usage

### Starting a Workflow

```go
// Initialize orchestrator
orchestrator := riverjobs.NewWorkflowOrchestrator(riverManager, workflowManager, logger)

// Start user signup workflow
workflowID, err := orchestrator.StartWorkflow(ctx, "user_signup", map[string]interface{}{
    "email": "user@example.com",
    "organization_name": "Acme Corp",
})
```

### Enqueuing Jobs

```go
// Enqueue a job with custom options
args := riverjobs.DrugTestOrderArgs{
    BaseJobArgs: riverjobs.BaseJobArgs{
        WorkflowID:   workflowID,
        WorkflowType: "drug_test_order", 
        StepName:     "order_test",
        OrgID:        "org_123",
        UserID:       "user_456",
    },
    EmployeeID: "emp_789",
    TestType:   "pre_employment",
    RushOrder:  true,
}

result, err := riverClient.Insert(ctx, args, &river.InsertOpts{
    Queue:    riverjobs.QueueCritical,
    Priority: 2,
})
```

### Checking Workflow Status

```go
// Get workflow execution details
execution, err := workflowManager.GetWorkflowExecution(ctx, workflowID)

// Get workflow summary with steps
summary, err := orchestrator.GetWorkflowSummary(ctx, workflowID)
```

## Development

### Testing

```bash
# Run unit tests
go test ./internal/pkg/riverjobs/...

# Run with coverage
go test -cover ./internal/pkg/riverjobs/...
```

### Adding New Workflows

1. Define job argument types in `examples.go`
2. Implement worker structs with required methods
3. Register workers in `cmd_river_worker.go`
4. Update this documentation

## Production Deployment

### Scaling

- Run multiple worker instances for horizontal scaling
- Use different queue configurations per instance
- Monitor queue depths and worker utilization

### Monitoring

- Set up alerts for stuck workflows
- Monitor job failure rates
- Track workflow completion times
- Use River's built-in metrics for job queue health

### Backup and Recovery

- Regular PostgreSQL backups include job queue state
- Failed jobs can be retried manually through River's API
- Workflow state is preserved and recoverable

## Troubleshooting

### Common Issues

1. **Workers Not Starting**: Check database connection and environment variables
2. **Jobs Not Processing**: Verify worker registration and queue configuration
3. **High Memory Usage**: Reduce MaxWorkers or optimize job logic
4. **Stuck Workflows**: Check background monitoring logs for details

### Debug Mode

Enable debug logging by setting log level to debug in the worker configuration.

---

For more information, see:
- [River Documentation](https://riverqueue.com/docs)
- [GoFrame Documentation](https://goframe.org/)
- [V1 Consortium Main Documentation](../README.MD)