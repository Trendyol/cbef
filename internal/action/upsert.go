package action

import (
	"context"
	"fmt"
	"time"

	"github.com/couchbase/gocb/v2"
	"github.com/trendyol/cbef/internal/model"
)

// Upsert inserts or updates an eventing function.
func (a *action) Upsert(ctx context.Context, f *model.Function) error {
	s := gocb.EventingFunction{
		Name:               f.Name,
		Code:               f.Code,
		Version:            f.Version,
		EnforceSchema:      f.EnforceSchema,
		HandlerUUID:        f.HandlerUUID,
		FunctionInstanceID: f.FunctionInstanceID,
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
			CPPWorkerThreadCount:   f.Settings.CPPWorkerThreadCount,
			DCPStreamBoundary:      gocb.EventingFunctionDCPBoundary(f.Settings.DCPStreamBoundary),
			Description:            f.Settings.Description,
			DeploymentStatus:       gocb.EventingFunctionDeploymentStatus(f.Settings.DeploymentStatus),
			ProcessingStatus:       gocb.EventingFunctionProcessingStatus(f.Settings.ProcessingStatus),
			LanguageCompatibility:  gocb.EventingFunctionLanguageCompatibility(f.Settings.LanguageCompatibility),
			LogLevel:               gocb.EventingFunctionLogLevel(f.Settings.LogLevel),
			ExecutionTimeout:       f.Settings.ExecutionTimeout,
			LCBInstCapacity:        f.Settings.LCBInstCapacity,
			LCBRetryCount:          f.Settings.LCBRetryCount,
			LCBTimeout:             f.Settings.LCBTimeout,
			QueryConsistency:       gocb.QueryScanConsistency(f.Settings.QueryConsistency),
			NumTimerPartitions:     f.Settings.NumTimerPartitions,
			SockBatchSize:          f.Settings.SockBatchSize,
			TickDuration:           f.Settings.TickDuration,
			TimerContextSize:       f.Settings.TimerContextSize,
			UserPrefix:             f.Settings.UserPrefix,
			BucketCacheSize:        f.Settings.BucketCacheSize,
			BucketCacheAge:         f.Settings.BucketCacheAge,
			CurlMaxAllowedRespSize: f.Settings.CurlMaxAllowedRespSize,
			QueryPrepareAll:        f.Settings.QueryPrepareAll,
			WorkerCount:            f.Settings.WorkerCount,
			HandlerHeaders:         f.Settings.HandlerHeaders,
			HandlerFooters:         f.Settings.HandlerFooters,
			EnableAppLogRotation:   f.Settings.EnableAppLogRotation,
			AppLogDir:              f.Settings.AppLogDir,
			AppLogMaxSize:          f.Settings.AppLogMaxSize,
			AppLogMaxFiles:         f.Settings.AppLogMaxFiles,
			CheckpointInterval:     f.Settings.CheckpointInterval,
		},
	}

	for _, binding := range f.BucketBindings {
		s.BucketBindings = append(s.BucketBindings, gocb.EventingFunctionBucketBinding{
			Alias: binding.Alias,
			Name: gocb.EventingFunctionKeyspace{
				Bucket:     binding.Name.Bucket,
				Scope:      binding.Name.Scope,
				Collection: binding.Name.Collection,
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
		Timeout:       10 * time.Second,
		RetryStrategy: gocb.NewBestEffortRetryStrategy(nil),
		Context:       ctx,
	}); err != nil {
		return fmt.Errorf("failed to upsert function: %w", err)
	}

	return nil
}
