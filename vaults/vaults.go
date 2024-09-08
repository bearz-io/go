package vaults

import (
	"context"
	"time"
)

const (
	FEAT_PROPERTIES      = "properties"
	FEAT_VERSION         = "version"
	FEAT_EXPIRES_AT      = "expires_at"
	FEAT_TAGS            = "tags"
	FEAT_RESTORE_DELETED = "restore_deleted"
	FEAT_ROTATE          = "rotate"
	FEAT_ROTATE_INTERVAL = "rotate_interval"
	FEAT_BACKUP          = "backup"
)

func init() {
}

type SecretProperties interface {
	Key() string

	Enabled() bool

	ExpiresAt() *time.Time

	CreatedAt() *time.Time

	Version() *string

	Tags() map[string]*string
}

type SecretVault interface {
	HasFeature(name string) bool

	Get(key string) string

	GetAsync(ctx context.Context, key string) string

	Set(key, value string) error

	SetAsync(ctx context.Context, key, value string) error

	SetProperties(key string, version *string, props *SecretProperties) error

	SetPropertiesAsync(ctx context.Context, key string, version *string, props *SecretProperties) error

	Delete(key string) error

	DeleteAsync(ctx context.Context, key string) error

	Has(key string) bool

	HasAsync(ctx context.Context, key string) bool

	Keys() []string

	KeysAsync(ctx context.Context) []string

	ListProperties() []SecretProperties

	ListPropertiesAsync(ctx context.Context) []SecretProperties
}
