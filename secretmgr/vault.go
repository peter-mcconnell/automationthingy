package secretmgr

import (
	"fmt"
	"net/http"
	"time"

	"github.com/hashicorp/vault/api"
)

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

func (v *Vault) Get(path, key string) ([]byte, error) {
	var noresp []byte
	if key == "" {
		return noresp, fmt.Errorf("can not get blank key from secret manager")
	}
	token := v.Config.Token
	vaultAddr := v.Config.Addr
	client, err := api.NewClient(&api.Config{Address: vaultAddr, HttpClient: httpClient})
	if err != nil {
		return noresp, nil
	}
	client.SetToken(token)

	data, err := client.Logical().Read(path)
	if err != nil {
		return noresp, nil
	}
	if data == nil {
		return []byte{}, fmt.Errorf("no secrets found at path: %s", path)
	}
	if val, ok := data.Data[key]; !ok {
		return []byte{}, fmt.Errorf("%s not found in vault secret path %s", key, path)
	} else {
		return []byte(val.(string)), nil
	}
}
