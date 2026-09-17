# \SquadsAPI

All URIs are relative to *http://localhost:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SquadsGet**](SquadsAPI.md#SquadsGet) | **Get** /squads | List squads
[**SquadsIdGet**](SquadsAPI.md#SquadsIdGet) | **Get** /squads/{id} | Get squad by id



## SquadsGet

> SquadsGet200Response SquadsGet(ctx).Limit(limit).Offset(offset).Slug(slug).MembersAmountMin(membersAmountMin).SortFields(sortFields).Directions(directions).Execute()

List squads

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
	limit := int32(56) // int32 | Limit
	offset := int32(56) // int32 | Offset
	slug := "slug_example" // string | Slug (optional)
	membersAmountMin := int32(56) // int32 | Members amount min (optional)
	sortFields := "sortFields_example" // string | Sort fields (optional)
	directions := "directions_example" // string | Sort directions (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SquadsAPI.SquadsGet(context.Background()).Limit(limit).Offset(offset).Slug(slug).MembersAmountMin(membersAmountMin).SortFields(sortFields).Directions(directions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SquadsAPI.SquadsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SquadsGet`: SquadsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SquadsAPI.SquadsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSquadsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit | 
 **offset** | **int32** | Offset | 
 **slug** | **string** | Slug | 
 **membersAmountMin** | **int32** | Members amount min | 
 **sortFields** | **string** | Sort fields | 
 **directions** | **string** | Sort directions | 

### Return type

[**SquadsGet200Response**](SquadsGet200Response.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SquadsIdGet

> SquadsIdGet200Response SquadsIdGet(ctx, id).Execute()

Get squad by id

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
	id := int32(56) // int32 | Squad ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SquadsAPI.SquadsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SquadsAPI.SquadsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SquadsIdGet`: SquadsIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SquadsAPI.SquadsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Squad ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSquadsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SquadsIdGet200Response**](SquadsIdGet200Response.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

