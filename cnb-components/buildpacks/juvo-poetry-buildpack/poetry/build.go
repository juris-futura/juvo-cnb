package poetry

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/paketo-buildpacks/packit"
)

type Deps struct {
	Metadata struct {
		Dependencies []struct {
			Name    string `toml:"name"`
			Version string `toml:"version"`
		} `toml:"dependencies"`
	} `toml:"metadata"`
}

func Build() packit.BuildFunc {
	return func(ctx packit.BuildContext) (packit.BuildResult, error) {
		// Read the content of buildpack.toml. Well find poetry dep there
		fmt.Println("Reading Metadata File . . .")
		var file, err = os.Open(filepath.Join(ctx.CNBPath, "buildpack.toml"))
		if err != nil {
			return packit.BuildResult{}, err
		}

		fmt.Println("Decoding . . .")
		var m Deps
		_, err = toml.DecodeReader(file, &m)
		if err != nil {
			return packit.BuildResult{}, err
		}

		fmt.Println("Fetching Poetry Version . . .")
		poetryVersion, err := readPoetryVersion(m)
		if err != nil {
			return packit.BuildResult{}, err
		}
		fmt.Printf("Poetry Version: %s\n", poetryVersion)

		poetryLayer, err := ctx.Layers.Get("poetry")
		if err != nil {
			return packit.BuildResult{}, err
		}

		poetryLayer, err = poetryLayer.Reset()
		if err != nil {
			return packit.BuildResult{}, err
		}

		poetryLayer.Launch = true

		fmt.Println("Installing Virtual Env . . .")
		err = installVenv(poetryLayer.Path)
		if err != nil {
			return packit.BuildResult{}, err
		}

		return packit.BuildResult{
			Layers: []packit.Layer{poetryLayer},
		}, nil
	}
}

func installVenv(path string) error {
	privKey := os.Getenv("PRIV_SSH_ID")

	filename := filepath.Join(path, "priv_id")

	fd, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer fd.Close()
	defer os.Remove(filename)

	_, err = fd.WriteString(fmt.Sprintf("%s\n", privKey))
	if err != nil {
		return err
	}
	err = fd.Chmod(0600)
	if err != nil {
		return err
	}
	var gitCmd = fmt.Sprintf("GIT_SSH_COMMAND=ssh -i %s -o IdentitiesOnly=yes -o StrictHostKeyChecking=no -F /dev/null", filename)

	cmd := exec.Command("poetry", "install", "--only=main")
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, gitCmd)
	return RunCmd(cmd)
}

func RunCmd(cmd *exec.Cmd) error {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func readPoetryVersion(m Deps) (string, error) {
	var deps = m.Metadata.Dependencies
	for _, dep := range deps {
		if dep.Name == "poetry" {
			return dep.Version, nil
		}
	}
	return "", fmt.Errorf("`poetry` dependency not found")
}
