# \TransitGatewayRouteTableAPI

All URIs are relative to *https://tgw.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTgwRoute**](TransitGatewayRouteTableAPI.md#CreateTgwRoute) | **Post** /api/v1/route-tables/{route_table_id}/routes | Create TGW route
[**CreateTgwRouteTable**](TransitGatewayRouteTableAPI.md#CreateTgwRouteTable) | **Post** /api/v1/route-tables | Create TGW route table
[**CreateTgwRouteTableAssociation**](TransitGatewayRouteTableAPI.md#CreateTgwRouteTableAssociation) | **Post** /api/v1/route-tables/{route_table_id}/associations | Create TGW route table association
[**DeleteTgwRoute**](TransitGatewayRouteTableAPI.md#DeleteTgwRoute) | **Delete** /api/v1/route-tables/{route_table_id}/routes/{route_id} | Delete TGW route
[**DeleteTgwRouteTable**](TransitGatewayRouteTableAPI.md#DeleteTgwRouteTable) | **Delete** /api/v1/route-tables/{route_table_id} | Delete TGW route table
[**DeleteTgwRouteTableAssociation**](TransitGatewayRouteTableAPI.md#DeleteTgwRouteTableAssociation) | **Delete** /api/v1/route-tables/{route_table_id}/associations/{association_id} | Delete TGW route table association
[**GetTgwRouteTable**](TransitGatewayRouteTableAPI.md#GetTgwRouteTable) | **Get** /api/v1/route-tables/{route_table_id} | Get TGW route table
[**ListTgwRouteTableAssociations**](TransitGatewayRouteTableAPI.md#ListTgwRouteTableAssociations) | **Get** /api/v1/route-tables/{route_table_id}/associations | List TGW route table associations
[**ListTgwRouteTables**](TransitGatewayRouteTableAPI.md#ListTgwRouteTables) | **Get** /api/v1/route-tables | List TGW route tables
[**ListTgwRoutes**](TransitGatewayRouteTableAPI.md#ListTgwRoutes) | **Get** /api/v1/route-tables/{route_table_id}/routes | List TGW routes
[**UpdateTgwRoute**](TransitGatewayRouteTableAPI.md#UpdateTgwRoute) | **Put** /api/v1/route-tables/{route_table_id}/routes/{route_id} | Update TGW route
[**UpdateTgwRouteTable**](TransitGatewayRouteTableAPI.md#UpdateTgwRouteTable) | **Put** /api/v1/route-tables/{route_table_id} | Update TGW route table



## CreateTgwRoute

> CreateTgwRouteResponse CreateTgwRoute(ctx, routeTableId).CreateTgwRouteRequest(createTgwRouteRequest).Execute()

Create TGW route



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
	routeTableId := "routeTableId_example" // string | Route가 속한 Transit Gateway 라우팅 테이블 ID
	createTgwRouteRequest := *openapiclient.NewCreateTgwRouteRequest(*openapiclient.NewCreateTgwRoute("DestinationCidrBlock_example", "TgwAttachmentId_example")) // CreateTgwRouteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransitGatewayRouteTableAPI.CreateTgwRoute(context.Background(), routeTableId).CreateTgwRouteRequest(createTgwRouteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransitGatewayRouteTableAPI.CreateTgwRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTgwRoute`: CreateTgwRouteResponse
	fmt.Fprintf(os.Stdout, "Response from `TransitGatewayRouteTableAPI.CreateTgwRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeTableId** | **string** | Route가 속한 Transit Gateway 라우팅 테이블 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTgwRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createTgwRouteRequest** | [**CreateTgwRouteRequest**](CreateTgwRouteRequest.md) |  | 

### Return type

[**CreateTgwRouteResponse**](CreateTgwRouteResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTgwRouteTable

> CreateTgwRouteTableResponse CreateTgwRouteTable(ctx).CreateTgwRouteTableRequest(createTgwRouteTableRequest).Execute()

Create TGW route table



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
	createTgwRouteTableRequest := *openapiclient.NewCreateTgwRouteTableRequest(*openapiclient.NewCreateTgwRouteTable("TgwId_example", "Name_example")) // CreateTgwRouteTableRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransitGatewayRouteTableAPI.CreateTgwRouteTable(context.Background()).CreateTgwRouteTableRequest(createTgwRouteTableRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransitGatewayRouteTableAPI.CreateTgwRouteTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTgwRouteTable`: CreateTgwRouteTableResponse
	fmt.Fprintf(os.Stdout, "Response from `TransitGatewayRouteTableAPI.CreateTgwRouteTable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTgwRouteTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTgwRouteTableRequest** | [**CreateTgwRouteTableRequest**](CreateTgwRouteTableRequest.md) |  | 

### Return type

[**CreateTgwRouteTableResponse**](CreateTgwRouteTableResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTgwRouteTableAssociation

> CreateTgwRouteTableAssociationResponse CreateTgwRouteTableAssociation(ctx, routeTableId).CreateTgwRouteTableAssociationRequest(createTgwRouteTableAssociationRequest).Execute()

Create TGW route table association



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
	routeTableId := "routeTableId_example" // string | 연결할 Transit Gateway 라우팅 테이블 ID
	createTgwRouteTableAssociationRequest := *openapiclient.NewCreateTgwRouteTableAssociationRequest(*openapiclient.NewCreateTgwRouteTableAssociation("TgwAttachmentId_example")) // CreateTgwRouteTableAssociationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransitGatewayRouteTableAPI.CreateTgwRouteTableAssociation(context.Background(), routeTableId).CreateTgwRouteTableAssociationRequest(createTgwRouteTableAssociationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransitGatewayRouteTableAPI.CreateTgwRouteTableAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTgwRouteTableAssociation`: CreateTgwRouteTableAssociationResponse
	fmt.Fprintf(os.Stdout, "Response from `TransitGatewayRouteTableAPI.CreateTgwRouteTableAssociation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeTableId** | **string** | 연결할 Transit Gateway 라우팅 테이블 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTgwRouteTableAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createTgwRouteTableAssociationRequest** | [**CreateTgwRouteTableAssociationRequest**](CreateTgwRouteTableAssociationRequest.md) |  | 

### Return type

[**CreateTgwRouteTableAssociationResponse**](CreateTgwRouteTableAssociationResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTgwRoute

> DeleteTgwRoute(ctx, routeTableId, routeId).Execute()

Delete TGW route



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
	routeTableId := "routeTableId_example" // string | Route를 삭제할 Transit Gateway 라우팅 테이블 ID
	routeId := "routeId_example" // string | 삭제할 Route ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TransitGatewayRouteTableAPI.DeleteTgwRoute(context.Background(), routeTableId, routeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransitGatewayRouteTableAPI.DeleteTgwRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeTableId** | **string** | Route를 삭제할 Transit Gateway 라우팅 테이블 ID | 
**routeId** | **string** | 삭제할 Route ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTgwRouteRequest struct via the builder pattern


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


## DeleteTgwRouteTable

> DeleteTgwRouteTable(ctx, routeTableId).Execute()

Delete TGW route table



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
	routeTableId := "routeTableId_example" // string | 삭제할 라우팅 테이블 ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TransitGatewayRouteTableAPI.DeleteTgwRouteTable(context.Background(), routeTableId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransitGatewayRouteTableAPI.DeleteTgwRouteTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeTableId** | **string** | 삭제할 라우팅 테이블 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTgwRouteTableRequest struct via the builder pattern


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


## DeleteTgwRouteTableAssociation

> DeleteTgwRouteTableAssociation(ctx, routeTableId, associationId).Execute()

Delete TGW route table association



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
	routeTableId := "routeTableId_example" // string | Association을 해제할 Transit Gateway 라우팅 테이블 ID
	associationId := "associationId_example" // string | 해제할 Association ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TransitGatewayRouteTableAPI.DeleteTgwRouteTableAssociation(context.Background(), routeTableId, associationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransitGatewayRouteTableAPI.DeleteTgwRouteTableAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeTableId** | **string** | Association을 해제할 Transit Gateway 라우팅 테이블 ID | 
**associationId** | **string** | 해제할 Association ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTgwRouteTableAssociationRequest struct via the builder pattern


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


## GetTgwRouteTable

> GetTgwRouteTableResponse GetTgwRouteTable(ctx, routeTableId).Execute()

Get TGW route table



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
	routeTableId := "routeTableId_example" // string | 조회할 Transit Gateway 라우팅 테이블 ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransitGatewayRouteTableAPI.GetTgwRouteTable(context.Background(), routeTableId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransitGatewayRouteTableAPI.GetTgwRouteTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTgwRouteTable`: GetTgwRouteTableResponse
	fmt.Fprintf(os.Stdout, "Response from `TransitGatewayRouteTableAPI.GetTgwRouteTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeTableId** | **string** | 조회할 Transit Gateway 라우팅 테이블 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTgwRouteTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetTgwRouteTableResponse**](GetTgwRouteTableResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTgwRouteTableAssociations

> ListTgwRouteTableAssociationsResponse ListTgwRouteTableAssociations(ctx, routeTableId).ResourceName(resourceName).ResourceId(resourceId).ResourceProvisioningStatus(resourceProvisioningStatus).ResourceType(resourceType).ProvisioningStatus(provisioningStatus).ResourceAttachmentId(resourceAttachmentId).Offset(offset).Limit(limit).SortKeys(sortKeys).SortDirs(sortDirs).Execute()

List TGW route table associations



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
	routeTableId := "routeTableId_example" // string | 조회할 Transit Gateway 라우팅 테이블 ID - [List TGW route tables](/openapi/networking/tgw/list-tgw-route-tables)에서 확인
	resourceName := "resourceName_example" // string | 연결된 리소스 이름 (optional)
	resourceId := "resourceId_example" // string | 연결된 리소스 ID (optional)
	resourceProvisioningStatus := "resourceProvisioningStatus_example" // string | 리소스의 프로비저닝 상태 (optional)
	resourceType := openapiclient.ResourceType("vpc") // ResourceType | 연결된 리소스 유형 (optional)
	provisioningStatus := openapiclient.TGWAssociationProvisioningStatus("ACTIVE") // TGWAssociationProvisioningStatus | Association의 프로비저닝 상태 (optional)
	resourceAttachmentId := "resourceAttachmentId_example" // string | 연결된 Attachment ID (optional)
	offset := int32(56) // int32 | 조회 시작 위치 (optional)
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional)
	sortKeys := "sortKeys_example" // string | 정렬할 필드 (optional)
	sortDirs := "sortDirs_example" // string | 정렬 방향 (`asc`, `desc`) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransitGatewayRouteTableAPI.ListTgwRouteTableAssociations(context.Background(), routeTableId).ResourceName(resourceName).ResourceId(resourceId).ResourceProvisioningStatus(resourceProvisioningStatus).ResourceType(resourceType).ProvisioningStatus(provisioningStatus).ResourceAttachmentId(resourceAttachmentId).Offset(offset).Limit(limit).SortKeys(sortKeys).SortDirs(sortDirs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransitGatewayRouteTableAPI.ListTgwRouteTableAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTgwRouteTableAssociations`: ListTgwRouteTableAssociationsResponse
	fmt.Fprintf(os.Stdout, "Response from `TransitGatewayRouteTableAPI.ListTgwRouteTableAssociations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeTableId** | **string** | 조회할 Transit Gateway 라우팅 테이블 ID - [List TGW route tables](/openapi/networking/tgw/list-tgw-route-tables)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTgwRouteTableAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resourceName** | **string** | 연결된 리소스 이름 | 
 **resourceId** | **string** | 연결된 리소스 ID | 
 **resourceProvisioningStatus** | **string** | 리소스의 프로비저닝 상태 | 
 **resourceType** | [**ResourceType**](ResourceType.md) | 연결된 리소스 유형 | 
 **provisioningStatus** | [**TGWAssociationProvisioningStatus**](TGWAssociationProvisioningStatus.md) | Association의 프로비저닝 상태 | 
 **resourceAttachmentId** | **string** | 연결된 Attachment ID | 
 **offset** | **int32** | 조회 시작 위치 | 
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | 
 **sortKeys** | **string** | 정렬할 필드 | 
 **sortDirs** | **string** | 정렬 방향 (&#x60;asc&#x60;, &#x60;desc&#x60;) | 

### Return type

[**ListTgwRouteTableAssociationsResponse**](ListTgwRouteTableAssociationsResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTgwRouteTables

> ListTgwRouteTablesResponse ListTgwRouteTables(ctx).Id(id).Name(name).TgwId(tgwId).TgwName(tgwName).ProvisioningStatus(provisioningStatus).CreatedAt(createdAt).UpdatedAt(updatedAt).Offset(offset).Limit(limit).SortKeys(sortKeys).SortDirs(sortDirs).Execute()

List TGW route tables



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
	id := "id_example" // string | 라우팅 테이블 ID (optional)
	name := "name_example" // string | 라우팅 테이블 이름 (optional)
	tgwId := "tgwId_example" // string | Transit Gateway ID (optional)
	tgwName := "tgwName_example" // string | Transit Gateway 이름 (optional)
	provisioningStatus := openapiclient.TGWRouteTableProvisioningStatus("ACTIVE") // TGWRouteTableProvisioningStatus | 라우팅 테이블의 프로비저닝 상태 (optional)
	createdAt := "createdAt_example" // string | 리소스가 생성된 시간 - ISO 8601 형식 - UTC 기준 (optional)
	updatedAt := "updatedAt_example" // string | 리소스가 마지막으로 수정된 시간 - ISO 8601 형식 - UTC 기준 (optional)
	offset := int32(56) // int32 | 조회 시작 위치 (optional)
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional)
	sortKeys := "sortKeys_example" // string | 정렬할 필드 (optional)
	sortDirs := "sortDirs_example" // string | 정렬 방향 (`asc`, `desc`) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransitGatewayRouteTableAPI.ListTgwRouteTables(context.Background()).Id(id).Name(name).TgwId(tgwId).TgwName(tgwName).ProvisioningStatus(provisioningStatus).CreatedAt(createdAt).UpdatedAt(updatedAt).Offset(offset).Limit(limit).SortKeys(sortKeys).SortDirs(sortDirs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransitGatewayRouteTableAPI.ListTgwRouteTables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTgwRouteTables`: ListTgwRouteTablesResponse
	fmt.Fprintf(os.Stdout, "Response from `TransitGatewayRouteTableAPI.ListTgwRouteTables`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTgwRouteTablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | 라우팅 테이블 ID | 
 **name** | **string** | 라우팅 테이블 이름 | 
 **tgwId** | **string** | Transit Gateway ID | 
 **tgwName** | **string** | Transit Gateway 이름 | 
 **provisioningStatus** | [**TGWRouteTableProvisioningStatus**](TGWRouteTableProvisioningStatus.md) | 라우팅 테이블의 프로비저닝 상태 | 
 **createdAt** | **string** | 리소스가 생성된 시간 - ISO 8601 형식 - UTC 기준 | 
 **updatedAt** | **string** | 리소스가 마지막으로 수정된 시간 - ISO 8601 형식 - UTC 기준 | 
 **offset** | **int32** | 조회 시작 위치 | 
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | 
 **sortKeys** | **string** | 정렬할 필드 | 
 **sortDirs** | **string** | 정렬 방향 (&#x60;asc&#x60;, &#x60;desc&#x60;) | 

### Return type

[**ListTgwRouteTablesResponse**](ListTgwRouteTablesResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTgwRoutes

> ListTgwRoutesResponse ListTgwRoutes(ctx, routeTableId).DestinationCidrBlock(destinationCidrBlock).RouteType(routeType).ProvisioningStatus(provisioningStatus).ResourceType(resourceType).ResourceId(resourceId).ResourceName(resourceName).ResourceProvisioningStatus(resourceProvisioningStatus).ResourceAttachmentId(resourceAttachmentId).Offset(offset).Limit(limit).SortKeys(sortKeys).SortDirs(sortDirs).Execute()

List TGW routes



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
	routeTableId := "routeTableId_example" // string | 조회할 Transit Gateway 라우팅 테이블 ID - [List TGW route tables](/openapi/networking/tgw/list-tgw-route-tables)에서 확인
	destinationCidrBlock := "destinationCidrBlock_example" // string | 목적지 CIDR 블록 (optional)
	routeType := "routeType_example" // string | Route 유형 (optional)
	provisioningStatus := openapiclient.TGWRouteProvisioningStatus("ACTIVE") // TGWRouteProvisioningStatus | Route의 프로비저닝 상태 (optional)
	resourceType := openapiclient.ResourceType("vpc") // ResourceType | 연결된 리소스 유형 (optional)
	resourceId := "resourceId_example" // string | 연결된 리소스 ID (optional)
	resourceName := "resourceName_example" // string | 연결된 리소스 이름 (optional)
	resourceProvisioningStatus := "resourceProvisioningStatus_example" // string | 연결된 리소스의 프로비저닝 상태 (optional)
	resourceAttachmentId := "resourceAttachmentId_example" // string | 연결된 Attachment ID (optional)
	offset := int32(56) // int32 | 조회 시작 위치 (optional)
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional)
	sortKeys := "sortKeys_example" // string | 정렬할 필드 (optional)
	sortDirs := "sortDirs_example" // string | 정렬 방향 (`asc`, `desc`) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransitGatewayRouteTableAPI.ListTgwRoutes(context.Background(), routeTableId).DestinationCidrBlock(destinationCidrBlock).RouteType(routeType).ProvisioningStatus(provisioningStatus).ResourceType(resourceType).ResourceId(resourceId).ResourceName(resourceName).ResourceProvisioningStatus(resourceProvisioningStatus).ResourceAttachmentId(resourceAttachmentId).Offset(offset).Limit(limit).SortKeys(sortKeys).SortDirs(sortDirs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransitGatewayRouteTableAPI.ListTgwRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTgwRoutes`: ListTgwRoutesResponse
	fmt.Fprintf(os.Stdout, "Response from `TransitGatewayRouteTableAPI.ListTgwRoutes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeTableId** | **string** | 조회할 Transit Gateway 라우팅 테이블 ID - [List TGW route tables](/openapi/networking/tgw/list-tgw-route-tables)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTgwRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **destinationCidrBlock** | **string** | 목적지 CIDR 블록 | 
 **routeType** | **string** | Route 유형 | 
 **provisioningStatus** | [**TGWRouteProvisioningStatus**](TGWRouteProvisioningStatus.md) | Route의 프로비저닝 상태 | 
 **resourceType** | [**ResourceType**](ResourceType.md) | 연결된 리소스 유형 | 
 **resourceId** | **string** | 연결된 리소스 ID | 
 **resourceName** | **string** | 연결된 리소스 이름 | 
 **resourceProvisioningStatus** | **string** | 연결된 리소스의 프로비저닝 상태 | 
 **resourceAttachmentId** | **string** | 연결된 Attachment ID | 
 **offset** | **int32** | 조회 시작 위치 | 
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | 
 **sortKeys** | **string** | 정렬할 필드 | 
 **sortDirs** | **string** | 정렬 방향 (&#x60;asc&#x60;, &#x60;desc&#x60;) | 

### Return type

[**ListTgwRoutesResponse**](ListTgwRoutesResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTgwRoute

> UpdateTgwRouteResponse UpdateTgwRoute(ctx, routeTableId, routeId).UpdateTgwRouteRequest(updateTgwRouteRequest).Execute()

Update TGW route



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
	routeTableId := "routeTableId_example" // string | 수정할 Transit Gateway 라우팅 테이블 ID
	routeId := "routeId_example" // string | 수정할 Route ID
	updateTgwRouteRequest := *openapiclient.NewUpdateTgwRouteRequest(*openapiclient.NewUpdateTgwRoute("TgwAttachmentId_example")) // UpdateTgwRouteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransitGatewayRouteTableAPI.UpdateTgwRoute(context.Background(), routeTableId, routeId).UpdateTgwRouteRequest(updateTgwRouteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransitGatewayRouteTableAPI.UpdateTgwRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTgwRoute`: UpdateTgwRouteResponse
	fmt.Fprintf(os.Stdout, "Response from `TransitGatewayRouteTableAPI.UpdateTgwRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeTableId** | **string** | 수정할 Transit Gateway 라우팅 테이블 ID | 
**routeId** | **string** | 수정할 Route ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTgwRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateTgwRouteRequest** | [**UpdateTgwRouteRequest**](UpdateTgwRouteRequest.md) |  | 

### Return type

[**UpdateTgwRouteResponse**](UpdateTgwRouteResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTgwRouteTable

> UpdateTgwRouteTableResponse UpdateTgwRouteTable(ctx, routeTableId).UpdateTgwRouteTableRequest(updateTgwRouteTableRequest).Execute()

Update TGW route table



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
	routeTableId := "routeTableId_example" // string | 수정할 Transit Gateway 라우팅 테이블 ID
	updateTgwRouteTableRequest := *openapiclient.NewUpdateTgwRouteTableRequest(*openapiclient.NewUpdateTgwRouteTable("Name_example")) // UpdateTgwRouteTableRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransitGatewayRouteTableAPI.UpdateTgwRouteTable(context.Background(), routeTableId).UpdateTgwRouteTableRequest(updateTgwRouteTableRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransitGatewayRouteTableAPI.UpdateTgwRouteTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTgwRouteTable`: UpdateTgwRouteTableResponse
	fmt.Fprintf(os.Stdout, "Response from `TransitGatewayRouteTableAPI.UpdateTgwRouteTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeTableId** | **string** | 수정할 Transit Gateway 라우팅 테이블 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTgwRouteTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTgwRouteTableRequest** | [**UpdateTgwRouteTableRequest**](UpdateTgwRouteTableRequest.md) |  | 

### Return type

[**UpdateTgwRouteTableResponse**](UpdateTgwRouteTableResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

