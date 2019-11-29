# \InferenceApi

All URIs are relative to *https://dev.kibernetika.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InferenceInferenceVersionDelete**](InferenceApi.md#InferenceInferenceVersionDelete) | **Delete** /api/v0.2/workspace/{workspace}/inference/{inference}/versions/{version} | Delete inference&#39;s version
[**InferenceInferenceVersionInfo**](InferenceApi.md#InferenceInferenceVersionInfo) | **Get** /api/v0.2/workspace/{workspace}/inference/{inference}/versions/{version} | Return inference&#39;s info for specified version
[**InferenceInferenceVersionStart**](InferenceApi.md#InferenceInferenceVersionStart) | **Post** /api/v0.2/workspace/{workspace}/inference/{inference}/versions/{version}/start | Starts serving
[**InferenceInferenceVersionUpdate**](InferenceApi.md#InferenceInferenceVersionUpdate) | **Put** /api/v0.2/workspace/{workspace}/inference/{inference}/versions/{version} | Update inference&#39;s info for specified version


# **InferenceInferenceVersionDelete**
> InferenceInferenceVersionDelete(ctx, workspace, inference, version)
Delete inference's version

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| Workspace&#39;s name | 
  **inference** | **string**| Item&#39;s name (Inference) | 
  **version** | **string**| Inference&#39;s version | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InferenceInferenceVersionInfo**
> ModelsInferenceVersion InferenceInferenceVersionInfo(ctx, workspace, inference, version)
Return inference's info for specified version

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| Workspace&#39;s name | 
  **inference** | **string**| Item&#39;s name (Inference) | 
  **version** | **string**| Inference&#39;s version | 

### Return type

[**ModelsInferenceVersion**](models.InferenceVersion.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InferenceInferenceVersionStart**
> ModelsServing InferenceInferenceVersionStart(ctx, body, workspace, inference, version)
Starts serving

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**InferenceRunServingRequest**](InferenceRunServingRequest.md)|  | 
  **workspace** | **string**| Workspace&#39;s name | 
  **inference** | **string**| Item&#39;s name (Inference) | 
  **version** | **string**| Inference&#39;s version | 

### Return type

[**ModelsServing**](models.Serving.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InferenceInferenceVersionUpdate**
> ModelsInferenceVersion InferenceInferenceVersionUpdate(ctx, body, workspace, inference, version)
Update inference's info for specified version

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ModelsInferenceVersion**](ModelsInferenceVersion.md)|  | 
  **workspace** | **string**| Workspace&#39;s name | 
  **inference** | **string**| Item&#39;s name (Inference) | 
  **version** | **string**| Inference&#39;s version | 

### Return type

[**ModelsInferenceVersion**](models.InferenceVersion.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

