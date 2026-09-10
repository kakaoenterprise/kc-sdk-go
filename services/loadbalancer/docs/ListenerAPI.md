# \ListenerAPI

All URIs are relative to *https://load-balancer.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateListener**](ListenerAPI.md#CreateListener) | **Post** /api/v1/load-balancers/listeners | Create listener
[**DeleteListener**](ListenerAPI.md#DeleteListener) | **Delete** /api/v1/load-balancers/listeners/{listener_id} | Delete listener
[**GetListener**](ListenerAPI.md#GetListener) | **Get** /api/v1/load-balancers/listeners/{listener_id} | Get listener
[**ListListeners**](ListenerAPI.md#ListListeners) | **Get** /api/v1/load-balancers/listeners | List listeners
[**UpdateListener**](ListenerAPI.md#UpdateListener) | **Put** /api/v1/load-balancers/listeners/{listener_id} | Update listener



## CreateListener

> CreateListenerResponse CreateListener(ctx).XAuthToken(xAuthToken).CreateListenerRequest(createListenerRequest).Execute()

Create listener



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/loadbalancer"
)

func main() {
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	createListenerRequest := *openapiclient.NewCreateListenerRequest(*openapiclient.NewCreateListener("LoadBalancerId_example", openapiclient.Protocol("HTTP"), int32(123))) // CreateListenerRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListenerAPI.CreateListener(context.Background()).XAuthToken(xAuthToken).CreateListenerRequest(createListenerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListenerAPI.CreateListener``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateListener`: CreateListenerResponse
	fmt.Fprintf(os.Stdout, "Response from `ListenerAPI.CreateListener`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateListenerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **createListenerRequest** | [**CreateListenerRequest**](CreateListenerRequest.md) |  | 

### Return type

[**CreateListenerResponse**](CreateListenerResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteListener

> DeleteListener(ctx, listenerId).XAuthToken(xAuthToken).Execute()

Delete listener



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/loadbalancer"
)

func main() {
	listenerId := "listenerId_example" // string | 삭제할 리스너의 고유 ID <br/>- [List listeners](https://docs.kakaocloud.com/openapi/networking/lb/list-listeners)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ListenerAPI.DeleteListener(context.Background(), listenerId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListenerAPI.DeleteListener``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listenerId** | **string** | 삭제할 리스너의 고유 ID &lt;br/&gt;- [List listeners](https://docs.kakaocloud.com/openapi/networking/lb/list-listeners)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteListenerRequest struct via the builder pattern


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


## GetListener

> GetListenerResponse GetListener(ctx, listenerId).XAuthToken(xAuthToken).Execute()

Get listener



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/loadbalancer"
)

func main() {
	listenerId := "listenerId_example" // string | 조회할 리스너 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListenerAPI.GetListener(context.Background(), listenerId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListenerAPI.GetListener``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListener`: GetListenerResponse
	fmt.Fprintf(os.Stdout, "Response from `ListenerAPI.GetListener`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listenerId** | **string** | 조회할 리스너 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListenerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**GetListenerResponse**](GetListenerResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListListeners

> ListListenersResponse ListListeners(ctx).XAuthToken(xAuthToken).LoadBalancerId(loadBalancerId).Id(id).Protocol(protocol).ProtocolPort(protocolPort).ProvisioningStatus(provisioningStatus).OperatingStatus(operatingStatus).SecretName(secretName).SecretId(secretId).TlsCertificateId(tlsCertificateId).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Offset(offset).Limit(limit).Execute()

List listeners



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/loadbalancer"
)

func main() {
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	loadBalancerId := "loadBalancerId_example" // string | 연결된 로드 밸런서 ID (optional)
	id := "id_example" // string | 리스너의 고유 ID (optional)
	protocol := openapiclient.Protocol("HTTP") // Protocol | 리스너가 사용하는 프로토콜 (optional)
	protocolPort := "protocolPort_example" // string | 리스너가 수신하는 포트 번호 (optional)
	provisioningStatus := openapiclient.ProvisioningStatus("ACTIVE") // ProvisioningStatus | 프로비저닝 상태 (optional)
	operatingStatus := openapiclient.LoadBalancerOperatingStatus("ONLINE") // LoadBalancerOperatingStatus | 운영 상태 (optional)
	secretName := "secretName_example" // string | TLS 인증서 이름 (optional)
	secretId := "secretId_example" // string | TLS 인증서 ID (optional)
	tlsCertificateId := "tlsCertificateId_example" // string | 리스너에 연결된 TLS 인증서 ID (optional)
	createdAt := "createdAt_example" // string | 리소스가 생성된 시간 <br/> - ISO 8601 형식 <br/> - UTC 기준 (optional)
	updatedAt := "updatedAt_example" // string | 리소스가 마지막으로 수정된 시간 <br/> - ISO 8601 형식 <br/> - UTC 기준 (optional)
	sortKeys := "sortKeys_example" // string | 정렬할 필드를 콤마(,)로 구분   (optional)
	sortDirs := "sortDirs_example" // string | 정렬 방향 (`asc`, `desc`)  (optional)
	offset := int32(56) // int32 | 조회 시작 위치 (optional)
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListenerAPI.ListListeners(context.Background()).XAuthToken(xAuthToken).LoadBalancerId(loadBalancerId).Id(id).Protocol(protocol).ProtocolPort(protocolPort).ProvisioningStatus(provisioningStatus).OperatingStatus(operatingStatus).SecretName(secretName).SecretId(secretId).TlsCertificateId(tlsCertificateId).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListenerAPI.ListListeners``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListListeners`: ListListenersResponse
	fmt.Fprintf(os.Stdout, "Response from `ListenerAPI.ListListeners`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListListenersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **loadBalancerId** | **string** | 연결된 로드 밸런서 ID | 
 **id** | **string** | 리스너의 고유 ID | 
 **protocol** | [**Protocol**](Protocol.md) | 리스너가 사용하는 프로토콜 | 
 **protocolPort** | **string** | 리스너가 수신하는 포트 번호 | 
 **provisioningStatus** | [**ProvisioningStatus**](ProvisioningStatus.md) | 프로비저닝 상태 | 
 **operatingStatus** | [**LoadBalancerOperatingStatus**](LoadBalancerOperatingStatus.md) | 운영 상태 | 
 **secretName** | **string** | TLS 인증서 이름 | 
 **secretId** | **string** | TLS 인증서 ID | 
 **tlsCertificateId** | **string** | 리스너에 연결된 TLS 인증서 ID | 
 **createdAt** | **string** | 리소스가 생성된 시간 &lt;br/&gt; - ISO 8601 형식 &lt;br/&gt; - UTC 기준 | 
 **updatedAt** | **string** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO 8601 형식 &lt;br/&gt; - UTC 기준 | 
 **sortKeys** | **string** | 정렬할 필드를 콤마(,)로 구분   | 
 **sortDirs** | **string** | 정렬 방향 (&#x60;asc&#x60;, &#x60;desc&#x60;)  | 
 **offset** | **int32** | 조회 시작 위치 | 
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | 

### Return type

[**ListListenersResponse**](ListListenersResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateListener

> UpdateListenerResponse UpdateListener(ctx, listenerId).XAuthToken(xAuthToken).UpdateListenerRequest(updateListenerRequest).Execute()

Update listener



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/loadbalancer"
)

func main() {
	listenerId := "listenerId_example" // string | 수정할 리스너 ID <br/>- [List listeners](https://docs.kakaocloud.com/openapi/networking/lb/list-listeners)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	updateListenerRequest := *openapiclient.NewUpdateListenerRequest(*openapiclient.NewUpdateListener()) // UpdateListenerRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListenerAPI.UpdateListener(context.Background(), listenerId).XAuthToken(xAuthToken).UpdateListenerRequest(updateListenerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListenerAPI.UpdateListener``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateListener`: UpdateListenerResponse
	fmt.Fprintf(os.Stdout, "Response from `ListenerAPI.UpdateListener`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listenerId** | **string** | 수정할 리스너 ID &lt;br/&gt;- [List listeners](https://docs.kakaocloud.com/openapi/networking/lb/list-listeners)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateListenerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **updateListenerRequest** | [**UpdateListenerRequest**](UpdateListenerRequest.md) |  | 

### Return type

[**UpdateListenerResponse**](UpdateListenerResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

