# \MarketAPI

All URIs are relative to *http://localhost:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1MarketAppsNewGet**](MarketAPI.md#ApiV1MarketAppsNewGet) | **Get** /api/v1/market/apps/new | List new visible apps for external market
[**ApiV1MarketItemDefGet**](MarketAPI.md#ApiV1MarketItemDefGet) | **Get** /api/v1/market/item_def | List item defs for external market
[**ApiV1MarketItemGet**](MarketAPI.md#ApiV1MarketItemGet) | **Get** /api/v1/market/item | List items for external market
[**ApiV1MarketItemItemIdGet**](MarketAPI.md#ApiV1MarketItemItemIdGet) | **Get** /api/v1/market/item/{item_id} | Get any user item by id for external market
[**ApiV1MarketItemItemIdTransferPost**](MarketAPI.md#ApiV1MarketItemItemIdTransferPost) | **Post** /api/v1/market/item/{item_id}/transfer | Transfer an item to a user from external market
[**ApiV1MarketProfileAccountIdItemsByAppsGet**](MarketAPI.md#ApiV1MarketProfileAccountIdItemsByAppsGet) | **Get** /api/v1/market/profile/{account_id}/items-by-apps | market profile apps items
[**ApiV1MarketProfileAccountIdItemsByDisplayTypesGet**](MarketAPI.md#ApiV1MarketProfileAccountIdItemsByDisplayTypesGet) | **Get** /api/v1/market/profile/{account_id}/items-by-display-types | market profile app items



## ApiV1MarketAppsNewGet

> ApiV1MarketAppsNewGet200Response ApiV1MarketAppsNewGet(ctx).FromDate(fromDate).Limit(limit).SortFields(sortFields).Directions(directions).Execute()

List new visible apps for external market

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
	fromDate := "fromDate_example" // string | from date, RFC3339 or YYYY-MM-DD (optional)
	limit := int32(56) // int32 | limit, max 500 (optional)
	sortFields := "sortFields_example" // string | comma separated: created_at|name (optional)
	directions := "directions_example" // string | comma separated: asc|desc (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketAPI.ApiV1MarketAppsNewGet(context.Background()).FromDate(fromDate).Limit(limit).SortFields(sortFields).Directions(directions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketAPI.ApiV1MarketAppsNewGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1MarketAppsNewGet`: ApiV1MarketAppsNewGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MarketAPI.ApiV1MarketAppsNewGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1MarketAppsNewGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromDate** | **string** | from date, RFC3339 or YYYY-MM-DD | 
 **limit** | **int32** | limit, max 500 | 
 **sortFields** | **string** | comma separated: created_at|name | 
 **directions** | **string** | comma separated: asc|desc | 

### Return type

[**ApiV1MarketAppsNewGet200Response**](ApiV1MarketAppsNewGet200Response.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1MarketItemDefGet

> ApiV1MarketItemDefGet200Response ApiV1MarketItemDefGet(ctx).Limit(limit).Offset(offset).SortFields(sortFields).Directions(directions).Execute()

List item defs for external market

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
	limit := int32(56) // int32 | limit (optional)
	offset := int32(56) // int32 | offset (optional)
	sortFields := "sortFields_example" // string | comma separated sort fields (optional)
	directions := "directions_example" // string | comma separated: asc|desc (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketAPI.ApiV1MarketItemDefGet(context.Background()).Limit(limit).Offset(offset).SortFields(sortFields).Directions(directions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketAPI.ApiV1MarketItemDefGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1MarketItemDefGet`: ApiV1MarketItemDefGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MarketAPI.ApiV1MarketItemDefGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1MarketItemDefGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | limit | 
 **offset** | **int32** | offset | 
 **sortFields** | **string** | comma separated sort fields | 
 **directions** | **string** | comma separated: asc|desc | 

### Return type

[**ApiV1MarketItemDefGet200Response**](ApiV1MarketItemDefGet200Response.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1MarketItemGet

> ApiV1MarketItemGet200Response ApiV1MarketItemGet(ctx).Limit(limit).Offset(offset).SortFields(sortFields).Directions(directions).Execute()

List items for external market

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
	limit := int32(56) // int32 | limit (optional)
	offset := int32(56) // int32 | offset (optional)
	sortFields := "sortFields_example" // string | comma separated sort fields (optional)
	directions := "directions_example" // string | comma separated: asc|desc (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketAPI.ApiV1MarketItemGet(context.Background()).Limit(limit).Offset(offset).SortFields(sortFields).Directions(directions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketAPI.ApiV1MarketItemGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1MarketItemGet`: ApiV1MarketItemGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MarketAPI.ApiV1MarketItemGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1MarketItemGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | limit | 
 **offset** | **int32** | offset | 
 **sortFields** | **string** | comma separated sort fields | 
 **directions** | **string** | comma separated: asc|desc | 

### Return type

[**ApiV1MarketItemGet200Response**](ApiV1MarketItemGet200Response.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1MarketItemItemIdGet

> ApiV1MarketItemItemIdGet200Response ApiV1MarketItemItemIdGet(ctx, itemId).Execute()

Get any user item by id for external market

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
	itemId := "itemId_example" // string | Item ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketAPI.ApiV1MarketItemItemIdGet(context.Background(), itemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketAPI.ApiV1MarketItemItemIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1MarketItemItemIdGet`: ApiV1MarketItemItemIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MarketAPI.ApiV1MarketItemItemIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | Item ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1MarketItemItemIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiV1MarketItemItemIdGet200Response**](ApiV1MarketItemItemIdGet200Response.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1MarketItemItemIdTransferPost

> GithubComNotPlatformInternalServerTemplatesResponseTemplate ApiV1MarketItemItemIdTransferPost(ctx, itemId).Request(request).Execute()

Transfer an item to a user from external market

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
	itemId := "itemId_example" // string | Item ID
	request := *openapiclient.NewGithubComNotPlatformInternalModuleMarketsPresentationExternalmarketsModelTransferItemRequest() // GithubComNotPlatformInternalModuleMarketsPresentationExternalmarketsModelTransferItemRequest | request body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketAPI.ApiV1MarketItemItemIdTransferPost(context.Background(), itemId).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketAPI.ApiV1MarketItemItemIdTransferPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1MarketItemItemIdTransferPost`: GithubComNotPlatformInternalServerTemplatesResponseTemplate
	fmt.Fprintf(os.Stdout, "Response from `MarketAPI.ApiV1MarketItemItemIdTransferPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | Item ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1MarketItemItemIdTransferPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**GithubComNotPlatformInternalModuleMarketsPresentationExternalmarketsModelTransferItemRequest**](GithubComNotPlatformInternalModuleMarketsPresentationExternalmarketsModelTransferItemRequest.md) | request body | 

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


## ApiV1MarketProfileAccountIdItemsByAppsGet

> ApiV1MarketProfileAccountIdItemsByAppsGet200Response ApiV1MarketProfileAccountIdItemsByAppsGet(ctx, accountId).Execute()

market profile apps items

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
	accountId := "accountId_example" // string | account id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketAPI.ApiV1MarketProfileAccountIdItemsByAppsGet(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketAPI.ApiV1MarketProfileAccountIdItemsByAppsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1MarketProfileAccountIdItemsByAppsGet`: ApiV1MarketProfileAccountIdItemsByAppsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MarketAPI.ApiV1MarketProfileAccountIdItemsByAppsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1MarketProfileAccountIdItemsByAppsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiV1MarketProfileAccountIdItemsByAppsGet200Response**](ApiV1MarketProfileAccountIdItemsByAppsGet200Response.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1MarketProfileAccountIdItemsByDisplayTypesGet

> ApiV1MarketProfileAccountIdItemsByDisplayTypesGet200Response ApiV1MarketProfileAccountIdItemsByDisplayTypesGet(ctx, accountId).AppId(appId).Execute()

market profile app items

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
	accountId := "accountId_example" // string | account id
	appId := "appId_example" // string | app id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketAPI.ApiV1MarketProfileAccountIdItemsByDisplayTypesGet(context.Background(), accountId).AppId(appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketAPI.ApiV1MarketProfileAccountIdItemsByDisplayTypesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1MarketProfileAccountIdItemsByDisplayTypesGet`: ApiV1MarketProfileAccountIdItemsByDisplayTypesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MarketAPI.ApiV1MarketProfileAccountIdItemsByDisplayTypesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1MarketProfileAccountIdItemsByDisplayTypesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appId** | **string** | app id | 

### Return type

[**ApiV1MarketProfileAccountIdItemsByDisplayTypesGet200Response**](ApiV1MarketProfileAccountIdItemsByDisplayTypesGet200Response.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

