package juvo

import (
	"fmt"
	"io"

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

type BPMetadata struct {
	PoetryVersion string
	PythonVersion string
}

func ReadMetadata(r io.Reader) (BPMetadata, error) {
	fmt.Println("Decoding . . .")
	var m Deps
	_, err := toml.DecodeReader(r, &m)
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
