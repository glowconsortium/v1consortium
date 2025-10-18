package emailpkg

import (
	"context"
	"testing"
	"time"
)

// TestEmailConfigValidation tests the configuration validation
func TestEmailConfigValidation(t *testing.T) {
	tests := []struct {
		name        string
		config      *EmailConfig
		shouldError bool
	}{
		{
			name:        "nil config",
			config:      nil,
			shouldError: true,
		},
		{
			name: "invalid provider",
			config: &EmailConfig{
				Provider:         "invalid",
				DefaultFromEmail: "test@example.com",
			},
			shouldError: true,
		},
		{
			name: "aws ses missing region",
			config: &EmailConfig{
				Provider:         ProviderAWSSES,
				DefaultFromEmail: "test@example.com",
			},
			shouldError: true,
		},
		{
			name: "brevo missing api key",
			config: &EmailConfig{
				Provider:         ProviderBrevo,
				DefaultFromEmail: "test@example.com",
			},
			shouldError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewEmailService(tt.config)
			if tt.shouldError && err == nil {
				t.Errorf("expected error but got none")
			}
			if !tt.shouldError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

// TestEmailValidation tests email address validation
func TestEmailValidation(t *testing.T) {
	tests := []struct {
		email string
		valid bool
	}{
		{"test@example.com", true},
		{"user.name@example.com", true},
		{"user+tag@example.com", true},
		{"", false},
		{"invalid", false},
		{"@example.com", false},
		{"test@", false},
		{"test@@example.com", false},
	}

	for _, tt := range tests {
		t.Run(tt.email, func(t *testing.T) {
			result := isValidEmail(tt.email)
			if result != tt.valid {
				t.Errorf("isValidEmail(%q) = %v, want %v", tt.email, result, tt.valid)
			}
		})
	}
}

// TestEmailMessageValidation tests email message validation
func TestEmailMessageValidation(t *testing.T) {
	// Create a mock email service for testing
	config := &EmailConfig{
		Provider:         ProviderAWSSES,
		AWSRegion:        "us-east-1",
		DefaultFromEmail: "test@example.com",
		DefaultFromName:  "Test",
		Timeout:          30 * time.Second,
	}

	service := &EmailService{
		config: config,
	}

	tests := []struct {
		name        string
		message     *EmailMessage
		shouldError bool
	}{
		{
			name:        "nil message",
			message:     nil,
			shouldError: true,
		},
		{
			name: "valid message",
			message: &EmailMessage{
				From:     EmailAddress{Email: "test@example.com"},
				To:       []EmailAddress{{Email: "user@example.com"}},
				Subject:  "Test",
				TextBody: "Test body",
			},
			shouldError: false,
		},
		{
			name: "missing to addresses",
			message: &EmailMessage{
				From:     EmailAddress{Email: "test@example.com"},
				To:       []EmailAddress{},
				Subject:  "Test",
				TextBody: "Test body",
			},
			shouldError: true,
		},
		{
			name: "missing subject",
			message: &EmailMessage{
				From:     EmailAddress{Email: "test@example.com"},
				To:       []EmailAddress{{Email: "user@example.com"}},
				TextBody: "Test body",
			},
			shouldError: true,
		},
		{
			name: "missing body",
			message: &EmailMessage{
				From:    EmailAddress{Email: "test@example.com"},
				To:      []EmailAddress{{Email: "user@example.com"}},
				Subject: "Test",
			},
			shouldError: true,
		},
		{
			name: "invalid email address",
			message: &EmailMessage{
				From:     EmailAddress{Email: "test@example.com"},
				To:       []EmailAddress{{Email: "invalid-email"}},
				Subject:  "Test",
				TextBody: "Test body",
			},
			shouldError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := service.validateEmailMessage(tt.message)
			if tt.shouldError && err == nil {
				t.Errorf("expected error but got none")
			}
			if !tt.shouldError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

// TestTemplateEmailMessageValidation tests template email message validation
func TestTemplateEmailMessageValidation(t *testing.T) {
	// Create a mock email service for testing
	config := &EmailConfig{
		Provider:         ProviderAWSSES,
		AWSRegion:        "us-east-1",
		DefaultFromEmail: "test@example.com",
		DefaultFromName:  "Test",
		Timeout:          30 * time.Second,
	}

	service := &EmailService{
		config: config,
	}

	tests := []struct {
		name        string
		message     *TemplateEmailMessage
		shouldError bool
	}{
		{
			name: "valid template message",
			message: &TemplateEmailMessage{
				From:       EmailAddress{Email: "test@example.com"},
				To:         []EmailAddress{{Email: "user@example.com"}},
				TemplateID: "welcome-template",
			},
			shouldError: false,
		},
		{
			name: "missing template",
			message: &TemplateEmailMessage{
				From: EmailAddress{Email: "test@example.com"},
				To:   []EmailAddress{{Email: "user@example.com"}},
			},
			shouldError: true,
		},
		{
			name: "valid with template name",
			message: &TemplateEmailMessage{
				From:         EmailAddress{Email: "test@example.com"},
				To:           []EmailAddress{{Email: "user@example.com"}},
				TemplateName: "welcome-template",
			},
			shouldError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := service.validateTemplateEmailMessage(tt.message)
			if tt.shouldError && err == nil {
				t.Errorf("expected error but got none")
			}
			if !tt.shouldError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

// TestDefaultEmailConfig tests the default configuration
func TestDefaultEmailConfig(t *testing.T) {
	config := DefaultEmailConfig()

	if config.Provider != ProviderAWSSES {
		t.Errorf("expected default provider to be %s, got %s", ProviderAWSSES, config.Provider)
	}

	if config.AWSRegion != "us-east-1" {
		t.Errorf("expected default AWS region to be us-east-1, got %s", config.AWSRegion)
	}

	if config.BrevoAPIURL != "https://api.brevo.com/v3" {
		t.Errorf("expected default Brevo API URL to be https://api.brevo.com/v3, got %s", config.BrevoAPIURL)
	}

	if config.Timeout != 30*time.Second {
		t.Errorf("expected default timeout to be 30s, got %v", config.Timeout)
	}
}

// TestEmailError tests the EmailError type
func TestEmailError(t *testing.T) {
	err := &EmailError{
		Provider: ProviderAWSSES,
		Code:     "TEST_ERROR",
		Message:  "Test error message",
	}

	expected := "Test error message"
	if err.Error() != expected {
		t.Errorf("expected error message %q, got %q", expected, err.Error())
	}

	// Test with original error
	originalErr := context.DeadlineExceeded
	err.Original = originalErr

	expected = "Test error message: " + originalErr.Error()
	if err.Error() != expected {
		t.Errorf("expected error message %q, got %q", expected, err.Error())
	}

	// Test Unwrap
	if err.Unwrap() != originalErr {
		t.Errorf("expected Unwrap to return original error")
	}
}

// BenchmarkEmailValidation benchmarks the email validation function
func BenchmarkEmailValidation(t *testing.B) {
	emails := []string{
		"test@example.com",
		"user.name@example.com",
		"user+tag@example.com",
		"invalid",
		"@example.com",
		"test@",
	}

	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		for _, email := range emails {
			isValidEmail(email)
		}
	}
}
