# AuthApi

All URIs are relative to *http://localhost:8080/api/v1*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**authGeneratePost**](#authgeneratepost) | **POST** /auth/generate | Generate new auth|

# **authGeneratePost**
> AuthGeneratePost200Response authGeneratePost(request)


### Example

```typescript
import {
    AuthApi,
    Configuration,
    GithubComNotPlatformInternalModuleAuthPresentationAuthorizationHttpModelAuthGenerateRequest
} from 'not-games-sdk-public';

const configuration = new Configuration();
const apiInstance = new AuthApi(configuration);

let request: GithubComNotPlatformInternalModuleAuthPresentationAuthorizationHttpModelAuthGenerateRequest; //request body

const { status, data } = await apiInstance.authGeneratePost(
    request
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **request** | **GithubComNotPlatformInternalModuleAuthPresentationAuthorizationHttpModelAuthGenerateRequest**| request body | |


### Return type

**AuthGeneratePost200Response**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |
|**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

