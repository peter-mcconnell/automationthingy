/*Logic for interacting with GIT repos*/
package scm

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"

	"github.com/peter-mcconnell/automationthingy/config"
	"github.com/peter-mcconnell/automationthingy/secretmgr"
)

type Git struct {
	script config.ScriptData
}

func (g Git) Clone(dir string) error {
	var auth transport.AuthMethod
	if g.script.Source.Git.Secrettype != "" {
		// currently we only support basic private keys
		// TODO: PAC, passkey etc
		secretManager, err := secretmgr.GetSecretMgr(g.script.Source.Git.Secretref)
		if err != nil {
			return err
		}
		secret, err := secretManager.Get(g.script.Source.Git.Secretref)
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
			URL:           g.script.Source.Git.Repo,
			Auth:          auth,
			Progress:      os.Stdout,
			ReferenceName: plumbing.ReferenceName(g.script.Source.Git.Branch),
			SingleBranch:  true,
		},
	)
	return err
}
