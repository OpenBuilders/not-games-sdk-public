# \MarketAPI

All URIs are relative to *http://localhost:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1MarketAppsNewGet**](MarketAPI.md#ApiV1MarketAppsNewGet) | **Get** /api/v1/market/apps/new | List new visible apps for external market
[**ApiV1MarketProfileAccountIdItemsByAppsGet**](MarketAPI.md#ApiV1MarketProfileAccountIdItemsByAppsGet) | **Get** /api/v1/market/profile/{account_id}/items-by-apps | market profile apps items
[**ApiV1ProfileAccountIdItemsByDisplayTypesGet**](MarketAPI.md#ApiV1ProfileAccountIdItemsByDisplayTypesGet) | **Get** /api/v1/profile/{account_id}/items-by-display-types | market profile app items



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


## ApiV1ProfileAccountIdItemsByDisplayTypesGet

> ApiV1ProfileAccountIdItemsByDisplayTypesGet200Response ApiV1ProfileAccountIdItemsByDisplayTypesGet(ctx, accountId).AppId(appId).Execute()

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
	resp, r, err := apiClient.MarketAPI.ApiV1ProfileAccountIdItemsByDisplayTypesGet(context.Background(), accountId).AppId(appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketAPI.ApiV1ProfileAccountIdItemsByDisplayTypesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1ProfileAccountIdItemsByDisplayTypesGet`: ApiV1ProfileAccountIdItemsByDisplayTypesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MarketAPI.ApiV1ProfileAccountIdItemsByDisplayTypesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1ProfileAccountIdItemsByDisplayTypesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appId** | **string** | app id | 

### Return type

[**ApiV1ProfileAccountIdItemsByDisplayTypesGet200Response**](ApiV1ProfileAccountIdItemsByDisplayTypesGet200Response.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

