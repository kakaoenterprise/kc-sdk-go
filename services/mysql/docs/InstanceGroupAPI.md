# \InstanceGroupAPI

All URIs are relative to *https://mysql.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyMysqlParameterGroup**](InstanceGroupAPI.md#ApplyMysqlParameterGroup) | **Post** /api/v1/instance-groups/{instance_group_id}/parameter-groups | Apply MySQL parameter group
[**CreateMysqlInstanceGroup**](InstanceGroupAPI.md#CreateMysqlInstanceGroup) | **Post** /api/v1/instance-groups | Create MySQL instance group
[**DeleteMysqlInstanceGroup**](InstanceGroupAPI.md#DeleteMysqlInstanceGroup) | **Delete** /api/v1/instance-groups/{instance_group_id} | Delete MySQL instance group
[**ExtendMysqlInstanceGroupVolume**](InstanceGroupAPI.md#ExtendMysqlInstanceGroupVolume) | **Post** /api/v1/instance-groups/{instance_group_id}/extend-volume | Extend MySQL instance group volume
[**GetMysqlInstanceGroup**](InstanceGroupAPI.md#GetMysqlInstanceGroup) | **Get** /api/v1/instance-groups/{instance_group_id} | Get MySQL instance group
[**GetMysqlRestorableTime**](InstanceGroupAPI.md#GetMysqlRestorableTime) | **Get** /api/v1/instance-groups/{instance_group_id}/restorable-time | Get MySQL restorable time
[**ListMysqlInstanceGroups**](InstanceGroupAPI.md#ListMysqlInstanceGroups) | **Get** /api/v1/instance-groups | List MySQL instance groups
[**ListMysqlInstanceGroupsUsingCustomParameterGroup**](InstanceGroupAPI.md#ListMysqlInstanceGroupsUsingCustomParameterGroup) | **Get** /api/v1/custom-parameter-groups/{custom_parameter_group_id}/instance-groups | List MySQL instance groups using custom parameter group
[**ListMysqlInstanceGroupsUsingDefaultParameterGroup**](InstanceGroupAPI.md#ListMysqlInstanceGroupsUsingDefaultParameterGroup) | **Get** /api/v1/default-parameter-groups/{default_parameter_group_id}/instance-groups | List MySQL instance groups using default parameter group
[**RetryMysqlParameterGroupSync**](InstanceGroupAPI.md#RetryMysqlParameterGroupSync) | **Post** /api/v1/instance-groups/{instance_group_id}/parameter-groups/retry | Retry MySQL parameter group sync
[**ScaleInMysqlInstanceGroup**](InstanceGroupAPI.md#ScaleInMysqlInstanceGroup) | **Post** /api/v1/instance-groups/{instance_group_id}/scale-in | Scale in MySQL instance group
[**ScaleOutMysqlInstanceGroup**](InstanceGroupAPI.md#ScaleOutMysqlInstanceGroup) | **Post** /api/v1/instance-groups/{instance_group_id}/scale-out | Scale out MySQL instance group
[**SwitchoverMysqlInstanceGroup**](InstanceGroupAPI.md#SwitchoverMysqlInstanceGroup) | **Post** /api/v1/instance-groups/{instance_group_id}/switch-over | Switchover MySQL instance group
[**UpdateMysqlInstanceGroupBackupSchedule**](InstanceGroupAPI.md#UpdateMysqlInstanceGroupBackupSchedule) | **Post** /api/v1/instance-groups/{instance_group_id}/backup-schedules/{backup_schedule_id} | Update MySQL backup schedule 
[**UpdateMysqlSecurityGroups**](InstanceGroupAPI.md#UpdateMysqlSecurityGroups) | **Patch** /api/v1/instance-groups/{instance_group_id}/security-groups | Update MySQL security groups



## ApplyMysqlParameterGroup

> ApplyMysqlParameterGroup(ctx, instanceGroupId).XAuthToken(xAuthToken).ApplyMysqlParameterGroupRequest(applyMysqlParameterGroupRequest).Execute()

Apply MySQL parameter group



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
	instanceGroupId := "instanceGroupId_example" // string | 대상 MySQL 인스턴스 그룹 ID <br/>- [List MySQL instance groups](https://docs.kakaocloud.com/openapi/data-store/mysql/list-mysql-instance-groups)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	applyMysqlParameterGroupRequest := *openapiclient.NewApplyMysqlParameterGroupRequest(*openapiclient.NewApplyMysqlParameterGroup("Id_example", openapiclient.ParameterGroupType("DEFAULT"))) // ApplyMysqlParameterGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstanceGroupAPI.ApplyMysqlParameterGroup(context.Background(), instanceGroupId).XAuthToken(xAuthToken).ApplyMysqlParameterGroupRequest(applyMysqlParameterGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceGroupAPI.ApplyMysqlParameterGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceGroupId** | **string** | 대상 MySQL 인스턴스 그룹 ID &lt;br/&gt;- [List MySQL instance groups](https://docs.kakaocloud.com/openapi/data-store/mysql/list-mysql-instance-groups)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplyMysqlParameterGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **applyMysqlParameterGroupRequest** | [**ApplyMysqlParameterGroupRequest**](ApplyMysqlParameterGroupRequest.md) |  | 

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


## CreateMysqlInstanceGroup

> CreateMysqlInstanceGroupResponse CreateMysqlInstanceGroup(ctx).XAuthToken(xAuthToken).CreateMysqlInstanceGroupRequest(createMysqlInstanceGroupRequest).Execute()

Create MySQL instance group



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
	createMysqlInstanceGroupRequest := *openapiclient.NewCreateMysqlInstanceGroupRequest(*openapiclient.NewCreateMysqlInstanceGroup("Name_example", *openapiclient.NewNetworkInfoRequest([]string{"SecurityGroupIds_example"}, *openapiclient.NewSubnetInfoRequest(int32(123), "SubnetId_example")), *openapiclient.NewSpecContentRequest("DatabaseUserName_example", "DatabaseUserPassword_example", int32(123), "EngineVersion_example", "FlavorId_example", int32(123), int32(123)), *openapiclient.NewBackupScheduleRequest(false), *openapiclient.NewParameterGroupRequest(openapiclient.ParameterGroupType("DEFAULT"), "Id_example"))) // CreateMysqlInstanceGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceGroupAPI.CreateMysqlInstanceGroup(context.Background()).XAuthToken(xAuthToken).CreateMysqlInstanceGroupRequest(createMysqlInstanceGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceGroupAPI.CreateMysqlInstanceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMysqlInstanceGroup`: CreateMysqlInstanceGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `InstanceGroupAPI.CreateMysqlInstanceGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMysqlInstanceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **createMysqlInstanceGroupRequest** | [**CreateMysqlInstanceGroupRequest**](CreateMysqlInstanceGroupRequest.md) |  | 

### Return type

[**CreateMysqlInstanceGroupResponse**](CreateMysqlInstanceGroupResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMysqlInstanceGroup

> DeleteMysqlInstanceGroup(ctx, instanceGroupId).KeepAutoBackup(keepAutoBackup).XAuthToken(xAuthToken).Execute()

Delete MySQL instance group



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
	instanceGroupId := "instanceGroupId_example" // string | 대상 MySQL 인스턴스 그룹 ID <br/>- [List MySQL instance groups](https://docs.kakaocloud.com/openapi/data-store/mysql/list-mysql-instance-groups)에서 확인
	keepAutoBackup := true // bool | 자동 백업을 유지할지 여부
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstanceGroupAPI.DeleteMysqlInstanceGroup(context.Background(), instanceGroupId).KeepAutoBackup(keepAutoBackup).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceGroupAPI.DeleteMysqlInstanceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceGroupId** | **string** | 대상 MySQL 인스턴스 그룹 ID &lt;br/&gt;- [List MySQL instance groups](https://docs.kakaocloud.com/openapi/data-store/mysql/list-mysql-instance-groups)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMysqlInstanceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **keepAutoBackup** | **bool** | 자동 백업을 유지할지 여부 | 
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


## ExtendMysqlInstanceGroupVolume

> ExtendMysqlInstanceGroupVolume(ctx, instanceGroupId).XAuthToken(xAuthToken).ExtendMysqlInstanceGroupVolumeRequest(extendMysqlInstanceGroupVolumeRequest).Execute()

Extend MySQL instance group volume



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
	instanceGroupId := "instanceGroupId_example" // string | 대상 MySQL 인스턴스 그룹 ID <br/>- [List MySQL instance groups](https://docs.kakaocloud.com/openapi/data-store/mysql/list-mysql-instance-groups)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	extendMysqlInstanceGroupVolumeRequest := *openapiclient.NewExtendMysqlInstanceGroupVolumeRequest(*openapiclient.NewExtendMysqlInstanceGroupVolume(int32(123), int32(123))) // ExtendMysqlInstanceGroupVolumeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstanceGroupAPI.ExtendMysqlInstanceGroupVolume(context.Background(), instanceGroupId).XAuthToken(xAuthToken).ExtendMysqlInstanceGroupVolumeRequest(extendMysqlInstanceGroupVolumeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceGroupAPI.ExtendMysqlInstanceGroupVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceGroupId** | **string** | 대상 MySQL 인스턴스 그룹 ID &lt;br/&gt;- [List MySQL instance groups](https://docs.kakaocloud.com/openapi/data-store/mysql/list-mysql-instance-groups)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiExtendMysqlInstanceGroupVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **extendMysqlInstanceGroupVolumeRequest** | [**ExtendMysqlInstanceGroupVolumeRequest**](ExtendMysqlInstanceGroupVolumeRequest.md) |  | 

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


## GetMysqlInstanceGroup

> GetMysqlInstanceGroupResponse GetMysqlInstanceGroup(ctx, instanceGroupId).XAuthToken(xAuthToken).Execute()

Get MySQL instance group



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
	instanceGroupId := "instanceGroupId_example" // string | 대상 MySQL 인스턴스 그룹 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceGroupAPI.GetMysqlInstanceGroup(context.Background(), instanceGroupId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceGroupAPI.GetMysqlInstanceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMysqlInstanceGroup`: GetMysqlInstanceGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `InstanceGroupAPI.GetMysqlInstanceGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceGroupId** | **string** | 대상 MySQL 인스턴스 그룹 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMysqlInstanceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**GetMysqlInstanceGroupResponse**](GetMysqlInstanceGroupResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMysqlRestorableTime

> GetMysqlRestorableTimeResponse GetMysqlRestorableTime(ctx, instanceGroupId).XAuthToken(xAuthToken).Execute()

Get MySQL restorable time



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
	instanceGroupId := "instanceGroupId_example" // string | 대상 MySQL 인스턴스 그룹 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceGroupAPI.GetMysqlRestorableTime(context.Background(), instanceGroupId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceGroupAPI.GetMysqlRestorableTime``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMysqlRestorableTime`: GetMysqlRestorableTimeResponse
	fmt.Fprintf(os.Stdout, "Response from `InstanceGroupAPI.GetMysqlRestorableTime`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceGroupId** | **string** | 대상 MySQL 인스턴스 그룹 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMysqlRestorableTimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**GetMysqlRestorableTimeResponse**](GetMysqlRestorableTimeResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMysqlInstanceGroups

> ListMysqlInstanceGroupsResponse ListMysqlInstanceGroups(ctx).XAuthToken(xAuthToken).Execute()

List MySQL instance groups



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
	resp, r, err := apiClient.InstanceGroupAPI.ListMysqlInstanceGroups(context.Background()).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceGroupAPI.ListMysqlInstanceGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMysqlInstanceGroups`: ListMysqlInstanceGroupsResponse
	fmt.Fprintf(os.Stdout, "Response from `InstanceGroupAPI.ListMysqlInstanceGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMysqlInstanceGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**ListMysqlInstanceGroupsResponse**](ListMysqlInstanceGroupsResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMysqlInstanceGroupsUsingCustomParameterGroup

> ListMysqlInstanceGroupsUsingCustomParameterGroupResponse ListMysqlInstanceGroupsUsingCustomParameterGroup(ctx, customParameterGroupId).XAuthToken(xAuthToken).Execute()

List MySQL instance groups using custom parameter group



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
	customParameterGroupId := "customParameterGroupId_example" // string | 커스텀 MySQL 파라미터 그룹 ID <br/>- [List MySQL custom parameter groups](https://docs.kakaocloud.com/openapi/data-store/mysql/list-mysql-custom-parameter-groups)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceGroupAPI.ListMysqlInstanceGroupsUsingCustomParameterGroup(context.Background(), customParameterGroupId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceGroupAPI.ListMysqlInstanceGroupsUsingCustomParameterGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMysqlInstanceGroupsUsingCustomParameterGroup`: ListMysqlInstanceGroupsUsingCustomParameterGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `InstanceGroupAPI.ListMysqlInstanceGroupsUsingCustomParameterGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customParameterGroupId** | **string** | 커스텀 MySQL 파라미터 그룹 ID &lt;br/&gt;- [List MySQL custom parameter groups](https://docs.kakaocloud.com/openapi/data-store/mysql/list-mysql-custom-parameter-groups)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMysqlInstanceGroupsUsingCustomParameterGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**ListMysqlInstanceGroupsUsingCustomParameterGroupResponse**](ListMysqlInstanceGroupsUsingCustomParameterGroupResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMysqlInstanceGroupsUsingDefaultParameterGroup

> ListMysqlInstanceGroupsUsingDefaultParameterGroupResponse ListMysqlInstanceGroupsUsingDefaultParameterGroup(ctx, defaultParameterGroupId).XAuthToken(xAuthToken).Execute()

List MySQL instance groups using default parameter group



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
	defaultParameterGroupId := "defaultParameterGroupId_example" // string | 기본 MySQL 파라미터 그룹 ID <br/>- [List MySQL default parameter groups](https://docs.kakaocloud.com/openapi/data-store/mysql/list-mysql-default-parameter-groups)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceGroupAPI.ListMysqlInstanceGroupsUsingDefaultParameterGroup(context.Background(), defaultParameterGroupId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceGroupAPI.ListMysqlInstanceGroupsUsingDefaultParameterGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMysqlInstanceGroupsUsingDefaultParameterGroup`: ListMysqlInstanceGroupsUsingDefaultParameterGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `InstanceGroupAPI.ListMysqlInstanceGroupsUsingDefaultParameterGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**defaultParameterGroupId** | **string** | 기본 MySQL 파라미터 그룹 ID &lt;br/&gt;- [List MySQL default parameter groups](https://docs.kakaocloud.com/openapi/data-store/mysql/list-mysql-default-parameter-groups)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMysqlInstanceGroupsUsingDefaultParameterGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**ListMysqlInstanceGroupsUsingDefaultParameterGroupResponse**](ListMysqlInstanceGroupsUsingDefaultParameterGroupResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/mysql"
)

func main() {
	instanceGroupId := "instanceGroupId_example" // string | 대상 MySQL 인스턴스 그룹 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstanceGroupAPI.RetryMysqlParameterGroupSync(context.Background(), instanceGroupId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceGroupAPI.RetryMysqlParameterGroupSync``: %v\n", err)
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


## ScaleInMysqlInstanceGroup

> ScaleInMysqlInstanceGroup(ctx, instanceGroupId).XAuthToken(xAuthToken).ScaleInMysqlInstanceGroupRequest(scaleInMysqlInstanceGroupRequest).Execute()

Scale in MySQL instance group



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
	instanceGroupId := "instanceGroupId_example" // string | 대상 MySQL 인스턴스 그룹 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	scaleInMysqlInstanceGroupRequest := *openapiclient.NewScaleInMysqlInstanceGroupRequest(*openapiclient.NewScaleInMysqlInstanceGroup([]string{"InstanceIds_example"})) // ScaleInMysqlInstanceGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstanceGroupAPI.ScaleInMysqlInstanceGroup(context.Background(), instanceGroupId).XAuthToken(xAuthToken).ScaleInMysqlInstanceGroupRequest(scaleInMysqlInstanceGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceGroupAPI.ScaleInMysqlInstanceGroup``: %v\n", err)
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

Other parameters are passed through a pointer to a apiScaleInMysqlInstanceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **scaleInMysqlInstanceGroupRequest** | [**ScaleInMysqlInstanceGroupRequest**](ScaleInMysqlInstanceGroupRequest.md) |  | 

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


## ScaleOutMysqlInstanceGroup

> ScaleOutMysqlInstanceGroup(ctx, instanceGroupId).XAuthToken(xAuthToken).ScaleOutMysqlInstanceGroupRequest(scaleOutMysqlInstanceGroupRequest).Execute()

Scale out MySQL instance group



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
	instanceGroupId := "instanceGroupId_example" // string | 대상 MySQL 인스턴스 그룹 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	scaleOutMysqlInstanceGroupRequest := *openapiclient.NewScaleOutMysqlInstanceGroupRequest(*openapiclient.NewScaleOutMysqlInstanceGroup([]openapiclient.SubnetInfoRequest{*openapiclient.NewSubnetInfoRequest(int32(123), "SubnetId_example")})) // ScaleOutMysqlInstanceGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstanceGroupAPI.ScaleOutMysqlInstanceGroup(context.Background(), instanceGroupId).XAuthToken(xAuthToken).ScaleOutMysqlInstanceGroupRequest(scaleOutMysqlInstanceGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceGroupAPI.ScaleOutMysqlInstanceGroup``: %v\n", err)
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

Other parameters are passed through a pointer to a apiScaleOutMysqlInstanceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **scaleOutMysqlInstanceGroupRequest** | [**ScaleOutMysqlInstanceGroupRequest**](ScaleOutMysqlInstanceGroupRequest.md) |  | 

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


## SwitchoverMysqlInstanceGroup

> SwitchoverMysqlInstanceGroup(ctx, instanceGroupId).XAuthToken(xAuthToken).Execute()

Switchover MySQL instance group



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
	instanceGroupId := "instanceGroupId_example" // string | 대상 MySQL 인스턴스 그룹 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstanceGroupAPI.SwitchoverMysqlInstanceGroup(context.Background(), instanceGroupId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceGroupAPI.SwitchoverMysqlInstanceGroup``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSwitchoverMysqlInstanceGroupRequest struct via the builder pattern


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


## UpdateMysqlInstanceGroupBackupSchedule

> UpdateMysqlInstanceGroupBackupSchedule(ctx, instanceGroupId, backupScheduleId).XAuthToken(xAuthToken).UpdateMysqlInstanceGroupBackupScheduleRequest(updateMysqlInstanceGroupBackupScheduleRequest).Execute()

Update MySQL backup schedule 



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
	instanceGroupId := "instanceGroupId_example" // string | 대상 MySQL 인스턴스 그룹 ID <br/>- [List MySQL instance groups](https://docs.kakaocloud.com/openapi/data-store/mysql/list-mysql-instance-groups)에서 확인
	backupScheduleId := "backupScheduleId_example" // string | 백업 스케줄 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	updateMysqlInstanceGroupBackupScheduleRequest := *openapiclient.NewUpdateMysqlInstanceGroupBackupScheduleRequest(*openapiclient.NewUpdateMysqlInstanceGroupBackupSchedule(false)) // UpdateMysqlInstanceGroupBackupScheduleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstanceGroupAPI.UpdateMysqlInstanceGroupBackupSchedule(context.Background(), instanceGroupId, backupScheduleId).XAuthToken(xAuthToken).UpdateMysqlInstanceGroupBackupScheduleRequest(updateMysqlInstanceGroupBackupScheduleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceGroupAPI.UpdateMysqlInstanceGroupBackupSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceGroupId** | **string** | 대상 MySQL 인스턴스 그룹 ID &lt;br/&gt;- [List MySQL instance groups](https://docs.kakaocloud.com/openapi/data-store/mysql/list-mysql-instance-groups)에서 확인 | 
**backupScheduleId** | **string** | 백업 스케줄 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMysqlInstanceGroupBackupScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **updateMysqlInstanceGroupBackupScheduleRequest** | [**UpdateMysqlInstanceGroupBackupScheduleRequest**](UpdateMysqlInstanceGroupBackupScheduleRequest.md) |  | 

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


## UpdateMysqlSecurityGroups

> UpdateMysqlSecurityGroups(ctx, instanceGroupId).XAuthToken(xAuthToken).UpdateMysqlSecurityGroupsRequest(updateMysqlSecurityGroupsRequest).Execute()

Update MySQL security groups



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
	instanceGroupId := "instanceGroupId_example" // string | 대상 MySQL 인스턴스 그룹 ID <br/>- [List MySQL instance groups](https://docs.kakaocloud.com/openapi/data-store/mysql/list-mysql-instance-groups)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	updateMysqlSecurityGroupsRequest := *openapiclient.NewUpdateMysqlSecurityGroupsRequest(*openapiclient.NewUpdateMysqlSecurityGroups([]string{"SecurityGroupIds_example"})) // UpdateMysqlSecurityGroupsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstanceGroupAPI.UpdateMysqlSecurityGroups(context.Background(), instanceGroupId).XAuthToken(xAuthToken).UpdateMysqlSecurityGroupsRequest(updateMysqlSecurityGroupsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceGroupAPI.UpdateMysqlSecurityGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceGroupId** | **string** | 대상 MySQL 인스턴스 그룹 ID &lt;br/&gt;- [List MySQL instance groups](https://docs.kakaocloud.com/openapi/data-store/mysql/list-mysql-instance-groups)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMysqlSecurityGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **updateMysqlSecurityGroupsRequest** | [**UpdateMysqlSecurityGroupsRequest**](UpdateMysqlSecurityGroupsRequest.md) |  | 

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

