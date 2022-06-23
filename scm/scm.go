package scm

import "github.com/peter-mcconnell/automationthingy/config"

type Scm interface {
	Clone(string) error
}

func GetScm(script config.ScriptData) Scm {
	// we only support git currently
	return Git{
		script: script,
	}
}
