package testsuites

import (
	"github.com/onsi/ginkgo"
	v1 "k8s.io/api/core/v1"
	clientset "k8s.io/client-go/kubernetes"
	"sigs.k8s.io/azurefile-csi-driver/test/e2e/driver"
)

type PreProvisionedVolumeSubpathTest struct {
	CSIDriver driver.PreProvisionedVolumeTestDriver
	Pod       PodDetails
}

func (t *PreProvisionedVolumeSubpathTest) Run(client clientset.Interface, namespace *v1.Namespace) {

	tpod, cleanup := t.Pod.SetupWithPreProvisionedVolumeWithSubpath(client, namespace, t.CSIDriver)
	// defer must be called here for resources not get removed before using them
	for i := range cleanup {
		defer cleanup[i]()
	}

	ginkgo.By("deploying the pod")
	tpod.Create()
	defer tpod.Cleanup()
	ginkgo.By("checking that the pods command exits with no error")
	tpod.WaitForSuccess()
}
