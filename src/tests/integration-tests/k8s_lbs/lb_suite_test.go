package k8s_lbs_test

import (
	"testing"
	"tests/config"
	"tests/test_helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestK8sLb(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "K8sLb Suite")
}

var (
	runner      *test_helpers.KubectlRunner
	nginxLBSpec = test_helpers.PathFromRoot("specs/nginx-lb.yml")
	testconfig  *config.Config
)

var _ = BeforeSuite(func() {
	var err error
	testconfig, err = config.InitConfig()
	Expect(err).NotTo(HaveOccurred())

	runner = test_helpers.NewKubectlRunner(testconfig.Kubernetes.PathToKubeConfig)
	runner.RunKubectlCommand("create", "namespace", runner.Namespace()).Wait("60s")
})

var _ = AfterSuite(func() {
	if runner != nil {
		runner.RunKubectlCommand("delete", "namespace", runner.Namespace()).Wait("60s")
	}
})

func K8SLBDescribe(description string, callback func()) bool {
	return Describe("[k8s_lb]", func() {
		BeforeEach(func() {
			if !testconfig.IntegrationTests.IncludeK8SLB {
				Skip(`Skipping this test suite because Config.IntegrationTests.IncludeK8SLB is set to 'false'.`)
			}
		})
		Describe(description, callback)
	})
}
