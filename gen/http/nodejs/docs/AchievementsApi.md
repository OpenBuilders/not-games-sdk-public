# AchievementsApi

All URIs are relative to *http://localhost:8080/api/v1*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**achievementsGet**](#achievementsget) | **GET** /achievements | List achievements|
|[**achievementsIdGet**](#achievementsidget) | **GET** /achievements/{id} | Get achievement|
|[**achievementsPost**](#achievementspost) | **POST** /achievements | Create achievement|

# **achievementsGet**
> AchievementsGet200Response achievementsGet()


### Example

```typescript
import {
    AchievementsApi,
    Configuration
} from 'not-games-sdk-public';

const configuration = new Configuration();
const apiInstance = new AchievementsApi(configuration);

let appId: string; //app id (default to undefined)
let limit: number; //limit (default to undefined)
let offset: number; //offset (default to undefined)

const { status, data } = await apiInstance.achievementsGet(
    appId,
    limit,
    offset
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **appId** | [**string**] | app id | defaults to undefined|
| **limit** | [**number**] | limit | defaults to undefined|
| **offset** | [**number**] | offset | defaults to undefined|


### Return type

**AchievementsGet200Response**

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

# **achievementsIdGet**
> AchievementsIdGet200Response achievementsIdGet()


### Example

```typescript
import {
    AchievementsApi,
    Configuration
} from 'not-games-sdk-public';

const configuration = new Configuration();
const apiInstance = new AchievementsApi(configuration);

let id: string; //achievement id (default to undefined)

const { status, data } = await apiInstance.achievementsIdGet(
    id
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **id** | [**string**] | achievement id | defaults to undefined|


### Return type

**AchievementsIdGet200Response**

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

# **achievementsPost**
> AchievementsPost200Response achievementsPost(request)


### Example

```typescript
import {
    AchievementsApi,
    Configuration,
    InternalModuleAchievementPresentationHttpCreateAchievementsRequest
} from 'not-games-sdk-public';

const configuration = new Configuration();
const apiInstance = new AchievementsApi(configuration);

let request: InternalModuleAchievementPresentationHttpCreateAchievementsRequest; //request body

const { status, data } = await apiInstance.achievementsPost(
    request
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **request** | **InternalModuleAchievementPresentationHttpCreateAchievementsRequest**| request body | |


### Return type

**AchievementsPost200Response**

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

