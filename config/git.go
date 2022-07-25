/*Logic for interacting with GIT repos*/
package config

import (
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"

	"github.com/peter-mcconnell/automationthingy/secretmgr"
)

type Git struct {
	Logger Logger
	Config *Config
}

func (g Git) Clone(source GitSource, dest string) error {
	g.Logger.Debugf("git cloning %s into %s", source.Repo, dest)
	var auth transport.AuthMethod
	if source.Secrettype != "" {
		// currently we only support basic private keys
		// TODO: PAC, passkey etc
		secretManager, err := secretmgr.GetSecretMgr(source.Secretref, &g.Config.Secretmgr)
		if err != nil {
			return err
		}
		secret, err := secretManager.Get(source.Secretref, "cert")
		if err != nil {
			return err
		}
		auth, err = ssh.NewPublicKeys("git", secret, "")
		if err != nil {
			return err
		}
	}
	g.Logger.Infof(" - cloning into %s", dest)
	_, err := git.PlainClone(
		dest,
		false,
		&git.CloneOptions{
			URL:           source.Repo,
			Auth:          auth,
			Progress:      os.Stdout,
			ReferenceName: plumbing.ReferenceName(source.Branch),
			SingleBranch:  true,
		},
	)
	return err
}
