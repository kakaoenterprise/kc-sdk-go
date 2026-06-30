# \InstanceGroupsParameterGroupsAPI

All URIs are relative to *https://mysql.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetryMysqlParameterGroupSync**](InstanceGroupsParameterGroupsAPI.md#RetryMysqlParameterGroupSync) | **Post** /api/v1/instance-groups/{instance_group_id}/parameter-groups/retry | Retry MySQL parameter group sync



## RetryMysqlParameterGroupSync

> RetryMysqlParameterGroupSync(ctx, instanceGroupId).XAuthToken(xAuthToken).Execute()

Retry MySQL parameter group sync



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/services/mysql"
)

func main() {
	instanceGroupId := "instanceGroupId_example" // string | 대상 MySQL 인스턴스 그룹 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstanceGroupsParameterGroupsAPI.RetryMysqlParameterGroupSync(context.Background(), instanceGroupId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceGroupsParameterGroupsAPI.RetryMysqlParameterGroupSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceGroupId** | **string** | 대상 MySQL 인스턴스 그룹 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetryMysqlParameterGroupSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

 (empty response body)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

