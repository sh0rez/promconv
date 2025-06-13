package k8s

// The name of the cluster
type AttrClusterName string // k8s.cluster.name

func (AttrClusterName) Development()    {}
func (AttrClusterName) Recommended()    {}
func (AttrClusterName) Key() string     { return "k8s_cluster_name" }
func (a AttrClusterName) Value() string { return string(a) }

// A pseudo-ID for the cluster, set to the UID of the `kube-system` namespace.
//
// K8s doesn't have support for obtaining a cluster ID. If this is ever
// added, we will recommend collecting the `k8s.cluster.uid` through the
// official APIs. In the meantime, we are able to use the `uid` of the
// `kube-system` namespace as a proxy for cluster ID. Read on for the
// rationale.
//
// Every object created in a K8s cluster is assigned a distinct UID. The
// `kube-system` namespace is used by Kubernetes itself and will exist
// for the lifetime of the cluster. Using the `uid` of the `kube-system`
// namespace is a reasonable proxy for the K8s ClusterID as it will only
// change if the cluster is rebuilt. Furthermore, Kubernetes UIDs are
// UUIDs as standardized by
// [ISO/IEC 9834-8 and ITU-T X.667].
// Which states:
//
// > If generated according to one of the mechanisms defined in Rec.
// > ITU-T X.667 | ISO/IEC 9834-8, a UUID is either guaranteed to be
// > different from all other UUIDs generated before 3603 A.D., or is
// > extremely likely to be different (depending on the mechanism chosen).
//
// Therefore, UIDs between clusters should be extremely unlikely to
// conflict
//
// [ISO/IEC 9834-8 and ITU-T X.667]: https://www.itu.int/ITU-T/studygroups/com17/oid.html
type AttrClusterUid string // k8s.cluster.uid

func (AttrClusterUid) Development()    {}
func (AttrClusterUid) Recommended()    {}
func (AttrClusterUid) Key() string     { return "k8s_cluster_uid" }
func (a AttrClusterUid) Value() string { return string(a) }

// The name of the Container from Pod specification, must be unique within a Pod. Container runtime usually uses different globally unique name (`container.name`)
type AttrContainerName string // k8s.container.name

func (AttrContainerName) Development()    {}
func (AttrContainerName) Recommended()    {}
func (AttrContainerName) Key() string     { return "k8s_container_name" }
func (a AttrContainerName) Value() string { return string(a) }

// Number of times the container was restarted. This attribute can be used to identify a particular container (running or stopped) within a container spec
type AttrContainerRestartCount string // k8s.container.restart_count

func (AttrContainerRestartCount) Development()    {}
func (AttrContainerRestartCount) Recommended()    {}
func (AttrContainerRestartCount) Key() string     { return "k8s_container_restart_count" }
func (a AttrContainerRestartCount) Value() string { return string(a) }

// Last terminated reason of the Container
type AttrContainerStatusLastTerminatedReason string // k8s.container.status.last_terminated_reason

func (AttrContainerStatusLastTerminatedReason) Development() {}
func (AttrContainerStatusLastTerminatedReason) Recommended() {}
func (AttrContainerStatusLastTerminatedReason) Key() string {
	return "k8s_container_status_last_terminated_reason"
}
func (a AttrContainerStatusLastTerminatedReason) Value() string { return string(a) }

// The reason for the container state. Corresponds to the `reason` field of the: [K8s ContainerStateWaiting] or [K8s ContainerStateTerminated]
//
// [K8s ContainerStateWaiting]: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#containerstatewaiting-v1-core
// [K8s ContainerStateTerminated]: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#containerstateterminated-v1-core
type AttrContainerStatusReason string // k8s.container.status.reason

func (AttrContainerStatusReason) Development()    {}
func (AttrContainerStatusReason) Recommended()    {}
func (AttrContainerStatusReason) Key() string     { return "k8s_container_status_reason" }
func (a AttrContainerStatusReason) Value() string { return string(a) }

const ContainerStatusReasonContainerCreating AttrContainerStatusReason = "ContainerCreating"
const ContainerStatusReasonCrashLoopBackOff AttrContainerStatusReason = "CrashLoopBackOff"
const ContainerStatusReasonCreateContainerConfigError AttrContainerStatusReason = "CreateContainerConfigError"
const ContainerStatusReasonErrImagePull AttrContainerStatusReason = "ErrImagePull"
const ContainerStatusReasonImagePullBackOff AttrContainerStatusReason = "ImagePullBackOff"
const ContainerStatusReasonOomKilled AttrContainerStatusReason = "OOMKilled"
const ContainerStatusReasonCompleted AttrContainerStatusReason = "Completed"
const ContainerStatusReasonError AttrContainerStatusReason = "Error"
const ContainerStatusReasonContainerCannotRun AttrContainerStatusReason = "ContainerCannotRun"

// The state of the container. [K8s ContainerState]
//
// [K8s ContainerState]: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#containerstate-v1-core
type AttrContainerStatusState string // k8s.container.status.state

func (AttrContainerStatusState) Development()    {}
func (AttrContainerStatusState) Recommended()    {}
func (AttrContainerStatusState) Key() string     { return "k8s_container_status_state" }
func (a AttrContainerStatusState) Value() string { return string(a) }

const ContainerStatusStateTerminated AttrContainerStatusState = "terminated"
const ContainerStatusStateRunning AttrContainerStatusState = "running"
const ContainerStatusStateWaiting AttrContainerStatusState = "waiting"

// The cronjob annotation placed on the CronJob, the `<key>` being the annotation name, the value being the annotation value.
//
// Examples:
//
//   - An annotation `retries` with value `4` SHOULD be recorded as the
//     `k8s.cronjob.annotation.retries` attribute with value `"4"`.
//   - An annotation `data` with empty string value SHOULD be recorded as
//     the `k8s.cronjob.annotation.data` attribute with value `""`
type AttrCronjobAnnotation string // k8s.cronjob.annotation

func (AttrCronjobAnnotation) Development()    {}
func (AttrCronjobAnnotation) Recommended()    {}
func (AttrCronjobAnnotation) Key() string     { return "k8s_cronjob_annotation" }
func (a AttrCronjobAnnotation) Value() string { return string(a) }

// The label placed on the CronJob, the `<key>` being the label name, the value being the label value.
//
// Examples:
//
//   - A label `type` with value `weekly` SHOULD be recorded as the
//     `k8s.cronjob.label.type` attribute with value `"weekly"`.
//   - A label `automated` with empty string value SHOULD be recorded as
//     the `k8s.cronjob.label.automated` attribute with value `""`
type AttrCronjobLabel string // k8s.cronjob.label

func (AttrCronjobLabel) Development()    {}
func (AttrCronjobLabel) Recommended()    {}
func (AttrCronjobLabel) Key() string     { return "k8s_cronjob_label" }
func (a AttrCronjobLabel) Value() string { return string(a) }

// The name of the CronJob
type AttrCronjobName string // k8s.cronjob.name

func (AttrCronjobName) Development()    {}
func (AttrCronjobName) Recommended()    {}
func (AttrCronjobName) Key() string     { return "k8s_cronjob_name" }
func (a AttrCronjobName) Value() string { return string(a) }

// The UID of the CronJob
type AttrCronjobUid string // k8s.cronjob.uid

func (AttrCronjobUid) Development()    {}
func (AttrCronjobUid) Recommended()    {}
func (AttrCronjobUid) Key() string     { return "k8s_cronjob_uid" }
func (a AttrCronjobUid) Value() string { return string(a) }

// The annotation key-value pairs placed on the DaemonSet.
//
// The `<key>` being the annotation name, the value being the annotation value, even if the value is empty
type AttrDaemonsetAnnotation string // k8s.daemonset.annotation

func (AttrDaemonsetAnnotation) Development()    {}
func (AttrDaemonsetAnnotation) Recommended()    {}
func (AttrDaemonsetAnnotation) Key() string     { return "k8s_daemonset_annotation" }
func (a AttrDaemonsetAnnotation) Value() string { return string(a) }

// The label key-value pairs placed on the DaemonSet.
//
// The `<key>` being the label name, the value being the label value, even if the value is empty
type AttrDaemonsetLabel string // k8s.daemonset.label

func (AttrDaemonsetLabel) Development()    {}
func (AttrDaemonsetLabel) Recommended()    {}
func (AttrDaemonsetLabel) Key() string     { return "k8s_daemonset_label" }
func (a AttrDaemonsetLabel) Value() string { return string(a) }

// The name of the DaemonSet
type AttrDaemonsetName string // k8s.daemonset.name

func (AttrDaemonsetName) Development()    {}
func (AttrDaemonsetName) Recommended()    {}
func (AttrDaemonsetName) Key() string     { return "k8s_daemonset_name" }
func (a AttrDaemonsetName) Value() string { return string(a) }

// The UID of the DaemonSet
type AttrDaemonsetUid string // k8s.daemonset.uid

func (AttrDaemonsetUid) Development()    {}
func (AttrDaemonsetUid) Recommended()    {}
func (AttrDaemonsetUid) Key() string     { return "k8s_daemonset_uid" }
func (a AttrDaemonsetUid) Value() string { return string(a) }

// The annotation key-value pairs placed on the Deployment.
//
// The `<key>` being the annotation name, the value being the annotation value, even if the value is empty
type AttrDeploymentAnnotation string // k8s.deployment.annotation

func (AttrDeploymentAnnotation) Development()    {}
func (AttrDeploymentAnnotation) Recommended()    {}
func (AttrDeploymentAnnotation) Key() string     { return "k8s_deployment_annotation" }
func (a AttrDeploymentAnnotation) Value() string { return string(a) }

// The label key-value pairs placed on the Deployment.
//
// The `<key>` being the label name, the value being the label value, even if the value is empty
type AttrDeploymentLabel string // k8s.deployment.label

func (AttrDeploymentLabel) Development()    {}
func (AttrDeploymentLabel) Recommended()    {}
func (AttrDeploymentLabel) Key() string     { return "k8s_deployment_label" }
func (a AttrDeploymentLabel) Value() string { return string(a) }

// The name of the Deployment
type AttrDeploymentName string // k8s.deployment.name

func (AttrDeploymentName) Development()    {}
func (AttrDeploymentName) Recommended()    {}
func (AttrDeploymentName) Key() string     { return "k8s_deployment_name" }
func (a AttrDeploymentName) Value() string { return string(a) }

// The UID of the Deployment
type AttrDeploymentUid string // k8s.deployment.uid

func (AttrDeploymentUid) Development()    {}
func (AttrDeploymentUid) Recommended()    {}
func (AttrDeploymentUid) Key() string     { return "k8s_deployment_uid" }
func (a AttrDeploymentUid) Value() string { return string(a) }

// The name of the horizontal pod autoscaler
type AttrHpaName string // k8s.hpa.name

func (AttrHpaName) Development()    {}
func (AttrHpaName) Recommended()    {}
func (AttrHpaName) Key() string     { return "k8s_hpa_name" }
func (a AttrHpaName) Value() string { return string(a) }

// The API version of the target resource to scale for the HorizontalPodAutoscaler.
//
// This maps to the `apiVersion` field in the `scaleTargetRef` of the HPA spec
type AttrHpaScaletargetrefApiVersion string // k8s.hpa.scaletargetref.api_version

func (AttrHpaScaletargetrefApiVersion) Development()    {}
func (AttrHpaScaletargetrefApiVersion) Recommended()    {}
func (AttrHpaScaletargetrefApiVersion) Key() string     { return "k8s_hpa_scaletargetref_api_version" }
func (a AttrHpaScaletargetrefApiVersion) Value() string { return string(a) }

// The kind of the target resource to scale for the HorizontalPodAutoscaler.
//
// This maps to the `kind` field in the `scaleTargetRef` of the HPA spec
type AttrHpaScaletargetrefKind string // k8s.hpa.scaletargetref.kind

func (AttrHpaScaletargetrefKind) Development()    {}
func (AttrHpaScaletargetrefKind) Recommended()    {}
func (AttrHpaScaletargetrefKind) Key() string     { return "k8s_hpa_scaletargetref_kind" }
func (a AttrHpaScaletargetrefKind) Value() string { return string(a) }

// The name of the target resource to scale for the HorizontalPodAutoscaler.
//
// This maps to the `name` field in the `scaleTargetRef` of the HPA spec
type AttrHpaScaletargetrefName string // k8s.hpa.scaletargetref.name

func (AttrHpaScaletargetrefName) Development()    {}
func (AttrHpaScaletargetrefName) Recommended()    {}
func (AttrHpaScaletargetrefName) Key() string     { return "k8s_hpa_scaletargetref_name" }
func (a AttrHpaScaletargetrefName) Value() string { return string(a) }

// The UID of the horizontal pod autoscaler
type AttrHpaUid string // k8s.hpa.uid

func (AttrHpaUid) Development()    {}
func (AttrHpaUid) Recommended()    {}
func (AttrHpaUid) Key() string     { return "k8s_hpa_uid" }
func (a AttrHpaUid) Value() string { return string(a) }

// The annotation key-value pairs placed on the Job.
//
// The `<key>` being the annotation name, the value being the annotation value, even if the value is empty
type AttrJobAnnotation string // k8s.job.annotation

func (AttrJobAnnotation) Development()    {}
func (AttrJobAnnotation) Recommended()    {}
func (AttrJobAnnotation) Key() string     { return "k8s_job_annotation" }
func (a AttrJobAnnotation) Value() string { return string(a) }

// The label key-value pairs placed on the Job.
//
// The `<key>` being the label name, the value being the label value, even if the value is empty
type AttrJobLabel string // k8s.job.label

func (AttrJobLabel) Development()    {}
func (AttrJobLabel) Recommended()    {}
func (AttrJobLabel) Key() string     { return "k8s_job_label" }
func (a AttrJobLabel) Value() string { return string(a) }

// The name of the Job
type AttrJobName string // k8s.job.name

func (AttrJobName) Development()    {}
func (AttrJobName) Recommended()    {}
func (AttrJobName) Key() string     { return "k8s_job_name" }
func (a AttrJobName) Value() string { return string(a) }

// The UID of the Job
type AttrJobUid string // k8s.job.uid

func (AttrJobUid) Development()    {}
func (AttrJobUid) Recommended()    {}
func (AttrJobUid) Key() string     { return "k8s_job_uid" }
func (a AttrJobUid) Value() string { return string(a) }

// The annotation key-value pairs placed on the Namespace.
//
// The `<key>` being the annotation name, the value being the annotation value, even if the value is empty
type AttrNamespaceAnnotation string // k8s.namespace.annotation

func (AttrNamespaceAnnotation) Development()    {}
func (AttrNamespaceAnnotation) Recommended()    {}
func (AttrNamespaceAnnotation) Key() string     { return "k8s_namespace_annotation" }
func (a AttrNamespaceAnnotation) Value() string { return string(a) }

// The label key-value pairs placed on the Namespace.
//
// The `<key>` being the label name, the value being the label value, even if the value is empty
type AttrNamespaceLabel string // k8s.namespace.label

func (AttrNamespaceLabel) Development()    {}
func (AttrNamespaceLabel) Recommended()    {}
func (AttrNamespaceLabel) Key() string     { return "k8s_namespace_label" }
func (a AttrNamespaceLabel) Value() string { return string(a) }

// The name of the namespace that the pod is running in
type AttrNamespaceName string // k8s.namespace.name

func (AttrNamespaceName) Development()    {}
func (AttrNamespaceName) Recommended()    {}
func (AttrNamespaceName) Key() string     { return "k8s_namespace_name" }
func (a AttrNamespaceName) Value() string { return string(a) }

// The phase of the K8s namespace.
//
// This attribute aligns with the `phase` field of the
// [K8s NamespaceStatus]
//
// [K8s NamespaceStatus]: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#namespacestatus-v1-core
type AttrNamespacePhase string // k8s.namespace.phase

func (AttrNamespacePhase) Development()    {}
func (AttrNamespacePhase) Recommended()    {}
func (AttrNamespacePhase) Key() string     { return "k8s_namespace_phase" }
func (a AttrNamespacePhase) Value() string { return string(a) }

const NamespacePhaseActive AttrNamespacePhase = "active"
const NamespacePhaseTerminating AttrNamespacePhase = "terminating"

// The annotation placed on the Node, the `<key>` being the annotation name, the value being the annotation value, even if the value is empty.
//
// Examples:
//
//   - An annotation `node.alpha.kubernetes.io/ttl` with value `0` SHOULD be recorded as
//     the `k8s.node.annotation.node.alpha.kubernetes.io/ttl` attribute with value `"0"`.
//   - An annotation `data` with empty string value SHOULD be recorded as
//     the `k8s.node.annotation.data` attribute with value `""`
type AttrNodeAnnotation string // k8s.node.annotation

func (AttrNodeAnnotation) Development()    {}
func (AttrNodeAnnotation) Recommended()    {}
func (AttrNodeAnnotation) Key() string     { return "k8s_node_annotation" }
func (a AttrNodeAnnotation) Value() string { return string(a) }

// The label placed on the Node, the `<key>` being the label name, the value being the label value, even if the value is empty.
//
// Examples:
//
//   - A label `kubernetes.io/arch` with value `arm64` SHOULD be recorded
//     as the `k8s.node.label.kubernetes.io/arch` attribute with value `"arm64"`.
//   - A label `data` with empty string value SHOULD be recorded as
//     the `k8s.node.label.data` attribute with value `""`
type AttrNodeLabel string // k8s.node.label

func (AttrNodeLabel) Development()    {}
func (AttrNodeLabel) Recommended()    {}
func (AttrNodeLabel) Key() string     { return "k8s_node_label" }
func (a AttrNodeLabel) Value() string { return string(a) }

// The name of the Node
type AttrNodeName string // k8s.node.name

func (AttrNodeName) Development()    {}
func (AttrNodeName) Recommended()    {}
func (AttrNodeName) Key() string     { return "k8s_node_name" }
func (a AttrNodeName) Value() string { return string(a) }

// The UID of the Node
type AttrNodeUid string // k8s.node.uid

func (AttrNodeUid) Development()    {}
func (AttrNodeUid) Recommended()    {}
func (AttrNodeUid) Key() string     { return "k8s_node_uid" }
func (a AttrNodeUid) Value() string { return string(a) }

// The annotation placed on the Pod, the `<key>` being the annotation name, the value being the annotation value.
//
// Examples:
//
//   - An annotation `kubernetes.io/enforce-mountable-secrets` with value `true` SHOULD be recorded as
//     the `k8s.pod.annotation.kubernetes.io/enforce-mountable-secrets` attribute with value `"true"`.
//   - An annotation `mycompany.io/arch` with value `x64` SHOULD be recorded as
//     the `k8s.pod.annotation.mycompany.io/arch` attribute with value `"x64"`.
//   - An annotation `data` with empty string value SHOULD be recorded as
//     the `k8s.pod.annotation.data` attribute with value `""`
type AttrPodAnnotation string // k8s.pod.annotation

func (AttrPodAnnotation) Development()    {}
func (AttrPodAnnotation) Recommended()    {}
func (AttrPodAnnotation) Key() string     { return "k8s_pod_annotation" }
func (a AttrPodAnnotation) Value() string { return string(a) }

// The label placed on the Pod, the `<key>` being the label name, the value being the label value.
//
// Examples:
//
//   - A label `app` with value `my-app` SHOULD be recorded as
//     the `k8s.pod.label.app` attribute with value `"my-app"`.
//   - A label `mycompany.io/arch` with value `x64` SHOULD be recorded as
//     the `k8s.pod.label.mycompany.io/arch` attribute with value `"x64"`.
//   - A label `data` with empty string value SHOULD be recorded as
//     the `k8s.pod.label.data` attribute with value `""`
type AttrPodLabel string // k8s.pod.label

func (AttrPodLabel) Development()    {}
func (AttrPodLabel) Recommended()    {}
func (AttrPodLabel) Key() string     { return "k8s_pod_label" }
func (a AttrPodLabel) Value() string { return string(a) }

// Deprecated, use `k8s.pod.label` instead
type AttrPodLabels string // k8s.pod.labels

func (AttrPodLabels) Development()    {}
func (AttrPodLabels) Recommended()    {}
func (AttrPodLabels) Key() string     { return "k8s_pod_labels" }
func (a AttrPodLabels) Value() string { return string(a) }

// The name of the Pod
type AttrPodName string // k8s.pod.name

func (AttrPodName) Development()    {}
func (AttrPodName) Recommended()    {}
func (AttrPodName) Key() string     { return "k8s_pod_name" }
func (a AttrPodName) Value() string { return string(a) }

// The UID of the Pod
type AttrPodUid string // k8s.pod.uid

func (AttrPodUid) Development()    {}
func (AttrPodUid) Recommended()    {}
func (AttrPodUid) Key() string     { return "k8s_pod_uid" }
func (a AttrPodUid) Value() string { return string(a) }

// The annotation key-value pairs placed on the ReplicaSet.
//
// The `<key>` being the annotation name, the value being the annotation value, even if the value is empty
type AttrReplicasetAnnotation string // k8s.replicaset.annotation

func (AttrReplicasetAnnotation) Development()    {}
func (AttrReplicasetAnnotation) Recommended()    {}
func (AttrReplicasetAnnotation) Key() string     { return "k8s_replicaset_annotation" }
func (a AttrReplicasetAnnotation) Value() string { return string(a) }

// The label key-value pairs placed on the ReplicaSet.
//
// The `<key>` being the label name, the value being the label value, even if the value is empty
type AttrReplicasetLabel string // k8s.replicaset.label

func (AttrReplicasetLabel) Development()    {}
func (AttrReplicasetLabel) Recommended()    {}
func (AttrReplicasetLabel) Key() string     { return "k8s_replicaset_label" }
func (a AttrReplicasetLabel) Value() string { return string(a) }

// The name of the ReplicaSet
type AttrReplicasetName string // k8s.replicaset.name

func (AttrReplicasetName) Development()    {}
func (AttrReplicasetName) Recommended()    {}
func (AttrReplicasetName) Key() string     { return "k8s_replicaset_name" }
func (a AttrReplicasetName) Value() string { return string(a) }

// The UID of the ReplicaSet
type AttrReplicasetUid string // k8s.replicaset.uid

func (AttrReplicasetUid) Development()    {}
func (AttrReplicasetUid) Recommended()    {}
func (AttrReplicasetUid) Key() string     { return "k8s_replicaset_uid" }
func (a AttrReplicasetUid) Value() string { return string(a) }

// The name of the replication controller
type AttrReplicationcontrollerName string // k8s.replicationcontroller.name

func (AttrReplicationcontrollerName) Development()    {}
func (AttrReplicationcontrollerName) Recommended()    {}
func (AttrReplicationcontrollerName) Key() string     { return "k8s_replicationcontroller_name" }
func (a AttrReplicationcontrollerName) Value() string { return string(a) }

// The UID of the replication controller
type AttrReplicationcontrollerUid string // k8s.replicationcontroller.uid

func (AttrReplicationcontrollerUid) Development()    {}
func (AttrReplicationcontrollerUid) Recommended()    {}
func (AttrReplicationcontrollerUid) Key() string     { return "k8s_replicationcontroller_uid" }
func (a AttrReplicationcontrollerUid) Value() string { return string(a) }

// The name of the resource quota
type AttrResourcequotaName string // k8s.resourcequota.name

func (AttrResourcequotaName) Development()    {}
func (AttrResourcequotaName) Recommended()    {}
func (AttrResourcequotaName) Key() string     { return "k8s_resourcequota_name" }
func (a AttrResourcequotaName) Value() string { return string(a) }

// The UID of the resource quota
type AttrResourcequotaUid string // k8s.resourcequota.uid

func (AttrResourcequotaUid) Development()    {}
func (AttrResourcequotaUid) Recommended()    {}
func (AttrResourcequotaUid) Key() string     { return "k8s_resourcequota_uid" }
func (a AttrResourcequotaUid) Value() string { return string(a) }

// The annotation key-value pairs placed on the StatefulSet.
//
// The `<key>` being the annotation name, the value being the annotation value, even if the value is empty
type AttrStatefulsetAnnotation string // k8s.statefulset.annotation

func (AttrStatefulsetAnnotation) Development()    {}
func (AttrStatefulsetAnnotation) Recommended()    {}
func (AttrStatefulsetAnnotation) Key() string     { return "k8s_statefulset_annotation" }
func (a AttrStatefulsetAnnotation) Value() string { return string(a) }

// The label key-value pairs placed on the StatefulSet.
//
// The `<key>` being the label name, the value being the label value, even if the value is empty
type AttrStatefulsetLabel string // k8s.statefulset.label

func (AttrStatefulsetLabel) Development()    {}
func (AttrStatefulsetLabel) Recommended()    {}
func (AttrStatefulsetLabel) Key() string     { return "k8s_statefulset_label" }
func (a AttrStatefulsetLabel) Value() string { return string(a) }

// The name of the StatefulSet
type AttrStatefulsetName string // k8s.statefulset.name

func (AttrStatefulsetName) Development()    {}
func (AttrStatefulsetName) Recommended()    {}
func (AttrStatefulsetName) Key() string     { return "k8s_statefulset_name" }
func (a AttrStatefulsetName) Value() string { return string(a) }

// The UID of the StatefulSet
type AttrStatefulsetUid string // k8s.statefulset.uid

func (AttrStatefulsetUid) Development()    {}
func (AttrStatefulsetUid) Recommended()    {}
func (AttrStatefulsetUid) Key() string     { return "k8s_statefulset_uid" }
func (a AttrStatefulsetUid) Value() string { return string(a) }

// The name of the K8s volume
type AttrVolumeName string // k8s.volume.name

func (AttrVolumeName) Development()    {}
func (AttrVolumeName) Recommended()    {}
func (AttrVolumeName) Key() string     { return "k8s_volume_name" }
func (a AttrVolumeName) Value() string { return string(a) }

// The type of the K8s volume
type AttrVolumeType string // k8s.volume.type

func (AttrVolumeType) Development()    {}
func (AttrVolumeType) Recommended()    {}
func (AttrVolumeType) Key() string     { return "k8s_volume_type" }
func (a AttrVolumeType) Value() string { return string(a) }

const VolumeTypePersistentVolumeClaim AttrVolumeType = "persistentVolumeClaim"
const VolumeTypeConfigMap AttrVolumeType = "configMap"
const VolumeTypeDownwardApi AttrVolumeType = "downwardAPI"
const VolumeTypeEmptyDir AttrVolumeType = "emptyDir"
const VolumeTypeSecret AttrVolumeType = "secret"
const VolumeTypeLocal AttrVolumeType = "local"

/* State {
    name: "attr.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "ctx": {
            "attributes": [
                {
                    "brief": "The name of the cluster.\n",
                    "examples": [
                        "opentelemetry-cluster",
                    ],
                    "name": "k8s.cluster.name",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "A pseudo-ID for the cluster, set to the UID of the `kube-system` namespace.\n",
                    "examples": [
                        "218fc5a9-a5f1-4b54-aa05-46717d0ab26d",
                    ],
                    "name": "k8s.cluster.uid",
                    "note": "K8s doesn't have support for obtaining a cluster ID. If this is ever\nadded, we will recommend collecting the `k8s.cluster.uid` through the\nofficial APIs. In the meantime, we are able to use the `uid` of the\n`kube-system` namespace as a proxy for cluster ID. Read on for the\nrationale.\n\nEvery object created in a K8s cluster is assigned a distinct UID. The\n`kube-system` namespace is used by Kubernetes itself and will exist\nfor the lifetime of the cluster. Using the `uid` of the `kube-system`\nnamespace is a reasonable proxy for the K8s ClusterID as it will only\nchange if the cluster is rebuilt. Furthermore, Kubernetes UIDs are\nUUIDs as standardized by\n[ISO/IEC 9834-8 and ITU-T X.667](https://www.itu.int/ITU-T/studygroups/com17/oid.html).\nWhich states:\n\n> If generated according to one of the mechanisms defined in Rec.\n> ITU-T X.667 | ISO/IEC 9834-8, a UUID is either guaranteed to be\n> different from all other UUIDs generated before 3603 A.D., or is\n> extremely likely to be different (depending on the mechanism chosen).\n\nTherefore, UIDs between clusters should be extremely unlikely to\nconflict.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The name of the Container from Pod specification, must be unique within a Pod. Container runtime usually uses different globally unique name (`container.name`).\n",
                    "examples": [
                        "redis",
                    ],
                    "name": "k8s.container.name",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Number of times the container was restarted. This attribute can be used to identify a particular container (running or stopped) within a container spec.\n",
                    "name": "k8s.container.restart_count",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "Last terminated reason of the Container.\n",
                    "examples": [
                        "Evicted",
                        "Error",
                    ],
                    "name": "k8s.container.status.last_terminated_reason",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The reason for the container state. Corresponds to the `reason` field of the: [K8s ContainerStateWaiting](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#containerstatewaiting-v1-core) or [K8s ContainerStateTerminated](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#containerstateterminated-v1-core)\n",
                    "examples": [
                        "ContainerCreating",
                        "CrashLoopBackOff",
                        "CreateContainerConfigError",
                        "ErrImagePull",
                        "ImagePullBackOff",
                        "OOMKilled",
                        "Completed",
                        "Error",
                        "ContainerCannotRun",
                    ],
                    "name": "k8s.container.status.reason",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "The container is being created.",
                                "deprecated": none,
                                "id": "container_creating",
                                "note": none,
                                "stability": "development",
                                "value": "ContainerCreating",
                            },
                            {
                                "brief": "The container is in a crash loop back off state.",
                                "deprecated": none,
                                "id": "crash_loop_back_off",
                                "note": none,
                                "stability": "development",
                                "value": "CrashLoopBackOff",
                            },
                            {
                                "brief": "There was an error creating the container configuration.",
                                "deprecated": none,
                                "id": "create_container_config_error",
                                "note": none,
                                "stability": "development",
                                "value": "CreateContainerConfigError",
                            },
                            {
                                "brief": "There was an error pulling the container image.",
                                "deprecated": none,
                                "id": "err_image_pull",
                                "note": none,
                                "stability": "development",
                                "value": "ErrImagePull",
                            },
                            {
                                "brief": "The container image pull is in back off state.",
                                "deprecated": none,
                                "id": "image_pull_back_off",
                                "note": none,
                                "stability": "development",
                                "value": "ImagePullBackOff",
                            },
                            {
                                "brief": "The container was killed due to out of memory.",
                                "deprecated": none,
                                "id": "oom_killed",
                                "note": none,
                                "stability": "development",
                                "value": "OOMKilled",
                            },
                            {
                                "brief": "The container has completed execution.",
                                "deprecated": none,
                                "id": "completed",
                                "note": none,
                                "stability": "development",
                                "value": "Completed",
                            },
                            {
                                "brief": "There was an error with the container.",
                                "deprecated": none,
                                "id": "error",
                                "note": none,
                                "stability": "development",
                                "value": "Error",
                            },
                            {
                                "brief": "The container cannot run.",
                                "deprecated": none,
                                "id": "container_cannot_run",
                                "note": none,
                                "stability": "development",
                                "value": "ContainerCannotRun",
                            },
                        ],
                    },
                },
                {
                    "brief": "The state of the container. [K8s ContainerState](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#containerstate-v1-core)\n",
                    "examples": [
                        "terminated",
                        "running",
                        "waiting",
                    ],
                    "name": "k8s.container.status.state",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "The container has terminated.",
                                "deprecated": none,
                                "id": "terminated",
                                "note": none,
                                "stability": "development",
                                "value": "terminated",
                            },
                            {
                                "brief": "The container is running.",
                                "deprecated": none,
                                "id": "running",
                                "note": none,
                                "stability": "development",
                                "value": "running",
                            },
                            {
                                "brief": "The container is waiting.",
                                "deprecated": none,
                                "id": "waiting",
                                "note": none,
                                "stability": "development",
                                "value": "waiting",
                            },
                        ],
                    },
                },
                {
                    "brief": "The cronjob annotation placed on the CronJob, the `<key>` being the annotation name, the value being the annotation value.\n",
                    "examples": [
                        "4",
                        "",
                    ],
                    "name": "k8s.cronjob.annotation",
                    "note": "Examples:\n\n- An annotation `retries` with value `4` SHOULD be recorded as the\n  `k8s.cronjob.annotation.retries` attribute with value `\"4\"`.\n- An annotation `data` with empty string value SHOULD be recorded as\n  the `k8s.cronjob.annotation.data` attribute with value `\"\"`.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "template[string]",
                },
                {
                    "brief": "The label placed on the CronJob, the `<key>` being the label name, the value being the label value.\n",
                    "examples": [
                        "weekly",
                        "",
                    ],
                    "name": "k8s.cronjob.label",
                    "note": "Examples:\n\n- A label `type` with value `weekly` SHOULD be recorded as the\n  `k8s.cronjob.label.type` attribute with value `\"weekly\"`.\n- A label `automated` with empty string value SHOULD be recorded as\n  the `k8s.cronjob.label.automated` attribute with value `\"\"`.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "template[string]",
                },
                {
                    "brief": "The name of the CronJob.\n",
                    "examples": [
                        "opentelemetry",
                    ],
                    "name": "k8s.cronjob.name",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The UID of the CronJob.\n",
                    "examples": [
                        "275ecb36-5aa8-4c2a-9c47-d8bb681b9aff",
                    ],
                    "name": "k8s.cronjob.uid",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The annotation key-value pairs placed on the DaemonSet.\n",
                    "examples": [
                        "k8s.daemonset.annotation.replicas=1",
                        "k8s.daemonset.annotation.data=",
                    ],
                    "name": "k8s.daemonset.annotation",
                    "note": "The `<key>` being the annotation name, the value being the annotation value, even if the value is empty.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "template[string]",
                },
                {
                    "brief": "The label key-value pairs placed on the DaemonSet.\n",
                    "examples": [
                        "k8s.daemonset.label.app=guestbook",
                        "k8s.daemonset.label.injected=",
                    ],
                    "name": "k8s.daemonset.label",
                    "note": "The `<key>` being the label name, the value being the label value, even if the value is empty.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "template[string]",
                },
                {
                    "brief": "The name of the DaemonSet.\n",
                    "examples": [
                        "opentelemetry",
                    ],
                    "name": "k8s.daemonset.name",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The UID of the DaemonSet.\n",
                    "examples": [
                        "275ecb36-5aa8-4c2a-9c47-d8bb681b9aff",
                    ],
                    "name": "k8s.daemonset.uid",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The annotation key-value pairs placed on the Deployment.\n",
                    "examples": [
                        "k8s.deployment.annotation.replicas=1",
                        "k8s.deployment.annotation.data=",
                    ],
                    "name": "k8s.deployment.annotation",
                    "note": "The `<key>` being the annotation name, the value being the annotation value, even if the value is empty.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "template[string]",
                },
                {
                    "brief": "The label key-value pairs placed on the Deployment.\n",
                    "examples": [
                        "k8s.deployment.label.app=guestbook",
                        "k8s.deployment.label.injected=",
                    ],
                    "name": "k8s.deployment.label",
                    "note": "The `<key>` being the label name, the value being the label value, even if the value is empty.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "template[string]",
                },
                {
                    "brief": "The name of the Deployment.\n",
                    "examples": [
                        "opentelemetry",
                    ],
                    "name": "k8s.deployment.name",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The UID of the Deployment.\n",
                    "examples": [
                        "275ecb36-5aa8-4c2a-9c47-d8bb681b9aff",
                    ],
                    "name": "k8s.deployment.uid",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The name of the horizontal pod autoscaler.\n",
                    "examples": [
                        "opentelemetry",
                    ],
                    "name": "k8s.hpa.name",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The API version of the target resource to scale for the HorizontalPodAutoscaler.\n",
                    "examples": [
                        "apps/v1",
                        "autoscaling/v2",
                    ],
                    "name": "k8s.hpa.scaletargetref.api_version",
                    "note": "This maps to the `apiVersion` field in the `scaleTargetRef` of the HPA spec.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The kind of the target resource to scale for the HorizontalPodAutoscaler.\n",
                    "examples": [
                        "Deployment",
                        "StatefulSet",
                    ],
                    "name": "k8s.hpa.scaletargetref.kind",
                    "note": "This maps to the `kind` field in the `scaleTargetRef` of the HPA spec.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The name of the target resource to scale for the HorizontalPodAutoscaler.\n",
                    "examples": [
                        "my-deployment",
                        "my-statefulset",
                    ],
                    "name": "k8s.hpa.scaletargetref.name",
                    "note": "This maps to the `name` field in the `scaleTargetRef` of the HPA spec.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The UID of the horizontal pod autoscaler.\n",
                    "examples": [
                        "275ecb36-5aa8-4c2a-9c47-d8bb681b9aff",
                    ],
                    "name": "k8s.hpa.uid",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The annotation key-value pairs placed on the Job.\n",
                    "examples": [
                        "k8s.job.annotation.number=1",
                        "k8s.job.annotation.data=",
                    ],
                    "name": "k8s.job.annotation",
                    "note": "The `<key>` being the annotation name, the value being the annotation value, even if the value is empty.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "template[string]",
                },
                {
                    "brief": "The label key-value pairs placed on the Job.\n",
                    "examples": [
                        "k8s.job.label.jobtype=ci",
                        "k8s.job.label.automated=",
                    ],
                    "name": "k8s.job.label",
                    "note": "The `<key>` being the label name, the value being the label value, even if the value is empty.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "template[string]",
                },
                {
                    "brief": "The name of the Job.\n",
                    "examples": [
                        "opentelemetry",
                    ],
                    "name": "k8s.job.name",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The UID of the Job.\n",
                    "examples": [
                        "275ecb36-5aa8-4c2a-9c47-d8bb681b9aff",
                    ],
                    "name": "k8s.job.uid",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The annotation key-value pairs placed on the Namespace.\n",
                    "examples": [
                        "k8s.namespace.annotation.ttl=0",
                        "k8s.namespace.annotation.data=",
                    ],
                    "name": "k8s.namespace.annotation",
                    "note": "The `<key>` being the annotation name, the value being the annotation value, even if the value is empty.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "template[string]",
                },
                {
                    "brief": "The label key-value pairs placed on the Namespace.\n",
                    "examples": [
                        "k8s.namespace.label.kubernetes.io/metadata.name=default",
                        "k8s.namespace.label.data=",
                    ],
                    "name": "k8s.namespace.label",
                    "note": "The `<key>` being the label name, the value being the label value, even if the value is empty.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "template[string]",
                },
                {
                    "brief": "The name of the namespace that the pod is running in.\n",
                    "examples": [
                        "default",
                    ],
                    "name": "k8s.namespace.name",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The phase of the K8s namespace.\n",
                    "examples": [
                        "active",
                        "terminating",
                    ],
                    "name": "k8s.namespace.phase",
                    "note": "This attribute aligns with the `phase` field of the\n[K8s NamespaceStatus](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#namespacestatus-v1-core)\n",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "Active namespace phase as described by [K8s API](https://pkg.go.dev/k8s.io/api@v0.31.3/core/v1#NamespacePhase)",
                                "deprecated": none,
                                "id": "active",
                                "note": none,
                                "stability": "development",
                                "value": "active",
                            },
                            {
                                "brief": "Terminating namespace phase as described by [K8s API](https://pkg.go.dev/k8s.io/api@v0.31.3/core/v1#NamespacePhase)",
                                "deprecated": none,
                                "id": "terminating",
                                "note": none,
                                "stability": "development",
                                "value": "terminating",
                            },
                        ],
                    },
                },
                {
                    "brief": "The annotation placed on the Node, the `<key>` being the annotation name, the value being the annotation value, even if the value is empty.\n",
                    "examples": [
                        "0",
                        "",
                    ],
                    "name": "k8s.node.annotation",
                    "note": "Examples:\n\n- An annotation `node.alpha.kubernetes.io/ttl` with value `0` SHOULD be recorded as\n  the `k8s.node.annotation.node.alpha.kubernetes.io/ttl` attribute with value `\"0\"`.\n- An annotation `data` with empty string value SHOULD be recorded as\n  the `k8s.node.annotation.data` attribute with value `\"\"`.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "template[string]",
                },
                {
                    "brief": "The label placed on the Node, the `<key>` being the label name, the value being the label value, even if the value is empty.\n",
                    "examples": [
                        "arm64",
                        "",
                    ],
                    "name": "k8s.node.label",
                    "note": "Examples:\n\n- A label `kubernetes.io/arch` with value `arm64` SHOULD be recorded\n  as the `k8s.node.label.kubernetes.io/arch` attribute with value `\"arm64\"`.\n- A label `data` with empty string value SHOULD be recorded as\n  the `k8s.node.label.data` attribute with value `\"\"`.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "template[string]",
                },
                {
                    "brief": "The name of the Node.\n",
                    "examples": [
                        "node-1",
                    ],
                    "name": "k8s.node.name",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The UID of the Node.\n",
                    "examples": [
                        "1eb3a0c6-0477-4080-a9cb-0cb7db65c6a2",
                    ],
                    "name": "k8s.node.uid",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The annotation placed on the Pod, the `<key>` being the annotation name, the value being the annotation value.\n",
                    "examples": [
                        "true",
                        "x64",
                        "",
                    ],
                    "name": "k8s.pod.annotation",
                    "note": "Examples:\n\n- An annotation `kubernetes.io/enforce-mountable-secrets` with value `true` SHOULD be recorded as\n  the `k8s.pod.annotation.kubernetes.io/enforce-mountable-secrets` attribute with value `\"true\"`.\n- An annotation `mycompany.io/arch` with value `x64` SHOULD be recorded as\n  the `k8s.pod.annotation.mycompany.io/arch` attribute with value `\"x64\"`.\n- An annotation `data` with empty string value SHOULD be recorded as\n  the `k8s.pod.annotation.data` attribute with value `\"\"`.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "template[string]",
                },
                {
                    "brief": "The label placed on the Pod, the `<key>` being the label name, the value being the label value.\n",
                    "examples": [
                        "my-app",
                        "x64",
                        "",
                    ],
                    "name": "k8s.pod.label",
                    "note": "Examples:\n\n- A label `app` with value `my-app` SHOULD be recorded as\n  the `k8s.pod.label.app` attribute with value `\"my-app\"`.\n- A label `mycompany.io/arch` with value `x64` SHOULD be recorded as\n  the `k8s.pod.label.mycompany.io/arch` attribute with value `\"x64\"`.\n- A label `data` with empty string value SHOULD be recorded as\n  the `k8s.pod.label.data` attribute with value `\"\"`.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "template[string]",
                },
                {
                    "brief": "Deprecated, use `k8s.pod.label` instead.",
                    "deprecated": {
                        "note": "Replaced by `k8s.pod.label`.",
                        "reason": "renamed",
                        "renamed_to": "k8s.pod.label",
                    },
                    "examples": [
                        "my-app",
                    ],
                    "name": "k8s.pod.labels",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "template[string]",
                },
                {
                    "brief": "The name of the Pod.\n",
                    "examples": [
                        "opentelemetry-pod-autoconf",
                    ],
                    "name": "k8s.pod.name",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The UID of the Pod.\n",
                    "examples": [
                        "275ecb36-5aa8-4c2a-9c47-d8bb681b9aff",
                    ],
                    "name": "k8s.pod.uid",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The annotation key-value pairs placed on the ReplicaSet.\n",
                    "examples": [
                        "k8s.replicaset.annotation.replicas=0",
                        "k8s.replicaset.annotation.data=",
                    ],
                    "name": "k8s.replicaset.annotation",
                    "note": "The `<key>` being the annotation name, the value being the annotation value, even if the value is empty.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "template[string]",
                },
                {
                    "brief": "The label key-value pairs placed on the ReplicaSet.\n",
                    "examples": [
                        "k8s.replicaset.label.app=guestbook",
                        "k8s.replicaset.label.injected=",
                    ],
                    "name": "k8s.replicaset.label",
                    "note": "The `<key>` being the label name, the value being the label value, even if the value is empty.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "template[string]",
                },
                {
                    "brief": "The name of the ReplicaSet.\n",
                    "examples": [
                        "opentelemetry",
                    ],
                    "name": "k8s.replicaset.name",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The UID of the ReplicaSet.\n",
                    "examples": [
                        "275ecb36-5aa8-4c2a-9c47-d8bb681b9aff",
                    ],
                    "name": "k8s.replicaset.uid",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The name of the replication controller.\n",
                    "examples": [
                        "opentelemetry",
                    ],
                    "name": "k8s.replicationcontroller.name",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The UID of the replication controller.\n",
                    "examples": [
                        "275ecb36-5aa8-4c2a-9c47-d8bb681b9aff",
                    ],
                    "name": "k8s.replicationcontroller.uid",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The name of the resource quota.\n",
                    "examples": [
                        "opentelemetry",
                    ],
                    "name": "k8s.resourcequota.name",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The UID of the resource quota.\n",
                    "examples": [
                        "275ecb36-5aa8-4c2a-9c47-d8bb681b9aff",
                    ],
                    "name": "k8s.resourcequota.uid",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The annotation key-value pairs placed on the StatefulSet.\n",
                    "examples": [
                        "k8s.statefulset.annotation.replicas=1",
                        "k8s.statefulset.annotation.data=",
                    ],
                    "name": "k8s.statefulset.annotation",
                    "note": "The `<key>` being the annotation name, the value being the annotation value, even if the value is empty.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "template[string]",
                },
                {
                    "brief": "The label key-value pairs placed on the StatefulSet.\n",
                    "examples": [
                        "k8s.statefulset.label.app=guestbook",
                        "k8s.statefulset.label.injected=",
                    ],
                    "name": "k8s.statefulset.label",
                    "note": "The `<key>` being the label name, the value being the label value, even if the value is empty.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "template[string]",
                },
                {
                    "brief": "The name of the StatefulSet.\n",
                    "examples": [
                        "opentelemetry",
                    ],
                    "name": "k8s.statefulset.name",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The UID of the StatefulSet.\n",
                    "examples": [
                        "275ecb36-5aa8-4c2a-9c47-d8bb681b9aff",
                    ],
                    "name": "k8s.statefulset.uid",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The name of the K8s volume.\n",
                    "examples": [
                        "volume0",
                    ],
                    "name": "k8s.volume.name",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The type of the K8s volume.\n",
                    "examples": [
                        "emptyDir",
                        "persistentVolumeClaim",
                    ],
                    "name": "k8s.volume.type",
                    "requirement_level": "recommended",
                    "root_namespace": "k8s",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "A [persistentVolumeClaim](https://v1-30.docs.kubernetes.io/docs/concepts/storage/volumes/#persistentvolumeclaim) volume",
                                "deprecated": none,
                                "id": "persistent_volume_claim",
                                "note": none,
                                "stability": "development",
                                "value": "persistentVolumeClaim",
                            },
                            {
                                "brief": "A [configMap](https://v1-30.docs.kubernetes.io/docs/concepts/storage/volumes/#configmap) volume",
                                "deprecated": none,
                                "id": "config_map",
                                "note": none,
                                "stability": "development",
                                "value": "configMap",
                            },
                            {
                                "brief": "A [downwardAPI](https://v1-30.docs.kubernetes.io/docs/concepts/storage/volumes/#downwardapi) volume",
                                "deprecated": none,
                                "id": "downward_api",
                                "note": none,
                                "stability": "development",
                                "value": "downwardAPI",
                            },
                            {
                                "brief": "An [emptyDir](https://v1-30.docs.kubernetes.io/docs/concepts/storage/volumes/#emptydir) volume",
                                "deprecated": none,
                                "id": "empty_dir",
                                "note": none,
                                "stability": "development",
                                "value": "emptyDir",
                            },
                            {
                                "brief": "A [secret](https://v1-30.docs.kubernetes.io/docs/concepts/storage/volumes/#secret) volume",
                                "deprecated": none,
                                "id": "secret",
                                "note": none,
                                "stability": "development",
                                "value": "secret",
                            },
                            {
                                "brief": "A [local](https://v1-30.docs.kubernetes.io/docs/concepts/storage/volumes/#local) volume",
                                "deprecated": none,
                                "id": "local",
                                "note": none,
                                "stability": "development",
                                "value": "local",
                            },
                        ],
                    },
                },
            ],
            "root_namespace": "k8s",
        },
    },
    env: Environment {
        globals: {
            "concat_if": weaver_forge::extensions::util::concat_if,
            "cycler": minijinja_contrib::globals::cycler,
            "debug": minijinja::functions::builtins::debug,
            "dict": minijinja::functions::builtins::dict,
            "joiner": minijinja_contrib::globals::joiner,
            "namespace": minijinja::functions::builtins::namespace,
            "params": {
                "params": {},
            },
            "range": minijinja::functions::builtins::range,
            "template": {},
        },
        tests: [
            "!=",
            "<",
            "<=",
            "==",
            ">",
            ">=",
            "array",
            "boolean",
            "defined",
            "deprecated",
            "divisibleby",
            "endingwith",
            "enum",
            "enum_type",
            "eq",
            "equalto",
            "escaped",
            "even",
            "experimental",
            "false",
            "filter",
            "float",
            "ge",
            "greaterthan",
            "gt",
            "in",
            "int",
            "integer",
            "iterable",
            "le",
            "lessthan",
            "lower",
            "lt",
            "mapping",
            "ne",
            "none",
            "number",
            "odd",
            "safe",
            "sameas",
            "sequence",
            "simple_type",
            "stable",
            "startingwith",
            "string",
            "template_type",
            "test",
            "true",
            "undefined",
            "upper",
        ],
        filters: [
            "abs",
            "acronym",
            "ansi_bg_black",
            "ansi_bg_blue",
            "ansi_bg_bright_black",
            "ansi_bg_bright_blue",
            "ansi_bg_bright_cyan",
            "ansi_bg_bright_green",
            "ansi_bg_bright_magenta",
            "ansi_bg_bright_red",
            "ansi_bg_bright_white",
            "ansi_bg_bright_yellow",
            "ansi_bg_cyan",
            "ansi_bg_green",
            "ansi_bg_magenta",
            "ansi_bg_red",
            "ansi_bg_white",
            "ansi_bg_yellow",
            "ansi_black",
            "ansi_blue",
            "ansi_bold",
            "ansi_bright_black",
            "ansi_bright_blue",
            "ansi_bright_cyan",
            "ansi_bright_green",
            "ansi_bright_magenta",
            "ansi_bright_red",
            "ansi_bright_white",
            "ansi_bright_yellow",
            "ansi_cyan",
            "ansi_green",
            "ansi_italic",
            "ansi_magenta",
            "ansi_red",
            "ansi_strikethrough",
            "ansi_underline",
            "ansi_white",
            "ansi_yellow",
            "attr",
            "attribute_id",
            "attribute_namespace",
            "attribute_registry_file",
            "attribute_registry_namespace",
            "attribute_registry_title",
            "attribute_sort",
            "batch",
            "body_fields",
            "bool",
            "camel_case",
            "camel_case_const",
            "capitalize",
            "capitalize_first",
            "comment",
            "comment_with_prefix",
            "count",
            "d",
            "default",
            "dictsort",
            "e",
            "enum_type",
            "escape",
            "filesizeformat",
            "first",
            "flatten",
            "float",
            "groupby",
            "indent",
            "instantiated_type",
            "int",
            "items",
            "join",
            "kebab_case",
            "kebab_case_const",
            "last",
            "length",
            "lines",
            "list",
            "lower",
            "lower_case",
            "map",
            "map_text",
            "markdown_to_html",
            "max",
            "metric_namespace",
            "min",
            "not_required",
            "pascal_case",
            "pascal_case_const",
            "pluralize",
            "pprint",
            "print_member_value",
            "regex_replace",
            "reject",
            "rejectattr",
            "replace",
            "required",
            "reverse",
            "round",
            "safe",
            "screaming_kebab_case",
            "screaming_snake_case",
            "screaming_snake_case_const",
            "select",
            "selectattr",
            "slice",
            "snake_case",
            "snake_case_const",
            "sort",
            "split",
            "split_id",
            "string",
            "striptags",
            "sum",
            "title",
            "title_case",
            "tojson",
            "toyaml",
            "trim",
            "truncate",
            "type_mapping",
            "unique",
            "upper",
            "upper_case",
            "urlencode",
        ],
        templates: [
            "attr.go.j2",
        ],
    },
} */
