package secretmgr

import (
	"fmt"
	"net/http"
	"time"

	"github.com/hashicorp/vault/api"
)

type Vault struct{}

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

func (v *Vault) Get(key string) ([]byte, error) {
	var noresp []byte
	if key == "" {
		return noresp, fmt.Errorf("can not get blank key from secret manager")
	}
	token := "root"
	vaultAddr := "http://10.43.56.96:8200"
	client, err := api.NewClient(&api.Config{Address: vaultAddr, HttpClient: httpClient})
	if err != nil {
		return noresp, nil
	}
	client.SetToken(token)

	data, err := client.Logical().Read(key)
	if err != nil {
		return noresp, nil
	}
	return []byte(data.Data["cert"].(string)), nil
}
