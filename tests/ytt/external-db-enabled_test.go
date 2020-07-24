package ytt

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("External DB Enabled", func() {

	var templates []string

	BeforeEach(func() {
		templates = []string{
			pathToFile("config/postgres"),
			pathToFile("tests/ytt/postgres/postgres-values.yml"),
			pathToFile("tests/ytt/postgres/cf-values.yml"),
		}
	})

	It("should do something", func() {
		ctx := NewRenderingContext(templates...).WithData(
			map[string]string{
			})

		Expect(ctx).To(ProduceYAML(WithDocument("cf-db",
			RepresentingNamespace().WithName("cf-db"),
		)))
	})
})
