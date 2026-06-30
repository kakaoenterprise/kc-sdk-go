# \DefaultParameterGroupsAPI

All URIs are relative to *https://mysql.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMysqlDefaultParameterGroup**](DefaultParameterGroupsAPI.md#GetMysqlDefaultParameterGroup) | **Get** /api/v1/default-parameter-groups/{default_parameter_group_id} | Get MySQL default parameter group
[**ListMysqlDefaultParameterGroupEvents**](DefaultParameterGroupsAPI.md#ListMysqlDefaultParameterGroupEvents) | **Get** /api/v1/default-parameter-groups/{default_parameter_group_id}/events | List MySQL default parameter group events
[**ListMysqlDefaultParameterGroups**](DefaultParameterGroupsAPI.md#ListMysqlDefaultParameterGroups) | **Get** /api/v1/default-parameter-groups | List MySQL default parameter groups
[**ListMysqlInstanceGroupsUsingDefaultParameterGroup**](DefaultParameterGroupsAPI.md#ListMysqlInstanceGroupsUsingDefaultParameterGroup) | **Get** /api/v1/default-parameter-groups/{default_parameter_group_id}/instance-groups | List MySQL instance groups using default parameter group



## GetMysqlDefaultParameterGroup

> GetMySQLDefaultParameterGroupResponseModel GetMysqlDefaultParameterGroup(ctx, defaultParameterGroupId).XAuthToken(xAuthToken).Execute()

Get MySQL default parameter group



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
	defaultParameterGroupId := "defaultParameterGroupId_example" // string | 기본 MySQL 파라미터 그룹 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultParameterGroupsAPI.GetMysqlDefaultParameterGroup(context.Background(), defaultParameterGroupId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultParameterGroupsAPI.GetMysqlDefaultParameterGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMysqlDefaultParameterGroup`: GetMySQLDefaultParameterGroupResponseModel
	fmt.Fprintf(os.Stdout, "Response from `DefaultParameterGroupsAPI.GetMysqlDefaultParameterGroup`: %v\n", resp)
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

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**GetMySQLDefaultParameterGroupResponseModel**](GetMySQLDefaultParameterGroupResponseModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMysqlDefaultParameterGroupEvents

> GetMySQLDefaultParameterGroupEventsResponseModel ListMysqlDefaultParameterGroupEvents(ctx, defaultParameterGroupId).XAuthToken(xAuthToken).Execute()

List MySQL default parameter group events



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
	defaultParameterGroupId := "defaultParameterGroupId_example" // string | 기본 MySQL 파라미터 그룹 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultParameterGroupsAPI.ListMysqlDefaultParameterGroupEvents(context.Background(), defaultParameterGroupId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultParameterGroupsAPI.ListMysqlDefaultParameterGroupEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMysqlDefaultParameterGroupEvents`: GetMySQLDefaultParameterGroupEventsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `DefaultParameterGroupsAPI.ListMysqlDefaultParameterGroupEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**defaultParameterGroupId** | **string** | 기본 MySQL 파라미터 그룹 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMysqlDefaultParameterGroupEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**GetMySQLDefaultParameterGroupEventsResponseModel**](GetMySQLDefaultParameterGroupEventsResponseModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMysqlDefaultParameterGroups

> GetMySQLDefaultParameterGroupsResponseModel ListMysqlDefaultParameterGroups(ctx).XAuthToken(xAuthToken).ShowInstanceGroupsInfo(showInstanceGroupsInfo).Execute()

List MySQL default parameter groups



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
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	showInstanceGroupsInfo := true // bool | 각 기본 파라미터 그룹에 연결된 인스턴스 그룹 정보를 함께 조회할지 여부 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultParameterGroupsAPI.ListMysqlDefaultParameterGroups(context.Background()).XAuthToken(xAuthToken).ShowInstanceGroupsInfo(showInstanceGroupsInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultParameterGroupsAPI.ListMysqlDefaultParameterGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMysqlDefaultParameterGroups`: GetMySQLDefaultParameterGroupsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `DefaultParameterGroupsAPI.ListMysqlDefaultParameterGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMysqlDefaultParameterGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **showInstanceGroupsInfo** | **bool** | 각 기본 파라미터 그룹에 연결된 인스턴스 그룹 정보를 함께 조회할지 여부 | 

### Return type

[**GetMySQLDefaultParameterGroupsResponseModel**](GetMySQLDefaultParameterGroupsResponseModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMysqlInstanceGroupsUsingDefaultParameterGroup

> MysqlV1ApiListMysqlInstanceGroupsUsingDefaultParameterGroupModelGetMySQLInstanceGroupsUsingDefaultParameterGroupResponseModel ListMysqlInstanceGroupsUsingDefaultParameterGroup(ctx, defaultParameterGroupId).XAuthToken(xAuthToken).Execute()

List MySQL instance groups using default parameter group



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
	defaultParameterGroupId := "defaultParameterGroupId_example" // string | 기본 MySQL 파라미터 그룹 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultParameterGroupsAPI.ListMysqlInstanceGroupsUsingDefaultParameterGroup(context.Background(), defaultParameterGroupId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultParameterGroupsAPI.ListMysqlInstanceGroupsUsingDefaultParameterGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMysqlInstanceGroupsUsingDefaultParameterGroup`: MysqlV1ApiListMysqlInstanceGroupsUsingDefaultParameterGroupModelGetMySQLInstanceGroupsUsingDefaultParameterGroupResponseModel
	fmt.Fprintf(os.Stdout, "Response from `DefaultParameterGroupsAPI.ListMysqlInstanceGroupsUsingDefaultParameterGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**defaultParameterGroupId** | **string** | 기본 MySQL 파라미터 그룹 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMysqlInstanceGroupsUsingDefaultParameterGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**MysqlV1ApiListMysqlInstanceGroupsUsingDefaultParameterGroupModelGetMySQLInstanceGroupsUsingDefaultParameterGroupResponseModel**](MysqlV1ApiListMysqlInstanceGroupsUsingDefaultParameterGroupModelGetMySQLInstanceGroupsUsingDefaultParameterGroupResponseModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

