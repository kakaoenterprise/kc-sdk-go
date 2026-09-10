# \NodeAPI

All URIs are relative to *https://kubernetes-engine.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetClusterNode**](NodeAPI.md#GetClusterNode) | **Get** /api/v1/clusters/{cluster_name}/nodes/{node_name} | Get cluster node  
[**GetClusterNodeDetails**](NodeAPI.md#GetClusterNodeDetails) | **Get** /api/v1/clusters/{cluster_name}/nodes/{node_name}/details | Get cluster node details



## GetClusterNode

> GetClusterNodeResponse GetClusterNode(ctx, clusterName, nodeName).XAuthToken(xAuthToken).Execute()

Get cluster node  



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
	nodeName := "nodeName_example" // string | 대상 노드 이름 <br/>- [List cluster nodes](https://docs.kakaocloud.com/openapi/container-pack/k8se/list-cluster-nodes)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeAPI.GetClusterNode(context.Background(), clusterName, nodeName).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeAPI.GetClusterNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClusterNode`: GetClusterNodeResponse
	fmt.Fprintf(os.Stdout, "Response from `NodeAPI.GetClusterNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | 대상 클러스터 이름 | 
**nodeName** | **string** | 대상 노드 이름 &lt;br/&gt;- [List cluster nodes](https://docs.kakaocloud.com/openapi/container-pack/k8se/list-cluster-nodes)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**GetClusterNodeResponse**](GetClusterNodeResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterNodeDetails

> GetClusterNodeDetailsResponse GetClusterNodeDetails(ctx, clusterName, nodeName).XAuthToken(xAuthToken).Execute()

Get cluster node details



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
	nodeName := "nodeName_example" // string | 대상 노드 이름 <br/>- [List cluster nodes](https://docs.kakaocloud.com/openapi/container-pack/k8se/list-cluster-nodes)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeAPI.GetClusterNodeDetails(context.Background(), clusterName, nodeName).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeAPI.GetClusterNodeDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClusterNodeDetails`: GetClusterNodeDetailsResponse
	fmt.Fprintf(os.Stdout, "Response from `NodeAPI.GetClusterNodeDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | 대상 클러스터 이름 | 
**nodeName** | **string** | 대상 노드 이름 &lt;br/&gt;- [List cluster nodes](https://docs.kakaocloud.com/openapi/container-pack/k8se/list-cluster-nodes)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterNodeDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**GetClusterNodeDetailsResponse**](GetClusterNodeDetailsResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

