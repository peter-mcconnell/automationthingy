package scm

import "github.com/peter-mcconnell/automationthingy/model"

type Scm interface {
	Clone(string) error
}

func GetScm(job model.JobData) Scm {
	// we only support git currently
	return Git{
		job: job,
	}
}
