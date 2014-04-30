package wamp_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestWamp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Wamp Suite")
}
