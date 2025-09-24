# \LoadBalancerListenerAPI

All URIs are relative to *https://load-balancer.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateListener**](LoadBalancerListenerAPI.md#CreateListener) | **Post** /api/v1/load-balancers/listeners | Create listener
[**DeleteListener**](LoadBalancerListenerAPI.md#DeleteListener) | **Delete** /api/v1/load-balancers/listeners/{listener_id} | Delete listener
[**GetListener**](LoadBalancerListenerAPI.md#GetListener) | **Get** /api/v1/load-balancers/listeners/{listener_id} | Get listener
[**ListListeners**](LoadBalancerListenerAPI.md#ListListeners) | **Get** /api/v1/load-balancers/listeners | List listeners
[**UpdateListener**](LoadBalancerListenerAPI.md#UpdateListener) | **Put** /api/v1/load-balancers/listeners/{listener_id} | Update listener



## CreateListener

> BnsLoadBalancerV1ApiCreateListenerModelResponseListenerModel CreateListener(ctx).XAuthToken(xAuthToken).BodyCreateListener(bodyCreateListener).Execute()

Create listener



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
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	bodyCreateListener := *openapiclient.NewBodyCreateListener(*openapiclient.NewCreateListener("LoadBalancerId_example", openapiclient.Protocol("HTTP"), int32(123))) // BodyCreateListener | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadBalancerListenerAPI.CreateListener(context.Background()).XAuthToken(xAuthToken).BodyCreateListener(bodyCreateListener).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerListenerAPI.CreateListener``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateListener`: BnsLoadBalancerV1ApiCreateListenerModelResponseListenerModel
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancerListenerAPI.CreateListener`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateListenerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyCreateListener** | [**BodyCreateListener**](BodyCreateListener.md) |  | 

### Return type

[**BnsLoadBalancerV1ApiCreateListenerModelResponseListenerModel**](BnsLoadBalancerV1ApiCreateListenerModelResponseListenerModel.md)

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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	listenerId := "listenerId_example" // string | 삭제할 리스너의 고유 ID 
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LoadBalancerListenerAPI.DeleteListener(context.Background(), listenerId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerListenerAPI.DeleteListener``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listenerId** | **string** | 삭제할 리스너의 고유 ID  | 

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

> BnsLoadBalancerV1ApiGetListenerModelResponseListenerModel GetListener(ctx, listenerId).XAuthToken(xAuthToken).Execute()

Get listener



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
	listenerId := "listenerId_example" // string | 조회할 리스너 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadBalancerListenerAPI.GetListener(context.Background(), listenerId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerListenerAPI.GetListener``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListener`: BnsLoadBalancerV1ApiGetListenerModelResponseListenerModel
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancerListenerAPI.GetListener`: %v\n", resp)
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

[**BnsLoadBalancerV1ApiGetListenerModelResponseListenerModel**](BnsLoadBalancerV1ApiGetListenerModelResponseListenerModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListListeners

> ListenerListModel ListListeners(ctx).XAuthToken(xAuthToken).Id(id).LoadBalancerId(loadBalancerId).Protocol(protocol).ProtocolPort(protocolPort).ProvisioningStatus(provisioningStatus).OperatingStatus(operatingStatus).SecretName(secretName).SecretId(secretId).TlsCertificateId(tlsCertificateId).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Limit(limit).Offset(offset).Execute()

List listeners



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
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	id := "id_example" // string | 리스너의 고유 ID  (optional)
	loadBalancerId := "loadBalancerId_example" // string | 연결된 로드 밸런서 ID  (optional)
	protocol := openapiclient.Protocol("HTTP") // Protocol | 리스너가 사용하는 프로토콜 <br/> - `HTTP`: HTTP <br/> - `TCP`: TCP <br/> - `UDP`: UDP <br/> - `TERMINATED_HTTPS`: SSL 종료 처리된 HTTPS (optional)
	protocolPort := "protocolPort_example" // string | 리스너가 수신하는 포트 번호  (optional)
	provisioningStatus := openapiclient.ProvisioningStatus("ACTIVE") // ProvisioningStatus | 프로비저닝 상태 <br/> - `ACTIVE`: 활성 <br/> - `DELETED`: 삭제됨 <br/> - `ERROR`: 오류 <br/> - `PENDING_CREATE`: 생성 대기 중 <br/> - `PENDING_UPDATE`: 업데이트 대기 중 <br/> - `PENDING_DELETE`: 삭제 대기 중 (optional)
	operatingStatus := openapiclient.LoadBalancerOperatingStatus("ONLINE") // LoadBalancerOperatingStatus | 운영 상태 <br/> - `ONLINE`: 온라인 <br/> - `DRAINING`: 연결 종료 중 <br/> - `OFFLINE`: 오프라인 <br/> - `DEGRADED`: 성능 저하 <br/> - `ERROR`: 오류 <br/> - `NO_MONITOR`: 모니터링 없음 (optional)
	secretName := "secretName_example" // string | TLS 인증서 이름  (optional)
	secretId := "secretId_example" // string | TLS 인증서 ID  (optional)
	tlsCertificateId := "tlsCertificateId_example" // string | 리스너에 연결된 TLS 인증서 ID  (optional)
	createdAt := "createdAt_example" // string | 리소스가 생성된 시간 <br/> - ISO_8601 형식  <br/> - UTC 기준 (optional)
	updatedAt := "updatedAt_example" // string | 리소스가 마지막으로 수정된 시간 <br/> - ISO_8601 형식  <br/> - UTC 기준 (optional)
	sortKeys := "sortKeys_example" // string | 정렬할 필드를 콤마(,)로 구분   (optional) (default to "created_at")
	sortDirs := "sortDirs_example" // string | 정렬 방향 (`asc`, `desc`)   (optional) (default to "desc")
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional) (default to 20)
	offset := int32(56) // int32 | 조회 시작 위치 (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadBalancerListenerAPI.ListListeners(context.Background()).XAuthToken(xAuthToken).Id(id).LoadBalancerId(loadBalancerId).Protocol(protocol).ProtocolPort(protocolPort).ProvisioningStatus(provisioningStatus).OperatingStatus(operatingStatus).SecretName(secretName).SecretId(secretId).TlsCertificateId(tlsCertificateId).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerListenerAPI.ListListeners``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListListeners`: ListenerListModel
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancerListenerAPI.ListListeners`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListListenersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **id** | **string** | 리스너의 고유 ID  | 
 **loadBalancerId** | **string** | 연결된 로드 밸런서 ID  | 
 **protocol** | [**Protocol**](Protocol.md) | 리스너가 사용하는 프로토콜 &lt;br/&gt; - &#x60;HTTP&#x60;: HTTP &lt;br/&gt; - &#x60;TCP&#x60;: TCP &lt;br/&gt; - &#x60;UDP&#x60;: UDP &lt;br/&gt; - &#x60;TERMINATED_HTTPS&#x60;: SSL 종료 처리된 HTTPS | 
 **protocolPort** | **string** | 리스너가 수신하는 포트 번호  | 
 **provisioningStatus** | [**ProvisioningStatus**](ProvisioningStatus.md) | 프로비저닝 상태 &lt;br/&gt; - &#x60;ACTIVE&#x60;: 활성 &lt;br/&gt; - &#x60;DELETED&#x60;: 삭제됨 &lt;br/&gt; - &#x60;ERROR&#x60;: 오류 &lt;br/&gt; - &#x60;PENDING_CREATE&#x60;: 생성 대기 중 &lt;br/&gt; - &#x60;PENDING_UPDATE&#x60;: 업데이트 대기 중 &lt;br/&gt; - &#x60;PENDING_DELETE&#x60;: 삭제 대기 중 | 
 **operatingStatus** | [**LoadBalancerOperatingStatus**](LoadBalancerOperatingStatus.md) | 운영 상태 &lt;br/&gt; - &#x60;ONLINE&#x60;: 온라인 &lt;br/&gt; - &#x60;DRAINING&#x60;: 연결 종료 중 &lt;br/&gt; - &#x60;OFFLINE&#x60;: 오프라인 &lt;br/&gt; - &#x60;DEGRADED&#x60;: 성능 저하 &lt;br/&gt; - &#x60;ERROR&#x60;: 오류 &lt;br/&gt; - &#x60;NO_MONITOR&#x60;: 모니터링 없음 | 
 **secretName** | **string** | TLS 인증서 이름  | 
 **secretId** | **string** | TLS 인증서 ID  | 
 **tlsCertificateId** | **string** | 리스너에 연결된 TLS 인증서 ID  | 
 **createdAt** | **string** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
 **updatedAt** | **string** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
 **sortKeys** | **string** | 정렬할 필드를 콤마(,)로 구분   | [default to &quot;created_at&quot;]
 **sortDirs** | **string** | 정렬 방향 (&#x60;asc&#x60;, &#x60;desc&#x60;)   | [default to &quot;desc&quot;]
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | [default to 20]
 **offset** | **int32** | 조회 시작 위치 | [default to 0]

### Return type

[**ListenerListModel**](ListenerListModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateListener

> ResponseEditListenerModel UpdateListener(ctx, listenerId).XAuthToken(xAuthToken).BodyUpdateListener(bodyUpdateListener).Execute()

Update listener



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
	listenerId := "listenerId_example" // string | 수정할 리스너 ID 
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	bodyUpdateListener := *openapiclient.NewBodyUpdateListener(*openapiclient.NewEditListener()) // BodyUpdateListener | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadBalancerListenerAPI.UpdateListener(context.Background(), listenerId).XAuthToken(xAuthToken).BodyUpdateListener(bodyUpdateListener).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerListenerAPI.UpdateListener``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateListener`: ResponseEditListenerModel
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancerListenerAPI.UpdateListener`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listenerId** | **string** | 수정할 리스너 ID  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateListenerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyUpdateListener** | [**BodyUpdateListener**](BodyUpdateListener.md) |  | 

### Return type

[**ResponseEditListenerModel**](ResponseEditListenerModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

