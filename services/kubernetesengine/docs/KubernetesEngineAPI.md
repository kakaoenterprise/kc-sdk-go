# \KubernetesEngineAPI

All URIs are relative to *https://kubernetes-engine.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceAgent**](KubernetesEngineAPI.md#CreateServiceAgent) | **Post** /api/v1/service-agents | Create service agent
[**GetKubernetesEngineQuota**](KubernetesEngineAPI.md#GetKubernetesEngineQuota) | **Get** /api/v1/quotas | Get Kubernetes Engine quotas 
[**GetServiceAgent**](KubernetesEngineAPI.md#GetServiceAgent) | **Get** /api/v1/service-agents | Get service agent 
[**ListAvailableKubernetesVersions**](KubernetesEngineAPI.md#ListAvailableKubernetesVersions) | **Get** /api/v1/versions | List available Kubernetes versions
[**ListNodePoolImages**](KubernetesEngineAPI.md#ListNodePoolImages) | **Get** /api/v1/images | List node pool images  



## CreateServiceAgent

> interface{} CreateServiceAgent(ctx).XAuthToken(xAuthToken).Execute()

Create service agent



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
	resp, r, err := apiClient.KubernetesEngineAPI.CreateServiceAgent(context.Background()).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesEngineAPI.CreateServiceAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateServiceAgent`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `KubernetesEngineAPI.CreateServiceAgent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceAgentRequest struct via the builder pattern


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


## GetKubernetesEngineQuota

> GetKubernetesEngineQuotaResponse GetKubernetesEngineQuota(ctx).XAuthToken(xAuthToken).Execute()

Get Kubernetes Engine quotas 



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
	resp, r, err := apiClient.KubernetesEngineAPI.GetKubernetesEngineQuota(context.Background()).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesEngineAPI.GetKubernetesEngineQuota``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKubernetesEngineQuota`: GetKubernetesEngineQuotaResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesEngineAPI.GetKubernetesEngineQuota`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKubernetesEngineQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**GetKubernetesEngineQuotaResponse**](GetKubernetesEngineQuotaResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceAgent

> GetServiceAgentResponse GetServiceAgent(ctx).XAuthToken(xAuthToken).Execute()

Get service agent 



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
	resp, r, err := apiClient.KubernetesEngineAPI.GetServiceAgent(context.Background()).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesEngineAPI.GetServiceAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceAgent`: GetServiceAgentResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesEngineAPI.GetServiceAgent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**GetServiceAgentResponse**](GetServiceAgentResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAvailableKubernetesVersions

> ListAvailableKubernetesVersionsResponse ListAvailableKubernetesVersions(ctx).XAuthToken(xAuthToken).Execute()

List available Kubernetes versions



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
	resp, r, err := apiClient.KubernetesEngineAPI.ListAvailableKubernetesVersions(context.Background()).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesEngineAPI.ListAvailableKubernetesVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAvailableKubernetesVersions`: ListAvailableKubernetesVersionsResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesEngineAPI.ListAvailableKubernetesVersions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAvailableKubernetesVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**ListAvailableKubernetesVersionsResponse**](ListAvailableKubernetesVersionsResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNodePoolImages

> ListNodePoolImagesResponse ListNodePoolImages(ctx).XAuthToken(xAuthToken).OsDistro(osDistro).InstanceType(instanceType).IsGpuType(isGpuType).K8sVersion(k8sVersion).Execute()

List node pool images  



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
	osDistro := "osDistro_example" // string | 운영체제 배포판 (optional)
	instanceType := "instanceType_example" // string | 인스턴스 유형 <br/> - `vm`: Virtual Machine 유형  <br/> - `bm`: Bare Metal Server 유형 (optional)
	isGpuType := true // bool | GPU 지원 이미지 여부 <br/> - `true`: GPU 장비를 지원하는 이미지<br/> - `false`: CPU 전용 이미지 (optional)
	k8sVersion := "k8sVersion_example" // string | 대상 Kubernetes 버전 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesEngineAPI.ListNodePoolImages(context.Background()).XAuthToken(xAuthToken).OsDistro(osDistro).InstanceType(instanceType).IsGpuType(isGpuType).K8sVersion(k8sVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesEngineAPI.ListNodePoolImages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNodePoolImages`: ListNodePoolImagesResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesEngineAPI.ListNodePoolImages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNodePoolImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **osDistro** | **string** | 운영체제 배포판 | 
 **instanceType** | **string** | 인스턴스 유형 &lt;br/&gt; - &#x60;vm&#x60;: Virtual Machine 유형  &lt;br/&gt; - &#x60;bm&#x60;: Bare Metal Server 유형 | 
 **isGpuType** | **bool** | GPU 지원 이미지 여부 &lt;br/&gt; - &#x60;true&#x60;: GPU 장비를 지원하는 이미지&lt;br/&gt; - &#x60;false&#x60;: CPU 전용 이미지 | 
 **k8sVersion** | **string** | 대상 Kubernetes 버전 | 

### Return type

[**ListNodePoolImagesResponse**](ListNodePoolImagesResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

