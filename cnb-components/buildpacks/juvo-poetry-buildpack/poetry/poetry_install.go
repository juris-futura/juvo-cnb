package poetry

import (
	"fmt"
	"os"
)

type PoetryInstall struct {
	KeyFilePath string
	OnlyMain    bool
}

func (install PoetryInstall) MkCmd() (CommandDescriptor, error) {
	var args = []string{"install"}
	if install.OnlyMain {
		args = append(args, "--only=main")
	}

	var env = make(map[string]string)
	var keyfile = install.KeyFilePath
	if _, err := os.Stat(keyfile); err != nil {
		return CommandDescriptor{}, err
	}
	var envSshCmd = fmt.Sprintf("ssh -i %s -o IdentitiesOnly=yes -o StrictHostKeyChecking=no -F /dev/null", keyfile)
	env["GIT_SSH_COMMAND"] = envSshCmd

	return CommandDescriptor{
		Cmd:  "poetry",
		Args: args,
		Env:  env,
	}, nil
}
