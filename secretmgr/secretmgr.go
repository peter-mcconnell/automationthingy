package secretmgr

import (
	"fmt"
)

type SecretMgr interface {
	Get(key string) ([]byte, error)
}

func GetSecretMgr(secretRef string) (SecretMgr, error) {
	// we only support vault currently
	secretManager := &Vault{}
	if secretRef == "" {
		return secretManager, fmt.Errorf("unsupported secret manager type: %s", secretRef)
	}
	return secretManager, nil
}
