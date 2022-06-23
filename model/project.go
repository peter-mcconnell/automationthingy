package model

import (
	"github.com/google/uuid"
	"github.com/peter-mcconnell/automationthingy/config"
)

type Project struct{}

type ProjectData struct {
	ID          uuid.UUID
	Name        string
	Description string
	Scripts     []config.ScriptData
}

func (p *Project) GetAll() []ProjectData {
	script := Script{}
	all_scripts := script.GetAll()
	projects := []ProjectData{
		{
			ID:          uuid.New(),
			Name:        "some random ops job 1",
			Description: "this is some random ops job 1",
			Scripts:     all_scripts,
		},
		{
			ID:          uuid.New(),
			Name:        "some random ops job 2",
			Description: "this is some random ops job 2",
			Scripts:     all_scripts,
		},
	}
	return projects
}

func (p *Project) GetOne(id uuid.UUID) ProjectData {
	all_projects := p.GetAll()
	return all_projects[0]
}
