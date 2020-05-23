# API Docs
This Document documents the types introduced by the FoundationDB Operator to be consumed by users.
> Note this document is generated from code comments. When contributing a change to this document please do so by changing the code comments.

## Table of Contents
* [BackupGenerationStatus](#backupgenerationstatus)
* [ClusterGenerationStatus](#clustergenerationstatus)
* [ClusterHealth](#clusterhealth)
* [ConnectionString](#connectionstring)
* [ContainerOverrides](#containeroverrides)
* [DataCenter](#datacenter)
* [DatabaseConfiguration](#databaseconfiguration)
* [FdbVersion](#fdbversion)
* [FoundationDBBackup](#foundationdbbackup)
* [FoundationDBBackupList](#foundationdbbackuplist)
* [FoundationDBBackupSpec](#foundationdbbackupspec)
* [FoundationDBBackupStatus](#foundationdbbackupstatus)
* [FoundationDBBackupStatusBackupDetails](#foundationdbbackupstatusbackupdetails)
* [FoundationDBCluster](#foundationdbcluster)
* [FoundationDBClusterAutomationOptions](#foundationdbclusterautomationoptions)
* [FoundationDBClusterFaultDomain](#foundationdbclusterfaultdomain)
* [FoundationDBClusterList](#foundationdbclusterlist)
* [FoundationDBClusterSpec](#foundationdbclusterspec)
* [FoundationDBClusterStatus](#foundationdbclusterstatus)
* [FoundationDBLiveBackupStatus](#foundationdblivebackupstatus)
* [FoundationDBLiveBackupStatusState](#foundationdblivebackupstatusstate)
* [FoundationDBRestore](#foundationdbrestore)
* [FoundationDBRestoreList](#foundationdbrestorelist)
* [FoundationDBRestoreSpec](#foundationdbrestorespec)
* [FoundationDBRestoreStatus](#foundationdbrestorestatus)
* [FoundationDBStatus](#foundationdbstatus)
* [FoundationDBStatusBackupInfo](#foundationdbstatusbackupinfo)
* [FoundationDBStatusBackupTag](#foundationdbstatusbackuptag)
* [FoundationDBStatusClientDBStatus](#foundationdbstatusclientdbstatus)
* [FoundationDBStatusClusterClientInfo](#foundationdbstatusclusterclientinfo)
* [FoundationDBStatusClusterInfo](#foundationdbstatusclusterinfo)
* [FoundationDBStatusConnectedClient](#foundationdbstatusconnectedclient)
* [FoundationDBStatusCoordinator](#foundationdbstatuscoordinator)
* [FoundationDBStatusCoordinatorInfo](#foundationdbstatuscoordinatorinfo)
* [FoundationDBStatusDataStatistics](#foundationdbstatusdatastatistics)
* [FoundationDBStatusLayerInfo](#foundationdbstatuslayerinfo)
* [FoundationDBStatusLocalClientInfo](#foundationdbstatuslocalclientinfo)
* [FoundationDBStatusMovingData](#foundationdbstatusmovingdata)
* [FoundationDBStatusProcessInfo](#foundationdbstatusprocessinfo)
* [FoundationDBStatusSupportedVersion](#foundationdbstatussupportedversion)
* [ProcessAddress](#processaddress)
* [ProcessCounts](#processcounts)
* [Region](#region)
* [RequiredAddressSet](#requiredaddressset)
* [RoleCounts](#rolecounts)
* [VersionFlags](#versionflags)

## BackupGenerationStatus

BackupGenerationStatus stores information on which generations have reached different stages in reconciliation for the backup.

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| reconciled | Reconciled provides the last generation that was fully reconciled. | int64 | false |
| needsBackupAgentUpdate | NeedsBackupAgentUpdate provides the last generation that could not complete reconciliation because the backup agent deployment needs to be updated. | int64 | false |
| needsBackupStart | NeedsBackupStart provides the last generation that could not complete reconciliation because we need to start a backup. | int64 | false |
| needsBackupStop | NeedsBackupStart provides the last generation that could not complete reconciliation because we need to stop a backup. | int64 | false |
| needsBackupPauseToggle | NeedsBackupPauseToggle provides the last generation that needs to have a backup paused or resumed. | int64 | false |
| needsBackupModification | NeedsBackupReconfiguration provides the last generation that could not complete reconciliation because we need to modify backup parameters. | int64 | false |

[Back to TOC](#table-of-contents)

## ClusterGenerationStatus

ClusterGenerationStatus stores information on which generations have reached different stages in reconciliation for the cluster.

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| reconciled | Reconciled provides the last generation that was fully reconciled. | int64 | false |
| needsConfigurationChange | NeedsConfigurationChange provides the last generation that is pending a change to configuration. | int64 | false |
| needsBounce | NeedsBounce provides the last generation that is pending a bounce of fdbserver. | int64 | false |
| needsPodDeletion | NeedsPodDeletion provides the last generation that is pending pods being deleted and recreated. | int64 | false |
| needsShrink | NeedsShrink provides the last generation that is pending pods being excluded and removed. | int64 | false |
| needsGrow | NeedsGrow provides the last generation that is pending pods being added. | int64 | false |
| needsMonitorConfUpdate | NeedsMonitorConfUpdate provides the last generation that needs an update through the fdbmonitor conf. | int64 | false |
| missingDatabaseStatus | DatabaseUnavailable provides the last generation that could not complete reconciliation due to the database being unavailable. | int64 | false |
| hasExtraListeners | HasExtraListeners provides the last generation that could not complete reconciliation because it has more listeners than it is supposed to. | int64 | false |
| needsBackupAgentUpdate | NeedsBackupAgentUpdate provides the last generation that could not complete reconciliation because the backup agent deployment needs to be updated. **Deprecated: This needs to get moved into FoundationDBBackup** | int64 | false |

[Back to TOC](#table-of-contents)

## ClusterHealth

ClusterHealth represents different views into health in the cluster status.

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| available | Available reports whether the database is accepting reads and writes. | bool | false |
| healthy | Healthy reports whether the database is in a fully healthy state. | bool | false |
| fullReplication | FullReplication reports whether all data are fully replicated according to the current replication policy. | bool | false |
| dataMovementPriority | DataMovementPriority reports the priority of the highest-priority data movement in the cluster. | int | false |

[Back to TOC](#table-of-contents)

## ConnectionString

ConnectionString models the contents of a cluster file in a structured way

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| DatabaseName | DatabaseName provides an identifier for the database which persists across coordinator changes. | string | false |
| GenerationID | GenerationID provides a unique ID for the current generation of coordinators. | string | false |
| Coordinators | Coordinators provides the addresses of the current coordinators. | []string | false |

[Back to TOC](#table-of-contents)

## ContainerOverrides

ContainerOverrides provides options for customizing a container created by the operator.

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| enableTls | EnableTLS controls whether we should be listening on a TLS connection. | bool | false |
| peerVerificationRules | PeerVerificationRules provides the rules for what client certificates the process should accept. | string | false |
| env | Env provides environment variables.  **Deprecated: Use the PodTemplate field instead.** | [][corev1.EnvVar](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.17/#envvar-v1-core) | false |
| volumeMounts | VolumeMounts provides volume mounts.  **Deprecated: Use the PodTemplate field instead.** | [][corev1.VolumeMount](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.17/#volumemount-v1-core) | false |
| imageName | ImageName provides the name of the image to use for the container, without the version tag.  **Deprecated: Use the PodTemplate field instead.** | string | false |
| securityContext | SecurityContext provides the container's security context.  **Deprecated: Use the PodTemplate field instead.** | *[corev1.SecurityContext](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.17/#securitycontext-v1-core) | false |

[Back to TOC](#table-of-contents)

## DataCenter

DataCenter represents a data center in the region configuration

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| id | The ID of the data center. This must match the dcid locality field. | string | false |
| priority | The priority of this data center when we have to choose a location. Higher priorities are preferred over lower priorities. | int | false |
| satellite | Satellite indicates whether the data center is serving as a satellite for the region. A value of 1 indicates that it is a satellite, and a value of 0 indicates that it is not a satellite. | int | false |

[Back to TOC](#table-of-contents)

## DatabaseConfiguration

DatabaseConfiguration represents the configuration of the database

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| redundancy_mode | RedundancyMode defines the core replication factor for the database. | string | false |
| storage_engine | StorageEngine defines the storage engine the database uses. | string | false |
| usable_regions | UsableRegions defines how many regions the database should store data in. | int | false |
| regions | Regions defines the regions that the database can replicate in. | [][Region](#region) | false |
| RoleCounts | RoleCounts defines how many processes the database should recruit for each role. | [RoleCounts](#rolecounts) | true |
| VersionFlags | VersionFlags defines internal flags for testing new features in the database. | [VersionFlags](#versionflags) | true |

[Back to TOC](#table-of-contents)

## FdbVersion

FdbVersion represents a version of FoundationDB.  This provides convenience methods for checking features available in different versions.

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| Major | Major is the major version | int | false |
| Minor | Minor is the minor version | int | false |
| Patch | Patch is the patch version | int | false |

[Back to TOC](#table-of-contents)

## FoundationDBBackup

FoundationDBBackup is the Schema for the FoundationDB Backup API

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| metadata |  | [metav1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.17/#objectmeta-v1-meta) | false |
| spec |  | [FoundationDBBackupSpec](#foundationdbbackupspec) | false |
| status |  | [FoundationDBBackupStatus](#foundationdbbackupstatus) | false |

[Back to TOC](#table-of-contents)

## FoundationDBBackupList

FoundationDBBackupList contains a list of FoundationDBBackup

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| metadata |  | [metav1.ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.17/#listmeta-v1-meta) | false |
| items |  | [][FoundationDBBackup](#foundationdbbackup) | true |

[Back to TOC](#table-of-contents)

## FoundationDBBackupSpec

FoundationDBBackupSpec describes the desired state of the backup for a cluster.

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| version | The version of FoundationDB that the backup agents should run. | string | true |
| clusterName | The cluster this backup is for. | string | true |
| backupState | The desired state of the backup. The default is Running. | string | false |
| backupName | The name for the backup. The default is to use the name from the backup metadata. | string | false |
| accountName | The account name to use with the backup destination. | string | true |
| bucket | The backup bucket to write to. The default is to use \"fdb-backups\". | string | false |
| agentCount | AgentCount defines the number of backup agents to run. The default is run 2 agents. | *int | false |
| snapshotPeriodSeconds | The time window between new snapshots. This is measured in seconds. The default is 864,000, or 10 days. | *int | false |
| podTemplateSpec | PodTemplateSpec allows customizing the pod template for the backup agents. | *[corev1.PodTemplateSpec](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.17/#podtemplatespec-v1-core) | false |

[Back to TOC](#table-of-contents)

## FoundationDBBackupStatus

FoundationDBBackupStatus describes the current status of the backup for a cluster.

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| agentCount | AgentCount provides the number of agents that are up-to-date, ready, and not terminated. | int | false |
| deploymentConfigured | DeploymentConfigured indicates whether the deployment is correctly configured. | bool | false |
| backupDetails | BackupDetails provides information about the state of the backup in the cluster. | *[FoundationDBBackupStatusBackupDetails](#foundationdbbackupstatusbackupdetails) | false |
| generations | Generations provides information about the latest generation to be reconciled, or to reach other stages in reconciliation. | [BackupGenerationStatus](#backupgenerationstatus) | false |

[Back to TOC](#table-of-contents)

## FoundationDBBackupStatusBackupDetails

FoundationDBBackupStatusBackupDetails provides information about the state of the backup in the cluster.

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| url |  | string | false |
| running |  | bool | false |
| paused |  | bool | false |
| snapshotTime |  | int | false |

[Back to TOC](#table-of-contents)

## FoundationDBCluster

FoundationDBCluster is the Schema for the foundationdbclusters API

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| metadata |  | [metav1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.17/#objectmeta-v1-meta) | false |
| spec |  | [FoundationDBClusterSpec](#foundationdbclusterspec) | false |
| status |  | [FoundationDBClusterStatus](#foundationdbclusterstatus) | false |

[Back to TOC](#table-of-contents)

## FoundationDBClusterAutomationOptions

FoundationDBClusterAutomationOptions provides flags for enabling or disabling operations that can be performed on a cluster.

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| configureDatabase | ConfigureDatabase defines whether the operator is allowed to reconfigure the database. | *bool | false |
| killProcesses | KillProcesses defines whether the operator is allowed to bounce fdbserver processes. | *bool | false |
| deletePods | DeletePods defines whether the operator is allowed to delete pods in order to recreate them. | *bool | false |

[Back to TOC](#table-of-contents)

## FoundationDBClusterFaultDomain

FoundationDBClusterFaultDomain describes the fault domain that a cluster is replicated across.

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| key | Key provides a topology key for the fault domain to replicate across. | string | false |
| value | Value provides a harcoded value to use for the zoneid for the pods. | string | false |
| valueFrom | ValueFrom provides a field selector to use as the source of the fault domain. | string | false |
| zoneCount | ZoneCount provides the number of fault domains in the data center where these processes are running. This is only used in the `kubernetes-cluster` fault domain strategy. | int | false |
| zoneIndex | ZoneIndex provides the index of this Kubernetes cluster in the list of KCs in the data center. This is only used in the `kubernetes-cluster` fault domain strategy. | int | false |

[Back to TOC](#table-of-contents)

## FoundationDBClusterList

FoundationDBClusterList contains a list of FoundationDBCluster

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| metadata |  | [metav1.ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.17/#listmeta-v1-meta) | false |
| items |  | [][FoundationDBCluster](#foundationdbcluster) | true |

[Back to TOC](#table-of-contents)

## FoundationDBClusterSpec

FoundationDBClusterSpec defines the desired state of a cluster.

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| version | Version defines the version of FoundationDB the cluster should run. | string | true |
| sidecarVersions | SidecarVersions defines the build version of the sidecar to run. This maps an FDB version to the corresponding sidecar build version. | map[string]int | false |
| runningVersion | RunningVersion defines the version of FoundationDB that the cluster is currently running. | string | false |
| databaseConfiguration | DatabaseConfiguration defines the database configuration. | [DatabaseConfiguration](#databaseconfiguration) | false |
| configured | Configured defines whether we have configured the database yet. | bool | false |
| processCounts | ProcessCounts defines the number of processes to configure for each process class. You can generally omit this, to allow the operator to infer the process counts based on the database configuration. | [ProcessCounts](#processcounts) | false |
| connectionString | ConnectionString defines the contents of the cluster file. | string | false |
| faultDomain | FaultDomain defines the rules for what fault domain to replicate across. | [FoundationDBClusterFaultDomain](#foundationdbclusterfaultdomain) | false |
| customParameters | CustomParameters defines additional parameters to pass to the fdbserver processes. | []string | false |
| instancesToRemove | InstancesToRemove defines the instances that we should remove from the cluster. This list contains the instance IDs. | []string | false |
| pendingRemovals | PendingRemovals defines the processes that are pending removal. This maps the name of a pod to its IP address. If a value is left blank, the controller will provide the pod's current IP.  **Deprecated: This is for internal use only. To tell the operator to remove or replace a process, use InstancesToRemove.** | map[string]string | false |
| podTemplate | PodTemplate allows customizing the FoundationDB pods. | *[corev1.PodTemplateSpec](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.17/#podtemplatespec-v1-core) | false |
| volumeClaim | VolumeClaim allows customizing the persistent volume claim for the FoundationDB pods. | *[corev1.PersistentVolumeClaim](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.17/#persistentvolumeclaim-v1-core) | false |
| configMap | ConfigMap allows customizing the config map the operator creates. | *[corev1.ConfigMap](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.17/#configmap-v1-core) | false |
| mainContainer | MainContainer defines customization for the foundationdb container. | [ContainerOverrides](#containeroverrides) | false |
| sidecarContainer | SidecarContainer defines customization for the foundationdb-kubernetes-sidecar container. | [ContainerOverrides](#containeroverrides) | false |
| trustedCAs | TrustedCAs defines a list of root CAs the cluster should trust, in PEM format. | []string | false |
| sidecarVariables | SidecarVariables defines Ccustom variables that the sidecar should make available for substitution in the monitor conf file. | []string | false |
| logGroup | LogGroup defines the log group to use for the trace logs for the cluster. | string | false |
| dataCenter | DataCenter defines the data center where these processes are running. | string | false |
| dataHall | DataHall defines the data hall where these processes are running. | string | false |
| automationOptions | AutomationOptions defines customization for enabling or disabling certain operations in the operator. | [FoundationDBClusterAutomationOptions](#foundationdbclusterautomationoptions) | false |
| instanceIDPrefix | InstanceIDPrefix defines a prefix to append to the instance IDs in the locality fields. | string | false |
| updatePodsByReplacement | UpdatePodsByReplacement determines whether we should update pod config by replacing the pods rather than deleting them. | bool | false |
| sidecarVersion | SidecarVersion defines the build version of the sidecar to use.  **Deprecated: Use SidecarVersions instead.** | int | false |
| podLabels | PodLabels defines custom labels to apply to the FDB pods.  **Deprecated: Use the PodTemplate field instead.** | map[string]string | false |
| resources | Resources defines the resource requirements for the foundationdb containers.  **Deprecated: Use the PodTemplate field instead.** | *[corev1.ResourceRequirements](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.17/#resourcerequirements-v1-core) | false |
| initContainers | InitContainers defines custom init containers for the FDB pods.  **Deprecated: Use the PodTemplate field instead.** | [][corev1.Container](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.17/#container-v1-core) | false |
| containers | Containers defines custom containers for the FDB pods.  **Deprecated: Use the PodTemplate field instead.** | [][corev1.Container](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.17/#container-v1-core) | false |
| volumes | Volumes defines custom volumes for the FDB pods.  **Deprecated: Use the PodTemplate field instead.** | []corev1.Volume | false |
| podSecurityContext | PodSecurityContext defines the security context to apply to the FDB pods.  **Deprecated: Use the PodTemplate field instead.** | *[corev1.PodSecurityContext](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.17/#podsecuritycontext-v1-core) | false |
| automountServiceAccountToken | AutomountServiceAccountToken defines whether we should automount the service account tokens in the FDB pods.  **Deprecated: Use the PodTemplate field instead.** | *bool | false |
| nextInstanceID | NextInstanceID defines the ID to use when creating the next instance.  **Deprecated: This is no longer used.** | int | false |
| storageClass | StorageClass defines the storage class for the volumes in the cluster.  **Deprecated: Use the VolumeClaim field instead.** | *string | false |
| volumeSize | VolumeSize defines the size of the volume to use for stateful processes.  **Deprecated: Use the VolumeClaim field instead.** | string | false |

[Back to TOC](#table-of-contents)

## FoundationDBClusterStatus

FoundationDBClusterStatus defines the observed state of FoundationDBCluster

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| processCounts | ProcessCounts defines the number of processes that are currently running in the cluster. | [ProcessCounts](#processcounts) | false |
| incorrectProcesses | IncorrectProcesses provides the processes that do not have the correct configuration.  This will map the instance ID to the timestamp when we observed the incorrect configuration. | map[string]int64 | false |
| incorrectPods | IncorrectPods provides the pods that do not have the correct spec.  This will contain the name of the pod. | []string | false |
| missingProcesses | MissingProcesses provides the processes that are not reporting to the cluster. This will map the names of the pod to the timestamp when we observed that the process was missing. | map[string]int64 | false |
| databaseConfiguration | DatabaseConfiguration provides the running configuration of the database. | [DatabaseConfiguration](#databaseconfiguration) | false |
| generations | Generations provides information about the latest generation to be reconciled, or to reach other stages at which reconciliation can halt. | [ClusterGenerationStatus](#clustergenerationstatus) | false |
| health | Health provides information about the health of the database. | [ClusterHealth](#clusterhealth) | false |
| requiredAddresses | RequiredAddresses define that addresses that we need to enable for the processes in the cluster. | [RequiredAddressSet](#requiredaddressset) | false |
| hasIncorrectConfigMap | HasIncorrectConfigMap indicates whether the latest config map is out of date with the cluster spec. | bool | false |

[Back to TOC](#table-of-contents)

## FoundationDBLiveBackupStatus

FoundationDBLiveBackupStatus describes the live status of the backup for a cluster, as provided by the backup status command.

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| DestinationURL | DestinationURL provides the URL that the backup is being written to. | string | false |
| SnapshotIntervalSeconds | SnapshotIntervalSeconds provides the interval of the snapshots. | int | false |
| Status | Status provides the current state of the backup. | [FoundationDBLiveBackupStatusState](#foundationdblivebackupstatusstate) | false |
| BackupAgentsPaused | BackupAgentsPaused describes whether the backup agents are paused. | bool | false |

[Back to TOC](#table-of-contents)

## FoundationDBLiveBackupStatusState

FoundationDBLiveBackupStatusState provides the state of a backup in the backup status.

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| Running | Running determines whether the backup is currently running. | bool | false |

[Back to TOC](#table-of-contents)

## FoundationDBRestore

FoundationDBRestore is the Schema for the FoundationDB Restore API

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| metadata |  | [metav1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.17/#objectmeta-v1-meta) | false |
| spec |  | [FoundationDBRestoreSpec](#foundationdbrestorespec) | false |
| status |  | [FoundationDBRestoreStatus](#foundationdbrestorestatus) | false |

[Back to TOC](#table-of-contents)

## FoundationDBRestoreList

FoundationDBRestoreList contains a list of FoundationDBRestore objects.

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| metadata |  | [metav1.ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.17/#listmeta-v1-meta) | false |
| items |  | [][FoundationDBRestore](#foundationdbrestore) | true |

[Back to TOC](#table-of-contents)

## FoundationDBRestoreSpec

FoundationDBRestoreSpec describes the desired state of the backup for a cluster.

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| destinationClusterName | DestinationClusterName provides the name of the cluster that the data is being restored into. | string | true |
| backupURL | BackupURL provides the URL for the backup. | string | true |

[Back to TOC](#table-of-contents)

## FoundationDBRestoreStatus

FoundationDBRestoreStatus describes the current status of the restore for a cluster.

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| running | Running describes whether the restore is currently running. | bool | false |

[Back to TOC](#table-of-contents)

## FoundationDBStatus

FoundationDBStatus describes the status of the cluster as provided by FoundationDB itself.

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| client | Client provides the client section of the status. | [FoundationDBStatusLocalClientInfo](#foundationdbstatuslocalclientinfo) | false |
| cluster | Cluster provides the cluster section of the status. | [FoundationDBStatusClusterInfo](#foundationdbstatusclusterinfo) | false |

[Back to TOC](#table-of-contents)

## FoundationDBStatusBackupInfo

FoundationDBStatusBackupInfo provides information about backups that have been started.

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| paused | Paused tells whether the backups are paused. | bool | false |
| tags | Tags provides information about specific backups. | map[string][FoundationDBStatusBackupTag](#foundationdbstatusbackuptag) | false |

[Back to TOC](#table-of-contents)

## FoundationDBStatusBackupTag

FoundationDBStatusBackupTag provides information about a backup under a tag in the cluster status.

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| current_container |  | string | false |
| running_backup |  | bool | false |
| running_backup_is_restorable |  | bool | false |

[Back to TOC](#table-of-contents)

## FoundationDBStatusClientDBStatus

FoundationDBStatusClientDBStatus represents the databaseStatus field in the JSON database status

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| available | Available indicates whether the database is accepting traffic. | bool | false |
| healthy | Healthy indicates whether the database is fully healthy. | bool | false |

[Back to TOC](#table-of-contents)

## FoundationDBStatusClusterClientInfo

FoundationDBStatusClusterClientInfo represents the connected client details in the cluster status.

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| count | Count provides the number of clients connected to the database. | int | false |
| supported_versions | SupportedVersions provides information about the versions supported by the connected clients. | [][FoundationDBStatusSupportedVersion](#foundationdbstatussupportedversion) | false |

[Back to TOC](#table-of-contents)

## FoundationDBStatusClusterInfo

FoundationDBStatusClusterInfo describes the \"cluster\" portion of the cluster status

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| configuration | DatabaseConfiguration describes the current configuration of the database. | [DatabaseConfiguration](#databaseconfiguration) | false |
| processes | Processes provides details on the processes that are reporting to the cluster. | map[string][FoundationDBStatusProcessInfo](#foundationdbstatusprocessinfo) | false |
| data | Data provides information about the data in the database. | [FoundationDBStatusDataStatistics](#foundationdbstatusdatastatistics) | false |
| full_replication | FullReplication indicates whether the database is fully replicated. | bool | false |
| clients | Clients provides information about clients that are connected to the database. | [FoundationDBStatusClusterClientInfo](#foundationdbstatusclusterclientinfo) | false |
| layers | Layers provides information about layers that are running against the cluster. | [FoundationDBStatusLayerInfo](#foundationdbstatuslayerinfo) | false |

[Back to TOC](#table-of-contents)

## FoundationDBStatusConnectedClient

FoundationDBStatusConnectedClient provides information about a client that is connected to the database.

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| address | Address provides the address the client is connecting from. | string | false |
| log_group | LogGroup provides the trace log group the client has set. | string | false |

[Back to TOC](#table-of-contents)

## FoundationDBStatusCoordinator

FoundationDBStatusCoordinator contains information about one of the coordinators.

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| address | Address provides the coordinator's address. | string | false |
| reachable | Reachable indicates whether the coordinator is reachable. | bool | false |

[Back to TOC](#table-of-contents)

## FoundationDBStatusCoordinatorInfo

FoundationDBStatusCoordinatorInfo contains information about the client's connection to the coordinators.

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| coordinators | Coordinators provides a list with coordinator details. | [][FoundationDBStatusCoordinator](#foundationdbstatuscoordinator) | false |

[Back to TOC](#table-of-contents)

## FoundationDBStatusDataStatistics

FoundationDBStatusDataStatistics provides information about the data in the database

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| total_kv_size_bytes | KVBytes provides the total Key Value Bytes in the database. | int | false |
| moving_data | MovingData provides information about the current data movement. | [FoundationDBStatusMovingData](#foundationdbstatusmovingdata) | false |

[Back to TOC](#table-of-contents)

## FoundationDBStatusLayerInfo

FoundationDBStatusLayerInfo provides information about layers that are running against the cluster.

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| backup | Backup provides information about backups that have been started. | [FoundationDBStatusBackupInfo](#foundationdbstatusbackupinfo) | false |

[Back to TOC](#table-of-contents)

## FoundationDBStatusLocalClientInfo

FoundationDBStatusLocalClientInfo contains information about the client connection from the process getting the status.

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| coordinators | Coordinators provides information about the cluster's coordinators. | [FoundationDBStatusCoordinatorInfo](#foundationdbstatuscoordinatorinfo) | false |
| database_status | DatabaseStatus provides a summary of the database's health. | [FoundationDBStatusClientDBStatus](#foundationdbstatusclientdbstatus) | false |

[Back to TOC](#table-of-contents)

## FoundationDBStatusMovingData

FoundationDBStatusMovingData provides information about the current data movement

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| highest_priority | HighestPriority provides the priority of the highest-priority data movement. | int | false |
| in_flight_bytes | InFlightBytes provides how many bytes are being actively moved. | int | false |
| in_queue_bytes | InQueueBytes provides how many bytes are pending data movement. | int | false |

[Back to TOC](#table-of-contents)

## FoundationDBStatusProcessInfo

FoundationDBStatusProcessInfo describes the \"processes\" portion of the cluster status

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| address | Address provides the address of the process. | string | false |
| class_type | ProcessClass provides the process class the process has been given. | string | false |
| command_line | CommandLine provides the command-line invocation for the process. | string | false |
| excluded | Excluded indicates whether the process has been excluded. | bool | false |
| locality | The locality information for the process. | map[string]string | false |
| version | The version of FoundationDB the process is running. | string | false |

[Back to TOC](#table-of-contents)

## FoundationDBStatusSupportedVersion

FoundationDBStatusSupportedVersion provides information about a version of FDB supported by the connected clients.

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| client_version | ClientVersion provides the version of FDB the client is connecting through. | string | false |
| connected_clients | ConnectedClient provides the clients that are using this version. | [][FoundationDBStatusConnectedClient](#foundationdbstatusconnectedclient) | true |
| max_protocol_clients | MaxProtocolClients provides the clients that are using this version as their highest supported protocol version. | [][FoundationDBStatusConnectedClient](#foundationdbstatusconnectedclient) | true |
| protocol_version | ProtocolVersion is the version of the wire protocol the client is using. | string | false |
| source_version | SourceVersion is the version of the source code that the client library was built from. | string | false |

[Back to TOC](#table-of-contents)

## ProcessAddress

ProcessAddress provides a structured address for a process.

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| IPAddress |  | string | false |
| Port |  | int | false |
| Flags |  | map[string]bool | false |

[Back to TOC](#table-of-contents)

## ProcessCounts

ProcessCounts represents the number of processes we have for each valid process class.  If one of the counts in the spec is set to 0, we will infer the process count for that class from the role counts. If one of the counts in the spec is set to -1, we will not create any processes for that class. See GetProcessCountsWithDefaults for more information on the rules for inferring process counts.

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| storage | Storage defines the number of storage class processes. | int | false |
| transaction | Transaction defines the number of transaction class processes. | int | false |
| stateless | Stateless defines the number of stateless class processes. | int | false |
| resolution | Resolution defines the number of resolution class processes. | int | false |
| unset |  | int | false |
| log |  | int | false |
| master |  | int | false |
| cluster_controller |  | int | false |
| proxy |  | int | false |
| resolver |  | int | false |
| router |  | int | false |
| ratekeeper |  | int | false |
| data_distributor |  | int | false |

[Back to TOC](#table-of-contents)

## Region

Region represents a region in the database configuration

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| datacenters | The data centers in this region. | [][DataCenter](#datacenter) | false |
| satellite_logs | The number of satellite logs that we should recruit. | int | false |
| satellite_redundancy_mode | The replication strategy for satellite logs. | string | false |

[Back to TOC](#table-of-contents)

## RequiredAddressSet

RequiredAddressSet provides settings for which addresses we need to listen on.

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| tls | TLS defines whether we need to listen on a TLS address. | bool | false |
| nonTLS | NonTLS defines whether we need to listen on a non-TLS address. | bool | false |

[Back to TOC](#table-of-contents)

## RoleCounts

RoleCounts represents the roles whose counts can be customized.

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| storage |  | int | false |
| logs |  | int | false |
| proxies |  | int | false |
| resolvers |  | int | false |
| log_routers |  | int | false |
| remote_logs |  | int | false |

[Back to TOC](#table-of-contents)

## VersionFlags

VersionFlags defines internal flags for new features in the database.

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| log_spill |  | int | false |
| log_version |  | int | false |

[Back to TOC](#table-of-contents)
