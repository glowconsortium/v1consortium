package emailpkg

import (
	"context"
	"fmt"
	"time"
)

// NewEmailService creates a new email service with the specified configuration
func NewEmailService(config *EmailConfig) (*EmailService, error) {
	if config == nil {
		return nil, NewEmailError("", ErrCodeInvalidConfig, MsgInvalidConfig)
	}

	var sender EmailSender
	var err error

	switch config.Provider {
	case ProviderAWSSES:
		sender, err = NewAWSSESSender(config)
	case ProviderBrevo:
		sender, err = NewBrevoSender(config)
	default:
		return nil, NewEmailError("", ErrCodeUnsupportedProvider,
			fmt.Sprintf("%s: %s", MsgUnsupportedProvider, config.Provider))
	}

	if err != nil {
		return nil, err
	}

	if err := sender.ValidateConfig(); err != nil {
		return nil, NewEmailErrorWithCause(config.Provider, ErrCodeConfigValidationFailed,
			MsgConfigValidationFailed, err)
	}

	return &EmailService{
		sender: sender,
		config: config,
	}, nil
}

// SendEmail sends a basic email message
func (s *EmailService) SendEmail(ctx context.Context, message *EmailMessage) (*EmailResponse, error) {
	if message == nil {
		return nil, NewEmailError(s.config.Provider, ErrCodeInvalidMessage, MsgInvalidMessage)
	}

	// Set default from address if not provided
	if message.From.Email == "" {
		message.From.Email = s.config.DefaultFromEmail
		message.From.Name = s.config.DefaultFromName
	}

	// Validate message
	if err := s.validateEmailMessage(message); err != nil {
		return nil, err
	}

	return s.sender.SendEmail(ctx, message)
}

// SendTemplateEmail sends a template-based email message
func (s *EmailService) SendTemplateEmail(ctx context.Context, message *TemplateEmailMessage) (*EmailResponse, error) {
	if message == nil {
		return nil, NewEmailError(s.config.Provider, ErrCodeInvalidMessage, "template email message is required")
	}

	// Set default from address if not provided
	if message.From.Email == "" {
		message.From.Email = s.config.DefaultFromEmail
		message.From.Name = s.config.DefaultFromName
	}

	// Validate message
	if err := s.validateTemplateEmailMessage(message); err != nil {
		return nil, err
	}

	return s.sender.SendTemplateEmail(ctx, message)
}

// SendSimpleEmail sends a simple email with basic parameters
func (s *EmailService) SendSimpleEmail(ctx context.Context, to []string, subject, textBody, htmlBody string) (*EmailResponse, error) {
	toAddresses := make([]EmailAddress, len(to))
	for i, email := range to {
		toAddresses[i] = EmailAddress{Email: email}
	}

	message := &EmailMessage{
		From: EmailAddress{
			Email: s.config.DefaultFromEmail,
			Name:  s.config.DefaultFromName,
		},
		To:       toAddresses,
		Subject:  subject,
		TextBody: textBody,
		HTMLBody: htmlBody,
	}

	return s.SendEmail(ctx, message)
}

// GetProvider returns the configured email provider
func (s *EmailService) GetProvider() EmailProvider {
	return s.config.Provider
}

// GetConfig returns a copy of the email configuration (without sensitive data)
func (s *EmailService) GetConfig() *EmailConfig {
	configCopy := *s.config
	// Remove sensitive information
	configCopy.AWSAccessKeyID = ""
	configCopy.AWSSecretAccessKey = ""
	configCopy.BrevoAPIKey = ""
	return &configCopy
}

// validateEmailMessage validates a basic email message
func (s *EmailService) validateEmailMessage(message *EmailMessage) error {
	if message == nil {
		return NewEmailError(s.config.Provider, ErrCodeInvalidMessage, MsgInvalidMessage)
	}

	if message.From.Email == "" {
		return NewEmailError(s.config.Provider, ErrCodeInvalidFromEmail, MsgInvalidFromEmail)
	}

	if len(message.To) == 0 {
		return NewEmailError(s.config.Provider, ErrCodeInvalidToEmail, MsgInvalidToEmail)
	}

	if message.Subject == "" {
		return NewEmailError(s.config.Provider, ErrCodeInvalidSubject, MsgInvalidSubject)
	}

	if message.TextBody == "" && message.HTMLBody == "" {
		return NewEmailError(s.config.Provider, ErrCodeInvalidBody, MsgInvalidBody)
	}

	// Validate email addresses
	for _, addr := range append(append(message.To, message.CC...), message.BCC...) {
		if !isValidEmail(addr.Email) {
			return NewEmailError(s.config.Provider, ErrCodeInvalidEmailAddress,
				fmt.Sprintf("%s: %s", MsgInvalidEmailAddress, addr.Email))
		}
	}

	return nil
}

// validateTemplateEmailMessage validates a template email message
func (s *EmailService) validateTemplateEmailMessage(message *TemplateEmailMessage) error {
	if message.From.Email == "" {
		return NewEmailError(s.config.Provider, ErrCodeInvalidFromEmail, MsgInvalidFromEmail)
	}

	if len(message.To) == 0 {
		return NewEmailError(s.config.Provider, ErrCodeInvalidToEmail, MsgInvalidToEmail)
	}

	if message.TemplateID == "" && message.TemplateName == "" {
		return NewEmailError(s.config.Provider, ErrCodeInvalidTemplate, MsgInvalidTemplate)
	}

	// Validate email addresses
	for _, addr := range append(append(message.To, message.CC...), message.BCC...) {
		if !isValidEmail(addr.Email) {
			return NewEmailError(s.config.Provider, ErrCodeInvalidEmailAddress,
				fmt.Sprintf("%s: %s", MsgInvalidEmailAddress, addr.Email))
		}
	}

	return nil
}

// isValidEmail performs basic email validation
func isValidEmail(email string) bool {
	// Basic email validation - you might want to use a more robust solution
	if len(email) < 3 {
		return false
	}

	atIndex := -1
	for i, char := range email {
		if char == '@' {
			if atIndex >= 0 {
				return false // Multiple @ symbols
			}
			atIndex = i
		}
	}

	if atIndex <= 0 || atIndex >= len(email)-1 {
		return false // @ at beginning/end or not found
	}

	return true
}

// DefaultEmailConfig returns a default email configuration
func DefaultEmailConfig() *EmailConfig {
	return &EmailConfig{
		Provider:         ProviderAWSSES,
		AWSRegion:        "us-east-1",
		BrevoAPIURL:      "https://api.brevo.com/v3",
		DefaultFromEmail: "",
		DefaultFromName:  "",
		Timeout:          30 * time.Second,
	}
}
