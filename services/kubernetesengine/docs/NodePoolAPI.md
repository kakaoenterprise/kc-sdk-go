# \NodePoolAPI

All URIs are relative to *https://kubernetes-engine.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNodePool**](NodePoolAPI.md#CreateNodePool) | **Post** /api/v1/clusters/{cluster_name}/node-pools | Create node pool
[**CreateNodePoolScheduledScaling**](NodePoolAPI.md#CreateNodePoolScheduledScaling) | **Post** /api/v1/clusters/{cluster_name}/node-pools/{node_pool_name}/scheduled-scaling | Create node pool scheduled scaling 
[**DeleteNodePool**](NodePoolAPI.md#DeleteNodePool) | **Delete** /api/v1/clusters/{cluster_name}/node-pools/{node_pool_name} | Delete node pool 
[**DeleteNodePoolScheduledScaling**](NodePoolAPI.md#DeleteNodePoolScheduledScaling) | **Delete** /api/v1/clusters/{cluster_name}/node-pools/{node_pool_name}/scheduled-scaling/{schedule_name} | Delete node pool scheduled scaling  
[**GetNodePool**](NodePoolAPI.md#GetNodePool) | **Get** /api/v1/clusters/{cluster_name}/node-pools/{node_pool_name} | Get node pool   
[**ListNodePoolNodes**](NodePoolAPI.md#ListNodePoolNodes) | **Get** /api/v1/clusters/{cluster_name}/node-pools/{node_pool_name}/nodes | List node pool nodes  
[**ListNodePoolScheduledScalings**](NodePoolAPI.md#ListNodePoolScheduledScalings) | **Get** /api/v1/clusters/{cluster_name}/node-pools/{node_pool_name}/scheduled-scaling | List node pool scheduled scalings  
[**ListNodePools**](NodePoolAPI.md#ListNodePools) | **Get** /api/v1/clusters/{cluster_name}/node-pools | List node pools  
[**SetNodePoolNodeLabel**](NodePoolAPI.md#SetNodePoolNodeLabel) | **Put** /api/v1/clusters/{cluster_name}/node-pools/{node_pool_name}/labels | Set node pool node label   
[**SetNodePoolResourceBasedAutoScaling**](NodePoolAPI.md#SetNodePoolResourceBasedAutoScaling) | **Put** /api/v1/clusters/{cluster_name}/node-pools/{node_pool_name}/autoscaling | Set node pool resource-based auto scaling   
[**SetNodePoolSecurityGroups**](NodePoolAPI.md#SetNodePoolSecurityGroups) | **Put** /api/v1/clusters/{cluster_name}/node-pools/{node_pool_name}/security-groups | Set node pool security groups 
[**SetNodePoolUserScript**](NodePoolAPI.md#SetNodePoolUserScript) | **Put** /api/v1/clusters/{cluster_name}/node-pools/{node_pool_name}/user-data | Set node pool user script   
[**UpdateNodePool**](NodePoolAPI.md#UpdateNodePool) | **Put** /api/v1/clusters/{cluster_name}/node-pools/{node_pool_name} | Update node pool  
[**UpgradeNodePool**](NodePoolAPI.md#UpgradeNodePool) | **Post** /api/v1/clusters/{cluster_name}/node-pools/{node_pool_name}/upgrade | Upgrade node pool   



## CreateNodePool

> interface{} CreateNodePool(ctx, clusterName).XAuthToken(xAuthToken).CreateNodePoolRequest(createNodePoolRequest).Execute()

Create node pool



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/kubernetesengine"
)

func main() {
	clusterName := "clusterName_example" // string | 대상 클러스터 이름
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	createNodePoolRequest := *openapiclient.NewCreateNodePoolRequest(*openapiclient.NewCreateNodePool("Name_example", "FlavorId_example", int32(123), "SshKeyName_example", *openapiclient.NewVpcInfoRequest("Id_example", []string{"Subnets_example"}), "ImageId_example")) // CreateNodePoolRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodePoolAPI.CreateNodePool(context.Background(), clusterName).XAuthToken(xAuthToken).CreateNodePoolRequest(createNodePoolRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodePoolAPI.CreateNodePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNodePool`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `NodePoolAPI.CreateNodePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | 대상 클러스터 이름 | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNodePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **createNodePoolRequest** | [**CreateNodePoolRequest**](CreateNodePoolRequest.md) |  | 

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


## CreateNodePoolScheduledScaling

> interface{} CreateNodePoolScheduledScaling(ctx, clusterName, nodePoolName).XAuthToken(xAuthToken).CreateNodePoolScheduledScalingRequest(createNodePoolScheduledScalingRequest).Execute()

Create node pool scheduled scaling 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/kubernetesengine"
)

func main() {
	clusterName := "clusterName_example" // string | 대상 클러스터 이름
	nodePoolName := "nodePoolName_example" // string | 대상 노드 풀 이름
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	createNodePoolScheduledScalingRequest := *openapiclient.NewCreateNodePoolScheduledScalingRequest(*openapiclient.NewCreateNodePoolScheduledScaling("Name_example", openapiclient.SchedulingType("cron"), int32(123), "StartTime_example")) // CreateNodePoolScheduledScalingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodePoolAPI.CreateNodePoolScheduledScaling(context.Background(), clusterName, nodePoolName).XAuthToken(xAuthToken).CreateNodePoolScheduledScalingRequest(createNodePoolScheduledScalingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodePoolAPI.CreateNodePoolScheduledScaling``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNodePoolScheduledScaling`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `NodePoolAPI.CreateNodePoolScheduledScaling`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | 대상 클러스터 이름 | 
**nodePoolName** | **string** | 대상 노드 풀 이름 | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNodePoolScheduledScalingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **createNodePoolScheduledScalingRequest** | [**CreateNodePoolScheduledScalingRequest**](CreateNodePoolScheduledScalingRequest.md) |  | 

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


## DeleteNodePool

> DeleteNodePool(ctx, clusterName, nodePoolName).XAuthToken(xAuthToken).Execute()

Delete node pool 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/kubernetesengine"
)

func main() {
	clusterName := "clusterName_example" // string | 대상 클러스터 이름
	nodePoolName := "nodePoolName_example" // string | 대상 노드 풀 이름
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NodePoolAPI.DeleteNodePool(context.Background(), clusterName, nodePoolName).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodePoolAPI.DeleteNodePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | 대상 클러스터 이름 | 
**nodePoolName** | **string** | 대상 노드 풀 이름 | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNodePoolRequest struct via the builder pattern


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
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/kubernetesengine"
)

func main() {
	clusterName := "clusterName_example" // string | 대상 클러스터 이름
	nodePoolName := "nodePoolName_example" // string | 대상 노드 풀 이름
	scheduleName := "scheduleName_example" // string | 대상 스케줄 이름
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NodePoolAPI.DeleteNodePoolScheduledScaling(context.Background(), clusterName, nodePoolName, scheduleName).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodePoolAPI.DeleteNodePoolScheduledScaling``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | 대상 클러스터 이름 | 
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


## GetNodePool

> GetNodePoolResponse GetNodePool(ctx, clusterName, nodePoolName).XAuthToken(xAuthToken).Execute()

Get node pool   



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/kubernetesengine"
)

func main() {
	clusterName := "clusterName_example" // string | 대상 클러스터 이름
	nodePoolName := "nodePoolName_example" // string | 대상 노드 풀 이름 <br/>- [List node pools](https://docs.kakaocloud.com/openapi/container-pack/k8se/list-node-pools)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodePoolAPI.GetNodePool(context.Background(), clusterName, nodePoolName).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodePoolAPI.GetNodePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNodePool`: GetNodePoolResponse
	fmt.Fprintf(os.Stdout, "Response from `NodePoolAPI.GetNodePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | 대상 클러스터 이름 | 
**nodePoolName** | **string** | 대상 노드 풀 이름 &lt;br/&gt;- [List node pools](https://docs.kakaocloud.com/openapi/container-pack/k8se/list-node-pools)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**GetNodePoolResponse**](GetNodePoolResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNodePoolNodes

> ListNodePoolNodesResponse ListNodePoolNodes(ctx, clusterName, nodePoolName).XAuthToken(xAuthToken).Execute()

List node pool nodes  



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/kubernetesengine"
)

func main() {
	clusterName := "clusterName_example" // string | 대상 클러스터 이름
	nodePoolName := "nodePoolName_example" // string | 대상 노드 풀 이름 <br/>- [List node pools](https://docs.kakaocloud.com/openapi/container-pack/k8se/list-node-pools)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodePoolAPI.ListNodePoolNodes(context.Background(), clusterName, nodePoolName).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodePoolAPI.ListNodePoolNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNodePoolNodes`: ListNodePoolNodesResponse
	fmt.Fprintf(os.Stdout, "Response from `NodePoolAPI.ListNodePoolNodes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | 대상 클러스터 이름 | 
**nodePoolName** | **string** | 대상 노드 풀 이름 &lt;br/&gt;- [List node pools](https://docs.kakaocloud.com/openapi/container-pack/k8se/list-node-pools)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNodePoolNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**ListNodePoolNodesResponse**](ListNodePoolNodesResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNodePoolScheduledScalings

> ListNodePoolScheduledScalingsResponse ListNodePoolScheduledScalings(ctx, clusterName, nodePoolName).XAuthToken(xAuthToken).Execute()

List node pool scheduled scalings  



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/kubernetesengine"
)

func main() {
	clusterName := "clusterName_example" // string | 대상 클러스터 이름
	nodePoolName := "nodePoolName_example" // string | 대상 노드 풀 이름 <br/>- [List node pools](https://docs.kakaocloud.com/openapi/container-pack/k8se/list-node-pools)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodePoolAPI.ListNodePoolScheduledScalings(context.Background(), clusterName, nodePoolName).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodePoolAPI.ListNodePoolScheduledScalings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNodePoolScheduledScalings`: ListNodePoolScheduledScalingsResponse
	fmt.Fprintf(os.Stdout, "Response from `NodePoolAPI.ListNodePoolScheduledScalings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | 대상 클러스터 이름 | 
**nodePoolName** | **string** | 대상 노드 풀 이름 &lt;br/&gt;- [List node pools](https://docs.kakaocloud.com/openapi/container-pack/k8se/list-node-pools)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNodePoolScheduledScalingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**ListNodePoolScheduledScalingsResponse**](ListNodePoolScheduledScalingsResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNodePools

> ListNodePoolsResponse ListNodePools(ctx, clusterName).XAuthToken(xAuthToken).Execute()

List node pools  



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/kubernetesengine"
)

func main() {
	clusterName := "clusterName_example" // string | 대상 클러스터 이름
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodePoolAPI.ListNodePools(context.Background(), clusterName).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodePoolAPI.ListNodePools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNodePools`: ListNodePoolsResponse
	fmt.Fprintf(os.Stdout, "Response from `NodePoolAPI.ListNodePools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | 대상 클러스터 이름 | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNodePoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**ListNodePoolsResponse**](ListNodePoolsResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetNodePoolNodeLabel

> SetNodePoolNodeLabelResponse SetNodePoolNodeLabel(ctx, clusterName, nodePoolName).XAuthToken(xAuthToken).SetNodePoolNodeLabelRequest(setNodePoolNodeLabelRequest).Execute()

Set node pool node label   



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/kubernetesengine"
)

func main() {
	clusterName := "clusterName_example" // string | 대상 클러스터 이름
	nodePoolName := "nodePoolName_example" // string | 대상 노드 풀 이름
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	setNodePoolNodeLabelRequest := *openapiclient.NewSetNodePoolNodeLabelRequest(*openapiclient.NewSetNodePoolNodeLabel()) // SetNodePoolNodeLabelRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodePoolAPI.SetNodePoolNodeLabel(context.Background(), clusterName, nodePoolName).XAuthToken(xAuthToken).SetNodePoolNodeLabelRequest(setNodePoolNodeLabelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodePoolAPI.SetNodePoolNodeLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetNodePoolNodeLabel`: SetNodePoolNodeLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `NodePoolAPI.SetNodePoolNodeLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | 대상 클러스터 이름 | 
**nodePoolName** | **string** | 대상 노드 풀 이름 | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetNodePoolNodeLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **setNodePoolNodeLabelRequest** | [**SetNodePoolNodeLabelRequest**](SetNodePoolNodeLabelRequest.md) |  | 

### Return type

[**SetNodePoolNodeLabelResponse**](SetNodePoolNodeLabelResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetNodePoolResourceBasedAutoScaling

> interface{} SetNodePoolResourceBasedAutoScaling(ctx, clusterName, nodePoolName).XAuthToken(xAuthToken).SetNodePoolResourceBasedAutoScalingRequest(setNodePoolResourceBasedAutoScalingRequest).Execute()

Set node pool resource-based auto scaling   



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/kubernetesengine"
)

func main() {
	clusterName := "clusterName_example" // string | 대상 클러스터 이름
	nodePoolName := "nodePoolName_example" // string | 대상 노드 풀 이름
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	setNodePoolResourceBasedAutoScalingRequest := *openapiclient.NewSetNodePoolResourceBasedAutoScalingRequest(*openapiclient.NewSetNodePoolResourceBasedAutoScaling(false)) // SetNodePoolResourceBasedAutoScalingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodePoolAPI.SetNodePoolResourceBasedAutoScaling(context.Background(), clusterName, nodePoolName).XAuthToken(xAuthToken).SetNodePoolResourceBasedAutoScalingRequest(setNodePoolResourceBasedAutoScalingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodePoolAPI.SetNodePoolResourceBasedAutoScaling``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetNodePoolResourceBasedAutoScaling`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `NodePoolAPI.SetNodePoolResourceBasedAutoScaling`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | 대상 클러스터 이름 | 
**nodePoolName** | **string** | 대상 노드 풀 이름 | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetNodePoolResourceBasedAutoScalingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **setNodePoolResourceBasedAutoScalingRequest** | [**SetNodePoolResourceBasedAutoScalingRequest**](SetNodePoolResourceBasedAutoScalingRequest.md) |  | 

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


## SetNodePoolSecurityGroups

> interface{} SetNodePoolSecurityGroups(ctx, clusterName, nodePoolName).XAuthToken(xAuthToken).SetNodePoolSecurityGroupsRequest(setNodePoolSecurityGroupsRequest).Execute()

Set node pool security groups 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/kubernetesengine"
)

func main() {
	clusterName := "clusterName_example" // string | 클러스터 이름
	nodePoolName := "nodePoolName_example" // string | 노드 풀 이름
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	setNodePoolSecurityGroupsRequest := *openapiclient.NewSetNodePoolSecurityGroupsRequest() // SetNodePoolSecurityGroupsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodePoolAPI.SetNodePoolSecurityGroups(context.Background(), clusterName, nodePoolName).XAuthToken(xAuthToken).SetNodePoolSecurityGroupsRequest(setNodePoolSecurityGroupsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodePoolAPI.SetNodePoolSecurityGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetNodePoolSecurityGroups`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `NodePoolAPI.SetNodePoolSecurityGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | 클러스터 이름 | 
**nodePoolName** | **string** | 노드 풀 이름 | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetNodePoolSecurityGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **setNodePoolSecurityGroupsRequest** | [**SetNodePoolSecurityGroupsRequest**](SetNodePoolSecurityGroupsRequest.md) |  | 

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


## SetNodePoolUserScript

> SetNodePoolUserScriptResponse SetNodePoolUserScript(ctx, clusterName, nodePoolName).XAuthToken(xAuthToken).SetNodePoolUserScriptRequest(setNodePoolUserScriptRequest).Execute()

Set node pool user script   



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/kubernetesengine"
)

func main() {
	clusterName := "clusterName_example" // string | 대상 클러스터 이름
	nodePoolName := "nodePoolName_example" // string | 대상 노드 풀 이름
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	setNodePoolUserScriptRequest := *openapiclient.NewSetNodePoolUserScriptRequest(*openapiclient.NewSetNodePoolUserScript("UserData_example")) // SetNodePoolUserScriptRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodePoolAPI.SetNodePoolUserScript(context.Background(), clusterName, nodePoolName).XAuthToken(xAuthToken).SetNodePoolUserScriptRequest(setNodePoolUserScriptRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodePoolAPI.SetNodePoolUserScript``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetNodePoolUserScript`: SetNodePoolUserScriptResponse
	fmt.Fprintf(os.Stdout, "Response from `NodePoolAPI.SetNodePoolUserScript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | 대상 클러스터 이름 | 
**nodePoolName** | **string** | 대상 노드 풀 이름 | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetNodePoolUserScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **setNodePoolUserScriptRequest** | [**SetNodePoolUserScriptRequest**](SetNodePoolUserScriptRequest.md) |  | 

### Return type

[**SetNodePoolUserScriptResponse**](SetNodePoolUserScriptResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNodePool

> interface{} UpdateNodePool(ctx, clusterName, nodePoolName).XAuthToken(xAuthToken).UpdateNodePoolRequest(updateNodePoolRequest).Execute()

Update node pool  



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/kubernetesengine"
)

func main() {
	clusterName := "clusterName_example" // string | 대상 클러스터 이름
	nodePoolName := "nodePoolName_example" // string | 대상 노드 풀 이름
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	updateNodePoolRequest := *openapiclient.NewUpdateNodePoolRequest(*openapiclient.NewUpdateNodePool()) // UpdateNodePoolRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodePoolAPI.UpdateNodePool(context.Background(), clusterName, nodePoolName).XAuthToken(xAuthToken).UpdateNodePoolRequest(updateNodePoolRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodePoolAPI.UpdateNodePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNodePool`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `NodePoolAPI.UpdateNodePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | 대상 클러스터 이름 | 
**nodePoolName** | **string** | 대상 노드 풀 이름 | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNodePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **updateNodePoolRequest** | [**UpdateNodePoolRequest**](UpdateNodePoolRequest.md) |  | 

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


## UpgradeNodePool

> interface{} UpgradeNodePool(ctx, clusterName, nodePoolName).XAuthToken(xAuthToken).Execute()

Upgrade node pool   



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/kubernetesengine"
)

func main() {
	clusterName := "clusterName_example" // string | 대상 클러스터 이름
	nodePoolName := "nodePoolName_example" // string | 대상 노드 풀 이름
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodePoolAPI.UpgradeNodePool(context.Background(), clusterName, nodePoolName).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodePoolAPI.UpgradeNodePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpgradeNodePool`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `NodePoolAPI.UpgradeNodePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | 대상 클러스터 이름 | 
**nodePoolName** | **string** | 대상 노드 풀 이름 | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeNodePoolRequest struct via the builder pattern


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

