# DropsApi

All URIs are relative to *http://localhost:8080/api/v1*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**dropsIdGet**](#dropsidget) | **GET** /drops/{id} | Check user eligibility for a drop|
|[**dropsIdWhitelistPost**](#dropsidwhitelistpost) | **POST** /drops/{id}/whitelist | Whitelist users for a drop|
|[**dropsPost**](#dropspost) | **POST** /drops | Create a new drop|

# **dropsIdGet**
> DropsIdGet200Response dropsIdGet()

Check user eligibility for a drop. Account ID must be provided in authorization token. This only used for game to know if user can claim drop.

### Example

```typescript
import {
    DropsApi,
    Configuration
} from 'not-games-sdk-public';

const configuration = new Configuration();
const apiInstance = new DropsApi(configuration);

let id: string; //Drop ID (default to undefined)

const { status, data } = await apiInstance.dropsIdGet(
    id
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **id** | [**string**] | Drop ID | defaults to undefined|


### Return type

**DropsIdGet200Response**

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |
|**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **dropsIdWhitelistPost**
> GithubComNotPlatformInternalServerTemplatesResponseTemplate dropsIdWhitelistPost(request)


### Example

```typescript
import {
    DropsApi,
    Configuration,
    InternalModuleDropsPresentationDropsHttpWhitelistDropUserRequest
} from 'not-games-sdk-public';

const configuration = new Configuration();
const apiInstance = new DropsApi(configuration);

let id: string; //Drop ID (default to undefined)
let request: InternalModuleDropsPresentationDropsHttpWhitelistDropUserRequest; //Whitelist Drop User Request

const { status, data } = await apiInstance.dropsIdWhitelistPost(
    id,
    request
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **request** | **InternalModuleDropsPresentationDropsHttpWhitelistDropUserRequest**| Whitelist Drop User Request | |
| **id** | [**string**] | Drop ID | defaults to undefined|


### Return type

**GithubComNotPlatformInternalServerTemplatesResponseTemplate**

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |
|**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **dropsPost**
> DropsPost200Response dropsPost(request)

Create a new drop. App ID must be provided in authorization token

### Example

```typescript
import {
    DropsApi,
    Configuration,
    InternalModuleDropsPresentationDropsHttpCreateDropRequest
} from 'not-games-sdk-public';

const configuration = new Configuration();
const apiInstance = new DropsApi(configuration);

let request: InternalModuleDropsPresentationDropsHttpCreateDropRequest; //Create Drop Request

const { status, data } = await apiInstance.dropsPost(
    request
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **request** | **InternalModuleDropsPresentationDropsHttpCreateDropRequest**| Create Drop Request | |


### Return type

**DropsPost200Response**

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |
|**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

