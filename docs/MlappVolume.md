# MlappVolume

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterStorage** | **string** |  | [optional] [default to null]
**Dataset** | [***MlappDatasetSource**](mlapp.DatasetSource.md) |  | [optional] [default to null]
**DatasetFS** | [***MlappDatasetFsSource**](mlapp.DatasetFSSource.md) |  | [optional] [default to null]
**EmptyDir** | [***V1EmptyDirVolumeSource**](v1.EmptyDirVolumeSource.md) |  | [optional] [default to null]
**FlexVolume** | [***V1FlexVolumeSource**](v1.FlexVolumeSource.md) |  | [optional] [default to null]
**GitRepo** | [***MlappGitRepoVolumeSource**](mlapp.GitRepoVolumeSource.md) |  | [optional] [default to null]
**HostPath** | [***V1HostPathVolumeSource**](v1.HostPathVolumeSource.md) |  | [optional] [default to null]
**IsLibDir** | **bool** |  | [optional] [default to null]
**IsTrainLogDir** | **bool** |  | [optional] [default to null]
**IsWorkspaceLocal** | **bool** |  | [optional] [default to null]
**Model** | [***MlappModelSource**](mlapp.ModelSource.md) |  | [optional] [default to null]
**MountPath** | **string** |  | [optional] [default to null]
**Name** | **string** |  | [default to null]
**Nfs** | [***V1NfsVolumeSource**](v1.NFSVolumeSource.md) |  | [optional] [default to null]
**PersistentStorage** | [***MlappPersistentStorage**](mlapp.PersistentStorage.md) |  | [optional] [default to null]
**PersistentVolumeClaim** | [***V1PersistentVolumeClaimVolumeSource**](v1.PersistentVolumeClaimVolumeSource.md) |  | [optional] [default to null]
**ReadOnly** | **bool** |  | [optional] [default to null]
**S3bucket** | [***MlappS3BucketSource**](mlapp.S3BucketSource.md) |  | [optional] [default to null]
**SubPath** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


