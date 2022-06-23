/*Logic for cloning script repos*/
package scm

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/peter-mcconnell/automationthingy/config"
)

func CloneScriptRepos(scripts []config.ScriptData) error {
	var errs []string
	for _, script := range scripts {
		source_control := GetScm(script)
		dir := fmt.Sprintf("scripts/%s", script.ID)
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			err := source_control.Clone(dir)
			if err != nil {
				errs = append(errs, fmt.Sprintf(" - [error] failed cloning %s into %s. script id: %s", err.Error(), dir, script.ID))
			}
		} else {
			fmt.Printf(" - skipping clone of %s as it already exists on disk\n", script.ID)
		}
	}
	if len(errs) != 0 {
		return errors.New(strings.Join(errs, "\n"))
	}
	return nil
}
