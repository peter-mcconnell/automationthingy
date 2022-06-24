package types

import "github.com/google/uuid"

type SecretMgrs struct {
	Vault Vault `json:"vault"`
}

type Vault struct {
	Addr          string `json:"addr"`
	TokenFilePath string `json:"tokenfilepath"`
}

type ScriptSources struct {
	Git  []GitScriptSource  `json:"git,omitempty"`
	Disk []DiskScriptSource `json:"disk,omitempty"`
}
type GitScriptSource struct {
	Repo       string `json:"repo"`
	Branch     string `json:"branch"`
	Secrettype string `json:"secrettype,omitempty"`
	Secretref  string `json:"secretref,omitempty"`
	Rbac       []Rbac `json:"rbac"`
}

type DiskScriptSource struct {
	Path string `json:"path"`
}

type SourceConfig struct {
	Scripts []SourceScriptData `json:"scripts"`
}

type SourceScriptData struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Desc    string    `json:"desc"`
	Workdir string    `json:"workdir"`
	Command string    `json:"command"`
}

type GitSource struct {
	Repo       string `json:"repo"`
	Branch     string `json:"branch"`
	Secrettype string `json:"secrettype,omitempty"`
	Secretref  string `json:"secretref,omitempty"`
}

type DiskSource struct {
	Path string `json:"path"`
}

type ScriptSource struct {
	Git  GitSource  `json:"git,omitempty"`
	Disk DiskSource `json:"disk,omitempty"`
}

type ScriptRbacRole struct {
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
}

type ScriptRbac struct {
	Roles []ScriptRbacRole `json:"roles"`
}

type ScriptData struct {
	ID         uuid.UUID    `json:"id"`
	Name       string       `json:"name"`
	Desc       string       `json:"desc"`
	Source     ScriptSource `json:"source"`
	Rbac       []ScriptRbac `json:"rbac"`
	Categories []string     `json:"categories"`
}

type Rbac struct {
	Roles            []RbacRole         `json:"roles"`
	Grouprolemapping []RbacGroupMapping `json:"grouprolemapping"`
	Userrolemapping  []RbacUserMapping  `json:"userrolemapping"`
}

type RbacRole struct {
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
}

type RbacGroupMapping struct {
	Name  string   `json:"name"`
	Roles []string `json:"roles"`
}

type RbacUserMapping struct {
	Name  string   `json:"name"`
	Roles []string `json:"roles"`
}
