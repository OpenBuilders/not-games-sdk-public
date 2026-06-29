# \InventoryAPI

All URIs are relative to *http://localhost:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InventoryByItemdefGet**](InventoryAPI.md#InventoryByItemdefGet) | **Get** /inventory/by_itemdef | List items
[**InventoryItemsGeneratePost**](InventoryAPI.md#InventoryItemsGeneratePost) | **Post** /inventory/items/generate | Generate item
[**InventoryItemsGet**](InventoryAPI.md#InventoryItemsGet) | **Get** /inventory/items | List items
[**InventoryItemsItemIdConsumePost**](InventoryAPI.md#InventoryItemsItemIdConsumePost) | **Post** /inventory/items/{item_id}/consume | Consume item
[**InventoryItemsItemIdGet**](InventoryAPI.md#InventoryItemsItemIdGet) | **Get** /inventory/items/{item_id} | Get item
[**InventoryItemsItemIdIncrementQuantityPost**](InventoryAPI.md#InventoryItemsItemIdIncrementQuantityPost) | **Post** /inventory/items/{item_id}/increment-quantity | Consume item
[**InventoryItemsItemIdPut**](InventoryAPI.md#InventoryItemsItemIdPut) | **Put** /inventory/items/{item_id} | Update item
[**InventoryItemsItemIdTransferLogPost**](InventoryAPI.md#InventoryItemsItemIdTransferLogPost) | **Post** /inventory/items/{item_id}/transfer_log | Get transfer logs
[**InventoryItemsItemIdTransferPost**](InventoryAPI.md#InventoryItemsItemIdTransferPost) | **Post** /inventory/items/{item_id}/transfer | Transfer item to another account
[**InventoryItemsItemIdTransferPreflightPost**](InventoryAPI.md#InventoryItemsItemIdTransferPreflightPost) | **Post** /inventory/items/{item_id}/transfer-preflight | Check transfer item possibility to another account
[**InventoryItemsUnpackPost**](InventoryAPI.md#InventoryItemsUnpackPost) | **Post** /inventory/items/unpack | Unpack item



## InventoryByItemdefGet

> InventoryByItemdefGet200Response InventoryByItemdefGet(ctx).Limit(limit).Offset(offset).ItemDefId(itemDefId).SortFields(sortFields).Directions(directions).Execute()

List items

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
	limit := int32(56) // int32 | limit
	offset := int32(56) // int32 | offset
	itemDefId := "itemDefId_example" // string | item_def_id
	sortFields := "sortFields_example" // string | comma separated sort fields (optional)
	directions := "directions_example" // string | comma separated sort directions (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryAPI.InventoryByItemdefGet(context.Background()).Limit(limit).Offset(offset).ItemDefId(itemDefId).SortFields(sortFields).Directions(directions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.InventoryByItemdefGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryByItemdefGet`: InventoryByItemdefGet200Response
	fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.InventoryByItemdefGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInventoryByItemdefGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | limit | 
 **offset** | **int32** | offset | 
 **itemDefId** | **string** | item_def_id | 
 **sortFields** | **string** | comma separated sort fields | 
 **directions** | **string** | comma separated sort directions | 

### Return type

[**InventoryByItemdefGet200Response**](InventoryByItemdefGet200Response.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryItemsGeneratePost

> InventoryItemsGeneratePost200Response InventoryItemsGeneratePost(ctx).Request(request).Execute()

Generate item

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
	request := *openapiclient.NewGithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelGenerateItemRequest() // GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelGenerateItemRequest | request body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryAPI.InventoryItemsGeneratePost(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.InventoryItemsGeneratePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryItemsGeneratePost`: InventoryItemsGeneratePost200Response
	fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.InventoryItemsGeneratePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInventoryItemsGeneratePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelGenerateItemRequest**](GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelGenerateItemRequest.md) | request body | 

### Return type

[**InventoryItemsGeneratePost200Response**](InventoryItemsGeneratePost200Response.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryItemsGet

> InventoryByItemdefGet200Response InventoryItemsGet(ctx).Limit(limit).Offset(offset).Execute()

List items

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
	limit := int32(56) // int32 | limit
	offset := int32(56) // int32 | offset

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryAPI.InventoryItemsGet(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.InventoryItemsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryItemsGet`: InventoryByItemdefGet200Response
	fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.InventoryItemsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInventoryItemsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | limit | 
 **offset** | **int32** | offset | 

### Return type

[**InventoryByItemdefGet200Response**](InventoryByItemdefGet200Response.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryItemsItemIdConsumePost

> GithubComNotPlatformInternalServerTemplatesResponseTemplate InventoryItemsItemIdConsumePost(ctx, itemId).AccountId(accountId).Request(request).Execute()

Consume item

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
	itemId := "itemId_example" // string | item id (uuid)
	accountId := "accountId_example" // string | account id (int64)
	request := *openapiclient.NewGithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelConsumeItemRequest() // GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelConsumeItemRequest | request body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryAPI.InventoryItemsItemIdConsumePost(context.Background(), itemId).AccountId(accountId).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.InventoryItemsItemIdConsumePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryItemsItemIdConsumePost`: GithubComNotPlatformInternalServerTemplatesResponseTemplate
	fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.InventoryItemsItemIdConsumePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | item id (uuid) | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryItemsItemIdConsumePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **string** | account id (int64) | 
 **request** | [**GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelConsumeItemRequest**](GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelConsumeItemRequest.md) | request body | 

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


## InventoryItemsItemIdGet

> InventoryItemsItemIdGet200Response InventoryItemsItemIdGet(ctx, itemId).AppId(appId).Execute()

Get item

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
	itemId := "itemId_example" // string | item id (uuid)
	appId := int32(56) // int32 | app id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryAPI.InventoryItemsItemIdGet(context.Background(), itemId).AppId(appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.InventoryItemsItemIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryItemsItemIdGet`: InventoryItemsItemIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.InventoryItemsItemIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | item id (uuid) | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryItemsItemIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appId** | **int32** | app id | 

### Return type

[**InventoryItemsItemIdGet200Response**](InventoryItemsItemIdGet200Response.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryItemsItemIdIncrementQuantityPost

> GithubComNotPlatformInternalServerTemplatesResponseTemplate InventoryItemsItemIdIncrementQuantityPost(ctx, itemId).AccountId(accountId).Request(request).Execute()

Consume item

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
	itemId := "itemId_example" // string | item id (uuid)
	accountId := "accountId_example" // string | account id (int64)
	request := *openapiclient.NewGithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelConsumeItemRequest() // GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelConsumeItemRequest | request body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryAPI.InventoryItemsItemIdIncrementQuantityPost(context.Background(), itemId).AccountId(accountId).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.InventoryItemsItemIdIncrementQuantityPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryItemsItemIdIncrementQuantityPost`: GithubComNotPlatformInternalServerTemplatesResponseTemplate
	fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.InventoryItemsItemIdIncrementQuantityPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | item id (uuid) | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryItemsItemIdIncrementQuantityPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **string** | account id (int64) | 
 **request** | [**GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelConsumeItemRequest**](GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelConsumeItemRequest.md) | request body | 

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


## InventoryItemsItemIdPut

> GithubComNotPlatformInternalServerTemplatesResponseTemplate InventoryItemsItemIdPut(ctx, itemId).Request(request).Execute()

Update item

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
	itemId := "itemId_example" // string | item id (uuid)
	request := *openapiclient.NewGithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelUpdateItemRequest() // GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelUpdateItemRequest | request body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryAPI.InventoryItemsItemIdPut(context.Background(), itemId).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.InventoryItemsItemIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryItemsItemIdPut`: GithubComNotPlatformInternalServerTemplatesResponseTemplate
	fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.InventoryItemsItemIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | item id (uuid) | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryItemsItemIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelUpdateItemRequest**](GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelUpdateItemRequest.md) | request body | 

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


## InventoryItemsItemIdTransferLogPost

> GithubComNotPlatformInternalServerTemplatesResponseTemplate InventoryItemsItemIdTransferLogPost(ctx, itemId).Limit(limit).Offset(offset).Execute()

Get transfer logs

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
	itemId := "itemId_example" // string | item id (uuid)
	limit := int32(56) // int32 | limit
	offset := int32(56) // int32 | offset

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryAPI.InventoryItemsItemIdTransferLogPost(context.Background(), itemId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.InventoryItemsItemIdTransferLogPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryItemsItemIdTransferLogPost`: GithubComNotPlatformInternalServerTemplatesResponseTemplate
	fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.InventoryItemsItemIdTransferLogPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | item id (uuid) | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryItemsItemIdTransferLogPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | limit | 
 **offset** | **int32** | offset | 

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


## InventoryItemsItemIdTransferPost

> GithubComNotPlatformInternalServerTemplatesResponseTemplate InventoryItemsItemIdTransferPost(ctx, itemId).Request(request).Execute()

Transfer item to another account

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
	itemId := "itemId_example" // string | item id (uuid)
	request := *openapiclient.NewGithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelTransferItemRequest() // GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelTransferItemRequest | request body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryAPI.InventoryItemsItemIdTransferPost(context.Background(), itemId).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.InventoryItemsItemIdTransferPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryItemsItemIdTransferPost`: GithubComNotPlatformInternalServerTemplatesResponseTemplate
	fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.InventoryItemsItemIdTransferPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | item id (uuid) | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryItemsItemIdTransferPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelTransferItemRequest**](GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelTransferItemRequest.md) | request body | 

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


## InventoryItemsItemIdTransferPreflightPost

> InventoryItemsItemIdTransferPreflightPost200Response InventoryItemsItemIdTransferPreflightPost(ctx, itemId).Request(request).Execute()

Check transfer item possibility to another account

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
	itemId := "itemId_example" // string | item id (uuid)
	request := *openapiclient.NewGithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelTransferItemRequest() // GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelTransferItemRequest | request body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryAPI.InventoryItemsItemIdTransferPreflightPost(context.Background(), itemId).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.InventoryItemsItemIdTransferPreflightPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryItemsItemIdTransferPreflightPost`: InventoryItemsItemIdTransferPreflightPost200Response
	fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.InventoryItemsItemIdTransferPreflightPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | item id (uuid) | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryItemsItemIdTransferPreflightPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelTransferItemRequest**](GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelTransferItemRequest.md) | request body | 

### Return type

[**InventoryItemsItemIdTransferPreflightPost200Response**](InventoryItemsItemIdTransferPreflightPost200Response.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryItemsUnpackPost

> InventoryItemsUnpackPost200Response InventoryItemsUnpackPost(ctx).Request(request).Execute()

Unpack item

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
	request := *openapiclient.NewGithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelUnpackItemRequest() // GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelUnpackItemRequest | request body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryAPI.InventoryItemsUnpackPost(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.InventoryItemsUnpackPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryItemsUnpackPost`: InventoryItemsUnpackPost200Response
	fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.InventoryItemsUnpackPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInventoryItemsUnpackPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelUnpackItemRequest**](GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelUnpackItemRequest.md) | request body | 

### Return type

[**InventoryItemsUnpackPost200Response**](InventoryItemsUnpackPost200Response.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

