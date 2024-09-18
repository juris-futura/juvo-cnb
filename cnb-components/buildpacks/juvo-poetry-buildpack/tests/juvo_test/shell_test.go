package juvo_test

import (
	"github.com/juris-futura/juvo-poetry-buildpack/juvo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Shell", func() {
	It("Executes a simple command", func() {
		ex := juvo.Executor{}
		err := ex.ExecuteStep(juvo.CommandDescriptor{Cmd: "ls"})
		Expect(err).ToNot(HaveOccurred())
	})
	Context("When we pass an environment", func() {
		It("Has access to the environment", func() {
			ex := juvo.Executor{}
			err := ex.ExecuteStep(juvo.CommandDescriptor{
				Cmd:  "bash",
				Args: []string{"-c", "echo $ENVVAR"},
				Env:  map[string]string{"ENVVAR": "Hello"},
			})
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
