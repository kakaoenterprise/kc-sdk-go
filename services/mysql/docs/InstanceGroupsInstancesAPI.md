# \InstanceGroupsInstancesAPI

All URIs are relative to *https://mysql.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportMysqlInstanceLogs**](InstanceGroupsInstancesAPI.md#ExportMysqlInstanceLogs) | **Post** /api/v1/instance-groups/{instance_group_id}/instances/{instance_id}/export-logs | Export MySQL instance logs



## ExportMysqlInstanceLogs

> ExportMysqlInstanceLogs(ctx, instanceGroupId, instanceId).XAuthToken(xAuthToken).BodyExportMysqlInstanceLogs(bodyExportMysqlInstanceLogs).Execute()

Export MySQL instance logs



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
	instanceId := "instanceId_example" // string | 대상 MySQL 인스턴스 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	bodyExportMysqlInstanceLogs := *openapiclient.NewBodyExportMysqlInstanceLogs(*openapiclient.NewInstanceRequestModel("Bucket_example", []openapiclient.LogInfoRequestModel{*openapiclient.NewLogInfoRequestModel(openapiclient.LogType("GENERAL_LOG"))}, "Path_example", "UserCredentialId_example", "UserCredentialSecret_example")) // BodyExportMysqlInstanceLogs | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstanceGroupsInstancesAPI.ExportMysqlInstanceLogs(context.Background(), instanceGroupId, instanceId).XAuthToken(xAuthToken).BodyExportMysqlInstanceLogs(bodyExportMysqlInstanceLogs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceGroupsInstancesAPI.ExportMysqlInstanceLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceGroupId** | **string** | 대상 MySQL 인스턴스 그룹 ID | 
**instanceId** | **string** | 대상 MySQL 인스턴스 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportMysqlInstanceLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyExportMysqlInstanceLogs** | [**BodyExportMysqlInstanceLogs**](BodyExportMysqlInstanceLogs.md) |  | 

### Return type

 (empty response body)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

