# RegistryApi

All URIs are relative to *http://localhost:8080/api/v1*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**registryAchievementsAchievementDefIdDelete**](#registryachievementsachievementdefiddelete) | **DELETE** /registry/achievements/{achievement_def_id} | Delete achievement-def|
|[**registryAchievementsAchievementDefIdGet**](#registryachievementsachievementdefidget) | **GET** /registry/achievements/{achievement_def_id} | Get achievement-def|
|[**registryAchievementsAchievementDefIdPut**](#registryachievementsachievementdefidput) | **PUT** /registry/achievements/{achievement_def_id} | Update achievement-def|
|[**registryAchievementsGet**](#registryachievementsget) | **GET** /registry/achievements | List achievement-def|
|[**registryAchievementsPost**](#registryachievementspost) | **POST** /registry/achievements | Create achievement-def|
|[**registryAppsGet**](#registryappsget) | **GET** /registry/apps | List apps|
|[**registryItemsGet**](#registryitemsget) | **GET** /registry/items | List item-def|
|[**registryItemsItemDefIdDelete**](#registryitemsitemdefiddelete) | **DELETE** /registry/items/{item_def_id} | Delete item-def|
|[**registryItemsItemDefIdGet**](#registryitemsitemdefidget) | **GET** /registry/items/{item_def_id} | Get item-def|
|[**registryItemsItemDefIdPut**](#registryitemsitemdefidput) | **PUT** /registry/items/{item_def_id} | Update item-def|
|[**registryItemsPost**](#registryitemspost) | **POST** /registry/items | Create item-def|

# **registryAchievementsAchievementDefIdDelete**
> GithubComNotPlatformInternalServerTemplatesResponseTemplate registryAchievementsAchievementDefIdDelete()


### Example

```typescript
import {
    RegistryApi,
    Configuration
} from 'not-games-sdk-public';

const configuration = new Configuration();
const apiInstance = new RegistryApi(configuration);

let achievementDefId: string; //achievement_def_id (default to undefined)

const { status, data } = await apiInstance.registryAchievementsAchievementDefIdDelete(
    achievementDefId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **achievementDefId** | [**string**] | achievement_def_id | defaults to undefined|


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

# **registryAchievementsAchievementDefIdGet**
> RegistryAchievementsAchievementDefIdGet200Response registryAchievementsAchievementDefIdGet()


### Example

```typescript
import {
    RegistryApi,
    Configuration
} from 'not-games-sdk-public';

const configuration = new Configuration();
const apiInstance = new RegistryApi(configuration);

let achievementDefId: string; //achievement_def_id (default to undefined)

const { status, data } = await apiInstance.registryAchievementsAchievementDefIdGet(
    achievementDefId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **achievementDefId** | [**string**] | achievement_def_id | defaults to undefined|


### Return type

**RegistryAchievementsAchievementDefIdGet200Response**

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

# **registryAchievementsAchievementDefIdPut**
> GithubComNotPlatformInternalServerTemplatesResponseTemplate registryAchievementsAchievementDefIdPut(request)


### Example

```typescript
import {
    RegistryApi,
    Configuration,
    InternalModuleRegistryPresentationAchievementHttpUpdateRequest
} from 'not-games-sdk-public';

const configuration = new Configuration();
const apiInstance = new RegistryApi(configuration);

let achievementDefId: string; //achievement_def_id (default to undefined)
let request: InternalModuleRegistryPresentationAchievementHttpUpdateRequest; //request body

const { status, data } = await apiInstance.registryAchievementsAchievementDefIdPut(
    achievementDefId,
    request
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **request** | **InternalModuleRegistryPresentationAchievementHttpUpdateRequest**| request body | |
| **achievementDefId** | [**string**] | achievement_def_id | defaults to undefined|


### Return type

**GithubComNotPlatformInternalServerTemplatesResponseTemplate**

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

# **registryAchievementsGet**
> RegistryAchievementsGet200Response registryAchievementsGet()


### Example

```typescript
import {
    RegistryApi,
    Configuration
} from 'not-games-sdk-public';

const configuration = new Configuration();
const apiInstance = new RegistryApi(configuration);

let limit: number; //limit (default to undefined)
let offset: number; //offset (default to undefined)

const { status, data } = await apiInstance.registryAchievementsGet(
    limit,
    offset
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **limit** | [**number**] | limit | defaults to undefined|
| **offset** | [**number**] | offset | defaults to undefined|


### Return type

**RegistryAchievementsGet200Response**

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

# **registryAchievementsPost**
> GithubComNotPlatformInternalServerTemplatesResponseTemplate registryAchievementsPost(request)


### Example

```typescript
import {
    RegistryApi,
    Configuration,
    InternalModuleRegistryPresentationAchievementHttpCreateRequest
} from 'not-games-sdk-public';

const configuration = new Configuration();
const apiInstance = new RegistryApi(configuration);

let request: InternalModuleRegistryPresentationAchievementHttpCreateRequest; //request body

const { status, data } = await apiInstance.registryAchievementsPost(
    request
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **request** | **InternalModuleRegistryPresentationAchievementHttpCreateRequest**| request body | |


### Return type

**GithubComNotPlatformInternalServerTemplatesResponseTemplate**

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

# **registryAppsGet**
> RegistryAppsGet200Response registryAppsGet()


### Example

```typescript
import {
    RegistryApi,
    Configuration
} from 'not-games-sdk-public';

const configuration = new Configuration();
const apiInstance = new RegistryApi(configuration);

let limit: number; //limit (default to undefined)
let offset: number; //offset (default to undefined)

const { status, data } = await apiInstance.registryAppsGet(
    limit,
    offset
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **limit** | [**number**] | limit | defaults to undefined|
| **offset** | [**number**] | offset | defaults to undefined|


### Return type

**RegistryAppsGet200Response**

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

# **registryItemsGet**
> RegistryItemsGet200Response registryItemsGet()


### Example

```typescript
import {
    RegistryApi,
    Configuration
} from 'not-games-sdk-public';

const configuration = new Configuration();
const apiInstance = new RegistryApi(configuration);

let limit: number; //limit (default to undefined)
let offset: number; //offset (default to undefined)

const { status, data } = await apiInstance.registryItemsGet(
    limit,
    offset
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **limit** | [**number**] | limit | defaults to undefined|
| **offset** | [**number**] | offset | defaults to undefined|


### Return type

**RegistryItemsGet200Response**

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

# **registryItemsItemDefIdDelete**
> GithubComNotPlatformInternalServerTemplatesResponseTemplate registryItemsItemDefIdDelete()


### Example

```typescript
import {
    RegistryApi,
    Configuration
} from 'not-games-sdk-public';

const configuration = new Configuration();
const apiInstance = new RegistryApi(configuration);

let itemDefId: string; //item_def_id (default to undefined)

const { status, data } = await apiInstance.registryItemsItemDefIdDelete(
    itemDefId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **itemDefId** | [**string**] | item_def_id | defaults to undefined|


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

# **registryItemsItemDefIdGet**
> RegistryItemsItemDefIdGet200Response registryItemsItemDefIdGet()


### Example

```typescript
import {
    RegistryApi,
    Configuration
} from 'not-games-sdk-public';

const configuration = new Configuration();
const apiInstance = new RegistryApi(configuration);

let itemDefId: string; //item_def_id (default to undefined)

const { status, data } = await apiInstance.registryItemsItemDefIdGet(
    itemDefId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **itemDefId** | [**string**] | item_def_id | defaults to undefined|


### Return type

**RegistryItemsItemDefIdGet200Response**

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

# **registryItemsItemDefIdPut**
> GithubComNotPlatformInternalServerTemplatesResponseTemplate registryItemsItemDefIdPut(request)


### Example

```typescript
import {
    RegistryApi,
    Configuration,
    InternalModuleRegistryPresentationItemHttpUpdateRequest
} from 'not-games-sdk-public';

const configuration = new Configuration();
const apiInstance = new RegistryApi(configuration);

let itemDefId: string; //item_def_id (default to undefined)
let request: InternalModuleRegistryPresentationItemHttpUpdateRequest; //request body

const { status, data } = await apiInstance.registryItemsItemDefIdPut(
    itemDefId,
    request
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **request** | **InternalModuleRegistryPresentationItemHttpUpdateRequest**| request body | |
| **itemDefId** | [**string**] | item_def_id | defaults to undefined|


### Return type

**GithubComNotPlatformInternalServerTemplatesResponseTemplate**

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

# **registryItemsPost**
> GithubComNotPlatformInternalServerTemplatesResponseTemplate registryItemsPost(request)


### Example

```typescript
import {
    RegistryApi,
    Configuration,
    InternalModuleRegistryPresentationItemHttpCreateRequest
} from 'not-games-sdk-public';

const configuration = new Configuration();
const apiInstance = new RegistryApi(configuration);

let request: InternalModuleRegistryPresentationItemHttpCreateRequest; //request body

const { status, data } = await apiInstance.registryItemsPost(
    request
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **request** | **InternalModuleRegistryPresentationItemHttpCreateRequest**| request body | |


### Return type

**GithubComNotPlatformInternalServerTemplatesResponseTemplate**

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

