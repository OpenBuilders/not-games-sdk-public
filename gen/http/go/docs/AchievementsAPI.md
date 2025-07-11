# \AchievementsAPI

All URIs are relative to *http://localhost:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AchievementsGet**](AchievementsAPI.md#AchievementsGet) | **Get** /achievements | List achievements
[**AchievementsIdGet**](AchievementsAPI.md#AchievementsIdGet) | **Get** /achievements/{id} | Get achievement
[**AchievementsPost**](AchievementsAPI.md#AchievementsPost) | **Post** /achievements | Create achievement



## AchievementsGet

> AchievementsGet200Response AchievementsGet(ctx).AppId(appId).Limit(limit).Offset(offset).Execute()

List achievements

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
	appId := "appId_example" // string | app id
	limit := int32(56) // int32 | limit
	offset := int32(56) // int32 | offset

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AchievementsAPI.AchievementsGet(context.Background()).AppId(appId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AchievementsAPI.AchievementsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AchievementsGet`: AchievementsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AchievementsAPI.AchievementsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAchievementsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appId** | **string** | app id | 
 **limit** | **int32** | limit | 
 **offset** | **int32** | offset | 

### Return type

[**AchievementsGet200Response**](AchievementsGet200Response.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AchievementsIdGet

> AchievementsIdGet200Response AchievementsIdGet(ctx, id).Execute()

Get achievement

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
	id := "id_example" // string | achievement id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AchievementsAPI.AchievementsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AchievementsAPI.AchievementsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AchievementsIdGet`: AchievementsIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AchievementsAPI.AchievementsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | achievement id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAchievementsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AchievementsIdGet200Response**](AchievementsIdGet200Response.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AchievementsPost

> AchievementsPost200Response AchievementsPost(ctx).Request(request).Execute()

Create achievement

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
	request := *openapiclient.NewInternalModuleAchievementPresentationHttpCreateAchievementsRequest() // InternalModuleAchievementPresentationHttpCreateAchievementsRequest | request body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AchievementsAPI.AchievementsPost(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AchievementsAPI.AchievementsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AchievementsPost`: AchievementsPost200Response
	fmt.Fprintf(os.Stdout, "Response from `AchievementsAPI.AchievementsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAchievementsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**InternalModuleAchievementPresentationHttpCreateAchievementsRequest**](InternalModuleAchievementPresentationHttpCreateAchievementsRequest.md) | request body | 

### Return type

[**AchievementsPost200Response**](AchievementsPost200Response.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

