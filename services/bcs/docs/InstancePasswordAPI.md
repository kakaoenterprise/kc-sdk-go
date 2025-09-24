# \InstancePasswordAPI

All URIs are relative to *https://bcs.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInstancePassword**](InstancePasswordAPI.md#GetInstancePassword) | **Get** /api/v1/instances/{instance_id}/password | Get instance password



## GetInstancePassword

> InstancePasswordModel GetInstancePassword(ctx, instanceId).XAuthToken(xAuthToken).Execute()

Get instance password



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
	instanceId := "instanceId_example" // string | 인스턴스의 고유 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstancePasswordAPI.GetInstancePassword(context.Background(), instanceId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancePasswordAPI.GetInstancePassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstancePassword`: InstancePasswordModel
	fmt.Fprintf(os.Stdout, "Response from `InstancePasswordAPI.GetInstancePassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstancePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**InstancePasswordModel**](InstancePasswordModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

