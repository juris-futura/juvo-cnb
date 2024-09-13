package juvo_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestJuvo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Poetry Suite")
}

type MockFs struct {
	Files map[string]string
}

func (fs MockFs) FileExists(filepath string) bool {
	for k := range fs.Files {
		if k == filepath {
			return true
		}
	}
	return false
}
