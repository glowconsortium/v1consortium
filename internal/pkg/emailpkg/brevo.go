package emailpkg

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// BrevoSender implements EmailSender for Brevo (formerly SendinBlue)
type BrevoSender struct {
	client     *http.Client
	config     *EmailConfig
	apiBaseURL string
}

// BrevoEmailRequest represents the request structure for Brevo API
type BrevoEmailRequest struct {
	Sender      BrevoEmailAddress   `json:"sender"`
	To          []BrevoEmailAddress `json:"to"`
	CC          []BrevoEmailAddress `json:"cc,omitempty"`
	BCC         []BrevoEmailAddress `json:"bcc,omitempty"`
	Subject     string              `json:"subject,omitempty"`
	TextContent string              `json:"textContent,omitempty"`
	HTMLContent string              `json:"htmlContent,omitempty"`
	Headers     map[string]string   `json:"headers,omitempty"`
}

// BrevoTemplateEmailRequest represents the request structure for Brevo template API
type BrevoTemplateEmailRequest struct {
	Sender     BrevoEmailAddress      `json:"sender"`
	To         []BrevoEmailAddress    `json:"to"`
	CC         []BrevoEmailAddress    `json:"cc,omitempty"`
	BCC        []BrevoEmailAddress    `json:"bcc,omitempty"`
	TemplateID int64                  `json:"templateId"`
	Params     map[string]interface{} `json:"params,omitempty"`
	Headers    map[string]string      `json:"headers,omitempty"`
}

// BrevoEmailAddress represents an email address in Brevo format
type BrevoEmailAddress struct {
	Email string `json:"email"`
	Name  string `json:"name,omitempty"`
}

// BrevoEmailResponse represents the response from Brevo API
type BrevoEmailResponse struct {
	MessageID string `json:"messageId"`
}

// BrevoErrorResponse represents an error response from Brevo API
type BrevoErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// NewBrevoSender creates a new Brevo email sender
func NewBrevoSender(config *EmailConfig) (*BrevoSender, error) {
	if config == nil {
		return nil, NewEmailError(ProviderBrevo, ErrCodeInvalidConfig, "Brevo configuration is required")
	}

	apiBaseURL := config.BrevoAPIURL
	if apiBaseURL == "" {
		apiBaseURL = "https://api.brevo.com/v3"
	}

	client := &http.Client{
		Timeout: config.Timeout,
	}

	return &BrevoSender{
		client:     client,
		config:     config,
		apiBaseURL: apiBaseURL,
	}, nil
}

// SendEmail sends a basic email using Brevo API
func (s *BrevoSender) SendEmail(ctx context.Context, message *EmailMessage) (*EmailResponse, error) {
	// Convert to Brevo format
	brevoRequest := BrevoEmailRequest{
		Sender:      s.convertToBrevoAddress(message.From),
		To:          s.convertToBrevoAddresses(message.To),
		Subject:     message.Subject,
		TextContent: message.TextBody,
		HTMLContent: message.HTMLBody,
		Headers:     message.Headers,
	}

	if len(message.CC) > 0 {
		brevoRequest.CC = s.convertToBrevoAddresses(message.CC)
	}
	if len(message.BCC) > 0 {
		brevoRequest.BCC = s.convertToBrevoAddresses(message.BCC)
	}

	// Send request
	response, err := s.sendBrevoRequest(ctx, "/smtp/email", brevoRequest)
	if err != nil {
		return nil, err
	}

	var brevoResponse BrevoEmailResponse
	if err := json.Unmarshal(response, &brevoResponse); err != nil {
		return nil, NewEmailErrorWithCause(ProviderBrevo, ErrCodeResponseParseError,
			MsgResponseParseError, err)
	}

	return &EmailResponse{
		MessageID: brevoResponse.MessageID,
		Status:    "sent",
		Provider:  ProviderBrevo,
		SentAt:    time.Now(),
	}, nil
}

// SendTemplateEmail sends a template-based email using Brevo API
func (s *BrevoSender) SendTemplateEmail(ctx context.Context, message *TemplateEmailMessage) (*EmailResponse, error) {
	// Parse template ID
	var templateID int64
	if message.TemplateID != "" {
		if _, err := fmt.Sscanf(message.TemplateID, "%d", &templateID); err != nil {
			return nil, NewEmailErrorWithCause(ProviderBrevo, ErrCodeBrevoInvalidTemplateID,
				MsgBrevoInvalidTemplateID, err)
		}
	} else {
		return nil, NewEmailError(ProviderBrevo, ErrCodeBrevoMissingTemplateID,
			MsgBrevoMissingTemplateID)
	}

	// Convert to Brevo format
	brevoRequest := BrevoTemplateEmailRequest{
		Sender:     s.convertToBrevoAddress(message.From),
		To:         s.convertToBrevoAddresses(message.To),
		TemplateID: templateID,
		Params:     message.Variables,
		Headers:    message.Headers,
	}

	if len(message.CC) > 0 {
		brevoRequest.CC = s.convertToBrevoAddresses(message.CC)
	}
	if len(message.BCC) > 0 {
		brevoRequest.BCC = s.convertToBrevoAddresses(message.BCC)
	}

	// Send request
	response, err := s.sendBrevoRequest(ctx, "/smtp/email", brevoRequest)
	if err != nil {
		return nil, err
	}

	var brevoResponse BrevoEmailResponse
	if err := json.Unmarshal(response, &brevoResponse); err != nil {
		return nil, &EmailError{
			Provider: ProviderBrevo,
			Code:     "RESPONSE_PARSE_ERROR",
			Message:  "failed to parse Brevo response",
			Original: err,
		}
	}

	return &EmailResponse{
		MessageID: brevoResponse.MessageID,
		Status:    "sent",
		Provider:  ProviderBrevo,
		SentAt:    time.Now(),
	}, nil
}

// ValidateConfig validates the Brevo configuration
func (s *BrevoSender) ValidateConfig() error {
	if s.config.BrevoAPIKey == "" {
		return NewEmailError(ProviderBrevo, ErrCodeBrevoMissingAPIKey, MsgBrevoMissingAPIKey)
	}

	if s.config.DefaultFromEmail == "" {
		return NewEmailError(ProviderBrevo, ErrCodeMissingFromEmail, MsgMissingFromEmail)
	}

	return nil
}

// GetProvider returns the provider type
func (s *BrevoSender) GetProvider() EmailProvider {
	return ProviderBrevo
}

// sendBrevoRequest sends a request to the Brevo API
func (s *BrevoSender) sendBrevoRequest(ctx context.Context, endpoint string, payload interface{}) ([]byte, error) {
	// Marshal payload
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, NewEmailErrorWithCause(ProviderBrevo, ErrCodeRequestMarshalError,
			MsgRequestMarshalError, err)
	}

	// Create request
	url := s.apiBaseURL + endpoint
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, NewEmailErrorWithCause(ProviderBrevo, ErrCodeRequestCreateError,
			MsgRequestCreateError, err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("api-key", s.config.BrevoAPIKey)

	// Send request
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, NewEmailErrorWithCause(ProviderBrevo, ErrCodeRequestSendError,
			MsgRequestSendError, err)
	}
	defer resp.Body.Close()

	// Read response
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, NewEmailErrorWithCause(ProviderBrevo, ErrCodeResponseReadError,
			MsgResponseReadError, err)
	}

	// Check status code
	if resp.StatusCode >= 400 {
		var errorResp BrevoErrorResponse
		if err := json.Unmarshal(responseBody, &errorResp); err == nil {
			return nil, NewEmailError(ProviderBrevo, errorResp.Code, errorResp.Message)
		}
		return nil, NewEmailError(ProviderBrevo, ErrCodeBrevoHTTPError,
			fmt.Sprintf("HTTP %d: %s", resp.StatusCode, string(responseBody)))
	}

	return responseBody, nil
}

// convertToBrevoAddress converts EmailAddress to BrevoEmailAddress
func (s *BrevoSender) convertToBrevoAddress(addr EmailAddress) BrevoEmailAddress {
	return BrevoEmailAddress{
		Email: addr.Email,
		Name:  addr.Name,
	}
}

// convertToBrevoAddresses converts slice of EmailAddress to slice of BrevoEmailAddress
func (s *BrevoSender) convertToBrevoAddresses(addrs []EmailAddress) []BrevoEmailAddress {
	brevoAddrs := make([]BrevoEmailAddress, len(addrs))
	for i, addr := range addrs {
		brevoAddrs[i] = s.convertToBrevoAddress(addr)
	}
	return brevoAddrs
}
