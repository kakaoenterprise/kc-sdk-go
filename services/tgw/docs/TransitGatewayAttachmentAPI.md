# \TransitGatewayAttachmentAPI

All URIs are relative to *https://tgw.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApproveTgwAttachment**](TransitGatewayAttachmentAPI.md#ApproveTgwAttachment) | **Put** /api/v1/attachments/{attachment_id}/approve | Approve TGW attachment
[**CreateTgwAttachment**](TransitGatewayAttachmentAPI.md#CreateTgwAttachment) | **Post** /api/v1/attachments | Create TGW attachment
[**DeleteTgwAttachment**](TransitGatewayAttachmentAPI.md#DeleteTgwAttachment) | **Delete** /api/v1/attachments/{attachment_id} | Delete TGW attachment
[**GetTgwAttachment**](TransitGatewayAttachmentAPI.md#GetTgwAttachment) | **Get** /api/v1/attachments/{attachment_id} | Get TGW attachment
[**ListTgwAttachments**](TransitGatewayAttachmentAPI.md#ListTgwAttachments) | **Get** /api/v1/attachments | List TGW attachments
[**RejectTgwAttachment**](TransitGatewayAttachmentAPI.md#RejectTgwAttachment) | **Put** /api/v1/attachments/{attachment_id}/reject | Reject TGW attachment 
[**UpdateTgwAttachment**](TransitGatewayAttachmentAPI.md#UpdateTgwAttachment) | **Put** /api/v1/attachments/{attachment_id} | Update TGW attachment



## ApproveTgwAttachment

> ApproveTgwAttachmentResponse ApproveTgwAttachment(ctx, attachmentId).XAuthToken(xAuthToken).Execute()

Approve TGW attachment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/tgw"
)

func main() {
	attachmentId := "attachmentId_example" // string | 승인할 Transit Gateway Attachment ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransitGatewayAttachmentAPI.ApproveTgwAttachment(context.Background(), attachmentId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransitGatewayAttachmentAPI.ApproveTgwAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApproveTgwAttachment`: ApproveTgwAttachmentResponse
	fmt.Fprintf(os.Stdout, "Response from `TransitGatewayAttachmentAPI.ApproveTgwAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attachmentId** | **string** | 승인할 Transit Gateway Attachment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApproveTgwAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**ApproveTgwAttachmentResponse**](ApproveTgwAttachmentResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTgwAttachment

> CreateTgwAttachmentResponse CreateTgwAttachment(ctx).XAuthToken(xAuthToken).CreateTgwAttachmentRequest(createTgwAttachmentRequest).Execute()

Create TGW attachment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/tgw"
)

func main() {
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	createTgwAttachmentRequest := *openapiclient.NewCreateTgwAttachmentRequest(*openapiclient.NewCreateTgwAttachment("TgwId_example", "VpcId_example", []string{"SubnetIds_example"})) // CreateTgwAttachmentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransitGatewayAttachmentAPI.CreateTgwAttachment(context.Background()).XAuthToken(xAuthToken).CreateTgwAttachmentRequest(createTgwAttachmentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransitGatewayAttachmentAPI.CreateTgwAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTgwAttachment`: CreateTgwAttachmentResponse
	fmt.Fprintf(os.Stdout, "Response from `TransitGatewayAttachmentAPI.CreateTgwAttachment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTgwAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **createTgwAttachmentRequest** | [**CreateTgwAttachmentRequest**](CreateTgwAttachmentRequest.md) |  | 

### Return type

[**CreateTgwAttachmentResponse**](CreateTgwAttachmentResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTgwAttachment

> DeleteTgwAttachment(ctx, attachmentId).XAuthToken(xAuthToken).Execute()

Delete TGW attachment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/tgw"
)

func main() {
	attachmentId := "attachmentId_example" // string | 삭제할 Transit Gateway Attachment ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TransitGatewayAttachmentAPI.DeleteTgwAttachment(context.Background(), attachmentId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransitGatewayAttachmentAPI.DeleteTgwAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attachmentId** | **string** | 삭제할 Transit Gateway Attachment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTgwAttachmentRequest struct via the builder pattern


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


## GetTgwAttachment

> GetTgwAttachmentResponse GetTgwAttachment(ctx, attachmentId).XAuthToken(xAuthToken).Execute()

Get TGW attachment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/tgw"
)

func main() {
	attachmentId := "attachmentId_example" // string | 조회할 Transit Gateway Attachment ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransitGatewayAttachmentAPI.GetTgwAttachment(context.Background(), attachmentId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransitGatewayAttachmentAPI.GetTgwAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTgwAttachment`: GetTgwAttachmentResponse
	fmt.Fprintf(os.Stdout, "Response from `TransitGatewayAttachmentAPI.GetTgwAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attachmentId** | **string** | 조회할 Transit Gateway Attachment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTgwAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**GetTgwAttachmentResponse**](GetTgwAttachmentResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTgwAttachments

> ListTgwAttachmentsResponse ListTgwAttachments(ctx).XAuthToken(xAuthToken).Id(id).Name(name).TgwId(tgwId).TgwName(tgwName).ProvisioningStatus(provisioningStatus).ResourceId(resourceId).ResourceName(resourceName).RouteTableId(routeTableId).RouteTableName(routeTableName).CreatedAt(createdAt).UpdatedAt(updatedAt).Offset(offset).Limit(limit).SortKeys(sortKeys).SortDirs(sortDirs).Execute()

List TGW attachments



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/tgw"
)

func main() {
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	id := "id_example" // string | TGW Attachment ID (optional)
	name := "name_example" // string | Transit Gateway Attachment 이름 (optional)
	tgwId := "tgwId_example" // string | Transit Gateway ID (optional)
	tgwName := "tgwName_example" // string | Transit Gateway 이름 (optional)
	provisioningStatus := openapiclient.TGWAttachmentProvisioningStatus("ACTIVE") // TGWAttachmentProvisioningStatus | Transit Gateway Attachment의 프로비저닝 상태 (optional)
	resourceId := "resourceId_example" // string | 연결된 리소스 ID (optional)
	resourceName := "resourceName_example" // string | 연결된 리소스 이름 (optional)
	routeTableId := "routeTableId_example" // string | 연결된 라우팅 테이블 ID (optional)
	routeTableName := "routeTableName_example" // string | 연결된 라우팅 테이블 이름 (optional)
	createdAt := "createdAt_example" // string | 리소스가 생성된 시간 <br/> - ISO 8601 형식 <br/> - UTC 기준 (optional)
	updatedAt := "updatedAt_example" // string | 리소스가 마지막으로 수정된 시간 <br/> - ISO 8601 형식 <br/> - UTC 기준 (optional)
	offset := int32(56) // int32 | 조회 시작 위치 (optional)
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional)
	sortKeys := "sortKeys_example" // string | 정렬할 필드 (optional)
	sortDirs := "sortDirs_example" // string | 정렬 방향 (`asc`, `desc`)   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransitGatewayAttachmentAPI.ListTgwAttachments(context.Background()).XAuthToken(xAuthToken).Id(id).Name(name).TgwId(tgwId).TgwName(tgwName).ProvisioningStatus(provisioningStatus).ResourceId(resourceId).ResourceName(resourceName).RouteTableId(routeTableId).RouteTableName(routeTableName).CreatedAt(createdAt).UpdatedAt(updatedAt).Offset(offset).Limit(limit).SortKeys(sortKeys).SortDirs(sortDirs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransitGatewayAttachmentAPI.ListTgwAttachments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTgwAttachments`: ListTgwAttachmentsResponse
	fmt.Fprintf(os.Stdout, "Response from `TransitGatewayAttachmentAPI.ListTgwAttachments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTgwAttachmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **id** | **string** | TGW Attachment ID | 
 **name** | **string** | Transit Gateway Attachment 이름 | 
 **tgwId** | **string** | Transit Gateway ID | 
 **tgwName** | **string** | Transit Gateway 이름 | 
 **provisioningStatus** | [**TGWAttachmentProvisioningStatus**](TGWAttachmentProvisioningStatus.md) | Transit Gateway Attachment의 프로비저닝 상태 | 
 **resourceId** | **string** | 연결된 리소스 ID | 
 **resourceName** | **string** | 연결된 리소스 이름 | 
 **routeTableId** | **string** | 연결된 라우팅 테이블 ID | 
 **routeTableName** | **string** | 연결된 라우팅 테이블 이름 | 
 **createdAt** | **string** | 리소스가 생성된 시간 &lt;br/&gt; - ISO 8601 형식 &lt;br/&gt; - UTC 기준 | 
 **updatedAt** | **string** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO 8601 형식 &lt;br/&gt; - UTC 기준 | 
 **offset** | **int32** | 조회 시작 위치 | 
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | 
 **sortKeys** | **string** | 정렬할 필드 | 
 **sortDirs** | **string** | 정렬 방향 (&#x60;asc&#x60;, &#x60;desc&#x60;)   | 

### Return type

[**ListTgwAttachmentsResponse**](ListTgwAttachmentsResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RejectTgwAttachment

> RejectTgwAttachmentResponse RejectTgwAttachment(ctx, attachmentId).XAuthToken(xAuthToken).Execute()

Reject TGW attachment 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/tgw"
)

func main() {
	attachmentId := "attachmentId_example" // string | 거부할 Transit Gateway Attachment의 고유 ID <br/>- [List TGW attachments](https://docs.kakaocloud.com/openapi/networking/tgw/list-tgw-attachments)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransitGatewayAttachmentAPI.RejectTgwAttachment(context.Background(), attachmentId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransitGatewayAttachmentAPI.RejectTgwAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RejectTgwAttachment`: RejectTgwAttachmentResponse
	fmt.Fprintf(os.Stdout, "Response from `TransitGatewayAttachmentAPI.RejectTgwAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attachmentId** | **string** | 거부할 Transit Gateway Attachment의 고유 ID &lt;br/&gt;- [List TGW attachments](https://docs.kakaocloud.com/openapi/networking/tgw/list-tgw-attachments)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiRejectTgwAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**RejectTgwAttachmentResponse**](RejectTgwAttachmentResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTgwAttachment

> UpdateTgwAttachmentResponse UpdateTgwAttachment(ctx, attachmentId).XAuthToken(xAuthToken).UpdateTgwAttachmentRequest(updateTgwAttachmentRequest).Execute()

Update TGW attachment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/tgw"
)

func main() {
	attachmentId := "attachmentId_example" // string | 수정할 Transit Gateway Attachment ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	updateTgwAttachmentRequest := *openapiclient.NewUpdateTgwAttachmentRequest(*openapiclient.NewUpdateTgwAttachment([]string{"SubnetIds_example"})) // UpdateTgwAttachmentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransitGatewayAttachmentAPI.UpdateTgwAttachment(context.Background(), attachmentId).XAuthToken(xAuthToken).UpdateTgwAttachmentRequest(updateTgwAttachmentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransitGatewayAttachmentAPI.UpdateTgwAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTgwAttachment`: UpdateTgwAttachmentResponse
	fmt.Fprintf(os.Stdout, "Response from `TransitGatewayAttachmentAPI.UpdateTgwAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attachmentId** | **string** | 수정할 Transit Gateway Attachment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTgwAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **updateTgwAttachmentRequest** | [**UpdateTgwAttachmentRequest**](UpdateTgwAttachmentRequest.md) |  | 

### Return type

[**UpdateTgwAttachmentResponse**](UpdateTgwAttachmentResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

