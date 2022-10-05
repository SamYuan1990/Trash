package main

import (
	"os"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var (
	tmpDir, Bin                                      string
	Session0, Session1, Session2, Session3, Session4 *gexec.Session
	err                                              error
)

func TestConsensus(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Consensus Suite")
}

var _ = BeforeSuite(func() {
	tmpDir, err = os.MkdirTemp("", "test")
	Expect(err).NotTo(HaveOccurred())
	Bin, err = gexec.Build("main.go")
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterEach(func() {
	if Session0 != nil {
		Session0.Kill()
	}
	if Session1 != nil {
		Session1.Kill()
	}
	if Session2 != nil {
		Session2.Kill()
	}
	if Session3 != nil {
		Session3.Kill()
	}
	if Session4 != nil {
		Session4.Kill()
	}
})

var _ = AfterSuite(func() {
	os.RemoveAll(tmpDir)
	os.RemoveAll(Bin)
})
