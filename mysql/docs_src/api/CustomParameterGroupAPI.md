# \CustomParameterGroupAPI

All URIs are relative to *https://mysql.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMysqlCustomParameterGroup**](CustomParameterGroupAPI.md#CreateMysqlCustomParameterGroup) | **Post** /api/v1/custom-parameter-groups | Create MySQL custom parameter group
[**DeleteMysqlCustomParameterGroup**](CustomParameterGroupAPI.md#DeleteMysqlCustomParameterGroup) | **Delete** /api/v1/custom-parameter-groups/{custom_parameter_group_id} | Delete MySQL custom parameter group
[**GetMysqlCustomParameterGroup**](CustomParameterGroupAPI.md#GetMysqlCustomParameterGroup) | **Get** /api/v1/custom-parameter-groups/{custom_parameter_group_id} | Get MySQL custom parameter group
[**ListMysqlCustomParameterGroupEvents**](CustomParameterGroupAPI.md#ListMysqlCustomParameterGroupEvents) | **Get** /api/v1/custom-parameter-groups/{custom_parameter_group_id}/events | List MySQL custom parameter group events
[**ListMysqlCustomParameterGroups**](CustomParameterGroupAPI.md#ListMysqlCustomParameterGroups) | **Get** /api/v1/custom-parameter-groups | List MySQL custom parameter groups
[**ResetMysqlCustomParameterGroup**](CustomParameterGroupAPI.md#ResetMysqlCustomParameterGroup) | **Post** /api/v1/custom-parameter-groups/{custom_parameter_group_id}/reset | Reset MySQL custom parameter group
[**RollbackMysqlCustomParameterGroups**](CustomParameterGroupAPI.md#RollbackMysqlCustomParameterGroups) | **Post** /api/v1/custom-parameter-groups/{custom_parameter_group_id}/rollback | Rollback MySQL custom parameter groups
[**UpdateMysqlCustomParameterGroup**](CustomParameterGroupAPI.md#UpdateMysqlCustomParameterGroup) | **Patch** /api/v1/custom-parameter-groups/{custom_parameter_group_id} | Update MySQL custom parameter group



## CreateMysqlCustomParameterGroup

> CreateMysqlCustomParameterGroupResponse CreateMysqlCustomParameterGroup(ctx).CreateMysqlCustomParameterGroupRequest(createMysqlCustomParameterGroupRequest).Execute()

Create MySQL custom parameter group



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
	createMysqlCustomParameterGroupRequest := *openapiclient.NewCreateMysqlCustomParameterGroupRequest(*openapiclient.NewCreateMysqlCustomParameterGroup("Name_example", "SourceParameterGroupId_example", openapiclient.ParameterGroupType("DEFAULT"))) // CreateMysqlCustomParameterGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomParameterGroupAPI.CreateMysqlCustomParameterGroup(context.Background()).CreateMysqlCustomParameterGroupRequest(createMysqlCustomParameterGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomParameterGroupAPI.CreateMysqlCustomParameterGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMysqlCustomParameterGroup`: CreateMysqlCustomParameterGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomParameterGroupAPI.CreateMysqlCustomParameterGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMysqlCustomParameterGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createMysqlCustomParameterGroupRequest** | [**CreateMysqlCustomParameterGroupRequest**](CreateMysqlCustomParameterGroupRequest.md) |  | 

### Return type

[**CreateMysqlCustomParameterGroupResponse**](CreateMysqlCustomParameterGroupResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMysqlCustomParameterGroup

> DeleteMysqlCustomParameterGroup(ctx, customParameterGroupId).Execute()

Delete MySQL custom parameter group



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
	customParameterGroupId := "customParameterGroupId_example" // string | 커스텀 MySQL 파라미터 그룹 ID - [List MySQL custom parameter groups](/openapi/data-store/mysql/list-mysql-custom-parameter-groups)에서 확인

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomParameterGroupAPI.DeleteMysqlCustomParameterGroup(context.Background(), customParameterGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomParameterGroupAPI.DeleteMysqlCustomParameterGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customParameterGroupId** | **string** | 커스텀 MySQL 파라미터 그룹 ID - [List MySQL custom parameter groups](/openapi/data-store/mysql/list-mysql-custom-parameter-groups)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMysqlCustomParameterGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetMysqlCustomParameterGroup

> GetMysqlCustomParameterGroupResponse GetMysqlCustomParameterGroup(ctx, customParameterGroupId).Execute()

Get MySQL custom parameter group



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
	customParameterGroupId := "customParameterGroupId_example" // string | 조회할 MySQL 커스텀 파라미터 그룹 ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomParameterGroupAPI.GetMysqlCustomParameterGroup(context.Background(), customParameterGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomParameterGroupAPI.GetMysqlCustomParameterGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMysqlCustomParameterGroup`: GetMysqlCustomParameterGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomParameterGroupAPI.GetMysqlCustomParameterGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customParameterGroupId** | **string** | 조회할 MySQL 커스텀 파라미터 그룹 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMysqlCustomParameterGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetMysqlCustomParameterGroupResponse**](GetMysqlCustomParameterGroupResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMysqlCustomParameterGroupEvents

> ListMysqlCustomParameterGroupEventsResponse ListMysqlCustomParameterGroupEvents(ctx, customParameterGroupId).Execute()

List MySQL custom parameter group events



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
	customParameterGroupId := "customParameterGroupId_example" // string | 커스텀 MySQL 파라미터 그룹 ID - [List MySQL custom parameter groups](/openapi/data-store/mysql/list-mysql-custom-parameter-groups)에서 확인

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomParameterGroupAPI.ListMysqlCustomParameterGroupEvents(context.Background(), customParameterGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomParameterGroupAPI.ListMysqlCustomParameterGroupEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMysqlCustomParameterGroupEvents`: ListMysqlCustomParameterGroupEventsResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomParameterGroupAPI.ListMysqlCustomParameterGroupEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customParameterGroupId** | **string** | 커스텀 MySQL 파라미터 그룹 ID - [List MySQL custom parameter groups](/openapi/data-store/mysql/list-mysql-custom-parameter-groups)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMysqlCustomParameterGroupEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListMysqlCustomParameterGroupEventsResponse**](ListMysqlCustomParameterGroupEventsResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMysqlCustomParameterGroups

> ListMysqlCustomParameterGroupsResponse ListMysqlCustomParameterGroups(ctx).ShowInstanceGroupsInfo(showInstanceGroupsInfo).Execute()

List MySQL custom parameter groups



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
	showInstanceGroupsInfo := true // bool | 각 커스텀 파라미터 그룹에 연결된 인스턴스 그룹 정보를 함께 조회할지 여부 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomParameterGroupAPI.ListMysqlCustomParameterGroups(context.Background()).ShowInstanceGroupsInfo(showInstanceGroupsInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomParameterGroupAPI.ListMysqlCustomParameterGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMysqlCustomParameterGroups`: ListMysqlCustomParameterGroupsResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomParameterGroupAPI.ListMysqlCustomParameterGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMysqlCustomParameterGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **showInstanceGroupsInfo** | **bool** | 각 커스텀 파라미터 그룹에 연결된 인스턴스 그룹 정보를 함께 조회할지 여부 | 

### Return type

[**ListMysqlCustomParameterGroupsResponse**](ListMysqlCustomParameterGroupsResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetMysqlCustomParameterGroup

> ResetMysqlCustomParameterGroup(ctx, customParameterGroupId).ResetMysqlCustomParameterGroupRequest(resetMysqlCustomParameterGroupRequest).Execute()

Reset MySQL custom parameter group



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
	customParameterGroupId := "customParameterGroupId_example" // string | 커스텀 MySQL 파라미터 그룹 ID
	resetMysqlCustomParameterGroupRequest := *openapiclient.NewResetMysqlCustomParameterGroupRequest(*openapiclient.NewResetMysqlCustomParameterGroup()) // ResetMysqlCustomParameterGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomParameterGroupAPI.ResetMysqlCustomParameterGroup(context.Background(), customParameterGroupId).ResetMysqlCustomParameterGroupRequest(resetMysqlCustomParameterGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomParameterGroupAPI.ResetMysqlCustomParameterGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customParameterGroupId** | **string** | 커스텀 MySQL 파라미터 그룹 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetMysqlCustomParameterGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resetMysqlCustomParameterGroupRequest** | [**ResetMysqlCustomParameterGroupRequest**](ResetMysqlCustomParameterGroupRequest.md) |  | 

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


## RollbackMysqlCustomParameterGroups

> RollbackMysqlCustomParameterGroups(ctx, customParameterGroupId).Execute()

Rollback MySQL custom parameter groups



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
	customParameterGroupId := "customParameterGroupId_example" // string | 커스텀 MySQL 파라미터 그룹 ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomParameterGroupAPI.RollbackMysqlCustomParameterGroups(context.Background(), customParameterGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomParameterGroupAPI.RollbackMysqlCustomParameterGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customParameterGroupId** | **string** | 커스텀 MySQL 파라미터 그룹 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRollbackMysqlCustomParameterGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## UpdateMysqlCustomParameterGroup

> UpdateMysqlCustomParameterGroup(ctx, customParameterGroupId).UpdateMysqlCustomParameterGroupRequest(updateMysqlCustomParameterGroupRequest).Execute()

Update MySQL custom parameter group



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
	customParameterGroupId := "customParameterGroupId_example" // string | 커스텀 MySQL 파라미터 그룹 ID - [List MySQL custom parameter groups](/openapi/data-store/mysql/list-mysql-custom-parameter-groups)에서 확인
	updateMysqlCustomParameterGroupRequest := *openapiclient.NewUpdateMysqlCustomParameterGroupRequest(*openapiclient.NewUpdateMysqlCustomParameterGroup()) // UpdateMysqlCustomParameterGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomParameterGroupAPI.UpdateMysqlCustomParameterGroup(context.Background(), customParameterGroupId).UpdateMysqlCustomParameterGroupRequest(updateMysqlCustomParameterGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomParameterGroupAPI.UpdateMysqlCustomParameterGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customParameterGroupId** | **string** | 커스텀 MySQL 파라미터 그룹 ID - [List MySQL custom parameter groups](/openapi/data-store/mysql/list-mysql-custom-parameter-groups)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMysqlCustomParameterGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateMysqlCustomParameterGroupRequest** | [**UpdateMysqlCustomParameterGroupRequest**](UpdateMysqlCustomParameterGroupRequest.md) |  | 

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

