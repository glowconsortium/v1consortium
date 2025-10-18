# Email Package

A flexible Go email package that supports multiple email service providers including AWS SES and Brevo (formerly SendinBlue). The package provides a unified interface for sending both regular emails and template-based emails.

## Features

- **Multi-provider support**: AWS SES and Brevo
- **Template emails**: Send emails using pre-defined templates
- **String emails**: Send emails with custom content
- **Unified interface**: Same API regardless of provider
- **Error handling**: Comprehensive error types and messages
- **Validation**: Email address and message validation
- **Configuration**: Flexible configuration options

## Installation

```bash
# For AWS SES support (using AWS SDK v2)
go get github.com/aws/aws-sdk-go-v2/aws
go get github.com/aws/aws-sdk-go-v2/config
go get github.com/aws/aws-sdk-go-v2/credentials
go get github.com/aws/aws-sdk-go-v2/service/ses

# Add to your go.mod (this package is internal)
```

## Quick Start

### AWS SES Configuration

```go
package main

import (
    "context"
    "log"
    "time"
    
    "path/to/your/internal/pkg/emailpkg"
)

func main() {
    // Configure AWS SES
    config := &emailpkg.EmailConfig{
        Provider:           emailpkg.ProviderAWSSES,
        AWSRegion:          "us-east-1",
        AWSAccessKeyID:     "your-access-key-id",     // Optional if using IAM roles
        AWSSecretAccessKey: "your-secret-access-key", // Optional if using IAM roles
        DefaultFromEmail:   "noreply@yourdomain.com",
        DefaultFromName:    "Your App Name",
        Timeout:            30 * time.Second,
    }

    // Create email service
    emailService, err := emailpkg.NewEmailService(config)
    if err != nil {
        log.Fatalf("Failed to create email service: %v", err)
    }

    // Send a simple email
    ctx := context.Background()
    response, err := emailService.SendSimpleEmail(
        ctx,
        []string{"user@example.com"},
        "Welcome!",
        "Welcome to our service!",
        "<h1>Welcome to our service!</h1>",
    )
    
    if err != nil {
        log.Printf("Error: %v", err)
        return
    }
    
    log.Printf("Email sent! Message ID: %s", response.MessageID)
}
```

### Brevo Configuration

```go
package main

import (
    "context"
    "log"
    "time"
    
    "path/to/your/internal/pkg/emailpkg"
)

func main() {
    // Configure Brevo
    config := &emailpkg.EmailConfig{
        Provider:         emailpkg.ProviderBrevo,
        BrevoAPIKey:      "your-brevo-api-key",
        DefaultFromEmail: "noreply@yourdomain.com", 
        DefaultFromName:  "Your App Name",
        Timeout:          30 * time.Second,
    }

    // Create email service
    emailService, err := emailpkg.NewEmailService(config)
    if err != nil {
        log.Fatalf("Failed to create email service: %v", err)
    }

    // Send email
    ctx := context.Background()
    response, err := emailService.SendSimpleEmail(
        ctx,
        []string{"user@example.com"},
        "Welcome!",
        "Welcome to our service!",
        "<h1>Welcome to our service!</h1>",
    )
    
    if err != nil {
        log.Printf("Error: %v", err)
        return
    }
    
    log.Printf("Email sent! Message ID: %s", response.MessageID)
}
```

## Advanced Usage

### Sending Complex Emails

```go
message := &emailpkg.EmailMessage{
    From: emailpkg.EmailAddress{
        Email: "support@yourdomain.com",
        Name:  "Support Team",
    },
    To: []emailpkg.EmailAddress{
        {Email: "user@example.com", Name: "John Doe"},
        {Email: "user2@example.com", Name: "Jane Smith"},
    },
    CC: []emailpkg.EmailAddress{
        {Email: "manager@yourdomain.com", Name: "Manager"},
    },
    BCC: []emailpkg.EmailAddress{
        {Email: "audit@yourdomain.com"},
    },
    Subject:  "Important Account Update",
    TextBody: "Your account has been updated. Please review the changes.",
    HTMLBody: "<h2>Account Update</h2><p>Your account has been updated. Please review the changes.</p>",
    Headers: map[string]string{
        "X-Priority": "1",
        "X-Category": "account-update",
    },
}

response, err := emailService.SendEmail(ctx, message)
if err != nil {
    log.Printf("Failed to send email: %v", err)
    return
}

log.Printf("Email sent successfully. Message ID: %s", response.MessageID)
```

### Sending Template Emails

#### AWS SES Template Email

```go
templateMessage := &emailpkg.TemplateEmailMessage{
    From: emailpkg.EmailAddress{
        Email: "noreply@yourdomain.com",
        Name:  "Your App",
    },
    To: []emailpkg.EmailAddress{
        {Email: "user@example.com", Name: "John Doe"},
    },
    TemplateID: "welcome-template", // AWS SES template name
    Variables: map[string]interface{}{
        "username":     "John",
        "company_name": "Your Company",
        "login_url":    "https://yourdomain.com/login",
    },
}

response, err := emailService.SendTemplateEmail(ctx, templateMessage)
if err != nil {
    log.Printf("Failed to send template email: %v", err)
    return
}

log.Printf("Template email sent successfully. Message ID: %s", response.MessageID)
```

#### Brevo Template Email

```go
templateMessage := &emailpkg.TemplateEmailMessage{
    From: emailpkg.EmailAddress{
        Email: "noreply@yourdomain.com",
        Name:  "Your App",
    },
    To: []emailpkg.EmailAddress{
        {Email: "user@example.com", Name: "John Doe"},
    },
    TemplateID: "1", // Brevo template ID (must be numeric)
    Variables: map[string]interface{}{
        "FNAME":        "John",          // Brevo template variables
        "COMPANY_NAME": "Your Company",
        "LOGIN_URL":    "https://yourdomain.com/login",
    },
}

response, err := emailService.SendTemplateEmail(ctx, templateMessage)
```

## Configuration Options

### EmailConfig Fields

| Field | Type | Description | AWS SES | Brevo |
|-------|------|-------------|---------|-------|
| `Provider` | `EmailProvider` | Email service provider | ✓ | ✓ |
| `AWSRegion` | `string` | AWS region | ✓ | ✗ |
| `AWSAccessKeyID` | `string` | AWS access key (optional if using IAM) | ✓ | ✗ |
| `AWSSecretAccessKey` | `string` | AWS secret key (optional if using IAM) | ✓ | ✗ |
| `BrevoAPIKey` | `string` | Brevo API key | ✗ | ✓ |
| `BrevoAPIURL` | `string` | Brevo API URL (default: https://api.brevo.com/v3) | ✗ | ✓ |
| `DefaultFromEmail` | `string` | Default sender email address | ✓ | ✓ |
| `DefaultFromName` | `string` | Default sender name | ✓ | ✓ |
| `Timeout` | `time.Duration` | Request timeout | ✓ | ✓ |

### Environment Variables

You can use environment variables to configure the email service:

```bash
# General
EMAIL_PROVIDER=aws_ses # or brevo
EMAIL_FROM_EMAIL=noreply@yourdomain.com
EMAIL_FROM_NAME="Your App Name"
EMAIL_TIMEOUT=30s

# AWS SES
AWS_REGION=us-east-1
AWS_ACCESS_KEY_ID=your-access-key
AWS_SECRET_ACCESS_KEY=your-secret-key

# Brevo
BREVO_API_KEY=your-brevo-api-key
BREVO_API_URL=https://api.brevo.com/v3
```

## Error Handling

The package provides structured error handling:

```go
response, err := emailService.SendEmail(ctx, message)
if err != nil {
    if emailErr, ok := err.(*emailpkg.EmailError); ok {
        log.Printf("Email error - Provider: %s, Code: %s, Message: %s", 
            emailErr.Provider, emailErr.Code, emailErr.Message)
        
        if emailErr.Original != nil {
            log.Printf("Original error: %v", emailErr.Original)
        }
    } else {
        log.Printf("Unknown error: %v", err)
    }
    return
}
```

### Common Error Codes

- `INVALID_CONFIG`: Configuration validation failed
- `INVALID_MESSAGE`: Message validation failed
- `INVALID_EMAIL_ADDRESS`: Invalid email address format
- `SEND_ERROR`: Failed to send email
- `SEND_TEMPLATE_ERROR`: Failed to send template email
- `HTTP_ERROR`: HTTP request failed (Brevo)
- `SESSION_ERROR`: AWS session creation failed (AWS SES)

## Provider-Specific Notes

### AWS SES

- Uses AWS SDK v2 (the recommended version)
- Requires AWS credentials (access key/secret key or IAM role)
- Template names are strings
- Supports both verified domains and individual email addresses
- Requires email addresses to be verified in sandbox mode
- Template variables are JSON strings

### Brevo

- Requires API key from Brevo dashboard
- Template IDs must be numeric
- Template variables use Brevo's naming convention (often UPPERCASE)
- Different API endpoints for different operations
- Built-in rate limiting and retry logic

## Dependencies

### AWS SES
```go
github.com/aws/aws-sdk-go-v2/aws
github.com/aws/aws-sdk-go-v2/config
github.com/aws/aws-sdk-go-v2/credentials
github.com/aws/aws-sdk-go-v2/service/ses
```

### Brevo
- Only standard library dependencies (net/http, encoding/json)

## Testing

The package includes comprehensive examples in `examples.go`. To test with your configuration:

1. Update the configuration in the examples
2. Run the examples with your actual credentials
3. Check your email provider's dashboard for delivery status

## Security Considerations

- Never commit API keys or credentials to version control
- Use environment variables or secure configuration management
- Implement proper access controls for email sending
- Validate and sanitize email content to prevent injection attacks
- Use HTTPS for all API communications
- Rotate API keys regularly

## Contributing

When adding new email providers:

1. Implement the `EmailSender` interface
2. Add the provider constant to `EmailProvider` type
3. Update the `NewEmailService` function
4. Add configuration fields to `EmailConfig`
5. Add examples and documentation
6. Add appropriate error handling

## License

This package is part of the internal codebase and follows the project's licensing terms.