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
	// +optional
	ImagePullPolicy string `json:"imagePullPolicy,omitempty"`

	// Name of Scheduler to use for pod placement (defaults to “default-scheduler”)
	// +optional
	SchedulerName string `json:"schedulerName,omitempty"`

	// Kubernetes Affinity rules that control how pods are assigned to particular nodes.
	// +optional
	Affinity corev1.Affinity `json:"affinity,omitempty"`

	// Pod's tolerations for Kubernetes node's taint
	// +optional
	Tolerations []corev1.Toleration `json:"tolerations,omitempty"`

	// resource requirements for the pod containers
	// +optional
	Resources corev1.ResourceRequirements `json:"resources,omitempty"`

	// ServiceTemplate is a template used to create Kubernetes services
	// +optional
	ServiceTemplate corev1.Service `json:"serviceTemplate,omitempty"`

	// TopologySpreadConstraint https://kubernetes.io/docs/concepts/scheduling-eviction/topology-spread-constraints/
	// +optional
	TopologySpreadConstraints []corev1.TopologySpreadConstraint `json:"topologySpreadConstraints,omitempty"`
}
