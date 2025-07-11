# \RegistryAPI

All URIs are relative to *http://localhost:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RegistryAchievementsAchievementDefIdDelete**](RegistryAPI.md#RegistryAchievementsAchievementDefIdDelete) | **Delete** /registry/achievements/{achievement_def_id} | Delete achievement-def
[**RegistryAchievementsAchievementDefIdGet**](RegistryAPI.md#RegistryAchievementsAchievementDefIdGet) | **Get** /registry/achievements/{achievement_def_id} | Get achievement-def
[**RegistryAchievementsAchievementDefIdPut**](RegistryAPI.md#RegistryAchievementsAchievementDefIdPut) | **Put** /registry/achievements/{achievement_def_id} | Update achievement-def
[**RegistryAchievementsGet**](RegistryAPI.md#RegistryAchievementsGet) | **Get** /registry/achievements | List achievement-def
[**RegistryAchievementsPost**](RegistryAPI.md#RegistryAchievementsPost) | **Post** /registry/achievements | Create achievement-def
[**RegistryAppsGet**](RegistryAPI.md#RegistryAppsGet) | **Get** /registry/apps | List apps
[**RegistryItemsGet**](RegistryAPI.md#RegistryItemsGet) | **Get** /registry/items | List item-def
[**RegistryItemsItemDefIdDelete**](RegistryAPI.md#RegistryItemsItemDefIdDelete) | **Delete** /registry/items/{item_def_id} | Delete item-def
[**RegistryItemsItemDefIdGet**](RegistryAPI.md#RegistryItemsItemDefIdGet) | **Get** /registry/items/{item_def_id} | Get item-def
[**RegistryItemsItemDefIdPut**](RegistryAPI.md#RegistryItemsItemDefIdPut) | **Put** /registry/items/{item_def_id} | Update item-def
[**RegistryItemsPost**](RegistryAPI.md#RegistryItemsPost) | **Post** /registry/items | Create item-def



## RegistryAchievementsAchievementDefIdDelete

> GithubComNotPlatformInternalServerTemplatesResponseTemplate RegistryAchievementsAchievementDefIdDelete(ctx, achievementDefId).Execute()

Delete achievement-def

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
	achievementDefId := "achievementDefId_example" // string | achievement_def_id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegistryAPI.RegistryAchievementsAchievementDefIdDelete(context.Background(), achievementDefId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistryAPI.RegistryAchievementsAchievementDefIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegistryAchievementsAchievementDefIdDelete`: GithubComNotPlatformInternalServerTemplatesResponseTemplate
	fmt.Fprintf(os.Stdout, "Response from `RegistryAPI.RegistryAchievementsAchievementDefIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**achievementDefId** | **string** | achievement_def_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegistryAchievementsAchievementDefIdDeleteRequest struct via the builder pattern


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


## RegistryAchievementsAchievementDefIdGet

> RegistryAchievementsAchievementDefIdGet200Response RegistryAchievementsAchievementDefIdGet(ctx, achievementDefId).Execute()

Get achievement-def

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
	achievementDefId := "achievementDefId_example" // string | achievement_def_id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegistryAPI.RegistryAchievementsAchievementDefIdGet(context.Background(), achievementDefId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistryAPI.RegistryAchievementsAchievementDefIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegistryAchievementsAchievementDefIdGet`: RegistryAchievementsAchievementDefIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `RegistryAPI.RegistryAchievementsAchievementDefIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**achievementDefId** | **string** | achievement_def_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegistryAchievementsAchievementDefIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RegistryAchievementsAchievementDefIdGet200Response**](RegistryAchievementsAchievementDefIdGet200Response.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegistryAchievementsAchievementDefIdPut

> GithubComNotPlatformInternalServerTemplatesResponseTemplate RegistryAchievementsAchievementDefIdPut(ctx, achievementDefId).Request(request).Execute()

Update achievement-def

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
	achievementDefId := "achievementDefId_example" // string | achievement_def_id
	request := *openapiclient.NewInternalModuleRegistryPresentationAchievementHttpUpdateRequest() // InternalModuleRegistryPresentationAchievementHttpUpdateRequest | request body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegistryAPI.RegistryAchievementsAchievementDefIdPut(context.Background(), achievementDefId).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistryAPI.RegistryAchievementsAchievementDefIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegistryAchievementsAchievementDefIdPut`: GithubComNotPlatformInternalServerTemplatesResponseTemplate
	fmt.Fprintf(os.Stdout, "Response from `RegistryAPI.RegistryAchievementsAchievementDefIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**achievementDefId** | **string** | achievement_def_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegistryAchievementsAchievementDefIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**InternalModuleRegistryPresentationAchievementHttpUpdateRequest**](InternalModuleRegistryPresentationAchievementHttpUpdateRequest.md) | request body | 

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


## RegistryAchievementsGet

> RegistryAchievementsGet200Response RegistryAchievementsGet(ctx).Limit(limit).Offset(offset).Execute()

List achievement-def

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
	resp, r, err := apiClient.RegistryAPI.RegistryAchievementsGet(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistryAPI.RegistryAchievementsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegistryAchievementsGet`: RegistryAchievementsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `RegistryAPI.RegistryAchievementsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegistryAchievementsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | limit | 
 **offset** | **int32** | offset | 

### Return type

[**RegistryAchievementsGet200Response**](RegistryAchievementsGet200Response.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegistryAchievementsPost

> GithubComNotPlatformInternalServerTemplatesResponseTemplate RegistryAchievementsPost(ctx).Request(request).Execute()

Create achievement-def

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
	request := *openapiclient.NewInternalModuleRegistryPresentationAchievementHttpCreateRequest() // InternalModuleRegistryPresentationAchievementHttpCreateRequest | request body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegistryAPI.RegistryAchievementsPost(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistryAPI.RegistryAchievementsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegistryAchievementsPost`: GithubComNotPlatformInternalServerTemplatesResponseTemplate
	fmt.Fprintf(os.Stdout, "Response from `RegistryAPI.RegistryAchievementsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegistryAchievementsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**InternalModuleRegistryPresentationAchievementHttpCreateRequest**](InternalModuleRegistryPresentationAchievementHttpCreateRequest.md) | request body | 

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


## RegistryAppsGet

> RegistryAppsGet200Response RegistryAppsGet(ctx).Limit(limit).Offset(offset).Execute()

List apps

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
	resp, r, err := apiClient.RegistryAPI.RegistryAppsGet(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistryAPI.RegistryAppsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegistryAppsGet`: RegistryAppsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `RegistryAPI.RegistryAppsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegistryAppsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | limit | 
 **offset** | **int32** | offset | 

### Return type

[**RegistryAppsGet200Response**](RegistryAppsGet200Response.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegistryItemsGet

> RegistryItemsGet200Response RegistryItemsGet(ctx).Limit(limit).Offset(offset).Execute()

List item-def

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
	resp, r, err := apiClient.RegistryAPI.RegistryItemsGet(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistryAPI.RegistryItemsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegistryItemsGet`: RegistryItemsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `RegistryAPI.RegistryItemsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegistryItemsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | limit | 
 **offset** | **int32** | offset | 

### Return type

[**RegistryItemsGet200Response**](RegistryItemsGet200Response.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegistryItemsItemDefIdDelete

> GithubComNotPlatformInternalServerTemplatesResponseTemplate RegistryItemsItemDefIdDelete(ctx, itemDefId).Execute()

Delete item-def

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
	itemDefId := "itemDefId_example" // string | item_def_id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegistryAPI.RegistryItemsItemDefIdDelete(context.Background(), itemDefId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistryAPI.RegistryItemsItemDefIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegistryItemsItemDefIdDelete`: GithubComNotPlatformInternalServerTemplatesResponseTemplate
	fmt.Fprintf(os.Stdout, "Response from `RegistryAPI.RegistryItemsItemDefIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemDefId** | **string** | item_def_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegistryItemsItemDefIdDeleteRequest struct via the builder pattern


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


## RegistryItemsItemDefIdGet

> RegistryItemsItemDefIdGet200Response RegistryItemsItemDefIdGet(ctx, itemDefId).Execute()

Get item-def

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
	itemDefId := "itemDefId_example" // string | item_def_id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegistryAPI.RegistryItemsItemDefIdGet(context.Background(), itemDefId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistryAPI.RegistryItemsItemDefIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegistryItemsItemDefIdGet`: RegistryItemsItemDefIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `RegistryAPI.RegistryItemsItemDefIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemDefId** | **string** | item_def_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegistryItemsItemDefIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RegistryItemsItemDefIdGet200Response**](RegistryItemsItemDefIdGet200Response.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegistryItemsItemDefIdPut

> GithubComNotPlatformInternalServerTemplatesResponseTemplate RegistryItemsItemDefIdPut(ctx, itemDefId).Request(request).Execute()

Update item-def

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
	itemDefId := "itemDefId_example" // string | item_def_id
	request := *openapiclient.NewInternalModuleRegistryPresentationItemHttpUpdateRequest() // InternalModuleRegistryPresentationItemHttpUpdateRequest | request body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegistryAPI.RegistryItemsItemDefIdPut(context.Background(), itemDefId).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistryAPI.RegistryItemsItemDefIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegistryItemsItemDefIdPut`: GithubComNotPlatformInternalServerTemplatesResponseTemplate
	fmt.Fprintf(os.Stdout, "Response from `RegistryAPI.RegistryItemsItemDefIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemDefId** | **string** | item_def_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegistryItemsItemDefIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**InternalModuleRegistryPresentationItemHttpUpdateRequest**](InternalModuleRegistryPresentationItemHttpUpdateRequest.md) | request body | 

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


## RegistryItemsPost

> GithubComNotPlatformInternalServerTemplatesResponseTemplate RegistryItemsPost(ctx).Request(request).Execute()

Create item-def

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
	request := *openapiclient.NewInternalModuleRegistryPresentationItemHttpCreateRequest() // InternalModuleRegistryPresentationItemHttpCreateRequest | request body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegistryAPI.RegistryItemsPost(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistryAPI.RegistryItemsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegistryItemsPost`: GithubComNotPlatformInternalServerTemplatesResponseTemplate
	fmt.Fprintf(os.Stdout, "Response from `RegistryAPI.RegistryItemsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegistryItemsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**InternalModuleRegistryPresentationItemHttpCreateRequest**](InternalModuleRegistryPresentationItemHttpCreateRequest.md) | request body | 

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

