# \VPCRouteTableRouteAPI

All URIs are relative to *https://vpc.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRoute**](VPCRouteTableRouteAPI.md#AddRoute) | **Post** /api/v1/route-tables/{route_table_id}/routes | Add route
[**DeleteRoute**](VPCRouteTableRouteAPI.md#DeleteRoute) | **Delete** /api/v1/route-tables/{route_table_id}/routes/{route_id} | Delete route
[**UpdateRoute**](VPCRouteTableRouteAPI.md#UpdateRoute) | **Put** /api/v1/route-tables/{route_table_id}/routes/{route_id} | Update route



## AddRoute

> BnsVpcV1ApiAddRouteModelResponseRouteModel AddRoute(ctx, routeTableId).XAuthToken(xAuthToken).BodyAddRoute(bodyAddRoute).Execute()

Add route



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
	routeTableId := "routeTableId_example" // string | 라우팅을 추가할 대상 라우팅 테이블의 ID <br/> - [List route tables](https://docs.kakaocloud.com/openapi/bns/vpc/list-route-tables) API 참고
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	bodyAddRoute := *openapiclient.NewBodyAddRoute(*openapiclient.NewCreateRouteModel(openapiclient.RouteTableRouteType("igw"), "TargetId_example", "Destination_example")) // BodyAddRoute | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPCRouteTableRouteAPI.AddRoute(context.Background(), routeTableId).XAuthToken(xAuthToken).BodyAddRoute(bodyAddRoute).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCRouteTableRouteAPI.AddRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddRoute`: BnsVpcV1ApiAddRouteModelResponseRouteModel
	fmt.Fprintf(os.Stdout, "Response from `VPCRouteTableRouteAPI.AddRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeTableId** | **string** | 라우팅을 추가할 대상 라우팅 테이블의 ID &lt;br/&gt; - [List route tables](https://docs.kakaocloud.com/openapi/bns/vpc/list-route-tables) API 참고 | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyAddRoute** | [**BodyAddRoute**](BodyAddRoute.md) |  | 

### Return type

[**BnsVpcV1ApiAddRouteModelResponseRouteModel**](BnsVpcV1ApiAddRouteModelResponseRouteModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRoute

> DeleteRoute(ctx, routeTableId, routeId).XAuthToken(xAuthToken).Execute()

Delete route



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
	routeId := "routeId_example" // string | 삭제할 라우팅 ID  
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VPCRouteTableRouteAPI.DeleteRoute(context.Background(), routeTableId, routeId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCRouteTableRouteAPI.DeleteRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeTableId** | **string** | 삭제할 라우팅 테이블 ID | 
**routeId** | **string** | 삭제할 라우팅 ID   | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRouteRequest struct via the builder pattern


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


## UpdateRoute

> BnsVpcV1ApiUpdateRouteModelResponseRouteModel UpdateRoute(ctx, routeTableId, routeId).XAuthToken(xAuthToken).BodyUpdateRoute(bodyUpdateRoute).Execute()

Update route



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
	routeId := "routeId_example" // string | 라우팅의 고유 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	bodyUpdateRoute := *openapiclient.NewBodyUpdateRoute(*openapiclient.NewEditRouteModel(openapiclient.RouteTableRouteType("igw"), "TargetId_example", "Destination_example")) // BodyUpdateRoute | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPCRouteTableRouteAPI.UpdateRoute(context.Background(), routeTableId, routeId).XAuthToken(xAuthToken).BodyUpdateRoute(bodyUpdateRoute).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCRouteTableRouteAPI.UpdateRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRoute`: BnsVpcV1ApiUpdateRouteModelResponseRouteModel
	fmt.Fprintf(os.Stdout, "Response from `VPCRouteTableRouteAPI.UpdateRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeTableId** | **string** | 라우팅 테이블의 고유 ID | 
**routeId** | **string** | 라우팅의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyUpdateRoute** | [**BodyUpdateRoute**](BodyUpdateRoute.md) |  | 

### Return type

[**BnsVpcV1ApiUpdateRouteModelResponseRouteModel**](BnsVpcV1ApiUpdateRouteModelResponseRouteModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

