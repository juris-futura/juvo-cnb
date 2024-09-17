package juvo

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type Deps struct {
	Metadata struct {
		Dependencies []struct {
			Name    string `toml:"name"`
			Version string `toml:"version"`
		} `toml:"dependencies"`
	} `toml:"metadata"`
}

type MetaInput struct {
	BuildpackMetadataPath string
}

type BPMetadata struct {
	PoetryVersion string
	PythonVersion string
}

func (input MetaInput) ReadMetadata() (BPMetadata, error) {
	fmt.Println("Reading Metadata File . . .")
	var file, err = os.Open(input.BuildpackMetadataPath)
	if err != nil {
		return BPMetadata{}, err
	}
	defer file.Close()

	fmt.Println("Decoding . . .")
	var m Deps
	_, err = toml.DecodeReader(file, &m)
	if err != nil {
		return BPMetadata{}, err
	}

	poetryVersion, err := readVersion("poetry", m)
	if err != nil {
		return BPMetadata{}, err
	}

	pythonVersion, err := readVersion("python", m)
	if err != nil {
		return BPMetadata{}, err
	}

	return BPMetadata{
		PoetryVersion: poetryVersion,
		PythonVersion: pythonVersion,
	}, nil
}

func readVersion(nm string, m Deps) (string, error) {
	var deps = m.Metadata.Dependencies
	for _, dep := range deps {
		if dep.Name == nm {
			return dep.Version, nil
		}
	}
	return "", fmt.Errorf("`%s` dependency not found", nm)
}
