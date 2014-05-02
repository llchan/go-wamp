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
				i := Integer(0)
				v = &i
				Expect(v.Value()).To(Equal(Integer(0)))
				Expect(v.Validate()).ToNot(HaveOccurred())
			})
		})

		Context("Id", func() {
			It("Should work for 0", func() {
				i := Id(0)
				v = &i
				Expect(v.Validate()).ToNot(HaveOccurred())
			})
			It("Should not work for MaxId + 1", func() {
				i := Id(MaxId + 1)
				v = &i
				Expect(v.Validate()).To(HaveOccurred())
				Expect(MaxId).To(BeEquivalentTo(1 << 53))
			})
		})

		Context("Uri", func() {
			It("Should work for abc", func() {
				u := Uri("abc")
				v = &u
				Expect(v.Validate()).ToNot(HaveOccurred())
			})
			It("Should work for abc.def.ghi", func() {
				u := Uri("abc.def.ghi")
				v = &u
				Expect(v.Validate()).ToNot(HaveOccurred())
			})
			It("Should not work for .abc.def", func() {
				u := Uri(".abc.def")
				v = &u
				Expect(v.Validate()).To(HaveOccurred())
			})
			It("Should not work for abc.#def", func() {
				u := Uri("abc.#def")
				v = &u
				Expect(v.Validate()).To(HaveOccurred())
			})
			It("Should not work for a bc.def", func() {
				u := Uri("a bc.def")
				v = &u
				Expect(v.Validate()).To(HaveOccurred())
			})
			It("Should not work for abc..def", func() {
				u := Uri("abc..def")
				v = &u
				Expect(v.Validate()).To(HaveOccurred())
			})
		})
	})
})
