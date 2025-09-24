# \InstanceRunAnActionAPI

All URIs are relative to *https://bcs.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInstanceConsoleLog**](InstanceRunAnActionAPI.md#GetInstanceConsoleLog) | **Get** /api/v1/instances/{instance_id}/console-logs | Get instance console log
[**HardRebootInstance**](InstanceRunAnActionAPI.md#HardRebootInstance) | **Post** /api/v1/instances/{instance_id}/hard-reboot | Hard reboot instance
[**RebuildInstance**](InstanceRunAnActionAPI.md#RebuildInstance) | **Post** /api/v1/instances/{instance_id}/rebuild | Rebuild instance
[**ResizeInstance**](InstanceRunAnActionAPI.md#ResizeInstance) | **Post** /api/v1/instances/{instance_id}/resize | Resize instance
[**ShelveInstance**](InstanceRunAnActionAPI.md#ShelveInstance) | **Post** /api/v1/instances/{instance_id}/shelve | Shelve instance
[**SoftRebootInstance**](InstanceRunAnActionAPI.md#SoftRebootInstance) | **Post** /api/v1/instances/{instance_id}/soft-reboot | Soft reboot instance 
[**StartInstance**](InstanceRunAnActionAPI.md#StartInstance) | **Post** /api/v1/instances/{instance_id}/start | Start instance
[**StopInstance**](InstanceRunAnActionAPI.md#StopInstance) | **Post** /api/v1/instances/{instance_id}/stop | Stop instance
[**UnshelveInstance**](InstanceRunAnActionAPI.md#UnshelveInstance) | **Post** /api/v1/instances/{instance_id}/unshelve | Unshelve instance



## GetInstanceConsoleLog

> InstanceConsoleLogModel GetInstanceConsoleLog(ctx, instanceId).XAuthToken(xAuthToken).Lines(lines).Execute()

Get instance console log



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
	lines := int32(56) // int32 | 콘솔 로그의 마지막 줄부터 조회할 라인(line)의 수 <br/> - 미 지정 시, 에러 반환   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceRunAnActionAPI.GetInstanceConsoleLog(context.Background(), instanceId).XAuthToken(xAuthToken).Lines(lines).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceRunAnActionAPI.GetInstanceConsoleLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstanceConsoleLog`: InstanceConsoleLogModel
	fmt.Fprintf(os.Stdout, "Response from `InstanceRunAnActionAPI.GetInstanceConsoleLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceConsoleLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **lines** | **int32** | 콘솔 로그의 마지막 줄부터 조회할 라인(line)의 수 &lt;br/&gt; - 미 지정 시, 에러 반환   | 

### Return type

[**InstanceConsoleLogModel**](InstanceConsoleLogModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HardRebootInstance

> interface{} HardRebootInstance(ctx, instanceId).XAuthToken(xAuthToken).Execute()

Hard reboot instance



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
	resp, r, err := apiClient.InstanceRunAnActionAPI.HardRebootInstance(context.Background(), instanceId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceRunAnActionAPI.HardRebootInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HardRebootInstance`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `InstanceRunAnActionAPI.HardRebootInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiHardRebootInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

**interface{}**

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RebuildInstance

> ResponseRebuildInstanceModel RebuildInstance(ctx, instanceId).XAuthToken(xAuthToken).BodyRebuildInstance(bodyRebuildInstance).Execute()

Rebuild instance



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
	bodyRebuildInstance := *openapiclient.NewBodyRebuildInstance(*openapiclient.NewRequestRebuildInstanceModel("ImageId_example")) // BodyRebuildInstance | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceRunAnActionAPI.RebuildInstance(context.Background(), instanceId).XAuthToken(xAuthToken).BodyRebuildInstance(bodyRebuildInstance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceRunAnActionAPI.RebuildInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RebuildInstance`: ResponseRebuildInstanceModel
	fmt.Fprintf(os.Stdout, "Response from `InstanceRunAnActionAPI.RebuildInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebuildInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyRebuildInstance** | [**BodyRebuildInstance**](BodyRebuildInstance.md) |  | 

### Return type

[**ResponseRebuildInstanceModel**](ResponseRebuildInstanceModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResizeInstance

> interface{} ResizeInstance(ctx, instanceId).XAuthToken(xAuthToken).BodyResizeInstance(bodyResizeInstance).Execute()

Resize instance



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
	bodyResizeInstance := *openapiclient.NewBodyResizeInstance(*openapiclient.NewResizeInstanceModel("Id_example")) // BodyResizeInstance | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceRunAnActionAPI.ResizeInstance(context.Background(), instanceId).XAuthToken(xAuthToken).BodyResizeInstance(bodyResizeInstance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceRunAnActionAPI.ResizeInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResizeInstance`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `InstanceRunAnActionAPI.ResizeInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiResizeInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyResizeInstance** | [**BodyResizeInstance**](BodyResizeInstance.md) |  | 

### Return type

**interface{}**

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShelveInstance

> interface{} ShelveInstance(ctx, instanceId).XAuthToken(xAuthToken).Execute()

Shelve instance



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
	resp, r, err := apiClient.InstanceRunAnActionAPI.ShelveInstance(context.Background(), instanceId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceRunAnActionAPI.ShelveInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShelveInstance`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `InstanceRunAnActionAPI.ShelveInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShelveInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

**interface{}**

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SoftRebootInstance

> interface{} SoftRebootInstance(ctx, instanceId).XAuthToken(xAuthToken).Execute()

Soft reboot instance 



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
	resp, r, err := apiClient.InstanceRunAnActionAPI.SoftRebootInstance(context.Background(), instanceId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceRunAnActionAPI.SoftRebootInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SoftRebootInstance`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `InstanceRunAnActionAPI.SoftRebootInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSoftRebootInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

**interface{}**

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartInstance

> interface{} StartInstance(ctx, instanceId).XAuthToken(xAuthToken).Execute()

Start instance



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
	resp, r, err := apiClient.InstanceRunAnActionAPI.StartInstance(context.Background(), instanceId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceRunAnActionAPI.StartInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartInstance`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `InstanceRunAnActionAPI.StartInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

**interface{}**

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopInstance

> interface{} StopInstance(ctx, instanceId).XAuthToken(xAuthToken).Execute()

Stop instance



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
	resp, r, err := apiClient.InstanceRunAnActionAPI.StopInstance(context.Background(), instanceId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceRunAnActionAPI.StopInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopInstance`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `InstanceRunAnActionAPI.StopInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

**interface{}**

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnshelveInstance

> interface{} UnshelveInstance(ctx, instanceId).XAuthToken(xAuthToken).Execute()

Unshelve instance



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
	resp, r, err := apiClient.InstanceRunAnActionAPI.UnshelveInstance(context.Background(), instanceId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceRunAnActionAPI.UnshelveInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnshelveInstance`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `InstanceRunAnActionAPI.UnshelveInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnshelveInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

**interface{}**

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

