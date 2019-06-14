package watermarkhorizontalpodautoscaler

import (
	"reflect"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	whpav1alpha1 "github.com/datadog/whpa/pkg/apis/whpa/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

func TestReconcileWatermarkHorizontalPodAutoscaler_Reconcile(t *testing.T) {
	// Register operator types with the runtime scheme.
	s := scheme.Scheme
	s.AddKnownTypes(whpav1alpha1.SchemeGroupVersion, &whpav1alpha1.WatermarkHorizontalPodAutoscaler{})

	nwhpa1 := &whpav1alpha1.WatermarkHorizontalPodAutoscaler{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "foo",
			Namespace: "bar",
		},
		Spec: whpav1alpha1.WatermarkHorizontalPodAutoscalerSpec{
			Replicas: 3,
		},
	}

	type fields struct {
		client client.Client
		scheme *runtime.Scheme
	}
	type args struct {
		request reconcile.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    reconcile.Result
		wantErr bool
	}{
		{
			name: "new object",
			fields: fields{
				scheme: s,
				client: fake.NewFakeClient([]runtime.Object{nwhpa1}...),
			},
			args: args{
				request: reconcile.Request{
					NamespacedName: types.NamespacedName{Name: "foo", Namespace: "bar"},
				},
			},
			want:    reconcile.Result{},
			wantErr: false,
		},
		{
			name: "new object not found",
			fields: fields{
				scheme: s,
				client: fake.NewFakeClient([]runtime.Object{nwhpa1}...),
			},
			args: args{
				request: reconcile.Request{
					NamespacedName: types.NamespacedName{Name: "foo2", Namespace: "bar"},
				},
			},
			want:    reconcile.Result{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ReconcileWatermarkHorizontalPodAutoscaler{
				client: tt.fields.client,
				scheme: tt.fields.scheme,
			}
			got, err := r.Reconcile(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReconcileWatermarkHorizontalPodAutoscaler.Reconcile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReconcileWatermarkHorizontalPodAutoscaler.Reconcile() = %v, want %v", got, tt.want)
			}
		})
	}
}
