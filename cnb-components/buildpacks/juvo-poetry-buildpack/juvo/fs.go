package juvo

import (
	"os"
)

type Fs interface {
	FileExists(string) bool
}

type PhysicalFs struct{}

func (_ PhysicalFs) FileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	return err == nil
}
