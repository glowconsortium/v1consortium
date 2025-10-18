package emailpkg

import (
	"errors"
	"fmt"
)

// Error codes for different types of email errors
const (
	// Configuration errors
	ErrCodeInvalidConfig          = "INVALID_CONFIG"
	ErrCodeMissingConfig          = "MISSING_CONFIG"
	ErrCodeUnsupportedProvider    = "UNSUPPORTED_PROVIDER"
	ErrCodeConfigValidationFailed = "CONFIG_VALIDATION_FAILED"

	// AWS SES specific errors
	ErrCodeAWSConfigLoadError = "AWS_CONFIG_LOAD_ERROR"
	ErrCodeAWSMissingRegion   = "AWS_MISSING_REGION"
	ErrCodeAWSSessionError    = "AWS_SESSION_ERROR"
	ErrCodeAWSSendError       = "AWS_SEND_ERROR"
	ErrCodeAWSTemplateError   = "AWS_TEMPLATE_ERROR"

	// Brevo specific errors
	ErrCodeBrevoMissingAPIKey     = "BREVO_MISSING_API_KEY"
	ErrCodeBrevoInvalidTemplateID = "BREVO_INVALID_TEMPLATE_ID"
	ErrCodeBrevoMissingTemplateID = "BREVO_MISSING_TEMPLATE_ID"
	ErrCodeBrevoRequestError      = "BREVO_REQUEST_ERROR"
	ErrCodeBrevoResponseError     = "BREVO_RESPONSE_ERROR"
	ErrCodeBrevoHTTPError         = "BREVO_HTTP_ERROR"

	// Message validation errors
	ErrCodeInvalidMessage      = "INVALID_MESSAGE"
	ErrCodeInvalidFromEmail    = "INVALID_FROM_EMAIL"
	ErrCodeInvalidToEmail      = "INVALID_TO_EMAIL"
	ErrCodeInvalidEmailAddress = "INVALID_EMAIL_ADDRESS"
	ErrCodeInvalidSubject      = "INVALID_SUBJECT"
	ErrCodeInvalidBody         = "INVALID_BODY"
	ErrCodeInvalidTemplate     = "INVALID_TEMPLATE"
	ErrCodeMissingFromEmail    = "MISSING_FROM_EMAIL"
	ErrCodeMissingToEmail      = "MISSING_TO_EMAIL"
	ErrCodeMissingSubject      = "MISSING_SUBJECT"
	ErrCodeMissingBody         = "MISSING_BODY"
	ErrCodeMissingTemplate     = "MISSING_TEMPLATE"

	// Network and communication errors
	ErrCodeNetworkError        = "NETWORK_ERROR"
	ErrCodeTimeoutError        = "TIMEOUT_ERROR"
	ErrCodeHTTPError           = "HTTP_ERROR"
	ErrCodeRequestMarshalError = "REQUEST_MARSHAL_ERROR"
	ErrCodeResponseParseError  = "RESPONSE_PARSE_ERROR"
	ErrCodeRequestCreateError  = "REQUEST_CREATE_ERROR"
	ErrCodeRequestSendError    = "REQUEST_SEND_ERROR"
	ErrCodeResponseReadError   = "RESPONSE_READ_ERROR"

	// Rate limiting and quota errors
	ErrCodeRateLimitExceeded = "RATE_LIMIT_EXCEEDED"
	ErrCodeQuotaExceeded     = "QUOTA_EXCEEDED"
	ErrCodeThrottled         = "THROTTLED"

	// Authentication and authorization errors
	ErrCodeUnauthorized       = "UNAUTHORIZED"
	ErrCodeInvalidAPIKey      = "INVALID_API_KEY"
	ErrCodeInvalidCredentials = "INVALID_CREDENTIALS"
	ErrCodePermissionDenied   = "PERMISSION_DENIED"

	// Service specific errors
	ErrCodeServiceUnavailable = "SERVICE_UNAVAILABLE"
	ErrCodeInternalError      = "INTERNAL_ERROR"
	ErrCodeUnknownError       = "UNKNOWN_ERROR"
)

// Error messages for common scenarios
const (
	MsgInvalidConfig          = "email configuration is required"
	MsgUnsupportedProvider    = "unsupported email provider"
	MsgConfigValidationFailed = "email service configuration validation failed"

	MsgAWSMissingRegion  = "AWS region is required"
	MsgAWSConfigLoad     = "failed to load AWS configuration"
	MsgAWSSendFailed     = "failed to send email via AWS SES"
	MsgAWSTemplateFailed = "failed to send template email via AWS SES"

	MsgBrevoMissingAPIKey     = "Brevo API key is required"
	MsgBrevoInvalidTemplateID = "template ID must be a number for Brevo"
	MsgBrevoMissingTemplateID = "template ID is required for Brevo (template name not supported)"
	MsgBrevoSendFailed        = "failed to send email via Brevo"

	MsgInvalidMessage      = "email message is required"
	MsgInvalidFromEmail    = "from email address is required"
	MsgInvalidToEmail      = "at least one recipient email address is required"
	MsgInvalidEmailAddress = "invalid email address"
	MsgInvalidSubject      = "email subject is required"
	MsgInvalidBody         = "either text body or HTML body is required"
	MsgInvalidTemplate     = "either template ID or template name is required"
	MsgMissingFromEmail    = "default from email is required"

	MsgNetworkError        = "network error occurred"
	MsgTimeoutError        = "request timeout"
	MsgRequestMarshalError = "failed to marshal request payload"
	MsgResponseParseError  = "failed to parse response"
	MsgRequestCreateError  = "failed to create HTTP request"
	MsgRequestSendError    = "failed to send HTTP request"
	MsgResponseReadError   = "failed to read response body"

	MsgRateLimitExceeded = "rate limit exceeded"
	MsgQuotaExceeded     = "quota exceeded"
	MsgThrottled         = "request throttled"

	MsgUnauthorized       = "unauthorized access"
	MsgInvalidAPIKey      = "invalid API key"
	MsgInvalidCredentials = "invalid credentials"
	MsgPermissionDenied   = "permission denied"

	MsgServiceUnavailable = "email service unavailable"
	MsgInternalError      = "internal server error"
	MsgUnknownError       = "unknown error occurred"
)

// Predefined error variables for common scenarios
var (
	// Configuration errors
	ErrInvalidConfig       = errors.New("invalid email configuration")
	ErrUnsupportedProvider = errors.New("unsupported email provider")
	ErrMissingConfig       = errors.New("missing email configuration")

	// Message validation errors
	ErrInvalidMessage      = errors.New("invalid email message")
	ErrMissingFromEmail    = errors.New("missing from email address")
	ErrMissingToEmail      = errors.New("missing recipient email addresses")
	ErrInvalidEmailAddress = errors.New("invalid email address format")
	ErrMissingSubject      = errors.New("missing email subject")
	ErrMissingBody         = errors.New("missing email body")
	ErrInvalidTemplate     = errors.New("invalid template specification")

	// Network errors
	ErrNetworkFailure = errors.New("network communication failed")
	ErrTimeout        = errors.New("request timeout")
	ErrServiceDown    = errors.New("email service unavailable")

	// Authentication errors
	ErrUnauthorized       = errors.New("unauthorized access")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrPermissionDenied   = errors.New("permission denied")

	// Rate limiting errors
	ErrRateLimitExceeded = errors.New("rate limit exceeded")
	ErrQuotaExceeded     = errors.New("quota exceeded")
)

// NewEmailError creates a new EmailError with the specified provider, code, and message
func NewEmailError(provider EmailProvider, code, message string) *EmailError {
	return &EmailError{
		Provider: provider,
		Code:     code,
		Message:  message,
	}
}

// NewEmailErrorWithCause creates a new EmailError with an underlying cause
func NewEmailErrorWithCause(provider EmailProvider, code, message string, cause error) *EmailError {
	return &EmailError{
		Provider: provider,
		Code:     code,
		Message:  message,
		Original: cause,
	}
}

// NewConfigError creates a configuration-related error
func NewConfigError(provider EmailProvider, message string) *EmailError {
	return &EmailError{
		Provider: provider,
		Code:     ErrCodeInvalidConfig,
		Message:  message,
	}
}

// NewValidationError creates a validation-related error
func NewValidationError(provider EmailProvider, field, value string) *EmailError {
	return &EmailError{
		Provider: provider,
		Code:     ErrCodeInvalidMessage,
		Message:  fmt.Sprintf("validation failed for field '%s': %s", field, value),
	}
}

// NewNetworkError creates a network-related error
func NewNetworkError(provider EmailProvider, cause error) *EmailError {
	return &EmailError{
		Provider: provider,
		Code:     ErrCodeNetworkError,
		Message:  MsgNetworkError,
		Original: cause,
	}
}

// NewAuthError creates an authentication-related error
func NewAuthError(provider EmailProvider, message string) *EmailError {
	return &EmailError{
		Provider: provider,
		Code:     ErrCodeUnauthorized,
		Message:  message,
	}
}

// NewRateLimitError creates a rate limiting error
func NewRateLimitError(provider EmailProvider) *EmailError {
	return &EmailError{
		Provider: provider,
		Code:     ErrCodeRateLimitExceeded,
		Message:  MsgRateLimitExceeded,
	}
}

// NewServiceError creates a service-related error
func NewServiceError(provider EmailProvider, message string, cause error) *EmailError {
	return &EmailError{
		Provider: provider,
		Code:     ErrCodeServiceUnavailable,
		Message:  message,
		Original: cause,
	}
}

// IsConfigError checks if the error is a configuration error
func IsConfigError(err error) bool {
	if emailErr, ok := err.(*EmailError); ok {
		return emailErr.Code == ErrCodeInvalidConfig ||
			emailErr.Code == ErrCodeMissingConfig ||
			emailErr.Code == ErrCodeConfigValidationFailed ||
			emailErr.Code == ErrCodeUnsupportedProvider
	}
	return false
}

// IsValidationError checks if the error is a validation error
func IsValidationError(err error) bool {
	if emailErr, ok := err.(*EmailError); ok {
		return emailErr.Code == ErrCodeInvalidMessage ||
			emailErr.Code == ErrCodeInvalidFromEmail ||
			emailErr.Code == ErrCodeInvalidToEmail ||
			emailErr.Code == ErrCodeInvalidEmailAddress ||
			emailErr.Code == ErrCodeInvalidSubject ||
			emailErr.Code == ErrCodeInvalidBody ||
			emailErr.Code == ErrCodeInvalidTemplate ||
			emailErr.Code == ErrCodeMissingFromEmail ||
			emailErr.Code == ErrCodeMissingToEmail ||
			emailErr.Code == ErrCodeMissingSubject ||
			emailErr.Code == ErrCodeMissingBody ||
			emailErr.Code == ErrCodeMissingTemplate
	}
	return false
}

// IsNetworkError checks if the error is a network-related error
func IsNetworkError(err error) bool {
	if emailErr, ok := err.(*EmailError); ok {
		return emailErr.Code == ErrCodeNetworkError ||
			emailErr.Code == ErrCodeTimeoutError ||
			emailErr.Code == ErrCodeHTTPError ||
			emailErr.Code == ErrCodeRequestSendError ||
			emailErr.Code == ErrCodeResponseReadError
	}
	return false
}

// IsAuthError checks if the error is an authentication/authorization error
func IsAuthError(err error) bool {
	if emailErr, ok := err.(*EmailError); ok {
		return emailErr.Code == ErrCodeUnauthorized ||
			emailErr.Code == ErrCodeInvalidAPIKey ||
			emailErr.Code == ErrCodeInvalidCredentials ||
			emailErr.Code == ErrCodePermissionDenied
	}
	return false
}

// IsRateLimitError checks if the error is a rate limiting error
func IsRateLimitError(err error) bool {
	if emailErr, ok := err.(*EmailError); ok {
		return emailErr.Code == ErrCodeRateLimitExceeded ||
			emailErr.Code == ErrCodeQuotaExceeded ||
			emailErr.Code == ErrCodeThrottled
	}
	return false
}

// IsRetryableError checks if the error is retryable
func IsRetryableError(err error) bool {
	return IsNetworkError(err) || IsRateLimitError(err) || IsServiceError(err)
}

// IsServiceError checks if the error is a service-related error
func IsServiceError(err error) bool {
	if emailErr, ok := err.(*EmailError); ok {
		return emailErr.Code == ErrCodeServiceUnavailable ||
			emailErr.Code == ErrCodeInternalError ||
			emailErr.Code == ErrCodeUnknownError
	}
	return false
}

// IsAWSError checks if the error is AWS SES specific
func IsAWSError(err error) bool {
	if emailErr, ok := err.(*EmailError); ok {
		return emailErr.Provider == ProviderAWSSES
	}
	return false
}

// IsBrevoError checks if the error is Brevo specific
func IsBrevoError(err error) bool {
	if emailErr, ok := err.(*EmailError); ok {
		return emailErr.Provider == ProviderBrevo
	}
	return false
}

// GetErrorCode extracts the error code from an EmailError
func GetErrorCode(err error) string {
	if emailErr, ok := err.(*EmailError); ok {
		return emailErr.Code
	}
	return ErrCodeUnknownError
}

// GetErrorProvider extracts the provider from an EmailError
func GetErrorProvider(err error) EmailProvider {
	if emailErr, ok := err.(*EmailError); ok {
		return emailErr.Provider
	}
	return ""
}

// FormatEmailError formats an EmailError for logging or display
func FormatEmailError(err error) string {
	if emailErr, ok := err.(*EmailError); ok {
		if emailErr.Original != nil {
			return fmt.Sprintf("[%s:%s] %s: %v",
				emailErr.Provider, emailErr.Code, emailErr.Message, emailErr.Original)
		}
		return fmt.Sprintf("[%s:%s] %s",
			emailErr.Provider, emailErr.Code, emailErr.Message)
	}
	return err.Error()
}
