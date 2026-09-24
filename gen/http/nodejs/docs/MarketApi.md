# MarketApi

All URIs are relative to *http://localhost:8080/api/v1*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**apiV1MarketAppsNewGet**](#apiv1marketappsnewget) | **GET** /api/v1/market/apps/new | List new visible apps for external market|
|[**apiV1MarketItemDefGet**](#apiv1marketitemdefget) | **GET** /api/v1/market/item_def | List item defs for external market|
|[**apiV1MarketItemGet**](#apiv1marketitemget) | **GET** /api/v1/market/item | List items for external market|
|[**apiV1MarketItemItemIdGet**](#apiv1marketitemitemidget) | **GET** /api/v1/market/item/{item_id} | Get any user item by id for external market|
|[**apiV1MarketItemItemIdTransferPost**](#apiv1marketitemitemidtransferpost) | **POST** /api/v1/market/item/{item_id}/transfer | Transfer an item to a user from external market|
|[**apiV1MarketProfileAccountIdItemsByAppsGet**](#apiv1marketprofileaccountiditemsbyappsget) | **GET** /api/v1/market/profile/{account_id}/items-by-apps | market profile apps items|
|[**apiV1MarketProfileAccountIdItemsByDisplayTypesGet**](#apiv1marketprofileaccountiditemsbydisplaytypesget) | **GET** /api/v1/market/profile/{account_id}/items-by-display-types | market profile app items|

# **apiV1MarketAppsNewGet**
> ApiV1MarketAppsNewGet200Response apiV1MarketAppsNewGet()


### Example

```typescript
import {
    MarketApi,
    Configuration
} from 'notgamessdkpublic';

const configuration = new Configuration();
const apiInstance = new MarketApi(configuration);

let fromDate: string; //from date, RFC3339 or YYYY-MM-DD (optional) (default to undefined)
let limit: number; //limit, max 500 (optional) (default to undefined)
let sortFields: string; //comma separated: created_at|name (optional) (default to undefined)
let directions: string; //comma separated: asc|desc (optional) (default to undefined)

const { status, data } = await apiInstance.apiV1MarketAppsNewGet(
    fromDate,
    limit,
    sortFields,
    directions
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **fromDate** | [**string**] | from date, RFC3339 or YYYY-MM-DD | (optional) defaults to undefined|
| **limit** | [**number**] | limit, max 500 | (optional) defaults to undefined|
| **sortFields** | [**string**] | comma separated: created_at|name | (optional) defaults to undefined|
| **directions** | [**string**] | comma separated: asc|desc | (optional) defaults to undefined|


### Return type

**ApiV1MarketAppsNewGet200Response**

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

# **apiV1MarketItemDefGet**
> ApiV1MarketItemDefGet200Response apiV1MarketItemDefGet()


### Example

```typescript
import {
    MarketApi,
    Configuration
} from 'notgamessdkpublic';

const configuration = new Configuration();
const apiInstance = new MarketApi(configuration);

let limit: number; //limit (optional) (default to undefined)
let offset: number; //offset (optional) (default to undefined)
let sortFields: string; //comma separated sort fields (optional) (default to undefined)
let directions: string; //comma separated: asc|desc (optional) (default to undefined)

const { status, data } = await apiInstance.apiV1MarketItemDefGet(
    limit,
    offset,
    sortFields,
    directions
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **limit** | [**number**] | limit | (optional) defaults to undefined|
| **offset** | [**number**] | offset | (optional) defaults to undefined|
| **sortFields** | [**string**] | comma separated sort fields | (optional) defaults to undefined|
| **directions** | [**string**] | comma separated: asc|desc | (optional) defaults to undefined|


### Return type

**ApiV1MarketItemDefGet200Response**

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |
|**400** | Bad Request |  -  |
|**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **apiV1MarketItemGet**
> ApiV1MarketItemGet200Response apiV1MarketItemGet()


### Example

```typescript
import {
    MarketApi,
    Configuration
} from 'notgamessdkpublic';

const configuration = new Configuration();
const apiInstance = new MarketApi(configuration);

let limit: number; //limit (optional) (default to undefined)
let offset: number; //offset (optional) (default to undefined)
let sortFields: string; //comma separated sort fields (optional) (default to undefined)
let directions: string; //comma separated: asc|desc (optional) (default to undefined)

const { status, data } = await apiInstance.apiV1MarketItemGet(
    limit,
    offset,
    sortFields,
    directions
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **limit** | [**number**] | limit | (optional) defaults to undefined|
| **offset** | [**number**] | offset | (optional) defaults to undefined|
| **sortFields** | [**string**] | comma separated sort fields | (optional) defaults to undefined|
| **directions** | [**string**] | comma separated: asc|desc | (optional) defaults to undefined|


### Return type

**ApiV1MarketItemGet200Response**

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |
|**400** | Bad Request |  -  |
|**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **apiV1MarketItemItemIdGet**
> ApiV1MarketItemItemIdGet200Response apiV1MarketItemItemIdGet()


### Example

```typescript
import {
    MarketApi,
    Configuration
} from 'notgamessdkpublic';

const configuration = new Configuration();
const apiInstance = new MarketApi(configuration);

let itemId: string; //Item ID (default to undefined)

const { status, data } = await apiInstance.apiV1MarketItemItemIdGet(
    itemId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **itemId** | [**string**] | Item ID | defaults to undefined|


### Return type

**ApiV1MarketItemItemIdGet200Response**

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |
|**400** | Bad Request |  -  |
|**404** | Not Found |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **apiV1MarketItemItemIdTransferPost**
> GithubComNotPlatformInternalServerTemplatesResponseTemplate apiV1MarketItemItemIdTransferPost(request)


### Example

```typescript
import {
    MarketApi,
    Configuration,
    GithubComNotPlatformInternalModuleMarketsPresentationExternalmarketsModelTransferItemRequest
} from 'notgamessdkpublic';

const configuration = new Configuration();
const apiInstance = new MarketApi(configuration);

let itemId: string; //Item ID (default to undefined)
let request: GithubComNotPlatformInternalModuleMarketsPresentationExternalmarketsModelTransferItemRequest; //request body

const { status, data } = await apiInstance.apiV1MarketItemItemIdTransferPost(
    itemId,
    request
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **request** | **GithubComNotPlatformInternalModuleMarketsPresentationExternalmarketsModelTransferItemRequest**| request body | |
| **itemId** | [**string**] | Item ID | defaults to undefined|


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
|**400** | Bad Request |  -  |
|**404** | Not Found |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **apiV1MarketProfileAccountIdItemsByAppsGet**
> ApiV1MarketProfileAccountIdItemsByAppsGet200Response apiV1MarketProfileAccountIdItemsByAppsGet()


### Example

```typescript
import {
    MarketApi,
    Configuration
} from 'notgamessdkpublic';

const configuration = new Configuration();
const apiInstance = new MarketApi(configuration);

let accountId: string; //account id (default to undefined)

const { status, data } = await apiInstance.apiV1MarketProfileAccountIdItemsByAppsGet(
    accountId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **accountId** | [**string**] | account id | defaults to undefined|


### Return type

**ApiV1MarketProfileAccountIdItemsByAppsGet200Response**

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

# **apiV1MarketProfileAccountIdItemsByDisplayTypesGet**
> ApiV1MarketProfileAccountIdItemsByDisplayTypesGet200Response apiV1MarketProfileAccountIdItemsByDisplayTypesGet()


### Example

```typescript
import {
    MarketApi,
    Configuration
} from 'notgamessdkpublic';

const configuration = new Configuration();
const apiInstance = new MarketApi(configuration);

let accountId: string; //account id (default to undefined)
let appId: string; //app id (default to undefined)

const { status, data } = await apiInstance.apiV1MarketProfileAccountIdItemsByDisplayTypesGet(
    accountId,
    appId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **accountId** | [**string**] | account id | defaults to undefined|
| **appId** | [**string**] | app id | defaults to undefined|


### Return type

**ApiV1MarketProfileAccountIdItemsByDisplayTypesGet200Response**

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

