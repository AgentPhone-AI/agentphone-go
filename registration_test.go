package agentphone

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestRegistrationService_GetStatusAndDetails(t *testing.T) {
	var requests []string
	client, server := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		requests = append(requests, r.Method+" "+r.URL.Path)
		w.WriteHeader(http.StatusOK)
		if r.URL.Path == "/register/status" {
			w.Write([]byte(`{"campaign_status":"approved","message":"Registration approved"}`))
			return
		}
		w.Write([]byte(`{"campaign_status":"approved","brand":{"business_name":"Example Inc","ein":"12-3456789"},"campaign":{"use_case":"customer_care","message_samples":["Hello"]}}`))
	})
	defer server.Close()

	status, err := client.Registration.GetStatus(context.Background())
	if err != nil {
		t.Fatalf("GetStatus() error: %v", err)
	}
	reg, err := client.Registration.GetDetails(context.Background())
	if err != nil {
		t.Fatalf("GetDetails() error: %v", err)
	}
	if status.Status != "approved" || reg.Status != "approved" || reg.Brand == nil || reg.Brand.BusinessName != "Example Inc" || reg.Campaign == nil || reg.Campaign.MessageSamples[0] != "Hello" {
		t.Errorf("status = %+v, registration = %+v", status, reg)
	}
	want := []string{"GET /register/status", "GET /register"}
	if len(requests) != len(want) {
		t.Fatalf("requests = %v, want %v", requests, want)
	}
	for i := range want {
		if requests[i] != want[i] {
			t.Errorf("requests[%d] = %q, want %q", i, requests[i], want[i])
		}
	}
}

func TestRegistrationService_RegisterAndUpdate(t *testing.T) {
	var requests []string
	client, server := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		requests = append(requests, r.Method+" "+r.URL.Path)
		var body map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("decode request body: %v", err)
		}
		if body["business_name"] != "Example Inc" || body["use_cases"] == nil {
			t.Errorf("request body = %v", body)
		}
		if r.Method == http.MethodPost && (body["message_samples"] == nil || body["phone_number"] != "+15551234567") {
			t.Errorf("registration POST body = %v", body)
		}
		if _, ok := body["legalName"]; ok {
			t.Errorf("request body = %v, unexpectedly contains legacy legalName", body)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"pending"}`))
	})
	defer server.Close()

	if _, err := client.Registration.Register(context.Background(), &RegisterA2PParams{
		RegistrationType: "sole_proprietor",
		FirstName:        "Ada",
		LastName:         "Lovelace",
		Email:            "ada@example.com",
		PhoneNumber:      "+15551234567",
		StreetAddress:    "1 Main St",
		City:             "San Francisco",
		State:            "CA",
		ZipCode:          "94105",
		Description:      "We send service updates to customers who opted in.",
		UseCases:         []string{"customer_care"},
		MessageSamples:   []string{"Example Inc: your support request is updated."},
		ConsentFlow:      "Customers opt in on our website before receiving messages.",
		OptInImageURL:    "https://example.com/opt-in.png",
		BusinessName:     "Example Inc",
	}); err != nil {
		t.Fatalf("Register() error: %v", err)
	}
	businessName := "Example Inc"
	useCaseList := []string{"customer_care"}
	if _, err := client.Registration.Update(context.Background(), &UpdateA2PRegistrationParams{BusinessName: &businessName, UseCases: &useCaseList}); err != nil {
		t.Fatalf("Update() error: %v", err)
	}
	want := []string{"POST /register", "PATCH /register"}
	if len(requests) != len(want) {
		t.Fatalf("requests = %v, want %v", requests, want)
	}
	for i := range want {
		if requests[i] != want[i] {
			t.Errorf("requests[%d] = %q, want %q", i, requests[i], want[i])
		}
	}
}
