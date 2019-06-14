package watermarkhorizontalpodautoscaler

import (
	"context"

	whpav1alpha1 "github.com/datadog/whpa/pkg/apis/whpa/v1alpha1"

	corev1 "k8s.io/api/core/v1"
	apiequality "k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_watermarkhorizontalpodautoscaler")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new WatermarkHorizontalPodAutoscaler Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileWatermarkHorizontalPodAutoscaler{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("watermarkhorizontalpodautoscaler-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource WatermarkHorizontalPodAutoscaler
	err = c.Watch(&source.Kind{Type: &whpav1alpha1.WatermarkHorizontalPodAutoscaler{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner WatermarkHorizontalPodAutoscaler
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &whpav1alpha1.WatermarkHorizontalPodAutoscaler{},
	})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileWatermarkHorizontalPodAutoscaler implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileWatermarkHorizontalPodAutoscaler{}

// ReconcileWatermarkHorizontalPodAutoscaler reconciles a WatermarkHorizontalPodAutoscaler object
type ReconcileWatermarkHorizontalPodAutoscaler struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a WatermarkHorizontalPodAutoscaler object and makes changes based on the state read
// and what is in the WatermarkHorizontalPodAutoscaler.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileWatermarkHorizontalPodAutoscaler) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling WatermarkHorizontalPodAutoscaler")

	// Fetch the WatermarkHorizontalPodAutoscaler instance
	instance := &whpav1alpha1.WatermarkHorizontalPodAutoscaler{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	replicas := instance.Spec.Replicas

	// Check if this Pod already exists
	foundPods := &corev1.PodList{}
	options := &client.ListOptions{
		LabelSelector: labels.Set{"app": instance.Name}.AsSelectorPreValidated(),
	}
	if err = r.client.List(context.TODO(), options, foundPods); err != nil {
		return reconcile.Result{}, err
	}

	if len(foundPods.Items) < int(replicas) {
		// Define a new Pod object
		pod := newPodForCR(instance)

		// Set WatermarkHorizontalPodAutoscaler instance as the owner and controller
		if err := controllerutil.SetControllerReference(instance, pod, r.scheme); err != nil {
			return reconcile.Result{}, err
		}
		reqLogger.Info("Creating a new Pod", "Pod.Namespace", pod.Namespace, "Pod.Name", pod.Name)
		err = r.client.Create(context.TODO(), pod)
		if err != nil {
			return reconcile.Result{}, err
		}

		// Pod created successfully - don't requeue
		return reconcile.Result{}, nil
	} else if len(foundPods.Items) > int(replicas) {
		pod := foundPods.Items[0]
		err = r.client.Delete(context.TODO(), &pod)
		if err != nil {
			return reconcile.Result{}, err
		}
	}

	newInstance := instance.DeepCopy()
	newInstance.Status.PodReady = int32(len(foundPods.Items))
	if !apiequality.Semantic.DeepEqual(instance.Status, newInstance.Status) {
		reqLogger.Info("Updating the status of WatermarkHorizontalPodAutoscaler", instance.Namespace, "/", instance.Name)
		if err = r.client.Status().Update(context.TODO(), newInstance); err != nil {
			return reconcile.Result{}, err
		}
	}

	return reconcile.Result{}, nil
}

// newPodForCR returns a busybox pod with the same name/namespace as the cr
func newPodForCR(cr *whpav1alpha1.WatermarkHorizontalPodAutoscaler) *corev1.Pod {
	labels := map[string]string{
		"app": cr.Name,
	}
	return &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: cr.Name + "-pod-",
			Namespace:    cr.Namespace,
			Labels:       labels,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:    "busybox",
					Image:   "busybox",
					Command: []string{"sleep", "3600"},
				},
			},
		},
	}
}
