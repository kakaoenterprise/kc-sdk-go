# \CustomParameterGroupsAPI

All URIs are relative to *https://mysql.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMysqlCustomParameterGroup**](CustomParameterGroupsAPI.md#CreateMysqlCustomParameterGroup) | **Post** /api/v1/custom-parameter-groups | Create MySQL custom parameter group
[**DeleteMysqlCustomParameterGroup**](CustomParameterGroupsAPI.md#DeleteMysqlCustomParameterGroup) | **Delete** /api/v1/custom-parameter-groups/{custom_parameter_group_id} | Delete MySQL custom parameter group
[**GetMysqlCustomParameterGroup**](CustomParameterGroupsAPI.md#GetMysqlCustomParameterGroup) | **Get** /api/v1/custom-parameter-groups/{custom_parameter_group_id} | Get MySQL custom parameter group
[**ListMysqlCustomParameterGroupEvents**](CustomParameterGroupsAPI.md#ListMysqlCustomParameterGroupEvents) | **Get** /api/v1/custom-parameter-groups/{custom_parameter_group_id}/events | List MySQL custom parameter group events
[**ListMysqlCustomParameterGroups**](CustomParameterGroupsAPI.md#ListMysqlCustomParameterGroups) | **Get** /api/v1/custom-parameter-groups | List MySQL custom parameter groups
[**ListMysqlInstanceGroupsUsingCustomParameterGroup**](CustomParameterGroupsAPI.md#ListMysqlInstanceGroupsUsingCustomParameterGroup) | **Get** /api/v1/custom-parameter-groups/{custom_parameter_group_id}/instance-groups | List MySQL instance groups using custom parameter group
[**ResetMysqlCustomParameterGroup**](CustomParameterGroupsAPI.md#ResetMysqlCustomParameterGroup) | **Post** /api/v1/custom-parameter-groups/{custom_parameter_group_id}/reset | Reset MySQL custom parameter group
[**RollbackMysqlCustomParameterGroups**](CustomParameterGroupsAPI.md#RollbackMysqlCustomParameterGroups) | **Post** /api/v1/custom-parameter-groups/{custom_parameter_group_id}/rollback | Rollback MySQL custom parameter groups
[**UpdateMysqlCustomParameterGroup**](CustomParameterGroupsAPI.md#UpdateMysqlCustomParameterGroup) | **Patch** /api/v1/custom-parameter-groups/{custom_parameter_group_id} | Update MySQL custom parameter group



## CreateMysqlCustomParameterGroup

> CreateMySQLCustomParameterGroupResponseModel CreateMysqlCustomParameterGroup(ctx).XAuthToken(xAuthToken).BodyCreateMysqlCustomParameterGroup(bodyCreateMysqlCustomParameterGroup).Execute()

Create MySQL custom parameter group



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
	bodyCreateMysqlCustomParameterGroup := *openapiclient.NewBodyCreateMysqlCustomParameterGroup(*openapiclient.NewMysqlV1ApiCreateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel("Name_example", "SourceParameterGroupId_example", openapiclient.ParameterGroupType("DEFAULT"))) // BodyCreateMysqlCustomParameterGroup | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomParameterGroupsAPI.CreateMysqlCustomParameterGroup(context.Background()).XAuthToken(xAuthToken).BodyCreateMysqlCustomParameterGroup(bodyCreateMysqlCustomParameterGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomParameterGroupsAPI.CreateMysqlCustomParameterGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMysqlCustomParameterGroup`: CreateMySQLCustomParameterGroupResponseModel
	fmt.Fprintf(os.Stdout, "Response from `CustomParameterGroupsAPI.CreateMysqlCustomParameterGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMysqlCustomParameterGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyCreateMysqlCustomParameterGroup** | [**BodyCreateMysqlCustomParameterGroup**](BodyCreateMysqlCustomParameterGroup.md) |  | 

### Return type

[**CreateMySQLCustomParameterGroupResponseModel**](CreateMySQLCustomParameterGroupResponseModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMysqlCustomParameterGroup

> DeleteMysqlCustomParameterGroup(ctx, customParameterGroupId).XAuthToken(xAuthToken).Execute()

Delete MySQL custom parameter group



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
	customParameterGroupId := "customParameterGroupId_example" // string | 커스텀 MySQL 파라미터 그룹 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomParameterGroupsAPI.DeleteMysqlCustomParameterGroup(context.Background(), customParameterGroupId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomParameterGroupsAPI.DeleteMysqlCustomParameterGroup``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteMysqlCustomParameterGroupRequest struct via the builder pattern


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


## GetMysqlCustomParameterGroup

> GetMySQLCustomParameterGroupResponseModel GetMysqlCustomParameterGroup(ctx, customParameterGroupId).XAuthToken(xAuthToken).Execute()

Get MySQL custom parameter group



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
	customParameterGroupId := "customParameterGroupId_example" // string | 조회할 MySQL 커스텀 파라미터 그룹 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomParameterGroupsAPI.GetMysqlCustomParameterGroup(context.Background(), customParameterGroupId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomParameterGroupsAPI.GetMysqlCustomParameterGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMysqlCustomParameterGroup`: GetMySQLCustomParameterGroupResponseModel
	fmt.Fprintf(os.Stdout, "Response from `CustomParameterGroupsAPI.GetMysqlCustomParameterGroup`: %v\n", resp)
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

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**GetMySQLCustomParameterGroupResponseModel**](GetMySQLCustomParameterGroupResponseModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMysqlCustomParameterGroupEvents

> GetMySQLCustomParameterGroupEventsResponseModel ListMysqlCustomParameterGroupEvents(ctx, customParameterGroupId).XAuthToken(xAuthToken).Execute()

List MySQL custom parameter group events



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
	customParameterGroupId := "customParameterGroupId_example" // string | 커스텀 MySQL 파라미터 그룹 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomParameterGroupsAPI.ListMysqlCustomParameterGroupEvents(context.Background(), customParameterGroupId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomParameterGroupsAPI.ListMysqlCustomParameterGroupEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMysqlCustomParameterGroupEvents`: GetMySQLCustomParameterGroupEventsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `CustomParameterGroupsAPI.ListMysqlCustomParameterGroupEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customParameterGroupId** | **string** | 커스텀 MySQL 파라미터 그룹 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMysqlCustomParameterGroupEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**GetMySQLCustomParameterGroupEventsResponseModel**](GetMySQLCustomParameterGroupEventsResponseModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMysqlCustomParameterGroups

> GetMySQLCustomParameterGroupsResponseModel ListMysqlCustomParameterGroups(ctx).XAuthToken(xAuthToken).ShowInstanceGroupsInfo(showInstanceGroupsInfo).Execute()

List MySQL custom parameter groups



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
	showInstanceGroupsInfo := true // bool | 각 커스텀 파라미터 그룹에 연결된 인스턴스 그룹 정보를 함께 조회할지 여부 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomParameterGroupsAPI.ListMysqlCustomParameterGroups(context.Background()).XAuthToken(xAuthToken).ShowInstanceGroupsInfo(showInstanceGroupsInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomParameterGroupsAPI.ListMysqlCustomParameterGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMysqlCustomParameterGroups`: GetMySQLCustomParameterGroupsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `CustomParameterGroupsAPI.ListMysqlCustomParameterGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMysqlCustomParameterGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **showInstanceGroupsInfo** | **bool** | 각 커스텀 파라미터 그룹에 연결된 인스턴스 그룹 정보를 함께 조회할지 여부 | 

### Return type

[**GetMySQLCustomParameterGroupsResponseModel**](GetMySQLCustomParameterGroupsResponseModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMysqlInstanceGroupsUsingCustomParameterGroup

> MysqlV1ApiListMysqlInstanceGroupsUsingCustomParameterGroupModelGetMySQLInstanceGroupsUsingDefaultParameterGroupResponseModel ListMysqlInstanceGroupsUsingCustomParameterGroup(ctx, customParameterGroupId).XAuthToken(xAuthToken).Execute()

List MySQL instance groups using custom parameter group



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
	customParameterGroupId := "customParameterGroupId_example" // string | 커스텀 MySQL 파라미터 그룹 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomParameterGroupsAPI.ListMysqlInstanceGroupsUsingCustomParameterGroup(context.Background(), customParameterGroupId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomParameterGroupsAPI.ListMysqlInstanceGroupsUsingCustomParameterGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMysqlInstanceGroupsUsingCustomParameterGroup`: MysqlV1ApiListMysqlInstanceGroupsUsingCustomParameterGroupModelGetMySQLInstanceGroupsUsingDefaultParameterGroupResponseModel
	fmt.Fprintf(os.Stdout, "Response from `CustomParameterGroupsAPI.ListMysqlInstanceGroupsUsingCustomParameterGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customParameterGroupId** | **string** | 커스텀 MySQL 파라미터 그룹 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMysqlInstanceGroupsUsingCustomParameterGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**MysqlV1ApiListMysqlInstanceGroupsUsingCustomParameterGroupModelGetMySQLInstanceGroupsUsingDefaultParameterGroupResponseModel**](MysqlV1ApiListMysqlInstanceGroupsUsingCustomParameterGroupModelGetMySQLInstanceGroupsUsingDefaultParameterGroupResponseModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetMysqlCustomParameterGroup

> ResetMysqlCustomParameterGroup(ctx, customParameterGroupId).XAuthToken(xAuthToken).BodyResetMysqlCustomParameterGroup(bodyResetMysqlCustomParameterGroup).Execute()

Reset MySQL custom parameter group



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
	customParameterGroupId := "customParameterGroupId_example" // string | 커스텀 MySQL 파라미터 그룹 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	bodyResetMysqlCustomParameterGroup := *openapiclient.NewBodyResetMysqlCustomParameterGroup(*openapiclient.NewMysqlV1ApiResetMysqlCustomParameterGroupModelCustomParameterGroupRequestModel()) // BodyResetMysqlCustomParameterGroup | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomParameterGroupsAPI.ResetMysqlCustomParameterGroup(context.Background(), customParameterGroupId).XAuthToken(xAuthToken).BodyResetMysqlCustomParameterGroup(bodyResetMysqlCustomParameterGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomParameterGroupsAPI.ResetMysqlCustomParameterGroup``: %v\n", err)
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

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyResetMysqlCustomParameterGroup** | [**BodyResetMysqlCustomParameterGroup**](BodyResetMysqlCustomParameterGroup.md) |  | 

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

> RollbackMysqlCustomParameterGroups(ctx, customParameterGroupId).XAuthToken(xAuthToken).Execute()

Rollback MySQL custom parameter groups



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
	customParameterGroupId := "customParameterGroupId_example" // string | 커스텀 MySQL 파라미터 그룹 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomParameterGroupsAPI.RollbackMysqlCustomParameterGroups(context.Background(), customParameterGroupId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomParameterGroupsAPI.RollbackMysqlCustomParameterGroups``: %v\n", err)
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


## UpdateMysqlCustomParameterGroup

> UpdateMysqlCustomParameterGroup(ctx, customParameterGroupId).XAuthToken(xAuthToken).BodyUpdateMysqlCustomParameterGroup(bodyUpdateMysqlCustomParameterGroup).Execute()

Update MySQL custom parameter group



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
	customParameterGroupId := "customParameterGroupId_example" // string | 커스텀 MySQL 파라미터 그룹 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	bodyUpdateMysqlCustomParameterGroup := *openapiclient.NewBodyUpdateMysqlCustomParameterGroup(*openapiclient.NewMysqlV1ApiUpdateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel()) // BodyUpdateMysqlCustomParameterGroup | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomParameterGroupsAPI.UpdateMysqlCustomParameterGroup(context.Background(), customParameterGroupId).XAuthToken(xAuthToken).BodyUpdateMysqlCustomParameterGroup(bodyUpdateMysqlCustomParameterGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomParameterGroupsAPI.UpdateMysqlCustomParameterGroup``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateMysqlCustomParameterGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyUpdateMysqlCustomParameterGroup** | [**BodyUpdateMysqlCustomParameterGroup**](BodyUpdateMysqlCustomParameterGroup.md) |  | 

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

