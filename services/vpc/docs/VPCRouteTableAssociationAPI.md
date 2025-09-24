# \VPCRouteTableAssociationAPI

All URIs are relative to *https://vpc.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListRouteTableAssociations**](VPCRouteTableAssociationAPI.md#ListRouteTableAssociations) | **Get** /api/v1/route-tables/{route_table_id}/associations | List route table associations 
[**UpdateRouteTableAssociation**](VPCRouteTableAssociationAPI.md#UpdateRouteTableAssociation) | **Put** /api/v1/route-tables/{route_table_id}/associations/{association_id} | Update route table association



## ListRouteTableAssociations

> RouteTableAssociationListModel ListRouteTableAssociations(ctx, routeTableId).XAuthToken(xAuthToken).Limit(limit).Offset(offset).Execute()

List route table associations 



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
	routeTableId := "routeTableId_example" // string | 라우팅 테이블의 고유 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional) (default to 20)
	offset := int32(56) // int32 | 조회 시작 위치 (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPCRouteTableAssociationAPI.ListRouteTableAssociations(context.Background(), routeTableId).XAuthToken(xAuthToken).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCRouteTableAssociationAPI.ListRouteTableAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRouteTableAssociations`: RouteTableAssociationListModel
	fmt.Fprintf(os.Stdout, "Response from `VPCRouteTableAssociationAPI.ListRouteTableAssociations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeTableId** | **string** | 라우팅 테이블의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRouteTableAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | [default to 20]
 **offset** | **int32** | 조회 시작 위치 | [default to 0]

### Return type

[**RouteTableAssociationListModel**](RouteTableAssociationListModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRouteTableAssociation

> ResponseRouteTableAssociationModel UpdateRouteTableAssociation(ctx, routeTableId, associationId).XAuthToken(xAuthToken).BodyUpdateRouteTableAssociation(bodyUpdateRouteTableAssociation).Execute()

Update route table association



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
	routeTableId := "routeTableId_example" // string | 현재 연결에 매핑된 라우팅 테이블 ID
	associationId := "associationId_example" // string | 수정할 라우팅 테이블과 서브넷 연결 정보를 식별하는 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	bodyUpdateRouteTableAssociation := *openapiclient.NewBodyUpdateRouteTableAssociation(*openapiclient.NewEditAssociationModel("TargetRouteTableId_example")) // BodyUpdateRouteTableAssociation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPCRouteTableAssociationAPI.UpdateRouteTableAssociation(context.Background(), routeTableId, associationId).XAuthToken(xAuthToken).BodyUpdateRouteTableAssociation(bodyUpdateRouteTableAssociation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCRouteTableAssociationAPI.UpdateRouteTableAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRouteTableAssociation`: ResponseRouteTableAssociationModel
	fmt.Fprintf(os.Stdout, "Response from `VPCRouteTableAssociationAPI.UpdateRouteTableAssociation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeTableId** | **string** | 현재 연결에 매핑된 라우팅 테이블 ID | 
**associationId** | **string** | 수정할 라우팅 테이블과 서브넷 연결 정보를 식별하는 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRouteTableAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyUpdateRouteTableAssociation** | [**BodyUpdateRouteTableAssociation**](BodyUpdateRouteTableAssociation.md) |  | 

### Return type

[**ResponseRouteTableAssociationModel**](ResponseRouteTableAssociationModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

