package config

import "github.com/google/uuid"

type Config struct {
	Rbac    Rbac         `json:"rbac"`
	Scripts []ScriptData `json:"scripts"`
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
	Roles  []RbacRole  `json:"roles"`
	Groups []RbacGroup `json:"groups"`
}

type RbacRole struct {
	Name string `json:"name"`
}

type RbacGroup struct {
	Name  string   `json:"name"`
	Roles []string `json:"roles"`
}
