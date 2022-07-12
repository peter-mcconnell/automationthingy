package model

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/peter-mcconnell/automationthingy/config"
	"github.com/peter-mcconnell/automationthingy/types"
)

type Script struct {
	Config config.Config
}

func (s *Script) GetAll() []types.ScriptData {
	return []types.ScriptData{}
}

func (s *Script) GetOne(id uuid.UUID) (types.ScriptData, error) {
	if _, ok := s.Config.ScriptIndex[id]; ok {
		return s.Config.Scripts[s.Config.ScriptIndex[id]], nil
	}
	return types.ScriptData{}, fmt.Errorf("no script found with id: %s", id)
}
