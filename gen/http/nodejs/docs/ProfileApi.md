# ProfileApi

All URIs are relative to *http://localhost:8080/api/v1*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**profileAccountIdGet**](#profileaccountidget) | **GET** /profile/{account_id} | Get account|

# **profileAccountIdGet**
> ProfileAccountIdGet200Response profileAccountIdGet()


### Example

```typescript
import {
    ProfileApi,
    Configuration
} from 'not-games-sdk-public';

const configuration = new Configuration();
const apiInstance = new ProfileApi(configuration);

let accountId: string; //account id (default to undefined)

const { status, data } = await apiInstance.profileAccountIdGet(
    accountId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **accountId** | [**string**] | account id | defaults to undefined|


### Return type

**ProfileAccountIdGet200Response**

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

