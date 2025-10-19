package dbosworkflow

import (
	"context"
	"database/sql"
)

type OnboardingTransactions struct {
	DB *sql.DB
}

// CreateUser creates a new user record transactionally
func (t *OnboardingTransactions) CreateUser(ctx context.Context, user User) (string, error) {
	var userID string
	err := t.DB.QueryRowContext(ctx,
		`INSERT INTO users (email, password_hash, first_name, last_name, provider, provider_id, email_verified, status)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		 RETURNING id`,
		user.Email, user.PasswordHash, user.FirstName, user.LastName,
		user.Provider, user.ProviderID, user.EmailVerified, user.Status,
	).Scan(&userID)
	return userID, err
}

// CreateCompany creates a new company record transactionally
func (t *OnboardingTransactions) CreateCompany(ctx context.Context, company Company) (string, error) {
	var companyID string
	err := t.DB.QueryRowContext(ctx,
		`INSERT INTO companies (name, size, industry, owner_id)
		 VALUES ($1, $2, $3, $4)
		 RETURNING id`,
		company.Name, company.Size, company.Industry, company.OwnerID,
	).Scan(&companyID)
	return companyID, err
}

// LinkUserToCompany creates the association between user and company
func (t *OnboardingTransactions) LinkUserToCompany(ctx context.Context, userID, companyID string) error {
	_, err := t.DB.ExecContext(ctx,
		`INSERT INTO user_companies (user_id, company_id, role)
		 VALUES ($1, $2, $3)`,
		userID, companyID, "owner",
	)
	return err
}

// UpdateUserEmailVerified marks user email as verified
func (t *OnboardingTransactions) UpdateUserEmailVerified(ctx context.Context, userID string) error {
	_, err := t.DB.ExecContext(ctx,
		`UPDATE users SET email_verified = true, status = $1, updated_at = NOW()
		 WHERE id = $2`,
		StatusEmailVerified, userID,
	)
	return err
}

// CreateSubscription creates a subscription record transactionally
func (t *OnboardingTransactions) CreateSubscription(ctx context.Context, sub Subscription) (string, error) {
	var subID string
	err := t.DB.QueryRowContext(ctx,
		`INSERT INTO subscriptions (user_id, company_id, tier, status, billing_interval, payment_provider, provider_subscription_id)
		 VALUES ($1, $2, $3, $4, $5, $6, $7)
		 RETURNING id`,
		sub.UserID, sub.CompanyID, sub.Tier, sub.Status,
		sub.BillingInterval, sub.PaymentProvider, sub.ProviderSubID,
	).Scan(&subID)
	return subID, err
}

// ActivateSubscription updates subscription and user status
func (t *OnboardingTransactions) ActivateSubscription(ctx context.Context, subscriptionID, userID string) error {
	tx, err := t.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Update subscription status
	_, err = tx.ExecContext(ctx,
		`UPDATE subscriptions SET status = $1 WHERE id = $2`,
		"active", subscriptionID,
	)
	if err != nil {
		return err
	}

	// Update user status
	_, err = tx.ExecContext(ctx,
		`UPDATE users SET status = $1, updated_at = NOW() WHERE id = $2`,
		StatusSubscribed, userID,
	)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// UpdateOnboardingWorkflowState updates workflow state transactionally
func (t *OnboardingTransactions) UpdateOnboardingWorkflowState(ctx context.Context, state OnboardingWorkflowState) error {
	_, err := t.DB.ExecContext(ctx,
		`INSERT INTO onboarding_workflow_states 
		 (workflow_id, user_id, company_id, status, current_step, email_verified, subscription_id, error_message, created_at, updated_at)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW(), NOW())
		 ON CONFLICT (workflow_id) 
		 DO UPDATE SET status = $4, current_step = $5, email_verified = $6, 
		     subscription_id = $7, error_message = $8, updated_at = NOW()`,
		state.WorkflowID, state.UserID, state.CompanyID, state.Status,
		state.CurrentStep, state.EmailVerified, state.SubscriptionID, state.ErrorMessage,
	)
	return err
}

// GetWorkflowState retrieves workflow state
func (t *OnboardingTransactions) GetWorkflowState(ctx context.Context, workflowID string) (*OnboardingWorkflowState, error) {
	var state OnboardingWorkflowState
	err := t.DB.QueryRowContext(ctx,
		`SELECT workflow_id, user_id, company_id, status, email_verified, 
		        subscription_id, current_step, error_message, created_at, updated_at
		 FROM onboarding_workflow_states WHERE workflow_id = $1`,
		workflowID,
	).Scan(
		&state.WorkflowID, &state.UserID, &state.CompanyID, &state.Status,
		&state.EmailVerified, &state.SubscriptionID, &state.CurrentStep,
		&state.ErrorMessage, &state.CreatedAt, &state.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &state, nil
}
