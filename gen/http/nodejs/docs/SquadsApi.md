# SquadsApi

All URIs are relative to *http://localhost:8080/api/v1*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**squadsGet**](#squadsget) | **GET** /squads | List squads|
|[**squadsIdGet**](#squadsidget) | **GET** /squads/{id} | Get squad by id|

# **squadsGet**
> SquadsGet200Response squadsGet()


### Example

```typescript
import {
    SquadsApi,
    Configuration
} from 'notgamessdkpublic';

const configuration = new Configuration();
const apiInstance = new SquadsApi(configuration);

let limit: number; //Limit (default to undefined)
let offset: number; //Offset (default to undefined)
let name: string; //Name (optional) (default to undefined)
let slug: string; //Slug (optional) (default to undefined)
let chatId: number; //Telegram chat ID (optional) (default to undefined)
let prefix: string; //Squad prefix (optional) (default to undefined)
let membersAmount: number; //Members amount (optional) (default to undefined)
let membersAmountMin: number; //Members amount min (optional) (default to undefined)
let membersAmountMax: number; //Members amount max (optional) (default to undefined)
let createdAtGte: string; //Created at greater than or equal, RFC3339 (optional) (default to undefined)
let createdAtLte: string; //Created at less than or equal, RFC3339 (optional) (default to undefined)
let sortFields: string; //Sort fields (optional) (default to undefined)
let directions: string; //Sort directions (optional) (default to undefined)

const { status, data } = await apiInstance.squadsGet(
    limit,
    offset,
    name,
    slug,
    chatId,
    prefix,
    membersAmount,
    membersAmountMin,
    membersAmountMax,
    createdAtGte,
    createdAtLte,
    sortFields,
    directions
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **limit** | [**number**] | Limit | defaults to undefined|
| **offset** | [**number**] | Offset | defaults to undefined|
| **name** | [**string**] | Name | (optional) defaults to undefined|
| **slug** | [**string**] | Slug | (optional) defaults to undefined|
| **chatId** | [**number**] | Telegram chat ID | (optional) defaults to undefined|
| **prefix** | [**string**] | Squad prefix | (optional) defaults to undefined|
| **membersAmount** | [**number**] | Members amount | (optional) defaults to undefined|
| **membersAmountMin** | [**number**] | Members amount min | (optional) defaults to undefined|
| **membersAmountMax** | [**number**] | Members amount max | (optional) defaults to undefined|
| **createdAtGte** | [**string**] | Created at greater than or equal, RFC3339 | (optional) defaults to undefined|
| **createdAtLte** | [**string**] | Created at less than or equal, RFC3339 | (optional) defaults to undefined|
| **sortFields** | [**string**] | Sort fields | (optional) defaults to undefined|
| **directions** | [**string**] | Sort directions | (optional) defaults to undefined|


### Return type

**SquadsGet200Response**

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

# **squadsIdGet**
> SquadsIdGet200Response squadsIdGet()


### Example

```typescript
import {
    SquadsApi,
    Configuration
} from 'notgamessdkpublic';

const configuration = new Configuration();
const apiInstance = new SquadsApi(configuration);

let id: number; //Squad ID (default to undefined)

const { status, data } = await apiInstance.squadsIdGet(
    id
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **id** | [**number**] | Squad ID | defaults to undefined|


### Return type

**SquadsIdGet200Response**

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |
|**401** | Unauthorized |  -  |
|**403** | Forbidden |  -  |
|**404** | Not Found |  -  |
|**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

