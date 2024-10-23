package keyring

import (
	"github.com/99designs/keyring"
	kr "github.com/99designs/keyring"
	"github.com/bearz-io/go/errors"
	"github.com/bearz-io/vaults"
)

type KeyRingSecretVault struct {
	config kr.Config
}

type KeyRingSecretVaultOptions struct {
	Service string `json:"service" yaml:"service"`
}

func NewKeyRingSecretVault(options KeyRingSecretVaultOptions) *KeyRingSecretVault {
	return &KeyRingSecretVault{
		config: kr.Config{
			ServiceName: options.Service,
			AllowedBackends: []keyring.BackendType{
				keyring.WinCredBackend,
				keyring.KeychainBackend,
				keyring.SecretServiceBackend,
				keyring.PassBackend,
				keyring.FileBackend,
			},
		},
	}
}

func (s *KeyRingSecretVault) VaultName() string {
	return s.config.ServiceName
}

func (s *KeyRingSecretVault) ProviderName() string {
	return "keyring"
}

func (s *KeyRingSecretVault) GetSecretValue(key string, params *vaults.GetSecretValueParams) (*string, error) {
	k, err := keyring.Open(s.config)
	item, err := k.Get(key)
	if err != nil {
		return nil, err
	}

	value := string(item.Data)
	return &value, nil
}

func (s *KeyRingSecretVault) SetSecretValue(key string, value string, params *vaults.SetSecretValueParams) error {
	k, err := keyring.Open(s.config)
	if err != nil {
		return err
	}

	return k.Set(keyring.Item{
		Key:  key,
		Data: []byte(value),
	})
}

func (s *KeyRingSecretVault) DeleteSecret(key string, params *vaults.DeleteSecretParams) error {
	k, err := keyring.Open(s.config)
	if err != nil {
		return err
	}

	return k.Remove(key)
}

func (s *KeyRingSecretVault) ListSecretNames(params *vaults.ListSecretsParams) ([]string, error) {
	k, err := keyring.Open(s.config)
	if err != nil {
		return nil, err
	}

	names, err := k.Keys()
	if err != nil {
		return nil, err
	}

	return names, nil
}

func (s *KeyRingSecretVault) GetSecretProperties(key string, params *vaults.GetSecretPropertiesParams) (*vaults.SecretProperties, error) {
	return nil, errors.ErrNotImplemented.WithStack().WithMessage("get secret properties is not implemented for keyring")
}

func (s *KeyRingSecretVault) SetSecretProperties(key string, params *vaults.SetSecretPropertiesParams) error {
	return errors.ErrNotImplemented.WithStack().WithMessage("get secret properties is not implemented for keyring")
}
