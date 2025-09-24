# \IdentityAPI

All URIs are relative to *https://iam.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IssueToken**](IdentityAPI.md#IssueToken) | **Post** /identity/v3/auth/tokens | 인증 토큰 발급 API
[**ValidateToken**](IdentityAPI.md#ValidateToken) | **Get** /identity/v3/auth/tokens | 토큰 유효성 검증 및 정보 확인



## IssueToken

> TokenResponse IssueToken(ctx).SwaggerIdPwdRequest(swaggerIdPwdRequest).Execute()

인증 토큰 발급 API



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
	swaggerIdPwdRequest := *openapiclient.NewSwaggerIdPwdRequest() // SwaggerIdPwdRequest | ID&PW 인증. Project Scope 기준

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityAPI.IssueToken(context.Background()).SwaggerIdPwdRequest(swaggerIdPwdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityAPI.IssueToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IssueToken`: TokenResponse
	fmt.Fprintf(os.Stdout, "Response from `IdentityAPI.IssueToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIssueTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **swaggerIdPwdRequest** | [**SwaggerIdPwdRequest**](SwaggerIdPwdRequest.md) | ID&amp;PW 인증. Project Scope 기준 | 

### Return type

[**TokenResponse**](TokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateToken

> TokenResponse ValidateToken(ctx).XAuthToken(xAuthToken).XSubjectToken(xSubjectToken).Execute()

토큰 유효성 검증 및 정보 확인



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
	xAuthToken := "xAuthToken_example" // string | authentication token
	xSubjectToken := "xSubjectToken_example" // string | authentication token

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityAPI.ValidateToken(context.Background()).XAuthToken(xAuthToken).XSubjectToken(xSubjectToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityAPI.ValidateToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateToken`: TokenResponse
	fmt.Fprintf(os.Stdout, "Response from `IdentityAPI.ValidateToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | authentication token | 
 **xSubjectToken** | **string** | authentication token | 

### Return type

[**TokenResponse**](TokenResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

