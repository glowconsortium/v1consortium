-- Migration: Adding Workflow Models for River Jobs
-- Created: 2025-10-16
-- Purpose: Create tables for workflow execution tracking and River job queue system

-- Enable necessary extensions
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Workflow executions table - tracks high-level workflow state
CREATE TABLE workflow_executions (
    workflow_id VARCHAR(255) PRIMARY KEY DEFAULT uuid_generate_v4()::text,
    workflow_type VARCHAR(100) NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'pending',
    context JSONB DEFAULT '{}',
    current_step VARCHAR(100),
    started_at TIMESTAMP WITH TIME ZONE,
    completed_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),

    args_hash VARCHAR(64),
    -- Organization and user tracking
    org_id VARCHAR(255),
    user_id UUID REFERENCES auth.users(id),
    
    -- Error tracking
    error_message TEXT,
    retry_count INTEGER DEFAULT 0,
    
    -- Constraints
    CONSTRAINT valid_status CHECK (status IN ('pending', 'running', 'completed', 'failed', 'cancelled'))
);

-- Workflow steps table - tracks individual job executions within workflows
CREATE TABLE workflow_steps (
    step_id VARCHAR(255) PRIMARY KEY DEFAULT uuid_generate_v4()::text,
    workflow_id VARCHAR(255) NOT NULL REFERENCES workflow_executions(workflow_id) ON DELETE CASCADE,
    step_name VARCHAR(100) NOT NULL,
    step_order INTEGER DEFAULT 0,
    status VARCHAR(50) NOT NULL DEFAULT 'pending',
    
    -- Job data
    input_data JSONB DEFAULT '{}',
    output_data JSONB DEFAULT '{}',
    
    -- Timing
    started_at TIMESTAMP WITH TIME ZONE,
    completed_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    
    -- Error handling
    error_message TEXT,
    retry_count INTEGER DEFAULT 0,
    max_retries INTEGER DEFAULT 3,
    
    -- River job tracking
    river_job_id BIGINT,
    queue_name VARCHAR(100),
    
    -- Constraints
    CONSTRAINT valid_step_status CHECK (status IN ('pending', 'running', 'completed', 'failed', 'skipped'))
);

-- Workflow templates table - defines reusable workflow patterns
CREATE TABLE workflow_templates (
    template_id VARCHAR(255) PRIMARY KEY DEFAULT uuid_generate_v4()::text,
    template_name VARCHAR(100) NOT NULL UNIQUE,
    description TEXT,
    workflow_type VARCHAR(100) NOT NULL,
    
    -- Template definition
    steps_config JSONB NOT NULL, -- Array of step definitions
    default_context JSONB DEFAULT '{}',
    
    -- Settings
    is_active BOOLEAN DEFAULT true,
    timeout_minutes INTEGER DEFAULT 60,
    max_retries INTEGER DEFAULT 3,
    
    -- Metadata
    created_by UUID REFERENCES auth.users(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    version INTEGER DEFAULT 1
);

-- Workflow metrics table - for performance tracking and analytics
CREATE TABLE workflow_metrics (
    metric_id VARCHAR(255) PRIMARY KEY DEFAULT uuid_generate_v4()::text,
    workflow_id VARCHAR(255) REFERENCES workflow_executions(workflow_id) ON DELETE CASCADE,
    workflow_type VARCHAR(100) NOT NULL,
    
    -- Performance metrics
    total_duration_ms BIGINT,
    step_count INTEGER,
    retry_count INTEGER,
    
    -- Timestamps
    recorded_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    
    -- Outcome
    final_status VARCHAR(50),
    
    -- Organization tracking
    org_id VARCHAR(255)
);

-- Create indexes for performance
CREATE INDEX idx_workflow_executions_status ON workflow_executions(status);
CREATE INDEX idx_workflow_executions_type ON workflow_executions(workflow_type);
CREATE INDEX idx_workflow_executions_org_id ON workflow_executions(org_id);
CREATE INDEX idx_workflow_executions_user_id ON workflow_executions(user_id);
CREATE INDEX idx_workflow_executions_created_at ON workflow_executions(created_at);
CREATE INDEX idx_workflow_executions_status_created ON workflow_executions(status, created_at);

CREATE INDEX idx_workflow_steps_workflow_id ON workflow_steps(workflow_id);
CREATE INDEX idx_workflow_steps_status ON workflow_steps(status);
CREATE INDEX idx_workflow_steps_step_name ON workflow_steps(step_name);
CREATE INDEX idx_workflow_steps_river_job_id ON workflow_steps(river_job_id);
CREATE INDEX idx_workflow_steps_workflow_step_order ON workflow_steps(workflow_id, step_order);

CREATE INDEX idx_workflow_templates_type ON workflow_templates(workflow_type);
CREATE INDEX idx_workflow_templates_active ON workflow_templates(is_active);

CREATE INDEX idx_workflow_metrics_type ON workflow_metrics(workflow_type);
CREATE INDEX idx_workflow_metrics_org_id ON workflow_metrics(org_id);
CREATE INDEX idx_workflow_metrics_recorded_at ON workflow_metrics(recorded_at);

-- Create updated_at trigger function
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Add updated_at triggers
CREATE TRIGGER update_workflow_executions_updated_at
    BEFORE UPDATE ON workflow_executions
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_workflow_steps_updated_at
    BEFORE UPDATE ON workflow_steps
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_workflow_templates_updated_at
    BEFORE UPDATE ON workflow_templates
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- Row Level Security (RLS) policies
ALTER TABLE workflow_executions ENABLE ROW LEVEL SECURITY;
ALTER TABLE workflow_steps ENABLE ROW LEVEL SECURITY;
ALTER TABLE workflow_templates ENABLE ROW LEVEL SECURITY;
ALTER TABLE workflow_metrics ENABLE ROW LEVEL SECURITY;

-- Policies for workflow_executions
CREATE POLICY "Users can view their own workflow executions" ON workflow_executions
    FOR SELECT USING (auth.uid() = user_id);

CREATE POLICY "Users can insert their own workflow executions" ON workflow_executions
    FOR INSERT WITH CHECK (auth.uid() = user_id);

CREATE POLICY "Users can update their own workflow executions" ON workflow_executions
    FOR UPDATE USING (auth.uid() = user_id);

-- Service role can access all workflow data (for background jobs)
CREATE POLICY "Service role full access to workflow executions" ON workflow_executions
    FOR ALL USING (auth.jwt() ->> 'role' = 'service_role');

-- Policies for workflow_steps
CREATE POLICY "Users can view steps for their workflows" ON workflow_steps
    FOR SELECT USING (
        EXISTS (
            SELECT 1 FROM workflow_executions 
            WHERE workflow_executions.workflow_id = workflow_steps.workflow_id 
            AND workflow_executions.user_id = auth.uid()
        )
    );

CREATE POLICY "Service role full access to workflow steps" ON workflow_steps
    FOR ALL USING (auth.jwt() ->> 'role' = 'service_role');

-- Policies for workflow_templates (admin only for creation/updates)
CREATE POLICY "Anyone can view active workflow templates" ON workflow_templates
    FOR SELECT USING (is_active = true);

CREATE POLICY "Service role full access to workflow templates" ON workflow_templates
    FOR ALL USING (auth.jwt() ->> 'role' = 'service_role');

-- Policies for workflow_metrics (read-only for users, full access for service)
CREATE POLICY "Users can view metrics for their workflows" ON workflow_metrics
    FOR SELECT USING (
        EXISTS (
            SELECT 1 FROM workflow_executions 
            WHERE workflow_executions.workflow_id = workflow_metrics.workflow_id 
            AND workflow_executions.user_id = auth.uid()
        )
    );

CREATE POLICY "Service role full access to workflow metrics" ON workflow_metrics
    FOR ALL USING (auth.jwt() ->> 'role' = 'service_role');

-- Insert default workflow templates
INSERT INTO workflow_templates (template_name, workflow_type, description, steps_config) VALUES
(
    'User Signup Flow',
    'user_signup',
    'Complete user registration with organization and subscription setup',
    '[
        {"step": "user_signup", "worker": "UserSignupWorker", "queue": "default", "timeout_minutes": 5},
        {"step": "create_organization", "worker": "CreateOrganizationWorker", "queue": "default", "timeout_minutes": 5},
        {"step": "process_subscription", "worker": "ProcessSubscriptionWorker", "queue": "critical", "timeout_minutes": 10}
    ]'::jsonb
),
(
    'Drug Test Order Flow',
    'drug_test_order',
    'Order drug test and send notifications',
    '[
        {"step": "order_test", "worker": "DrugTestOrderWorker", "queue": "external", "timeout_minutes": 10},
        {"step": "send_notification", "worker": "SendTestNotificationWorker", "queue": "notification", "timeout_minutes": 3}
    ]'::jsonb
),
(
    'Background Check Flow',
    'background_check',
    'Order background check and process results',
    '[
        {"step": "order_background_check", "worker": "BackgroundCheckWorker", "queue": "external", "timeout_minutes": 15},
        {"step": "process_results", "worker": "ProcessBackgroundCheckWorker", "queue": "default", "timeout_minutes": 5},
        {"step": "send_notification", "worker": "SendNotificationWorker", "queue": "notification", "timeout_minutes": 3}
    ]'::jsonb
);

-- Create views for common queries
CREATE VIEW workflow_summary AS
SELECT 
    we.workflow_id,
    we.workflow_type,
    we.status,
    we.created_at,
    we.started_at,
    we.completed_at,
    we.org_id,
    we.user_id,
    COUNT(ws.step_id) as total_steps,
    COUNT(CASE WHEN ws.status = 'completed' THEN 1 END) as completed_steps,
    COUNT(CASE WHEN ws.status = 'failed' THEN 1 END) as failed_steps,
    COALESCE(we.completed_at, NOW()) - we.started_at as duration
FROM workflow_executions we
LEFT JOIN workflow_steps ws ON we.workflow_id = ws.workflow_id
GROUP BY we.workflow_id, we.workflow_type, we.status, we.created_at, we.started_at, we.completed_at, we.org_id, we.user_id;

-- Grant permissions on the view
GRANT SELECT ON workflow_summary TO authenticated;

-- Comments for documentation
COMMENT ON TABLE workflow_executions IS 'Tracks high-level workflow execution state for River job workflows';
COMMENT ON TABLE workflow_steps IS 'Tracks individual job executions within workflows';
COMMENT ON TABLE workflow_templates IS 'Reusable workflow pattern definitions';
COMMENT ON TABLE workflow_metrics IS 'Performance and analytics data for workflow executions';
COMMENT ON VIEW workflow_summary IS 'Aggregated view of workflow execution progress and timing';



-- Add args_hash column to workflow_executions table for duplicate prevention
-- This migration should be run to support the workflow deduplication feature



-- Create index for efficient duplicate checking
CREATE INDEX IF NOT EXISTS idx_workflow_executions_dedup 
ON workflow_executions(workflow_type, org_id, args_hash) 
WHERE status IN ('pending', 'running');

-- Add comment to document the purpose
COMMENT ON COLUMN workflow_executions.args_hash IS 'SHA256 hash of job arguments for duplicate prevention';