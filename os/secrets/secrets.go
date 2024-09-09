package secrets

import (
	"github.com/99designs/keyring"
)

const TEST = "TEST"

func init() {
}

var config = keyring.Config{
	ServiceName: "bearz",
	AllowedBackends: []keyring.BackendType{
		keyring.WinCredBackend,
		keyring.KeychainBackend,
		keyring.SecretServiceBackend,
		keyring.PassBackend,
		keyring.FileBackend,
	},
}

func SetKeyringConfig(c *keyring.Config) {
	config = *c
}

func cloneConfig(service string) *keyring.Config {
	return &keyring.Config{
		ServiceName:                    service,
		AllowedBackends:                config.AllowedBackends,
		KeychainName:                   config.KeychainName,
		KeychainTrustApplication:       config.KeychainTrustApplication,
		KeychainSynchronizable:         config.KeychainSynchronizable,
		KeychainAccessibleWhenUnlocked: config.KeychainAccessibleWhenUnlocked,
		KeychainPasswordFunc:           config.KeychainPasswordFunc,
		FilePasswordFunc:               config.FilePasswordFunc,
		FileDir:                        config.FileDir,
		KeyCtlScope:                    config.KeyCtlScope,
		KeyCtlPerm:                     config.KeyCtlPerm,
		KWalletAppID:                   config.KWalletAppID,
		LibSecretCollectionName:        config.LibSecretCollectionName,
		PassDir:                        config.PassDir,
		PassCmd:                        config.PassCmd,
		PassPrefix:                     config.PassPrefix,
		WinCredPrefix:                  config.WinCredPrefix,
	}
}

func ListSecretsNames(service string) ([]string, error) {
	c := cloneConfig(service)
	k, err := keyring.Open(*c)
	if err != nil {
		return nil, err
	}

	return k.Keys()
}

func GetSecret(service, name string) (string, error) {
	c := cloneConfig(service)
	k, err := keyring.Open(*c)
	if err != nil {
		return "", err
	}

	item, err := k.Get(name)
	if err != nil {
		return "", err
	}

	return string(item.Data), nil
}

func GetSecretBytes(service, name string) ([]byte, error) {
	c := cloneConfig(service)
	k, err := keyring.Open(*c)
	if err != nil {
		return nil, err
	}

	item, err := k.Get(name)
	if err != nil {
		return nil, err
	}

	return item.Data, nil
}

func SetSecret(service, name, value string) error {
	c := cloneConfig(service)
	k, err := keyring.Open(*c)
	if err != nil {
		return err
	}

	return k.Set(keyring.Item{
		Key:  name,
		Data: []byte(value),
	})
}

func SetSecretBytes(service, name string, value []byte) error {
	c := cloneConfig(service)
	k, err := keyring.Open(*c)
	if err != nil {
		return err
	}

	return k.Set(keyring.Item{
		Key:  name,
		Data: value,
	})
}

func DeleteSecret(service, name string) error {
	c := cloneConfig(service)
	k, err := keyring.Open(*c)
	if err != nil {
		return err
	}

	return k.Remove(name)
}
