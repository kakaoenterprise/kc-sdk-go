# \ScalingAPI

All URIs are relative to *https://kubernetes-engine.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNodePoolScheduledScaling**](ScalingAPI.md#CreateNodePoolScheduledScaling) | **Post** /api/v1/clusters/{cluster_name}/node-pools/{node_pool_name}/scheduled-scaling | Create node pool scheduled scaling 
[**DeleteNodePoolScheduledScaling**](ScalingAPI.md#DeleteNodePoolScheduledScaling) | **Delete** /api/v1/clusters/{cluster_name}/node-pools/{node_pool_name}/scheduled-scaling/{schedule_name} | Delete node pool scheduled scaling  
[**ListNodePoolScheduledScalings**](ScalingAPI.md#ListNodePoolScheduledScalings) | **Get** /api/v1/clusters/{cluster_name}/node-pools/{node_pool_name}/scheduled-scaling | List node pool scheduled scalings  
[**SetNodePoolResourceBasedAutoScaling**](ScalingAPI.md#SetNodePoolResourceBasedAutoScaling) | **Put** /api/v1/clusters/{cluster_name}/node-pools/{node_pool_name}/autoscaling | Set node pool resource-based auto scaling   



## CreateNodePoolScheduledScaling

> interface{} CreateNodePoolScheduledScaling(ctx, clusterName, nodePoolName).XAuthToken(xAuthToken).CreateK8sClusterNodePoolScalingScheduleRequestModel(createK8sClusterNodePoolScalingScheduleRequestModel).Execute()

Create node pool scheduled scaling 



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
	clusterName := "clusterName_example" // string | 대상 클러스터 이름 
	nodePoolName := "nodePoolName_example" // string | 대상 노드 풀 이름
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	createK8sClusterNodePoolScalingScheduleRequestModel := *openapiclient.NewCreateK8sClusterNodePoolScalingScheduleRequestModel(*openapiclient.NewScheduleRequestModel("Name_example", openapiclient.SchedulingType("cron"), int32(123), "StartTime_example")) // CreateK8sClusterNodePoolScalingScheduleRequestModel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScalingAPI.CreateNodePoolScheduledScaling(context.Background(), clusterName, nodePoolName).XAuthToken(xAuthToken).CreateK8sClusterNodePoolScalingScheduleRequestModel(createK8sClusterNodePoolScalingScheduleRequestModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScalingAPI.CreateNodePoolScheduledScaling``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNodePoolScheduledScaling`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ScalingAPI.CreateNodePoolScheduledScaling`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | 대상 클러스터 이름  | 
**nodePoolName** | **string** | 대상 노드 풀 이름 | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNodePoolScheduledScalingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **createK8sClusterNodePoolScalingScheduleRequestModel** | [**CreateK8sClusterNodePoolScalingScheduleRequestModel**](CreateK8sClusterNodePoolScalingScheduleRequestModel.md) |  | 

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


## DeleteNodePoolScheduledScaling

> DeleteNodePoolScheduledScaling(ctx, clusterName, nodePoolName, scheduleName).XAuthToken(xAuthToken).Execute()

Delete node pool scheduled scaling  



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
	clusterName := "clusterName_example" // string | 대상 클러스터 이름 
	nodePoolName := "nodePoolName_example" // string | 대상 노드 풀 이름
	scheduleName := "scheduleName_example" // string | 대상 스케줄 이름
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ScalingAPI.DeleteNodePoolScheduledScaling(context.Background(), clusterName, nodePoolName, scheduleName).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScalingAPI.DeleteNodePoolScheduledScaling``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | 대상 클러스터 이름  | 
**nodePoolName** | **string** | 대상 노드 풀 이름 | 
**scheduleName** | **string** | 대상 스케줄 이름 | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNodePoolScheduledScalingRequest struct via the builder pattern


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


## ListNodePoolScheduledScalings

> GetK8sClusterNodePoolScalingScheduleResponseModel ListNodePoolScheduledScalings(ctx, clusterName, nodePoolName).XAuthToken(xAuthToken).Execute()

List node pool scheduled scalings  



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
	clusterName := "clusterName_example" // string | 대상 클러스터 이름 
	nodePoolName := "nodePoolName_example" // string | 대상 노드 풀 이름
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScalingAPI.ListNodePoolScheduledScalings(context.Background(), clusterName, nodePoolName).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScalingAPI.ListNodePoolScheduledScalings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNodePoolScheduledScalings`: GetK8sClusterNodePoolScalingScheduleResponseModel
	fmt.Fprintf(os.Stdout, "Response from `ScalingAPI.ListNodePoolScheduledScalings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | 대상 클러스터 이름  | 
**nodePoolName** | **string** | 대상 노드 풀 이름 | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNodePoolScheduledScalingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**GetK8sClusterNodePoolScalingScheduleResponseModel**](GetK8sClusterNodePoolScalingScheduleResponseModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetNodePoolResourceBasedAutoScaling

> interface{} SetNodePoolResourceBasedAutoScaling(ctx, clusterName, nodePoolName).XAuthToken(xAuthToken).UpdateKubernetesEngineClusterNodePoolScalingResourceRequestModel(updateKubernetesEngineClusterNodePoolScalingResourceRequestModel).Execute()

Set node pool resource-based auto scaling   



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
	clusterName := "clusterName_example" // string | 대상 클러스터 이름 
	nodePoolName := "nodePoolName_example" // string | 대상 노드 풀 이름
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	updateKubernetesEngineClusterNodePoolScalingResourceRequestModel := *openapiclient.NewUpdateKubernetesEngineClusterNodePoolScalingResourceRequestModel(*openapiclient.NewNodePoolScalingResourceRequestModel(false)) // UpdateKubernetesEngineClusterNodePoolScalingResourceRequestModel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScalingAPI.SetNodePoolResourceBasedAutoScaling(context.Background(), clusterName, nodePoolName).XAuthToken(xAuthToken).UpdateKubernetesEngineClusterNodePoolScalingResourceRequestModel(updateKubernetesEngineClusterNodePoolScalingResourceRequestModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScalingAPI.SetNodePoolResourceBasedAutoScaling``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetNodePoolResourceBasedAutoScaling`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ScalingAPI.SetNodePoolResourceBasedAutoScaling`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | 대상 클러스터 이름  | 
**nodePoolName** | **string** | 대상 노드 풀 이름 | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetNodePoolResourceBasedAutoScalingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **updateKubernetesEngineClusterNodePoolScalingResourceRequestModel** | [**UpdateKubernetesEngineClusterNodePoolScalingResourceRequestModel**](UpdateKubernetesEngineClusterNodePoolScalingResourceRequestModel.md) |  | 

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

