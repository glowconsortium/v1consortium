package riverjobs

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/riverqueue/river"
	"github.com/riverqueue/river/riverdriver/riverpgxv5"
	"github.com/riverqueue/river/rivertype"
)

// RiverManager manages River clients and workers
type RiverManager struct {
	Client  *river.Client[pgx.Tx]
	Workers *river.Workers
	Config  *Config
	logger  *slog.Logger
}

// NewRiverManager creates a new River manager
func NewRiverManager(ctx context.Context, config *Config, logger *slog.Logger) (*RiverManager, error) {
	// Initialize database connection
	dbPool, err := pgxpool.New(ctx, config.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Initialize workers registry
	workers := river.NewWorkers()

	// Default queue configuration if none provided
	if config.Queues == nil {
		config.Queues = map[string]river.QueueConfig{
			river.QueueDefault: {MaxWorkers: 100},
			QueueCritical:      {MaxWorkers: 50},
			QueueScheduled:     {MaxWorkers: 25},
			QueueNotification:  {MaxWorkers: 25},
			QueueExternal:      {MaxWorkers: 10},
		}
	}

	// Create River client
	riverClient, err := river.NewClient(riverpgxv5.New(dbPool), &river.Config{
		Queues:  config.Queues,
		Workers: workers,
		Logger:  logger,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create River client: %w", err)
	}

	manager := &RiverManager{
		Client:  riverClient,
		Workers: workers,
		Config:  config,
		logger:  logger,
	}

	return manager, nil
}

// RegisterWorker registers a worker for a specific job type
func (rm *RiverManager) RegisterWorker(worker river.Worker[river.JobArgs]) {
	river.AddWorker(rm.Workers, worker)
}

// Start starts the River job processing
func (rm *RiverManager) Start(ctx context.Context) error {
	return rm.Client.Start(ctx)
}

// Stop stops the River job processing
func (rm *RiverManager) Stop(ctx context.Context) error {
	return rm.Client.Stop(ctx)
}

// Insert inserts a job
func (rm *RiverManager) Insert(ctx context.Context, args river.JobArgs, opts *river.InsertOpts) (*rivertype.JobInsertResult, error) {
	return rm.Client.Insert(ctx, args, opts)
}

// InsertTx inserts a job in a transaction
func (rm *RiverManager) InsertTx(ctx context.Context, tx pgx.Tx, args river.JobArgs, opts *river.InsertOpts) (*rivertype.JobInsertResult, error) {
	return rm.Client.InsertTx(ctx, tx, args, opts)
}

// JobClient interface for abstraction over River client for RPC/HTTP
type JobClient interface {
	Insert(ctx context.Context, args river.JobArgs, opts *river.InsertOpts) (*rivertype.JobInsertResult, error)
	InsertTx(ctx context.Context, tx pgx.Tx, args river.JobArgs, opts *river.InsertOpts) (*rivertype.JobInsertResult, error)
	Stop(ctx context.Context) error
}

// HTTPJobClient wraps River client for HTTP usage
type HTTPJobClient struct {
	client *river.Client[pgx.Tx]
}

// NewHTTPJobClient creates a new HTTP job client
func NewHTTPJobClient(client *river.Client[pgx.Tx]) *HTTPJobClient {
	return &HTTPJobClient{client: client}
}

func (h *HTTPJobClient) Insert(ctx context.Context, args river.JobArgs, opts *river.InsertOpts) (*rivertype.JobInsertResult, error) {
	return h.client.Insert(ctx, args, opts)
}

func (h *HTTPJobClient) InsertTx(ctx context.Context, tx pgx.Tx, args river.JobArgs, opts *river.InsertOpts) (*rivertype.JobInsertResult, error) {
	return h.client.InsertTx(ctx, tx, args, opts)
}

func (h *HTTPJobClient) Stop(ctx context.Context) error {
	return h.client.Stop(ctx)
}

// RPCJobClient interface for RPC implementations
type RPCJobClient interface {
	JobClient
	// Additional RPC-specific methods can be added here
}

// ClientManager manages both HTTP and RPC clients
type ClientManager struct {
	HTTPClient *HTTPJobClient
	RPCClient  RPCJobClient
	logger     *slog.Logger
}

// NewClientManager creates a new client manager
func NewClientManager(riverClient *river.Client[pgx.Tx], logger *slog.Logger) *ClientManager {
	return &ClientManager{
		HTTPClient: NewHTTPJobClient(riverClient),
		logger:     logger,
	}
}

// SetRPCClient sets the RPC client
func (cm *ClientManager) SetRPCClient(client RPCJobClient) {
	cm.RPCClient = client
}

// GetClient returns the appropriate client based on the protocol
func (cm *ClientManager) GetClient(protocol string) JobClient {
	switch protocol {
	case "rpc", "grpc":
		if cm.RPCClient != nil {
			return cm.RPCClient
		}
		cm.logger.Warn("RPC client requested but not available, falling back to HTTP client")
		fallthrough
	default:
		return cm.HTTPClient
	}
}
