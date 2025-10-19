package emailpkg

import (
	"context"
	"time"
)

// EmailProvider represents the type of email service provider
type EmailProvider string

const (
	ProviderAWSSES EmailProvider = "aws_ses"
	ProviderBrevo  EmailProvider = "brevo"
)

// EmailConfig holds the configuration for email services
type EmailConfig struct {
	Provider EmailProvider `json:"provider" yaml:"provider"`

	// AWS SES Configuration
	AWSRegion          string `json:"aws_region,omitempty" yaml:"aws_region,omitempty"`
	AWSAccessKeyID     string `json:"aws_access_key_id,omitempty" yaml:"aws_access_key_id,omitempty"`
	AWSSecretAccessKey string `json:"aws_secret_access_key,omitempty" yaml:"aws_secret_access_key,omitempty"`

	// Brevo Configuration
	BrevoAPIKey string `json:"brevo_api_key,omitempty" yaml:"brevo_api_key,omitempty"`
	BrevoAPIURL string `json:"brevo_api_url,omitempty" yaml:"brevo_api_url,omitempty"`

	// Common settings
	DefaultFromEmail string        `json:"default_from_email" yaml:"default_from_email"`
	DefaultFromName  string        `json:"default_from_name" yaml:"default_from_name"`
	Timeout          time.Duration `json:"timeout" yaml:"timeout"`
}

// EmailAddress represents an email address with optional name
type EmailAddress struct {
	Email string `json:"email"`
	Name  string `json:"name,omitempty"`
}

// Attachment represents an email attachment
type Attachment struct {
	Filename    string `json:"filename"`
	Content     []byte `json:"content"`
	ContentType string `json:"content_type"`
}

// EmailMessage represents a basic email message
type EmailMessage struct {
	From        EmailAddress      `json:"from"`
	To          []EmailAddress    `json:"to"`
	CC          []EmailAddress    `json:"cc,omitempty"`
	BCC         []EmailAddress    `json:"bcc,omitempty"`
	Subject     string            `json:"subject"`
	TextBody    string            `json:"text_body,omitempty"`
	HTMLBody    string            `json:"html_body,omitempty"`
	Attachments []Attachment      `json:"attachments,omitempty"`
	Headers     map[string]string `json:"headers,omitempty"`
}

// TemplateEmailMessage represents a template-based email message
type TemplateEmailMessage struct {
	From         EmailAddress           `json:"from"`
	To           []EmailAddress         `json:"to"`
	CC           []EmailAddress         `json:"cc,omitempty"`
	BCC          []EmailAddress         `json:"bcc,omitempty"`
	TemplateID   string                 `json:"template_id"`
	TemplateName string                 `json:"template_name,omitempty"`
	Variables    map[string]interface{} `json:"variables,omitempty"`
	Attachments  []Attachment           `json:"attachments,omitempty"`
	Headers      map[string]string      `json:"headers,omitempty"`
}

// EmailResponse represents the response from sending an email
type EmailResponse struct {
	MessageID string        `json:"message_id"`
	Status    string        `json:"status"`
	Provider  EmailProvider `json:"provider"`
	SentAt    time.Time     `json:"sent_at"`
	Error     string        `json:"error,omitempty"`
}

// EmailSender interface defines the contract for email sending services
type EmailSender interface {
	// SendEmail sends a basic email message
	SendEmail(ctx context.Context, message *EmailMessage) (*EmailResponse, error)

	// SendTemplateEmail sends a template-based email message
	SendTemplateEmail(ctx context.Context, message *TemplateEmailMessage) (*EmailResponse, error)

	// ValidateConfig validates the email service configuration
	ValidateConfig() error

	// GetProvider returns the email provider type
	GetProvider() EmailProvider
}

// EmailService is the main service that manages email sending
type EmailService struct {
	sender EmailSender
	config *EmailConfig
}

// EmailError represents an email-specific error
type EmailError struct {
	Provider EmailProvider `json:"provider"`
	Code     string        `json:"code"`
	Message  string        `json:"message"`
	Original error         `json:"-"`
}

func (e *EmailError) Error() string {
	if e.Original != nil {
		return e.Message + ": " + e.Original.Error()
	}
	return e.Message
}

func (e *EmailError) Unwrap() error {
	return e.Original
}
