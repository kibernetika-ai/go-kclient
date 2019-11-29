# MlappGitRepoVolumeSource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** |  | [optional] [default to null]
**AccountId** | **string** |  | [optional] [default to null]
**Directory** | **string** | Target directory name. Must not contain or start with &#39;..&#39;.  If &#39;.&#39; is supplied, the volume directory will be the git repository.  Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name. | [optional] [default to null]
**PrivateKey** | **string** |  | [optional] [default to null]
**Repository** | **string** | Repository URL | [default to null]
**Revision** | **string** | Commit hash for the specified revision. | [optional] [default to null]
**UserName** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


