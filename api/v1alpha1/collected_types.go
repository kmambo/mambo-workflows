package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// NodeSelectorApplyConfiguration represents an declarative configuration of the NodeSelector type for use
// with apply.
// +kubebuilder:object:root=true
type NodeSelectorApplyConfiguration struct {
	NodeSelectorTerms []NodeSelectorTermApplyConfiguration `json:"nodeSelectorTerms,omitempty"`
}

// ObjectFieldSelectorApplyConfiguration represents a declarative configuration of the ObjectFieldSelector type for use
// with apply.
// +kubebuilder:object:root=true
type ObjectFieldSelectorApplyConfiguration struct {
	APIVersion *string `json:"apiVersion,omitempty"`
	FieldPath  *string `json:"fieldPath,omitempty"`
}

// LocalVolumeSourceApplyConfiguration represents a declarative configuration of the LocalVolumeSource type for use
// with apply.
// +kubebuilder:object:root=true
type LocalVolumeSourceApplyConfiguration struct {
	Path   *string `json:"path,omitempty"`
	FSType *string `json:"fsType,omitempty"`
}

// PersistentVolumeClaimTemplateApplyConfiguration represents a declarative configuration of the PersistentVolumeClaimTemplate type for use
// with apply.
// +kubebuilder:object:root=true
type PersistentVolumeClaimTemplateApplyConfiguration struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              *PersistentVolumeClaimSpecApplyConfiguration `json:"spec,omitempty"`
}

// EphemeralVolumeSourceApplyConfiguration represents a declarative configuration of the EphemeralVolumeSource type for use
// with apply.
// +kubebuilder:object:root=true
type EphemeralVolumeSourceApplyConfiguration struct {
	VolumeClaimTemplate *PersistentVolumeClaimTemplateApplyConfiguration `json:"volumeClaimTemplate,omitempty"`
}

// ResourceFieldSelectorApplyConfiguration represents a declarative configuration of the ResourceFieldSelector type for use
// with apply.
// +kubebuilder:object:root=true
type ResourceFieldSelectorApplyConfiguration struct {
	ContainerName *string            `json:"containerName,omitempty"`
	Resource      *string            `json:"resource,omitempty"`
	Divisor       *resource.Quantity `json:"divisor,omitempty"`
}

// ContainerResizePolicyApplyConfiguration represents a declarative configuration of the ContainerResizePolicy type for use
// with apply.
// +kubebuilder:object:root=true
type ContainerResizePolicyApplyConfiguration struct {
	ResourceName  *string `json:"resourceName,omitempty"`
	RestartPolicy string  `json:"restartPolicy,omitempty"`
}

// VolumeSourceApplyConfiguration represents a declarative configuration of the VolumeSource type for use
// with apply.
// +kubebuilder:object:root=true
type VolumeSourceApplyConfiguration struct {
	EmptyDir              *EmptyDirVolumeSourceApplyConfiguration              `json:"emptyDir,omitempty"`
	Secret                *SecretVolumeSourceApplyConfiguration                `json:"secret,omitempty"`
	NFS                   *NFSVolumeSourceApplyConfiguration                   `json:"nfs,omitempty"`
	ISCSI                 *ISCSIVolumeSourceApplyConfiguration                 `json:"iscsi,omitempty"`
	PersistentVolumeClaim *PersistentVolumeClaimVolumeSourceApplyConfiguration `json:"persistentVolumeClaim,omitempty"`
	DownwardAPI           *DownwardAPIVolumeSourceApplyConfiguration           `json:"downwardAPI,omitempty"`
	ConfigMap             *ConfigMapVolumeSourceApplyConfiguration             `json:"configMap,omitempty"`
	Projected             *ProjectedVolumeSourceApplyConfiguration             `json:"projected,omitempty"`
	StorageOS             *StorageOSVolumeSourceApplyConfiguration             `json:"storageos,omitempty"`
	CSI                   *CSIVolumeSourceApplyConfiguration                   `json:"csi,omitempty"`
	Ephemeral             *EphemeralVolumeSourceApplyConfiguration             `json:"ephemeral,omitempty"`
}

// HTTPGetActionApplyConfiguration represents a declarative configuration of the HTTPGetAction type for use
// with apply.
// +kubebuilder:object:root=true
type HTTPGetActionApplyConfiguration struct {
	Path        *string                        `json:"path,omitempty"`
	Port        *intstr.IntOrString            `json:"port,omitempty"`
	Host        *string                        `json:"host,omitempty"`
	Scheme      *string                        `json:"scheme,omitempty"`
	HTTPHeaders []HTTPHeaderApplyConfiguration `json:"httpHeaders,omitempty"`
}

// FlockerVolumeSourceApplyConfiguration represents a declarative configuration of the FlockerVolumeSource type for use
// with apply.
// +kubebuilder:object:root=true
type FlockerVolumeSourceApplyConfiguration struct {
	DatasetName *string `json:"datasetName,omitempty"`
	DatasetUUID *string `json:"datasetUUID,omitempty"`
}

// NFSVolumeSourceApplyConfiguration represents a declarative configuration of the NFSVolumeSource type for use
// with apply.
// +kubebuilder:object:root=true
type NFSVolumeSourceApplyConfiguration struct {
	Server   *string `json:"server,omitempty"`
	Path     *string `json:"path,omitempty"`
	ReadOnly *bool   `json:"readOnly,omitempty"`
}

// EventApplyConfiguration represents a declarative configuration of the Event type for use
// with apply.
// +kubebuilder:object:root=true
type EventApplyConfiguration struct {
	metav1.TypeMeta     `json:",inline"`
	metav1.ObjectMeta   `json:"metadata,omitempty"`
	InvolvedObject      *ObjectReferenceApplyConfiguration `json:"involvedObject,omitempty"`
	Reason              *string                            `json:"reason,omitempty"`
	Message             *string                            `json:"message,omitempty"`
	Source              *EventSourceApplyConfiguration     `json:"source,omitempty"`
	FirstTimestamp      *metav1.Time                       `json:"firstTimestamp,omitempty"`
	LastTimestamp       *metav1.Time                       `json:"lastTimestamp,omitempty"`
	Count               *int32                             `json:"count,omitempty"`
	Type                *string                            `json:"type,omitempty"`
	EventTime           *metav1.MicroTime                  `json:"eventTime,omitempty"`
	Series              *EventSeriesApplyConfiguration     `json:"series,omitempty"`
	Action              *string                            `json:"action,omitempty"`
	Related             *ObjectReferenceApplyConfiguration `json:"related,omitempty"`
	ReportingController *string                            `json:"reportingComponent,omitempty"`
	ReportingInstance   *string                            `json:"reportingInstance,omitempty"`
}

// LoadBalancerIngressApplyConfiguration represents a declarative configuration of the LoadBalancerIngress type for use
// with apply.
// +kubebuilder:object:root=true
type LoadBalancerIngressApplyConfiguration struct {
	IP       *string                        `json:"ip,omitempty"`
	Hostname *string                        `json:"hostname,omitempty"`
	Ports    []PortStatusApplyConfiguration `json:"ports,omitempty"`
}

// PersistentVolumeClaimSpecApplyConfiguration represents a declarative configuration of the PersistentVolumeClaimSpec type for use
// with apply.
// +kubebuilder:object:root=true
type PersistentVolumeClaimSpecApplyConfiguration struct {
	AccessModes      []string                                     `json:"accessModes,omitempty"`
	Selector         LabelSelectorApplyConfiguration              `json:"selector,omitempty"`
	Resources        *ResourceRequirementsApplyConfiguration      `json:"resources,omitempty"`
	VolumeName       *string                                      `json:"volumeName,omitempty"`
	StorageClassName *string                                      `json:"storageClassName,omitempty"`
	VolumeMode       *string                                      `json:"volumeMode,omitempty"`
	DataSource       *TypedLocalObjectReferenceApplyConfiguration `json:"dataSource,omitempty"`
	DataSourceRef    *TypedObjectReferenceApplyConfiguration      `json:"dataSourceRef,omitempty"`
}

// ProjectedVolumeSourceApplyConfiguration represents a declarative configuration of the ProjectedVolumeSource type for use
// with apply.
// +kubebuilder:object:root=true
type ProjectedVolumeSourceApplyConfiguration struct {
	Sources     []VolumeProjectionApplyConfiguration `json:"sources,omitempty"`
	DefaultMode *int32                               `json:"defaultMode,omitempty"`
}

// PodTemplateApplyConfiguration represents a declarative configuration of the PodTemplate type for use
// with apply.
// +kubebuilder:object:root=true
type PodTemplateApplyConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Template          *PodTemplateSpecApplyConfiguration `json:"template,omitempty"`
}

// ConfigMapNodeConfigSourceApplyConfiguration represents a declarative configuration of the ConfigMapNodeConfigSource type for use
// with apply.
// +kubebuilder:object:root=true
type ConfigMapNodeConfigSourceApplyConfiguration struct {
	Namespace        *string    `json:"namespace,omitempty"`
	Name             *string    `json:"name,omitempty"`
	UID              *types.UID `json:"uid,omitempty"`
	ResourceVersion  *string    `json:"resourceVersion,omitempty"`
	KubeletConfigKey *string    `json:"kubeletConfigKey,omitempty"`
}

// TaintApplyConfiguration represents a declarative configuration of the Taint type for use
// with apply.
// +kubebuilder:object:root=true
type TaintApplyConfiguration struct {
	Key       *string      `json:"key,omitempty"`
	Value     *string      `json:"value,omitempty"`
	Effect    *string      `json:"effect,omitempty"`
	TimeAdded *metav1.Time `json:"timeAdded,omitempty"`
}

// PodReadinessGateApplyConfiguration represents a declarative configuration of the PodReadinessGate type for use
// with apply.
// +kubebuilder:object:root=true
type PodReadinessGateApplyConfiguration struct {
	ConditionType *string `json:"conditionType,omitempty"`
}

// ContainerStatusApplyConfiguration represents a declarative configuration of the ContainerStatus type for use
// with apply.
// +kubebuilder:object:root=true
type ContainerStatusApplyConfiguration struct {
	Name                 *string                                 `json:"name,omitempty"`
	State                *ContainerStateApplyConfiguration       `json:"state,omitempty"`
	LastTerminationState *ContainerStateApplyConfiguration       `json:"lastState,omitempty"`
	Ready                *bool                                   `json:"ready,omitempty"`
	RestartCount         *int32                                  `json:"restartCount,omitempty"`
	Image                *string                                 `json:"image,omitempty"`
	ImageID              *string                                 `json:"imageID,omitempty"`
	ContainerID          *string                                 `json:"containerID,omitempty"`
	Started              *bool                                   `json:"started,omitempty"`
	AllocatedResources   *map[string]resource.Quantity           `json:"allocatedResources,omitempty"`
	Resources            *ResourceRequirementsApplyConfiguration `json:"resources,omitempty"`
}

// VolumeMountApplyConfiguration represents a declarative configuration of the VolumeMount type for use
// with apply.
// +kubebuilder:object:root=true
type VolumeMountApplyConfiguration struct {
	Name             *string `json:"name,omitempty"`
	ReadOnly         *bool   `json:"readOnly,omitempty"`
	MountPath        *string `json:"mountPath,omitempty"`
	SubPath          *string `json:"subPath,omitempty"`
	MountPropagation *string `json:"mountPropagation,omitempty"`
	SubPathExpr      *string `json:"subPathExpr,omitempty"`
}

// EventSourceApplyConfiguration represents a declarative configuration of the EventSource type for use
// with apply.
// +kubebuilder:object:root=true
type EventSourceApplyConfiguration struct {
	Component *string `json:"component,omitempty"`
	Host      *string `json:"host,omitempty"`
}

// ConfigMapEnvSourceApplyConfiguration represents a declarative configuration of the ConfigMapEnvSource type for use
// with apply.
// +kubebuilder:object:root=true
type ConfigMapEnvSourceApplyConfiguration struct {
	LocalObjectReferenceApplyConfiguration `json:",inline"`
	Optional                               *bool `json:"optional,omitempty"`
}

// DownwardAPIVolumeSourceApplyConfiguration represents a declarative configuration of the DownwardAPIVolumeSource type for use
// with apply.
// +kubebuilder:object:root=true
type DownwardAPIVolumeSourceApplyConfiguration struct {
	Items       []DownwardAPIVolumeFileApplyConfiguration `json:"items,omitempty"`
	DefaultMode *int32                                    `json:"defaultMode,omitempty"`
}

// QuobyteVolumeSourceApplyConfiguration represents a declarative configuration of the QuobyteVolumeSource type for use
// with apply.
// +kubebuilder:object:root=true
type QuobyteVolumeSourceApplyConfiguration struct {
	Registry *string `json:"registry,omitempty"`
	Volume   *string `json:"volume,omitempty"`
	ReadOnly *bool   `json:"readOnly,omitempty"`
	User     *string `json:"user,omitempty"`
	Group    *string `json:"group,omitempty"`
	Tenant   *string `json:"tenant,omitempty"`
}

// FCVolumeSourceApplyConfiguration represents a declarative configuration of the FCVolumeSource type for use
// with apply.
// +kubebuilder:object:root=true
type FCVolumeSourceApplyConfiguration struct {
	TargetWWNs []string `json:"targetWWNs,omitempty"`
	Lun        *int32   `json:"lun,omitempty"`
	FSType     *string  `json:"fsType,omitempty"`
	ReadOnly   *bool    `json:"readOnly,omitempty"`
	WWIDs      []string `json:"wwids,omitempty"`
}

// TypedObjectReferenceApplyConfiguration represents a declarative configuration of the TypedObjectReference type for use
// with apply.
// +kubebuilder:object:root=true
type TypedObjectReferenceApplyConfiguration struct {
	APIGroup  *string `json:"apiGroup,omitempty"`
	Kind      *string `json:"kind,omitempty"`
	Name      *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
}

// PodDNSConfigOptionApplyConfiguration represents a declarative configuration of the PodDNSConfigOption type for use
// with apply.
// +kubebuilder:object:root=true
type PodDNSConfigOptionApplyConfiguration struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

// RBDPersistentVolumeSourceApplyConfiguration represents a declarative configuration of the RBDPersistentVolumeSource type for use
// with apply.
// +kubebuilder:object:root=true
type RBDPersistentVolumeSourceApplyConfiguration struct {
	CephMonitors []string                           `json:"monitors,omitempty"`
	RBDImage     *string                            `json:"image,omitempty"`
	FSType       *string                            `json:"fsType,omitempty"`
	RBDPool      *string                            `json:"pool,omitempty"`
	RadosUser    *string                            `json:"user,omitempty"`
	Keyring      *string                            `json:"keyring,omitempty"`
	SecretRef    *SecretReferenceApplyConfiguration `json:"secretRef,omitempty"`
	ReadOnly     *bool                              `json:"readOnly,omitempty"`
}

// SecretKeySelectorApplyConfiguration represents a declarative configuration of the SecretKeySelector type for use
// with apply.
// +kubebuilder:object:root=true
type SecretKeySelectorApplyConfiguration struct {
	LocalObjectReferenceApplyConfiguration `json:",inline"`
	Key                                    *string `json:"key,omitempty"`
	Optional                               *bool   `json:"optional,omitempty"`
}

// PersistentVolumeStatusApplyConfiguration represents a declarative configuration of the PersistentVolumeStatus type for use
// with apply.
// +kubebuilder:object:root=true
type PersistentVolumeStatusApplyConfiguration struct {
	Phase   *string `json:"phase,omitempty"`
	Message *string `json:"message,omitempty"`
	Reason  *string `json:"reason,omitempty"`
}

// KeyToPathApplyConfiguration represents a declarative configuration of the KeyToPath type for use
// with apply.
// +kubebuilder:object:root=true
type KeyToPathApplyConfiguration struct {
	Key  *string `json:"key,omitempty"`
	Path *string `json:"path,omitempty"`
	Mode *int32  `json:"mode,omitempty"`
}

// VsphereVirtualDiskVolumeSourceApplyConfiguration represents a declarative configuration of the VsphereVirtualDiskVolumeSource type for use
// with apply.
// +kubebuilder:object:root=true
type VsphereVirtualDiskVolumeSourceApplyConfiguration struct {
	VolumePath        *string `json:"volumePath,omitempty"`
	FSType            *string `json:"fsType,omitempty"`
	StoragePolicyName *string `json:"storagePolicyName,omitempty"`
	StoragePolicyID   *string `json:"storagePolicyID,omitempty"`
}

// ContainerImageApplyConfiguration represents a declarative configuration of the ContainerImage type for use
// with apply.
// +kubebuilder:object:root=true
type ContainerImageApplyConfiguration struct {
	Names     []string `json:"names,omitempty"`
	SizeBytes *int64   `json:"sizeBytes,omitempty"`
}

// ResourceQuotaStatusApplyConfiguration represents a declarative configuration of the ResourceQuotaStatus type for use
// with apply.
// +kubebuilder:object:root=true
type ResourceQuotaStatusApplyConfiguration struct {
	Hard *map[string]resource.Quantity `json:"hard,omitempty"`
	Used *map[string]resource.Quantity `json:"used,omitempty"`
}

// ISCSIVolumeSourceApplyConfiguration represents a declarative configuration of the ISCSIVolumeSource type for use
// with apply.
// +kubebuilder:object:root=true
type ISCSIVolumeSourceApplyConfiguration struct {
	TargetPortal      *string                                 `json:"targetPortal,omitempty"`
	IQN               *string                                 `json:"iqn,omitempty"`
	Lun               *int32                                  `json:"lun,omitempty"`
	ISCSIInterface    *string                                 `json:"iscsiInterface,omitempty"`
	FSType            *string                                 `json:"fsType,omitempty"`
	ReadOnly          *bool                                   `json:"readOnly,omitempty"`
	Portals           []string                                `json:"portals,omitempty"`
	DiscoveryCHAPAuth *bool                                   `json:"chapAuthDiscovery,omitempty"`
	SessionCHAPAuth   *bool                                   `json:"chapAuthSession,omitempty"`
	SecretRef         *LocalObjectReferenceApplyConfiguration `json:"secretRef,omitempty"`
	InitiatorName     *string                                 `json:"initiatorName,omitempty"`
}

// PersistentVolumeClaimVolumeSourceApplyConfiguration represents a declarative configuration of the PersistentVolumeClaimVolumeSource type for use
// with apply.
// +kubebuilder:object:root=true
type PersistentVolumeClaimVolumeSourceApplyConfiguration struct {
	ClaimName *string `json:"claimName,omitempty"`
	ReadOnly  *bool   `json:"readOnly,omitempty"`
}

// GRPCActionApplyConfiguration represents a declarative configuration of the GRPCAction type for use
// with apply.
// +kubebuilder:object:root=true
type GRPCActionApplyConfiguration struct {
	Port    *int32  `json:"port,omitempty"`
	Service *string `json:"service,omitempty"`
}

// CSIPersistentVolumeSourceApplyConfiguration represents a declarative configuration of the CSIPersistentVolumeSource type for use
// with apply.
// +kubebuilder:object:root=true
type CSIPersistentVolumeSourceApplyConfiguration struct {
	Driver                     *string                            `json:"driver,omitempty"`
	VolumeHandle               *string                            `json:"volumeHandle,omitempty"`
	ReadOnly                   *bool                              `json:"readOnly,omitempty"`
	FSType                     *string                            `json:"fsType,omitempty"`
	VolumeAttributes           map[string]string                  `json:"volumeAttributes,omitempty"`
	ControllerPublishSecretRef *SecretReferenceApplyConfiguration `json:"controllerPublishSecretRef,omitempty"`
	NodeStageSecretRef         *SecretReferenceApplyConfiguration `json:"nodeStageSecretRef,omitempty"`
	NodePublishSecretRef       *SecretReferenceApplyConfiguration `json:"nodePublishSecretRef,omitempty"`
	ControllerExpandSecretRef  *SecretReferenceApplyConfiguration `json:"controllerExpandSecretRef,omitempty"`
	NodeExpandSecretRef        *SecretReferenceApplyConfiguration `json:"nodeExpandSecretRef,omitempty"`
}

// ContainerStateApplyConfiguration represents a declarative configuration of the ContainerState type for use
// with apply.
// +kubebuilder:object:root=true
type ContainerStateApplyConfiguration struct {
	Waiting    *ContainerStateWaitingApplyConfiguration    `json:"waiting,omitempty"`
	Running    *ContainerStateRunningApplyConfiguration    `json:"running,omitempty"`
	Terminated *ContainerStateTerminatedApplyConfiguration `json:"terminated,omitempty"`
}

// PodSchedulingGateApplyConfiguration represents a declarative configuration of the PodSchedulingGate type for use
// with apply.
// +kubebuilder:object:root=true
type PodSchedulingGateApplyConfiguration struct {
	Name *string `json:"name,omitempty"`
}

// PodApplyConfiguration represents a declarative configuration of the Pod type for use
// with apply.
// +kubebuilder:object:root=true
type PodApplyConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              *PodSpecApplyConfiguration   `json:"spec,omitempty"`
	Status            *PodStatusApplyConfiguration `json:"status,omitempty"`
}

// PersistentVolumeSourceApplyConfiguration represents a declarative configuration of the PersistentVolumeSource type for use
// with apply.
// +kubebuilder:object:root=true
type PersistentVolumeSourceApplyConfiguration struct {
	HostPath  *HostPathVolumeSourceApplyConfiguration            `json:"hostPath,omitempty"`
	NFS       *NFSVolumeSourceApplyConfiguration                 `json:"nfs,omitempty"`
	RBD       *RBDPersistentVolumeSourceApplyConfiguration       `json:"rbd,omitempty"`
	ISCSI     *ISCSIPersistentVolumeSourceApplyConfiguration     `json:"iscsi,omitempty"`
	Local     *LocalVolumeSourceApplyConfiguration               `json:"local,omitempty"`
	StorageOS *StorageOSPersistentVolumeSourceApplyConfiguration `json:"storageos,omitempty"`
	CSI       *CSIPersistentVolumeSourceApplyConfiguration       `json:"csi,omitempty"`
}

// ScopedResourceSelectorRequirementApplyConfiguration represents a declarative configuration of the ScopedResourceSelectorRequirement type for use
// with apply.
// +kubebuilder:object:root=true
type ScopedResourceSelectorRequirementApplyConfiguration struct {
	ScopeName *string  `json:"scopeName,omitempty"`
	Operator  *string  `json:"operator,omitempty"`
	Values    []string `json:"values,omitempty"`
}

// LocalObjectReferenceApplyConfiguration represents a declarative configuration of the LocalObjectReference type for use
// with apply.
// +kubebuilder:object:root=true
type LocalObjectReferenceApplyConfiguration struct {
	Name *string `json:"name,omitempty"`
}

// +kubebuilder:object:root=true
type AffinityApplyConfiguration struct {
	NodeAffinity    *NodeAffinityApplyConfiguration    `json:"nodeAffinity,omitempty"`
	PodAffinity     *PodAffinityApplyConfiguration     `json:"podAffinity,omitempty"`
	PodAntiAffinity *PodAntiAffinityApplyConfiguration `json:"podAntiAffinity,omitempty"`
}

// StorageOSPersistentVolumeSourceApplyConfiguration represents a declarative configuration of the StorageOSPersistentVolumeSource type for use
// with apply.
// +kubebuilder:object:root=true
type StorageOSPersistentVolumeSourceApplyConfiguration struct {
	VolumeName      *string                            `json:"volumeName,omitempty"`
	VolumeNamespace *string                            `json:"volumeNamespace,omitempty"`
	FSType          *string                            `json:"fsType,omitempty"`
	ReadOnly        *bool                              `json:"readOnly,omitempty"`
	SecretRef       *ObjectReferenceApplyConfiguration `json:"secretRef,omitempty"`
}

// NodeAffinityApplyConfiguration represents a declarative configuration of the NodeAffinity type for use
// with apply.
// +kubebuilder:object:root=true
type NodeAffinityApplyConfiguration struct {
	RequiredDuringSchedulingIgnoredDuringExecution  NodeSelectorApplyConfiguration              `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
	PreferredDuringSchedulingIgnoredDuringExecution []PreferredSchedulingTermApplyConfiguration `json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
}

// ResourceRequirementsApplyConfiguration represents a declarative configuration of the ResourceRequirements type for use
// with apply.
// +kubebuilder:object:root=true
type ResourceRequirementsApplyConfiguration struct {
	Limits   *map[string]resource.Quantity     `json:"limits,omitempty"`
	Requests *map[string]resource.Quantity     `json:"requests,omitempty"`
	Claims   []ResourceClaimApplyConfiguration `json:"claims,omitempty"`
}

// SecretEnvSourceApplyConfiguration represents a declarative configuration of the SecretEnvSource type for use
// with apply.
// +kubebuilder:object:root=true
type SecretEnvSourceApplyConfiguration struct {
	LocalObjectReferenceApplyConfiguration `json:",inline"`
	Optional                               *bool `json:"optional,omitempty"`
}

// PodOSApplyConfiguration represents a declarative configuration of the PodOS type for use
// with apply.
// +kubebuilder:object:root=true
type PodOSApplyConfiguration struct {
	Name *string `json:"name,omitempty"`
}

// SessionAffinityConfigApplyConfiguration represents a declarative configuration of the SessionAffinityConfig type for use
// with apply.
// +kubebuilder:object:root=true
type SessionAffinityConfigApplyConfiguration struct {
	ClientIP *ClientIPConfigApplyConfiguration `json:"clientIP,omitempty"`
}

// CSIVolumeSourceApplyConfiguration represents a declarative configuration of the CSIVolumeSource type for use
// with apply.
// +kubebuilder:object:root=true
type CSIVolumeSourceApplyConfiguration struct {
	Driver               *string                                 `json:"driver,omitempty"`
	ReadOnly             *bool                                   `json:"readOnly,omitempty"`
	FSType               *string                                 `json:"fsType,omitempty"`
	VolumeAttributes     map[string]string                       `json:"volumeAttributes,omitempty"`
	NodePublishSecretRef *LocalObjectReferenceApplyConfiguration `json:"nodePublishSecretRef,omitempty"`
}

// EnvFromSourceApplyConfiguration represents a declarative configuration of the EnvFromSource type for use
// with apply.
// +kubebuilder:object:root=true
type EnvFromSourceApplyConfiguration struct {
	Prefix       *string                               `json:"prefix,omitempty"`
	ConfigMapRef *ConfigMapEnvSourceApplyConfiguration `json:"configMapRef,omitempty"`
	SecretRef    *SecretEnvSourceApplyConfiguration    `json:"secretRef,omitempty"`
}

// PodSpecApplyConfiguration represents a declarative configuration of the PodSpec type for use
// with apply.
// +kubebuilder:object:root=true
type PodSpecApplyConfiguration struct {
	Volumes                       []VolumeApplyConfiguration                   `json:"volumes,omitempty"`
	InitContainers                []ContainerApplyConfiguration                `json:"initContainers,omitempty"`
	Containers                    []ContainerApplyConfiguration                `json:"containers,omitempty"`
	RestartPolicy                 *string                                      `json:"restartPolicy,omitempty"`
	TerminationGracePeriodSeconds *int64                                       `json:"terminationGracePeriodSeconds,omitempty"`
	ActiveDeadlineSeconds         *int64                                       `json:"activeDeadlineSeconds,omitempty"`
	DNSPolicy                     *string                                      `json:"dnsPolicy,omitempty"`
	NodeSelector                  map[string]string                            `json:"nodeSelector,omitempty"`
	ServiceAccountName            *string                                      `json:"serviceAccountName,omitempty"`
	DeprecatedServiceAccount      *string                                      `json:"serviceAccount,omitempty"`
	AutomountServiceAccountToken  *bool                                        `json:"automountServiceAccountToken,omitempty"`
	NodeName                      *string                                      `json:"nodeName,omitempty"`
	HostNetwork                   *bool                                        `json:"hostNetwork,omitempty"`
	HostPID                       *bool                                        `json:"hostPID,omitempty"`
	HostIPC                       *bool                                        `json:"hostIPC,omitempty"`
	ShareProcessNamespace         *bool                                        `json:"shareProcessNamespace,omitempty"`
	SecurityContext               *PodSecurityContextApplyConfiguration        `json:"securityContext,omitempty"`
	ImagePullSecrets              []LocalObjectReferenceApplyConfiguration     `json:"imagePullSecrets,omitempty"`
	Hostname                      *string                                      `json:"hostname,omitempty"`
	Subdomain                     *string                                      `json:"subdomain,omitempty"`
	Affinity                      *AffinityApplyConfiguration                  `json:"affinity,omitempty"`
	SchedulerName                 *string                                      `json:"schedulerName,omitempty"`
	Tolerations                   []TolerationApplyConfiguration               `json:"tolerations,omitempty"`
	HostAliases                   []HostAliasApplyConfiguration                `json:"hostAliases,omitempty"`
	PriorityClassName             *string                                      `json:"priorityClassName,omitempty"`
	Priority                      *int32                                       `json:"priority,omitempty"`
	DNSConfig                     *PodDNSConfigApplyConfiguration              `json:"dnsConfig,omitempty"`
	ReadinessGates                []PodReadinessGateApplyConfiguration         `json:"readinessGates,omitempty"`
	RuntimeClassName              *string                                      `json:"runtimeClassName,omitempty"`
	EnableServiceLinks            *bool                                        `json:"enableServiceLinks,omitempty"`
	PreemptionPolicy              *string                                      `json:"preemptionPolicy,omitempty"`
	Overhead                      *map[string]resource.Quantity                `json:"overhead,omitempty"`
	TopologySpreadConstraints     []TopologySpreadConstraintApplyConfiguration `json:"topologySpreadConstraints,omitempty"`
	SetHostnameAsFQDN             *bool                                        `json:"setHostnameAsFQDN,omitempty"`
	OS                            *PodOSApplyConfiguration                     `json:"os,omitempty"`
	HostUsers                     *bool                                        `json:"hostUsers,omitempty"`
	SchedulingGates               []PodSchedulingGateApplyConfiguration        `json:"schedulingGates,omitempty"`
	ResourceClaims                []PodResourceClaimApplyConfiguration         `json:"resourceClaims,omitempty"`
}

// EndpointPortApplyConfiguration represents a declarative configuration of the EndpointPort type for use
// with apply.
// +kubebuilder:object:root=true
type EndpointPortApplyConfiguration struct {
	Name        *string `json:"name,omitempty"`
	Port        *int32  `json:"port,omitempty"`
	Protocol    *string `json:"protocol,omitempty"`
	AppProtocol *string `json:"appProtocol,omitempty"`
}

// HTTPHeaderApplyConfiguration represents a declarative configuration of the HTTPHeader type for use
// with apply.
// +kubebuilder:object:root=true
type HTTPHeaderApplyConfiguration struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

// PersistentVolumeClaimConditionApplyConfiguration represents a declarative configuration of the PersistentVolumeClaimCondition type for use
// with apply.
// +kubebuilder:object:root=true
type PersistentVolumeClaimConditionApplyConfiguration struct {
	Type               *string      `json:"type,omitempty"`
	Status             *string      `json:"status,omitempty"`
	LastProbeTime      *metav1.Time `json:"lastProbeTime,omitempty"`
	LastTransitionTime *metav1.Time `json:"lastTransitionTime,omitempty"`
	Reason             *string      `json:"reason,omitempty"`
	Message            *string      `json:"message,omitempty"`
}

// ContainerPortApplyConfiguration represents a declarative configuration of the ContainerPort type for use
// with apply.
// +kubebuilder:object:root=true
type ContainerPortApplyConfiguration struct {
	Name          *string `json:"name,omitempty"`
	HostPort      *int32  `json:"hostPort,omitempty"`
	ContainerPort *int32  `json:"containerPort,omitempty"`
	Protocol      string  `json:"protocol,omitempty"`
	HostIP        *string `json:"hostIP,omitempty"`
}

// DownwardAPIProjectionApplyConfiguration represents a declarative configuration of the DownwardAPIProjection type for use
// with apply.
// +kubebuilder:object:root=true
type DownwardAPIProjectionApplyConfiguration struct {
	Items []DownwardAPIVolumeFileApplyConfiguration `json:"items,omitempty"`
}

// WindowsSecurityContextOptionsApplyConfiguration represents a declarative configuration of the WindowsSecurityContextOptions type for use
// with apply.
// +kubebuilder:object:root=true
type WindowsSecurityContextOptionsApplyConfiguration struct {
	GMSACredentialSpecName *string `json:"gmsaCredentialSpecName,omitempty"`
	GMSACredentialSpec     *string `json:"gmsaCredentialSpec,omitempty"`
	RunAsUserName          *string `json:"runAsUserName,omitempty"`
	HostProcess            *bool   `json:"hostProcess,omitempty"`
}

// ScopeSelectorApplyConfiguration represents a declarative configuration of the ScopeSelector type for use
// with apply.
// +kubebuilder:object:root=true
type ScopeSelectorApplyConfiguration struct {
	MatchExpressions []ScopedResourceSelectorRequirementApplyConfiguration `json:"matchExpressions,omitempty"`
}

// VolumeProjectionApplyConfiguration represents a declarative configuration of the VolumeProjection type for use
// with apply.
// +kubebuilder:object:root=true
type VolumeProjectionApplyConfiguration struct {
	Secret              *SecretProjectionApplyConfiguration              `json:"secret,omitempty"`
	DownwardAPI         *DownwardAPIProjectionApplyConfiguration         `json:"downwardAPI,omitempty"`
	ConfigMap           *ConfigMapProjectionApplyConfiguration           `json:"configMap,omitempty"`
	ServiceAccountToken *ServiceAccountTokenProjectionApplyConfiguration `json:"serviceAccountToken,omitempty"`
}

// NodeSelectorRequirementApplyConfiguration represents a declarative configuration of the NodeSelectorRequirement type for use
// with apply.
// +kubebuilder:object:root=true
type NodeSelectorRequirementApplyConfiguration struct {
	Key      *string  `json:"key,omitempty"`
	Operator *string  `json:"operator,omitempty"`
	Values   []string `json:"values,omitempty"`
}

// ServiceAccountApplyConfiguration represents a declarative configuration of the ServiceAccount type for use
// with apply.
// +kubebuilder:object:root=true
type ServiceAccountApplyConfiguration struct {
	metav1.TypeMeta              `json:",inline"`
	metav1.ObjectMeta            `json:"metadata,omitempty"`
	Secrets                      []ObjectReferenceApplyConfiguration      `json:"secrets,omitempty"`
	ImagePullSecrets             []LocalObjectReferenceApplyConfiguration `json:"imagePullSecrets,omitempty"`
	AutomountServiceAccountToken *bool                                    `json:"automountServiceAccountToken,omitempty"`
}

// ClaimSourceApplyConfiguration represents a declarative configuration of the ClaimSource type for use
// with apply.
// +kubebuilder:object:root=true
type ClaimSourceApplyConfiguration struct {
	ResourceClaimName         *string `json:"resourceClaimName,omitempty"`
	ResourceClaimTemplateName *string `json:"resourceClaimTemplateName,omitempty"`
}

// ContainerStateTerminatedApplyConfiguration represents a declarative configuration of the ContainerStateTerminated type for use
// with apply.
// +kubebuilder:object:root=true
type ContainerStateTerminatedApplyConfiguration struct {
	ExitCode    *int32       `json:"exitCode,omitempty"`
	Signal      *int32       `json:"signal,omitempty"`
	Reason      *string      `json:"reason,omitempty"`
	Message     *string      `json:"message,omitempty"`
	StartedAt   *metav1.Time `json:"startedAt,omitempty"`
	FinishedAt  *metav1.Time `json:"finishedAt,omitempty"`
	ContainerID *string      `json:"containerID,omitempty"`
}

// PortStatusApplyConfiguration represents a declarative configuration of the PortStatus type for use
// with apply.
// +kubebuilder:object:root=true
type PortStatusApplyConfiguration struct {
	Port     *int32  `json:"port,omitempty"`
	Protocol string  `json:"protocol,omitempty"`
	Error    *string `json:"error,omitempty"`
}

// StorageOSVolumeSourceApplyConfiguration represents a declarative configuration of the StorageOSVolumeSource type for use
// with apply.
// +kubebuilder:object:root=true
type StorageOSVolumeSourceApplyConfiguration struct {
	VolumeName      *string                                 `json:"volumeName,omitempty"`
	VolumeNamespace *string                                 `json:"volumeNamespace,omitempty"`
	FSType          *string                                 `json:"fsType,omitempty"`
	ReadOnly        *bool                                   `json:"readOnly,omitempty"`
	SecretRef       *LocalObjectReferenceApplyConfiguration `json:"secretRef,omitempty"`
}

// PodStatusApplyConfiguration represents a declarative configuration of the PodStatus type for use
// with apply.
// +kubebuilder:object:root=true
type PodStatusApplyConfiguration struct {
	Phase                      *string                             `json:"phase,omitempty"`
	Conditions                 []PodConditionApplyConfiguration    `json:"conditions,omitempty"`
	Message                    *string                             `json:"message,omitempty"`
	Reason                     *string                             `json:"reason,omitempty"`
	NominatedNodeName          *string                             `json:"nominatedNodeName,omitempty"`
	HostIP                     *string                             `json:"hostIP,omitempty"`
	PodIP                      *string                             `json:"podIP,omitempty"`
	PodIPs                     []PodIPApplyConfiguration           `json:"podIPs,omitempty"`
	StartTime                  *metav1.Time                        `json:"startTime,omitempty"`
	InitContainerStatuses      []ContainerStatusApplyConfiguration `json:"initContainerStatuses,omitempty"`
	ContainerStatuses          []ContainerStatusApplyConfiguration `json:"containerStatuses,omitempty"`
	QOSClass                   *string                             `json:"qosClass,omitempty"`
	EphemeralContainerStatuses []ContainerStatusApplyConfiguration `json:"ephemeralContainerStatuses,omitempty"`
	Resize                     *string                             `json:"resize,omitempty"`
}

// ContainerStateRunningApplyConfiguration represents a declarative configuration of the ContainerStateRunning type for use
// with apply.
// +kubebuilder:object:root=true
type ContainerStateRunningApplyConfiguration struct {
	StartedAt *metav1.Time `json:"startedAt,omitempty"`
}

// ServiceAccountTokenProjectionApplyConfiguration represents a declarative configuration of the ServiceAccountTokenProjection type for use
// with apply.
// +kubebuilder:object:root=true
type ServiceAccountTokenProjectionApplyConfiguration struct {
	Audience          *string `json:"audience,omitempty"`
	ExpirationSeconds *int64  `json:"expirationSeconds,omitempty"`
	Path              *string `json:"path,omitempty"`
}

// ProbeApplyConfiguration represents a declarative configuration of the Probe type for use
// with apply.
// +kubebuilder:object:root=true
type ProbeApplyConfiguration struct {
	ProbeHandlerApplyConfiguration `json:",inline"`
	InitialDelaySeconds            *int32 `json:"initialDelaySeconds,omitempty"`
	TimeoutSeconds                 *int32 `json:"timeoutSeconds,omitempty"`
	PeriodSeconds                  *int32 `json:"periodSeconds,omitempty"`
	SuccessThreshold               *int32 `json:"successThreshold,omitempty"`
	FailureThreshold               *int32 `json:"failureThreshold,omitempty"`
	TerminationGracePeriodSeconds  *int64 `json:"terminationGracePeriodSeconds,omitempty"`
}

// ObjectReferenceApplyConfiguration represents a declarative configuration of the ObjectReference type for use
// with apply.
// +kubebuilder:object:root=true
type ObjectReferenceApplyConfiguration struct {
	Kind            *string    `json:"kind,omitempty"`
	Namespace       *string    `json:"namespace,omitempty"`
	Name            *string    `json:"name,omitempty"`
	UID             *types.UID `json:"uid,omitempty"`
	APIVersion      *string    `json:"apiVersion,omitempty"`
	ResourceVersion *string    `json:"resourceVersion,omitempty"`
	FieldPath       *string    `json:"fieldPath,omitempty"`
}

// LoadBalancerStatusApplyConfiguration represents a declarative configuration of the LoadBalancerStatus type for use
// with apply.
// +kubebuilder:object:root=true
type LoadBalancerStatusApplyConfiguration struct {
	Ingress []LoadBalancerIngressApplyConfiguration `json:"ingress,omitempty"`
}

// LifecycleApplyConfiguration represents a declarative configuration of the Lifecycle type for use
// with apply.
// +kubebuilder:object:root=true
type LifecycleApplyConfiguration struct {
	PostStart *LifecycleHandlerApplyConfiguration `json:"postStart,omitempty"`
	PreStop   *LifecycleHandlerApplyConfiguration `json:"preStop,omitempty"`
}

// FlexPersistentVolumeSourceApplyConfiguration represents a declarative configuration of the FlexPersistentVolumeSource type for use
// with apply.
// +kubebuilder:object:root=true
type FlexPersistentVolumeSourceApplyConfiguration struct {
	Driver    *string                            `json:"driver,omitempty"`
	FSType    *string                            `json:"fsType,omitempty"`
	SecretRef *SecretReferenceApplyConfiguration `json:"secretRef,omitempty"`
	ReadOnly  *bool                              `json:"readOnly,omitempty"`
	Options   map[string]string                  `json:"options,omitempty"`
}

// PodAntiAffinityApplyConfiguration represents a declarative configuration of the PodAntiAffinity type for use
// with apply.
// +kubebuilder:object:root=true
type PodAntiAffinityApplyConfiguration struct {
	RequiredDuringSchedulingIgnoredDuringExecution  []PodAffinityTermApplyConfiguration         `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
	PreferredDuringSchedulingIgnoredDuringExecution []WeightedPodAffinityTermApplyConfiguration `json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
}

// NodeConfigSourceApplyConfiguration represents a declarative configuration of the NodeConfigSource type for use
// with apply.
// +kubebuilder:object:root=true
type NodeConfigSourceApplyConfiguration struct {
	ConfigMap *ConfigMapNodeConfigSourceApplyConfiguration `json:"configMap,omitempty"`
}

// EndpointAddressApplyConfiguration represents a declarative configuration of the EndpointAddress type for use
// with apply.
// +kubebuilder:object:root=true
type EndpointAddressApplyConfiguration struct {
	IP        *string                            `json:"ip,omitempty"`
	Hostname  *string                            `json:"hostname,omitempty"`
	NodeName  *string                            `json:"nodeName,omitempty"`
	TargetRef *ObjectReferenceApplyConfiguration `json:"targetRef,omitempty"`
}

// SecretApplyConfiguration represents a declarative configuration of the Secret type for use
// with apply.
// +kubebuilder:object:root=true
type SecretApplyConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Immutable         *bool             `json:"immutable,omitempty"`
	Data              map[string][]byte `json:"data,omitempty"`
	StringData        map[string]string `json:"stringData,omitempty"`
	Type              *string           `json:"type,omitempty"`
}

// ProbeHandlerApplyConfiguration represents a declarative configuration of the ProbeHandler type for use
// with apply.
// +kubebuilder:object:root=true
type ProbeHandlerApplyConfiguration struct {
	Exec      *ExecActionApplyConfiguration      `json:"exec,omitempty"`
	HTTPGet   *HTTPGetActionApplyConfiguration   `json:"httpGet,omitempty"`
	TCPSocket *TCPSocketActionApplyConfiguration `json:"tcpSocket,omitempty"`
	GRPC      *GRPCActionApplyConfiguration      `json:"grpc,omitempty"`
}

// CapabilitiesApplyConfiguration represents a declarative configuration of the Capabilities type for use
// with apply.
// +kubebuilder:object:root=true
type CapabilitiesApplyConfiguration struct {
	Add  []string `json:"add,omitempty"`
	Drop []string `json:"drop,omitempty"`
}

// PortworxVolumeSourceApplyConfiguration represents a declarative configuration of the PortworxVolumeSource type for use
// with apply.
// +kubebuilder:object:root=true
type PortworxVolumeSourceApplyConfiguration struct {
	VolumeID *string `json:"volumeID,omitempty"`
	FSType   *string `json:"fsType,omitempty"`
	ReadOnly *bool   `json:"readOnly,omitempty"`
}

// ScaleIOVolumeSourceApplyConfiguration represents a declarative configuration of the ScaleIOVolumeSource type for use
// with apply.
// +kubebuilder:object:root=true
type ScaleIOVolumeSourceApplyConfiguration struct {
	Gateway          *string                                 `json:"gateway,omitempty"`
	System           *string                                 `json:"system,omitempty"`
	SecretRef        *LocalObjectReferenceApplyConfiguration `json:"secretRef,omitempty"`
	SSLEnabled       *bool                                   `json:"sslEnabled,omitempty"`
	ProtectionDomain *string                                 `json:"protectionDomain,omitempty"`
	StoragePool      *string                                 `json:"storagePool,omitempty"`
	StorageMode      *string                                 `json:"storageMode,omitempty"`
	VolumeName       *string                                 `json:"volumeName,omitempty"`
	FSType           *string                                 `json:"fsType,omitempty"`
	ReadOnly         *bool                                   `json:"readOnly,omitempty"`
}

// TopologySelectorTermApplyConfiguration represents a declarative configuration of the TopologySelectorTerm type for use
// with apply.
// +kubebuilder:object:root=true
type TopologySelectorTermApplyConfiguration struct {
	MatchLabelExpressions []TopologySelectorLabelRequirementApplyConfiguration `json:"matchLabelExpressions,omitempty"`
}

// TypedLocalObjectReferenceApplyConfiguration represents a declarative configuration of the TypedLocalObjectReference type for use
// with apply.
// +kubebuilder:object:root=true
type TypedLocalObjectReferenceApplyConfiguration struct {
	APIGroup *string `json:"apiGroup,omitempty"`
	Kind     *string `json:"kind,omitempty"`
	Name     *string `json:"name,omitempty"`
}

// GitRepoVolumeSourceApplyConfiguration represents a declarative configuration of the GitRepoVolumeSource type for use
// with apply.
// +kubebuilder:object:root=true
type GitRepoVolumeSourceApplyConfiguration struct {
	Repository *string `json:"repository,omitempty"`
	Revision   *string `json:"revision,omitempty"`
	Directory  *string `json:"directory,omitempty"`
}

// PodDNSConfigApplyConfiguration represents a declarative configuration of the PodDNSConfig type for use
// with apply.
// +kubebuilder:object:root=true
type PodDNSConfigApplyConfiguration struct {
	Nameservers []string                               `json:"nameservers,omitempty"`
	Searches    []string                               `json:"searches,omitempty"`
	Options     []PodDNSConfigOptionApplyConfiguration `json:"options,omitempty"`
}

// DownwardAPIVolumeFileApplyConfiguration represents a declarative configuration of the DownwardAPIVolumeFile type for use
// with apply.
// +kubebuilder:object:root=true
type DownwardAPIVolumeFileApplyConfiguration struct {
	Path             *string                                  `json:"path,omitempty"`
	FieldRef         *ObjectFieldSelectorApplyConfiguration   `json:"fieldRef,omitempty"`
	ResourceFieldRef *ResourceFieldSelectorApplyConfiguration `json:"resourceFieldRef,omitempty"`
	Mode             *int32                                   `json:"mode,omitempty"`
}

// LifecycleHandlerApplyConfiguration represents a declarative configuration of the LifecycleHandler type for use
// with apply.
// +kubebuilder:object:root=true
type LifecycleHandlerApplyConfiguration struct {
	Exec      *ExecActionApplyConfiguration      `json:"exec,omitempty"`
	HTTPGet   *HTTPGetActionApplyConfiguration   `json:"httpGet,omitempty"`
	TCPSocket *TCPSocketActionApplyConfiguration `json:"tcpSocket,omitempty"`
}

// PodIPApplyConfiguration represents a declarative configuration of the PodIP type for use
// with apply.
// +kubebuilder:object:root=true
type PodIPApplyConfiguration struct {
	IP *string `json:"ip,omitempty"`
}

// SecurityContextApplyConfiguration represents a declarative configuration of the SecurityContext type for use
// with apply.
// +kubebuilder:object:root=true
type SecurityContextApplyConfiguration struct {
	Capabilities             *CapabilitiesApplyConfiguration                  `json:"capabilities,omitempty"`
	Privileged               *bool                                            `json:"privileged,omitempty"`
	SELinuxOptions           *SELinuxOptionsApplyConfiguration                `json:"seLinuxOptions,omitempty"`
	WindowsOptions           *WindowsSecurityContextOptionsApplyConfiguration `json:"windowsOptions,omitempty"`
	RunAsUser                *int64                                           `json:"runAsUser,omitempty"`
	RunAsGroup               *int64                                           `json:"runAsGroup,omitempty"`
	RunAsNonRoot             *bool                                            `json:"runAsNonRoot,omitempty"`
	ReadOnlyRootFilesystem   *bool                                            `json:"readOnlyRootFilesystem,omitempty"`
	AllowPrivilegeEscalation *bool                                            `json:"allowPrivilegeEscalation,omitempty"`
	ProcMount                *string                                          `json:"procMount,omitempty"`
	SeccompProfile           *SeccompProfileApplyConfiguration                `json:"seccompProfile,omitempty"`
}

// VolumeDeviceApplyConfiguration represents a declarative configuration of the VolumeDevice type for use
// with apply.
// +kubebuilder:object:root=true
type VolumeDeviceApplyConfiguration struct {
	Name       *string `json:"name,omitempty"`
	DevicePath *string `json:"devicePath,omitempty"`
}

// NodeConfigStatusApplyConfiguration represents a declarative configuration of the NodeConfigStatus type for use
// with apply.
// +kubebuilder:object:root=true
type NodeConfigStatusApplyConfiguration struct {
	Assigned      *NodeConfigSourceApplyConfiguration `json:"assigned,omitempty"`
	Active        *NodeConfigSourceApplyConfiguration `json:"active,omitempty"`
	LastKnownGood *NodeConfigSourceApplyConfiguration `json:"lastKnownGood,omitempty"`
	Error         *string                             `json:"error,omitempty"`
}

// EventSeriesApplyConfiguration represents a declarative configuration of the EventSeries type for use
// with apply.
// +kubebuilder:object:root=true
type EventSeriesApplyConfiguration struct {
	Count            *int32            `json:"count,omitempty"`
	LastObservedTime *metav1.MicroTime `json:"lastObservedTime,omitempty"`
}

// ConfigMapApplyConfiguration represents a declarative configuration of the ConfigMap type for use
// with apply.
// +kubebuilder:object:root=true
type ConfigMapApplyConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Immutable         *bool             `json:"immutable,omitempty"`
	Data              map[string]string `json:"data,omitempty"`
	BinaryData        map[string][]byte `json:"binaryData,omitempty"`
}

// EnvVarApplyConfiguration represents a declarative configuration of the EnvVar type for use
// with apply.
// +kubebuilder:object:root=true
type EnvVarApplyConfiguration struct {
	Name      *string                         `json:"name,omitempty"`
	Value     *string                         `json:"value,omitempty"`
	ValueFrom *EnvVarSourceApplyConfiguration `json:"valueFrom,omitempty"`
}

// EndpointSubsetApplyConfiguration represents a declarative configuration of the EndpointSubset type for use
// with apply.
// +kubebuilder:object:root=true
type EndpointSubsetApplyConfiguration struct {
	Addresses         []EndpointAddressApplyConfiguration `json:"addresses,omitempty"`
	NotReadyAddresses []EndpointAddressApplyConfiguration `json:"notReadyAddresses,omitempty"`
	Ports             []EndpointPortApplyConfiguration    `json:"ports,omitempty"`
}

// ResourceClaimApplyConfiguration represents a declarative configuration of the ResourceClaim type for use
// with apply.
// +kubebuilder:object:root=true
type ResourceClaimApplyConfiguration struct {
	Name *string `json:"name,omitempty"`
}

// VolumeApplyConfiguration represents a declarative configuration of the Volume type for use
// with apply.
// +kubebuilder:object:root=true
type VolumeApplyConfiguration struct {
	Name                           *string `json:"name,omitempty"`
	VolumeSourceApplyConfiguration `json:",inline"`
}

// NodeSystemInfoApplyConfiguration represents a declarative configuration of the NodeSystemInfo type for use
// with apply.
// +kubebuilder:object:root=true
type NodeSystemInfoApplyConfiguration struct {
	MachineID               *string `json:"machineID,omitempty"`
	SystemUUID              *string `json:"systemUUID,omitempty"`
	BootID                  *string `json:"bootID,omitempty"`
	KernelVersion           *string `json:"kernelVersion,omitempty"`
	OSImage                 *string `json:"osImage,omitempty"`
	ContainerRuntimeVersion *string `json:"containerRuntimeVersion,omitempty"`
	KubeletVersion          *string `json:"kubeletVersion,omitempty"`
	KubeProxyVersion        *string `json:"kubeProxyVersion,omitempty"`
	OperatingSystem         *string `json:"operatingSystem,omitempty"`
	Architecture            *string `json:"architecture,omitempty"`
}

// ContainerApplyConfiguration represents a declarative configuration of the Container type for use
// with apply.
// +kubebuilder:object:root=true
type ContainerApplyConfiguration struct {
	Name                     *string                                   `json:"name,omitempty"`
	Image                    *string                                   `json:"image,omitempty"`
	Command                  []string                                  `json:"command,omitempty"`
	Args                     []string                                  `json:"args,omitempty"`
	WorkingDir               *string                                   `json:"workingDir,omitempty"`
	Ports                    []ContainerPortApplyConfiguration         `json:"ports,omitempty"`
	EnvFrom                  []EnvFromSourceApplyConfiguration         `json:"envFrom,omitempty"`
	Env                      []EnvVarApplyConfiguration                `json:"env,omitempty"`
	Resources                *ResourceRequirementsApplyConfiguration   `json:"resources,omitempty"`
	ResizePolicy             []ContainerResizePolicyApplyConfiguration `json:"resizePolicy,omitempty"`
	VolumeMounts             []VolumeMountApplyConfiguration           `json:"volumeMounts,omitempty"`
	VolumeDevices            []VolumeDeviceApplyConfiguration          `json:"volumeDevices,omitempty"`
	LivenessProbe            *ProbeApplyConfiguration                  `json:"livenessProbe,omitempty"`
	ReadinessProbe           *ProbeApplyConfiguration                  `json:"readinessProbe,omitempty"`
	StartupProbe             *ProbeApplyConfiguration                  `json:"startupProbe,omitempty"`
	Lifecycle                *LifecycleApplyConfiguration              `json:"lifecycle,omitempty"`
	TerminationMessagePath   *string                                   `json:"terminationMessagePath,omitempty"`
	TerminationMessagePolicy *string                                   `json:"terminationMessagePolicy,omitempty"`
	ImagePullPolicy          *string                                   `json:"imagePullPolicy,omitempty"`
	SecurityContext          *SecurityContextApplyConfiguration        `json:"securityContext,omitempty"`
	Stdin                    *bool                                     `json:"stdin,omitempty"`
	StdinOnce                *bool                                     `json:"stdinOnce,omitempty"`
	TTY                      *bool                                     `json:"tty,omitempty"`
}

// SecretReferenceApplyConfiguration represents a declarative configuration of the SecretReference type for use
// with apply.
// +kubebuilder:object:root=true
type SecretReferenceApplyConfiguration struct {
	Name      *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
}

// SecretProjectionApplyConfiguration represents a declarative configuration of the SecretProjection type for use
// with apply.
// +kubebuilder:object:root=true
type SecretProjectionApplyConfiguration struct {
	LocalObjectReferenceApplyConfiguration `json:",inline"`
	Items                                  []KeyToPathApplyConfiguration `json:"items,omitempty"`
	Optional                               *bool                         `json:"optional,omitempty"`
}

// PersistentVolumeSpecApplyConfiguration represents a declarative configuration of the PersistentVolumeSpec type for use
// with apply.
// +kubebuilder:object:root=true
type PersistentVolumeSpecApplyConfiguration struct {
	Capacity                                 *map[string]resource.Quantity `json:"capacity,omitempty"`
	PersistentVolumeSourceApplyConfiguration `json:",inline"`
	AccessModes                              []string                              `json:"accessModes,omitempty"`
	ClaimRef                                 *ObjectReferenceApplyConfiguration    `json:"claimRef,omitempty"`
	PersistentVolumeReclaimPolicy            *string                               `json:"persistentVolumeReclaimPolicy,omitempty"`
	StorageClassName                         *string                               `json:"storageClassName,omitempty"`
	MountOptions                             []string                              `json:"mountOptions,omitempty"`
	VolumeMode                               *string                               `json:"volumeMode,omitempty"`
	NodeAffinity                             *VolumeNodeAffinityApplyConfiguration `json:"nodeAffinity,omitempty"`
}

// SELinuxOptionsApplyConfiguration represents a declarative configuration of the SELinuxOptions type for use
// with apply.
// +kubebuilder:object:root=true
type SELinuxOptionsApplyConfiguration struct {
	User  *string `json:"user,omitempty"`
	Role  *string `json:"role,omitempty"`
	Type  *string `json:"type,omitempty"`
	Level *string `json:"level,omitempty"`
}

// PodTemplateSpecApplyConfiguration represents a declarative configuration of the PodTemplateSpec type for use
// with apply.
// +kubebuilder:object:root=true
type PodTemplateSpecApplyConfiguration struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              *PodSpecApplyConfiguration `json:"spec,omitempty"`
}

// FlexVolumeSourceApplyConfiguration represents a declarative configuration of the FlexVolumeSource type for use
// with apply.
// +kubebuilder:object:root=true
type FlexVolumeSourceApplyConfiguration struct {
	Driver    *string                                 `json:"driver,omitempty"`
	FSType    *string                                 `json:"fsType,omitempty"`
	SecretRef *LocalObjectReferenceApplyConfiguration `json:"secretRef,omitempty"`
	ReadOnly  *bool                                   `json:"readOnly,omitempty"`
	Options   map[string]string                       `json:"options,omitempty"`
}

// ConfigMapProjectionApplyConfiguration represents a declarative configuration of the ConfigMapProjection type for use
// with apply.
// +kubebuilder:object:root=true
type ConfigMapProjectionApplyConfiguration struct {
	LocalObjectReferenceApplyConfiguration `json:",inline"`
	Items                                  []KeyToPathApplyConfiguration `json:"items,omitempty"`
	Optional                               *bool                         `json:"optional,omitempty"`
}

// GlusterfsVolumeSourceApplyConfiguration represents a declarative configuration of the GlusterfsVolumeSource type for use
// with apply.
// +kubebuilder:object:root=true
type GlusterfsVolumeSourceApplyConfiguration struct {
	EndpointsName *string `json:"endpoints,omitempty"`
	Path          *string `json:"path,omitempty"`
	ReadOnly      *bool   `json:"readOnly,omitempty"`
}

// PodAffinityTermApplyConfiguration represents a declarative configuration of the PodAffinityTerm type for use
// with apply.
// +kubebuilder:object:root=true
type PodAffinityTermApplyConfiguration struct {
	LabelSelector     LabelSelectorApplyConfiguration `json:"labelSelector,omitempty"`
	Namespaces        []string                        `json:"namespaces,omitempty"`
	TopologyKey       *string                         `json:"topologyKey,omitempty"`
	NamespaceSelector LabelSelectorApplyConfiguration `json:"namespaceSelector,omitempty"`
}

// EndpointsApplyConfiguration represents a declarative configuration of the Endpoints type for use
// with apply.
// +kubebuilder:object:root=true
type EndpointsApplyConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Subsets           []EndpointSubsetApplyConfiguration `json:"subsets,omitempty"`
}

// WeightedPodAffinityTermApplyConfiguration represents a declarative configuration of the WeightedPodAffinityTerm type for use
// with apply.
// +kubebuilder:object:root=true
type WeightedPodAffinityTermApplyConfiguration struct {
	Weight          *int32                             `json:"weight,omitempty"`
	PodAffinityTerm *PodAffinityTermApplyConfiguration `json:"podAffinityTerm,omitempty"`
}

// TolerationApplyConfiguration represents a declarative configuration of the Toleration type for use
// with apply.
// +kubebuilder:object:root=true
type TolerationApplyConfiguration struct {
	Key               *string `json:"key,omitempty"`
	Operator          *string `json:"operator,omitempty"`
	Value             *string `json:"value,omitempty"`
	Effect            *string `json:"effect,omitempty"`
	TolerationSeconds *int64  `json:"tolerationSeconds,omitempty"`
}

// HostPathVolumeSourceApplyConfiguration represents a declarative configuration of the HostPathVolumeSource type for use
// with apply.
// +kubebuilder:object:root=true
type HostPathVolumeSourceApplyConfiguration struct {
	Path *string `json:"path,omitempty"`
	Type *string `json:"type,omitempty"`
}

// TCPSocketActionApplyConfiguration represents a declarative configuration of the TCPSocketAction type for use
// with apply.
// +kubebuilder:object:root=true
type TCPSocketActionApplyConfiguration struct {
	Port *intstr.IntOrString `json:"port,omitempty"`
	Host *string             `json:"host,omitempty"`
}

// PodConditionApplyConfiguration represents a declarative configuration of the PodCondition type for use
// with apply.
// +kubebuilder:object:root=true
type PodConditionApplyConfiguration struct {
	Type               *string      `json:"type,omitempty"`
	Status             *string      `json:"status,omitempty"`
	LastProbeTime      *metav1.Time `json:"lastProbeTime,omitempty"`
	LastTransitionTime *metav1.Time `json:"lastTransitionTime,omitempty"`
	Reason             *string      `json:"reason,omitempty"`
	Message            *string      `json:"message,omitempty"`
}

// PersistentVolumeApplyConfiguration represents a declarative configuration of the PersistentVolume type for use
// with apply.
// +kubebuilder:object:root=true
type PersistentVolumeApplyConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              *PersistentVolumeSpecApplyConfiguration   `json:"spec,omitempty"`
	Status            *PersistentVolumeStatusApplyConfiguration `json:"status,omitempty"`
}

// ConfigMapVolumeSourceApplyConfiguration represents a declarative configuration of the ConfigMapVolumeSource type for use
// with apply.
// +kubebuilder:object:root=true
type ConfigMapVolumeSourceApplyConfiguration struct {
	LocalObjectReferenceApplyConfiguration `json:",inline"`
	Items                                  []KeyToPathApplyConfiguration `json:"items,omitempty"`
	DefaultMode                            *int32                        `json:"defaultMode,omitempty"`
	Optional                               *bool                         `json:"optional,omitempty"`
}

// SecretVolumeSourceApplyConfiguration represents a declarative configuration of the SecretVolumeSource type for use
// with apply.
// +kubebuilder:object:root=true
type SecretVolumeSourceApplyConfiguration struct {
	SecretName  *string                       `json:"secretName,omitempty"`
	Items       []KeyToPathApplyConfiguration `json:"items,omitempty"`
	DefaultMode *int32                        `json:"defaultMode,omitempty"`
	Optional    *bool                         `json:"optional,omitempty"`
}

// EnvVarSourceApplyConfiguration represents a declarative configuration of the EnvVarSource type for use
// with apply.
// +kubebuilder:object:root=true
type EnvVarSourceApplyConfiguration struct {
	FieldRef         *ObjectFieldSelectorApplyConfiguration   `json:"fieldRef,omitempty"`
	ResourceFieldRef *ResourceFieldSelectorApplyConfiguration `json:"resourceFieldRef,omitempty"`
	ConfigMapKeyRef  *ConfigMapKeySelectorApplyConfiguration  `json:"configMapKeyRef,omitempty"`
	SecretKeyRef     *SecretKeySelectorApplyConfiguration     `json:"secretKeyRef,omitempty"`
}

// ISCSIPersistentVolumeSourceApplyConfiguration represents a declarative configuration of the ISCSIPersistentVolumeSource type for use
// with apply.
// +kubebuilder:object:root=true
type ISCSIPersistentVolumeSourceApplyConfiguration struct {
	TargetPortal      *string                            `json:"targetPortal,omitempty"`
	IQN               *string                            `json:"iqn,omitempty"`
	Lun               *int32                             `json:"lun,omitempty"`
	ISCSIInterface    *string                            `json:"iscsiInterface,omitempty"`
	FSType            *string                            `json:"fsType,omitempty"`
	ReadOnly          *bool                              `json:"readOnly,omitempty"`
	Portals           []string                           `json:"portals,omitempty"`
	DiscoveryCHAPAuth *bool                              `json:"chapAuthDiscovery,omitempty"`
	SessionCHAPAuth   *bool                              `json:"chapAuthSession,omitempty"`
	SecretRef         *SecretReferenceApplyConfiguration `json:"secretRef,omitempty"`
	InitiatorName     *string                            `json:"initiatorName,omitempty"`
}

// TopologySpreadConstraintApplyConfiguration represents a declarative configuration of the TopologySpreadConstraint type for use
// with apply.
// +kubebuilder:object:root=true
type TopologySpreadConstraintApplyConfiguration struct {
	MaxSkew            *int32                          `json:"maxSkew,omitempty"`
	TopologyKey        *string                         `json:"topologyKey,omitempty"`
	WhenUnsatisfiable  *string                         `json:"whenUnsatisfiable,omitempty"`
	LabelSelector      LabelSelectorApplyConfiguration `json:"labelSelector,omitempty"`
	MinDomains         *int32                          `json:"minDomains,omitempty"`
	NodeAffinityPolicy *string                         `json:"nodeAffinityPolicy,omitempty"`
	NodeTaintsPolicy   *string                         `json:"nodeTaintsPolicy,omitempty"`
	MatchLabelKeys     []string                        `json:"matchLabelKeys,omitempty"`
}

// PodResourceClaimApplyConfiguration represents a declarative configuration of the PodResourceClaim type for use
// with apply.
// +kubebuilder:object:root=true
type PodResourceClaimApplyConfiguration struct {
	Name   *string                        `json:"name,omitempty"`
	Source *ClaimSourceApplyConfiguration `json:"source,omitempty"`
}

// ClientIPConfigApplyConfiguration represents a declarative configuration of the ClientIPConfig type for use
// with apply.
// +kubebuilder:object:root=true
type ClientIPConfigApplyConfiguration struct {
	TimeoutSeconds *int32 `json:"timeoutSeconds,omitempty"`
}

// SeccompProfileApplyConfiguration represents a declarative configuration of the SeccompProfile type for use
// with apply.
// +kubebuilder:object:root=true
type SeccompProfileApplyConfiguration struct {
	Type             *string `json:"type,omitempty"`
	LocalhostProfile *string `json:"localhostProfile,omitempty"`
}

// HostAliasApplyConfiguration represents a declarative configuration of the HostAlias type for use
// with apply.
// +kubebuilder:object:root=true
type HostAliasApplyConfiguration struct {
	IP        *string  `json:"ip,omitempty"`
	Hostnames []string `json:"hostnames,omitempty"`
}

// VolumeNodeAffinityApplyConfiguration represents a declarative configuration of the VolumeNodeAffinity type for use
// with apply.
// +kubebuilder:object:root=true
type VolumeNodeAffinityApplyConfiguration struct {
	Required NodeSelectorApplyConfiguration `json:"required,omitempty"`
}

// ExecActionApplyConfiguration represents a declarative configuration of the ExecAction type for use
// with apply.
// +kubebuilder:object:root=true
type ExecActionApplyConfiguration struct {
	Command []string `json:"command,omitempty"`
}

// SysctlApplyConfiguration represents a declarative configuration of the Sysctl type for use
// with apply.
// +kubebuilder:object:root=true
type SysctlApplyConfiguration struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NodeSelectorTermApplyConfiguration represents a declarative configuration of the NodeSelectorTerm type for use
// with apply.
// +kubebuilder:object:root=true
type NodeSelectorTermApplyConfiguration struct {
	MatchExpressions []NodeSelectorRequirementApplyConfiguration `json:"matchExpressions,omitempty"`
	MatchFields      []NodeSelectorRequirementApplyConfiguration `json:"matchFields,omitempty"`
}

// RBDVolumeSourceApplyConfiguration represents a declarative configuration of the RBDVolumeSource type for use
// with apply.
// +kubebuilder:object:root=true
type RBDVolumeSourceApplyConfiguration struct {
	CephMonitors []string                                `json:"monitors,omitempty"`
	RBDImage     *string                                 `json:"image,omitempty"`
	FSType       *string                                 `json:"fsType,omitempty"`
	RBDPool      *string                                 `json:"pool,omitempty"`
	RadosUser    *string                                 `json:"user,omitempty"`
	Keyring      *string                                 `json:"keyring,omitempty"`
	SecretRef    *LocalObjectReferenceApplyConfiguration `json:"secretRef,omitempty"`
	ReadOnly     *bool                                   `json:"readOnly,omitempty"`
}

// TopologySelectorLabelRequirementApplyConfiguration represents a declarative configuration of the TopologySelectorLabelRequirement type for use
// with apply.
// +kubebuilder:object:root=true
type TopologySelectorLabelRequirementApplyConfiguration struct {
	Key    *string  `json:"key,omitempty"`
	Values []string `json:"values,omitempty"`
}

// PodAffinityApplyConfiguration represents a declarative configuration of the PodAffinity type for use
// with apply.
// +kubebuilder:object:root=true
type PodAffinityApplyConfiguration struct {
	RequiredDuringSchedulingIgnoredDuringExecution  []PodAffinityTermApplyConfiguration         `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
	PreferredDuringSchedulingIgnoredDuringExecution []WeightedPodAffinityTermApplyConfiguration `json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
}

// ReplicationControllerConditionApplyConfiguration represents a declarative configuration of the ReplicationControllerCondition type for use
// with apply.
// +kubebuilder:object:root=true
type ReplicationControllerConditionApplyConfiguration struct {
	Type               *string      `json:"type,omitempty"`
	Status             *string      `json:"status,omitempty"`
	LastTransitionTime *metav1.Time `json:"lastTransitionTime,omitempty"`
	Reason             *string      `json:"reason,omitempty"`
	Message            *string      `json:"message,omitempty"`
}

// ConfigMapKeySelectorApplyConfiguration represents a declarative configuration of the ConfigMapKeySelector type for use
// with apply.
// +kubebuilder:object:root=true
type ConfigMapKeySelectorApplyConfiguration struct {
	LocalObjectReferenceApplyConfiguration `json:",inline"`
	Key                                    *string `json:"key,omitempty"`
	Optional                               *bool   `json:"optional,omitempty"`
}

// PreferredSchedulingTermApplyConfiguration represents a declarative configuration of the PreferredSchedulingTerm type for use
// with apply.
// +kubebuilder:object:root=true
type PreferredSchedulingTermApplyConfiguration struct {
	Weight     *int32                              `json:"weight,omitempty"`
	Preference *NodeSelectorTermApplyConfiguration `json:"preference,omitempty"`
}

// PodSecurityContextApplyConfiguration represents a declarative configuration of the PodSecurityContext type for use
// with apply.
// +kubebuilder:object:root=true
type PodSecurityContextApplyConfiguration struct {
	SELinuxOptions      *SELinuxOptionsApplyConfiguration                `json:"seLinuxOptions,omitempty"`
	WindowsOptions      *WindowsSecurityContextOptionsApplyConfiguration `json:"windowsOptions,omitempty"`
	RunAsUser           *int64                                           `json:"runAsUser,omitempty"`
	RunAsGroup          *int64                                           `json:"runAsGroup,omitempty"`
	RunAsNonRoot        *bool                                            `json:"runAsNonRoot,omitempty"`
	SupplementalGroups  []int64                                          `json:"supplementalGroups,omitempty"`
	FSGroup             *int64                                           `json:"fsGroup,omitempty"`
	Sysctls             []SysctlApplyConfiguration                       `json:"sysctls,omitempty"`
	FSGroupChangePolicy *string                                          `json:"fsGroupChangePolicy,omitempty"`
	SeccompProfile      *SeccompProfileApplyConfiguration                `json:"seccompProfile,omitempty"`
}

// ContainerStateWaitingApplyConfiguration represents a declarative configuration of the ContainerStateWaiting type for use
// with apply.
// +kubebuilder:object:root=true
type ContainerStateWaitingApplyConfiguration struct {
	Reason  *string `json:"reason,omitempty"`
	Message *string `json:"message,omitempty"`
}

// AttachedVolumeApplyConfiguration represents a declarative configuration of the AttachedVolume type for use
// with apply.
// +kubebuilder:object:root=true
type AttachedVolumeApplyConfiguration struct {
	Name       *string `json:"name,omitempty"`
	DevicePath *string `json:"devicePath,omitempty"`
}

// EmptyDirVolumeSourceApplyConfiguration represents a declarative configuration of the EmptyDirVolumeSource type for use
// with apply.
// +kubebuilder:object:root=true
type EmptyDirVolumeSourceApplyConfiguration struct {
	Medium    *string            `json:"medium,omitempty"`
	SizeLimit *resource.Quantity `json:"sizeLimit,omitempty"`
}

// LabelSelectorApplyConfiguration represents an declarative configuration of the LabelSelector type for use
// with apply.
// +kubebuilder:object:root=true
type LabelSelectorApplyConfiguration struct {
	MatchLabels      map[string]string                            `json:"matchLabels,omitempty"`
	MatchExpressions []LabelSelectorRequirementApplyConfiguration `json:"matchExpressions,omitempty"`
}

// LabelSelectorRequirementApplyConfiguration represents an declarative configuration of the LabelSelectorRequirement type for use
// with apply.
// +kubebuilder:object:root=true
type LabelSelectorRequirementApplyConfiguration struct {
	Key      *string  `json:"key,omitempty"`
	Operator *string  `json:"operator,omitempty"`
	Values   []string `json:"values,omitempty"`
}
