package requirements_test

import (
	"github.com/blang/semver"
	"github.com/cloudfoundry/cli/cf/configuration/coreconfig"
	"github.com/cloudfoundry/cli/cf/requirements"

	testconfig "github.com/cloudfoundry/cli/testhelpers/configuration"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MaxAPIVersionRequirement", func() {
	var (
		config      coreconfig.Repository
		requirement requirements.MaxAPIVersionRequirement
	)

	BeforeEach(func() {
		config = testconfig.NewRepository()
		maximumVersion, err := semver.Make("1.2.3")
		Expect(err).NotTo(HaveOccurred())

		requirement = requirements.NewMaxAPIVersionRequirement(config, "version-restricted-feature", maximumVersion)
	})

	Context("Execute", func() {
		Context("when the config's api version is less than the maximum version", func() {
			BeforeEach(func() {
				config.SetApiVersion("1.2.2")
			})

			It("succeeds", func() {
				err := requirement.Execute()
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when the config's api version is equal to the maximum version", func() {
			BeforeEach(func() {
				config.SetApiVersion("1.2.3")
			})

			It("succeeds", func() {
				err := requirement.Execute()
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when the config's api version is greater than the maximum version", func() {
			BeforeEach(func() {
				config.SetApiVersion("1.2.4")
			})

			It("returns an error", func() {
				err := requirement.Execute()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("version-restricted-feature only works up to CF API version 1.2.3. Your target is 1.2.4."))
			})
		})

		Context("when the config's api version can not be parsed", func() {
			BeforeEach(func() {
				config.SetApiVersion("-")
			})

			It("returns an error", func() {
				err := requirement.Execute()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Unable to parse CC API Version '-'"))
			})
		})

		Context("when the config's api version is empty", func() {
			BeforeEach(func() {
				config.SetApiVersion("")
			})

			It("returns an error", func() {
				err := requirement.Execute()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Unable to determine CC API Version. Please log in again."))
			})
		})
	})
})