package secretmgr

type Vault struct {
	Config *ConfigVault
}

type ConfigVault struct {
	Token string `json:"token"`
	Addr  string `json:"addr"`
}

type ConfigSecretMgr struct {
	Vault ConfigVault `json:"vault"`
}
