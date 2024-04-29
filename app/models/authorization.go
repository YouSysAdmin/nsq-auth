package models

type Authorization struct {
	Topic       string   `json:"topic" yaml:"topic"`
	Channels    []string `json:"channels" yaml:"channels"`
	Permissions []string `json:"permissions" yaml:"permissions"`
}
