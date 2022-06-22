package executor

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/peter-mcconnell/automationthingy/model"
)

type LocalExecutor struct {
	ID string
}

func (e *LocalExecutor) Execute() {
	job_model := model.Job{}
	job_uuid, err := uuid.Parse(e.ID)
	if err != nil {
		fmt.Errorf("failed to parse uuid %s", e.ID)
	}
	job := job_model.GetOne(job_uuid)

	fmt.Println(job.Name)
	fmt.Println(job.Description)
	fmt.Println(job.Repo)
}
