package juvo_test

import (
	"testing"

	"github.com/juris-futura/juvo-poetry-buildpack/juvo"
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

func (fs MockFs) FileExists(filepath string) (bool, error) {
	for k := range fs.Files {
		if k == filepath {
			return true, nil
		}
	}
	return false, nil
}

func (fs MockFs) ParseMetadataFromFile(_ string) (juvo.BPMetadata, error) {
	return juvo.BPMetadata{}, nil
}

type MockExecutor struct{}

func (_ MockExecutor) ExecuteStep(_ juvo.Executable) error {
	return nil
}
