package model

import "github.com/google/uuid"

type Project struct{}

type ProjectData struct {
	ID          uuid.UUID
	Name        string
	Description string
	Jobs        []JobData
}

func (p *Project) GetAll() []ProjectData {
	job := Job{}
	all_jobs := job.GetAll()
	projects := []ProjectData{
		{
			ID:          uuid.New(),
			Name:        "some random ops job 1",
			Description: "this is some random ops job 1",
			Jobs:        all_jobs,
		},
		{
			ID:          uuid.New(),
			Name:        "some random ops job 2",
			Description: "this is some random ops job 2",
			Jobs:        all_jobs,
		},
	}
	return projects
}

func (p *Project) GetOne(id uuid.UUID) ProjectData {
	all_projects := p.GetAll()
	return all_projects[0]
}
