package models

type StateResponse struct {
	TTL            int             `json:"ttl" yaml:"ttl"`
	Authorizations []Authorization `json:"authorizations" yaml:"authorizations"`
	Identity       string          `json:"identity" yaml:"identity"`
	IdentityURL    string          `json:"identity_url" yaml:"identity_url"`
}
