# \RouteTablesAPI

All URIs are relative to *https://tgw.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTgwRoute**](RouteTablesAPI.md#CreateTgwRoute) | **Post** /api/v1/route-tables/{route_table_id}/routes | Create TGW route
[**CreateTgwRouteTable**](RouteTablesAPI.md#CreateTgwRouteTable) | **Post** /api/v1/route-tables | Create TGW route table
[**CreateTgwRouteTableAssociation**](RouteTablesAPI.md#CreateTgwRouteTableAssociation) | **Post** /api/v1/route-tables/{route_table_id}/associations | Create TGW route table association
[**DeleteTgwRoute**](RouteTablesAPI.md#DeleteTgwRoute) | **Delete** /api/v1/route-tables/{route_table_id}/routes/{route_id} | Delete TGW route
[**DeleteTgwRouteTable**](RouteTablesAPI.md#DeleteTgwRouteTable) | **Delete** /api/v1/route-tables/{route_table_id} | Delete TGW route table
[**DeleteTgwRouteTableAssociation**](RouteTablesAPI.md#DeleteTgwRouteTableAssociation) | **Delete** /api/v1/route-tables/{route_table_id}/associations/{association_id} | Delete TGW route table association
[**GetTgwRouteTable**](RouteTablesAPI.md#GetTgwRouteTable) | **Get** /api/v1/route-tables/{route_table_id} | Get TGW route table
[**ListTgwRouteTableAssociations**](RouteTablesAPI.md#ListTgwRouteTableAssociations) | **Get** /api/v1/route-tables/{route_table_id}/associations | List TGW route table associations
[**ListTgwRouteTables**](RouteTablesAPI.md#ListTgwRouteTables) | **Get** /api/v1/route-tables | List TGW route tables
[**ListTgwRoutes**](RouteTablesAPI.md#ListTgwRoutes) | **Get** /api/v1/route-tables/{route_table_id}/routes | List TGW routes
[**UpdateTgwRouteTable**](RouteTablesAPI.md#UpdateTgwRouteTable) | **Put** /api/v1/route-tables/{route_table_id} | Update TGW route table
[**UpdateTgwRouteTableRoute**](RouteTablesAPI.md#UpdateTgwRouteTableRoute) | **Put** /api/v1/route-tables/{route_table_id}/routes/{route_id} | Update TGW route



## CreateTgwRoute

> BnsTgwV1ApiCreateTgwRouteModelCreateTgwRouteTableRouteResponseModel CreateTgwRoute(ctx, routeTableId).XAuthToken(xAuthToken).CreateTgwRouteTableRouteRequestModel(createTgwRouteTableRouteRequestModel).Execute()

Create TGW route



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
	routeTableId := "routeTableId_example" // string | Route가 속한 Transit Gateway 라우팅 테이블 ID 
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	createTgwRouteTableRouteRequestModel := *openapiclient.NewCreateTgwRouteTableRouteRequestModel(*openapiclient.NewBnsTgwV1ApiCreateTgwRouteModelTgwRouteRequestModel("DestinationCidrBlock_example", "TgwAttachmentId_example")) // CreateTgwRouteTableRouteRequestModel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteTablesAPI.CreateTgwRoute(context.Background(), routeTableId).XAuthToken(xAuthToken).CreateTgwRouteTableRouteRequestModel(createTgwRouteTableRouteRequestModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteTablesAPI.CreateTgwRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTgwRoute`: BnsTgwV1ApiCreateTgwRouteModelCreateTgwRouteTableRouteResponseModel
	fmt.Fprintf(os.Stdout, "Response from `RouteTablesAPI.CreateTgwRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeTableId** | **string** | Route가 속한 Transit Gateway 라우팅 테이블 ID  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTgwRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **createTgwRouteTableRouteRequestModel** | [**CreateTgwRouteTableRouteRequestModel**](CreateTgwRouteTableRouteRequestModel.md) |  | 

### Return type

[**BnsTgwV1ApiCreateTgwRouteModelCreateTgwRouteTableRouteResponseModel**](BnsTgwV1ApiCreateTgwRouteModelCreateTgwRouteTableRouteResponseModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTgwRouteTable

> BnsTgwV1ApiCreateTgwRouteTableModelCreateTgwRouteTableResponseModel CreateTgwRouteTable(ctx).XAuthToken(xAuthToken).CreateTgwRouteTableRequestModel(createTgwRouteTableRequestModel).Execute()

Create TGW route table



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
	createTgwRouteTableRequestModel := *openapiclient.NewCreateTgwRouteTableRequestModel(*openapiclient.NewBnsTgwV1ApiCreateTgwRouteTableModelTgwRouteTableRequestModel("TgwId_example", "Name_example")) // CreateTgwRouteTableRequestModel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteTablesAPI.CreateTgwRouteTable(context.Background()).XAuthToken(xAuthToken).CreateTgwRouteTableRequestModel(createTgwRouteTableRequestModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteTablesAPI.CreateTgwRouteTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTgwRouteTable`: BnsTgwV1ApiCreateTgwRouteTableModelCreateTgwRouteTableResponseModel
	fmt.Fprintf(os.Stdout, "Response from `RouteTablesAPI.CreateTgwRouteTable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTgwRouteTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **createTgwRouteTableRequestModel** | [**CreateTgwRouteTableRequestModel**](CreateTgwRouteTableRequestModel.md) |  | 

### Return type

[**BnsTgwV1ApiCreateTgwRouteTableModelCreateTgwRouteTableResponseModel**](BnsTgwV1ApiCreateTgwRouteTableModelCreateTgwRouteTableResponseModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTgwRouteTableAssociation

> CreateTgwRouteTableAssociationResponseModel CreateTgwRouteTableAssociation(ctx, routeTableId).XAuthToken(xAuthToken).CreateTgwRouteTableAssociationRequestModel(createTgwRouteTableAssociationRequestModel).Execute()

Create TGW route table association



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
	routeTableId := "routeTableId_example" // string | 연결할 Transit Gateway 라우팅 테이블 ID 
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	createTgwRouteTableAssociationRequestModel := *openapiclient.NewCreateTgwRouteTableAssociationRequestModel(*openapiclient.NewTgwRouteTableAssociationRequestModel("TgwAttachmentId_example")) // CreateTgwRouteTableAssociationRequestModel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteTablesAPI.CreateTgwRouteTableAssociation(context.Background(), routeTableId).XAuthToken(xAuthToken).CreateTgwRouteTableAssociationRequestModel(createTgwRouteTableAssociationRequestModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteTablesAPI.CreateTgwRouteTableAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTgwRouteTableAssociation`: CreateTgwRouteTableAssociationResponseModel
	fmt.Fprintf(os.Stdout, "Response from `RouteTablesAPI.CreateTgwRouteTableAssociation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeTableId** | **string** | 연결할 Transit Gateway 라우팅 테이블 ID  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTgwRouteTableAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **createTgwRouteTableAssociationRequestModel** | [**CreateTgwRouteTableAssociationRequestModel**](CreateTgwRouteTableAssociationRequestModel.md) |  | 

### Return type

[**CreateTgwRouteTableAssociationResponseModel**](CreateTgwRouteTableAssociationResponseModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTgwRoute

> DeleteTgwRoute(ctx, routeTableId, routeId).XAuthToken(xAuthToken).Execute()

Delete TGW route



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
	routeTableId := "routeTableId_example" // string | Route를 삭제할 Transit Gateway 라우팅 테이블 ID 
	routeId := "routeId_example" // string | 삭제할 Route ID 
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RouteTablesAPI.DeleteTgwRoute(context.Background(), routeTableId, routeId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteTablesAPI.DeleteTgwRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeTableId** | **string** | Route를 삭제할 Transit Gateway 라우팅 테이블 ID  | 
**routeId** | **string** | 삭제할 Route ID  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTgwRouteRequest struct via the builder pattern


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


## DeleteTgwRouteTable

> DeleteTgwRouteTable(ctx, routeTableId).XAuthToken(xAuthToken).Execute()

Delete TGW route table



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
	routeTableId := "routeTableId_example" // string | 삭제할 라우팅 테이블 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RouteTablesAPI.DeleteTgwRouteTable(context.Background(), routeTableId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteTablesAPI.DeleteTgwRouteTable``: %v\n", err)
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


## DeleteTgwRouteTableAssociation

> DeleteTgwRouteTableAssociation(ctx, routeTableId, associationId).XAuthToken(xAuthToken).Execute()

Delete TGW route table association



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
	routeTableId := "routeTableId_example" // string | Association을 해제할 Transit Gateway 라우팅 테이블 ID 
	associationId := "associationId_example" // string | 해제할 Association ID 
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RouteTablesAPI.DeleteTgwRouteTableAssociation(context.Background(), routeTableId, associationId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteTablesAPI.DeleteTgwRouteTableAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeTableId** | **string** | Association을 해제할 Transit Gateway 라우팅 테이블 ID  | 
**associationId** | **string** | 해제할 Association ID  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTgwRouteTableAssociationRequest struct via the builder pattern


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


## GetTgwRouteTable

> GetTgwRouteTableResponseModel GetTgwRouteTable(ctx, routeTableId).XAuthToken(xAuthToken).Execute()

Get TGW route table



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
	routeTableId := "routeTableId_example" // string | 조회할 Transit Gateway 라우팅 테이블 ID 
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteTablesAPI.GetTgwRouteTable(context.Background(), routeTableId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteTablesAPI.GetTgwRouteTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTgwRouteTable`: GetTgwRouteTableResponseModel
	fmt.Fprintf(os.Stdout, "Response from `RouteTablesAPI.GetTgwRouteTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeTableId** | **string** | 조회할 Transit Gateway 라우팅 테이블 ID  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTgwRouteTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**GetTgwRouteTableResponseModel**](GetTgwRouteTableResponseModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTgwRouteTableAssociations

> GetTgwRouteTableAssociationsResponseModel ListTgwRouteTableAssociations(ctx, routeTableId).XAuthToken(xAuthToken).ResourceName(resourceName).ResourceId(resourceId).ResourceProvisioningStatus(resourceProvisioningStatus).ResourceType(resourceType).ProvisioningStatus(provisioningStatus).ResourceAttachmentId(resourceAttachmentId).Offset(offset).Limit(limit).SortKeys(sortKeys).SortDirs(sortDirs).Execute()

List TGW route table associations



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
	routeTableId := "routeTableId_example" // string | 조회할 Transit Gateway 라우팅 테이블 ID 
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	resourceName := "resourceName_example" // string | 연결된 리소스 이름  (optional)
	resourceId := "resourceId_example" // string | 연결된 리소스 ID  (optional)
	resourceProvisioningStatus := "resourceProvisioningStatus_example" // string | 리소스의 프로비저닝 상태   (optional)
	resourceType := openapiclient.ResourceType("vpc") // ResourceType | 연결된 리소스 유형  (optional)
	provisioningStatus := openapiclient.ProvisioningStatus("ACTIVE") // ProvisioningStatus | Association의 프로비저닝 상태   (optional)
	resourceAttachmentId := "resourceAttachmentId_example" // string | 연결된 Attachment ID  (optional)
	offset := int32(56) // int32 | 조회 시작 위치 (optional) (default to 0)
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional) (default to 20)
	sortKeys := "sortKeys_example" // string | 정렬할 필드  (optional) (default to "created_at")
	sortDirs := "sortDirs_example" // string | 정렬 방향 (`asc`, `desc`)    (optional) (default to "desc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteTablesAPI.ListTgwRouteTableAssociations(context.Background(), routeTableId).XAuthToken(xAuthToken).ResourceName(resourceName).ResourceId(resourceId).ResourceProvisioningStatus(resourceProvisioningStatus).ResourceType(resourceType).ProvisioningStatus(provisioningStatus).ResourceAttachmentId(resourceAttachmentId).Offset(offset).Limit(limit).SortKeys(sortKeys).SortDirs(sortDirs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteTablesAPI.ListTgwRouteTableAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTgwRouteTableAssociations`: GetTgwRouteTableAssociationsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `RouteTablesAPI.ListTgwRouteTableAssociations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeTableId** | **string** | 조회할 Transit Gateway 라우팅 테이블 ID  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTgwRouteTableAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **resourceName** | **string** | 연결된 리소스 이름  | 
 **resourceId** | **string** | 연결된 리소스 ID  | 
 **resourceProvisioningStatus** | **string** | 리소스의 프로비저닝 상태   | 
 **resourceType** | [**ResourceType**](ResourceType.md) | 연결된 리소스 유형  | 
 **provisioningStatus** | [**ProvisioningStatus**](ProvisioningStatus.md) | Association의 프로비저닝 상태   | 
 **resourceAttachmentId** | **string** | 연결된 Attachment ID  | 
 **offset** | **int32** | 조회 시작 위치 | [default to 0]
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | [default to 20]
 **sortKeys** | **string** | 정렬할 필드  | [default to &quot;created_at&quot;]
 **sortDirs** | **string** | 정렬 방향 (&#x60;asc&#x60;, &#x60;desc&#x60;)    | [default to &quot;desc&quot;]

### Return type

[**GetTgwRouteTableAssociationsResponseModel**](GetTgwRouteTableAssociationsResponseModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTgwRouteTables

> GetTgwRouteTablesResponseModel ListTgwRouteTables(ctx).XAuthToken(xAuthToken).Id(id).Name(name).TgwId(tgwId).TgwName(tgwName).ProvisioningStatus(provisioningStatus).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Limit(limit).Offset(offset).Execute()

List TGW route tables



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
	id := "id_example" // string | 라우팅 테이블 ID  (optional)
	name := "name_example" // string | 라우팅 테이블 이름  (optional)
	tgwId := "tgwId_example" // string | Transit Gateway ID  (optional)
	tgwName := "tgwName_example" // string | Transit Gateway 이름  (optional)
	provisioningStatus := openapiclient.ProvisioningStatus("ACTIVE") // ProvisioningStatus | 라우팅 테이블의 프로비저닝 상태   (optional)
	createdAt := "createdAt_example" // string | 리소스가 생성된 시간 <br/> - ISO_8601 형식  <br/> - UTC 기준 (optional)
	updatedAt := "updatedAt_example" // string | 리소스가 마지막으로 수정된 시간 <br/> - ISO_8601 형식  <br/> - UTC 기준 (optional)
	sortKeys := "sortKeys_example" // string | 정렬할 필드  (optional) (default to "created_at")
	sortDirs := "sortDirs_example" // string | 정렬 방향 (`asc`, `desc`)    (optional) (default to "desc")
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional) (default to 20)
	offset := int32(56) // int32 | 조회 시작 위치 (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteTablesAPI.ListTgwRouteTables(context.Background()).XAuthToken(xAuthToken).Id(id).Name(name).TgwId(tgwId).TgwName(tgwName).ProvisioningStatus(provisioningStatus).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteTablesAPI.ListTgwRouteTables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTgwRouteTables`: GetTgwRouteTablesResponseModel
	fmt.Fprintf(os.Stdout, "Response from `RouteTablesAPI.ListTgwRouteTables`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTgwRouteTablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **id** | **string** | 라우팅 테이블 ID  | 
 **name** | **string** | 라우팅 테이블 이름  | 
 **tgwId** | **string** | Transit Gateway ID  | 
 **tgwName** | **string** | Transit Gateway 이름  | 
 **provisioningStatus** | [**ProvisioningStatus**](ProvisioningStatus.md) | 라우팅 테이블의 프로비저닝 상태   | 
 **createdAt** | **string** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
 **updatedAt** | **string** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
 **sortKeys** | **string** | 정렬할 필드  | [default to &quot;created_at&quot;]
 **sortDirs** | **string** | 정렬 방향 (&#x60;asc&#x60;, &#x60;desc&#x60;)    | [default to &quot;desc&quot;]
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | [default to 20]
 **offset** | **int32** | 조회 시작 위치 | [default to 0]

### Return type

[**GetTgwRouteTablesResponseModel**](GetTgwRouteTablesResponseModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTgwRoutes

> GetTgwRouteTableRoutesResponseModel ListTgwRoutes(ctx, routeTableId).XAuthToken(xAuthToken).DestinationCidrBlock(destinationCidrBlock).RouteType(routeType).ProvisioningStatus(provisioningStatus).ResourceType(resourceType).ResourceId(resourceId).ResourceName(resourceName).ResourceProvisioningStatus(resourceProvisioningStatus).ResourceAttachmentId(resourceAttachmentId).Offset(offset).Limit(limit).SortKeys(sortKeys).SortDirs(sortDirs).Execute()

List TGW routes



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
	routeTableId := "routeTableId_example" // string | 조회할 Transit Gateway 라우팅 테이블 ID 
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	destinationCidrBlock := "destinationCidrBlock_example" // string | 목적지 CIDR 블록  (optional)
	routeType := "routeType_example" // string | Route 유형 (예: `static`, `propagated`)  (optional)
	provisioningStatus := openapiclient.ProvisioningStatus("ACTIVE") // ProvisioningStatus | Route의 프로비저닝 상태   (optional)
	resourceType := "resourceType_example" // string | 연결된 리소스 유형  (optional)
	resourceId := "resourceId_example" // string | 연결된 리소스 ID  (optional)
	resourceName := "resourceName_example" // string | 연결된 리소스 이름  (optional)
	resourceProvisioningStatus := "resourceProvisioningStatus_example" // string | 연결된 리소스의 프로비저닝 상태   (optional)
	resourceAttachmentId := "resourceAttachmentId_example" // string | 연결된 Attachment ID  (optional)
	offset := int32(56) // int32 | 조회 시작 위치 (optional) (default to 0)
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional) (default to 20)
	sortKeys := "sortKeys_example" // string | 정렬할 필드  (optional) (default to "created_at")
	sortDirs := "sortDirs_example" // string | 정렬 방향 (`asc`, `desc`)    (optional) (default to "desc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteTablesAPI.ListTgwRoutes(context.Background(), routeTableId).XAuthToken(xAuthToken).DestinationCidrBlock(destinationCidrBlock).RouteType(routeType).ProvisioningStatus(provisioningStatus).ResourceType(resourceType).ResourceId(resourceId).ResourceName(resourceName).ResourceProvisioningStatus(resourceProvisioningStatus).ResourceAttachmentId(resourceAttachmentId).Offset(offset).Limit(limit).SortKeys(sortKeys).SortDirs(sortDirs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteTablesAPI.ListTgwRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTgwRoutes`: GetTgwRouteTableRoutesResponseModel
	fmt.Fprintf(os.Stdout, "Response from `RouteTablesAPI.ListTgwRoutes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeTableId** | **string** | 조회할 Transit Gateway 라우팅 테이블 ID  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTgwRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **destinationCidrBlock** | **string** | 목적지 CIDR 블록  | 
 **routeType** | **string** | Route 유형 (예: &#x60;static&#x60;, &#x60;propagated&#x60;)  | 
 **provisioningStatus** | [**ProvisioningStatus**](ProvisioningStatus.md) | Route의 프로비저닝 상태   | 
 **resourceType** | **string** | 연결된 리소스 유형  | 
 **resourceId** | **string** | 연결된 리소스 ID  | 
 **resourceName** | **string** | 연결된 리소스 이름  | 
 **resourceProvisioningStatus** | **string** | 연결된 리소스의 프로비저닝 상태   | 
 **resourceAttachmentId** | **string** | 연결된 Attachment ID  | 
 **offset** | **int32** | 조회 시작 위치 | [default to 0]
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | [default to 20]
 **sortKeys** | **string** | 정렬할 필드  | [default to &quot;created_at&quot;]
 **sortDirs** | **string** | 정렬 방향 (&#x60;asc&#x60;, &#x60;desc&#x60;)    | [default to &quot;desc&quot;]

### Return type

[**GetTgwRouteTableRoutesResponseModel**](GetTgwRouteTableRoutesResponseModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTgwRouteTable

> BnsTgwV1ApiUpdateTgwRouteTableModelCreateTgwRouteTableResponseModel UpdateTgwRouteTable(ctx, routeTableId).XAuthToken(xAuthToken).UpdateTgwRouteTableRequestModel(updateTgwRouteTableRequestModel).Execute()

Update TGW route table



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
	routeTableId := "routeTableId_example" // string | 수정할 Transit Gateway 라우팅 테이블 ID  
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	updateTgwRouteTableRequestModel := *openapiclient.NewUpdateTgwRouteTableRequestModel(*openapiclient.NewBnsTgwV1ApiUpdateTgwRouteTableModelTgwRouteTableRequestModel("Name_example")) // UpdateTgwRouteTableRequestModel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteTablesAPI.UpdateTgwRouteTable(context.Background(), routeTableId).XAuthToken(xAuthToken).UpdateTgwRouteTableRequestModel(updateTgwRouteTableRequestModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteTablesAPI.UpdateTgwRouteTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTgwRouteTable`: BnsTgwV1ApiUpdateTgwRouteTableModelCreateTgwRouteTableResponseModel
	fmt.Fprintf(os.Stdout, "Response from `RouteTablesAPI.UpdateTgwRouteTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeTableId** | **string** | 수정할 Transit Gateway 라우팅 테이블 ID   | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTgwRouteTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **updateTgwRouteTableRequestModel** | [**UpdateTgwRouteTableRequestModel**](UpdateTgwRouteTableRequestModel.md) |  | 

### Return type

[**BnsTgwV1ApiUpdateTgwRouteTableModelCreateTgwRouteTableResponseModel**](BnsTgwV1ApiUpdateTgwRouteTableModelCreateTgwRouteTableResponseModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTgwRouteTableRoute

> BnsTgwV1ApiUpdateTgwRouteModelCreateTgwRouteTableRouteResponseModel UpdateTgwRouteTableRoute(ctx, routeTableId, routeId).XAuthToken(xAuthToken).UpdateTgwRouteTableRouteRequestModel(updateTgwRouteTableRouteRequestModel).Execute()

Update TGW route



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
	routeTableId := "routeTableId_example" // string | 수정할 Transit Gateway 라우팅 테이블 ID 
	routeId := "routeId_example" // string | 수정할 Route ID  
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	updateTgwRouteTableRouteRequestModel := *openapiclient.NewUpdateTgwRouteTableRouteRequestModel(*openapiclient.NewBnsTgwV1ApiUpdateTgwRouteModelTgwRouteRequestModel("TgwAttachmentId_example")) // UpdateTgwRouteTableRouteRequestModel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteTablesAPI.UpdateTgwRouteTableRoute(context.Background(), routeTableId, routeId).XAuthToken(xAuthToken).UpdateTgwRouteTableRouteRequestModel(updateTgwRouteTableRouteRequestModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteTablesAPI.UpdateTgwRouteTableRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTgwRouteTableRoute`: BnsTgwV1ApiUpdateTgwRouteModelCreateTgwRouteTableRouteResponseModel
	fmt.Fprintf(os.Stdout, "Response from `RouteTablesAPI.UpdateTgwRouteTableRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeTableId** | **string** | 수정할 Transit Gateway 라우팅 테이블 ID  | 
**routeId** | **string** | 수정할 Route ID   | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTgwRouteTableRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **updateTgwRouteTableRouteRequestModel** | [**UpdateTgwRouteTableRouteRequestModel**](UpdateTgwRouteTableRouteRequestModel.md) |  | 

### Return type

[**BnsTgwV1ApiUpdateTgwRouteModelCreateTgwRouteTableRouteResponseModel**](BnsTgwV1ApiUpdateTgwRouteModelCreateTgwRouteTableRouteResponseModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

