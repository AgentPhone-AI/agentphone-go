package agentphone

import (
	"context"
	"net/http"
)

// VerificationService handles the /verify endpoints: one-time-code phone number verification (send a code, then check what the user entered).
type VerificationService struct {
	client *Client
}

// SendVerificationParams are the parameters for sending a verification code.
type SendVerificationParams struct {
	PhoneNumber string `json:"phone_number"`
}

// SendVerificationResponse is the response from Send.
type SendVerificationResponse struct {
	Status         string `json:"status,omitempty"`
	VerificationID string `json:"verification_id"`
}

// Send, sends a one-time verification code to a phone number.
func (s *VerificationService) Send(ctx context.Context, params *SendVerificationParams) (*SendVerificationResponse, error) {
	var resp SendVerificationResponse
	err := s.client.request(ctx, http.MethodPost, "/verify/send", params, &resp)
	return &resp, err
}

// CheckVerificationParams are the parameters for checking a code the user entered.
type CheckVerificationParams struct {
	VerificationID string `json:"verification_id"`
	Code           string `json:"code"`
}

// CheckVerificationResponse is the response from Check.
type CheckVerificationResponse struct {
	Result            string `json:"result"`
	AttemptsRemaining *int   `json:"attempts_remaining,omitempty"`
}

// Check verifies the code a user entered against the one most recently sent to their number.
func (s *VerificationService) Check(ctx context.Context, params *CheckVerificationParams) (*CheckVerificationResponse, error) {
	var resp CheckVerificationResponse
	err := s.client.request(ctx, http.MethodPost, "/verify/check", params, &resp)
	return &resp, err
}
