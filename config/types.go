package config

import (
	"github.com/google/uuid"
	"github.com/peter-mcconnell/automationthingy/secretmgr"
)

type Config struct {
	Rbac        Rbac                      `json:"rbac"`
	Scripts     []Script                  `json:"scripts"`
	Sources     Sources                   `json:"sources"`
	Secretmgrs  SecretMgrs                `json:"secretmgrs"`
	General     General                   `json:"general"`
	Secretmgr   secretmgr.ConfigSecretMgr `json:"secretmgr"`
	Logger      Logger
	ScriptIndex map[uuid.UUID]int
}

type SourceConfig struct {
	Scripts []Script `json:"scripts"`
}

type Logger interface {
	Infof(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Fatalf(format string, v ...interface{})
	Debugf(format string, v ...interface{})
	Info(args ...interface{})
	Error(args ...interface{})
	Warn(args ...interface{})
	Fatal(args ...interface{})
	Debug(args ...interface{})
}

type SecretMgrs struct {
	Vault Vault `json:"vault"`
}

type Vault struct {
	Addr          string `json:"addr"`
	TokenFilePath string `json:"tokenfilepath"`
}

type Sources struct {
	Git  []GitSource  `json:"git,omitempty"`
	Disk []DiskSource `json:"disk,omitempty"`
}
type GitSource struct {
	Repo       string `json:"repo"`
	Branch     string `json:"branch"`
	Secrettype string `json:"secrettype,omitempty"`
	Secretref  string `json:"secretref,omitempty"`
	Rbac       []Rbac `json:"rbac"`
}

type DiskSource struct {
	Path string `json:"path"`
}

type SourceScriptData struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Desc       string    `json:"desc"`
	Workdir    string    `json:"workdir"`
	Command    string    `json:"command"`
	Categories []string  `json:"categories"`
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

type KubernetesJob struct {
	Namespace               string `json:"namespace,omitempty"`
	Ttlsecondsafterfinished int32  `json:"ttlsecondsafterfinished,omitempty"`
	Image                   string `json:"image,omitempty"`
	Cluster                 string `json:"cluster,omitempty"`
	Imagepullpolicy         string `json:"imagepullpolicy,omitempty"`
	Restartpolicy           string `json:"restartpolicy,omitempty"`
	Backofflimit            int32  `json:"backofflimit,omitempty"`
	Parallelism             int32  `json:"parallelism,omitempty"`
	Completions             int32  `json:"completions,omitempty"`
}

type Script struct {
	ID            uuid.UUID     `json:"id"`
	Name          string        `json:"name"`
	Desc          string        `json:"desc"`
	Source        ScriptSource  `json:"source"`
	Rbac          []ScriptRbac  `json:"rbac"`
	Categories    []string      `json:"categories"`
	Command       []string      `json:"command"`
	Workdir       string        `json:"workdir,omitempty"`
	Kubernetesjob KubernetesJob `json:"kubernetesjob,omitempty"`
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

type GithubAuth struct {
	Clientid    string `json:"clientid"`
	Redirecturi string `json:"redirecturi"`
	Secrettype  string `json:"secrettype"`
	Secretref   string `json:"secretref"`
}

type Auth struct {
	Github GithubAuth `json:"github,omitempty"`
}

type Api struct {
	Port string `json:"port"`
	Host string `json:"host"`
	Auth Auth   `json:"auth,omitempty"`
}

type Web struct {
	Port string `json:"port"`
	Host string `json:"host"`
}

type General struct {
	Web Web `json:"web"`
	Api Api `json:"api"`
}
