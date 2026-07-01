# InventoryApi

All URIs are relative to *http://localhost:8080/api/v1*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**inventoryByItemdefGet**](#inventorybyitemdefget) | **GET** /inventory/by_itemdef | List items|
|[**inventoryItemsGeneratePost**](#inventoryitemsgeneratepost) | **POST** /inventory/items/generate | Generate item|
|[**inventoryItemsGet**](#inventoryitemsget) | **GET** /inventory/items | List items|
|[**inventoryItemsItemIdConsumePost**](#inventoryitemsitemidconsumepost) | **POST** /inventory/items/{item_id}/consume | Consume item|
|[**inventoryItemsItemIdGet**](#inventoryitemsitemidget) | **GET** /inventory/items/{item_id} | Get item|
|[**inventoryItemsItemIdIncrementQuantityPost**](#inventoryitemsitemidincrementquantitypost) | **POST** /inventory/items/{item_id}/increment-quantity | Consume item|
|[**inventoryItemsItemIdPut**](#inventoryitemsitemidput) | **PUT** /inventory/items/{item_id} | Update item|
|[**inventoryItemsItemIdTransferLogPost**](#inventoryitemsitemidtransferlogpost) | **POST** /inventory/items/{item_id}/transfer_log | Get transfer logs|
|[**inventoryItemsItemIdTransferPost**](#inventoryitemsitemidtransferpost) | **POST** /inventory/items/{item_id}/transfer | Transfer item to another account|
|[**inventoryItemsItemIdTransferPreflightPost**](#inventoryitemsitemidtransferpreflightpost) | **POST** /inventory/items/{item_id}/transfer-preflight | Check transfer item possibility to another account|
|[**inventoryItemsUnpackPost**](#inventoryitemsunpackpost) | **POST** /inventory/items/unpack | Unpack item|

# **inventoryByItemdefGet**
> InventoryByItemdefGet200Response inventoryByItemdefGet()


### Example

```typescript
import {
    InventoryApi,
    Configuration
} from 'not-games-sdk-public';

const configuration = new Configuration();
const apiInstance = new InventoryApi(configuration);

let limit: number; //limit (default to undefined)
let offset: number; //offset (default to undefined)
let itemDefId: string; //item_def_id (default to undefined)
let sortFields: string; //comma separated sort fields (optional) (default to undefined)
let directions: string; //comma separated sort directions (optional) (default to undefined)

const { status, data } = await apiInstance.inventoryByItemdefGet(
    limit,
    offset,
    itemDefId,
    sortFields,
    directions
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **limit** | [**number**] | limit | defaults to undefined|
| **offset** | [**number**] | offset | defaults to undefined|
| **itemDefId** | [**string**] | item_def_id | defaults to undefined|
| **sortFields** | [**string**] | comma separated sort fields | (optional) defaults to undefined|
| **directions** | [**string**] | comma separated sort directions | (optional) defaults to undefined|


### Return type

**InventoryByItemdefGet200Response**

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

# **inventoryItemsGeneratePost**
> InventoryItemsGeneratePost200Response inventoryItemsGeneratePost(request)


### Example

```typescript
import {
    InventoryApi,
    Configuration,
    GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelGenerateItemRequest
} from 'not-games-sdk-public';

const configuration = new Configuration();
const apiInstance = new InventoryApi(configuration);

let request: GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelGenerateItemRequest; //request body

const { status, data } = await apiInstance.inventoryItemsGeneratePost(
    request
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **request** | **GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelGenerateItemRequest**| request body | |


### Return type

**InventoryItemsGeneratePost200Response**

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

# **inventoryItemsGet**
> InventoryByItemdefGet200Response inventoryItemsGet()


### Example

```typescript
import {
    InventoryApi,
    Configuration
} from 'not-games-sdk-public';

const configuration = new Configuration();
const apiInstance = new InventoryApi(configuration);

let limit: number; //limit (default to undefined)
let offset: number; //offset (default to undefined)

const { status, data } = await apiInstance.inventoryItemsGet(
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

**InventoryByItemdefGet200Response**

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

# **inventoryItemsItemIdConsumePost**
> GithubComNotPlatformInternalServerTemplatesResponseTemplate inventoryItemsItemIdConsumePost(request)


### Example

```typescript
import {
    InventoryApi,
    Configuration,
    GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelConsumeItemRequest
} from 'not-games-sdk-public';

const configuration = new Configuration();
const apiInstance = new InventoryApi(configuration);

let itemId: string; //item id (uuid) (default to undefined)
let accountId: string; //account id (int64) (default to undefined)
let request: GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelConsumeItemRequest; //request body

const { status, data } = await apiInstance.inventoryItemsItemIdConsumePost(
    itemId,
    accountId,
    request
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **request** | **GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelConsumeItemRequest**| request body | |
| **itemId** | [**string**] | item id (uuid) | defaults to undefined|
| **accountId** | [**string**] | account id (int64) | defaults to undefined|


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

# **inventoryItemsItemIdGet**
> InventoryItemsItemIdGet200Response inventoryItemsItemIdGet()


### Example

```typescript
import {
    InventoryApi,
    Configuration
} from 'not-games-sdk-public';

const configuration = new Configuration();
const apiInstance = new InventoryApi(configuration);

let itemId: string; //item id (uuid) (default to undefined)
let appId: number; //app id (optional) (default to undefined)

const { status, data } = await apiInstance.inventoryItemsItemIdGet(
    itemId,
    appId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **itemId** | [**string**] | item id (uuid) | defaults to undefined|
| **appId** | [**number**] | app id | (optional) defaults to undefined|


### Return type

**InventoryItemsItemIdGet200Response**

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

# **inventoryItemsItemIdIncrementQuantityPost**
> GithubComNotPlatformInternalServerTemplatesResponseTemplate inventoryItemsItemIdIncrementQuantityPost(request)


### Example

```typescript
import {
    InventoryApi,
    Configuration,
    GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelConsumeItemRequest
} from 'not-games-sdk-public';

const configuration = new Configuration();
const apiInstance = new InventoryApi(configuration);

let itemId: string; //item id (uuid) (default to undefined)
let accountId: string; //account id (int64) (default to undefined)
let request: GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelConsumeItemRequest; //request body

const { status, data } = await apiInstance.inventoryItemsItemIdIncrementQuantityPost(
    itemId,
    accountId,
    request
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **request** | **GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelConsumeItemRequest**| request body | |
| **itemId** | [**string**] | item id (uuid) | defaults to undefined|
| **accountId** | [**string**] | account id (int64) | defaults to undefined|


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

# **inventoryItemsItemIdPut**
> GithubComNotPlatformInternalServerTemplatesResponseTemplate inventoryItemsItemIdPut(request)


### Example

```typescript
import {
    InventoryApi,
    Configuration,
    GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelUpdateItemRequest
} from 'not-games-sdk-public';

const configuration = new Configuration();
const apiInstance = new InventoryApi(configuration);

let itemId: string; //item id (uuid) (default to undefined)
let request: GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelUpdateItemRequest; //request body

const { status, data } = await apiInstance.inventoryItemsItemIdPut(
    itemId,
    request
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **request** | **GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelUpdateItemRequest**| request body | |
| **itemId** | [**string**] | item id (uuid) | defaults to undefined|


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

# **inventoryItemsItemIdTransferLogPost**
> GithubComNotPlatformInternalServerTemplatesResponseTemplate inventoryItemsItemIdTransferLogPost()


### Example

```typescript
import {
    InventoryApi,
    Configuration
} from 'not-games-sdk-public';

const configuration = new Configuration();
const apiInstance = new InventoryApi(configuration);

let itemId: string; //item id (uuid) (default to undefined)
let limit: number; //limit (default to undefined)
let offset: number; //offset (default to undefined)

const { status, data } = await apiInstance.inventoryItemsItemIdTransferLogPost(
    itemId,
    limit,
    offset
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **itemId** | [**string**] | item id (uuid) | defaults to undefined|
| **limit** | [**number**] | limit | defaults to undefined|
| **offset** | [**number**] | offset | defaults to undefined|


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

# **inventoryItemsItemIdTransferPost**
> GithubComNotPlatformInternalServerTemplatesResponseTemplate inventoryItemsItemIdTransferPost(request)


### Example

```typescript
import {
    InventoryApi,
    Configuration,
    GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelTransferItemRequest
} from 'not-games-sdk-public';

const configuration = new Configuration();
const apiInstance = new InventoryApi(configuration);

let itemId: string; //item id (uuid) (default to undefined)
let request: GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelTransferItemRequest; //request body

const { status, data } = await apiInstance.inventoryItemsItemIdTransferPost(
    itemId,
    request
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **request** | **GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelTransferItemRequest**| request body | |
| **itemId** | [**string**] | item id (uuid) | defaults to undefined|


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

# **inventoryItemsItemIdTransferPreflightPost**
> InventoryItemsItemIdTransferPreflightPost200Response inventoryItemsItemIdTransferPreflightPost(request)


### Example

```typescript
import {
    InventoryApi,
    Configuration,
    GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelTransferItemRequest
} from 'not-games-sdk-public';

const configuration = new Configuration();
const apiInstance = new InventoryApi(configuration);

let itemId: string; //item id (uuid) (default to undefined)
let request: GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelTransferItemRequest; //request body

const { status, data } = await apiInstance.inventoryItemsItemIdTransferPreflightPost(
    itemId,
    request
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **request** | **GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelTransferItemRequest**| request body | |
| **itemId** | [**string**] | item id (uuid) | defaults to undefined|


### Return type

**InventoryItemsItemIdTransferPreflightPost200Response**

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

# **inventoryItemsUnpackPost**
> InventoryItemsUnpackPost200Response inventoryItemsUnpackPost(request)


### Example

```typescript
import {
    InventoryApi,
    Configuration,
    GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelUnpackItemRequest
} from 'not-games-sdk-public';

const configuration = new Configuration();
const apiInstance = new InventoryApi(configuration);

let request: GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelUnpackItemRequest; //request body

const { status, data } = await apiInstance.inventoryItemsUnpackPost(
    request
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **request** | **GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelUnpackItemRequest**| request body | |


### Return type

**InventoryItemsUnpackPost200Response**

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

