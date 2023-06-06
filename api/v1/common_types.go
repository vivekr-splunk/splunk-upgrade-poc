package v1

import (
	corev1 "k8s.io/api/core/v1"
)

// Phase is used to represent the current phase of a custom resource
// +kubebuilder:validation:Enum=Pending;Ready;Updating;ScalingUp;ScalingDown;Terminating;Error
type Phase string

const (
	// PhasePending means a custom resource has just been created and is not yet ready
	PhasePending Phase = "Pending"

	// PhaseReady means a custom resource is ready and up to date
	PhaseReady Phase = "Ready"

	// PhaseUpdating means a custom resource is in the process of updating to a new desired state (spec)
	PhaseUpdating Phase = "Updating"

	// PhaseScalingUp means a custom resource is in the process of scaling up
	PhaseScalingUp Phase = "ScalingUp"

	// PhaseScalingDown means a custom resource is in the process of scaling down
	PhaseScalingDown Phase = "ScalingDown"

	// PhaseTerminating means a custom resource is in the process of being removed
	PhaseTerminating Phase = "Terminating"

	// PhaseError means an error occured with custom resource management
	PhaseError Phase = "Error"
)

// Spec defines a subset of the desired state of parameters that are common across all CRD types
type Spec struct {
	// Image to use for Splunk pod containers (overrides RELATED_IMAGE_SPLUNK_ENTERPRISE environment variables)
	Image string `json:"image"`

	// Sets pull policy for all images (either “Always” or the default: “IfNotPresent”)
	// +kubebuilder:validation:Enum=Always;IfNotPresent
	ImagePullPolicy string `json:"imagePullPolicy"`

	// Name of Scheduler to use for pod placement (defaults to “default-scheduler”)
	SchedulerName string `json:"schedulerName"`

	// Kubernetes Affinity rules that control how pods are assigned to particular nodes.
	Affinity corev1.Affinity `json:"affinity"`

	// Pod's tolerations for Kubernetes node's taint
	Tolerations []corev1.Toleration `json:"tolerations,omitempty"`

	// resource requirements for the pod containers
	Resources corev1.ResourceRequirements `json:"resources"`

	// ServiceTemplate is a template used to create Kubernetes services
	ServiceTemplate corev1.Service `json:"serviceTemplate"`

	// TopologySpreadConstraint https://kubernetes.io/docs/concepts/scheduling-eviction/topology-spread-constraints/
	TopologySpreadConstraints []corev1.TopologySpreadConstraint `json:"topologySpreadConstraints,omitempty"`
}
