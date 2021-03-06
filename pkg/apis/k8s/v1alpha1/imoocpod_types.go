package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ImoocPodSpec defines the desired state of ImoocPod
type ImoocPodSpec struct {
	Replicas int `json:"replicas"`
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// ImoocPodStatus defines the observed state of ImoocPod
type ImoocPodStatus struct {
	Replicas int `json:"replicas"`
	PodNames []string `json:"podNames"`
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ImoocPod is the Schema for the imoocpods API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=imoocpods,scope=Namespaced
type ImoocPod struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ImoocPodSpec   `json:"spec,omitempty"`
	Status ImoocPodStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ImoocPodList contains a list of ImoocPod
type ImoocPodList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ImoocPod `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ImoocPod{}, &ImoocPodList{})
}
