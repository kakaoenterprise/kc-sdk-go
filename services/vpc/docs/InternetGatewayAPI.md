# \InternetGatewayAPI

All URIs are relative to *https://vpc.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIgwAttachment**](InternetGatewayAPI.md#CreateIgwAttachment) | **Post** /api/v1/igws/{igw_id}/attachments | Create IGW attachment 
[**CreateInternetGateway**](InternetGatewayAPI.md#CreateInternetGateway) | **Post** /api/v1/igws | Create internet gateway 
[**DeleteIgwAttachment**](InternetGatewayAPI.md#DeleteIgwAttachment) | **Delete** /api/v1/igws/{igw_id}/attachments/{igw_attachment_id} | Delete IGW attachment 
[**DeleteInternetGateway**](InternetGatewayAPI.md#DeleteInternetGateway) | **Delete** /api/v1/igws/{igw_id} | Delete internet gateway 
[**GetInternetGateway**](InternetGatewayAPI.md#GetInternetGateway) | **Get** /api/v1/igws/{igw_id} | Get internet gateway 
[**ListInternetGateways**](InternetGatewayAPI.md#ListInternetGateways) | **Get** /api/v1/igws | List internet gateways 
[**UpdateInternetGateway**](InternetGatewayAPI.md#UpdateInternetGateway) | **Put** /api/v1/igws/{igw_id} | Update internet gateway 



## CreateIgwAttachment

> CreateIgwAttachmentResponse CreateIgwAttachment(ctx, igwId).XAuthToken(xAuthToken).CreateIgwAttachmentRequest(createIgwAttachmentRequest).Execute()

Create IGW attachment 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/vpc"
)

func main() {
	igwId := "igwId_example" // string | 연결할 인터넷 게이트웨이의 고유 ID <br/>- [List internet gateways](https://docs.kakaocloud.com/openapi/networking/vpc/list-internet-gateways)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	createIgwAttachmentRequest := *openapiclient.NewCreateIgwAttachmentRequest(*openapiclient.NewCreateIgwAttachment("VpcId_example")) // CreateIgwAttachmentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InternetGatewayAPI.CreateIgwAttachment(context.Background(), igwId).XAuthToken(xAuthToken).CreateIgwAttachmentRequest(createIgwAttachmentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternetGatewayAPI.CreateIgwAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIgwAttachment`: CreateIgwAttachmentResponse
	fmt.Fprintf(os.Stdout, "Response from `InternetGatewayAPI.CreateIgwAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**igwId** | **string** | 연결할 인터넷 게이트웨이의 고유 ID &lt;br/&gt;- [List internet gateways](https://docs.kakaocloud.com/openapi/networking/vpc/list-internet-gateways)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIgwAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **createIgwAttachmentRequest** | [**CreateIgwAttachmentRequest**](CreateIgwAttachmentRequest.md) |  | 

### Return type

[**CreateIgwAttachmentResponse**](CreateIgwAttachmentResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInternetGateway

> CreateInternetGatewayResponse CreateInternetGateway(ctx).XAuthToken(xAuthToken).CreateInternetGatewayRequest(createInternetGatewayRequest).Execute()

Create internet gateway 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/vpc"
)

func main() {
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	createInternetGatewayRequest := *openapiclient.NewCreateInternetGatewayRequest(*openapiclient.NewCreateInternetGateway()) // CreateInternetGatewayRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InternetGatewayAPI.CreateInternetGateway(context.Background()).XAuthToken(xAuthToken).CreateInternetGatewayRequest(createInternetGatewayRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternetGatewayAPI.CreateInternetGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInternetGateway`: CreateInternetGatewayResponse
	fmt.Fprintf(os.Stdout, "Response from `InternetGatewayAPI.CreateInternetGateway`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInternetGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **createInternetGatewayRequest** | [**CreateInternetGatewayRequest**](CreateInternetGatewayRequest.md) |  | 

### Return type

[**CreateInternetGatewayResponse**](CreateInternetGatewayResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIgwAttachment

> DeleteIgwAttachment(ctx, igwId, igwAttachmentId).XAuthToken(xAuthToken).Execute()

Delete IGW attachment 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/vpc"
)

func main() {
	igwId := "igwId_example" // string | 인터넷 게이트웨이의 고유 ID
	igwAttachmentId := "igwAttachmentId_example" // string | 삭제할 IGW Attachment의 고유 ID <br/>- [List internet gateways](https://docs.kakaocloud.com/openapi/networking/vpc/list-internet-gateways)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InternetGatewayAPI.DeleteIgwAttachment(context.Background(), igwId, igwAttachmentId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternetGatewayAPI.DeleteIgwAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**igwId** | **string** | 인터넷 게이트웨이의 고유 ID | 
**igwAttachmentId** | **string** | 삭제할 IGW Attachment의 고유 ID &lt;br/&gt;- [List internet gateways](https://docs.kakaocloud.com/openapi/networking/vpc/list-internet-gateways)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIgwAttachmentRequest struct via the builder pattern


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


## DeleteInternetGateway

> DeleteInternetGateway(ctx, igwId).XAuthToken(xAuthToken).Execute()

Delete internet gateway 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/vpc"
)

func main() {
	igwId := "igwId_example" // string | 삭제할 인터넷 게이트웨이의 고유 ID <br/>- [List internet gateways](https://docs.kakaocloud.com/openapi/networking/vpc/list-internet-gateways)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InternetGatewayAPI.DeleteInternetGateway(context.Background(), igwId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternetGatewayAPI.DeleteInternetGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**igwId** | **string** | 삭제할 인터넷 게이트웨이의 고유 ID &lt;br/&gt;- [List internet gateways](https://docs.kakaocloud.com/openapi/networking/vpc/list-internet-gateways)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInternetGatewayRequest struct via the builder pattern


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


## GetInternetGateway

> GetInternetGatewayResponse GetInternetGateway(ctx, igwId).XAuthToken(xAuthToken).Execute()

Get internet gateway 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/vpc"
)

func main() {
	igwId := "igwId_example" // string | 조회할 인터넷 게이트웨이의 고유 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InternetGatewayAPI.GetInternetGateway(context.Background(), igwId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternetGatewayAPI.GetInternetGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInternetGateway`: GetInternetGatewayResponse
	fmt.Fprintf(os.Stdout, "Response from `InternetGatewayAPI.GetInternetGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**igwId** | **string** | 조회할 인터넷 게이트웨이의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInternetGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**GetInternetGatewayResponse**](GetInternetGatewayResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInternetGateways

> ListInternetGatewaysResponse ListInternetGateways(ctx).XAuthToken(xAuthToken).Id(id).Name(name).VpcId(vpcId).VpcName(vpcName).VpcCidrBlock(vpcCidrBlock).AttachmentId(attachmentId).AttachmentProvisioningStatus(attachmentProvisioningStatus).NatIp(natIp).ProvisioningStatus(provisioningStatus).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Offset(offset).Limit(limit).Execute()

List internet gateways 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/vpc"
)

func main() {
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	id := "id_example" // string | 인터넷 게이트웨이의 고유 ID (optional)
	name := "name_example" // string | 인터넷 게이트웨이 이름 (optional)
	vpcId := "vpcId_example" // string | 인터넷 게이트웨이가 연결된 VPC의 고유 ID (optional)
	vpcName := "vpcName_example" // string | 인터넷 게이트웨이가 연결된 VPC 이름 (optional)
	vpcCidrBlock := "vpcCidrBlock_example" // string | 인터넷 게이트웨이가 연결된 VPC의 CIDR 블록 (optional)
	attachmentId := "attachmentId_example" // string | IGW Attachment의 고유 ID (optional)
	attachmentProvisioningStatus := "attachmentProvisioningStatus_example" // string | IGW Attachment의 프로비저닝 상태 <br/>- 사용 가능한 값: `ACTIVE`, `DELETED`, `ERROR`, `PENDING_CREATE`, `PENDING_UPDATE`, `PENDING_DELETE` (optional)
	natIp := "natIp_example" // string | 인터넷 게이트웨이에 할당된 NAT IP 주소 (optional)
	provisioningStatus := "provisioningStatus_example" // string | 인터넷 게이트웨이의 프로비저닝 상태 <br/>- 사용 가능한 값: `ACTIVE`, `DELETED`, `ERROR`, `PENDING_CREATE`, `PENDING_UPDATE`, `PENDING_DELETE` (optional)
	createdAt := "createdAt_example" // string | 리소스가 생성된 시간 <br/> - ISO 8601 형식 <br/> - UTC 기준 (optional)
	updatedAt := "updatedAt_example" // string | 리소스가 마지막으로 수정된 시간 <br/> - ISO 8601 형식 <br/> - UTC 기준 (optional)
	sortKeys := "sortKeys_example" // string | 정렬할 필드 (optional)
	sortDirs := "sortDirs_example" // string | 정렬 방향 (`asc`, `desc`) (optional)
	offset := int32(56) // int32 | 조회 시작 위치 (optional)
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InternetGatewayAPI.ListInternetGateways(context.Background()).XAuthToken(xAuthToken).Id(id).Name(name).VpcId(vpcId).VpcName(vpcName).VpcCidrBlock(vpcCidrBlock).AttachmentId(attachmentId).AttachmentProvisioningStatus(attachmentProvisioningStatus).NatIp(natIp).ProvisioningStatus(provisioningStatus).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternetGatewayAPI.ListInternetGateways``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInternetGateways`: ListInternetGatewaysResponse
	fmt.Fprintf(os.Stdout, "Response from `InternetGatewayAPI.ListInternetGateways`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInternetGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **id** | **string** | 인터넷 게이트웨이의 고유 ID | 
 **name** | **string** | 인터넷 게이트웨이 이름 | 
 **vpcId** | **string** | 인터넷 게이트웨이가 연결된 VPC의 고유 ID | 
 **vpcName** | **string** | 인터넷 게이트웨이가 연결된 VPC 이름 | 
 **vpcCidrBlock** | **string** | 인터넷 게이트웨이가 연결된 VPC의 CIDR 블록 | 
 **attachmentId** | **string** | IGW Attachment의 고유 ID | 
 **attachmentProvisioningStatus** | **string** | IGW Attachment의 프로비저닝 상태 &lt;br/&gt;- 사용 가능한 값: &#x60;ACTIVE&#x60;, &#x60;DELETED&#x60;, &#x60;ERROR&#x60;, &#x60;PENDING_CREATE&#x60;, &#x60;PENDING_UPDATE&#x60;, &#x60;PENDING_DELETE&#x60; | 
 **natIp** | **string** | 인터넷 게이트웨이에 할당된 NAT IP 주소 | 
 **provisioningStatus** | **string** | 인터넷 게이트웨이의 프로비저닝 상태 &lt;br/&gt;- 사용 가능한 값: &#x60;ACTIVE&#x60;, &#x60;DELETED&#x60;, &#x60;ERROR&#x60;, &#x60;PENDING_CREATE&#x60;, &#x60;PENDING_UPDATE&#x60;, &#x60;PENDING_DELETE&#x60; | 
 **createdAt** | **string** | 리소스가 생성된 시간 &lt;br/&gt; - ISO 8601 형식 &lt;br/&gt; - UTC 기준 | 
 **updatedAt** | **string** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO 8601 형식 &lt;br/&gt; - UTC 기준 | 
 **sortKeys** | **string** | 정렬할 필드 | 
 **sortDirs** | **string** | 정렬 방향 (&#x60;asc&#x60;, &#x60;desc&#x60;) | 
 **offset** | **int32** | 조회 시작 위치 | 
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | 

### Return type

[**ListInternetGatewaysResponse**](ListInternetGatewaysResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInternetGateway

> UpdateInternetGatewayResponse UpdateInternetGateway(ctx, igwId).XAuthToken(xAuthToken).UpdateInternetGatewayRequest(updateInternetGatewayRequest).Execute()

Update internet gateway 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/vpc"
)

func main() {
	igwId := "igwId_example" // string | 수정할 인터넷 게이트웨이의 고유 ID <br/>- [List internet gateways](https://docs.kakaocloud.com/openapi/networking/vpc/list-internet-gateways)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	updateInternetGatewayRequest := *openapiclient.NewUpdateInternetGatewayRequest(*openapiclient.NewUpdateInternetGateway()) // UpdateInternetGatewayRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InternetGatewayAPI.UpdateInternetGateway(context.Background(), igwId).XAuthToken(xAuthToken).UpdateInternetGatewayRequest(updateInternetGatewayRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternetGatewayAPI.UpdateInternetGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateInternetGateway`: UpdateInternetGatewayResponse
	fmt.Fprintf(os.Stdout, "Response from `InternetGatewayAPI.UpdateInternetGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**igwId** | **string** | 수정할 인터넷 게이트웨이의 고유 ID &lt;br/&gt;- [List internet gateways](https://docs.kakaocloud.com/openapi/networking/vpc/list-internet-gateways)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInternetGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **updateInternetGatewayRequest** | [**UpdateInternetGatewayRequest**](UpdateInternetGatewayRequest.md) |  | 

### Return type

[**UpdateInternetGatewayResponse**](UpdateInternetGatewayResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

