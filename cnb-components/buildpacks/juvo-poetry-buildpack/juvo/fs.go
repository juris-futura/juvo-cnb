package juvo

import (
	"os"
	"path/filepath"
)

type Fs interface {
	FileExists(string) (bool, error)
	ParseMetadataFromFile(string) (BPMetadata, error)
}

type PhysicalFs struct{}

func (_ PhysicalFs) FileExists(filepath string) (bool, error) {
	_, err := os.Stat(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (_ PhysicalFs) ParseMetadataFromFile(path string) (BPMetadata, error) {
	// Read the content of buildpack.toml. Well find poetry dep there
	file, err := os.Open(filepath.Join(path, "buildpack.toml"))
	if err != nil {
		return BPMetadata{}, err
	}
	defer file.Close()
	return ReadMetadata(file)
}
