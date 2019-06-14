package e2e

import (
	"context"
	goctx "context"
	"testing"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	apis "github.com/datadog/whpa/pkg/apis"
	whpav1alpha1 "github.com/datadog/whpa/pkg/apis/whpa/v1alpha1"

	framework "github.com/operator-framework/operator-sdk/pkg/test"
	"github.com/operator-framework/operator-sdk/pkg/test/e2eutil"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/wait"

	dynclient "sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	retryInterval        = time.Second * 5
	timeout              = time.Second * 60
	cleanupRetryInterval = time.Second * 1
	cleanupTimeout       = time.Second * 5
)

func TestController(t *testing.T) {
	whpaList := &whpav1alpha1.WatermarkHorizontalPodAutoscalerList{}
	err := framework.AddToFrameworkScheme(apis.AddToScheme, whpaList)
	if err != nil {
		t.Fatalf("failed to add custom resource scheme to framework: %v", err)
	}
	// run subtests
	t.Run("group", func(t *testing.T) {
		t.Run("Deployment", InitialDeployment)
	})
}

func InitialDeployment(t *testing.T) {
	t.Parallel()
	ns, ctx, f := initTestFwkResources(t, "whpa")
	defer ctx.Cleanup()

	nwhpa1 := &whpav1alpha1.WatermarkHorizontalPodAutoscaler{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "foo",
			Namespace: ns,
		},
		Spec: whpav1alpha1.WatermarkHorizontalPodAutoscalerSpec{
			Replicas: 3,
		},
	}
	f.Client.Create(goctx.TODO(), nwhpa1, &framework.CleanupOptions{TestContext: ctx, Timeout: cleanupTimeout, RetryInterval: cleanupRetryInterval})

	doneFunc := func(whpa *whpav1alpha1.WatermarkHorizontalPodAutoscaler) (bool, error) {
		if whpa.Status.PodReady == 3 {
			return true, nil
		}
		return false, nil
	}

	err := WaitForFuncOnHPA(t, f.Client, ns, "foo", doneFunc, retryInterval, timeout)
	if err != nil {
		t.Fatal(err)
	}
}

func initTestFwkResources(t *testing.T, deploymentName string) (string, *framework.TestCtx, *framework.Framework) {
	ctx := framework.NewTestCtx(t)
	err := ctx.InitializeClusterResources(&framework.CleanupOptions{TestContext: ctx, Timeout: cleanupTimeout, RetryInterval: cleanupRetryInterval})
	if err != nil {
		t.Fatalf("failed to initialize cluster resources: %v", err)
	}
	t.Log("Initialized cluster resources")
	namespace, err := ctx.GetNamespace()
	if err != nil {
		t.Fatal(err)
	}
	// get global framework variables
	f := framework.Global
	// wait for ddaemonset-controller to be ready
	err = e2eutil.WaitForDeployment(t, f.KubeClient, namespace, deploymentName, 1, retryInterval, timeout)
	if err != nil {
		t.Fatal(err)
	}
	return namespace, ctx, f
}

// WaitForFuncOnDDaemonset used to wait a valid condition on a DDaemonSet
func WaitForFuncOnHPA(t *testing.T, client framework.FrameworkClient, namespace, name string, f func(whpa *whpav1alpha1.WatermarkHorizontalPodAutoscaler) (bool, error), retryInterval, timeout time.Duration) error {
	return wait.Poll(retryInterval, timeout, func() (bool, error) {
		objKey := dynclient.ObjectKey{
			Namespace: namespace,
			Name:      name,
		}
		h := &whpav1alpha1.WatermarkHorizontalPodAutoscaler{}
		err := client.Get(context.TODO(), objKey, h)
		if err != nil {
			if apierrors.IsNotFound(err) {
				t.Logf("Waiting for availability of %s WatermarkHorizontalPodAutoscaler\n", name)
				return false, nil
			}
			return false, err
		}

		ok, err := f(h)
		t.Logf("Waiting for condition function to be true ok for %s WatermarkHorizontalPodAutoscaler (%t/%v)\n", name, ok, err)
		return ok, err
	})
}
