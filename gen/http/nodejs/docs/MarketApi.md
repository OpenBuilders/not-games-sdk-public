# MarketApi

All URIs are relative to *http://localhost:8080/api/v1*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**apiV1MarketAppsNewGet**](#apiv1marketappsnewget) | **GET** /api/v1/market/apps/new | List new visible apps for external market|
|[**apiV1MarketProfileAccountIdItemsByAppsGet**](#apiv1marketprofileaccountiditemsbyappsget) | **GET** /api/v1/market/profile/{account_id}/items-by-apps | market profile apps items|
|[**apiV1ProfileAccountIdItemsByDisplayTypesGet**](#apiv1profileaccountiditemsbydisplaytypesget) | **GET** /api/v1/profile/{account_id}/items-by-display-types | market profile app items|

# **apiV1MarketAppsNewGet**
> ApiV1MarketAppsNewGet200Response apiV1MarketAppsNewGet()


### Example

```typescript
import {
    MarketApi,
    Configuration
} from 'not-games-sdk-public';

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

# **apiV1MarketProfileAccountIdItemsByAppsGet**
> ApiV1MarketProfileAccountIdItemsByAppsGet200Response apiV1MarketProfileAccountIdItemsByAppsGet()


### Example

```typescript
import {
    MarketApi,
    Configuration
} from 'not-games-sdk-public';

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

# **apiV1ProfileAccountIdItemsByDisplayTypesGet**
> ApiV1ProfileAccountIdItemsByDisplayTypesGet200Response apiV1ProfileAccountIdItemsByDisplayTypesGet()


### Example

```typescript
import {
    MarketApi,
    Configuration
} from 'not-games-sdk-public';

const configuration = new Configuration();
const apiInstance = new MarketApi(configuration);

let accountId: string; //account id (default to undefined)
let appId: string; //app id (default to undefined)

const { status, data } = await apiInstance.apiV1ProfileAccountIdItemsByDisplayTypesGet(
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

**ApiV1ProfileAccountIdItemsByDisplayTypesGet200Response**

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

