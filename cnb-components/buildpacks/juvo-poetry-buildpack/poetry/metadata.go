package poetry

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

type Metadata struct {
	PoetryVersion string
}

func (input MetaInput) ReadMetadata() (Metadata, error) {
	fmt.Println("Reading Metadata File . . .")
	var file, err = os.Open(input.BuildpackMetadataPath)
	if err != nil {
		return Metadata{}, err
	}
	defer file.Close()

	fmt.Println("Decoding . . .")
	var m Deps
	_, err = toml.DecodeReader(file, &m)
	if err != nil {
		return Metadata{}, err
	}

	poetryVersion, err := readPoetryVersion(m)
	if err != nil {
		return Metadata{}, err
	}

	return Metadata{PoetryVersion: poetryVersion}, nil
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
