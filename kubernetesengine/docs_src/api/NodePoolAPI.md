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

> CreateNodePool(ctx, clusterName).CreateNodePoolRequest(createNodePoolRequest).Execute()

Create node pool



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/kubernetesengine"
)

func main() {
	clusterName := "clusterName_example" // string | 대상 클러스터 이름
	createNodePoolRequest := *openapiclient.NewCreateNodePoolRequest(*openapiclient.NewCreateNodePool("Name_example", "FlavorId_example", int32(123), "SshKeyName_example", *openapiclient.NewVpcInfoRequest("Id_example", []string{"Subnets_example"}), "ImageId_example")) // CreateNodePoolRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NodePoolAPI.CreateNodePool(context.Background(), clusterName).CreateNodePoolRequest(createNodePoolRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodePoolAPI.CreateNodePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
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

 **createNodePoolRequest** | [**CreateNodePoolRequest**](CreateNodePoolRequest.md) |  | 

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


## CreateNodePoolScheduledScaling

> CreateNodePoolScheduledScaling(ctx, clusterName, nodePoolName).CreateNodePoolScheduledScalingRequest(createNodePoolScheduledScalingRequest).Execute()

Create node pool scheduled scaling 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/kubernetesengine"
)

func main() {
	clusterName := "clusterName_example" // string | 대상 클러스터 이름
	nodePoolName := "nodePoolName_example" // string | 대상 노드 풀 이름
	createNodePoolScheduledScalingRequest := *openapiclient.NewCreateNodePoolScheduledScalingRequest(*openapiclient.NewCreateNodePoolScheduledScaling("Name_example", openapiclient.SchedulingType("cron"), int32(123), "StartTime_example")) // CreateNodePoolScheduledScalingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NodePoolAPI.CreateNodePoolScheduledScaling(context.Background(), clusterName, nodePoolName).CreateNodePoolScheduledScalingRequest(createNodePoolScheduledScalingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodePoolAPI.CreateNodePoolScheduledScaling``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCreateNodePoolScheduledScalingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createNodePoolScheduledScalingRequest** | [**CreateNodePoolScheduledScalingRequest**](CreateNodePoolScheduledScalingRequest.md) |  | 

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


## DeleteNodePool

> DeleteNodePool(ctx, clusterName, nodePoolName).Execute()

Delete node pool 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/kubernetesengine"
)

func main() {
	clusterName := "clusterName_example" // string | 대상 클러스터 이름
	nodePoolName := "nodePoolName_example" // string | 대상 노드 풀 이름

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NodePoolAPI.DeleteNodePool(context.Background(), clusterName, nodePoolName).Execute()
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

> DeleteNodePoolScheduledScaling(ctx, clusterName, nodePoolName, scheduleName).Execute()

Delete node pool scheduled scaling  



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/kubernetesengine"
)

func main() {
	clusterName := "clusterName_example" // string | 대상 클러스터 이름
	nodePoolName := "nodePoolName_example" // string | 대상 노드 풀 이름
	scheduleName := "scheduleName_example" // string | 대상 스케줄 이름

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NodePoolAPI.DeleteNodePoolScheduledScaling(context.Background(), clusterName, nodePoolName, scheduleName).Execute()
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

> GetNodePoolResponse GetNodePool(ctx, clusterName, nodePoolName).Execute()

Get node pool   



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/kubernetesengine"
)

func main() {
	clusterName := "clusterName_example" // string | 대상 클러스터 이름
	nodePoolName := "nodePoolName_example" // string | 대상 노드 풀 이름 - [List node pools](/openapi/container-pack/k8se/list-node-pools)에서 확인

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodePoolAPI.GetNodePool(context.Background(), clusterName, nodePoolName).Execute()
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
**nodePoolName** | **string** | 대상 노드 풀 이름 - [List node pools](/openapi/container-pack/k8se/list-node-pools)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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

> ListNodePoolNodesResponse ListNodePoolNodes(ctx, clusterName, nodePoolName).Execute()

List node pool nodes  



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/kubernetesengine"
)

func main() {
	clusterName := "clusterName_example" // string | 대상 클러스터 이름
	nodePoolName := "nodePoolName_example" // string | 대상 노드 풀 이름 - [List node pools](/openapi/container-pack/k8se/list-node-pools)에서 확인

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodePoolAPI.ListNodePoolNodes(context.Background(), clusterName, nodePoolName).Execute()
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
**nodePoolName** | **string** | 대상 노드 풀 이름 - [List node pools](/openapi/container-pack/k8se/list-node-pools)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNodePoolNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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

> ListNodePoolScheduledScalingsResponse ListNodePoolScheduledScalings(ctx, clusterName, nodePoolName).Execute()

List node pool scheduled scalings  



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/kubernetesengine"
)

func main() {
	clusterName := "clusterName_example" // string | 대상 클러스터 이름
	nodePoolName := "nodePoolName_example" // string | 대상 노드 풀 이름 - [List node pools](/openapi/container-pack/k8se/list-node-pools)에서 확인

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodePoolAPI.ListNodePoolScheduledScalings(context.Background(), clusterName, nodePoolName).Execute()
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
**nodePoolName** | **string** | 대상 노드 풀 이름 - [List node pools](/openapi/container-pack/k8se/list-node-pools)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNodePoolScheduledScalingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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

> ListNodePoolsResponse ListNodePools(ctx, clusterName).Execute()

List node pools  



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/kubernetesengine"
)

func main() {
	clusterName := "clusterName_example" // string | 대상 클러스터 이름

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodePoolAPI.ListNodePools(context.Background(), clusterName).Execute()
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

> SetNodePoolNodeLabelResponse SetNodePoolNodeLabel(ctx, clusterName, nodePoolName).SetNodePoolNodeLabelRequest(setNodePoolNodeLabelRequest).Execute()

Set node pool node label   



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/kubernetesengine"
)

func main() {
	clusterName := "clusterName_example" // string | 대상 클러스터 이름
	nodePoolName := "nodePoolName_example" // string | 대상 노드 풀 이름
	setNodePoolNodeLabelRequest := *openapiclient.NewSetNodePoolNodeLabelRequest(*openapiclient.NewSetNodePoolNodeLabel()) // SetNodePoolNodeLabelRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodePoolAPI.SetNodePoolNodeLabel(context.Background(), clusterName, nodePoolName).SetNodePoolNodeLabelRequest(setNodePoolNodeLabelRequest).Execute()
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

> SetNodePoolResourceBasedAutoScaling(ctx, clusterName, nodePoolName).SetNodePoolResourceBasedAutoScalingRequest(setNodePoolResourceBasedAutoScalingRequest).Execute()

Set node pool resource-based auto scaling   



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/kubernetesengine"
)

func main() {
	clusterName := "clusterName_example" // string | 대상 클러스터 이름
	nodePoolName := "nodePoolName_example" // string | 대상 노드 풀 이름
	setNodePoolResourceBasedAutoScalingRequest := *openapiclient.NewSetNodePoolResourceBasedAutoScalingRequest(*openapiclient.NewSetNodePoolResourceBasedAutoScaling(false)) // SetNodePoolResourceBasedAutoScalingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NodePoolAPI.SetNodePoolResourceBasedAutoScaling(context.Background(), clusterName, nodePoolName).SetNodePoolResourceBasedAutoScalingRequest(setNodePoolResourceBasedAutoScalingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodePoolAPI.SetNodePoolResourceBasedAutoScaling``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSetNodePoolResourceBasedAutoScalingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **setNodePoolResourceBasedAutoScalingRequest** | [**SetNodePoolResourceBasedAutoScalingRequest**](SetNodePoolResourceBasedAutoScalingRequest.md) |  | 

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


## SetNodePoolSecurityGroups

> SetNodePoolSecurityGroups(ctx, clusterName, nodePoolName).SetNodePoolSecurityGroupsRequest(setNodePoolSecurityGroupsRequest).Execute()

Set node pool security groups 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/kubernetesengine"
)

func main() {
	clusterName := "clusterName_example" // string | 클러스터 이름
	nodePoolName := "nodePoolName_example" // string | 노드 풀 이름
	setNodePoolSecurityGroupsRequest := *openapiclient.NewSetNodePoolSecurityGroupsRequest() // SetNodePoolSecurityGroupsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NodePoolAPI.SetNodePoolSecurityGroups(context.Background(), clusterName, nodePoolName).SetNodePoolSecurityGroupsRequest(setNodePoolSecurityGroupsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodePoolAPI.SetNodePoolSecurityGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
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


 **setNodePoolSecurityGroupsRequest** | [**SetNodePoolSecurityGroupsRequest**](SetNodePoolSecurityGroupsRequest.md) |  | 

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


## SetNodePoolUserScript

> SetNodePoolUserScriptResponse SetNodePoolUserScript(ctx, clusterName, nodePoolName).SetNodePoolUserScriptRequest(setNodePoolUserScriptRequest).Execute()

Set node pool user script   



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/kubernetesengine"
)

func main() {
	clusterName := "clusterName_example" // string | 대상 클러스터 이름
	nodePoolName := "nodePoolName_example" // string | 대상 노드 풀 이름
	setNodePoolUserScriptRequest := *openapiclient.NewSetNodePoolUserScriptRequest(*openapiclient.NewSetNodePoolUserScript("UserData_example")) // SetNodePoolUserScriptRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodePoolAPI.SetNodePoolUserScript(context.Background(), clusterName, nodePoolName).SetNodePoolUserScriptRequest(setNodePoolUserScriptRequest).Execute()
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

> UpdateNodePool(ctx, clusterName, nodePoolName).UpdateNodePoolRequest(updateNodePoolRequest).Execute()

Update node pool  



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/kubernetesengine"
)

func main() {
	clusterName := "clusterName_example" // string | 대상 클러스터 이름
	nodePoolName := "nodePoolName_example" // string | 대상 노드 풀 이름
	updateNodePoolRequest := *openapiclient.NewUpdateNodePoolRequest(*openapiclient.NewUpdateNodePool()) // UpdateNodePoolRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NodePoolAPI.UpdateNodePool(context.Background(), clusterName, nodePoolName).UpdateNodePoolRequest(updateNodePoolRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodePoolAPI.UpdateNodePool``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateNodePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNodePoolRequest** | [**UpdateNodePoolRequest**](UpdateNodePoolRequest.md) |  | 

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


## UpgradeNodePool

> UpgradeNodePool(ctx, clusterName, nodePoolName).Execute()

Upgrade node pool   



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/kubernetesengine"
)

func main() {
	clusterName := "clusterName_example" // string | 대상 클러스터 이름
	nodePoolName := "nodePoolName_example" // string | 대상 노드 풀 이름

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NodePoolAPI.UpgradeNodePool(context.Background(), clusterName, nodePoolName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodePoolAPI.UpgradeNodePool``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpgradeNodePoolRequest struct via the builder pattern


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

