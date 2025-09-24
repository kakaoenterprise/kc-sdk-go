# \UpgradesAPI

All URIs are relative to *https://kubernetes-engine.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAvailableKubernetesVersions**](UpgradesAPI.md#ListAvailableKubernetesVersions) | **Get** /api/v1/versions | List available Kubernetes versions



## ListAvailableKubernetesVersions

> GetK8sUpgradeVersionsResponseModel ListAvailableKubernetesVersions(ctx).XAuthToken(xAuthToken).Execute()

List available Kubernetes versions



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
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpgradesAPI.ListAvailableKubernetesVersions(context.Background()).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpgradesAPI.ListAvailableKubernetesVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAvailableKubernetesVersions`: GetK8sUpgradeVersionsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `UpgradesAPI.ListAvailableKubernetesVersions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAvailableKubernetesVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**GetK8sUpgradeVersionsResponseModel**](GetK8sUpgradeVersionsResponseModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

