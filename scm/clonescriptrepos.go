/*Logic for cloning script repos*/
package scm

import (
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/peter-mcconnell/automationthingy/types"
)

func CloneScriptRepos(scriptSources types.ScriptSources) error {
	var errs []string

	// git repos
	git := Git{}
	for _, source := range scriptSources.Git {
		b64dir := base64.StdEncoding.EncodeToString([]byte(source.Repo))
		dir := fmt.Sprintf("scripts/%s", b64dir)
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			err := git.Clone(source, dir)
			if err != nil {
				errs = append(errs, fmt.Sprintf(" - [error] failed cloning %s into %s. script id: %s", err.Error(), dir, source.Repo))
			}
		} else {
			fmt.Printf(" - skipping clone of %s as it already exists on disk\n", source.Repo)
		}
	}
	if len(errs) != 0 {
		return errors.New(strings.Join(errs, "\n"))
	}
	return nil
}
