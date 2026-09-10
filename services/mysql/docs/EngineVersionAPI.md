# \EngineVersionAPI

All URIs are relative to *https://mysql.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAvailableMysqlEngineVersions**](EngineVersionAPI.md#ListAvailableMysqlEngineVersions) | **Get** /api/v1/engine-versions | List available MySQL engine versions



## ListAvailableMysqlEngineVersions

> ListAvailableMysqlEngineVersionsResponse ListAvailableMysqlEngineVersions(ctx).XAuthToken(xAuthToken).Execute()

List available MySQL engine versions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/mysql"
)

func main() {
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngineVersionAPI.ListAvailableMysqlEngineVersions(context.Background()).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngineVersionAPI.ListAvailableMysqlEngineVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAvailableMysqlEngineVersions`: ListAvailableMysqlEngineVersionsResponse
	fmt.Fprintf(os.Stdout, "Response from `EngineVersionAPI.ListAvailableMysqlEngineVersions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAvailableMysqlEngineVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**ListAvailableMysqlEngineVersionsResponse**](ListAvailableMysqlEngineVersionsResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

