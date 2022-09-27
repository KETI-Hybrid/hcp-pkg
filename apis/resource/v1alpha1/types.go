package v1alpha1

import (
	appsv1 "k8s.io/api/apps/v1"
	hpav2beta1 "k8s.io/api/autoscaling/v2beta1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	vpav1beta2 "k8s.io/autoscaler/vertical-pod-autoscaler/pkg/apis/autoscaling.k8s.io/v1beta2"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type HCPPod struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HCPPodSpec       `json:"spec,omitempty"`
	Status            corev1.PodStatus `json:"status,omitempty"`
}

type HCPPodSpec struct {
	RealPodSpec     corev1.PodSpec    `json:"realDeploymentSpec"`
	RealPodMetadata metav1.ObjectMeta `json:"metadata,omitempty"`

	//SchedlingStatus -Requested - Scheduled - Completed
	SchedulingStatus string              `json:"schedulingstatus,omitempty" protobuf:"bytes,11,opt,name=schedulingstatus"`
	SchedulingType   string              `json:"schedulingType,omitempty" protobuf:"bytes,3,opt,name=schedulingtype"`
	SchedulingResult HCPSchedulingResult `json:"schedulingresult,omitempty" protobuf:"bytes,11,opt,name=schedulingresult"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type HCPPodList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HCPPod `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type HCPDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HCPDeploymentSpec       `json:"spec,omitempty"`
	Status appsv1.DeploymentStatus `json:"status,omitempty"`
}

type HCPDeploymentSpec struct {
	UUID                   int                   `json:"uuid"`
	RealDeploymentSpec     appsv1.DeploymentSpec `json:"realDeploymentSpec"`
	RealDeploymentMetadata metav1.ObjectMeta     `json:"metadata,omitempty"`

	//SchedlingStatus -Requested - Scheduled - Completed
	SchedulingNeed     bool `json:"schedulingneed,omitempty" protobuf:"bytes,11,opt,name=schedulingneed"`
	SchedulingComplete bool `json:"schedulingcomplete,omitempty" protobuf:"bytes,11,opt,name=schedulingcomplete"`
	//SchedulingType     string              `json:"schedulingType,omitempty" protobuf:"bytes,3,opt,name=schedulingtype"`
	SchedulingResult HCPSchedulingResult `json:"schedulingresult,omitempty" protobuf:"bytes,11,opt,name=schedulingresult"`
}

type HCPSchedulingResult struct {
	Targets []Target `json:"targets,omitempty" protobuf:"bytes,11,opt,name=targets"`
}

type Target struct {
	Cluster  string `json:"cluster,omitempty" protobuf:"bytes,11,opt,name=cluster"`
	Node     string `json:"node,omitempty" protobuf:"bytes,11,opt,name=node"`
	Replicas *int32 `json:"replicas,omitempty" protobuf:"bytes,11,opt,name=replicas"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type HCPDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HCPDeployment `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type HCPHybridAutoScaler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HCPHybridAutoScalerSpec   `json:"spec,omitempty"`
	Status HCPHybridAutoScalerStatus `json:"status,omitempty"`
}

type HCPHybridAutoScalerSpec struct {
	Mode           string         `json:"mode"`
	TargetCluster  string         `json:"targetCluster"`
	WarningCount   int32          `json:"warningCount"`
	ScalingOptions ScalingOptions `json:"scalingOptions,omitempty" protobuf:"bytes,2,opt,name=scalingoptions"`
}

type ScalingOptions struct {
	//CpaTemplate CpaTemplate                        `json:"cpaTemplate,omitempty" protobuf:"bytes,1,opt,name=cpatemplate"`
	HpaTemplate hpav2beta1.HorizontalPodAutoscaler `json:"hpaTemplate,omitempty" protobuf:"bytes,2,opt,name=hpatemplate"`
	VpaTemplate vpav1beta2.VerticalPodAutoscaler   `json:"vpaTemplate,omitempty" protobuf:"bytes,3,opt,name=hpatemplate"`
}

type HCPHybridAutoScalerStatus struct {
	ResourceStatus     string                  `json:"resourceStatus"`
	FirstProcess       bool                    `json:"firstsProcess"`
	ScalingInProcess   bool                    `json:"scalingInProcess"`
	ExpandingInProcess bool                    `json:"expandingInProcess"`
	LastSpec           HCPHybridAutoScalerSpec `json:"lastSpec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type HCPHybridAutoScalerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HCPHybridAutoScaler `json:"items"`
}
