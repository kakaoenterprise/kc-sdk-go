# \TransitGatewayAPI

All URIs are relative to *https://tgw.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTransitGateway**](TransitGatewayAPI.md#CreateTransitGateway) | **Post** /api/v1/tgws | Create transit gateway
[**DeleteTransitGateway**](TransitGatewayAPI.md#DeleteTransitGateway) | **Delete** /api/v1/tgws/{tgw_id} | Delete transit gateway
[**GetTransitGateway**](TransitGatewayAPI.md#GetTransitGateway) | **Get** /api/v1/tgws/{tgw_id} | Get transit gateway
[**ListTgwSharedProjects**](TransitGatewayAPI.md#ListTgwSharedProjects) | **Get** /api/v1/tgws/{tgw_id}/projects | List TGW shared projects
[**ListTransitGateways**](TransitGatewayAPI.md#ListTransitGateways) | **Get** /api/v1/tgws | List transit gateways
[**ShareTransitGateway**](TransitGatewayAPI.md#ShareTransitGateway) | **Post** /api/v1/tgws/{tgw_id}/projects/{target_project_id} | Share transit gateway
[**UnshareTransitGateway**](TransitGatewayAPI.md#UnshareTransitGateway) | **Delete** /api/v1/tgws/{tgw_id}/projects/{target_project_id} | Unshare transit gateway
[**UpdateTransitGateway**](TransitGatewayAPI.md#UpdateTransitGateway) | **Put** /api/v1/tgws/{tgw_id} | Update transit gateway



## CreateTransitGateway

> CreateTransitGatewayResponse CreateTransitGateway(ctx).CreateTransitGatewayRequest(createTransitGatewayRequest).Execute()

Create transit gateway



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/tgw"
)

func main() {
	createTransitGatewayRequest := *openapiclient.NewCreateTransitGatewayRequest(*openapiclient.NewCreateTransitGateway("Name_example", *openapiclient.NewCreateTransitGatewayTgwOptionRequest(false, false))) // CreateTransitGatewayRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransitGatewayAPI.CreateTransitGateway(context.Background()).CreateTransitGatewayRequest(createTransitGatewayRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransitGatewayAPI.CreateTransitGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTransitGateway`: CreateTransitGatewayResponse
	fmt.Fprintf(os.Stdout, "Response from `TransitGatewayAPI.CreateTransitGateway`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTransitGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTransitGatewayRequest** | [**CreateTransitGatewayRequest**](CreateTransitGatewayRequest.md) |  | 

### Return type

[**CreateTransitGatewayResponse**](CreateTransitGatewayResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTransitGateway

> DeleteTransitGateway(ctx, tgwId).Execute()

Delete transit gateway



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/tgw"
)

func main() {
	tgwId := "tgwId_example" // string | 삭제할 Transit Gateway ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TransitGatewayAPI.DeleteTransitGateway(context.Background(), tgwId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransitGatewayAPI.DeleteTransitGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tgwId** | **string** | 삭제할 Transit Gateway ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTransitGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> GetTransitGatewayResponse GetTransitGateway(ctx, tgwId).Execute()

Get transit gateway



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/tgw"
)

func main() {
	tgwId := "tgwId_example" // string | 조회할 Transit Gateway ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransitGatewayAPI.GetTransitGateway(context.Background(), tgwId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransitGatewayAPI.GetTransitGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransitGateway`: GetTransitGatewayResponse
	fmt.Fprintf(os.Stdout, "Response from `TransitGatewayAPI.GetTransitGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tgwId** | **string** | 조회할 Transit Gateway ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransitGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetTransitGatewayResponse**](GetTransitGatewayResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTgwSharedProjects

> ListTgwSharedProjectsResponse ListTgwSharedProjects(ctx, tgwId).Execute()

List TGW shared projects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/tgw"
)

func main() {
	tgwId := "tgwId_example" // string | 조회할 Transit Gateway ID - [List transit gateways](/openapi/networking/tgw/list-transit-gateways)에서 확인

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransitGatewayAPI.ListTgwSharedProjects(context.Background(), tgwId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransitGatewayAPI.ListTgwSharedProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTgwSharedProjects`: ListTgwSharedProjectsResponse
	fmt.Fprintf(os.Stdout, "Response from `TransitGatewayAPI.ListTgwSharedProjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tgwId** | **string** | 조회할 Transit Gateway ID - [List transit gateways](/openapi/networking/tgw/list-transit-gateways)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTgwSharedProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListTgwSharedProjectsResponse**](ListTgwSharedProjectsResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransitGateways

> ListTransitGatewaysResponse ListTransitGateways(ctx).Id(id).Name(name).Region(region).IsShared(isShared).ProvisioningStatus(provisioningStatus).CreatedAt(createdAt).UpdatedAt(updatedAt).Offset(offset).Limit(limit).SortKeys(sortKeys).SortDirs(sortDirs).Execute()

List transit gateways



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/tgw"
)

func main() {
	id := "id_example" // string | Transit Gateway ID (optional)
	name := "name_example" // string | Transit Gateway 이름 (optional)
	region := "region_example" // string | Transit Gateway가 위치한 리전 (예: `kr-central-2`) (optional)
	isShared := true // bool | 공유된 Transit Gateway 여부 (예: `true`) (optional)
	provisioningStatus := openapiclient.TGWProvisioningStatus("ACTIVE") // TGWProvisioningStatus | Transit Gateway 프로비저닝 상태 (optional)
	createdAt := "createdAt_example" // string | 리소스가 생성된 시간 - ISO 8601 형식 - UTC 기준 (optional)
	updatedAt := "updatedAt_example" // string | 리소스가 마지막으로 수정된 시간 - ISO 8601 형식 - UTC 기준 (optional)
	offset := int32(56) // int32 | 조회 시작 위치 (optional)
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional)
	sortKeys := "sortKeys_example" // string | 정렬할 필드 (optional)
	sortDirs := "sortDirs_example" // string | 정렬 방향 (`asc`, `desc`) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransitGatewayAPI.ListTransitGateways(context.Background()).Id(id).Name(name).Region(region).IsShared(isShared).ProvisioningStatus(provisioningStatus).CreatedAt(createdAt).UpdatedAt(updatedAt).Offset(offset).Limit(limit).SortKeys(sortKeys).SortDirs(sortDirs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransitGatewayAPI.ListTransitGateways``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTransitGateways`: ListTransitGatewaysResponse
	fmt.Fprintf(os.Stdout, "Response from `TransitGatewayAPI.ListTransitGateways`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTransitGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Transit Gateway ID | 
 **name** | **string** | Transit Gateway 이름 | 
 **region** | **string** | Transit Gateway가 위치한 리전 (예: &#x60;kr-central-2&#x60;) | 
 **isShared** | **bool** | 공유된 Transit Gateway 여부 (예: &#x60;true&#x60;) | 
 **provisioningStatus** | [**TGWProvisioningStatus**](TGWProvisioningStatus.md) | Transit Gateway 프로비저닝 상태 | 
 **createdAt** | **string** | 리소스가 생성된 시간 - ISO 8601 형식 - UTC 기준 | 
 **updatedAt** | **string** | 리소스가 마지막으로 수정된 시간 - ISO 8601 형식 - UTC 기준 | 
 **offset** | **int32** | 조회 시작 위치 | 
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | 
 **sortKeys** | **string** | 정렬할 필드 | 
 **sortDirs** | **string** | 정렬 방향 (&#x60;asc&#x60;, &#x60;desc&#x60;) | 

### Return type

[**ListTransitGatewaysResponse**](ListTransitGatewaysResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShareTransitGateway

> ShareTransitGateway(ctx, tgwId, targetProjectId).Execute()

Share transit gateway



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/tgw"
)

func main() {
	tgwId := "tgwId_example" // string | 공유할 Transit Gateway ID
	targetProjectId := "targetProjectId_example" // string | 공유 대상 프로젝트 ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TransitGatewayAPI.ShareTransitGateway(context.Background(), tgwId, targetProjectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransitGatewayAPI.ShareTransitGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tgwId** | **string** | 공유할 Transit Gateway ID | 
**targetProjectId** | **string** | 공유 대상 프로젝트 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShareTransitGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## UnshareTransitGateway

> UnshareTransitGateway(ctx, tgwId, targetProjectId).Execute()

Unshare transit gateway



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/tgw"
)

func main() {
	tgwId := "tgwId_example" // string | 공유 해제할 Transit Gateway ID
	targetProjectId := "targetProjectId_example" // string | 공유 해제 대상 프로젝트 ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TransitGatewayAPI.UnshareTransitGateway(context.Background(), tgwId, targetProjectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransitGatewayAPI.UnshareTransitGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tgwId** | **string** | 공유 해제할 Transit Gateway ID | 
**targetProjectId** | **string** | 공유 해제 대상 프로젝트 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnshareTransitGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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

> UpdateTransitGatewayResponse UpdateTransitGateway(ctx, tgwId).UpdateTransitGatewayRequest(updateTransitGatewayRequest).Execute()

Update transit gateway



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/tgw"
)

func main() {
	tgwId := "tgwId_example" // string | 수정할 Transit Gateway ID
	updateTransitGatewayRequest := *openapiclient.NewUpdateTransitGatewayRequest(*openapiclient.NewUpdateTransitGateway()) // UpdateTransitGatewayRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransitGatewayAPI.UpdateTransitGateway(context.Background(), tgwId).UpdateTransitGatewayRequest(updateTransitGatewayRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransitGatewayAPI.UpdateTransitGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTransitGateway`: UpdateTransitGatewayResponse
	fmt.Fprintf(os.Stdout, "Response from `TransitGatewayAPI.UpdateTransitGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tgwId** | **string** | 수정할 Transit Gateway ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTransitGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTransitGatewayRequest** | [**UpdateTransitGatewayRequest**](UpdateTransitGatewayRequest.md) |  | 

### Return type

[**UpdateTransitGatewayResponse**](UpdateTransitGatewayResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

