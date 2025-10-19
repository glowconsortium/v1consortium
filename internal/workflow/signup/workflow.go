package signup

import (
	"time"
	"v1consortium/internal/pkg/riverjobs"
)

// SignupInput represents the input for the signup workflow
type SignupInput struct {
	Email            string                 `json:"email" v:"required|email"`
	Password         string                 `json:"password" v:"required|min-length:8"`
	FirstName        string                 `json:"first_name" v:"required|min-length:2"`
	LastName         string                 `json:"last_name" v:"required|min-length:2"`
	Role             string                 `json:"role" v:"required|in:driver,admin,owner"`
	OrganizationData map[string]interface{} `json:"organization_data"`
	Metadata         map[string]interface{} `json:"metadata"`
}

// NewSignupWorkflowDefinition creates the workflow definition for user signup
func NewSignupWorkflowDefinition() *riverjobs.WorkflowDefinition {
	return &riverjobs.WorkflowDefinition{
		Name:        "user_signup",
		Description: "Complete user signup process with organization setup",
		Steps: []riverjobs.StepDefinition{
			{
				Name:       "validate",
				Queue:      "default",
				Timeout:    2 * time.Minute,
				MaxRetries: 3,
				NextSteps:  []string{"create_user"},
				Required:   true,
			},
			{
				Name:       "create_user",
				Queue:      "default",
				Timeout:    3 * time.Minute,
				MaxRetries: 3,
				NextSteps:  []string{"create_organization"},
				Required:   true,
			},
			{
				Name:       "create_organization",
				Queue:      "default",
				Timeout:    3 * time.Minute,
				MaxRetries: 3,
				NextSteps:  []string{"setup_stripe"},
				Required:   true,
			},
			{
				Name:       "setup_stripe",
				Queue:      "external",
				Timeout:    5 * time.Minute,
				MaxRetries: 5,
				NextSteps:  []string{"send_verification"},
				Required:   false, // Can continue without stripe setup
			},
			{
				Name:       "send_verification",
				Queue:      "notification",
				Timeout:    2 * time.Minute,
				MaxRetries: 3,
				NextSteps:  []string{}, // Final step
				Required:   false,      // Can complete without email verification
			},
		},
		Config: riverjobs.WorkflowConfig{
			MaxRetries:     3,
			DefaultQueue:   "default",
			DefaultTimeout: 5 * time.Minute,
		},
	}
}
