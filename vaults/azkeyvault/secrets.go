package azkeyvault

import (
	"context"
	"errors"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/azsecrets"
	"github.com/bearz-io/vaults"
)

type AzSecretVaultOptions struct {
	Token        *azcore.TokenCredential
	Identity     bool   `json:"identity" yaml:"identity"`
	Uri          string `json:"uri" yaml:"uri"`
	ClientId     string `json:"clientId" yaml:"clientId"`
	ClientSecret string `json:"clientSecret" yaml:"clientSecret"`
	TenantId     string `json:"tenantId" yaml:"tenantId"`
}

type AzSecretsVault struct {
	options   AzSecretVaultOptions
	client    *azsecrets.Client
	vaultName string
}

func NewAzSecretsVault(options AzSecretVaultOptions) *AzSecretsVault {
	vaultName := options.Uri
	if strings.HasPrefix(vaultName, "https://") {
		vaultName = strings.TrimPrefix(vaultName, "https://")
	}

	if strings.HasSuffix(vaultName, "/") {
		vaultName = strings.TrimSuffix(vaultName, "/")
	}
	if strings.HasSuffix(vaultName, "vault.azure.net") {
		vaultName = strings.TrimSuffix(vaultName, "vault.azure.net")
	}

	return &AzSecretsVault{
		options:   options,
		vaultName: vaultName,
	}
}

func (s *AzSecretsVault) VaultName() string {
	return s.vaultName
}

func (s *AzSecretsVault) ProviderName() string {
	return "azkv"
}

func (s *AzSecretsVault) GetSecretValue(key string, params *vaults.GetSecretValueParams) (*string, error) {
	ctx := params.Context
	if ctx == nil {
		ctx = context.Background()
	}

	version := ""
	if params != nil && params.Version != "" {
		version = params.Version
	}

	secret, err := s.client.GetSecret(ctx, key, version, nil)
	if err != nil {
		return nil, err
	}

	value := secret.Value
	return value, nil
}

func (s *AzSecretsVault) GetSecretProperties(key string, params *vaults.GetSecretPropertiesParams) (*vaults.SecretProperties, error) {
	return nil, nil
}

func (s *AzSecretsVault) SetSecretValue(key string, value string, params *vaults.SetSecretValueParams) error {
	ctx := params.Context
	if ctx == nil {
		ctx = context.Background()
	}

	p := azsecrets.SetSecretParameters{
		Value: &value,
	}

	_, err := s.client.SetSecret(ctx, key, p, nil)
	if err != nil {
		return err
	}

	return nil
}

func (s *AzSecretsVault) SetSecretProperties(key string, params *vaults.SetSecretPropertiesParams) error {
	return nil
}

func (s *AzSecretsVault) DeleteSecret(key string, params *vaults.DeleteSecretParams) error {
	ctx := params.Context
	if ctx == nil {
		ctx = context.Background()
	}

	_, err := s.client.DeleteSecret(ctx, key, nil)
	if err != nil {
		return err
	}

	return nil
}

func (s *AzSecretsVault) ListSecretNames(params *vaults.ListSecretsParams) ([]string, error) {
	ctx := params.Context
	if ctx == nil {
		ctx = context.Background()
	}

	names := []string{}
	pager := s.client.NewListSecretPropertiesPager(nil)
	errs := make([]error, 0)
	for pager.More() {
		page, err := pager.NextPage(context.TODO())
		if err != nil {
			errs = append(errs, err)
		}
		for _, secret := range page.Value {
			names = append(names, secret.ID.Name())
		}
	}

	if len(errs) > 0 {
		return names, errors.Join(errs...)
	}

	return names, nil
}
