# \KubernetesEngineClusterAPI

All URIs are relative to *https://kubernetes-engine.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCluster**](KubernetesEngineClusterAPI.md#CreateCluster) | **Post** /api/v1/clusters | Create cluster
[**DeleteCluster**](KubernetesEngineClusterAPI.md#DeleteCluster) | **Delete** /api/v1/clusters/{cluster_name} | Delete cluster  
[**DeleteClusterNodes**](KubernetesEngineClusterAPI.md#DeleteClusterNodes) | **Delete** /api/v1/clusters/{cluster_name}/nodes | Delete cluster nodes 
[**GetCluster**](KubernetesEngineClusterAPI.md#GetCluster) | **Get** /api/v1/clusters/{cluster_name} | Get cluster  
[**GetClusterKubeconfig**](KubernetesEngineClusterAPI.md#GetClusterKubeconfig) | **Get** /api/v1/clusters/{cluster_name}/kubeconfig | Get cluster kubeconfig
[**GetClusterQuota**](KubernetesEngineClusterAPI.md#GetClusterQuota) | **Get** /api/v1/quotas/{cluster_name} | Get cluster quotas 
[**ListClusterNodes**](KubernetesEngineClusterAPI.md#ListClusterNodes) | **Get** /api/v1/clusters/{cluster_name}/nodes | List cluster nodes
[**ListClusterUpgradableVersions**](KubernetesEngineClusterAPI.md#ListClusterUpgradableVersions) | **Get** /api/v1/clusters/{cluster_name}/upgrade | List cluster upgradable versions
[**ListClusters**](KubernetesEngineClusterAPI.md#ListClusters) | **Get** /api/v1/clusters | List clusters
[**SetClusterNodesCordon**](KubernetesEngineClusterAPI.md#SetClusterNodesCordon) | **Post** /api/v1/clusters/{cluster_name}/nodes/cordon | Set cluster nodes cordon
[**UpdateCluster**](KubernetesEngineClusterAPI.md#UpdateCluster) | **Put** /api/v1/clusters/{cluster_name} | Update cluster
[**UpgradeCluster**](KubernetesEngineClusterAPI.md#UpgradeCluster) | **Post** /api/v1/clusters/{cluster_name}/upgrade | Upgrade cluster  



## CreateCluster

> interface{} CreateCluster(ctx).XAuthToken(xAuthToken).CreateClusterRequest(createClusterRequest).Execute()

Create cluster



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
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	createClusterRequest := *openapiclient.NewCreateClusterRequest(*openapiclient.NewCreateCluster("Name_example", "Version_example", *openapiclient.NewVpcInfoRequest("Id_example", []string{"Subnets_example"}), false, *openapiclient.NewClusterNetworkRequest(openapiclient.ClusterNetworkCNI("cilium")))) // CreateClusterRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesEngineClusterAPI.CreateCluster(context.Background()).XAuthToken(xAuthToken).CreateClusterRequest(createClusterRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesEngineClusterAPI.CreateCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCluster`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `KubernetesEngineClusterAPI.CreateCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **createClusterRequest** | [**CreateClusterRequest**](CreateClusterRequest.md) |  | 

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


## DeleteCluster

> DeleteCluster(ctx, clusterName).XAuthToken(xAuthToken).Execute()

Delete cluster  



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
	r, err := apiClient.KubernetesEngineClusterAPI.DeleteCluster(context.Background(), clusterName).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesEngineClusterAPI.DeleteCluster``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteClusterRequest struct via the builder pattern


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


## DeleteClusterNodes

> DeleteClusterNodes(ctx, clusterName).XAuthToken(xAuthToken).DeleteClusterNodesRequest(deleteClusterNodesRequest).Execute()

Delete cluster nodes 



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
	deleteClusterNodesRequest := *openapiclient.NewDeleteClusterNodesRequest(*openapiclient.NewDeleteClusterNodes(false, []string{"NodeNames_example"})) // DeleteClusterNodesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KubernetesEngineClusterAPI.DeleteClusterNodes(context.Background(), clusterName).XAuthToken(xAuthToken).DeleteClusterNodesRequest(deleteClusterNodesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesEngineClusterAPI.DeleteClusterNodes``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteClusterNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **deleteClusterNodesRequest** | [**DeleteClusterNodesRequest**](DeleteClusterNodesRequest.md) |  | 

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


## GetCluster

> GetClusterResponse GetCluster(ctx, clusterName).XAuthToken(xAuthToken).Execute()

Get cluster  



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
	resp, r, err := apiClient.KubernetesEngineClusterAPI.GetCluster(context.Background(), clusterName).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesEngineClusterAPI.GetCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCluster`: GetClusterResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesEngineClusterAPI.GetCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | 대상 클러스터 이름 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**GetClusterResponse**](GetClusterResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterKubeconfig

> string GetClusterKubeconfig(ctx, clusterName).XAuthToken(xAuthToken).Execute()

Get cluster kubeconfig



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
	resp, r, err := apiClient.KubernetesEngineClusterAPI.GetClusterKubeconfig(context.Background(), clusterName).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesEngineClusterAPI.GetClusterKubeconfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClusterKubeconfig`: string
	fmt.Fprintf(os.Stdout, "Response from `KubernetesEngineClusterAPI.GetClusterKubeconfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | 대상 클러스터 이름 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterKubeconfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

**string**

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/yaml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterQuota

> GetClusterQuotaResponse GetClusterQuota(ctx, clusterName).XAuthToken(xAuthToken).Execute()

Get cluster quotas 



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
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesEngineClusterAPI.GetClusterQuota(context.Background(), clusterName).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesEngineClusterAPI.GetClusterQuota``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClusterQuota`: GetClusterQuotaResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesEngineClusterAPI.GetClusterQuota`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | 클러스터 이름 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**GetClusterQuotaResponse**](GetClusterQuotaResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListClusterNodes

> ListClusterNodesResponse ListClusterNodes(ctx, clusterName).XAuthToken(xAuthToken).Execute()

List cluster nodes



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
	resp, r, err := apiClient.KubernetesEngineClusterAPI.ListClusterNodes(context.Background(), clusterName).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesEngineClusterAPI.ListClusterNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListClusterNodes`: ListClusterNodesResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesEngineClusterAPI.ListClusterNodes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | 대상 클러스터 이름 | 

### Other Parameters

Other parameters are passed through a pointer to a apiListClusterNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**ListClusterNodesResponse**](ListClusterNodesResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListClusterUpgradableVersions

> ListClusterUpgradableVersionsResponse ListClusterUpgradableVersions(ctx, clusterName).XAuthToken(xAuthToken).Execute()

List cluster upgradable versions



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
	resp, r, err := apiClient.KubernetesEngineClusterAPI.ListClusterUpgradableVersions(context.Background(), clusterName).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesEngineClusterAPI.ListClusterUpgradableVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListClusterUpgradableVersions`: ListClusterUpgradableVersionsResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesEngineClusterAPI.ListClusterUpgradableVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | 대상 클러스터 이름 | 

### Other Parameters

Other parameters are passed through a pointer to a apiListClusterUpgradableVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**ListClusterUpgradableVersionsResponse**](ListClusterUpgradableVersionsResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListClusters

> ListClustersResponse ListClusters(ctx).XAuthToken(xAuthToken).Execute()

List clusters



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
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesEngineClusterAPI.ListClusters(context.Background()).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesEngineClusterAPI.ListClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListClusters`: ListClustersResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesEngineClusterAPI.ListClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**ListClustersResponse**](ListClustersResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetClusterNodesCordon

> SetClusterNodesCordon(ctx, clusterName).XAuthToken(xAuthToken).SetClusterNodesCordonRequest(setClusterNodesCordonRequest).Execute()

Set cluster nodes cordon



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
	setClusterNodesCordonRequest := *openapiclient.NewSetClusterNodesCordonRequest(*openapiclient.NewSetClusterNodesCordon(false, []string{"NodeNames_example"})) // SetClusterNodesCordonRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KubernetesEngineClusterAPI.SetClusterNodesCordon(context.Background(), clusterName).XAuthToken(xAuthToken).SetClusterNodesCordonRequest(setClusterNodesCordonRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesEngineClusterAPI.SetClusterNodesCordon``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSetClusterNodesCordonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **setClusterNodesCordonRequest** | [**SetClusterNodesCordonRequest**](SetClusterNodesCordonRequest.md) |  | 

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


## UpdateCluster

> interface{} UpdateCluster(ctx, clusterName).XAuthToken(xAuthToken).UpdateClusterRequest(updateClusterRequest).Execute()

Update cluster



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
	updateClusterRequest := *openapiclient.NewUpdateClusterRequest(*openapiclient.NewUpdateCluster("Description_example")) // UpdateClusterRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesEngineClusterAPI.UpdateCluster(context.Background(), clusterName).XAuthToken(xAuthToken).UpdateClusterRequest(updateClusterRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesEngineClusterAPI.UpdateCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCluster`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `KubernetesEngineClusterAPI.UpdateCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | 대상 클러스터 이름 | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **updateClusterRequest** | [**UpdateClusterRequest**](UpdateClusterRequest.md) |  | 

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


## UpgradeCluster

> interface{} UpgradeCluster(ctx, clusterName).XAuthToken(xAuthToken).Execute()

Upgrade cluster  



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
	resp, r, err := apiClient.KubernetesEngineClusterAPI.UpgradeCluster(context.Background(), clusterName).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesEngineClusterAPI.UpgradeCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpgradeCluster`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `KubernetesEngineClusterAPI.UpgradeCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | 대상 클러스터 이름 | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeClusterRequest struct via the builder pattern


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

