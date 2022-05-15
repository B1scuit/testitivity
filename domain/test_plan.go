package domain

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

//go:generate controller-gen object paths=$GOFILE

type TestPlanSpec struct {
	Replicas    int    `json:"replicas"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type TestPlan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec TestPlanSpec `json:"spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type TestPlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []TestPlan `json:"items"`
}
