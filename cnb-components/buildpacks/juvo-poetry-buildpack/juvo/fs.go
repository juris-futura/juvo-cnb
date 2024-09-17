package juvo

import (
	"os"
	"path/filepath"
)

type Fs interface {
	FileExists(string) bool
	ParseMetadataFromFile(string) (BPMetadata, error)
}

type PhysicalFs struct{}

func (_ PhysicalFs) FileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	return err == nil
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
