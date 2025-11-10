# \TgwsAPI

All URIs are relative to *https://tgw.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTransitGateway**](TgwsAPI.md#CreateTransitGateway) | **Post** /api/v1/tgws | Create transit gateway
[**DeleteTransitGateway**](TgwsAPI.md#DeleteTransitGateway) | **Delete** /api/v1/tgws/{tgw_id} | Delete transit gateway
[**GetTransitGateway**](TgwsAPI.md#GetTransitGateway) | **Get** /api/v1/tgws/{tgw_id} | Get transit gateway
[**ListTgwSharedProjects**](TgwsAPI.md#ListTgwSharedProjects) | **Get** /api/v1/tgws/{tgw_id}/projects | List TGW shared projects
[**ListTransitGateways**](TgwsAPI.md#ListTransitGateways) | **Get** /api/v1/tgws | List transit gateways
[**ShareTransitGateway**](TgwsAPI.md#ShareTransitGateway) | **Post** /api/v1/tgws/{tgw_id}/projects/{target_project_id} | Share transit gateway
[**UnshareTransitGateway**](TgwsAPI.md#UnshareTransitGateway) | **Delete** /api/v1/tgws/{tgw_id}/projects/{target_project_id} | Unshare transit gateway
[**UpdateTransitGateway**](TgwsAPI.md#UpdateTransitGateway) | **Put** /api/v1/tgws/{tgw_id} | Update transit gateway



## CreateTransitGateway

> BnsTgwV1ApiCreateTransitGatewayModelCreateTgwResponseModel CreateTransitGateway(ctx).XAuthToken(xAuthToken).CreateTgwRequestModel(createTgwRequestModel).Execute()

Create transit gateway



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
	createTgwRequestModel := *openapiclient.NewCreateTgwRequestModel(*openapiclient.NewBnsTgwV1ApiCreateTransitGatewayModelTgwRequestModel("Name_example", *openapiclient.NewBnsTgwV1ApiCreateTransitGatewayModelTgwOptionRequestModel(false, false))) // CreateTgwRequestModel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TgwsAPI.CreateTransitGateway(context.Background()).XAuthToken(xAuthToken).CreateTgwRequestModel(createTgwRequestModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TgwsAPI.CreateTransitGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTransitGateway`: BnsTgwV1ApiCreateTransitGatewayModelCreateTgwResponseModel
	fmt.Fprintf(os.Stdout, "Response from `TgwsAPI.CreateTransitGateway`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTransitGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **createTgwRequestModel** | [**CreateTgwRequestModel**](CreateTgwRequestModel.md) |  | 

### Return type

[**BnsTgwV1ApiCreateTransitGatewayModelCreateTgwResponseModel**](BnsTgwV1ApiCreateTransitGatewayModelCreateTgwResponseModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTransitGateway

> DeleteTransitGateway(ctx, tgwId).XAuthToken(xAuthToken).Execute()

Delete transit gateway



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
	tgwId := "tgwId_example" // string | 삭제할 Transit Gateway ID 
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TgwsAPI.DeleteTransitGateway(context.Background(), tgwId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TgwsAPI.DeleteTransitGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tgwId** | **string** | 삭제할 Transit Gateway ID  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTransitGatewayRequest struct via the builder pattern


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


## GetTransitGateway

> GetTgwResponseModel GetTransitGateway(ctx, tgwId).XAuthToken(xAuthToken).Execute()

Get transit gateway



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
	tgwId := "tgwId_example" // string | 조회할 Transit Gateway ID 
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TgwsAPI.GetTransitGateway(context.Background(), tgwId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TgwsAPI.GetTransitGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransitGateway`: GetTgwResponseModel
	fmt.Fprintf(os.Stdout, "Response from `TgwsAPI.GetTransitGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tgwId** | **string** | 조회할 Transit Gateway ID  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransitGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**GetTgwResponseModel**](GetTgwResponseModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTgwSharedProjects

> GetTgwProjectsResponseModel ListTgwSharedProjects(ctx, tgwId).XAuthToken(xAuthToken).Execute()

List TGW shared projects



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
	tgwId := "tgwId_example" // string | 조회할 Transit Gateway ID 
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TgwsAPI.ListTgwSharedProjects(context.Background(), tgwId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TgwsAPI.ListTgwSharedProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTgwSharedProjects`: GetTgwProjectsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `TgwsAPI.ListTgwSharedProjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tgwId** | **string** | 조회할 Transit Gateway ID  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTgwSharedProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**GetTgwProjectsResponseModel**](GetTgwProjectsResponseModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransitGateways

> GetTgwsResponseModel ListTransitGateways(ctx).XAuthToken(xAuthToken).Id(id).Name(name).Region(region).IsShared(isShared).ProvisioningStatus(provisioningStatus).CreatedAt(createdAt).UpdatedAt(updatedAt).Offset(offset).Limit(limit).SortKeys(sortKeys).SortDirs(sortDirs).Execute()

List transit gateways



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
	id := "id_example" // string | Transit Gateway ID  (optional)
	name := "name_example" // string | Transit Gateway 이름  (optional)
	region := openapiclient.Region("kr-central-2") // Region | Transit Gateway가 위치한 리전 (예: `kr-central-2`)  (optional)
	isShared := true // bool | 공유된 Transit Gateway 여부 (예: `true`)  (optional)
	provisioningStatus := openapiclient.ProvisioningStatus("ACTIVE") // ProvisioningStatus | Transit Gateway 프로비저닝 상태  (optional)
	createdAt := "createdAt_example" // string | 리소스가 생성된 시간 <br/> - ISO_8601 형식  <br/> - UTC 기준 (optional)
	updatedAt := "updatedAt_example" // string | 리소스가 마지막으로 수정된 시간 <br/> - ISO_8601 형식  <br/> - UTC 기준 (optional)
	offset := int32(56) // int32 | 조회 시작 위치 (optional) (default to 0)
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional) (default to 20)
	sortKeys := "sortKeys_example" // string | 정렬할 필드 (optional) (default to "created_at")
	sortDirs := "sortDirs_example" // string | 정렬 방향 (`asc`, `desc`)  (optional) (default to "desc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TgwsAPI.ListTransitGateways(context.Background()).XAuthToken(xAuthToken).Id(id).Name(name).Region(region).IsShared(isShared).ProvisioningStatus(provisioningStatus).CreatedAt(createdAt).UpdatedAt(updatedAt).Offset(offset).Limit(limit).SortKeys(sortKeys).SortDirs(sortDirs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TgwsAPI.ListTransitGateways``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTransitGateways`: GetTgwsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `TgwsAPI.ListTransitGateways`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTransitGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **id** | **string** | Transit Gateway ID  | 
 **name** | **string** | Transit Gateway 이름  | 
 **region** | [**Region**](Region.md) | Transit Gateway가 위치한 리전 (예: &#x60;kr-central-2&#x60;)  | 
 **isShared** | **bool** | 공유된 Transit Gateway 여부 (예: &#x60;true&#x60;)  | 
 **provisioningStatus** | [**ProvisioningStatus**](ProvisioningStatus.md) | Transit Gateway 프로비저닝 상태  | 
 **createdAt** | **string** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
 **updatedAt** | **string** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
 **offset** | **int32** | 조회 시작 위치 | [default to 0]
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | [default to 20]
 **sortKeys** | **string** | 정렬할 필드 | [default to &quot;created_at&quot;]
 **sortDirs** | **string** | 정렬 방향 (&#x60;asc&#x60;, &#x60;desc&#x60;)  | [default to &quot;desc&quot;]

### Return type

[**GetTgwsResponseModel**](GetTgwsResponseModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShareTransitGateway

> interface{} ShareTransitGateway(ctx, tgwId, targetProjectId).XAuthToken(xAuthToken).Execute()

Share transit gateway



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
	tgwId := "tgwId_example" // string | 공유할 Transit Gateway ID 
	targetProjectId := "targetProjectId_example" // string | 공유 대상 프로젝트 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TgwsAPI.ShareTransitGateway(context.Background(), tgwId, targetProjectId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TgwsAPI.ShareTransitGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShareTransitGateway`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TgwsAPI.ShareTransitGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tgwId** | **string** | 공유할 Transit Gateway ID  | 
**targetProjectId** | **string** | 공유 대상 프로젝트 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShareTransitGatewayRequest struct via the builder pattern


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


## UnshareTransitGateway

> UnshareTransitGateway(ctx, tgwId, targetProjectId).XAuthToken(xAuthToken).Execute()

Unshare transit gateway



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
	tgwId := "tgwId_example" // string | 공유 해제할 Transit Gateway ID 
	targetProjectId := "targetProjectId_example" // string | 공유 해제 대상 프로젝트 ID 
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TgwsAPI.UnshareTransitGateway(context.Background(), tgwId, targetProjectId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TgwsAPI.UnshareTransitGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tgwId** | **string** | 공유 해제할 Transit Gateway ID  | 
**targetProjectId** | **string** | 공유 해제 대상 프로젝트 ID  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnshareTransitGatewayRequest struct via the builder pattern


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


## UpdateTransitGateway

> BnsTgwV1ApiUpdateTransitGatewayModelCreateTgwResponseModel UpdateTransitGateway(ctx, tgwId).XAuthToken(xAuthToken).UpdateTgwRequestModel(updateTgwRequestModel).Execute()

Update transit gateway



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
	tgwId := "tgwId_example" // string | 수정할 Transit Gateway ID  
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	updateTgwRequestModel := *openapiclient.NewUpdateTgwRequestModel(*openapiclient.NewBnsTgwV1ApiUpdateTransitGatewayModelTgwRequestModel()) // UpdateTgwRequestModel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TgwsAPI.UpdateTransitGateway(context.Background(), tgwId).XAuthToken(xAuthToken).UpdateTgwRequestModel(updateTgwRequestModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TgwsAPI.UpdateTransitGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTransitGateway`: BnsTgwV1ApiUpdateTransitGatewayModelCreateTgwResponseModel
	fmt.Fprintf(os.Stdout, "Response from `TgwsAPI.UpdateTransitGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tgwId** | **string** | 수정할 Transit Gateway ID   | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTransitGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **updateTgwRequestModel** | [**UpdateTgwRequestModel**](UpdateTgwRequestModel.md) |  | 

### Return type

[**BnsTgwV1ApiUpdateTransitGatewayModelCreateTgwResponseModel**](BnsTgwV1ApiUpdateTransitGatewayModelCreateTgwResponseModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

