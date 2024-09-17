package juvo

import (
	"fmt"
	"os"
	"os/exec"
)

type CommandDescriptor struct {
	Cmd  string
	Args []string
	Env  map[string]string
}

type Executable interface {
	MkCmd() (CommandDescriptor, error)
}

func (ctx CommandDescriptor) MkCmd() (CommandDescriptor, error) {
	return ctx, nil
}

func ExecuteStep(e Executable) error {
	var cd, err = e.MkCmd()
	if err != nil {
		return err
	}
	var cmd = exec.Command(cd.Cmd, cd.Args...)
	cmd.Env = os.Environ()

	for k, v := range cd.Env {
		var envvar = fmt.Sprintf("%s=%s", k, v)
		fmt.Println(envvar)
		cmd.Env = append(cmd.Env, envvar)
	}
	fmt.Println(cmd)
	return ExecuteCommand(cmd)
}

func ExecuteCommand(cmd *exec.Cmd) error {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
