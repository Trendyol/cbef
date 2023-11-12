package model

import "time"

type Function struct {
	Cluster            Cluster           `json:"cluster"`
	Name               string            `json:"name"`
	Code               string            `json:"-"`
	Version            string            `json:"version"`
	EnforceSchema      bool              `json:"enforce_schema"`
	HandlerUUID        int               `json:"handler_uuid"`
	FunctionInstanceID string            `json:"function_instance_id"`
	MetadataKeyspace   Keyspace          `json:"metadata_keyspace"`
	SourceKeyspace     Keyspace          `json:"source_keyspace"`
	BucketBindings     []BucketBinding   `json:"bucket_bindings"`
	URLBindings        []URLBinding      `json:"url_bindings"`
	ConstantBindings   []ConstantBinding `json:"constant_bindings"`
	Settings           Settings          `json:"settings"`
}

type Cluster struct {
	ConnectionString string `json:"connection_string"`
	User             string `json:"user"`
	Pass             string `json:"pass"`
}

type Keyspace struct {
	Bucket     string `json:"bucket"`
	Scope      string `json:"scope"`
	Collection string `json:"collection"`
}

type BucketBinding struct {
	Alias  string   `json:"alias"`
	Name   Keyspace `json:"name"`
	Access string   `json:"access"`
}

type URLBinding struct {
	Hostname               string `json:"hostname"`
	Alias                  string `json:"alias"`
	Auth                   Auth   `json:"auth"`
	AllowCookies           bool   `json:"allow_cookies"`
	ValidateSSLCertificate bool   `json:"validate_ssl_certificate"`
}

type Auth struct {
	Type  string `json:"type"`
	User  string `json:"user"`
	Pass  string `json:"pass"`
	Token string `json:"token"`
}

type ConstantBinding struct {
	Alias   string `json:"alias"`
	Literal string `json:"literal"`
}

type Settings struct {
	CPPWorkerThreadCount   int           `json:"cpp_worker_thread_count"`
	DCPStreamBoundary      string        `json:"dcp_stream_boundary"`
	Description            string        `json:"description"`
	DeploymentStatus       bool          `json:"deployment_status"`
	ProcessingStatus       bool          `json:"processing_status"`
	LanguageCompatibility  string        `json:"language_compatibility"`
	LogLevel               string        `json:"log_level"`
	ExecutionTimeout       time.Duration `json:"execution_timeout"`
	LCBInstCapacity        int           `json:"lcb_inst_capacity"`
	LCBRetryCount          int           `json:"lcb_retry_count"`
	LCBTimeout             time.Duration `json:"lcb_timeout"`
	QueryConsistency       int           `json:"query_consistency"`
	NumTimerPartitions     int           `json:"num_timer_partitions"`
	SockBatchSize          int           `json:"sock_batch_size"`
	TickDuration           time.Duration `json:"tick_duration"`
	TimerContextSize       int           `json:"timer_context_size"`
	UserPrefix             string        `json:"user_prefix"`
	BucketCacheSize        int           `json:"bucket_cache_size"`
	BucketCacheAge         int           `json:"bucket_cache_age"`
	CurlMaxAllowedRespSize int           `json:"curl_max_allowed_resp_size"`
	QueryPrepareAll        bool          `json:"query_prepare_all"`
	WorkerCount            int           `json:"worker_count"`
	HandlerHeaders         []string      `json:"handler_headers"`
	HandlerFooters         []string      `json:"handler_footers"`
	EnableAppLogRotation   bool          `json:"enable_app_log_rotation"`
	AppLogDir              string        `json:"app_log_dir"`
	AppLogMaxSize          int           `json:"app_log_max_size"`
	AppLogMaxFiles         int           `json:"app_log_max_files"`
	CheckpointInterval     time.Duration `json:"checkpoint_interval"`
}
