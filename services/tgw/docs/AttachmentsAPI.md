# \AttachmentsAPI

All URIs are relative to *https://tgw.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApproveTgwAttachment**](AttachmentsAPI.md#ApproveTgwAttachment) | **Put** /api/v1/attachments/{attachment_id}/approve | Approve TGW attachment
[**CreateTgwAttachment**](AttachmentsAPI.md#CreateTgwAttachment) | **Post** /api/v1/attachments | Create TGW attachment
[**DeleteTgwAttachment**](AttachmentsAPI.md#DeleteTgwAttachment) | **Delete** /api/v1/attachments/{attachment_id} | Delete TGW attachment
[**GetTgwAttachment**](AttachmentsAPI.md#GetTgwAttachment) | **Get** /api/v1/attachments/{attachment_id} | Get TGW attachment
[**ListTgwAttachments**](AttachmentsAPI.md#ListTgwAttachments) | **Get** /api/v1/attachments | List TGW attachments
[**UpdateTgwAttachment**](AttachmentsAPI.md#UpdateTgwAttachment) | **Put** /api/v1/attachments/{attachment_id} | Update TGW attachment



## ApproveTgwAttachment

> BnsTgwV1ApiApproveTgwAttachmentModelCreateTgwAttachmentResponseModel ApproveTgwAttachment(ctx, attachmentId).XAuthToken(xAuthToken).Execute()

Approve TGW attachment



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
	attachmentId := "attachmentId_example" // string | 승인할 Transit Gateway Attachment ID 
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttachmentsAPI.ApproveTgwAttachment(context.Background(), attachmentId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttachmentsAPI.ApproveTgwAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApproveTgwAttachment`: BnsTgwV1ApiApproveTgwAttachmentModelCreateTgwAttachmentResponseModel
	fmt.Fprintf(os.Stdout, "Response from `AttachmentsAPI.ApproveTgwAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attachmentId** | **string** | 승인할 Transit Gateway Attachment ID  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApproveTgwAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**BnsTgwV1ApiApproveTgwAttachmentModelCreateTgwAttachmentResponseModel**](BnsTgwV1ApiApproveTgwAttachmentModelCreateTgwAttachmentResponseModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTgwAttachment

> BnsTgwV1ApiCreateTgwAttachmentModelCreateTgwAttachmentResponseModel CreateTgwAttachment(ctx).XAuthToken(xAuthToken).CreateTgwAttachmentRequestModel(createTgwAttachmentRequestModel).Execute()

Create TGW attachment



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
	createTgwAttachmentRequestModel := *openapiclient.NewCreateTgwAttachmentRequestModel(*openapiclient.NewBnsTgwV1ApiCreateTgwAttachmentModelTgwAttachmentRequestModel("TgwId_example", "VpcId_example", []string{"SubnetIds_example"})) // CreateTgwAttachmentRequestModel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttachmentsAPI.CreateTgwAttachment(context.Background()).XAuthToken(xAuthToken).CreateTgwAttachmentRequestModel(createTgwAttachmentRequestModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttachmentsAPI.CreateTgwAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTgwAttachment`: BnsTgwV1ApiCreateTgwAttachmentModelCreateTgwAttachmentResponseModel
	fmt.Fprintf(os.Stdout, "Response from `AttachmentsAPI.CreateTgwAttachment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTgwAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **createTgwAttachmentRequestModel** | [**CreateTgwAttachmentRequestModel**](CreateTgwAttachmentRequestModel.md) |  | 

### Return type

[**BnsTgwV1ApiCreateTgwAttachmentModelCreateTgwAttachmentResponseModel**](BnsTgwV1ApiCreateTgwAttachmentModelCreateTgwAttachmentResponseModel.md)

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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	attachmentId := "attachmentId_example" // string | 삭제할 Transit Gateway Attachment ID 
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AttachmentsAPI.DeleteTgwAttachment(context.Background(), attachmentId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttachmentsAPI.DeleteTgwAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attachmentId** | **string** | 삭제할 Transit Gateway Attachment ID  | 

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

> GetTgwAttachmentResponseModel GetTgwAttachment(ctx, attachmentId).XAuthToken(xAuthToken).Execute()

Get TGW attachment



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
	attachmentId := "attachmentId_example" // string | 조회할 Transit Gateway Attachment ID 
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttachmentsAPI.GetTgwAttachment(context.Background(), attachmentId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttachmentsAPI.GetTgwAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTgwAttachment`: GetTgwAttachmentResponseModel
	fmt.Fprintf(os.Stdout, "Response from `AttachmentsAPI.GetTgwAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attachmentId** | **string** | 조회할 Transit Gateway Attachment ID  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTgwAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**GetTgwAttachmentResponseModel**](GetTgwAttachmentResponseModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTgwAttachments

> GetTgwAttachmentsResponseModel ListTgwAttachments(ctx).XAuthToken(xAuthToken).Id(id).TgwId(tgwId).TgwName(tgwName).ProvisioningStatus(provisioningStatus).ResourceId(resourceId).ResourceName(resourceName).RouteTableId(routeTableId).RouteTableName(routeTableName).CreatedAt(createdAt).UpdatedAt(updatedAt).Offset(offset).Limit(limit).SortKeys(sortKeys).SortDirs(sortDirs).Execute()

List TGW attachments



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
	id := "id_example" // string | TGW Attachment ID  (optional)
	tgwId := "tgwId_example" // string | Transit Gateway ID  (optional)
	tgwName := "tgwName_example" // string | Transit Gateway 이름  (optional)
	provisioningStatus := openapiclient.ProvisioningStatus("ACTIVE") // ProvisioningStatus | 프로비저닝 상태 (optional)
	resourceId := "resourceId_example" // string | 연결된 리소스 ID (optional)
	resourceName := "resourceName_example" // string | 연결된 리소스 이름 (optional)
	routeTableId := "routeTableId_example" // string | 연결된 라우팅 테이블 ID  (optional)
	routeTableName := "routeTableName_example" // string | 연결된 라우팅 테이블 이름 (optional)
	createdAt := "createdAt_example" // string | 리소스가 생성된 시간 <br/> - ISO_8601 형식  <br/> - UTC 기준 (optional)
	updatedAt := "updatedAt_example" // string | 리소스가 마지막으로 수정된 시간 <br/> - ISO_8601 형식  <br/> - UTC 기준 (optional)
	offset := int32(56) // int32 | 조회 시작 위치 (optional) (default to 0)
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional) (default to 20)
	sortKeys := "sortKeys_example" // string | 정렬할 필드 (optional) (default to "created_at")
	sortDirs := "sortDirs_example" // string | 정렬 방향 (`asc`, `desc`)   (optional) (default to "desc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttachmentsAPI.ListTgwAttachments(context.Background()).XAuthToken(xAuthToken).Id(id).TgwId(tgwId).TgwName(tgwName).ProvisioningStatus(provisioningStatus).ResourceId(resourceId).ResourceName(resourceName).RouteTableId(routeTableId).RouteTableName(routeTableName).CreatedAt(createdAt).UpdatedAt(updatedAt).Offset(offset).Limit(limit).SortKeys(sortKeys).SortDirs(sortDirs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttachmentsAPI.ListTgwAttachments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTgwAttachments`: GetTgwAttachmentsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `AttachmentsAPI.ListTgwAttachments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTgwAttachmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **id** | **string** | TGW Attachment ID  | 
 **tgwId** | **string** | Transit Gateway ID  | 
 **tgwName** | **string** | Transit Gateway 이름  | 
 **provisioningStatus** | [**ProvisioningStatus**](ProvisioningStatus.md) | 프로비저닝 상태 | 
 **resourceId** | **string** | 연결된 리소스 ID | 
 **resourceName** | **string** | 연결된 리소스 이름 | 
 **routeTableId** | **string** | 연결된 라우팅 테이블 ID  | 
 **routeTableName** | **string** | 연결된 라우팅 테이블 이름 | 
 **createdAt** | **string** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
 **updatedAt** | **string** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
 **offset** | **int32** | 조회 시작 위치 | [default to 0]
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | [default to 20]
 **sortKeys** | **string** | 정렬할 필드 | [default to &quot;created_at&quot;]
 **sortDirs** | **string** | 정렬 방향 (&#x60;asc&#x60;, &#x60;desc&#x60;)   | [default to &quot;desc&quot;]

### Return type

[**GetTgwAttachmentsResponseModel**](GetTgwAttachmentsResponseModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTgwAttachment

> BnsTgwV1ApiUpdateTgwAttachmentModelCreateTgwAttachmentResponseModel UpdateTgwAttachment(ctx, attachmentId).XAuthToken(xAuthToken).UpdateTgwAttachmentRequestModel(updateTgwAttachmentRequestModel).Execute()

Update TGW attachment



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
	attachmentId := "attachmentId_example" // string | 수정할 Transit Gateway Attachment ID  
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	updateTgwAttachmentRequestModel := *openapiclient.NewUpdateTgwAttachmentRequestModel(*openapiclient.NewBnsTgwV1ApiUpdateTgwAttachmentModelTgwAttachmentRequestModel([]string{"SubnetIds_example"})) // UpdateTgwAttachmentRequestModel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttachmentsAPI.UpdateTgwAttachment(context.Background(), attachmentId).XAuthToken(xAuthToken).UpdateTgwAttachmentRequestModel(updateTgwAttachmentRequestModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttachmentsAPI.UpdateTgwAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTgwAttachment`: BnsTgwV1ApiUpdateTgwAttachmentModelCreateTgwAttachmentResponseModel
	fmt.Fprintf(os.Stdout, "Response from `AttachmentsAPI.UpdateTgwAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attachmentId** | **string** | 수정할 Transit Gateway Attachment ID   | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTgwAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **updateTgwAttachmentRequestModel** | [**UpdateTgwAttachmentRequestModel**](UpdateTgwAttachmentRequestModel.md) |  | 

### Return type

[**BnsTgwV1ApiUpdateTgwAttachmentModelCreateTgwAttachmentResponseModel**](BnsTgwV1ApiUpdateTgwAttachmentModelCreateTgwAttachmentResponseModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

