package scm

import (
	"errors"
	"fmt"
	"strings"

	"github.com/peter-mcconnell/automationthingy/model"
)

func CloneJobRepos() error {
	jobModel := model.Job{}
	jobs := jobModel.GetAll()
	var errs []string
	for _, job := range jobs {
		source_control := GetScm(job)
		dir := fmt.Sprintf("scripts/%s", job.ID)
		err := source_control.Clone(dir)
		if err != nil {
			errs = append(errs, fmt.Sprintf(" - [error] failed cloning %s into %s. job id: %s", err.Error(), dir, job.ID))
		}
	}
	if len(errs) != 0 {
		return errors.New(strings.Join(errs, "\n"))
	}
	return nil
}
