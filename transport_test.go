package wamp_test

import (
	. "github.com/llchan/go-wamp"
	"github.com/llchan/go-wamp/transports/inproc"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Wamp", func() {
	Describe("Inproc Transport", func() {
		Context("Sending Hello", func() {
			It("should work", func() {
				t1, t2 := inproc.NewTransportPair(16)
				var m Message = &HelloMessage{Realm: "Test"}
				Expect(t1.Send(m)).ToNot(HaveOccurred())
				Expect(t2.Recv()).To(Equal(m))
			})
		})
	})
})
