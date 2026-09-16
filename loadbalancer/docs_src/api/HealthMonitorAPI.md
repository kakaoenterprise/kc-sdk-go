# \HealthMonitorAPI

All URIs are relative to *https://load-balancer.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHealthMonitor**](HealthMonitorAPI.md#CreateHealthMonitor) | **Post** /api/v1/load-balancers/health-monitors | Create health monitor
[**DeleteHealthMonitor**](HealthMonitorAPI.md#DeleteHealthMonitor) | **Delete** /api/v1/load-balancers/health-monitors/{health_monitor_id} | Delete health monitor
[**GetTargetGroupHealthCheckSubnets**](HealthMonitorAPI.md#GetTargetGroupHealthCheckSubnets) | **Get** /api/v1/load-balancers/target-groups/{target_group_id}/health-check-subnets | Get target group health check subnets
[**GetTargetGroupHealthMonitor**](HealthMonitorAPI.md#GetTargetGroupHealthMonitor) | **Get** /api/v1/load-balancers/health-monitors/{health_monitor_id} | Get target group health monitor
[**UpdateHealthMonitor**](HealthMonitorAPI.md#UpdateHealthMonitor) | **Put** /api/v1/load-balancers/health-monitors/{health_monitor_id} | Update health monitor



## CreateHealthMonitor

> CreateHealthMonitorResponse CreateHealthMonitor(ctx).CreateHealthMonitorRequest(createHealthMonitorRequest).Execute()

Create health monitor



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/loadbalancer"
)

func main() {
	createHealthMonitorRequest := *openapiclient.NewCreateHealthMonitorRequest(*openapiclient.NewCreateHealthMonitor(int32(123), int32(123), int32(123), "TargetGroupId_example", int32(123), openapiclient.HealthMonitorType("HTTP"))) // CreateHealthMonitorRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HealthMonitorAPI.CreateHealthMonitor(context.Background()).CreateHealthMonitorRequest(createHealthMonitorRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthMonitorAPI.CreateHealthMonitor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateHealthMonitor`: CreateHealthMonitorResponse
	fmt.Fprintf(os.Stdout, "Response from `HealthMonitorAPI.CreateHealthMonitor`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateHealthMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createHealthMonitorRequest** | [**CreateHealthMonitorRequest**](CreateHealthMonitorRequest.md) |  | 

### Return type

[**CreateHealthMonitorResponse**](CreateHealthMonitorResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteHealthMonitor

> DeleteHealthMonitor(ctx, healthMonitorId).Execute()

Delete health monitor



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/loadbalancer"
)

func main() {
	healthMonitorId := "healthMonitorId_example" // string | 삭제할 헬스 모니터의 ID - [List target groups](/openapi/networking/lb/list-target-groups)에서 확인

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HealthMonitorAPI.DeleteHealthMonitor(context.Background(), healthMonitorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthMonitorAPI.DeleteHealthMonitor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**healthMonitorId** | **string** | 삭제할 헬스 모니터의 ID - [List target groups](/openapi/networking/lb/list-target-groups)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHealthMonitorRequest struct via the builder pattern


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


## GetTargetGroupHealthCheckSubnets

> GetTargetGroupHealthCheckSubnetsResponse GetTargetGroupHealthCheckSubnets(ctx, targetGroupId).Execute()

Get target group health check subnets



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/loadbalancer"
)

func main() {
	targetGroupId := "targetGroupId_example" // string | 대상 그룹 ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HealthMonitorAPI.GetTargetGroupHealthCheckSubnets(context.Background(), targetGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthMonitorAPI.GetTargetGroupHealthCheckSubnets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTargetGroupHealthCheckSubnets`: GetTargetGroupHealthCheckSubnetsResponse
	fmt.Fprintf(os.Stdout, "Response from `HealthMonitorAPI.GetTargetGroupHealthCheckSubnets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetGroupId** | **string** | 대상 그룹 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTargetGroupHealthCheckSubnetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetTargetGroupHealthCheckSubnetsResponse**](GetTargetGroupHealthCheckSubnetsResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTargetGroupHealthMonitor

> GetTargetGroupHealthMonitorResponse GetTargetGroupHealthMonitor(ctx, healthMonitorId).Execute()

Get target group health monitor



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/loadbalancer"
)

func main() {
	healthMonitorId := "healthMonitorId_example" // string | 조회할 헬스 모니터 ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HealthMonitorAPI.GetTargetGroupHealthMonitor(context.Background(), healthMonitorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthMonitorAPI.GetTargetGroupHealthMonitor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTargetGroupHealthMonitor`: GetTargetGroupHealthMonitorResponse
	fmt.Fprintf(os.Stdout, "Response from `HealthMonitorAPI.GetTargetGroupHealthMonitor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**healthMonitorId** | **string** | 조회할 헬스 모니터 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTargetGroupHealthMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetTargetGroupHealthMonitorResponse**](GetTargetGroupHealthMonitorResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHealthMonitor

> UpdateHealthMonitorResponse UpdateHealthMonitor(ctx, healthMonitorId).UpdateHealthMonitorRequest(updateHealthMonitorRequest).Execute()

Update health monitor



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/loadbalancer"
)

func main() {
	healthMonitorId := "healthMonitorId_example" // string | 수정 대상 헬스 모니터의 ID - [List target groups](/openapi/networking/lb/list-target-groups)에서 확인
	updateHealthMonitorRequest := *openapiclient.NewUpdateHealthMonitorRequest(*openapiclient.NewUpdateHealthMonitor()) // UpdateHealthMonitorRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HealthMonitorAPI.UpdateHealthMonitor(context.Background(), healthMonitorId).UpdateHealthMonitorRequest(updateHealthMonitorRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthMonitorAPI.UpdateHealthMonitor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateHealthMonitor`: UpdateHealthMonitorResponse
	fmt.Fprintf(os.Stdout, "Response from `HealthMonitorAPI.UpdateHealthMonitor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**healthMonitorId** | **string** | 수정 대상 헬스 모니터의 ID - [List target groups](/openapi/networking/lb/list-target-groups)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHealthMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateHealthMonitorRequest** | [**UpdateHealthMonitorRequest**](UpdateHealthMonitorRequest.md) |  | 

### Return type

[**UpdateHealthMonitorResponse**](UpdateHealthMonitorResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

