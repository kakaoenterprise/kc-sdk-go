# \PublicIPAPI

All URIs are relative to *https://network.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssociatePublicIp**](PublicIPAPI.md#AssociatePublicIp) | **Put** /api/v1/public-ips/{public_ip_id}/network-interfaces/{network_interface_id} | Associate public IP
[**CreatePublicIp**](PublicIPAPI.md#CreatePublicIp) | **Post** /api/v1/public-ips | Create public IP
[**DeletePublicIp**](PublicIPAPI.md#DeletePublicIp) | **Delete** /api/v1/public-ips/{public_ip_id} | Delete public IP
[**DisassociatePublicIp**](PublicIPAPI.md#DisassociatePublicIp) | **Delete** /api/v1/public-ips/{public_ip_id}/network-interfaces/{network_interface_id} | Disassociate public IP
[**GetPublicIp**](PublicIPAPI.md#GetPublicIp) | **Get** /api/v1/public-ips/{public_ip_id} | Get public IP
[**ListPublicIps**](PublicIPAPI.md#ListPublicIps) | **Get** /api/v1/public-ips | List public IPs
[**UpdatePublicIp**](PublicIPAPI.md#UpdatePublicIp) | **Put** /api/v1/public-ips/{public_ip_id} | Update public IP



## AssociatePublicIp

> BnsNetworkV1ApiAssociatePublicIpModelResponsePublicIpModel AssociatePublicIp(ctx, publicIpId, networkInterfaceId).XAuthToken(xAuthToken).Execute()

Associate public IP



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
	publicIpId := "publicIpId_example" // string | 연결할 퍼블릭 IP의 ID 
	networkInterfaceId := "networkInterfaceId_example" // string | 퍼블릭 IP를 연결할 대상 네트워크 인터페이스의 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicIPAPI.AssociatePublicIp(context.Background(), publicIpId, networkInterfaceId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicIPAPI.AssociatePublicIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssociatePublicIp`: BnsNetworkV1ApiAssociatePublicIpModelResponsePublicIpModel
	fmt.Fprintf(os.Stdout, "Response from `PublicIPAPI.AssociatePublicIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicIpId** | **string** | 연결할 퍼블릭 IP의 ID  | 
**networkInterfaceId** | **string** | 퍼블릭 IP를 연결할 대상 네트워크 인터페이스의 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssociatePublicIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**BnsNetworkV1ApiAssociatePublicIpModelResponsePublicIpModel**](BnsNetworkV1ApiAssociatePublicIpModelResponsePublicIpModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePublicIp

> BnsNetworkV1ApiCreatePublicIpModelResponsePublicIpModel CreatePublicIp(ctx).XAuthToken(xAuthToken).BodyCreatePublicIp(bodyCreatePublicIp).Execute()

Create public IP



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
	bodyCreatePublicIp := *openapiclient.NewBodyCreatePublicIp(*openapiclient.NewCreatePublicIpModel()) // BodyCreatePublicIp | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicIPAPI.CreatePublicIp(context.Background()).XAuthToken(xAuthToken).BodyCreatePublicIp(bodyCreatePublicIp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicIPAPI.CreatePublicIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePublicIp`: BnsNetworkV1ApiCreatePublicIpModelResponsePublicIpModel
	fmt.Fprintf(os.Stdout, "Response from `PublicIPAPI.CreatePublicIp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePublicIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyCreatePublicIp** | [**BodyCreatePublicIp**](BodyCreatePublicIp.md) |  | 

### Return type

[**BnsNetworkV1ApiCreatePublicIpModelResponsePublicIpModel**](BnsNetworkV1ApiCreatePublicIpModelResponsePublicIpModel.md)

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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	publicIpId := "publicIpId_example" // string | 삭제할 퍼블릭 IP ID
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
**publicIpId** | **string** | 삭제할 퍼블릭 IP ID | 

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


## DisassociatePublicIp

> BnsNetworkV1ApiDisassociatePublicIpModelResponsePublicIpModel DisassociatePublicIp(ctx, publicIpId, networkInterfaceId).XAuthToken(xAuthToken).Execute()

Disassociate public IP



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
	publicIpId := "publicIpId_example" // string | 연결 해제할 퍼블릭 IP의 ID 
	networkInterfaceId := "networkInterfaceId_example" // string | 연결 해제 대상 네트워크 인터페이스의 ID  
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicIPAPI.DisassociatePublicIp(context.Background(), publicIpId, networkInterfaceId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicIPAPI.DisassociatePublicIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisassociatePublicIp`: BnsNetworkV1ApiDisassociatePublicIpModelResponsePublicIpModel
	fmt.Fprintf(os.Stdout, "Response from `PublicIPAPI.DisassociatePublicIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicIpId** | **string** | 연결 해제할 퍼블릭 IP의 ID  | 
**networkInterfaceId** | **string** | 연결 해제 대상 네트워크 인터페이스의 ID   | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisassociatePublicIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**BnsNetworkV1ApiDisassociatePublicIpModelResponsePublicIpModel**](BnsNetworkV1ApiDisassociatePublicIpModelResponsePublicIpModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicIp

> BnsNetworkV1ApiGetPublicIpModelResponsePublicIpModel GetPublicIp(ctx, publicIpId).XAuthToken(xAuthToken).Execute()

Get public IP



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
	publicIpId := "publicIpId_example" // string | 조회할 퍼블릭 IP ID 
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicIPAPI.GetPublicIp(context.Background(), publicIpId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicIPAPI.GetPublicIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicIp`: BnsNetworkV1ApiGetPublicIpModelResponsePublicIpModel
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

[**BnsNetworkV1ApiGetPublicIpModelResponsePublicIpModel**](BnsNetworkV1ApiGetPublicIpModelResponsePublicIpModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPublicIps

> PublicIpListModel ListPublicIps(ctx).XAuthToken(xAuthToken).Id(id).Status(status).PublicIp(publicIp).RelatedResourceId(relatedResourceId).RelatedResourceName(relatedResourceName).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Limit(limit).Offset(offset).Execute()

List public IPs



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
	id := "id_example" // string | 퍼블릭 IP의 고유 ID (optional)
	status := openapiclient.PublicIpStatus("available") // PublicIpStatus | 상태 <br/> - `available`: 사용 가능 <br/> - `in_use`: 사용 중 <br/> - `attaching`: 연결 중  (optional)
	publicIp := "publicIp_example" // string | 퍼블릭 IP 주소  (optional)
	relatedResourceId := "relatedResourceId_example" // string | 퍼블릭 IP가 연결된 리소스 ID (예: 네트워크 인터페이스 ID 등)  (optional)
	relatedResourceName := "relatedResourceName_example" // string | 퍼블릭 IP가 연결된 리소스 이름 (예: 네트워크 인터페이스 이름 등)  (optional)
	createdAt := "createdAt_example" // string | 리소스가 생성된 시간 <br/> - ISO_8601 형식  <br/> - UTC 기준 (optional)
	updatedAt := "updatedAt_example" // string | 리소스가 마지막으로 수정된 시간 <br/> - ISO_8601 형식  <br/> - UTC 기준 (optional)
	sortKeys := "sortKeys_example" // string | 정렬할 필드를 콤마(,)로 구분 (optional) (default to "created_at")
	sortDirs := "sortDirs_example" // string | 정렬 방향 (`asc`, `desc`)  (optional) (default to "desc")
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional) (default to 20)
	offset := int32(56) // int32 | 조회 시작 위치 (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicIPAPI.ListPublicIps(context.Background()).XAuthToken(xAuthToken).Id(id).Status(status).PublicIp(publicIp).RelatedResourceId(relatedResourceId).RelatedResourceName(relatedResourceName).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicIPAPI.ListPublicIps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPublicIps`: PublicIpListModel
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
 **status** | [**PublicIpStatus**](PublicIpStatus.md) | 상태 &lt;br/&gt; - &#x60;available&#x60;: 사용 가능 &lt;br/&gt; - &#x60;in_use&#x60;: 사용 중 &lt;br/&gt; - &#x60;attaching&#x60;: 연결 중  | 
 **publicIp** | **string** | 퍼블릭 IP 주소  | 
 **relatedResourceId** | **string** | 퍼블릭 IP가 연결된 리소스 ID (예: 네트워크 인터페이스 ID 등)  | 
 **relatedResourceName** | **string** | 퍼블릭 IP가 연결된 리소스 이름 (예: 네트워크 인터페이스 이름 등)  | 
 **createdAt** | **string** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
 **updatedAt** | **string** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
 **sortKeys** | **string** | 정렬할 필드를 콤마(,)로 구분 | [default to &quot;created_at&quot;]
 **sortDirs** | **string** | 정렬 방향 (&#x60;asc&#x60;, &#x60;desc&#x60;)  | [default to &quot;desc&quot;]
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | [default to 20]
 **offset** | **int32** | 조회 시작 위치 | [default to 0]

### Return type

[**PublicIpListModel**](PublicIpListModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePublicIp

> BnsNetworkV1ApiUpdatePublicIpModelResponsePublicIpModel UpdatePublicIp(ctx, publicIpId).XAuthToken(xAuthToken).BodyUpdatePublicIp(bodyUpdatePublicIp).Execute()

Update public IP



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
	publicIpId := "publicIpId_example" // string | 퍼블릭 IP의 고유 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	bodyUpdatePublicIp := *openapiclient.NewBodyUpdatePublicIp(*openapiclient.NewEditPublicIpModel()) // BodyUpdatePublicIp | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicIPAPI.UpdatePublicIp(context.Background(), publicIpId).XAuthToken(xAuthToken).BodyUpdatePublicIp(bodyUpdatePublicIp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicIPAPI.UpdatePublicIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePublicIp`: BnsNetworkV1ApiUpdatePublicIpModelResponsePublicIpModel
	fmt.Fprintf(os.Stdout, "Response from `PublicIPAPI.UpdatePublicIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicIpId** | **string** | 퍼블릭 IP의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePublicIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyUpdatePublicIp** | [**BodyUpdatePublicIp**](BodyUpdatePublicIp.md) |  | 

### Return type

[**BnsNetworkV1ApiUpdatePublicIpModelResponsePublicIpModel**](BnsNetworkV1ApiUpdatePublicIpModelResponsePublicIpModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

