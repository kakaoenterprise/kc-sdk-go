# \InstanceAPI

All URIs are relative to *https://mysql.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportMysqlInstanceLogs**](InstanceAPI.md#ExportMysqlInstanceLogs) | **Post** /api/v1/instance-groups/{instance_group_id}/instances/{instance_id}/export-logs | Export MySQL instance logs
[**ListMysqlInstanceTypesFlavors**](InstanceAPI.md#ListMysqlInstanceTypesFlavors) | **Get** /api/v1/flavors | List MySQL instance types (flavors)
[**ListMysqlInstances**](InstanceAPI.md#ListMysqlInstances) | **Get** /api/v1/instance-groups/{instance_group_id}/instances | List MySQL instances
[**RestartMysqlInstances**](InstanceAPI.md#RestartMysqlInstances) | **Post** /api/v1/instance-groups/{instance_group_id}/restart-instances | Restart MySQL instance



## ExportMysqlInstanceLogs

> ExportMysqlInstanceLogs(ctx, instanceGroupId, instanceId).ExportMysqlInstanceLogsRequest(exportMysqlInstanceLogsRequest).Execute()

Export MySQL instance logs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/mysql"
)

func main() {
	instanceGroupId := "instanceGroupId_example" // string | 대상 MySQL 인스턴스 그룹 ID
	instanceId := "instanceId_example" // string | 대상 MySQL 인스턴스 ID
	exportMysqlInstanceLogsRequest := *openapiclient.NewExportMysqlInstanceLogsRequest(*openapiclient.NewExportMysqlInstanceLogs("Bucket_example", []openapiclient.LogInfoRequest{*openapiclient.NewLogInfoRequest(openapiclient.LogType("GENERAL_LOG"))}, "Path_example", "UserCredentialId_example", "UserCredentialSecret_example")) // ExportMysqlInstanceLogsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstanceAPI.ExportMysqlInstanceLogs(context.Background(), instanceGroupId, instanceId).ExportMysqlInstanceLogsRequest(exportMysqlInstanceLogsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.ExportMysqlInstanceLogs``: %v\n", err)
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


 **exportMysqlInstanceLogsRequest** | [**ExportMysqlInstanceLogsRequest**](ExportMysqlInstanceLogsRequest.md) |  | 

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


## ListMysqlInstanceTypesFlavors

> ListMysqlInstanceTypesFlavorsResponse ListMysqlInstanceTypesFlavors(ctx).ShowAll(showAll).Execute()

List MySQL instance types (flavors)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/mysql"
)

func main() {
	showAll := true // bool | 모든 인스턴스 유형(Flavor)을 포함할지 여부 - `false`로 설정 시, 사용 중단된 인스턴스 유형을 모두 제외하여 반환 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.ListMysqlInstanceTypesFlavors(context.Background()).ShowAll(showAll).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.ListMysqlInstanceTypesFlavors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMysqlInstanceTypesFlavors`: ListMysqlInstanceTypesFlavorsResponse
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.ListMysqlInstanceTypesFlavors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMysqlInstanceTypesFlavorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **showAll** | **bool** | 모든 인스턴스 유형(Flavor)을 포함할지 여부 - &#x60;false&#x60;로 설정 시, 사용 중단된 인스턴스 유형을 모두 제외하여 반환 | 

### Return type

[**ListMysqlInstanceTypesFlavorsResponse**](ListMysqlInstanceTypesFlavorsResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMysqlInstances

> ListMysqlInstancesResponse ListMysqlInstances(ctx, instanceGroupId).Execute()

List MySQL instances



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/mysql"
)

func main() {
	instanceGroupId := "instanceGroupId_example" // string | 대상 MySQL 인스턴스 그룹 ID - [List MySQL instance groups](/openapi/data-store/mysql/list-mysql-instance-groups)에서 확인

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.ListMysqlInstances(context.Background(), instanceGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.ListMysqlInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMysqlInstances`: ListMysqlInstancesResponse
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.ListMysqlInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceGroupId** | **string** | 대상 MySQL 인스턴스 그룹 ID - [List MySQL instance groups](/openapi/data-store/mysql/list-mysql-instance-groups)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMysqlInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListMysqlInstancesResponse**](ListMysqlInstancesResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestartMysqlInstances

> RestartMysqlInstances(ctx, instanceGroupId).RestartMysqlInstancesRequest(restartMysqlInstancesRequest).Execute()

Restart MySQL instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/mysql"
)

func main() {
	instanceGroupId := "instanceGroupId_example" // string | 대상 MySQL 인스턴스 그룹 ID
	restartMysqlInstancesRequest := *openapiclient.NewRestartMysqlInstancesRequest(*openapiclient.NewRestartMysqlInstances([]string{"InstanceIds_example"})) // RestartMysqlInstancesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstanceAPI.RestartMysqlInstances(context.Background(), instanceGroupId).RestartMysqlInstancesRequest(restartMysqlInstancesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.RestartMysqlInstances``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRestartMysqlInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **restartMysqlInstancesRequest** | [**RestartMysqlInstancesRequest**](RestartMysqlInstancesRequest.md) |  | 

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

