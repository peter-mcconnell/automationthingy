package model

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/peter-mcconnell/automationthingy/config"
)

type Script struct {
	Config config.Config
}

func (s *Script) GetAll() []config.Script {
	return s.Config.Scripts
}

func (s *Script) GetOne(id uuid.UUID) (config.Script, error) {
	if _, ok := s.Config.ScriptIndex[id]; ok {
		return s.Config.Scripts[s.Config.ScriptIndex[id]], nil
	}
	return config.Script{}, fmt.Errorf("no script found with id: %s", id)
}
