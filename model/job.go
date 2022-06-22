package model

import (
	"fmt"

	"github.com/google/uuid"
)

type Job struct{}

type JobData struct {
	ID             uuid.UUID
	Name           string
	Description    string
	Repo           string
	Branch         string
	RepoSecretType string
	RepoSecretRef  string
}

func (j *Job) GetAll() []JobData {
	test_uuid, err := uuid.Parse("4aa1f842-e934-4c61-8bcc-b9562c02b220")
	if err != nil {
		fmt.Errorf("failed to parse uuid: %s", err)
	}
	test_uuid2, err := uuid.Parse("ee48530e-3319-4c0d-b928-86f1e70cde92")
	if err != nil {
		fmt.Errorf("failed to parse uuid: %s", err)
	}
	jobs := []JobData{
		{
			ID:             test_uuid2,
			Name:           "some random ops job 1",
			Description:    "this is some random ops job 1",
			Repo:           "git@github.com:peter-mcconnell/oneofmyautomationthingys.git",
			Branch:         "refs/heads/master",
			RepoSecretType: "vault",
			RepoSecretRef:  "kv-v1/keys/key-automationthingy",
		},
		{
			ID:             test_uuid,
			Name:           "some random ops job 2",
			Description:    "this is some random ops job 2",
			Repo:           "git@github.com:peter-mcconnell/oneofmyautomationthingys.git",
			Branch:         "refs/heads/master",
			RepoSecretType: "vault",
			RepoSecretRef:  "kv-v1/keys/key-automationthingy",
		},
	}
	return jobs
}

func (j *Job) GetOne(id uuid.UUID) JobData {
	jobs := j.GetAll()
	for i := 0; i < len(jobs); i++ {
		if jobs[i].ID == id {
			return jobs[i]
		}
	}
	return JobData{}
}
