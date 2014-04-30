package wamp_test

import (
	. "github.com/llchan/go-wamp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Wamp", func() {
	Context("Calling Dummy()", func() {
		It("Should return )", func() {
			Expect(Dummy()).To(Equal(0))
		})
	})
})
