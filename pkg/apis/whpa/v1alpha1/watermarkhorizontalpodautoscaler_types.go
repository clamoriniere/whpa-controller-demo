package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// WatermarkHorizontalPodAutoscalerSpec defines the desired state of WatermarkHorizontalPodAutoscaler
// +k8s:openapi-gen=true
type WatermarkHorizontalPodAutoscalerSpec struct {
	Replicas int32 `json:"replicas,omitempty"`
}

// WatermarkHorizontalPodAutoscalerStatus defines the observed state of WatermarkHorizontalPodAutoscaler
// +k8s:openapi-gen=true
type WatermarkHorizontalPodAutoscalerStatus struct {
	PodReady int32 `json:"podReady,omitempty"`
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WatermarkHorizontalPodAutoscaler is the Schema for the watermarkhorizontalpodautoscalers API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:subresource:scale:specpath=.spec.replica,statuspath=.status.podReady
// +kubebuilder:printcolumn:name="age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="replicas",type="integer",JSONPath=".spec.replicas"
// +kubebuilder:printcolumn:name="ready",type="integer",JSONPath=".status.podReady"
type WatermarkHorizontalPodAutoscaler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WatermarkHorizontalPodAutoscalerSpec   `json:"spec,omitempty"`
	Status WatermarkHorizontalPodAutoscalerStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WatermarkHorizontalPodAutoscalerList contains a list of WatermarkHorizontalPodAutoscaler
type WatermarkHorizontalPodAutoscalerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WatermarkHorizontalPodAutoscaler `json:"items"`
}

func init() {
	SchemeBuilder.Register(&WatermarkHorizontalPodAutoscaler{}, &WatermarkHorizontalPodAutoscalerList{})
}
