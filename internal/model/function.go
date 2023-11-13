package model

import "time"

// Function represents an eventing function.
type Function struct {
	Cluster          Cluster           `json:"cluster"`
	Code             string            `json:"-"`
	Name             string            `json:"name"`
	MetadataKeyspace Keyspace          `json:"metadata_keyspace"`
	SourceKeyspace   Keyspace          `json:"source_keyspace"`
	BucketBindings   []BucketBinding   `json:"bucket_bindings"`
	URLBindings      []URLBinding      `json:"url_bindings"`
	ConstantBindings []ConstantBinding `json:"constant_bindings"`
	Settings         Settings          `json:"settings"`
}

// Cluster represents an couchbase cluster.
type Cluster struct {
	ConnectionString string `json:"connection_string"`
	User             string `json:"user"`
	Pass             string `json:"pass"`
}

// Keyspace represents a triple of bucket, collection, and scope names.
type Keyspace struct {
	Bucket     string `json:"bucket"`
	Scope      string `json:"scope"`
	Collection string `json:"collection"`
}

// BucketBinding represents an eventing function binding allowing the function access to buckets, scopes, and collections.
type BucketBinding struct {
	Alias      string `json:"alias"`
	Bucket     string `json:"bucket"`
	Scope      string `json:"scope"`
	Collection string `json:"collection"`
	Access     string `json:"access"`
}

// URLBinding represents an eventing function binding allowing the function access external resources via cURL.
type URLBinding struct {
	Hostname               string `json:"hostname"`
	Alias                  string `json:"alias"`
	AllowCookies           bool   `json:"allow_cookies"`
	ValidateSSLCertificate bool   `json:"validate_ssl_certificate"`
	Auth                   Auth   `json:"auth"`
}

// Auth represents an authentication method for URLBinding for an eventing function.
type Auth struct {
	Type  string `json:"type"`
	User  string `json:"user"`
	Pass  string `json:"pass"`
	Token string `json:"token"`
}

// ConstantBinding represents an eventing function binding allowing the function to utilize global variables.
type ConstantBinding struct {
	Alias   string `json:"alias"`
	Literal string `json:"literal"`
}

// Settings are the settings for an eventing Function.
type Settings struct {
	DCPStreamBoundary     string        `json:"dcp_stream_boundary"`
	Description           string        `json:"description"`
	LogLevel              string        `json:"log_level"`
	QueryConsistency      uint          `json:"query_consistency"`
	WorkerCount           uint          `json:"worker_count"`
	LanguageCompatibility string        `json:"language_compatibility"`
	ExecutionTimeout      time.Duration `json:"execution_timeout"`
	TimerContextSize      uint          `json:"timer_context_size"`
}
