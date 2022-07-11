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
	for i := 0; i < len(s.Config.Scripts); i++ {
		fmt.Printf(">> %s\n", s.Config.Scripts[i].ID)
		if s.Config.Scripts[i].ID == id {
			return s.Config.Scripts[i], nil
		}
	}
	return types.ScriptData{}, fmt.Errorf("no script found with id: %s", id)
}
