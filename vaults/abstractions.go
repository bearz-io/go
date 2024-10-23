package vaults

import (
	"context"
	"time"
)

type GetSecretValueParams struct {
	Version string
	Context context.Context
}

type SetSecretValueParams struct {
	Context context.Context
}

type SetSecretPropertiesParams struct {
	Context context.Context

	Tags      *map[string]*string
	ExpiresAt *TimeSetter
}

type GetSecretPropertiesParams struct {
	Context context.Context
}

type SecretProperties struct {
	Name      *string
	CreatedAt *time.Time
	Version   *string
	Tags      *map[string]*string
	ExpiresAt *time.Time
}

type TimeSetter struct {
	Time  *time.Time
	Unset bool
}

type ListSecretsParams struct {
	Context context.Context
}

type DeleteSecretParams struct {
	Context context.Context
}

type SecretVault interface {
	VaultName() string

	ProviderName() string

	GetSecretValue(key string, params *GetSecretValueParams) (*string, error)

	GetSecretProperties(key string, params *GetSecretPropertiesParams) (*SecretProperties, error)

	SetSecretValue(key string, value string, params *SetSecretValueParams) error

	SetSecretProperties(key string, params *SetSecretPropertiesParams) error

	DeleteSecret(key string, params *DeleteSecretParams) error

	ListSecretNames(params *ListSecretsParams) ([]string, error)
}
