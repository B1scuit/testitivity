package k8s

import (
	"github.com/B1scuit/testitivity/domain"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const groupName = "testitivity.example.com"
const groupVersion = "v1"

var SchemeGroupVersion = schema.GroupVersion{Group: groupName, Version: groupVersion}

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&domain.TestPlan{},
		&domain.TestPlanList{},
	)

	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
