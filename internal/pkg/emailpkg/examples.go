package emailpkg

import (
	"context"
	"log"
	"time"
)

// Example demonstrates how to use the email package

// ExampleAWSSES shows how to set up and use AWS SES
func ExampleAWSSES() {
	// Create AWS SES configuration
	config := &EmailConfig{
		Provider:           ProviderAWSSES,
		AWSRegion:          "us-east-1",
		AWSAccessKeyID:     "your-access-key-id",
		AWSSecretAccessKey: "your-secret-access-key",
		DefaultFromEmail:   "noreply@yourdomain.com",
		DefaultFromName:    "Your App Name",
		Timeout:            30 * time.Second,
	}

	// Create email service
	emailService, err := NewEmailService(config)
	if err != nil {
		log.Fatalf("Failed to create email service: %v", err)
	}

	ctx := context.Background()

	// Send a simple email
	response, err := emailService.SendSimpleEmail(
		ctx,
		[]string{"user@example.com"},
		"Welcome to our service",
		"Thank you for signing up! This is the text version.",
		"<h1>Thank you for signing up!</h1><p>This is the HTML version.</p>",
	)
	if err != nil {
		log.Printf("Failed to send email: %v", err)
		return
	}

	log.Printf("Email sent successfully. Message ID: %s", response.MessageID)

	// Send a more complex email
	message := &EmailMessage{
		From: EmailAddress{
			Email: "support@yourdomain.com",
			Name:  "Support Team",
		},
		To: []EmailAddress{
			{Email: "user@example.com", Name: "John Doe"},
		},
		CC: []EmailAddress{
			{Email: "manager@yourdomain.com", Name: "Manager"},
		},
		Subject:  "Account Verification Required",
		TextBody: "Please verify your account by clicking the link.",
		HTMLBody: "<p>Please verify your account by <a href='#'>clicking here</a>.</p>",
		Headers: map[string]string{
			"X-Priority": "1",
		},
	}

	response, err = emailService.SendEmail(ctx, message)
	if err != nil {
		log.Printf("Failed to send complex email: %v", err)
		return
	}

	log.Printf("Complex email sent successfully. Message ID: %s", response.MessageID)

	// Send a template email (requires AWS SES template to be created)
	templateMessage := &TemplateEmailMessage{
		From: EmailAddress{
			Email: "noreply@yourdomain.com",
			Name:  "Your App",
		},
		To: []EmailAddress{
			{Email: "user@example.com", Name: "John Doe"},
		},
		TemplateID: "welcome-template",
		Variables: map[string]interface{}{
			"username":     "John",
			"company_name": "Your Company",
			"login_url":    "https://yourdomain.com/login",
		},
	}

	response, err = emailService.SendTemplateEmail(ctx, templateMessage)
	if err != nil {
		log.Printf("Failed to send template email: %v", err)
		return
	}

	log.Printf("Template email sent successfully. Message ID: %s", response.MessageID)
}

// ExampleBrevo shows how to set up and use Brevo
func ExampleBrevo() {
	// Create Brevo configuration
	config := &EmailConfig{
		Provider:         ProviderBrevo,
		BrevoAPIKey:      "your-brevo-api-key",
		BrevoAPIURL:      "https://api.brevo.com/v3", // Optional, this is the default
		DefaultFromEmail: "noreply@yourdomain.com",
		DefaultFromName:  "Your App Name",
		Timeout:          30 * time.Second,
	}

	// Create email service
	emailService, err := NewEmailService(config)
	if err != nil {
		log.Fatalf("Failed to create email service: %v", err)
	}

	ctx := context.Background()

	// Send a simple email
	response, err := emailService.SendSimpleEmail(
		ctx,
		[]string{"user@example.com"},
		"Welcome to our service",
		"Thank you for signing up! This is the text version.",
		"<h1>Thank you for signing up!</h1><p>This is the HTML version.</p>",
	)
	if err != nil {
		log.Printf("Failed to send email: %v", err)
		return
	}

	log.Printf("Email sent successfully. Message ID: %s", response.MessageID)

	// Send a template email (requires Brevo template to be created)
	// Note: Brevo requires numeric template IDs
	templateMessage := &TemplateEmailMessage{
		From: EmailAddress{
			Email: "noreply@yourdomain.com",
			Name:  "Your App",
		},
		To: []EmailAddress{
			{Email: "user@example.com", Name: "John Doe"},
		},
		TemplateID: "1", // Brevo template ID (numeric)
		Variables: map[string]interface{}{
			"FNAME":        "John",
			"COMPANY_NAME": "Your Company",
			"LOGIN_URL":    "https://yourdomain.com/login",
		},
	}

	response, err = emailService.SendTemplateEmail(ctx, templateMessage)
	if err != nil {
		log.Printf("Failed to send template email: %v", err)
		return
	}

	log.Printf("Template email sent successfully. Message ID: %s", response.MessageID)
}

// ExampleWithEnvironmentConfig shows how to load configuration from environment
func ExampleWithEnvironmentConfig() {
	// You would typically load these from environment variables or config files
	config := DefaultEmailConfig()

	// Override with your actual values
	config.Provider = ProviderAWSSES // or ProviderBrevo
	config.DefaultFromEmail = "noreply@yourdomain.com"
	config.DefaultFromName = "Your App Name"

	// For AWS SES
	if config.Provider == ProviderAWSSES {
		config.AWSRegion = "us-east-1"
		config.AWSAccessKeyID = "your-access-key"
		config.AWSSecretAccessKey = "your-secret-key"
	}

	// For Brevo
	if config.Provider == ProviderBrevo {
		config.BrevoAPIKey = "your-brevo-api-key"
	}

	emailService, err := NewEmailService(config)
	if err != nil {
		log.Fatalf("Failed to create email service: %v", err)
	}

	// Use the service...
	ctx := context.Background()
	response, err := emailService.SendSimpleEmail(
		ctx,
		[]string{"test@example.com"},
		"Test Email",
		"This is a test email.",
		"<p>This is a <strong>test</strong> email.</p>",
	)

	if err != nil {
		log.Printf("Failed to send email: %v", err)
		return
	}

	log.Printf("Email sent successfully. Provider: %s, Message ID: %s",
		response.Provider, response.MessageID)
}
