# \AuthAPI

All URIs are relative to *http://localhost:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthGeneratePost**](AuthAPI.md#AuthGeneratePost) | **Post** /auth/generate | Generate new auth



## AuthGeneratePost

> AuthGeneratePost200Response AuthGeneratePost(ctx).Request(request).Execute()

Generate new auth

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
	request := *openapiclient.NewGithubComNotPlatformInternalModuleAuthPresentationAuthorizationHttpModelAuthGenerateRequest() // GithubComNotPlatformInternalModuleAuthPresentationAuthorizationHttpModelAuthGenerateRequest | request body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.AuthGeneratePost(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.AuthGeneratePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthGeneratePost`: AuthGeneratePost200Response
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.AuthGeneratePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthGeneratePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**GithubComNotPlatformInternalModuleAuthPresentationAuthorizationHttpModelAuthGenerateRequest**](GithubComNotPlatformInternalModuleAuthPresentationAuthorizationHttpModelAuthGenerateRequest.md) | request body | 

### Return type

[**AuthGeneratePost200Response**](AuthGeneratePost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

