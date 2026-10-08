package agentphone

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestFeedbackService_Submit(t *testing.T) {
	client, server := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/feedback" {
			t.Errorf("request = %s %s, want POST /feedback", r.Method, r.URL.Path)
		}
		var body map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("decode feedback body: %v", err)
		}
		if body["category"] != "docs_mismatch" || body["note"] != "The request field name differs from the OpenAPI schema." {
			t.Errorf("feedback body = %v", body)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"feedback_id":"fb_1","already_submitted":false,"request_id":"req_1"}`))
	})
	defer server.Close()

	response, err := client.Feedback.Submit(context.Background(), &SubmitFeedbackParams{
		Category:  "docs_mismatch",
		Note:      "The request field name differs from the OpenAPI schema.",
		RequestID: "req_1",
	})
	if err != nil {
		t.Fatalf("Submit() error: %v", err)
	}
	if response.FeedbackID != "fb_1" || response.AlreadySubmitted || response.RequestID != "req_1" {
		t.Errorf("response = %+v", response)
	}
}
