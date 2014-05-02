package wamp_test

import (
	. "github.com/llchan/go-wamp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Wamp", func() {
	Describe("Message Parts", func() {
		var v MessagePart
		Context("Integer", func() {
			It("Should work for 0", func() {
				v = Integer(0)
				Expect(v.Value()).To(Equal(Integer(0)))
				Expect(v.Validate()).ToNot(HaveOccurred())
			})
		})

		Context("Id", func() {
			It("Should work for 0", func() {
				v = Id(0)
				Expect(v.Validate()).ToNot(HaveOccurred())
			})
			It("Should not work for MaxId + 1", func() {
				v = Id(MaxId + 1)
				Expect(v.Validate()).To(HaveOccurred())
				Expect(MaxId).To(BeEquivalentTo(1 << 53))
			})
		})

		Context("Uri", func() {
			It("Should work for abc", func() {
				v = Uri("abc")
				Expect(v.Validate()).ToNot(HaveOccurred())
			})
			It("Should work for abc.def.ghi", func() {
				v = Uri("abc.def.ghi")
				Expect(v.Validate()).ToNot(HaveOccurred())
			})
			It("Should not work for .abc.def", func() {
				v = Uri(".abc.def")
				Expect(v.Validate()).To(HaveOccurred())
			})
			It("Should not work for abc.#def", func() {
				v = Uri("abc.#def")
				Expect(v.Validate()).To(HaveOccurred())
			})
			It("Should not work for a bc.def", func() {
				v = Uri("a bc.def")
				Expect(v.Validate()).To(HaveOccurred())
			})
			It("Should not work for abc..def", func() {
				v = Uri("abc..def")
				Expect(v.Validate()).To(HaveOccurred())
			})
		})
	})
})
