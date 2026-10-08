package agentphone

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestVerificationService_SendAndCheck(t *testing.T) {
	var requests []string
	client, server := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		requests = append(requests, r.Method+" "+r.URL.Path)
		var body map[string]string
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("decode request body: %v", err)
		}
		if r.URL.Path == "/verify/send" && body["phone_number"] != "+14155550100" {
			t.Errorf("phone_number = %q", body["phone_number"])
		}
		w.WriteHeader(http.StatusOK)
		if r.URL.Path == "/verify/send" {
			w.Write([]byte(`{"status":"pending","verification_id":"ver_1"}`))
			return
		}
		if body["verification_id"] != "ver_1" {
			t.Errorf("verification_id = %q, want ver_1", body["verification_id"])
		}
		if body["code"] != "123456" {
			t.Errorf("code = %q, want 123456", body["code"])
		}
		w.Write([]byte(`{"result":"ok","attempts_remaining":2}`))
	})
	defer server.Close()

	sent, err := client.Verification.Send(context.Background(), &SendVerificationParams{PhoneNumber: "+14155550100"})
	if err != nil {
		t.Fatalf("Send() error: %v", err)
	}
	checked, err := client.Verification.Check(context.Background(), &CheckVerificationParams{VerificationID: "ver_1", Code: "123456"})
	if err != nil {
		t.Fatalf("Check() error: %v", err)
	}
	if sent.Status != "pending" || sent.VerificationID != "ver_1" || checked.Result != "ok" || checked.AttemptsRemaining == nil || *checked.AttemptsRemaining != 2 {
		t.Errorf("send response = %+v, check response = %+v", sent, checked)
	}
	want := []string{"POST /verify/send", "POST /verify/check"}
	if len(requests) != len(want) {
		t.Fatalf("requests = %v, want %v", requests, want)
	}
	for i := range want {
		if requests[i] != want[i] {
			t.Errorf("requests[%d] = %q, want %q", i, requests[i], want[i])
		}
	}
}
