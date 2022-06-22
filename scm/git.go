package scm

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"

	"github.com/peter-mcconnell/automationthingy/model"
	"github.com/peter-mcconnell/automationthingy/secretmgr"
)

type Git struct {
	job model.JobData
}

func (g Git) Clone(dir string) error {
	var auth transport.AuthMethod
	if g.job.RepoSecretType != "" {
		// currently we only support basic private keys
		// TODO: PAC, passkey etc
		secretManager, err := secretmgr.GetSecretMgr(g.job.RepoSecretRef)
		if err != nil {
			return err
		}
		secret, err := secretManager.Get(g.job.RepoSecretRef)
		if err != nil {
			return err
		}
		auth, err = ssh.NewPublicKeys("git", secret, "")
		if err != nil {
			return err
		}
	}
	fmt.Printf(" - cloning into %s\n", dir)
	_, err := git.PlainClone(
		dir,
		false,
		&git.CloneOptions{
			URL:           g.job.Repo,
			Auth:          auth,
			Progress:      os.Stdout,
			ReferenceName: plumbing.ReferenceName(g.job.Branch),
			SingleBranch:  true,
		},
	)
	return err
}
