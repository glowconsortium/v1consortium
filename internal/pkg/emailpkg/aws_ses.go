package emailpkg

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
)

// AWSSESSender implements EmailSender for AWS SES
type AWSSESSender struct {
	client *ses.Client
	config *EmailConfig
}

// NewAWSSESSender creates a new AWS SES email sender
func NewAWSSESSender(config *EmailConfig) (*AWSSESSender, error) {
	if config == nil {
		return nil, NewEmailError(ProviderAWSSES, ErrCodeInvalidConfig, "AWS SES configuration is required")
	}

	// Load AWS config
	ctx := context.Background()
	var awsConfig aws.Config
	var err error

	// Set up credentials if provided
	if config.AWSAccessKeyID != "" && config.AWSSecretAccessKey != "" {
		awsConfig, err = awsconfig.LoadDefaultConfig(ctx,
			awsconfig.WithRegion(config.AWSRegion),
			awsconfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
				config.AWSAccessKeyID,
				config.AWSSecretAccessKey,
				"",
			)),
		)
	} else {
		// Use default credential chain (IAM roles, environment variables, etc.)
		awsConfig, err = awsconfig.LoadDefaultConfig(ctx,
			awsconfig.WithRegion(config.AWSRegion),
		)
	}

	if err != nil {
		return nil, NewEmailErrorWithCause(ProviderAWSSES, ErrCodeAWSConfigLoadError,
			MsgAWSConfigLoad, err)
	}

	return &AWSSESSender{
		client: ses.NewFromConfig(awsConfig),
		config: config,
	}, nil
}

// SendEmail sends a basic email using AWS SES
func (s *AWSSESSender) SendEmail(ctx context.Context, message *EmailMessage) (*EmailResponse, error) {
	// Build destination
	destination := &types.Destination{}

	// Add To addresses
	if len(message.To) > 0 {
		toAddresses := make([]string, len(message.To))
		for i, addr := range message.To {
			toAddresses[i] = s.formatEmailAddress(addr)
		}
		destination.ToAddresses = toAddresses
	}

	// Add CC addresses
	if len(message.CC) > 0 {
		ccAddresses := make([]string, len(message.CC))
		for i, addr := range message.CC {
			ccAddresses[i] = s.formatEmailAddress(addr)
		}
		destination.CcAddresses = ccAddresses
	}

	// Add BCC addresses
	if len(message.BCC) > 0 {
		bccAddresses := make([]string, len(message.BCC))
		for i, addr := range message.BCC {
			bccAddresses[i] = s.formatEmailAddress(addr)
		}
		destination.BccAddresses = bccAddresses
	}

	// Build message
	sesMessage := &types.Message{
		Subject: &types.Content{
			Data:    aws.String(message.Subject),
			Charset: aws.String("UTF-8"),
		},
	}

	// Build body
	body := &types.Body{}
	if message.TextBody != "" {
		body.Text = &types.Content{
			Data:    aws.String(message.TextBody),
			Charset: aws.String("UTF-8"),
		}
	}
	if message.HTMLBody != "" {
		body.Html = &types.Content{
			Data:    aws.String(message.HTMLBody),
			Charset: aws.String("UTF-8"),
		}
	}
	sesMessage.Body = body

	// Send email
	input := &ses.SendEmailInput{
		Source:      aws.String(s.formatEmailAddress(message.From)),
		Destination: destination,
		Message:     sesMessage,
	}

	result, err := s.client.SendEmail(ctx, input)
	if err != nil {
		return nil, NewEmailErrorWithCause(ProviderAWSSES, ErrCodeAWSSendError, MsgAWSSendFailed, err)
	}

	return &EmailResponse{
		MessageID: aws.ToString(result.MessageId),
		Status:    "sent",
		Provider:  ProviderAWSSES,
		SentAt:    time.Now(),
	}, nil
}

// SendTemplateEmail sends a template-based email using AWS SES
func (s *AWSSESSender) SendTemplateEmail(ctx context.Context, message *TemplateEmailMessage) (*EmailResponse, error) {
	// Build destination
	destination := &types.Destination{}

	// Add To addresses
	if len(message.To) > 0 {
		toAddresses := make([]string, len(message.To))
		for i, addr := range message.To {
			toAddresses[i] = s.formatEmailAddress(addr)
		}
		destination.ToAddresses = toAddresses
	}

	// Add CC addresses
	if len(message.CC) > 0 {
		ccAddresses := make([]string, len(message.CC))
		for i, addr := range message.CC {
			ccAddresses[i] = s.formatEmailAddress(addr)
		}
		destination.CcAddresses = ccAddresses
	}

	// Add BCC addresses
	if len(message.BCC) > 0 {
		bccAddresses := make([]string, len(message.BCC))
		for i, addr := range message.BCC {
			bccAddresses[i] = s.formatEmailAddress(addr)
		}
		destination.BccAddresses = bccAddresses
	}

	// Build template data
	templateDataJSON := "{}"
	if len(message.Variables) > 0 {
		// Convert variables to JSON string
		var templateData strings.Builder
		templateData.WriteString("{")
		first := true
		for key, value := range message.Variables {
			if !first {
				templateData.WriteString(",")
			}
			templateData.WriteString(fmt.Sprintf(`"%s":"%v"`, key, value))
			first = false
		}
		templateData.WriteString("}")
		templateDataJSON = templateData.String()
	}

	// Use template ID or name
	templateName := message.TemplateID
	if templateName == "" {
		templateName = message.TemplateName
	}

	// Send templated email
	input := &ses.SendTemplatedEmailInput{
		Source:       aws.String(s.formatEmailAddress(message.From)),
		Destination:  destination,
		Template:     aws.String(templateName),
		TemplateData: aws.String(templateDataJSON),
	}

	result, err := s.client.SendTemplatedEmail(ctx, input)
	if err != nil {
		return nil, NewEmailErrorWithCause(ProviderAWSSES, ErrCodeAWSTemplateError, MsgAWSTemplateFailed, err)
	}

	return &EmailResponse{
		MessageID: aws.ToString(result.MessageId),
		Status:    "sent",
		Provider:  ProviderAWSSES,
		SentAt:    time.Now(),
	}, nil
}

// ValidateConfig validates the AWS SES configuration
func (s *AWSSESSender) ValidateConfig() error {
	if s.config.AWSRegion == "" {
		return NewEmailError(ProviderAWSSES, ErrCodeAWSMissingRegion, MsgAWSMissingRegion)
	}

	if s.config.DefaultFromEmail == "" {
		return NewEmailError(ProviderAWSSES, ErrCodeMissingFromEmail, MsgMissingFromEmail)
	}

	return nil
}

// GetProvider returns the provider type
func (s *AWSSESSender) GetProvider() EmailProvider {
	return ProviderAWSSES
}

// formatEmailAddress formats an email address for AWS SES
func (s *AWSSESSender) formatEmailAddress(addr EmailAddress) string {
	if addr.Name != "" {
		return fmt.Sprintf("%s <%s>", addr.Name, addr.Email)
	}
	return addr.Email
}
