# \FriendsAPI

All URIs are relative to *http://localhost:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FriendsFriendIdDelete**](FriendsAPI.md#FriendsFriendIdDelete) | **Delete** /friends/{friend_id} | Remove friend
[**FriendsGet**](FriendsAPI.md#FriendsGet) | **Get** /friends | List friends
[**FriendsPageGet**](FriendsAPI.md#FriendsPageGet) | **Get** /friends/page | Get friends page
[**FriendsRequestPost**](FriendsAPI.md#FriendsRequestPost) | **Post** /friends/request | Send friend request
[**FriendsRequestRequestIdCancelPut**](FriendsAPI.md#FriendsRequestRequestIdCancelPut) | **Put** /friends/request/{request_id}/cancel | Cancel outgoing friend request
[**FriendsRequestRequestIdRespondPut**](FriendsAPI.md#FriendsRequestRequestIdRespondPut) | **Put** /friends/request/{request_id}/respond | Respond to friend request
[**FriendsRequestsIncomingGet**](FriendsAPI.md#FriendsRequestsIncomingGet) | **Get** /friends/requests/incoming | List incoming friend requests
[**FriendsRequestsOutgoingGet**](FriendsAPI.md#FriendsRequestsOutgoingGet) | **Get** /friends/requests/outgoing | List outgoing friend requests



## FriendsFriendIdDelete

> GithubComNotPlatformInternalServerTemplatesResponseTemplate FriendsFriendIdDelete(ctx, friendId).Execute()

Remove friend

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	friendId := int32(56) // int32 | friend account id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FriendsAPI.FriendsFriendIdDelete(context.Background(), friendId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FriendsAPI.FriendsFriendIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FriendsFriendIdDelete`: GithubComNotPlatformInternalServerTemplatesResponseTemplate
	fmt.Fprintf(os.Stdout, "Response from `FriendsAPI.FriendsFriendIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**friendId** | **int32** | friend account id | 

### Other Parameters

Other parameters are passed through a pointer to a apiFriendsFriendIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GithubComNotPlatformInternalServerTemplatesResponseTemplate**](GithubComNotPlatformInternalServerTemplatesResponseTemplate.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FriendsGet

> FriendsGet200Response FriendsGet(ctx).Search(search).Limit(limit).Offset(offset).Execute()

List friends

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	search := "search_example" // string | search by username or name (optional)
	limit := int32(56) // int32 | limit (default 20, max 100) (optional)
	offset := int32(56) // int32 | offset (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FriendsAPI.FriendsGet(context.Background()).Search(search).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FriendsAPI.FriendsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FriendsGet`: FriendsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `FriendsAPI.FriendsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFriendsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | search by username or name | 
 **limit** | **int32** | limit (default 20, max 100) | 
 **offset** | **int32** | offset | 

### Return type

[**FriendsGet200Response**](FriendsGet200Response.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FriendsPageGet

> FriendsPageGet200Response FriendsPageGet(ctx).Execute()

Get friends page

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FriendsAPI.FriendsPageGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FriendsAPI.FriendsPageGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FriendsPageGet`: FriendsPageGet200Response
	fmt.Fprintf(os.Stdout, "Response from `FriendsAPI.FriendsPageGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFriendsPageGetRequest struct via the builder pattern


### Return type

[**FriendsPageGet200Response**](FriendsPageGet200Response.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FriendsRequestPost

> FriendsRequestPost200Response FriendsRequestPost(ctx).Request(request).Execute()

Send friend request

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	request := *openapiclient.NewInternalModuleFriendsPresentationHttpSendRequestBody() // InternalModuleFriendsPresentationHttpSendRequestBody | request body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FriendsAPI.FriendsRequestPost(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FriendsAPI.FriendsRequestPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FriendsRequestPost`: FriendsRequestPost200Response
	fmt.Fprintf(os.Stdout, "Response from `FriendsAPI.FriendsRequestPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFriendsRequestPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**InternalModuleFriendsPresentationHttpSendRequestBody**](InternalModuleFriendsPresentationHttpSendRequestBody.md) | request body | 

### Return type

[**FriendsRequestPost200Response**](FriendsRequestPost200Response.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FriendsRequestRequestIdCancelPut

> GithubComNotPlatformInternalServerTemplatesResponseTemplate FriendsRequestRequestIdCancelPut(ctx, requestId).Execute()

Cancel outgoing friend request

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	requestId := int32(56) // int32 | request id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FriendsAPI.FriendsRequestRequestIdCancelPut(context.Background(), requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FriendsAPI.FriendsRequestRequestIdCancelPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FriendsRequestRequestIdCancelPut`: GithubComNotPlatformInternalServerTemplatesResponseTemplate
	fmt.Fprintf(os.Stdout, "Response from `FriendsAPI.FriendsRequestRequestIdCancelPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **int32** | request id | 

### Other Parameters

Other parameters are passed through a pointer to a apiFriendsRequestRequestIdCancelPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GithubComNotPlatformInternalServerTemplatesResponseTemplate**](GithubComNotPlatformInternalServerTemplatesResponseTemplate.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FriendsRequestRequestIdRespondPut

> GithubComNotPlatformInternalServerTemplatesResponseTemplate FriendsRequestRequestIdRespondPut(ctx, requestId).Request(request).Execute()

Respond to friend request

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	requestId := int32(56) // int32 | request id
	request := *openapiclient.NewInternalModuleFriendsPresentationHttpRespondRequestBody() // InternalModuleFriendsPresentationHttpRespondRequestBody | request body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FriendsAPI.FriendsRequestRequestIdRespondPut(context.Background(), requestId).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FriendsAPI.FriendsRequestRequestIdRespondPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FriendsRequestRequestIdRespondPut`: GithubComNotPlatformInternalServerTemplatesResponseTemplate
	fmt.Fprintf(os.Stdout, "Response from `FriendsAPI.FriendsRequestRequestIdRespondPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **int32** | request id | 

### Other Parameters

Other parameters are passed through a pointer to a apiFriendsRequestRequestIdRespondPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**InternalModuleFriendsPresentationHttpRespondRequestBody**](InternalModuleFriendsPresentationHttpRespondRequestBody.md) | request body | 

### Return type

[**GithubComNotPlatformInternalServerTemplatesResponseTemplate**](GithubComNotPlatformInternalServerTemplatesResponseTemplate.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FriendsRequestsIncomingGet

> FriendsRequestsIncomingGet200Response FriendsRequestsIncomingGet(ctx).Limit(limit).Offset(offset).Execute()

List incoming friend requests

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	limit := int32(56) // int32 | limit (default 20, max 100) (optional)
	offset := int32(56) // int32 | offset (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FriendsAPI.FriendsRequestsIncomingGet(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FriendsAPI.FriendsRequestsIncomingGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FriendsRequestsIncomingGet`: FriendsRequestsIncomingGet200Response
	fmt.Fprintf(os.Stdout, "Response from `FriendsAPI.FriendsRequestsIncomingGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFriendsRequestsIncomingGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | limit (default 20, max 100) | 
 **offset** | **int32** | offset | 

### Return type

[**FriendsRequestsIncomingGet200Response**](FriendsRequestsIncomingGet200Response.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FriendsRequestsOutgoingGet

> FriendsRequestsIncomingGet200Response FriendsRequestsOutgoingGet(ctx).Limit(limit).Offset(offset).Execute()

List outgoing friend requests

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	limit := int32(56) // int32 | limit (default 20, max 100) (optional)
	offset := int32(56) // int32 | offset (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FriendsAPI.FriendsRequestsOutgoingGet(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FriendsAPI.FriendsRequestsOutgoingGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FriendsRequestsOutgoingGet`: FriendsRequestsIncomingGet200Response
	fmt.Fprintf(os.Stdout, "Response from `FriendsAPI.FriendsRequestsOutgoingGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFriendsRequestsOutgoingGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | limit (default 20, max 100) | 
 **offset** | **int32** | offset | 

### Return type

[**FriendsRequestsIncomingGet200Response**](FriendsRequestsIncomingGet200Response.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

