# \InstanceGroupsAPI

All URIs are relative to *https://mysql.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyMysqlParameterGroup**](InstanceGroupsAPI.md#ApplyMysqlParameterGroup) | **Post** /api/v1/instance-groups/{instance_group_id}/parameter-groups | Apply MySQL parameter group
[**CreateMysqlInstanceGroup**](InstanceGroupsAPI.md#CreateMysqlInstanceGroup) | **Post** /api/v1/instance-groups | Create MySQL instance group
[**DeleteMysqlInstanceGroup**](InstanceGroupsAPI.md#DeleteMysqlInstanceGroup) | **Delete** /api/v1/instance-groups/{instance_group_id} | Delete MySQL instance group
[**ExtendMysqlInstanceGroupVolume**](InstanceGroupsAPI.md#ExtendMysqlInstanceGroupVolume) | **Post** /api/v1/instance-groups/{instance_group_id}/extend-volume | Extend MySQL instance group volume
[**GetMysqlInstanceGroup**](InstanceGroupsAPI.md#GetMysqlInstanceGroup) | **Get** /api/v1/instance-groups/{instance_group_id} | Get MySQL instance group
[**GetMysqlRestorableTime**](InstanceGroupsAPI.md#GetMysqlRestorableTime) | **Get** /api/v1/instance-groups/{instance_group_id}/restorable-time | Get MySQL restorable time
[**ListMysqlInstanceGroups**](InstanceGroupsAPI.md#ListMysqlInstanceGroups) | **Get** /api/v1/instance-groups | List MySQL instance groups
[**ListMysqlInstances**](InstanceGroupsAPI.md#ListMysqlInstances) | **Get** /api/v1/instance-groups/{instance_group_id}/instances | List MySQL instances
[**RestartMysqlInstances**](InstanceGroupsAPI.md#RestartMysqlInstances) | **Post** /api/v1/instance-groups/{instance_group_id}/restart-instances | Restart MySQL instance
[**ScaleInMysqlInstanceGroup**](InstanceGroupsAPI.md#ScaleInMysqlInstanceGroup) | **Post** /api/v1/instance-groups/{instance_group_id}/scale-in | Scale in MySQL instance group
[**ScaleOutMysqlInstanceGroup**](InstanceGroupsAPI.md#ScaleOutMysqlInstanceGroup) | **Post** /api/v1/instance-groups/{instance_group_id}/scale-out | Scale out MySQL instance group
[**SwitchoverMysqlInstanceGroup**](InstanceGroupsAPI.md#SwitchoverMysqlInstanceGroup) | **Post** /api/v1/instance-groups/{instance_group_id}/switch-over | Switchover MySQL instance group
[**UpdateMysqlSecurityGroups**](InstanceGroupsAPI.md#UpdateMysqlSecurityGroups) | **Patch** /api/v1/instance-groups/{instance_group_id}/security-groups | Update MySQL security groups



## ApplyMysqlParameterGroup

> ApplyMysqlParameterGroup(ctx, instanceGroupId).XAuthToken(xAuthToken).BodyApplyMysqlParameterGroup(bodyApplyMysqlParameterGroup).Execute()

Apply MySQL parameter group



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
	bodyApplyMysqlParameterGroup := *openapiclient.NewBodyApplyMysqlParameterGroup(*openapiclient.NewMysqlV1ApiApplyMysqlParameterGroupModelParameterGroupRequestModel("Id_example", openapiclient.ParameterGroupType("DEFAULT"))) // BodyApplyMysqlParameterGroup | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstanceGroupsAPI.ApplyMysqlParameterGroup(context.Background(), instanceGroupId).XAuthToken(xAuthToken).BodyApplyMysqlParameterGroup(bodyApplyMysqlParameterGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceGroupsAPI.ApplyMysqlParameterGroup``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApplyMysqlParameterGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyApplyMysqlParameterGroup** | [**BodyApplyMysqlParameterGroup**](BodyApplyMysqlParameterGroup.md) |  | 

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

> CreateMySQLInstanceGroupResponseModel CreateMysqlInstanceGroup(ctx).XAuthToken(xAuthToken).BodyCreateMysqlInstanceGroup(bodyCreateMysqlInstanceGroup).Execute()

Create MySQL instance group



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
	bodyCreateMysqlInstanceGroup := *openapiclient.NewBodyCreateMysqlInstanceGroup(*openapiclient.NewMysqlV1ApiCreateMysqlInstanceGroupModelInstanceGroupRequestModel("Name_example", *openapiclient.NewNetworkInfoRequestModel([]string{"SecurityGroupIds_example"}, *openapiclient.NewMysqlV1ApiCreateMysqlInstanceGroupModelSubnetInfoRequestModel(int32(123), "SubnetId_example")), *openapiclient.NewSpecContentRequestModel("DatabaseUserName_example", "DatabaseUserPassword_example", int32(123), "EngineVersion_example", "FlavorId_example", int32(123), int32(123)), *openapiclient.NewMysqlV1ApiCreateMysqlInstanceGroupModelBackupScheduleRequestModel(false), *openapiclient.NewMysqlV1ApiCreateMysqlInstanceGroupModelParameterGroupRequestModel(openapiclient.ParameterGroupType("DEFAULT"), "Id_example"))) // BodyCreateMysqlInstanceGroup | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceGroupsAPI.CreateMysqlInstanceGroup(context.Background()).XAuthToken(xAuthToken).BodyCreateMysqlInstanceGroup(bodyCreateMysqlInstanceGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceGroupsAPI.CreateMysqlInstanceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMysqlInstanceGroup`: CreateMySQLInstanceGroupResponseModel
	fmt.Fprintf(os.Stdout, "Response from `InstanceGroupsAPI.CreateMysqlInstanceGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMysqlInstanceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyCreateMysqlInstanceGroup** | [**BodyCreateMysqlInstanceGroup**](BodyCreateMysqlInstanceGroup.md) |  | 

### Return type

[**CreateMySQLInstanceGroupResponseModel**](CreateMySQLInstanceGroupResponseModel.md)

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
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/services/mysql"
)

func main() {
	instanceGroupId := "instanceGroupId_example" // string | 대상 MySQL 인스턴스 그룹 ID
	keepAutoBackup := true // bool | 자동 백업을 유지할지 여부
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstanceGroupsAPI.DeleteMysqlInstanceGroup(context.Background(), instanceGroupId).KeepAutoBackup(keepAutoBackup).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceGroupsAPI.DeleteMysqlInstanceGroup``: %v\n", err)
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

> ExtendMysqlInstanceGroupVolume(ctx, instanceGroupId).XAuthToken(xAuthToken).BodyExtendMysqlInstanceGroupVolume(bodyExtendMysqlInstanceGroupVolume).Execute()

Extend MySQL instance group volume



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
	bodyExtendMysqlInstanceGroupVolume := *openapiclient.NewBodyExtendMysqlInstanceGroupVolume(*openapiclient.NewMysqlV1ApiExtendMysqlInstanceGroupVolumeModelInstanceGroupRequestModel(int32(123), int32(123))) // BodyExtendMysqlInstanceGroupVolume | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstanceGroupsAPI.ExtendMysqlInstanceGroupVolume(context.Background(), instanceGroupId).XAuthToken(xAuthToken).BodyExtendMysqlInstanceGroupVolume(bodyExtendMysqlInstanceGroupVolume).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceGroupsAPI.ExtendMysqlInstanceGroupVolume``: %v\n", err)
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

Other parameters are passed through a pointer to a apiExtendMysqlInstanceGroupVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyExtendMysqlInstanceGroupVolume** | [**BodyExtendMysqlInstanceGroupVolume**](BodyExtendMysqlInstanceGroupVolume.md) |  | 

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

> GetMySQLInstanceGroupResponseModel GetMysqlInstanceGroup(ctx, instanceGroupId).XAuthToken(xAuthToken).Execute()

Get MySQL instance group



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
	resp, r, err := apiClient.InstanceGroupsAPI.GetMysqlInstanceGroup(context.Background(), instanceGroupId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceGroupsAPI.GetMysqlInstanceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMysqlInstanceGroup`: GetMySQLInstanceGroupResponseModel
	fmt.Fprintf(os.Stdout, "Response from `InstanceGroupsAPI.GetMysqlInstanceGroup`: %v\n", resp)
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

[**GetMySQLInstanceGroupResponseModel**](GetMySQLInstanceGroupResponseModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMysqlRestorableTime

> GetMySQLInstanceGroupRestorableTimeResponseModel GetMysqlRestorableTime(ctx, instanceGroupId).XAuthToken(xAuthToken).Execute()

Get MySQL restorable time



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
	resp, r, err := apiClient.InstanceGroupsAPI.GetMysqlRestorableTime(context.Background(), instanceGroupId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceGroupsAPI.GetMysqlRestorableTime``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMysqlRestorableTime`: GetMySQLInstanceGroupRestorableTimeResponseModel
	fmt.Fprintf(os.Stdout, "Response from `InstanceGroupsAPI.GetMysqlRestorableTime`: %v\n", resp)
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

[**GetMySQLInstanceGroupRestorableTimeResponseModel**](GetMySQLInstanceGroupRestorableTimeResponseModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMysqlInstanceGroups

> GetMySQLInstanceGroupsResponseModel ListMysqlInstanceGroups(ctx).XAuthToken(xAuthToken).Execute()

List MySQL instance groups



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceGroupsAPI.ListMysqlInstanceGroups(context.Background()).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceGroupsAPI.ListMysqlInstanceGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMysqlInstanceGroups`: GetMySQLInstanceGroupsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `InstanceGroupsAPI.ListMysqlInstanceGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMysqlInstanceGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**GetMySQLInstanceGroupsResponseModel**](GetMySQLInstanceGroupsResponseModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMysqlInstances

> GetMySQLInstanceGroupInstancesResponseModel ListMysqlInstances(ctx, instanceGroupId).XAuthToken(xAuthToken).Execute()

List MySQL instances



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
	resp, r, err := apiClient.InstanceGroupsAPI.ListMysqlInstances(context.Background(), instanceGroupId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceGroupsAPI.ListMysqlInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMysqlInstances`: GetMySQLInstanceGroupInstancesResponseModel
	fmt.Fprintf(os.Stdout, "Response from `InstanceGroupsAPI.ListMysqlInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceGroupId** | **string** | 대상 MySQL 인스턴스 그룹 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMysqlInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**GetMySQLInstanceGroupInstancesResponseModel**](GetMySQLInstanceGroupInstancesResponseModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestartMysqlInstances

> RestartMysqlInstances(ctx, instanceGroupId).XAuthToken(xAuthToken).BodyRestartMysqlInstances(bodyRestartMysqlInstances).Execute()

Restart MySQL instance



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
	bodyRestartMysqlInstances := *openapiclient.NewBodyRestartMysqlInstances(*openapiclient.NewMysqlV1ApiRestartMysqlInstancesModelInstanceGroupRequestModel([]string{"InstanceIds_example"})) // BodyRestartMysqlInstances | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstanceGroupsAPI.RestartMysqlInstances(context.Background(), instanceGroupId).XAuthToken(xAuthToken).BodyRestartMysqlInstances(bodyRestartMysqlInstances).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceGroupsAPI.RestartMysqlInstances``: %v\n", err)
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

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyRestartMysqlInstances** | [**BodyRestartMysqlInstances**](BodyRestartMysqlInstances.md) |  | 

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


## ScaleInMysqlInstanceGroup

> ScaleInMysqlInstanceGroup(ctx, instanceGroupId).XAuthToken(xAuthToken).BodyScaleInMysqlInstanceGroup(bodyScaleInMysqlInstanceGroup).Execute()

Scale in MySQL instance group



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
	bodyScaleInMysqlInstanceGroup := *openapiclient.NewBodyScaleInMysqlInstanceGroup(*openapiclient.NewMysqlV1ApiScaleInMysqlInstanceGroupModelInstanceGroupRequestModel([]string{"InstanceIds_example"})) // BodyScaleInMysqlInstanceGroup | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstanceGroupsAPI.ScaleInMysqlInstanceGroup(context.Background(), instanceGroupId).XAuthToken(xAuthToken).BodyScaleInMysqlInstanceGroup(bodyScaleInMysqlInstanceGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceGroupsAPI.ScaleInMysqlInstanceGroup``: %v\n", err)
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
 **bodyScaleInMysqlInstanceGroup** | [**BodyScaleInMysqlInstanceGroup**](BodyScaleInMysqlInstanceGroup.md) |  | 

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

> ScaleOutMysqlInstanceGroup(ctx, instanceGroupId).XAuthToken(xAuthToken).BodyScaleOutMysqlInstanceGroup(bodyScaleOutMysqlInstanceGroup).Execute()

Scale out MySQL instance group



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
	bodyScaleOutMysqlInstanceGroup := *openapiclient.NewBodyScaleOutMysqlInstanceGroup(*openapiclient.NewMysqlV1ApiScaleOutMysqlInstanceGroupModelInstanceGroupRequestModel([]openapiclient.MysqlV1ApiScaleOutMysqlInstanceGroupModelSubnetInfoRequestModel{*openapiclient.NewMysqlV1ApiScaleOutMysqlInstanceGroupModelSubnetInfoRequestModel(int32(123), "SubnetId_example")})) // BodyScaleOutMysqlInstanceGroup | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstanceGroupsAPI.ScaleOutMysqlInstanceGroup(context.Background(), instanceGroupId).XAuthToken(xAuthToken).BodyScaleOutMysqlInstanceGroup(bodyScaleOutMysqlInstanceGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceGroupsAPI.ScaleOutMysqlInstanceGroup``: %v\n", err)
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
 **bodyScaleOutMysqlInstanceGroup** | [**BodyScaleOutMysqlInstanceGroup**](BodyScaleOutMysqlInstanceGroup.md) |  | 

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
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/services/mysql"
)

func main() {
	instanceGroupId := "instanceGroupId_example" // string | 대상 MySQL 인스턴스 그룹 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstanceGroupsAPI.SwitchoverMysqlInstanceGroup(context.Background(), instanceGroupId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceGroupsAPI.SwitchoverMysqlInstanceGroup``: %v\n", err)
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


## UpdateMysqlSecurityGroups

> UpdateMysqlSecurityGroups(ctx, instanceGroupId).XAuthToken(xAuthToken).BodyUpdateMysqlSecurityGroups(bodyUpdateMysqlSecurityGroups).Execute()

Update MySQL security groups



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
	bodyUpdateMysqlSecurityGroups := *openapiclient.NewBodyUpdateMysqlSecurityGroups(*openapiclient.NewMysqlV1ApiUpdateMysqlSecurityGroupsModelInstanceGroupRequestModel([]string{"SecurityGroupIds_example"})) // BodyUpdateMysqlSecurityGroups | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstanceGroupsAPI.UpdateMysqlSecurityGroups(context.Background(), instanceGroupId).XAuthToken(xAuthToken).BodyUpdateMysqlSecurityGroups(bodyUpdateMysqlSecurityGroups).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceGroupsAPI.UpdateMysqlSecurityGroups``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateMysqlSecurityGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyUpdateMysqlSecurityGroups** | [**BodyUpdateMysqlSecurityGroups**](BodyUpdateMysqlSecurityGroups.md) |  | 

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

