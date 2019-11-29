# \ServingApi

All URIs are relative to *https://dev.kibernetika.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServingDelete**](ServingApi.md#ServingDelete) | **Delete** /api/v0.2/workspace/{workspace}/serving/{serving} | Delete serving
[**ServingDisable**](ServingApi.md#ServingDisable) | **Post** /api/v0.2/workspace/{workspace}/serving/{serving}/disable | Disable serving
[**ServingEnable**](ServingApi.md#ServingEnable) | **Post** /api/v0.2/workspace/{workspace}/serving/{serving}/enable | Enable serving
[**ServingInfo**](ServingApi.md#ServingInfo) | **Get** /api/v0.2/workspace/{workspace}/serving/{serving} | Return serving&#39;s info
[**ServingTFProxyModel**](ServingApi.md#ServingTFProxyModel) | **Post** /api/v0.2/workspace/{workspace}/serving/{serving}/tfproxy/{port}/{model} | TF proxy to serving (model)
[**ServingTFProxyModelSignature**](ServingApi.md#ServingTFProxyModelSignature) | **Post** /api/v0.2/workspace/{workspace}/serving/{serving}/tfproxy/{port}/{model}/{signature} | TF proxy to serving (model, signature)
[**ServingTFProxyModelSignatureVersion**](ServingApi.md#ServingTFProxyModelSignatureVersion) | **Post** /api/v0.2/workspace/{workspace}/serving/{serving}/tfproxy/{port}/{model}/{signature}/{version} | TF proxy to serving (model, signature, version)
[**UpdateServing**](ServingApi.md#UpdateServing) | **Put** /api/v0.2/workspace/{workspace}/serving/{serving} | Update serving


# **ServingDelete**
> ServingDelete(ctx, workspace, serving, optional)
Delete serving

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| Workspace&#39;s name | 
  **serving** | **string**| Serving&#39;s Name or ID | 
 **optional** | ***ServingDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServingDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dependencies** | **optional.Bool**| Get only dependencies | 
 **confirm** | **optional.String**| String for confirmation | 
 **force** | **optional.Bool**| Force deletion | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServingDisable**
> ModelsServing ServingDisable(ctx, workspace, serving)
Disable serving

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| Workspace&#39;s name | 
  **serving** | **string**| Serving&#39;s Name or ID | 

### Return type

[**ModelsServing**](models.Serving.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServingEnable**
> ModelsServing ServingEnable(ctx, workspace, serving)
Enable serving

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| Workspace&#39;s name | 
  **serving** | **string**| Serving&#39;s Name or ID | 

### Return type

[**ModelsServing**](models.Serving.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServingInfo**
> ModelsServing ServingInfo(ctx, workspace, serving)
Return serving's info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| Workspace&#39;s name | 
  **serving** | **string**| Serving&#39;s Name or ID | 

### Return type

[**ModelsServing**](models.Serving.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServingTFProxyModel**
> ModelsArbitrary ServingTFProxyModel(ctx, model, body, workspace, serving, port)
TF proxy to serving (model)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | **string**| Serving model | 
  **body** | [**ModelsArbitrary**](ModelsArbitrary.md)|  | 
  **workspace** | **string**| Workspace&#39;s name | 
  **serving** | **string**| Serving&#39;s Name or ID | 
  **port** | **string**| Serving port | 

### Return type

[**ModelsArbitrary**](models.Arbitrary.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServingTFProxyModelSignature**
> ModelsArbitrary ServingTFProxyModelSignature(ctx, model, signature, body, workspace, serving, port)
TF proxy to serving (model, signature)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | **string**| Serving model | 
  **signature** | **string**| Serving signature | 
  **body** | [**ModelsArbitrary**](ModelsArbitrary.md)|  | 
  **workspace** | **string**| Workspace&#39;s name | 
  **serving** | **string**| Serving&#39;s Name or ID | 
  **port** | **string**| Serving port | 

### Return type

[**ModelsArbitrary**](models.Arbitrary.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServingTFProxyModelSignatureVersion**
> ModelsArbitrary ServingTFProxyModelSignatureVersion(ctx, model, signature, version, body, workspace, serving, port)
TF proxy to serving (model, signature, version)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | **string**| Serving model | 
  **signature** | **string**| Serving signature | 
  **version** | **string**| Serving version | 
  **body** | [**ModelsArbitrary**](ModelsArbitrary.md)|  | 
  **workspace** | **string**| Workspace&#39;s name | 
  **serving** | **string**| Serving&#39;s Name or ID | 
  **port** | **string**| Serving port | 

### Return type

[**ModelsArbitrary**](models.Arbitrary.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateServing**
> ModelsServing UpdateServing(ctx, body, workspace, serving)
Update serving

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MlappServing**](MlappServing.md)|  | 
  **workspace** | **string**| Workspace&#39;s name | 
  **serving** | **string**| Serving&#39;s Name or ID | 

### Return type

[**ModelsServing**](models.Serving.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

