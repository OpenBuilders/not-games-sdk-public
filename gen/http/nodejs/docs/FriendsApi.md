# FriendsApi

All URIs are relative to *http://localhost:8080/api/v1*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**friendsFriendIdDelete**](#friendsfriendiddelete) | **DELETE** /friends/{friend_id} | Remove friend|
|[**friendsGet**](#friendsget) | **GET** /friends | List friends|
|[**friendsPageGet**](#friendspageget) | **GET** /friends/page | Get friends page|
|[**friendsRequestPost**](#friendsrequestpost) | **POST** /friends/request | Send friend request|
|[**friendsRequestRequestIdCancelPut**](#friendsrequestrequestidcancelput) | **PUT** /friends/request/{request_id}/cancel | Cancel outgoing friend request|
|[**friendsRequestRequestIdRespondPut**](#friendsrequestrequestidrespondput) | **PUT** /friends/request/{request_id}/respond | Respond to friend request|
|[**friendsRequestsIncomingGet**](#friendsrequestsincomingget) | **GET** /friends/requests/incoming | List incoming friend requests|
|[**friendsRequestsOutgoingGet**](#friendsrequestsoutgoingget) | **GET** /friends/requests/outgoing | List outgoing friend requests|

# **friendsFriendIdDelete**
> GithubComNotPlatformInternalServerTemplatesResponseTemplate friendsFriendIdDelete()


### Example

```typescript
import {
    FriendsApi,
    Configuration
} from 'notgamessdkpublic';

const configuration = new Configuration();
const apiInstance = new FriendsApi(configuration);

let friendId: number; //friend account id (default to undefined)

const { status, data } = await apiInstance.friendsFriendIdDelete(
    friendId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **friendId** | [**number**] | friend account id | defaults to undefined|


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
|**400** | Bad Request |  -  |
|**401** | Unauthorized |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **friendsGet**
> FriendsGet200Response friendsGet()


### Example

```typescript
import {
    FriendsApi,
    Configuration
} from 'notgamessdkpublic';

const configuration = new Configuration();
const apiInstance = new FriendsApi(configuration);

let search: string; //search by username or name (optional) (default to undefined)
let limit: number; //limit (default 20, max 100) (optional) (default to undefined)
let offset: number; //offset (optional) (default to undefined)

const { status, data } = await apiInstance.friendsGet(
    search,
    limit,
    offset
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **search** | [**string**] | search by username or name | (optional) defaults to undefined|
| **limit** | [**number**] | limit (default 20, max 100) | (optional) defaults to undefined|
| **offset** | [**number**] | offset | (optional) defaults to undefined|


### Return type

**FriendsGet200Response**

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **friendsPageGet**
> FriendsPageGet200Response friendsPageGet()


### Example

```typescript
import {
    FriendsApi,
    Configuration
} from 'notgamessdkpublic';

const configuration = new Configuration();
const apiInstance = new FriendsApi(configuration);

const { status, data } = await apiInstance.friendsPageGet();
```

### Parameters
This endpoint does not have any parameters.


### Return type

**FriendsPageGet200Response**

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **friendsRequestPost**
> FriendsRequestPost200Response friendsRequestPost(request)


### Example

```typescript
import {
    FriendsApi,
    Configuration,
    InternalModuleFriendsPresentationHttpSendRequestBody
} from 'notgamessdkpublic';

const configuration = new Configuration();
const apiInstance = new FriendsApi(configuration);

let request: InternalModuleFriendsPresentationHttpSendRequestBody; //request body

const { status, data } = await apiInstance.friendsRequestPost(
    request
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **request** | **InternalModuleFriendsPresentationHttpSendRequestBody**| request body | |


### Return type

**FriendsRequestPost200Response**

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
|**401** | Unauthorized |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **friendsRequestRequestIdCancelPut**
> GithubComNotPlatformInternalServerTemplatesResponseTemplate friendsRequestRequestIdCancelPut()


### Example

```typescript
import {
    FriendsApi,
    Configuration
} from 'notgamessdkpublic';

const configuration = new Configuration();
const apiInstance = new FriendsApi(configuration);

let requestId: number; //request id (default to undefined)

const { status, data } = await apiInstance.friendsRequestRequestIdCancelPut(
    requestId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **requestId** | [**number**] | request id | defaults to undefined|


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
|**400** | Bad Request |  -  |
|**401** | Unauthorized |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **friendsRequestRequestIdRespondPut**
> GithubComNotPlatformInternalServerTemplatesResponseTemplate friendsRequestRequestIdRespondPut(request)


### Example

```typescript
import {
    FriendsApi,
    Configuration,
    InternalModuleFriendsPresentationHttpRespondRequestBody
} from 'notgamessdkpublic';

const configuration = new Configuration();
const apiInstance = new FriendsApi(configuration);

let requestId: number; //request id (default to undefined)
let request: InternalModuleFriendsPresentationHttpRespondRequestBody; //request body

const { status, data } = await apiInstance.friendsRequestRequestIdRespondPut(
    requestId,
    request
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **request** | **InternalModuleFriendsPresentationHttpRespondRequestBody**| request body | |
| **requestId** | [**number**] | request id | defaults to undefined|


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
|**401** | Unauthorized |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **friendsRequestsIncomingGet**
> FriendsRequestsIncomingGet200Response friendsRequestsIncomingGet()


### Example

```typescript
import {
    FriendsApi,
    Configuration
} from 'notgamessdkpublic';

const configuration = new Configuration();
const apiInstance = new FriendsApi(configuration);

let limit: number; //limit (default 20, max 100) (optional) (default to undefined)
let offset: number; //offset (optional) (default to undefined)

const { status, data } = await apiInstance.friendsRequestsIncomingGet(
    limit,
    offset
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **limit** | [**number**] | limit (default 20, max 100) | (optional) defaults to undefined|
| **offset** | [**number**] | offset | (optional) defaults to undefined|


### Return type

**FriendsRequestsIncomingGet200Response**

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **friendsRequestsOutgoingGet**
> FriendsRequestsIncomingGet200Response friendsRequestsOutgoingGet()


### Example

```typescript
import {
    FriendsApi,
    Configuration
} from 'notgamessdkpublic';

const configuration = new Configuration();
const apiInstance = new FriendsApi(configuration);

let limit: number; //limit (default 20, max 100) (optional) (default to undefined)
let offset: number; //offset (optional) (default to undefined)

const { status, data } = await apiInstance.friendsRequestsOutgoingGet(
    limit,
    offset
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **limit** | [**number**] | limit (default 20, max 100) | (optional) defaults to undefined|
| **offset** | [**number**] | offset | (optional) defaults to undefined|


### Return type

**FriendsRequestsIncomingGet200Response**

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

