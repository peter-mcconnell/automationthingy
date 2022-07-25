package secretmgr

import (
	"fmt"
)

type SecretMgr interface {
	Get(path, key string) ([]byte, error)
}

func GetSecretMgr(secretRef string, config *ConfigSecretMgr) (SecretMgr, error) {
	// we only support vault currently
	secretManager := &Vault{
		Config: &config.Vault,
	}
	if secretRef == "" {
		return secretManager, fmt.Errorf("unsupported secret manager type: %s", secretRef)
	}
	return secretManager, nil
}
