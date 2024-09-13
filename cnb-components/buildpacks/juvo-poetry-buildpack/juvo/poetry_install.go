package juvo

import (
	"fmt"
	"os"
	"path/filepath"
)

type PoetryInstall struct {
	KeyFilePath    string
	CtxPath        string
	OnlyMain       bool
	FallbackEnvVar string
}

func (install PoetryInstall) MkCmd() (CommandDescriptor, error) {
	var env = make(map[string]string)
	var keyfile, err = install.GetKeyfile()
	if err != nil {
		return CommandDescriptor{}, err
	}
	var envSshCmd = fmt.Sprintf("ssh -i %s -o IdentitiesOnly=yes -o StrictHostKeyChecking=no -F /dev/null", keyfile)
	env["GIT_SSH_COMMAND"] = envSshCmd

	var args = []string{"install"}
	if install.OnlyMain {
		args = append(args, "--only=main")
	}

	return CommandDescriptor{
		Cmd:  "poetry",
		Args: args,
		Env:  env,
	}, nil
}

func (i PoetryInstall) GetKeyfile() (string, error) {
	// If the binding path exists, then that's the keyfile
	var bindingPath = i.KeyFilePath
	var _, err = os.Stat(bindingPath)
	if err == nil {
		return bindingPath, nil
	}
	if os.IsNotExist(err) {
		return i.MakeKeyFile()
	}
	return "", err
}

func (i PoetryInstall) MakeKeyFile() (string, error) {
	var privkey = os.Getenv(i.FallbackEnvVar)
	if privkey == "" {
		return "", fmt.Errorf("%s not found and %s not set", i.KeyFilePath, i.FallbackEnvVar)
	}

	var keyfile = filepath.Join(i.CtxPath, "privkey")
	var fd, err = os.Create(keyfile)
	if err != nil {
		return "", err
	}
	defer fd.Close()

	if _, err = fd.WriteString(fmt.Sprintf("%s\n", privkey)); err != nil {
		return "", err
	}

	if err = fd.Chmod(0600); err != nil {
		return "", err
	}

	return keyfile, nil
}
