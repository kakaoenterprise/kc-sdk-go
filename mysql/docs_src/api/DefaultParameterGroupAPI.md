# \DefaultParameterGroupAPI

All URIs are relative to *https://mysql.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMysqlDefaultParameterGroup**](DefaultParameterGroupAPI.md#GetMysqlDefaultParameterGroup) | **Get** /api/v1/default-parameter-groups/{default_parameter_group_id} | Get MySQL default parameter group
[**ListMysqlDefaultParameterGroupEvents**](DefaultParameterGroupAPI.md#ListMysqlDefaultParameterGroupEvents) | **Get** /api/v1/default-parameter-groups/{default_parameter_group_id}/events | List MySQL default parameter group events
[**ListMysqlDefaultParameterGroups**](DefaultParameterGroupAPI.md#ListMysqlDefaultParameterGroups) | **Get** /api/v1/default-parameter-groups | List MySQL default parameter groups



## GetMysqlDefaultParameterGroup

> GetMysqlDefaultParameterGroupResponse GetMysqlDefaultParameterGroup(ctx, defaultParameterGroupId).Execute()

Get MySQL default parameter group



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
	defaultParameterGroupId := "defaultParameterGroupId_example" // string | 기본 MySQL 파라미터 그룹 ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultParameterGroupAPI.GetMysqlDefaultParameterGroup(context.Background(), defaultParameterGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultParameterGroupAPI.GetMysqlDefaultParameterGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMysqlDefaultParameterGroup`: GetMysqlDefaultParameterGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultParameterGroupAPI.GetMysqlDefaultParameterGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**defaultParameterGroupId** | **string** | 기본 MySQL 파라미터 그룹 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMysqlDefaultParameterGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetMysqlDefaultParameterGroupResponse**](GetMysqlDefaultParameterGroupResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMysqlDefaultParameterGroupEvents

> ListMysqlDefaultParameterGroupEventsResponse ListMysqlDefaultParameterGroupEvents(ctx, defaultParameterGroupId).Execute()

List MySQL default parameter group events



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
	defaultParameterGroupId := "defaultParameterGroupId_example" // string | 기본 MySQL 파라미터 그룹 ID - [List MySQL default parameter groups](/openapi/data-store/mysql/list-mysql-default-parameter-groups)에서 확인

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultParameterGroupAPI.ListMysqlDefaultParameterGroupEvents(context.Background(), defaultParameterGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultParameterGroupAPI.ListMysqlDefaultParameterGroupEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMysqlDefaultParameterGroupEvents`: ListMysqlDefaultParameterGroupEventsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultParameterGroupAPI.ListMysqlDefaultParameterGroupEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**defaultParameterGroupId** | **string** | 기본 MySQL 파라미터 그룹 ID - [List MySQL default parameter groups](/openapi/data-store/mysql/list-mysql-default-parameter-groups)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMysqlDefaultParameterGroupEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListMysqlDefaultParameterGroupEventsResponse**](ListMysqlDefaultParameterGroupEventsResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMysqlDefaultParameterGroups

> ListMysqlDefaultParameterGroupsResponse ListMysqlDefaultParameterGroups(ctx).ShowInstanceGroupsInfo(showInstanceGroupsInfo).Execute()

List MySQL default parameter groups



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
	showInstanceGroupsInfo := true // bool | 각 기본 파라미터 그룹에 연결된 인스턴스 그룹 정보를 함께 조회할지 여부 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultParameterGroupAPI.ListMysqlDefaultParameterGroups(context.Background()).ShowInstanceGroupsInfo(showInstanceGroupsInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultParameterGroupAPI.ListMysqlDefaultParameterGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMysqlDefaultParameterGroups`: ListMysqlDefaultParameterGroupsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultParameterGroupAPI.ListMysqlDefaultParameterGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMysqlDefaultParameterGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **showInstanceGroupsInfo** | **bool** | 각 기본 파라미터 그룹에 연결된 인스턴스 그룹 정보를 함께 조회할지 여부 | 

### Return type

[**ListMysqlDefaultParameterGroupsResponse**](ListMysqlDefaultParameterGroupsResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

