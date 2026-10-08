package agentphone

import (
	"context"
	"net/http"
)

// FeedbackService handles POST /feedback.
type FeedbackService struct {
	client *Client
}

// SubmitFeedbackParams describes an issue or suggestion about the API.
type SubmitFeedbackParams struct {
	Category  string   `json:"category"`
	Note      string   `json:"note"`
	RequestID string   `json:"request_id,omitempty"`
	URL       string   `json:"url,omitempty"`
	Tags      []string `json:"tags,omitempty"`
}

// SubmitFeedbackResponse is returned after feedback is stored or deduplicated.
type SubmitFeedbackResponse struct {
	FeedbackID       string `json:"feedback_id"`
	AlreadySubmitted bool   `json:"already_submitted"`
	RequestID        string `json:"request_id,omitempty"`
}

// Submit sends feedback to AgentPhone.
func (s *FeedbackService) Submit(ctx context.Context, params *SubmitFeedbackParams) (*SubmitFeedbackResponse, error) {
	var response SubmitFeedbackResponse
	err := s.client.request(ctx, http.MethodPost, "/feedback", params, &response)
	return &response, err
}
