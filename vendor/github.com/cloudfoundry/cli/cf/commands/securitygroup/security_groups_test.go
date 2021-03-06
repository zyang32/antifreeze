package securitygroup_test

import (
	"github.com/cloudfoundry/cli/cf/api/securitygroups/securitygroupsfakes"
	"github.com/cloudfoundry/cli/cf/commandregistry"
	"github.com/cloudfoundry/cli/cf/configuration/coreconfig"
	"github.com/cloudfoundry/cli/cf/errors"
	"github.com/cloudfoundry/cli/cf/models"
	"github.com/cloudfoundry/cli/flags"
	testcmd "github.com/cloudfoundry/cli/testhelpers/commands"
	testconfig "github.com/cloudfoundry/cli/testhelpers/configuration"
	testreq "github.com/cloudfoundry/cli/testhelpers/requirements"
	testterm "github.com/cloudfoundry/cli/testhelpers/terminal"

	"github.com/cloudfoundry/cli/cf/commands/securitygroup"
	. "github.com/cloudfoundry/cli/testhelpers/matchers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("list-security-groups command", func() {
	var (
		ui                  *testterm.FakeUI
		repo                *securitygroupsfakes.FakeSecurityGroupRepo
		requirementsFactory *testreq.FakeReqFactory
		configRepo          coreconfig.Repository
		deps                commandregistry.Dependency
	)

	updateCommandDependency := func(pluginCall bool) {
		deps.Ui = ui
		deps.RepoLocator = deps.RepoLocator.SetSecurityGroupRepository(repo)
		deps.Config = configRepo
		commandregistry.Commands.SetCommand(commandregistry.Commands.FindCommand("security-groups").SetDependency(deps, pluginCall))
	}

	BeforeEach(func() {
		ui = &testterm.FakeUI{}
		requirementsFactory = &testreq.FakeReqFactory{}
		repo = new(securitygroupsfakes.FakeSecurityGroupRepo)
		configRepo = testconfig.NewRepositoryWithDefaults()
	})

	runCommand := func(args ...string) bool {
		return testcmd.RunCliCommand("security-groups", args, requirementsFactory, updateCommandDependency, false)
	}

	Describe("requirements", func() {
		It("should fail if not logged in", func() {
			Expect(runCommand()).To(BeFalse())
		})

		Context("when arguments are provided", func() {
			var cmd commandregistry.Command
			var flagContext flags.FlagContext

			BeforeEach(func() {
				cmd = &securitygroup.SecurityGroups{}
				cmd.SetDependency(deps, false)
				flagContext = flags.NewFlagContext(cmd.MetaData().Flags)
			})

			It("should fail with usage", func() {
				flagContext.Parse("blahblah")

				reqs := cmd.Requirements(requirementsFactory, flagContext)

				err := testcmd.RunRequirements(reqs)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Incorrect Usage"))
				Expect(err.Error()).To(ContainSubstring("No argument required"))
			})
		})
	})

	Context("when logged in", func() {
		BeforeEach(func() {
			requirementsFactory.LoginSuccess = true
		})

		It("tells the user what it's about to do", func() {
			runCommand()
			Expect(ui.Outputs).To(ContainSubstrings(
				[]string{"Getting", "security groups", "my-user"},
			))
		})

		It("handles api errors with an error message", func() {
			repo.FindAllReturns([]models.SecurityGroup{}, errors.New("YO YO YO, ERROR YO"))

			runCommand()
			Expect(ui.Outputs).To(ContainSubstrings(
				[]string{"FAILED"},
			))
		})

		Context("when there are no security groups", func() {
			It("Should tell the user that there are no security groups", func() {
				repo.FindAllReturns([]models.SecurityGroup{}, nil)

				runCommand()
				Expect(ui.Outputs).To(ContainSubstrings([]string{"No security groups"}))
			})
		})

		Context("when there is at least one security group", func() {
			BeforeEach(func() {
				securityGroup := models.SecurityGroup{}
				securityGroup.Name = "my-group"
				securityGroup.Guid = "group-guid"

				repo.FindAllReturns([]models.SecurityGroup{securityGroup}, nil)
			})

			Describe("Where there are spaces assigned", func() {
				BeforeEach(func() {
					securityGroups := []models.SecurityGroup{
						{
							SecurityGroupFields: models.SecurityGroupFields{
								Name: "my-group",
								Guid: "group-guid",
							},
							Spaces: []models.Space{
								{
									SpaceFields:  models.SpaceFields{Guid: "my-space-guid-1", Name: "space-1"},
									Organization: models.OrganizationFields{Guid: "my-org-guid-1", Name: "org-1"},
								},
								{
									SpaceFields:  models.SpaceFields{Guid: "my-space-guid", Name: "space-2"},
									Organization: models.OrganizationFields{Guid: "my-org-guid-2", Name: "org-2"},
								},
							},
						},
					}

					repo.FindAllReturns(securityGroups, nil)
				})

				It("lists out the security group's: name, organization and space", func() {
					runCommand()
					Expect(ui.Outputs).To(ContainSubstrings(
						[]string{"Getting", "security group", "my-user"},
						[]string{"OK"},
						[]string{"#0", "my-group", "org-1", "space-1"},
					))

					//If there is a panic in this test, it is likely due to the following
					//Expectation to be false
					Expect(ui.Outputs).ToNot(ContainSubstrings(
						[]string{"#0", "my-group", "org-2", "space-2"},
					))
				})
			})

			Describe("Where there are no spaces assigned", func() {
				It("lists out the security group's: name", func() {
					runCommand()
					Expect(ui.Outputs).To(ContainSubstrings(
						[]string{"Getting", "security group", "my-user"},
						[]string{"OK"},
						[]string{"#0", "my-group"},
					))
				})
			})
		})
	})
})
