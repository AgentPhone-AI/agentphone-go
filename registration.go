package agentphone

import (
	"context"
	"net/http"
)

// RegistrationService handles the /register endpoints: US A2P 10DLC registration, required to send SMS to US numbers at scale, reference -> docs.agentphone.ai/error-handling's. This number is not registered for US A2P 10DLC messaging" error, and its note that raising a brand's daily volume limit requires brand vetting
type RegistrationService struct {
	client *Client
}

// RegistrationStatus is the response from GetStatus.
type RegistrationStatus struct {
	Status       string             `json:"campaign_status,omitempty"`
	Message      string             `json:"message,omitempty"`
	Stage        string             `json:"stage,omitempty"`
	Issue        *RegistrationIssue `json:"issue,omitempty"`
	SubAccountID string             `json:"sub_account_id,omitempty"`
}

// GetStatus returns a lightweight registration status check.
func (s *RegistrationService) GetStatus(ctx context.Context) (*RegistrationStatus, error) {
	var status RegistrationStatus
	err := s.client.request(ctx, http.MethodGet, "/register/status", nil, &status)
	return &status, err
}

// RegistrationBrand identifies the business sending messages.
type RegistrationBrand struct {
	BusinessName   string `json:"business_name,omitempty"`
	WebsiteURL     string `json:"website_url,omitempty"`
	Industry       string `json:"industry,omitempty"`
	EntityType     string `json:"entity_type,omitempty"`
	EIN            string `json:"ein,omitempty"`
	Email          string `json:"email,omitempty"`
	Phone          string `json:"phone,omitempty"`
	ContactName    string `json:"contact_name,omitempty"`
	StreetAddress  string `json:"street_address,omitempty"`
	StreetAddress2 string `json:"street_address_2,omitempty"`
	City           string `json:"city,omitempty"`
	State          string `json:"state,omitempty"`
	ZipCode        string `json:"zip_code,omitempty"`
	Country        string `json:"country,omitempty"`
}

// RegistrationCampaign describes the messaging use case being registered.
type RegistrationCampaign struct {
	UseCase        string   `json:"use_case,omitempty"`
	Description    string   `json:"description,omitempty"`
	MessageSamples []string `json:"message_samples,omitempty"`
	ConsentFlow    string   `json:"consent_flow,omitempty"`
	OptInImageURL  string   `json:"opt_in_image_url,omitempty"`
	Includes       []string `json:"includes,omitempty"`
}

// Registration is the account's full A2P 10DLC registration record.
type Registration struct {
	Status        string                `json:"campaign_status,omitempty"`
	Stage         string                `json:"stage,omitempty"`
	Issue         *RegistrationIssue    `json:"issue,omitempty"`
	SubAccountID  string                `json:"sub_account_id,omitempty"`
	Brand         *RegistrationBrand    `json:"brand,omitempty"`
	Campaign      *RegistrationCampaign `json:"campaign,omitempty"`
	OK            bool                  `json:"ok,omitempty"`
	Message       string                `json:"message,omitempty"`
	FeeCents      int                   `json:"fee_cents,omitempty"`
	UpdatedFields []string              `json:"updated_fields,omitempty"`
}

// RegistrationIssue explains why a registration needs changes.
type RegistrationIssue struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Since   string `json:"since,omitempty"`
}

// GetDetails returns the full registration record (brand and campaign details, not just the status).
func (s *RegistrationService) GetDetails(ctx context.Context) (*Registration, error) {
	var reg Registration
	err := s.client.request(ctx, http.MethodGet, "/register", nil, &reg)
	return &reg, err
}

// RegisterA2PParams contains the flat fields required to submit A2P 10DLC registration.
type RegisterA2PParams struct {
	RegistrationType  string   `json:"registration_type"`
	FirstName         string   `json:"first_name"`
	LastName          string   `json:"last_name"`
	Email             string   `json:"email"`
	PhoneNumber       string   `json:"phone_number"`
	BusinessName      string   `json:"business_name,omitempty"`
	BrandName         string   `json:"brand_name,omitempty"`
	OrganizationType  string   `json:"organization_type,omitempty"`
	ContactTitle      string   `json:"contact_title,omitempty"`
	ContactTitleOther string   `json:"contact_title_other,omitempty"`
	EIN               string   `json:"ein,omitempty"`
	WebsiteURL        string   `json:"website_url,omitempty"`
	Industry          string   `json:"industry,omitempty"`
	BusinessRegions   []string `json:"business_regions,omitempty"`
	StreetAddress     string   `json:"street_address"`
	StreetAddress2    string   `json:"street_address_2,omitempty"`
	City              string   `json:"city"`
	State             string   `json:"state"`
	ZipCode           string   `json:"zip_code"`
	Country           string   `json:"country,omitempty"`
	PlannedNumbers    *int     `json:"planned_numbers,omitempty"`
	Description       string   `json:"description"`
	UseCases          []string `json:"use_cases"`
	Volume            string   `json:"volume,omitempty"`
	MessageSamples    []string `json:"message_samples"`
	ConsentFlow       string   `json:"consent_flow"`
	PrivacyPolicyURL  string   `json:"privacy_policy_url,omitempty"`
	TermsOfServiceURL string   `json:"terms_of_service_url,omitempty"`
	OptInImageURL     string   `json:"opt_in_image_url"`
	Includes          []string `json:"includes,omitempty"`
	LinkSample        string   `json:"link_sample,omitempty"`
	CreateSubAccount  *bool    `json:"create_sub_account,omitempty"`
}

// Register starts or submits A2P 10DLC registration for the account.
func (s *RegistrationService) Register(ctx context.Context, params *RegisterA2PParams) (*Registration, error) {
	var reg Registration
	err := s.client.request(ctx, http.MethodPost, "/register", params, &reg)
	return &reg, err
}

// UpdateA2PRegistrationParams contains optional fields for updating a registration.
type UpdateA2PRegistrationParams struct {
	FirstName         *string   `json:"first_name,omitempty"`
	LastName          *string   `json:"last_name,omitempty"`
	Email             *string   `json:"email,omitempty"`
	PhoneNumber       *string   `json:"phone_number,omitempty"`
	BusinessName      *string   `json:"business_name,omitempty"`
	BrandName         *string   `json:"brand_name,omitempty"`
	OrganizationType  *string   `json:"organization_type,omitempty"`
	ContactTitle      *string   `json:"contact_title,omitempty"`
	ContactTitleOther *string   `json:"contact_title_other,omitempty"`
	EIN               *string   `json:"ein,omitempty"`
	WebsiteURL        *string   `json:"website_url,omitempty"`
	Industry          *string   `json:"industry,omitempty"`
	BusinessRegions   *[]string `json:"business_regions,omitempty"`
	StreetAddress     *string   `json:"street_address,omitempty"`
	StreetAddress2    *string   `json:"street_address_2,omitempty"`
	City              *string   `json:"city,omitempty"`
	State             *string   `json:"state,omitempty"`
	ZipCode           *string   `json:"zip_code,omitempty"`
	Country           *string   `json:"country,omitempty"`
	Description       *string   `json:"description,omitempty"`
	UseCases          *[]string `json:"use_cases,omitempty"`
	Volume            *string   `json:"volume,omitempty"`
	MessageSamples    *[]string `json:"message_samples,omitempty"`
	ConsentFlow       *string   `json:"consent_flow,omitempty"`
	PrivacyPolicyURL  *string   `json:"privacy_policy_url,omitempty"`
	TermsOfServiceURL *string   `json:"terms_of_service_url,omitempty"`
	OptInImageURL     *string   `json:"opt_in_image_url,omitempty"`
	Includes          *[]string `json:"includes,omitempty"`
	LinkSample        *string   `json:"link_sample,omitempty"`
}

// Update updates the account's A2P 10DLC registration details.
func (s *RegistrationService) Update(ctx context.Context, params *UpdateA2PRegistrationParams) (*Registration, error) {
	var reg Registration
	err := s.client.request(ctx, http.MethodPatch, "/register", params, &reg)
	return &reg, err
}
