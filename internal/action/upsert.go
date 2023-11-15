package action

import (
	"context"
	"fmt"

	"github.com/couchbase/gocb/v2"
	"github.com/trendyol/cbef/internal/model"
)

// Upsert inserts or updates an eventing function.
func (a *action) Upsert(ctx context.Context, f *model.Function) error {
	s := gocb.EventingFunction{
		Name: f.Name,
		Code: f.Code,
		MetadataKeyspace: gocb.EventingFunctionKeyspace{
			Bucket:     f.MetadataKeyspace.Bucket,
			Scope:      f.MetadataKeyspace.Scope,
			Collection: f.MetadataKeyspace.Collection,
		},
		SourceKeyspace: gocb.EventingFunctionKeyspace{
			Bucket:     f.SourceKeyspace.Bucket,
			Scope:      f.SourceKeyspace.Scope,
			Collection: f.SourceKeyspace.Collection,
		},
		Settings: gocb.EventingFunctionSettings{
			DCPStreamBoundary:     gocb.EventingFunctionDCPBoundary(f.Settings.DCPStreamBoundary),
			Description:           f.Settings.Description,
			LogLevel:              gocb.EventingFunctionLogLevel(f.Settings.LogLevel),
			QueryConsistency:      gocb.QueryScanConsistency(f.Settings.QueryConsistency),
			WorkerCount:           int(f.Settings.WorkerCount),
			LanguageCompatibility: gocb.EventingFunctionLanguageCompatibility(f.Settings.LanguageCompatibility),
			ExecutionTimeout:      f.Settings.ExecutionTimeout,
			TimerContextSize:      int(f.Settings.TimerContextSize),
			DeploymentStatus:      gocb.EventingFunctionDeploymentStatusDeployed,
			ProcessingStatus:      gocb.EventingFunctionProcessingStatusRunning,
		},
	}

	for _, binding := range f.BucketBindings {
		s.BucketBindings = append(s.BucketBindings, gocb.EventingFunctionBucketBinding{
			Alias: binding.Alias,
			Name: gocb.EventingFunctionKeyspace{
				Bucket:     binding.Bucket,
				Scope:      binding.Scope,
				Collection: binding.Collection,
			},
			Access: gocb.EventingFunctionBucketAccess(binding.Access),
		})
	}

	for _, binding := range f.URLBindings {
		b := gocb.EventingFunctionUrlBinding{
			Hostname:               binding.Hostname,
			Alias:                  binding.Alias,
			Auth:                   gocb.EventingFunctionUrlNoAuth{},
			AllowCookies:           binding.AllowCookies,
			ValidateSSLCertificate: binding.ValidateSSLCertificate,
		}

		if binding.Auth.Type == "basic" {
			b.Auth = gocb.EventingFunctionUrlAuthBasic{
				User: binding.Auth.User,
				Pass: binding.Auth.Pass,
			}
		}

		if binding.Auth.Type == "digest" {
			b.Auth = gocb.EventingFunctionUrlAuthDigest{
				User: binding.Auth.User,
				Pass: binding.Auth.Pass,
			}
		}

		if binding.Auth.Type == "bearer" {
			b.Auth = gocb.EventingFunctionUrlAuthBearer{
				BearerKey: binding.Auth.Token,
			}
		}

		s.UrlBindings = append(s.UrlBindings, b)
	}

	for _, binding := range f.ConstantBindings {
		s.ConstantBindings = append(s.ConstantBindings, gocb.EventingFunctionConstantBinding{
			Alias:   binding.Alias,
			Literal: binding.Literal,
		})
	}

	if err := a.cluster.EventingFunctions().UpsertFunction(s, &gocb.UpsertEventingFunctionOptions{
		Timeout:       a.timeout,
		RetryStrategy: gocb.NewBestEffortRetryStrategy(nil),
		Context:       ctx,
	}); err != nil {
		return fmt.Errorf("failed to upsert function: %w", err)
	}

	return nil
}
