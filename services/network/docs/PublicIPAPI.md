# \PublicIPAPI

All URIs are relative to *https://network.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePublicIp**](PublicIPAPI.md#CreatePublicIp) | **Post** /api/v1/public-ips | Create public IP
[**DeletePublicIp**](PublicIPAPI.md#DeletePublicIp) | **Delete** /api/v1/public-ips/{public_ip_id} | Delete public IP
[**GetPublicIp**](PublicIPAPI.md#GetPublicIp) | **Get** /api/v1/public-ips/{public_ip_id} | Get public IP
[**ListPublicIps**](PublicIPAPI.md#ListPublicIps) | **Get** /api/v1/public-ips | List public IPs
[**UpdatePublicIp**](PublicIPAPI.md#UpdatePublicIp) | **Put** /api/v1/public-ips/{public_ip_id} | Update public IP



## CreatePublicIp

> CreatePublicIpResponse CreatePublicIp(ctx).XAuthToken(xAuthToken).CreatePublicIpRequest(createPublicIpRequest).Execute()

Create public IP



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/network"
)

func main() {
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	createPublicIpRequest := *openapiclient.NewCreatePublicIpRequest(*openapiclient.NewCreatePublicIp()) // CreatePublicIpRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicIPAPI.CreatePublicIp(context.Background()).XAuthToken(xAuthToken).CreatePublicIpRequest(createPublicIpRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicIPAPI.CreatePublicIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePublicIp`: CreatePublicIpResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicIPAPI.CreatePublicIp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePublicIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **createPublicIpRequest** | [**CreatePublicIpRequest**](CreatePublicIpRequest.md) |  | 

### Return type

[**CreatePublicIpResponse**](CreatePublicIpResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePublicIp

> DeletePublicIp(ctx, publicIpId).XAuthToken(xAuthToken).Execute()

Delete public IP



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/network"
)

func main() {
	publicIpId := "publicIpId_example" // string | 삭제할 퍼블릭 IP ID <br/>- [List public IPs](https://docs.kakaocloud.com/openapi/networking/vpc/list-public-ips)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PublicIPAPI.DeletePublicIp(context.Background(), publicIpId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicIPAPI.DeletePublicIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicIpId** | **string** | 삭제할 퍼블릭 IP ID &lt;br/&gt;- [List public IPs](https://docs.kakaocloud.com/openapi/networking/vpc/list-public-ips)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePublicIpRequest struct via the builder pattern


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


## GetPublicIp

> GetPublicIpResponse GetPublicIp(ctx, publicIpId).XAuthToken(xAuthToken).Execute()

Get public IP



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/network"
)

func main() {
	publicIpId := "publicIpId_example" // string | 조회할 퍼블릭 IP ID 
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicIPAPI.GetPublicIp(context.Background(), publicIpId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicIPAPI.GetPublicIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicIp`: GetPublicIpResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicIPAPI.GetPublicIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicIpId** | **string** | 조회할 퍼블릭 IP ID  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**GetPublicIpResponse**](GetPublicIpResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPublicIps

> ListPublicIpsResponse ListPublicIps(ctx).XAuthToken(xAuthToken).Id(id).Status(status).RelatedResourceName(relatedResourceName).PublicIp(publicIp).RelatedResourceId(relatedResourceId).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Offset(offset).Limit(limit).Execute()

List public IPs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/network"
)

func main() {
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	id := "id_example" // string | 퍼블릭 IP의 고유 ID (optional)
	status := openapiclient.PublicIpStatus("available") // PublicIpStatus | 퍼블릭 IP 상태 (optional)
	relatedResourceName := "relatedResourceName_example" // string | 퍼블릭 IP가 연결된 리소스 이름 (예: 네트워크 인터페이스 이름 등) (optional)
	publicIp := "publicIp_example" // string | 퍼블릭 IP 주소 (optional)
	relatedResourceId := "relatedResourceId_example" // string | 퍼블릭 IP가 연결된 리소스 ID (예: 네트워크 인터페이스 ID 등) (optional)
	createdAt := "createdAt_example" // string | 리소스가 생성된 시간 <br/> - ISO 8601 형식 <br/> - UTC 기준 (optional)
	updatedAt := "updatedAt_example" // string | 리소스가 마지막으로 수정된 시간 <br/> - ISO 8601 형식 <br/> - UTC 기준 (optional)
	sortKeys := "sortKeys_example" // string | 정렬할 필드를 콤마(,)로 구분 (optional)
	sortDirs := "sortDirs_example" // string | 정렬 방향 (`asc`, `desc`)  (optional)
	offset := int32(56) // int32 | 조회 시작 위치 (optional)
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicIPAPI.ListPublicIps(context.Background()).XAuthToken(xAuthToken).Id(id).Status(status).RelatedResourceName(relatedResourceName).PublicIp(publicIp).RelatedResourceId(relatedResourceId).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicIPAPI.ListPublicIps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPublicIps`: ListPublicIpsResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicIPAPI.ListPublicIps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPublicIpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **id** | **string** | 퍼블릭 IP의 고유 ID | 
 **status** | [**PublicIpStatus**](PublicIpStatus.md) | 퍼블릭 IP 상태 | 
 **relatedResourceName** | **string** | 퍼블릭 IP가 연결된 리소스 이름 (예: 네트워크 인터페이스 이름 등) | 
 **publicIp** | **string** | 퍼블릭 IP 주소 | 
 **relatedResourceId** | **string** | 퍼블릭 IP가 연결된 리소스 ID (예: 네트워크 인터페이스 ID 등) | 
 **createdAt** | **string** | 리소스가 생성된 시간 &lt;br/&gt; - ISO 8601 형식 &lt;br/&gt; - UTC 기준 | 
 **updatedAt** | **string** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO 8601 형식 &lt;br/&gt; - UTC 기준 | 
 **sortKeys** | **string** | 정렬할 필드를 콤마(,)로 구분 | 
 **sortDirs** | **string** | 정렬 방향 (&#x60;asc&#x60;, &#x60;desc&#x60;)  | 
 **offset** | **int32** | 조회 시작 위치 | 
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | 

### Return type

[**ListPublicIpsResponse**](ListPublicIpsResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePublicIp

> UpdatePublicIpResponse UpdatePublicIp(ctx, publicIpId).XAuthToken(xAuthToken).UpdatePublicIpRequest(updatePublicIpRequest).Execute()

Update public IP



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/network"
)

func main() {
	publicIpId := "publicIpId_example" // string | 퍼블릭 IP의 고유 ID <br/>- [List public IPs](https://docs.kakaocloud.com/openapi/networking/vpc/list-public-ips)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	updatePublicIpRequest := *openapiclient.NewUpdatePublicIpRequest(*openapiclient.NewUpdatePublicIp()) // UpdatePublicIpRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicIPAPI.UpdatePublicIp(context.Background(), publicIpId).XAuthToken(xAuthToken).UpdatePublicIpRequest(updatePublicIpRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicIPAPI.UpdatePublicIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePublicIp`: UpdatePublicIpResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicIPAPI.UpdatePublicIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicIpId** | **string** | 퍼블릭 IP의 고유 ID &lt;br/&gt;- [List public IPs](https://docs.kakaocloud.com/openapi/networking/vpc/list-public-ips)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePublicIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **updatePublicIpRequest** | [**UpdatePublicIpRequest**](UpdatePublicIpRequest.md) |  | 

### Return type

[**UpdatePublicIpResponse**](UpdatePublicIpResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

