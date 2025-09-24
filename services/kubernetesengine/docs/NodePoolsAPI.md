# \NodePoolsAPI

All URIs are relative to *https://kubernetes-engine.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNodePool**](NodePoolsAPI.md#CreateNodePool) | **Post** /api/v1/clusters/{cluster_name}/node-pools | Create node pool
[**DeleteNodePool**](NodePoolsAPI.md#DeleteNodePool) | **Delete** /api/v1/clusters/{cluster_name}/node-pools/{node_pool_name} | Delete node pool 
[**GetNodePool**](NodePoolsAPI.md#GetNodePool) | **Get** /api/v1/clusters/{cluster_name}/node-pools/{node_pool_name} | Get node pool   
[**ListNodePoolNodes**](NodePoolsAPI.md#ListNodePoolNodes) | **Get** /api/v1/clusters/{cluster_name}/node-pools/{node_pool_name}/nodes | List node pool nodes  
[**ListNodePools**](NodePoolsAPI.md#ListNodePools) | **Get** /api/v1/clusters/{cluster_name}/node-pools | List node pools  
[**SetNodePoolNodeLabel**](NodePoolsAPI.md#SetNodePoolNodeLabel) | **Put** /api/v1/clusters/{cluster_name}/node-pools/{node_pool_name}/labels | Set node pool node label   
[**SetNodePoolUserScript**](NodePoolsAPI.md#SetNodePoolUserScript) | **Put** /api/v1/clusters/{cluster_name}/node-pools/{node_pool_name}/user-data | Set node pool user script   
[**UpdateNodePool**](NodePoolsAPI.md#UpdateNodePool) | **Put** /api/v1/clusters/{cluster_name}/node-pools/{node_pool_name} | Update node pool  
[**UpgradeNodePool**](NodePoolsAPI.md#UpgradeNodePool) | **Post** /api/v1/clusters/{cluster_name}/node-pools/{node_pool_name}/upgrade | Upgrade node pool   



## CreateNodePool

> interface{} CreateNodePool(ctx, clusterName).XAuthToken(xAuthToken).CreateK8sClusterNodePoolRequestModel(createK8sClusterNodePoolRequestModel).Execute()

Create node pool



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
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	createK8sClusterNodePoolRequestModel := *openapiclient.NewCreateK8sClusterNodePoolRequestModel(*openapiclient.NewKubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel("Name_example", "FlavorId_example", int32(123), int32(123), "SshKeyName_example", *openapiclient.NewVpcInfoRequestModel("Id_example", []string{"Subnets_example"}), "ImageId_example")) // CreateK8sClusterNodePoolRequestModel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodePoolsAPI.CreateNodePool(context.Background(), clusterName).XAuthToken(xAuthToken).CreateK8sClusterNodePoolRequestModel(createK8sClusterNodePoolRequestModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodePoolsAPI.CreateNodePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNodePool`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `NodePoolsAPI.CreateNodePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | 대상 클러스터 이름  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNodePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **createK8sClusterNodePoolRequestModel** | [**CreateK8sClusterNodePoolRequestModel**](CreateK8sClusterNodePoolRequestModel.md) |  | 

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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	clusterName := "clusterName_example" // string | 대상 클러스터 이름 
	nodePoolName := "nodePoolName_example" // string | 대상 노드 풀 이름
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NodePoolsAPI.DeleteNodePool(context.Background(), clusterName, nodePoolName).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodePoolsAPI.DeleteNodePool``: %v\n", err)
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


## GetNodePool

> GetK8sClusterNodePoolResponseModel GetNodePool(ctx, clusterName, nodePoolName).XAuthToken(xAuthToken).Execute()

Get node pool   



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
	resp, r, err := apiClient.NodePoolsAPI.GetNodePool(context.Background(), clusterName, nodePoolName).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodePoolsAPI.GetNodePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNodePool`: GetK8sClusterNodePoolResponseModel
	fmt.Fprintf(os.Stdout, "Response from `NodePoolsAPI.GetNodePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | 대상 클러스터 이름  | 
**nodePoolName** | **string** | 대상 노드 풀 이름 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**GetK8sClusterNodePoolResponseModel**](GetK8sClusterNodePoolResponseModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNodePoolNodes

> GetK8sClusterNodePoolNodesResponseModel ListNodePoolNodes(ctx, clusterName, nodePoolName).XAuthToken(xAuthToken).Execute()

List node pool nodes  



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
	resp, r, err := apiClient.NodePoolsAPI.ListNodePoolNodes(context.Background(), clusterName, nodePoolName).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodePoolsAPI.ListNodePoolNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNodePoolNodes`: GetK8sClusterNodePoolNodesResponseModel
	fmt.Fprintf(os.Stdout, "Response from `NodePoolsAPI.ListNodePoolNodes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | 대상 클러스터 이름  | 
**nodePoolName** | **string** | 대상 노드 풀 이름 | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNodePoolNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**GetK8sClusterNodePoolNodesResponseModel**](GetK8sClusterNodePoolNodesResponseModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNodePools

> GetK8sClusterNodePoolsResponseModel ListNodePools(ctx, clusterName).XAuthToken(xAuthToken).Execute()

List node pools  



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
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodePoolsAPI.ListNodePools(context.Background(), clusterName).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodePoolsAPI.ListNodePools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNodePools`: GetK8sClusterNodePoolsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `NodePoolsAPI.ListNodePools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | 대상 클러스터 이름  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNodePoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**GetK8sClusterNodePoolsResponseModel**](GetK8sClusterNodePoolsResponseModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetNodePoolNodeLabel

> UpdateK8sClusterNodePoolNodeLabelsResponseModel SetNodePoolNodeLabel(ctx, clusterName, nodePoolName).XAuthToken(xAuthToken).UpdateK8sClusterNodePoolNodeLabelsRequestModel(updateK8sClusterNodePoolNodeLabelsRequestModel).Execute()

Set node pool node label   



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
	updateK8sClusterNodePoolNodeLabelsRequestModel := *openapiclient.NewUpdateK8sClusterNodePoolNodeLabelsRequestModel(*openapiclient.NewNodeLabelsRequestModel()) // UpdateK8sClusterNodePoolNodeLabelsRequestModel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodePoolsAPI.SetNodePoolNodeLabel(context.Background(), clusterName, nodePoolName).XAuthToken(xAuthToken).UpdateK8sClusterNodePoolNodeLabelsRequestModel(updateK8sClusterNodePoolNodeLabelsRequestModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodePoolsAPI.SetNodePoolNodeLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetNodePoolNodeLabel`: UpdateK8sClusterNodePoolNodeLabelsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `NodePoolsAPI.SetNodePoolNodeLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | 대상 클러스터 이름  | 
**nodePoolName** | **string** | 대상 노드 풀 이름 | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetNodePoolNodeLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **updateK8sClusterNodePoolNodeLabelsRequestModel** | [**UpdateK8sClusterNodePoolNodeLabelsRequestModel**](UpdateK8sClusterNodePoolNodeLabelsRequestModel.md) |  | 

### Return type

[**UpdateK8sClusterNodePoolNodeLabelsResponseModel**](UpdateK8sClusterNodePoolNodeLabelsResponseModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetNodePoolUserScript

> UpdateK8sClusterNodePoolUserScriptResponseModel SetNodePoolUserScript(ctx, clusterName, nodePoolName).XAuthToken(xAuthToken).UpdateK8sClusterNodePoolUserScriptRequestModel(updateK8sClusterNodePoolUserScriptRequestModel).Execute()

Set node pool user script   



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
	updateK8sClusterNodePoolUserScriptRequestModel := *openapiclient.NewUpdateK8sClusterNodePoolUserScriptRequestModel(*openapiclient.NewNodePoolScriptRequestModel("UserData_example")) // UpdateK8sClusterNodePoolUserScriptRequestModel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodePoolsAPI.SetNodePoolUserScript(context.Background(), clusterName, nodePoolName).XAuthToken(xAuthToken).UpdateK8sClusterNodePoolUserScriptRequestModel(updateK8sClusterNodePoolUserScriptRequestModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodePoolsAPI.SetNodePoolUserScript``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetNodePoolUserScript`: UpdateK8sClusterNodePoolUserScriptResponseModel
	fmt.Fprintf(os.Stdout, "Response from `NodePoolsAPI.SetNodePoolUserScript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | 대상 클러스터 이름  | 
**nodePoolName** | **string** | 대상 노드 풀 이름 | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetNodePoolUserScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **updateK8sClusterNodePoolUserScriptRequestModel** | [**UpdateK8sClusterNodePoolUserScriptRequestModel**](UpdateK8sClusterNodePoolUserScriptRequestModel.md) |  | 

### Return type

[**UpdateK8sClusterNodePoolUserScriptResponseModel**](UpdateK8sClusterNodePoolUserScriptResponseModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNodePool

> interface{} UpdateNodePool(ctx, clusterName, nodePoolName).XAuthToken(xAuthToken).UpdateK8sClusterNodePoolRequestModel(updateK8sClusterNodePoolRequestModel).Execute()

Update node pool  



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
	updateK8sClusterNodePoolRequestModel := *openapiclient.NewUpdateK8sClusterNodePoolRequestModel(*openapiclient.NewKubernetesEngineV1ApiUpdateNodePoolModelNodePoolRequestModel()) // UpdateK8sClusterNodePoolRequestModel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodePoolsAPI.UpdateNodePool(context.Background(), clusterName, nodePoolName).XAuthToken(xAuthToken).UpdateK8sClusterNodePoolRequestModel(updateK8sClusterNodePoolRequestModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodePoolsAPI.UpdateNodePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNodePool`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `NodePoolsAPI.UpdateNodePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | 대상 클러스터 이름  | 
**nodePoolName** | **string** | 대상 노드 풀 이름 | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNodePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **updateK8sClusterNodePoolRequestModel** | [**UpdateK8sClusterNodePoolRequestModel**](UpdateK8sClusterNodePoolRequestModel.md) |  | 

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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	clusterName := "clusterName_example" // string | 대상 클러스터 이름 
	nodePoolName := "nodePoolName_example" // string | 대상 노드 풀 이름
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodePoolsAPI.UpgradeNodePool(context.Background(), clusterName, nodePoolName).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodePoolsAPI.UpgradeNodePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpgradeNodePool`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `NodePoolsAPI.UpgradeNodePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | 대상 클러스터 이름  | 
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

