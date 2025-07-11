# \DropsAPI

All URIs are relative to *http://localhost:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DropsIdGet**](DropsAPI.md#DropsIdGet) | **Get** /drops/{id} | Check user eligibility for a drop
[**DropsIdWhitelistPost**](DropsAPI.md#DropsIdWhitelistPost) | **Post** /drops/{id}/whitelist | Whitelist users for a drop
[**DropsPost**](DropsAPI.md#DropsPost) | **Post** /drops | Create a new drop



## DropsIdGet

> DropsIdGet200Response DropsIdGet(ctx, id).Execute()

Check user eligibility for a drop



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
	id := "id_example" // string | Drop ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DropsAPI.DropsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DropsAPI.DropsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DropsIdGet`: DropsIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DropsAPI.DropsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Drop ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDropsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DropsIdGet200Response**](DropsIdGet200Response.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DropsIdWhitelistPost

> GithubComNotPlatformInternalServerTemplatesResponseTemplate DropsIdWhitelistPost(ctx, id).Request(request).Execute()

Whitelist users for a drop

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
	id := "id_example" // string | Drop ID
	request := *openapiclient.NewInternalModuleDropsPresentationDropsHttpWhitelistDropUserRequest() // InternalModuleDropsPresentationDropsHttpWhitelistDropUserRequest | Whitelist Drop User Request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DropsAPI.DropsIdWhitelistPost(context.Background(), id).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DropsAPI.DropsIdWhitelistPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DropsIdWhitelistPost`: GithubComNotPlatformInternalServerTemplatesResponseTemplate
	fmt.Fprintf(os.Stdout, "Response from `DropsAPI.DropsIdWhitelistPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Drop ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDropsIdWhitelistPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**InternalModuleDropsPresentationDropsHttpWhitelistDropUserRequest**](InternalModuleDropsPresentationDropsHttpWhitelistDropUserRequest.md) | Whitelist Drop User Request | 

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


## DropsPost

> DropsPost200Response DropsPost(ctx).Request(request).Execute()

Create a new drop



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
	request := *openapiclient.NewInternalModuleDropsPresentationDropsHttpCreateDropRequest() // InternalModuleDropsPresentationDropsHttpCreateDropRequest | Create Drop Request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DropsAPI.DropsPost(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DropsAPI.DropsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DropsPost`: DropsPost200Response
	fmt.Fprintf(os.Stdout, "Response from `DropsAPI.DropsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDropsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**InternalModuleDropsPresentationDropsHttpCreateDropRequest**](InternalModuleDropsPresentationDropsHttpCreateDropRequest.md) | Create Drop Request | 

### Return type

[**DropsPost200Response**](DropsPost200Response.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

